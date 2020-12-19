# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUmaScopes**](UMAScopeApi.md#GetUmaScopes) | **Get** /uma/scopes | Provides UMA Scope description by scope id.

# **GetUmaScopes**
> UmaScopeDescription GetUmaScopes(ctx, id)
Provides UMA Scope description by scope id.

Provides UMA Scope description by scope id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Resource description ID. | 

### Return type

[**UmaScopeDescription**](UmaScopeDescription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

