# {{classname}}

All URIs are relative to *https://api4.thetvdb.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserInfo**](UserInfoApi.md#GetUserInfo) | **Get** /user | 
[**GetUserInfoById**](UserInfoApi.md#GetUserInfoById) | **Get** /user/{id} | 

# **GetUserInfo**
> InlineResponse20053 GetUserInfo(ctx, )


returns user info

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserInfoById**
> InlineResponse20053 GetUserInfoById(ctx, id)


returns user info by user id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

