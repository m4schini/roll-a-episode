# {{classname}}

All URIs are relative to *https://api4.thetvdb.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllSeasons**](SeasonsApi.md#GetAllSeasons) | **Get** /seasons | 
[**GetSeasonBase**](SeasonsApi.md#GetSeasonBase) | **Get** /seasons/{id} | 
[**GetSeasonExtended**](SeasonsApi.md#GetSeasonExtended) | **Get** /seasons/{id}/extended | 
[**GetSeasonTranslation**](SeasonsApi.md#GetSeasonTranslation) | **Get** /seasons/{id}/translations/{language} | 
[**GetSeasonTypes**](SeasonsApi.md#GetSeasonTypes) | **Get** /seasons/types | 

# **GetAllSeasons**
> InlineResponse20041 GetAllSeasons(ctx, optional)


returns list of seasons base records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SeasonsApiGetAllSeasonsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeasonsApiGetAllSeasonsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| page number | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeasonBase**
> InlineResponse20042 GetSeasonBase(ctx, id)


Returns season base record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeasonExtended**
> InlineResponse20043 GetSeasonExtended(ctx, id)


Returns season extended record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeasonTranslation**
> InlineResponse20020 GetSeasonTranslation(ctx, id, language)


Returns season translation record

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

# **GetSeasonTypes**
> InlineResponse20044 GetSeasonTypes(ctx, )


Returns season type records

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

