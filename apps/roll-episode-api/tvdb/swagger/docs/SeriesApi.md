# {{classname}}

All URIs are relative to *https://api4.thetvdb.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllSeries**](SeriesApi.md#GetAllSeries) | **Get** /series | 
[**GetSeriesArtworks**](SeriesApi.md#GetSeriesArtworks) | **Get** /series/{id}/artworks | 
[**GetSeriesBase**](SeriesApi.md#GetSeriesBase) | **Get** /series/{id} | 
[**GetSeriesBaseBySlug**](SeriesApi.md#GetSeriesBaseBySlug) | **Get** /series/slug/{slug} | 
[**GetSeriesEpisodes**](SeriesApi.md#GetSeriesEpisodes) | **Get** /series/{id}/episodes/{season-type} | 
[**GetSeriesExtended**](SeriesApi.md#GetSeriesExtended) | **Get** /series/{id}/extended | 
[**GetSeriesFilter**](SeriesApi.md#GetSeriesFilter) | **Get** /series/filter | 
[**GetSeriesNextAired**](SeriesApi.md#GetSeriesNextAired) | **Get** /series/{id}/nextAired | 
[**GetSeriesSeasonEpisodesTranslated**](SeriesApi.md#GetSeriesSeasonEpisodesTranslated) | **Get** /series/{id}/episodes/{season-type}/{lang} | 
[**GetSeriesTranslation**](SeriesApi.md#GetSeriesTranslation) | **Get** /series/{id}/translations/{language} | 

# **GetAllSeries**
> InlineResponse20045 GetAllSeries(ctx, optional)


returns list of series base records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SeriesApiGetAllSeriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeriesApiGetAllSeriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| page number | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeriesArtworks**
> InlineResponse20047 GetSeriesArtworks(ctx, id, optional)


Returns series artworks base on language and type. <br> Note&#58; Artwork type is an id that can be found using **_/artwork/types** endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 
 **optional** | ***SeriesApiGetSeriesArtworksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeriesApiGetSeriesArtworksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lang** | **optional.String**| lang | 
 **type_** | **optional.Int32**| type | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeriesBase**
> InlineResponse20046 GetSeriesBase(ctx, id)


Returns series base record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeriesBaseBySlug**
> InlineResponse20046 GetSeriesBaseBySlug(ctx, slug)


Returns series base record searched by slug

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slug** | **string**| slug | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeriesEpisodes**
> InlineResponse20048 GetSeriesEpisodes(ctx, page, id, seasonType, optional)


Returns series episodes from the specified season type, default returns the episodes in the series default season type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**|  | [default to 0]
  **id** | **float64**| id | 
  **seasonType** | **string**| season-type | 
 **optional** | ***SeriesApiGetSeriesEpisodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeriesApiGetSeriesEpisodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **season** | **optional.Int32**|  | [default to 0]
 **episodeNumber** | **optional.Int32**|  | [default to 0]
 **airDate** | **optional.String**| airDate of the episode, format is yyyy-mm-dd | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeriesExtended**
> InlineResponse20047 GetSeriesExtended(ctx, id, optional)


Returns series extended record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 
 **optional** | ***SeriesApiGetSeriesExtendedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeriesApiGetSeriesExtendedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **meta** | **optional.String**| meta | 
 **short** | **optional.Bool**| reduce the payload and returns the short version of this record without characters and artworks | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeriesFilter**
> InlineResponse20050 GetSeriesFilter(ctx, country, lang, optional)


Search series based on filter parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| country of origin | 
  **lang** | **string**| original language | 
 **optional** | ***SeriesApiGetSeriesFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SeriesApiGetSeriesFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **company** | **optional.Float64**| production company | 
 **contentRating** | **optional.Float64**| content rating id base on a country | 
 **genre** | **optional.Float64**| Genre id. This id can be found using **_/genres** endpoint. | 
 **sort** | **optional.String**| sort by results | 
 **sortType** | **optional.String**| sort type ascending or descending | 
 **status** | **optional.Float64**| status | 
 **year** | **optional.Float64**| release year | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeriesNextAired**
> InlineResponse20046 GetSeriesNextAired(ctx, id)


Returns series base record including the nextAired field. <br> Note&#58; nextAired was included in the base record endpoint but that field will deprecated in the future so developers should use the nextAired endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeriesSeasonEpisodesTranslated**
> InlineResponse20049 GetSeriesSeasonEpisodesTranslated(ctx, page, id, seasonType, lang)


Returns series base record with episodes from the specified season type and language. Default returns the episodes in the series default season type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**|  | [default to 0]
  **id** | **float64**| id | 
  **seasonType** | **string**| season-type | 
  **lang** | **string**|  | 

### Return type

[**InlineResponse20049**](inline_response_200_49.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSeriesTranslation**
> InlineResponse20020 GetSeriesTranslation(ctx, id, language)


Returns series translation record

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

