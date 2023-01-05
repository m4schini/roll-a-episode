# {{classname}}

All URIs are relative to *https://api4.thetvdb.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllPeople**](PeopleApi.md#GetAllPeople) | **Get** /people | 
[**GetPeopleBase**](PeopleApi.md#GetPeopleBase) | **Get** /people/{id} | 
[**GetPeopleExtended**](PeopleApi.md#GetPeopleExtended) | **Get** /people/{id}/extended | 
[**GetPeopleTranslation**](PeopleApi.md#GetPeopleTranslation) | **Get** /people/{id}/translations/{language} | 

# **GetAllPeople**
> InlineResponse20035 GetAllPeople(ctx, optional)


Returns a list of people base records with the basic attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PeopleApiGetAllPeopleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeopleApiGetAllPeopleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| page number | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPeopleBase**
> InlineResponse20036 GetPeopleBase(ctx, id)


Returns people base record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPeopleExtended**
> InlineResponse20037 GetPeopleExtended(ctx, id, optional)


Returns people extended record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**| id | 
 **optional** | ***PeopleApiGetPeopleExtendedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeopleApiGetPeopleExtendedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **meta** | **optional.String**| meta | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPeopleTranslation**
> InlineResponse20020 GetPeopleTranslation(ctx, id, language)


Returns people translation record

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

