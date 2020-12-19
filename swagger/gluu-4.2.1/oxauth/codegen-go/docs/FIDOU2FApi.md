# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FidoConfiguration**](FIDOU2FApi.md#FidoConfiguration) | **Get** /fido-configuration | Performs FIDO configuration data.
[**GetFidoU2fAuthentication**](FIDOU2FApi.md#GetFidoU2fAuthentication) | **Get** /fido/u2f/authentication | Performs FIDO U2F authentication of end-user.
[**GetFidoU2fRegistration**](FIDOU2FApi.md#GetFidoU2fRegistration) | **Get** /fido/u2f/registration | U2F device registration.
[**PostFidoU2fAuthentication**](FIDOU2FApi.md#PostFidoU2fAuthentication) | **Post** /fido/u2f/authentication | Performs FIDO U2F authentication of end-user.
[**PostFidoU2fRegistration**](FIDOU2FApi.md#PostFidoU2fRegistration) | **Post** /fido/u2f/registration | U2F device registration.

# **FidoConfiguration**
> U2fConfiguration FidoConfiguration(ctx, )
Performs FIDO configuration data.

Performs FIDO configuration data.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**U2fConfiguration**](U2fConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFidoU2fAuthentication**
> AuthenticateRequestMessage GetFidoU2fAuthentication(ctx, username, optional)
Performs FIDO U2F authentication of end-user.

Performs FIDO U2F authentication of end-user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**|  | 
 **optional** | ***FIDOU2FApiGetFidoU2fAuthenticationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FIDOU2FApiGetFidoU2fAuthenticationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyhandle** | **optional.String**|  | 
 **application** | **optional.String**| The application id that the RP would like to assert. | 
 **sessionId** | **optional.String**|  | 

### Return type

[**AuthenticateRequestMessage**](AuthenticateRequestMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFidoU2fRegistration**
> RegisterRequestMessage GetFidoU2fRegistration(ctx, username, optional)
U2F device registration.

U2F device registration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**|  | 
 **optional** | ***FIDOU2FApiGetFidoU2fRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FIDOU2FApiGetFidoU2fRegistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **application** | **optional.String**| The application id that the RP would like to assert. | 
 **sessionId** | **optional.String**|  | 
 **enrollmentCode** | **optional.String**|  | 

### Return type

[**RegisterRequestMessage**](RegisterRequestMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFidoU2fAuthentication**
> AuthenticateStatus PostFidoU2fAuthentication(ctx, optional)
Performs FIDO U2F authentication of end-user.

Performs FIDO U2F authentication of end-user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FIDOU2FApiPostFidoU2fAuthenticationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FIDOU2FApiPostFidoU2fAuthenticationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GetClientTokenParams**](GetClientTokenParams.md)|  | 

### Return type

[**AuthenticateStatus**](AuthenticateStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFidoU2fRegistration**
> RegisterStatus PostFidoU2fRegistration(ctx, optional)
U2F device registration.

U2F device registration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FIDOU2FApiPostFidoU2fRegistrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FIDOU2FApiPostFidoU2fRegistrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RegisterSiteParams**](RegisterSiteParams.md)|  | 

### Return type

[**RegisterStatus**](RegisterStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

