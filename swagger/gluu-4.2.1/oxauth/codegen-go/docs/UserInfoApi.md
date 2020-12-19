# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserinfo**](UserInfoApi.md#GetUserinfo) | **Get** /userinfo | Returns Claims about the authenticated End-User.
[**PostUserinfo**](UserInfoApi.md#PostUserinfo) | **Post** /userinfo | Returns Claims about the authenticated End-User.

# **GetUserinfo**
> UserClaimsDetails GetUserinfo(ctx, accessToken, optional)
Returns Claims about the authenticated End-User.

Returns Claims about the authenticated End-User.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**| OAuth 2.0 Access Token. | 
 **optional** | ***UserInfoApiGetUserinfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserInfoApiGetUserinfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**|  | 

### Return type

[**UserClaimsDetails**](User Claims details.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/jwt, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUserinfo**
> UserClaimsDetails PostUserinfo(ctx, accessToken, optional)
Returns Claims about the authenticated End-User.

Returns Claims about the authenticated End-User.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
 **optional** | ***UserInfoApiPostUserinfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserInfoApiPostUserinfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**| Client Authorization details that contains the access token along with other details. | 

### Return type

[**UserClaimsDetails**](User Claims details.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/jwt, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

