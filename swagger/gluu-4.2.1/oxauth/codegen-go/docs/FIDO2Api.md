# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttestationOptions**](FIDO2Api.md#AttestationOptions) | **Post** /fido2/attestation/options | Created new registration.
[**AttestationResult**](FIDO2Api.md#AttestationResult) | **Post** /fido2/attestation/result | FIDO2 attestation result.
[**GetFido2Configuration**](FIDO2Api.md#GetFido2Configuration) | **Get** /fido2/configuration | FIDO2 configuration
[**Options**](FIDO2Api.md#Options) | **Post** /fido2/assertion/options | FIDO2 Assertion Options
[**Result**](FIDO2Api.md#Result) | **Post** /fido2/assertion/result | FIDO2 Assertion Result - Parses and validates an assertion response from the client.

# **AttestationOptions**
> CredentialCreationOptions AttestationOptions(ctx, optional)
Created new registration.

Created new registration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FIDO2ApiAttestationOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FIDO2ApiAttestationOptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AttestationOptions**](AttestationOptions.md)|  | 

### Return type

[**CredentialCreationOptions**](CredentialCreationOptions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttestationResult**
> ErrorResponse AttestationResult(ctx, optional)
FIDO2 attestation result.

FIDO2 attestation result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FIDO2ApiAttestationResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FIDO2ApiAttestationResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AttestationOptions1**](AttestationOptions1.md)|  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFido2Configuration**
> Fido2Configuration GetFido2Configuration(ctx, )
FIDO2 configuration

FIDO2 configuration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Fido2Configuration**](FIDO2 configuration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Options**
> AssertionOptionsResponse Options(ctx, optional)
FIDO2 Assertion Options

FIDO2 Assertion Options

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FIDO2ApiOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FIDO2ApiOptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AssertionOptions**](AssertionOptions.md)|  | 

### Return type

[**AssertionOptionsResponse**](AssertionOptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Result**
> Result(ctx, optional)
FIDO2 Assertion Result - Parses and validates an assertion response from the client.

FIDO2 Assertion Result.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FIDO2ApiResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FIDO2ApiResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AssertionOptions1**](AssertionOptions1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

