package sql

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// DatabaseBlobAuditingPoliciesClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type DatabaseBlobAuditingPoliciesClient struct {
	ManagementClient
}

// NewDatabaseBlobAuditingPoliciesClient creates an instance of the DatabaseBlobAuditingPoliciesClient client.
func NewDatabaseBlobAuditingPoliciesClient(subscriptionID string) DatabaseBlobAuditingPoliciesClient {
	return NewDatabaseBlobAuditingPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabaseBlobAuditingPoliciesClientWithBaseURI creates an instance of the DatabaseBlobAuditingPoliciesClient
// client.
func NewDatabaseBlobAuditingPoliciesClientWithBaseURI(baseURI string, subscriptionID string) DatabaseBlobAuditingPoliciesClient {
	return DatabaseBlobAuditingPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a database's blob auditing policy.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of the
// database for which the blob auditing policy will be defined. blobAuditingPolicyName is the name of the blob auditing
// policy. parameters is the database blob auditing policy.
func (client DatabaseBlobAuditingPoliciesClient) CreateOrUpdate(resourceGroupName string, serverName string, databaseName string, blobAuditingPolicyName string, parameters DatabaseBlobAuditingPolicy) (result DatabaseBlobAuditingPolicy, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, serverName, databaseName, blobAuditingPolicyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DatabaseBlobAuditingPoliciesClient) CreateOrUpdatePreparer(resourceGroupName string, serverName string, databaseName string, blobAuditingPolicyName string, parameters DatabaseBlobAuditingPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blobAuditingPolicyName": autorest.Encode("path", blobAuditingPolicyName),
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/auditingSettings/{blobAuditingPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseBlobAuditingPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DatabaseBlobAuditingPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result DatabaseBlobAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a database's blob auditing policy.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of the
// database for which the blob audit policy is defined. blobAuditingPolicyName is the name of the blob auditing policy.
func (client DatabaseBlobAuditingPoliciesClient) Get(resourceGroupName string, serverName string, databaseName string, blobAuditingPolicyName string) (result DatabaseBlobAuditingPolicy, err error) {
	req, err := client.GetPreparer(resourceGroupName, serverName, databaseName, blobAuditingPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseBlobAuditingPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DatabaseBlobAuditingPoliciesClient) GetPreparer(resourceGroupName string, serverName string, databaseName string, blobAuditingPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blobAuditingPolicyName": autorest.Encode("path", blobAuditingPolicyName),
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/auditingSettings/{blobAuditingPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseBlobAuditingPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DatabaseBlobAuditingPoliciesClient) GetResponder(resp *http.Response) (result DatabaseBlobAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
