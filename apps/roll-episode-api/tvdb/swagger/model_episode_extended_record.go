/*
 * TVDB API V4
 *
 * Documentation of [TheTVDB](https://thetvdb.com/) API V4. All related information is linked from our [Github repo](https://github.com/thetvdb/v4-api). You might also want to use our [Postman collection] (https://www.getpostman.com/collections/7a9397ce69ff246f74d0) ## Authentication 1. Use the /login endpoint and provide your API key as \"apikey\". If you have a user-supported key, also provide your subscriber PIN as \"pin\". Otherwise completely remove \"pin\" from your call. 2. Executing this call will provide you with a bearer token, which is valid for 1 month. 3. Provide your bearer token for subsequent API calls by clicking Authorize below or including in the header of all direct API calls: `Authorization: Bearer [your-token]`  ## Notes 1. \"score\" is a field across almost all entities.  We generate scores for different types of entities in various ways, so no assumptions should be made about the meaning of this value.  It is simply used to hint at relative popularity for sorting purposes. 
 *
 * API version: 4.7.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// extended episode record
type EpisodeExtendedRecord struct {
	Aired string `json:"aired,omitempty"`
	AirsAfterSeason int32 `json:"airsAfterSeason,omitempty"`
	AirsBeforeEpisode int32 `json:"airsBeforeEpisode,omitempty"`
	AirsBeforeSeason int32         `json:"airsBeforeSeason,omitempty"`
	Awards []AwardBaseRecord       `json:"awards,omitempty"`
	Characters []Character         `json:"characters,omitempty"`
	Companies []Company            `json:"companies,omitempty"`
	ContentRatings []ContentRating `json:"contentRatings,omitempty"`
	// season, midseason, or series
	FinaleType string `json:"finaleType,omitempty"`
	Id int64 `json:"id,omitempty"`
	Image string `json:"image,omitempty"`
	ImageType int32 `json:"imageType,omitempty"`
	IsMovie int64 `json:"isMovie,omitempty"`
	LastUpdated string `json:"lastUpdated,omitempty"`
	LinkedMovie int32 `json:"linkedMovie,omitempty"`
	Name string `json:"name,omitempty"`
	NameTranslations []string            `json:"nameTranslations,omitempty"`
	Networks []Company                   `json:"networks,omitempty"`
	Nominations []AwardNomineeBaseRecord `json:"nominations,omitempty"`
	Number int32                         `json:"number,omitempty"`
	Overview string `json:"overview,omitempty"`
	OverviewTranslations []string `json:"overviewTranslations,omitempty"`
	ProductionCode string                `json:"productionCode,omitempty"`
	RemoteIds []RemoteId                 `json:"remoteIds,omitempty"`
	Runtime int32                        `json:"runtime,omitempty"`
	SeasonNumber int32                   `json:"seasonNumber,omitempty"`
	Seasons []SeasonBaseRecord           `json:"seasons,omitempty"`
	SeriesId int64                       `json:"seriesId,omitempty"`
	Studios []Company                    `json:"studios,omitempty"`
	TagOptions []TagOption               `json:"tagOptions,omitempty"`
	Trailers []Trailer                   `json:"trailers,omitempty"`
	Translations *TranslationExtended    `json:"translations,omitempty"`
	Year string                          `json:"year,omitempty"`
}
