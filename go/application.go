// Package svix this file is @generated DO NOT EDIT
package svix

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/svix/svix-webhooks/go/models"
)

type Application struct {
	client *SvixHttpClient
}

type ApplicationListOptions struct {
	// Limit the number of returned items
	Limit *uint64
	// The iterator returned from a prior invocation
	Iterator *string

	// The sorting order of the returned items
	Order *models.Ordering
}

type ApplicationCreateOptions struct {
	IdempotencyKey *string
}

// List of all the organization's applications.
func (application *Application) List(
	ctx context.Context,
	o *ApplicationListOptions,
) (*models.ListResponseApplicationOut, error) {
	pathMap := map[string]string{}
	queryMap := map[string]string{}
	headerMap := map[string]string{}
	var jsonBody []byte

	var err error
	if o != nil {
		serializeParamToMap("limit", o.Limit, queryMap, &err)
		serializeParamToMap("iterator", o.Iterator, queryMap, &err)
		serializeParamToMap("order", o.Order, queryMap, &err)
		if err != nil {
			return nil, err
		}
	}
	ret, err := executeRequest[models.ListResponseApplicationOut](
		ctx,
		application.client,
		"GET",
		"/api/v1/app",
		pathMap,
		queryMap,
		headerMap,
		jsonBody,
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Create a new application.
func (application *Application) Create(
	ctx context.Context,
	applicationIn *models.ApplicationIn,
	o *ApplicationCreateOptions,
) (*models.ApplicationOut, error) {
	if applicationIn == nil {
		return nil, fmt.Errorf("Application.Create(), applicationIn must not be nil")
	}
	pathMap := map[string]string{}
	queryMap := map[string]string{
		"get_if_exists": "false",
	}
	headerMap := map[string]string{}
	var jsonBody []byte

	var err error
	if o != nil {
		serializeParamToMap("idempotency-key", o.IdempotencyKey, headerMap, &err)
		if err != nil {
			return nil, err
		}
	}
	jsonBody, err = json.Marshal(applicationIn)
	if err != nil {
		return nil, err
	}
	ret, err := executeRequest[models.ApplicationOut](
		ctx,
		application.client,
		"POST",
		"/api/v1/app",
		pathMap,
		queryMap,
		headerMap,
		jsonBody,
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Get or create a new application.
func (application *Application) GetOrCreate(
	ctx context.Context,
	applicationIn *models.ApplicationIn,
	o *ApplicationCreateOptions,
) (*models.ApplicationOut, error) {
	if applicationIn == nil {
		return nil, fmt.Errorf("Application.GetOrCreate(), applicationIn must not be nil")
	}
	pathMap := map[string]string{}
	queryMap := map[string]string{
		"get_if_exists": "true",
	}
	headerMap := map[string]string{}
	var jsonBody []byte

	var err error
	if o != nil {
		serializeParamToMap("idempotency-key", o.IdempotencyKey, headerMap, &err)
		if err != nil {
			return nil, err
		}
	}
	jsonBody, err = json.Marshal(applicationIn)
	if err != nil {
		return nil, err
	}
	ret, err := executeRequest[models.ApplicationOut](
		ctx,
		application.client,
		"POST",
		"/api/v1/app",
		pathMap,
		queryMap,
		headerMap,
		jsonBody,
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Get an application.
func (application *Application) Get(
	ctx context.Context,
	appId string,
) (*models.ApplicationOut, error) {
	pathMap := map[string]string{
		"app_id": appId,
	}
	queryMap := map[string]string{}
	headerMap := map[string]string{}
	var jsonBody []byte

	ret, err := executeRequest[models.ApplicationOut](
		ctx,
		application.client,
		"GET",
		"/api/v1/app/{app_id}",
		pathMap,
		queryMap,
		headerMap,
		jsonBody,
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Update an application.
func (application *Application) Update(
	ctx context.Context,
	appId string,
	applicationIn *models.ApplicationIn,
) (*models.ApplicationOut, error) {
	if applicationIn == nil {
		return nil, fmt.Errorf("Application.Update(), applicationIn must not be nil")
	}
	pathMap := map[string]string{
		"app_id": appId,
	}
	queryMap := map[string]string{}
	headerMap := map[string]string{}
	var jsonBody []byte

	jsonBody, err := json.Marshal(applicationIn)
	if err != nil {
		return nil, err
	}
	ret, err := executeRequest[models.ApplicationOut](
		ctx,
		application.client,
		"PUT",
		"/api/v1/app/{app_id}",
		pathMap,
		queryMap,
		headerMap,
		jsonBody,
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Delete an application.
func (application *Application) Delete(
	ctx context.Context,
	appId string,
) error {
	pathMap := map[string]string{
		"app_id": appId,
	}
	queryMap := map[string]string{}
	headerMap := map[string]string{}
	var jsonBody []byte

	_, err := executeRequest[any](
		ctx,
		application.client,
		"DELETE",
		"/api/v1/app/{app_id}",
		pathMap,
		queryMap,
		headerMap,
		jsonBody,
	)
	if err != nil {
		return err
	}
	return nil
}

// Partially update an application.
func (application *Application) Patch(
	ctx context.Context,
	appId string,
	applicationPatch *models.ApplicationPatch,
) (*models.ApplicationOut, error) {
	if applicationPatch == nil {
		return nil, fmt.Errorf("Application.Patch(), applicationPatch must not be nil")
	}
	pathMap := map[string]string{
		"app_id": appId,
	}
	queryMap := map[string]string{}
	headerMap := map[string]string{}
	var jsonBody []byte

	jsonBody, err := json.Marshal(applicationPatch)
	if err != nil {
		return nil, err
	}
	ret, err := executeRequest[models.ApplicationOut](
		ctx,
		application.client,
		"PATCH",
		"/api/v1/app/{app_id}",
		pathMap,
		queryMap,
		headerMap,
		jsonBody,
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
