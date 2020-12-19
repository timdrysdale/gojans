# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRegister**](ClientRegistrationApi.md#DeleteRegister) | **Delete** /register | Deletes the client info for a previously registered client.

# **DeleteRegister**
> DeleteRegister(ctx, clientId, authorization)
Deletes the client info for a previously registered client.

The Client Registration Endpoint removes the Client Metadata for a previously registered client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| Client ID that identifies client. | 
  **authorization** | **string**| Authorization header carrying \\\&quot;registration_access_token\\\&quot; issued before as a Bearer token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

