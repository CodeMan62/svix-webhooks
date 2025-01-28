# this file is @generated
import typing as t

from pydantic import Field

from .common import SvixBaseModel


class EndpointIn(SvixBaseModel):
    channels: t.Optional[t.List[str]] = None
    """List of message channels this endpoint listens to (omit for all)."""

    description: t.Optional[str] = None

    disabled: t.Optional[bool] = None

    filter_types: t.Optional[t.List[str]] = Field(default=None, alias="filterTypes")

    metadata: t.Optional[t.Dict[str, str]] = None

    rate_limit: t.Optional[int] = Field(default=None, alias="rateLimit")

    secret: t.Optional[str] = None
    """The endpoint's verification secret.

    Format: `base64` encoded random bytes optionally prefixed with `whsec_`.
    It is recommended to not set this and let the server generate the secret."""

    uid: t.Optional[str] = None
    """Optional unique identifier for the endpoint."""

    url: str

    version: t.Optional[int] = None
