# {{classname}}

All URIs are relative to *https://api4.thetvdb.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserFavorites**](FavoritesApi.md#CreateUserFavorites) | **Post** /user/favorites | 
[**GetUserFavorites**](FavoritesApi.md#GetUserFavorites) | **Get** /user/favorites | 

# **CreateUserFavorites**
> CreateUserFavorites(ctx, optional)


creates a new user favorite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoritesApiCreateUserFavoritesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoritesApiCreateUserFavoritesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FavoriteRecord**](FavoriteRecord.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserFavorites**
> InlineResponse20054 GetUserFavorites(ctx, )


returns user favorites

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20054**](inline_response_200_54.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

