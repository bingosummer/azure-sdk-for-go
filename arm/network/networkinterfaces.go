package network

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// NetworkInterfaces Client
type NetworkInterfacesClient struct {
	NetworkResourceProviderClient
}

func NewNetworkInterfacesClient(subscriptionId string) NetworkInterfacesClient {
	return NewNetworkInterfacesClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewNetworkInterfacesClientWithBaseUri(baseUri string, subscriptionId string) NetworkInterfacesClient {
	return NetworkInterfacesClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// Delete the delete netwokInterface operation deletes the specified
// netwokInterface.
//
// resourceGroupName is the name of the resource group. networkInterfaceName
// is the name of the network interface.
func (client NetworkInterfacesClient) Delete(resourceGroupName string, networkInterfaceName string) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewDeleteRequest(resourceGroupName, networkInterfaceName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "Delete", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "Delete", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusAccepted))
	if err == nil {
		err = client.IsPollingAllowed(resp)
		if err == nil {
			resp, err = client.PollAsNeeded(resp)
		}
	}

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "Delete", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "Delete", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the Delete request.
func (client NetworkInterfacesClient) NewDeleteRequest(resourceGroupName string, networkInterfaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": url.QueryEscape(networkInterfaceName),
		"resourceGroupName":    url.QueryEscape(resourceGroupName),
		"subscriptionId":       url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.DeleteRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Delete request.
func (client NetworkInterfacesClient) DeleteRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"))
}

// Get the Get ntework interface operation retreives information about the
// specified network interface.
//
// resourceGroupName is the name of the resource group. networkInterfaceName
// is the name of the network interface.
func (client NetworkInterfacesClient) Get(resourceGroupName string, networkInterfaceName string) (result NetworkInterface, ae autorest.Error) {
	req, err := client.NewGetRequest(resourceGroupName, networkInterfaceName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "Get", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "Get", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "Get", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "Get", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Get request.
func (client NetworkInterfacesClient) NewGetRequest(resourceGroupName string, networkInterfaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": url.QueryEscape(networkInterfaceName),
		"resourceGroupName":    url.QueryEscape(resourceGroupName),
		"subscriptionId":       url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.GetRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Get request.
func (client NetworkInterfacesClient) GetRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"))
}

// CreateOrUpdate the Put NetworkInterface operation creates/updates a
// networkInterface
//
// resourceGroupName is the name of the resource group. networkInterfaceName
// is the name of the network interface. parameters is parameters supplied to
// the create/update NetworkInterface operation
func (client NetworkInterfacesClient) CreateOrUpdate(resourceGroupName string, networkInterfaceName string, parameters NetworkInterface) (result NetworkInterface, ae autorest.Error) {
	req, err := client.NewCreateOrUpdateRequest(resourceGroupName, networkInterfaceName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "CreateOrUpdate", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusCreated, http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "CreateOrUpdate", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "CreateOrUpdate", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the CreateOrUpdate request.
func (client NetworkInterfacesClient) NewCreateOrUpdateRequest(resourceGroupName string, networkInterfaceName string, parameters NetworkInterface) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": url.QueryEscape(networkInterfaceName),
		"resourceGroupName":    url.QueryEscape(resourceGroupName),
		"subscriptionId":       url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.CreateOrUpdateRequestPreparer(),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the CreateOrUpdate request.
func (client NetworkInterfacesClient) CreateOrUpdateRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"))
}

// ListAll the List networkInterfaces opertion retrieves all the
// networkInterfaces in a subscription.
func (client NetworkInterfacesClient) ListAll() (result NetworkInterfaceListResult, ae autorest.Error) {
	req, err := client.NewListAllRequest()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "ListAll", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "ListAll", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "ListAll", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "ListAll", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the ListAll request.
func (client NetworkInterfacesClient) NewListAllRequest() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListAllRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the ListAll request.
func (client NetworkInterfacesClient) ListAllRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkInterfaces"))
}

// List the List networkInterfaces opertion retrieves all the
// networkInterfaces in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client NetworkInterfacesClient) List(resourceGroupName string) (result NetworkInterfaceListResult, ae autorest.Error) {
	req, err := client.NewListRequest(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "List", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "List", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "List", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "network.NetworkInterfacesClient", "List", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the List request.
func (client NetworkInterfacesClient) NewListRequest(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the List request.
func (client NetworkInterfacesClient) ListRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces"))
}
