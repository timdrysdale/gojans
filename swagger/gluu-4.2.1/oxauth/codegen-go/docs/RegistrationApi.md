# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BcDeviceRegistration**](RegistrationApi.md#BcDeviceRegistration) | **Post** /bc-deviceRegistration | Performs backchannel device registration.
[**GetRegister**](RegistrationApi.md#GetRegister) | **Get** /register | Get client information for a previously registered client.
[**PostRegister**](RegistrationApi.md#PostRegister) | **Post** /register | Registers new client dynamically.
[**PutRegister**](RegistrationApi.md#PutRegister) | **Put** /register | Updates Client Metadata for a registered client.

# **BcDeviceRegistration**
> BcDeviceRegistration(ctx, idTokenHint, deviceRegistrationToken)
Performs backchannel device registration.

Performs backchannel device registration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idTokenHint** | **string**|  | 
  **deviceRegistrationToken** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegister**
> ClientResponse GetRegister(ctx, clientId, authorization)
Get client information for a previously registered client.

Get client information for a previously registered client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| Client ID that identifies client. | 
  **authorization** | **string**| Authorization header carrying \\\&quot;registration_access_token\\\&quot; issued before as a Bearer token | 

### Return type

[**ClientResponse**](ClientResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRegister**
> RegisterResponseParam PostRegister(ctx, optional)
Registers new client dynamically.

The Client Registration Endpoint is an OAuth 2.0 Protected Resource through which a new Client registration can be requested.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegistrationApiPostRegisterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegistrationApiPostRegisterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RegisterParams1**](RegisterParams1.md)|  | 

### Return type

[**RegisterResponseParam**](RegisterResponseParam.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutRegister**
> RegisterResponseParam PutRegister(ctx, clientId, authorization, optional)
Updates Client Metadata for a registered client.

Updates Client Metadata for a registered client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| Client ID that identifies client that must be updated by this request. | 
  **authorization** | **string**| Authorization header carrying \\\&quot;registration_access_token\\\&quot; issued before as a Bearer token | 
 **optional** | ***RegistrationApiPutRegisterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegistrationApiPutRegisterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of RegisterParams**](RegisterParams.md)|  | 

### Return type

[**RegisterResponseParam**](RegisterResponseParam.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

