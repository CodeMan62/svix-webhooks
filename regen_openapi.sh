#!/bin/bash

set -eo pipefail

OPENAPI_GIT_REV='98dbc5b090a5c8d72fe50962ee04b46fb9d7db20'

if [ -n "$1" ]; then
    curl "$1" | python -m json.tool > lib-openapi.json
fi

# Rust - using the new openapi-codegen
if ! command -v openapi-codegen >/dev/null; then
    if [[ -z "$GITHUB_WORKFLOW" ]]; then
        echo openapi-codegen is not installed. install using
        echo "cargo install --git https://github.com/svix/openapi-codegen --rev $OPENAPI_GIT_REV --locked"
        exit 1
    else
        cargo install --git https://github.com/svix/openapi-codegen --rev $OPENAPI_GIT_REV --locked
    fi
fi

(
    # Print commands we run
    set -x

    openapi-codegen generate \
        --template rust/templates/api_resource.rs.jinja \
        --input-file lib-openapi.json \
        --output-dir rust/src/api
    openapi-codegen generate \
        --template rust/templates/component_type.rs.jinja \
        --input-file lib-openapi.json \
        --output-dir rust/src/models

    # Remove APIs we may not (yet) want to expose
    rm rust/src/api/{environment,health}.rs

    # Remove .codegen.json files, its purpose is fulfilled already:
    # - The expected git rev of the tool is encoded at the top of this file.
    # - The lib-openapi.json used as input is also committed to this repo.
    rm rust/src/{api,models}/.codegen.json
)

cd $(dirname "$0")
mkdir -p .codegen-tmp
# OpenAPI version has to be overwritten to avoid broken codegen paths in OpenAPI generator.
# Spec version is overwritten to avoid unnecessary diffs on comments. Same for description.
# `additionalProperties: true` is removed because OpenAPI generator can't deal with it.
jq --indent 4 '
    .openapi = "3.0.2" |
    .info.version = "1.1.1" |
    del(.info.description) |
    .components.schemas |= with_entries(
        if .value | has("properties") then
            .value.properties |= with_entries(
                if .value.additionalProperties == true then
                    del(.value.additionalProperties)
                else .
                end
            )
        else .
        end
    )' \
    < lib-openapi.json \
    > .codegen-tmp/openapi.json

# For some languages, write a separate OpenAPI spec file where optional fields
# of non-`Patch` schemas are set to not be nullable, so the codegen doesn't wrap
# the struct fields in double options.
# Rust's serde and Go's encoding/json will respect both in deserialization
# anyways, but this frees users from having to think about the difference
# between absence of the field or setting it to `null` (which is only
# significant for patch request bodies).
jq --indent 4 '.components.schemas |= with_entries(
        if .key | endswith("Patch") then .
        else
            (.value.required // []) as $required |
            if .value | has("properties") then
                .value.properties |= with_entries(
                    if .key | IN($required[]) then .
                    else del(.value.nullable)
                    end
                )
            else .
            end
        end
    )' \
    < .codegen-tmp/openapi.json \
    > .codegen-tmp/openapi-less-null.json

# Print commands we run
set -x

yarn openapi-generator-cli generate -i .codegen-tmp/openapi.json -g ruby -o ruby -c ruby/openapi-generator-config.json -t ruby/templates

rm -rf .codegen-tmp
echo Note: Python generation is not executed automatically.
