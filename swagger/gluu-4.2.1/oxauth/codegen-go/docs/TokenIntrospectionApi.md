# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRptStatus**](TokenIntrospectionApi.md#GetRptStatus) | **Get** /rpt/status | The Introspection OAuth 2 Endpoint for RPT.
[**PostRptStatus**](TokenIntrospectionApi.md#PostRptStatus) | **Post** /rpt/status | The Introspection OAuth 2 Endpoint for RPT.

# **GetRptStatus**
> RptIntrospectionResponse GetRptStatus(ctx, authorization, token, optional)
The Introspection OAuth 2 Endpoint for RPT.

The Introspection OAuth 2 Endpoint for RPT.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **token** | **string**|  | 
 **optional** | ***TokenIntrospectionApiGetRptStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenIntrospectionApiGetRptStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenTypeHint** | **optional.String**|  | 

### Return type

[**RptIntrospectionResponse**](RptIntrospectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRptStatus**
> RptIntrospectionResponse1 PostRptStatus(ctx, authorization, optional)
The Introspection OAuth 2 Endpoint for RPT.

The Introspection OAuth 2 Endpoint for RPT.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Client Authorization details that contains the access token along with other details. | 
 **optional** | ***TokenIntrospectionApiPostRptStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenIntrospectionApiPostRptStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **optional.**|  | 
 **tokenTypeHint** | **optional.**|  | 

### Return type

[**RptIntrospectionResponse1**](RptIntrospectionResponse_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

