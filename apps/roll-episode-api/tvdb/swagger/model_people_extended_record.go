/*
 * TVDB API V4
 *
 * Documentation of [TheTVDB](https://thetvdb.com/) API V4. All related information is linked from our [Github repo](https://github.com/thetvdb/v4-api). You might also want to use our [Postman collection] (https://www.getpostman.com/collections/7a9397ce69ff246f74d0) ## Authentication 1. Use the /login endpoint and provide your API key as \"apikey\". If you have a user-supported key, also provide your subscriber PIN as \"pin\". Otherwise completely remove \"pin\" from your call. 2. Executing this call will provide you with a bearer token, which is valid for 1 month. 3. Provide your bearer token for subsequent API calls by clicking Authorize below or including in the header of all direct API calls: `Authorization: Bearer [your-token]`  ## Notes 1. \"score\" is a field across almost all entities.  We generate scores for different types of entities in various ways, so no assumptions should be made about the meaning of this value.  It is simply used to hint at relative popularity for sorting purposes. 
 *
 * API version: 4.7.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// extended people record
type PeopleExtendedRecord struct {
	Aliases []Alias                   `json:"aliases,omitempty"`
	Awards []AwardBaseRecord          `json:"awards,omitempty"`
	Biographies []Biography           `json:"biographies,omitempty"`
	Birth string                      `json:"birth,omitempty"`
	BirthPlace string                 `json:"birthPlace,omitempty"`
	Characters []Character            `json:"characters,omitempty"`
	Death string                      `json:"death,omitempty"`
	Gender int32 `json:"gender,omitempty"`
	Id int64 `json:"id,omitempty"`
	Image string `json:"image,omitempty"`
	LastUpdated string `json:"lastUpdated,omitempty"`
	Name string `json:"name,omitempty"`
	NameTranslations []string `json:"nameTranslations,omitempty"`
	OverviewTranslations []string     `json:"overviewTranslations,omitempty"`
	Races []Race                      `json:"races,omitempty"`
	RemoteIds []RemoteId              `json:"remoteIds,omitempty"`
	Score int64                       `json:"score,omitempty"`
	Slug string                       `json:"slug,omitempty"`
	TagOptions []TagOption            `json:"tagOptions,omitempty"`
	Translations *TranslationExtended `json:"translations,omitempty"`
}