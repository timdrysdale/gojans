# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BcAuthorize**](AuthorizationApi.md#BcAuthorize) | **Post** /bc-authorize | Performs backchannel authorization of the end-user.
[**GetAuthorize**](AuthorizationApi.md#GetAuthorize) | **Get** /authorize | The Authorization Endpoint performs Authentication of the End-User.
[**PostAuthorize**](AuthorizationApi.md#PostAuthorize) | **Post** /authorize | The Authorization Endpoint performs Authentication of the End-User.

# **BcAuthorize**
> BackchannelAuthorization BcAuthorize(ctx, clientId, scope, clientNotificationToken, acrValues, loginHintToken, idTokenHint, loginHint, bindingMessage, userCode, requestedExpiry)
Performs backchannel authorization of the end-user.

The Backchannel Authentication Endpoint is used to initiate an out-of-band authentication of the end-user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**|  | 
  **scope** | **string**|  | 
  **clientNotificationToken** | **string**|  | 
  **acrValues** | **string**|  | 
  **loginHintToken** | **string**|  | 
  **idTokenHint** | **string**|  | 
  **loginHint** | **string**|  | 
  **bindingMessage** | **string**|  | 
  **userCode** | **string**|  | 
  **requestedExpiry** | **int32**|  | 

### Return type

[**BackchannelAuthorization**](Backchannel Authorization.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthorize**
> GetAuthorize(ctx, scope, responseType, clientId, redirectUri, optional)
The Authorization Endpoint performs Authentication of the End-User.

End-User Authentication and Authorization done by sending the User Agent to the Authorization Endpoint using request parameters defined by OAuth 2.0 and OpenID Connect.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scope** | **string**| OpenID Connect requests MUST contain the openid scope value. If the openid scope value is not present, the behavior is entirely unspecified. Other scope values MAY be present. | 
  **responseType** | **string**| OAuth 2.0 Response Type value that determines the authorization processing flow to be used, including what parameters are returned from the endpoints used. | 
  **clientId** | **string**| OAuth 2.0 Client Identifier valid at the Authorization Server. | 
  **redirectUri** | **string**| Redirection URI to which the response will be sent. This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider. | 
 **optional** | ***AuthorizationApiGetAuthorizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiGetAuthorizeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **state** | **optional.String**| Opaque value used to maintain state between the request and the callback. | 
 **responseMode** | **optional.String**| Informs the Authorization Server of the mechanism to be used for returning parameters from the Authorization Endpoint. | 
 **nonce** | **optional.String**| String value used to associate a Client session with an ID Token, and to mitigate replay attacks. | 
 **display** | **optional.String**| ASCII string value that specifies how the Authorization Server displays the authentication and consent user interface pages to the End-User. | 
 **prompt** | **optional.String**| Space delimited, case sensitive list of ASCII string values that specifies whether the Authorization Server prompts the End-User for reauthentication and consent. The defined values are - none, login, consent, select_account. | 
 **maxAge** | **optional.Int32**| Maximum Authentication Age. Specifies the allowable elapsed time in seconds since the last time the End-User was actively authenticated by the OP. | 
 **uiLocales** | **optional.String**| End-User&#x27;s preferred languages and scripts for the user interface, represented as a space-separated list of BCP47 [RFC5646] language tag values, ordered by preference. | 
 **idTokenHint** | **optional.String**| ID Token previously issued by the Authorization Server being passed as a hint about the End-User&#x27;s current or past authenticated session with the Client. If the End-User identified by the ID Token is logged in or is logged in by the request, then the Authorization Server returns a positive response. | 
 **loginHint** | **optional.String**| Hint to the Authorization Server about the login identifier the End-User might use to log in (if necessary). | 
 **acrValues** | **optional.String**| Requested Authentication Context Class Reference values. Space-separated string that specifies the acr values that the Authorization Server is being requested to use for processing this Authentication Request, with the values appearing in order of preference. | 
 **amrValues** | **optional.String**| AMR Values. | 
 **request** | **optional.String**| This parameter enables OpenID Connect requests to be passed in a single, self-contained parameter and to be optionally signed and/or encrypted. The parameter value is a Request Object value. It represents the request as a JWT whose Claims are the request parameters. | 
 **requestUri** | **optional.String**| This parameter enables OpenID Connect requests to be passed by reference, rather than by value. The request_uri value is a URL using the https scheme referencing a resource containing a Request Object value, which is a JWT containing the request parameters. | 
 **requestSessionId** | **optional.String**| Request session id. | 
 **sessionId** | **optional.String**| Session id of this call. | 
 **originHeaders** | **optional.String**| Origin headers. Used in custom workflows. | 
 **codeChallenge** | **optional.String**| PKCE code challenge. | 
 **codeChallengeMethod** | **optional.String**| PKCE code challenge method. | 
 **customResponseHeaders** | **optional.String**| Custom Response Headers. | 
 **claims** | **optional.String**| Requested Claims. | 
 **authReqId** | **optional.String**| CIBA authentication request Id. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAuthorize**
> PostAuthorize(ctx, scope, responseType, clientId, redirectUri, state, responseMode, nonce, display, prompt, maxAge, uiLocales, idTokenHint, loginHint, acrValues, amrValues, request, requestUri, requestSessionId, sessionId, originHeaders, codeChallenge, codeChallengeMethod, customResponseHeaders, claims)
The Authorization Endpoint performs Authentication of the End-User.

End-User Authentication and Authorization done by sending the User Agent to the Authorization Endpoint using request parameters defined by OAuth 2.0 and OpenID Connect.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scope** | **string**|  | 
  **responseType** | **string**|  | 
  **clientId** | **string**|  | 
  **redirectUri** | **string**|  | 
  **state** | **string**|  | 
  **responseMode** | **string**|  | 
  **nonce** | **string**|  | 
  **display** | **string**|  | 
  **prompt** | **string**|  | 
  **maxAge** | **int32**|  | 
  **uiLocales** | **string**|  | 
  **idTokenHint** | **string**|  | 
  **loginHint** | **string**|  | 
  **acrValues** | **string**|  | 
  **amrValues** | **string**|  | 
  **request** | **string**|  | 
  **requestUri** | **string**|  | 
  **requestSessionId** | **string**|  | 
  **sessionId** | **string**|  | 
  **originHeaders** | **string**|  | 
  **codeChallenge** | **string**|  | 
  **codeChallengeMethod** | **string**|  | 
  **customResponseHeaders** | **string**|  | 
  **claims** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

