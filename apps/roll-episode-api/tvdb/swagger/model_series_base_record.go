/*
 * TVDB API V4
 *
 * Documentation of [TheTVDB](https://thetvdb.com/) API V4. All related information is linked from our [Github repo](https://github.com/thetvdb/v4-api). You might also want to use our [Postman collection] (https://www.getpostman.com/collections/7a9397ce69ff246f74d0) ## Authentication 1. Use the /login endpoint and provide your API key as \"apikey\". If you have a user-supported key, also provide your subscriber PIN as \"pin\". Otherwise completely remove \"pin\" from your call. 2. Executing this call will provide you with a bearer token, which is valid for 1 month. 3. Provide your bearer token for subsequent API calls by clicking Authorize below or including in the header of all direct API calls: `Authorization: Bearer [your-token]`  ## Notes 1. \"score\" is a field across almost all entities.  We generate scores for different types of entities in various ways, so no assumptions should be made about the meaning of this value.  It is simply used to hint at relative popularity for sorting purposes. 
 *
 * API version: 4.7.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The base record for a series. All series airs time like firstAired, lastAired, nextAired, etc. are in US EST for US series, and for all non-US series, the time of the show’s country capital or most populous city. For streaming services, is the official release time. See https://support.thetvdb.com/kb/faq.php?id=29.
type SeriesBaseRecord struct {
	Aliases []Alias              `json:"aliases,omitempty"`
	AverageRuntime int32         `json:"averageRuntime,omitempty"`
	Country string `json:"country,omitempty"`
	DefaultSeasonType int64      `json:"defaultSeasonType,omitempty"`
	Episodes []EpisodeBaseRecord `json:"episodes,omitempty"`
	FirstAired string            `json:"firstAired,omitempty"`
	Id int32 `json:"id,omitempty"`
	Image string `json:"image,omitempty"`
	IsOrderRandomized bool `json:"isOrderRandomized,omitempty"`
	LastAired string `json:"lastAired,omitempty"`
	LastUpdated string `json:"lastUpdated,omitempty"`
	Name string `json:"name,omitempty"`
	NameTranslations []string `json:"nameTranslations,omitempty"`
	NextAired string `json:"nextAired,omitempty"`
	OriginalCountry string `json:"originalCountry,omitempty"`
	OriginalLanguage string `json:"originalLanguage,omitempty"`
	OverviewTranslations []string `json:"overviewTranslations,omitempty"`
	Score float64 `json:"score,omitempty"`
	Slug string                  `json:"slug,omitempty"`
	Status *Status               `json:"status,omitempty"`
	Year string                  `json:"year,omitempty"`
}