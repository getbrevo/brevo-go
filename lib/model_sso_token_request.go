/*
 * Brevo API
 *
 * Brevo provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/brevo  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  |   | 406  | Error. Not Acceptable  |
 *
 * API version: 3.0.0
 * Contact: contact@brevo.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package lib

type SsoTokenRequest struct {
	// Id of the sub-account organization
	Id int64 `json:"id"`
	// User email of sub-account organization
	Email string `json:"email,omitempty"`
	// Set target after login success * automation - Redirect to Automation after login * email_campaign - Redirect to Email Campaign after login * contacts - Redirect to Contacts after login * landing_pages - Redirect to Landing Pages after login * email_transactional - Redirect to Email Transactional after login * senders - Redirect to Contacts after login * sms_campaign - Redirect to Sms Campaign after login * sms_transactional - Redirect to Sms Transactional after login
	Target string `json:"target,omitempty"`
}
