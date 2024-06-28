/*
 * Brevo API
 *
 * Brevo provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/brevo  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  |   | 406  | Error. Not Acceptable  |   | 422  | Error. Unprocessable Entity | 
 *
 * API version: 3.0.0
 * Contact: contact@brevo.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UpdateExternalFeed struct {
	// Name of the feed
	Name string `json:"name,omitempty"`
	// URL of the feed
	Url string `json:"url,omitempty"`
	// Auth type of the feed:   * `basic`   * `token`   * `noAuth` 
	AuthType string `json:"authType,omitempty"`
	// Username for authType `basic`
	Username string `json:"username,omitempty"`
	// Password for authType `basic`
	Password string `json:"password,omitempty"`
	// Token for authType `token`
	Token string `json:"token,omitempty"`
	// Custom headers for the feed
	Headers []GetExternalFeedByUuidHeaders `json:"headers,omitempty"`
	// Maximum number of retries on the feed url
	MaxRetries int32 `json:"maxRetries,omitempty"`
	// Toggle caching of feed url response
	Cache bool `json:"cache,omitempty"`
}
