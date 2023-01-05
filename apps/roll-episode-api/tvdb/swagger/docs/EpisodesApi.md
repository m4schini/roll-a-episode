# {{classname}}

All URIs are relative to *https://api4.thetvdb.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllEpisodes**](EpisodesApi.md#GetAllEpisodes) | **Get** /episodes | 
[**GetEpisodeBase**](EpisodesApi.md#GetEpisodeBase) | **Get** /episodes/{id} | 
[**GetEpisodeExtended**](EpisodesApi.md#GetEpisodeExtended) | **Get** /episodes/{id}/extended | 
[**GetEpisodeTranslation**](EpisodesApi.md#GetEpisodeTranslation) | **Get** /episodes/{id}/translations/{language} | 

# **GetAllEpisodes**
> InlineResponse20017 GetAllEpisodes(ctx, optional)


Returns a list of episodes base records with the basic attributes.<br> Note that all episodes are returned, even those that may not be included in a series' default season order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EpisodesApiGetAllEpisodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EpisodesApiGetAllEpisodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| page number | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEpisodeBase**
> InlineResponse20018 GetEpisodeBase(ctx, id)


Returns episode base record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEpisodeExtended**
> InlineResponse20019 GetEpisodeExtended(ctx, id, optional)


Returns episode extended record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 
 **optional** | ***EpisodesApiGetEpisodeExtendedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EpisodesApiGetEpisodeExtendedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **meta** | **optional.String**| meta | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEpisodeTranslation**
> InlineResponse20020 GetEpisodeTranslation(ctx, id, language)


Returns episode translation record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 
  **language** | **string**| language | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

