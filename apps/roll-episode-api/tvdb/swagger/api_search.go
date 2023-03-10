/*
 * TVDB API V4
 *
 * Documentation of [TheTVDB](https://thetvdb.com/) API V4. All related information is linked from our [Github repo](https://github.com/thetvdb/v4-api). You might also want to use our [Postman collection] (https://www.getpostman.com/collections/7a9397ce69ff246f74d0) ## Authentication 1. Use the /login endpoint and provide your API key as \"apikey\". If you have a user-supported key, also provide your subscriber PIN as \"pin\". Otherwise completely remove \"pin\" from your call. 2. Executing this call will provide you with a bearer token, which is valid for 1 month. 3. Provide your bearer token for subsequent API calls by clicking Authorize below or including in the header of all direct API calls: `Authorization: Bearer [your-token]`  ## Notes 1. \"score\" is a field across almost all entities.  We generate scores for different types of entities in various ways, so no assumptions should be made about the meaning of this value.  It is simply used to hint at relative popularity for sorting purposes.
 *
 * API version: 4.7.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type SearchApiService service

/*
SearchApiService
Our search index includes series, movies, people, and companies. Search is limited to 5k results max.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *SearchApiGetSearchResultsOpts - Optional Parameters:
     * @param "Query" (optional.String) -  The primary search string, which can include the main title for a record including all translations and aliases.
     * @param "Q" (optional.String) -  Alias of the \&quot;query\&quot; parameter.  Recommend using query instead as this field will eventually be deprecated.
     * @param "Type_" (optional.String) -  Restrict results to a specific entity type.  Can be movie, series, person, or company.
     * @param "Year" (optional.Float64) -  Restrict results to a specific year. Currently only used for series and movies.
     * @param "Company" (optional.String) -  Restrict results to a specific company (original network, production company, studio, etc).  As an example, \&quot;The Walking Dead\&quot; would have companies of \&quot;AMC\&quot;, \&quot;AMC+\&quot;, and \&quot;Disney+\&quot;.
     * @param "Country" (optional.String) -  Restrict results to a specific country of origin. Should contain a 3 character country code. Currently only used for series and movies.
     * @param "Director" (optional.String) -  Restrict results to a specific director.  Generally only used for movies.  Should include the full name of the director, such as \&quot;Steven Spielberg\&quot;.
     * @param "Language" (optional.String) -  Restrict results to a specific primary language.  Should include the 3 character language code.  Currently only used for series and movies.
     * @param "PrimaryType" (optional.String) -  Restrict results to a specific type of company.  Should include the full name of the type of company, such as \&quot;Production Company\&quot;.  Only used for companies.
     * @param "Network" (optional.String) -  Restrict results to a specific network.  Used for TV and TV movies, and functions the same as the company parameter with more specificity.
     * @param "RemoteId" (optional.String) -  Search for a specific remote id.  Allows searching for an IMDB or EIDR id, for example.
     * @param "Offset" (optional.Float64) -  Offset results.
     * @param "Limit" (optional.Float64) -  Limit results.
@return InlineResponse20039
*/

type SearchApiGetSearchResultsOpts struct {
	Query       optional.String
	Q           optional.String
	Type_       optional.String
	Year        optional.Float64
	Company     optional.String
	Country     optional.String
	Director    optional.String
	Language    optional.String
	PrimaryType optional.String
	Network     optional.String
	RemoteId    optional.String
	Offset      optional.Float64
	Limit       optional.Float64
}

func (a *SearchApiService) GetSearchResults(ctx context.Context, localVarOptionals *SearchApiGetSearchResultsOpts) (InlineResponse20039, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse20039
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Query.IsSet() {
		localVarQueryParams.Add("query", parameterToString(localVarOptionals.Query.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Q.IsSet() {
		localVarQueryParams.Add("q", parameterToString(localVarOptionals.Q.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Year.IsSet() {
		localVarQueryParams.Add("year", parameterToString(localVarOptionals.Year.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Company.IsSet() {
		localVarQueryParams.Add("company", parameterToString(localVarOptionals.Company.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Country.IsSet() {
		localVarQueryParams.Add("country", parameterToString(localVarOptionals.Country.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Director.IsSet() {
		localVarQueryParams.Add("director", parameterToString(localVarOptionals.Director.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Language.IsSet() {
		localVarQueryParams.Add("language", parameterToString(localVarOptionals.Language.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PrimaryType.IsSet() {
		localVarQueryParams.Add("primaryType", parameterToString(localVarOptionals.PrimaryType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Network.IsSet() {
		localVarQueryParams.Add("network", parameterToString(localVarOptionals.Network.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RemoteId.IsSet() {
		localVarQueryParams.Add("remote_id", parameterToString(localVarOptionals.RemoteId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20039
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SearchApiService
Search a series, movie, people, episode, company or season by specific remote id and returns a base record for that entity.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param remoteId Search for a specific remote id.  Allows searching for an IMDB or EIDR id, for example.

@return InlineResponse20040
*/
func (a *SearchApiService) GetSearchResultsByRemoteId(ctx context.Context, remoteId string) (InlineResponse20040, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse20040
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/search/remoteid/{remoteId}"
	localVarPath = strings.Replace(localVarPath, "{"+"remoteId"+"}", fmt.Sprintf("%v", remoteId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse20040
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
