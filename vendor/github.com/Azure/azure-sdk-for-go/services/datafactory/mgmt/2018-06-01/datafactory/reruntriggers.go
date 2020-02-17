package datafactory

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RerunTriggersClient is the the Azure Data Factory V2 management API provides a RESTful set of web services that
// interact with Azure Data Factory V2 services.
type RerunTriggersClient struct {
	BaseClient
}

// NewRerunTriggersClient creates an instance of the RerunTriggersClient client.
func NewRerunTriggersClient(subscriptionID string) RerunTriggersClient {
	return NewRerunTriggersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRerunTriggersClientWithBaseURI creates an instance of the RerunTriggersClient client.
func NewRerunTriggersClientWithBaseURI(baseURI string, subscriptionID string) RerunTriggersClient {
	return RerunTriggersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel cancels a trigger.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// triggerName - the trigger name.
// rerunTriggerName - the rerun trigger name.
func (client RerunTriggersClient) Cancel(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string) (result RerunTriggersCancelFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RerunTriggersClient.Cancel")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}},
		{TargetValue: rerunTriggerName,
			Constraints: []validation.Constraint{{Target: "rerunTriggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "rerunTriggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "rerunTriggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.RerunTriggersClient", "Cancel", err.Error())
	}

	req, err := client.CancelPreparer(ctx, resourceGroupName, factoryName, triggerName, rerunTriggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "Cancel", nil, "Failure preparing request")
		return
	}

	result, err = client.CancelSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "Cancel", result.Response(), "Failure sending request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client RerunTriggersClient) CancelPreparer(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"rerunTriggerName":  autorest.Encode("path", rerunTriggerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/rerunTriggers/{rerunTriggerName}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client RerunTriggersClient) CancelSender(req *http.Request) (future RerunTriggersCancelFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client RerunTriggersClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Create creates a rerun trigger.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// triggerName - the trigger name.
// rerunTriggerName - the rerun trigger name.
// rerunTumblingWindowTriggerActionParameters - rerun tumbling window trigger action parameters.
func (client RerunTriggersClient) Create(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string, rerunTumblingWindowTriggerActionParameters RerunTumblingWindowTriggerActionParameters) (result TriggerResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RerunTriggersClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}},
		{TargetValue: rerunTriggerName,
			Constraints: []validation.Constraint{{Target: "rerunTriggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "rerunTriggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "rerunTriggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}},
		{TargetValue: rerunTumblingWindowTriggerActionParameters,
			Constraints: []validation.Constraint{{Target: "rerunTumblingWindowTriggerActionParameters.StartTime", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "rerunTumblingWindowTriggerActionParameters.EndTime", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "rerunTumblingWindowTriggerActionParameters.MaxConcurrency", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "rerunTumblingWindowTriggerActionParameters.MaxConcurrency", Name: validation.InclusiveMaximum, Rule: int64(50), Chain: nil},
						{Target: "rerunTumblingWindowTriggerActionParameters.MaxConcurrency", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("datafactory.RerunTriggersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, factoryName, triggerName, rerunTriggerName, rerunTumblingWindowTriggerActionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client RerunTriggersClient) CreatePreparer(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string, rerunTumblingWindowTriggerActionParameters RerunTumblingWindowTriggerActionParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"rerunTriggerName":  autorest.Encode("path", rerunTriggerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/rerunTriggers/{rerunTriggerName}", pathParameters),
		autorest.WithJSON(rerunTumblingWindowTriggerActionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client RerunTriggersClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client RerunTriggersClient) CreateResponder(resp *http.Response) (result TriggerResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByTrigger lists rerun triggers by an original trigger name.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// triggerName - the trigger name.
func (client RerunTriggersClient) ListByTrigger(ctx context.Context, resourceGroupName string, factoryName string, triggerName string) (result RerunTriggerListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RerunTriggersClient.ListByTrigger")
		defer func() {
			sc := -1
			if result.rtlr.Response.Response != nil {
				sc = result.rtlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.RerunTriggersClient", "ListByTrigger", err.Error())
	}

	result.fn = client.listByTriggerNextResults
	req, err := client.ListByTriggerPreparer(ctx, resourceGroupName, factoryName, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "ListByTrigger", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTriggerSender(req)
	if err != nil {
		result.rtlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "ListByTrigger", resp, "Failure sending request")
		return
	}

	result.rtlr, err = client.ListByTriggerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "ListByTrigger", resp, "Failure responding to request")
	}

	return
}

// ListByTriggerPreparer prepares the ListByTrigger request.
func (client RerunTriggersClient) ListByTriggerPreparer(ctx context.Context, resourceGroupName string, factoryName string, triggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/rerunTriggers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByTriggerSender sends the ListByTrigger request. The method will close the
// http.Response Body if it receives an error.
func (client RerunTriggersClient) ListByTriggerSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByTriggerResponder handles the response to the ListByTrigger request. The method always
// closes the http.Response Body.
func (client RerunTriggersClient) ListByTriggerResponder(resp *http.Response) (result RerunTriggerListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByTriggerNextResults retrieves the next set of results, if any.
func (client RerunTriggersClient) listByTriggerNextResults(ctx context.Context, lastResults RerunTriggerListResponse) (result RerunTriggerListResponse, err error) {
	req, err := lastResults.rerunTriggerListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "listByTriggerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByTriggerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "listByTriggerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByTriggerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "listByTriggerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByTriggerComplete enumerates all values, automatically crossing page boundaries as required.
func (client RerunTriggersClient) ListByTriggerComplete(ctx context.Context, resourceGroupName string, factoryName string, triggerName string) (result RerunTriggerListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RerunTriggersClient.ListByTrigger")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByTrigger(ctx, resourceGroupName, factoryName, triggerName)
	return
}

// Start starts a trigger.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// triggerName - the trigger name.
// rerunTriggerName - the rerun trigger name.
func (client RerunTriggersClient) Start(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string) (result RerunTriggersStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RerunTriggersClient.Start")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}},
		{TargetValue: rerunTriggerName,
			Constraints: []validation.Constraint{{Target: "rerunTriggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "rerunTriggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "rerunTriggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.RerunTriggersClient", "Start", err.Error())
	}

	req, err := client.StartPreparer(ctx, resourceGroupName, factoryName, triggerName, rerunTriggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "Start", result.Response(), "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client RerunTriggersClient) StartPreparer(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"rerunTriggerName":  autorest.Encode("path", rerunTriggerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/rerunTriggers/{rerunTriggerName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client RerunTriggersClient) StartSender(req *http.Request) (future RerunTriggersStartFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client RerunTriggersClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stops a trigger.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// triggerName - the trigger name.
// rerunTriggerName - the rerun trigger name.
func (client RerunTriggersClient) Stop(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string) (result RerunTriggersStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RerunTriggersClient.Stop")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}},
		{TargetValue: rerunTriggerName,
			Constraints: []validation.Constraint{{Target: "rerunTriggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "rerunTriggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "rerunTriggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.RerunTriggersClient", "Stop", err.Error())
	}

	req, err := client.StopPreparer(ctx, resourceGroupName, factoryName, triggerName, rerunTriggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.RerunTriggersClient", "Stop", result.Response(), "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client RerunTriggersClient) StopPreparer(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"rerunTriggerName":  autorest.Encode("path", rerunTriggerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/rerunTriggers/{rerunTriggerName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client RerunTriggersClient) StopSender(req *http.Request) (future RerunTriggersStopFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client RerunTriggersClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
