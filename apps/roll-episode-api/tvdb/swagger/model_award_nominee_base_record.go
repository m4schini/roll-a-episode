/*
 * TVDB API V4
 *
 * Documentation of [TheTVDB](https://thetvdb.com/) API V4. All related information is linked from our [Github repo](https://github.com/thetvdb/v4-api). You might also want to use our [Postman collection] (https://www.getpostman.com/collections/7a9397ce69ff246f74d0) ## Authentication 1. Use the /login endpoint and provide your API key as \"apikey\". If you have a user-supported key, also provide your subscriber PIN as \"pin\". Otherwise completely remove \"pin\" from your call. 2. Executing this call will provide you with a bearer token, which is valid for 1 month. 3. Provide your bearer token for subsequent API calls by clicking Authorize below or including in the header of all direct API calls: `Authorization: Bearer [your-token]`  ## Notes 1. \"score\" is a field across almost all entities.  We generate scores for different types of entities in various ways, so no assumptions should be made about the meaning of this value.  It is simply used to hint at relative popularity for sorting purposes. 
 *
 * API version: 4.7.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// base award nominee record
type AwardNomineeBaseRecord struct {
	Character *Character       `json:"character,omitempty"`
	Details string             `json:"details,omitempty"`
	Episode *EpisodeBaseRecord `json:"episode,omitempty"`
	Id int64                   `json:"id,omitempty"`
	IsWinner bool              `json:"isWinner,omitempty"`
	Movie *MovieBaseRecord     `json:"movie,omitempty"`
	Series *SeriesBaseRecord   `json:"series,omitempty"`
	Year string                `json:"year,omitempty"`
	Category string `json:"category,omitempty"`
	Name string `json:"name,omitempty"`
}