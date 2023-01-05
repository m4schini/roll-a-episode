# {{classname}}

All URIs are relative to *https://api4.thetvdb.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearchResults**](SearchApi.md#GetSearchResults) | **Get** /search | 
[**GetSearchResultsByRemoteId**](SearchApi.md#GetSearchResultsByRemoteId) | **Get** /search/remoteid/{remoteId} | 

# **GetSearchResults**
> InlineResponse20039 GetSearchResults(ctx, optional)


Our search index includes series, movies, people, and companies. Search is limited to 5k results max.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiGetSearchResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiGetSearchResultsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The primary search string, which can include the main title for a record including all translations and aliases. | 
 **q** | **optional.String**| Alias of the \&quot;query\&quot; parameter.  Recommend using query instead as this field will eventually be deprecated. | 
 **type_** | **optional.String**| Restrict results to a specific entity type.  Can be movie, series, person, or company. | 
 **year** | **optional.Float64**| Restrict results to a specific year. Currently only used for series and movies. | 
 **company** | **optional.String**| Restrict results to a specific company (original network, production company, studio, etc).  As an example, \&quot;The Walking Dead\&quot; would have companies of \&quot;AMC\&quot;, \&quot;AMC+\&quot;, and \&quot;Disney+\&quot;. | 
 **country** | **optional.String**| Restrict results to a specific country of origin. Should contain a 3 character country code. Currently only used for series and movies. | 
 **director** | **optional.String**| Restrict results to a specific director.  Generally only used for movies.  Should include the full name of the director, such as \&quot;Steven Spielberg\&quot;. | 
 **language** | **optional.String**| Restrict results to a specific primary language.  Should include the 3 character language code.  Currently only used for series and movies. | 
 **primaryType** | **optional.String**| Restrict results to a specific type of company.  Should include the full name of the type of company, such as \&quot;Production Company\&quot;.  Only used for companies. | 
 **network** | **optional.String**| Restrict results to a specific network.  Used for TV and TV movies, and functions the same as the company parameter with more specificity. | 
 **remoteId** | **optional.String**| Search for a specific remote id.  Allows searching for an IMDB or EIDR id, for example. | 
 **offset** | **optional.Float64**| Offset results. | 
 **limit** | **optional.Float64**| Limit results. | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSearchResultsByRemoteId**
> InlineResponse20040 GetSearchResultsByRemoteId(ctx, remoteId)


Search a series, movie, people, episode, company or season by specific remote id and returns a base record for that entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **remoteId** | **string**| Search for a specific remote id.  Allows searching for an IMDB or EIDR id, for example. | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

