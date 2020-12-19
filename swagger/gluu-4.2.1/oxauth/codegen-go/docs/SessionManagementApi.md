# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndSession**](SessionManagementApi.md#EndSession) | **Get** /end_session | End current session.
[**RevokeSession**](SessionManagementApi.md#RevokeSession) | **Post** /revoke_session | Revoke all sessions for user.
[**SessionStatus**](SessionManagementApi.md#SessionStatus) | **Get** /session_status | Determine current sesion status.

# **EndSession**
> EndSession(ctx, optional)
End current session.

End current session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SessionManagementApiEndSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SessionManagementApiEndSessionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idTokenHint** | **optional.String**| Previously issued ID Token (id_token) passed to the logout endpoint as a hint about the End-User&#x27;s current authenticated session with the Client. This is used as an indication of the identity of the End-User that the RP is requesting be logged out by the OP. The OP need not be listed as an audience of the ID Token when it is used as an id_token_hint value. | 
 **postLogoutRedirectUri** | **optional.String**| URL to which the RP is requesting that the End-User&#x27;s User Agent be redirected after a logout has been performed. The value MUST have been previously registered with the OP, either using the post_logout_redirect_uris Registration parameter or via another mechanism. If supplied, the OP SHOULD honor this request following the logout. | 
 **state** | **optional.String**| Opaque value used by the RP to maintain state between the logout request and the callback to the endpoint specified by the post_logout_redirect_uri parameter. If included in the logout request, the OP passes this value back to the RP using the state query parameter when redirecting the User Agent back to the RP. | 
 **sessionId** | **optional.String**| Session Id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeSession**
> RevokeSession(ctx, userCriterionKey, userCriterionValue)
Revoke all sessions for user.

Revoke all sessions for user (requires revoke_session scope).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userCriterionKey** | **string**|  | 
  **userCriterionValue** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionStatus**
> SessionStateObject SessionStatus(ctx, )
Determine current sesion status.

Determine current sesion status.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SessionStateObject**](SessionStateObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

