/*
 * Brevo API
 *
 * Brevo provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/brevo  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  |   | 406  | Error. Not Acceptable  |   | 422  | Error. Unprocessable Entity |
 *
 * API version: 3.0.0
 * Contact: contact@brevo.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package lib

type AbTestCampaignResult struct {
	// Winning Campaign Info. pending = Campaign has been picked for sending and winning version is yet to be decided, tie = A tie happened between both the versions, notAvailable = Campaign has not yet been picked for sending.
	WinningVersion string `json:"winningVersion,omitempty"`
	// Criteria choosen for winning version (Open/Click)
	WinningCriteria string `json:"winningCriteria,omitempty"`
	// Subject Line of current winning version
	WinningSubjectLine string `json:"winningSubjectLine,omitempty"`
	// Open rate for current winning version
	OpenRate string `json:"openRate,omitempty"`
	// Click rate for current winning version
	ClickRate string `json:"clickRate,omitempty"`
	// Open/Click rate for the winner version
	WinningVersionRate string                            `json:"winningVersionRate,omitempty"`
	Statistics         *AbTestCampaignResultStatistics   `json:"statistics,omitempty"`
	ClickedLinks       *AbTestCampaignResultClickedLinks `json:"clickedLinks,omitempty"`
}
