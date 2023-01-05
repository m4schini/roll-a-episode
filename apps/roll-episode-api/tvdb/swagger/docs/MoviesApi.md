# {{classname}}

All URIs are relative to *https://api4.thetvdb.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllMovie**](MoviesApi.md#GetAllMovie) | **Get** /movies | 
[**GetMovieBase**](MoviesApi.md#GetMovieBase) | **Get** /movies/{id} | 
[**GetMovieBaseBySlug**](MoviesApi.md#GetMovieBaseBySlug) | **Get** /movies/slug/{slug} | 
[**GetMovieExtended**](MoviesApi.md#GetMovieExtended) | **Get** /movies/{id}/extended | 
[**GetMovieTranslation**](MoviesApi.md#GetMovieTranslation) | **Get** /movies/{id}/translations/{language} | 
[**GetMoviesFilter**](MoviesApi.md#GetMoviesFilter) | **Get** /movies/filter | 

# **GetAllMovie**
> InlineResponse20030 GetAllMovie(ctx, optional)


returns list of movie base records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MoviesApiGetAllMovieOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MoviesApiGetAllMovieOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| page number | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMovieBase**
> InlineResponse20031 GetMovieBase(ctx, id)


Returns movie base record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMovieBaseBySlug**
> InlineResponse20031 GetMovieBaseBySlug(ctx, slug)


Returns movie base record search by slug

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slug** | **string**| slug | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMovieExtended**
> InlineResponse20032 GetMovieExtended(ctx, id, optional)


Returns movie extended record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 
 **optional** | ***MoviesApiGetMovieExtendedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MoviesApiGetMovieExtendedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **meta** | **optional.String**| meta | 
 **short** | **optional.Bool**| reduce the payload and returns the short version of this record without characters, artworks and trailers. | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMovieTranslation**
> InlineResponse20020 GetMovieTranslation(ctx, id, language)


Returns movie translation record

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

# **GetMoviesFilter**
> InlineResponse20033 GetMoviesFilter(ctx, country, lang, optional)


Search movies based on filter parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| country of origin | 
  **lang** | **string**| original language | 
 **optional** | ***MoviesApiGetMoviesFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MoviesApiGetMoviesFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **company** | **optional.Float64**| production company | 
 **contentRating** | **optional.Float64**| content rating id base on a country | 
 **genre** | **optional.Float64**| genre | 
 **sort** | **optional.String**| sort by results | 
 **status** | **optional.Float64**| status | 
 **year** | **optional.Float64**| release year | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

