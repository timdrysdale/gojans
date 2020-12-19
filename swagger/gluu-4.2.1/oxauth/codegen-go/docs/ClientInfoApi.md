# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClientinfo**](ClientInfoApi.md#GetClientinfo) | **Get** /clientinfo | To get Claims details about the registered client.
[**PostClientinfo**](ClientInfoApi.md#PostClientinfo) | **Post** /clientinfo | To get Claims details about the registered client.

# **GetClientinfo**
> ClientInfoResponse GetClientinfo(ctx, optional)
To get Claims details about the registered client.

The ClientInfo Endpoint is an OAuth 2.0 Protected Resource that returns Claims about the registered client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientInfoApiGetClientinfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientInfoApiGetClientinfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **optional.String**|  | 
 **authorization** | **optional.String**|  | 

### Return type

[**ClientInfoResponse**](ClientInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostClientinfo**
> ClientInfoResponse PostClientinfo(ctx, accessToken, optional)
To get Claims details about the registered client.

The ClientInfo Endpoint is an OAuth 2.0 Protected Resource that returns Claims about the registered client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***ClientInfoApiPostClientinfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientInfoApiPostClientinfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**|  | 

### Return type

[**ClientInfoResponse**](ClientInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

