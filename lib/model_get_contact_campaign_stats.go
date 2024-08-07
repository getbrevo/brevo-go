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

// Campaign Statistics for the contact
type GetContactCampaignStats struct {
	MessagesSent      []GetExtendedContactDetailsStatisticsMessagesSent `json:"messagesSent,omitempty"`
	HardBounces       []GetExtendedContactDetailsStatisticsMessagesSent `json:"hardBounces,omitempty"`
	SoftBounces       []GetExtendedContactDetailsStatisticsMessagesSent `json:"softBounces,omitempty"`
	Complaints        []GetExtendedContactDetailsStatisticsMessagesSent `json:"complaints,omitempty"`
	Unsubscriptions   *GetContactCampaignStatsUnsubscriptions           `json:"unsubscriptions,omitempty"`
	Opened            []GetContactCampaignStatsOpened                   `json:"opened,omitempty"`
	Clicked           []GetContactCampaignStatsClicked                  `json:"clicked,omitempty"`
	TransacAttributes []GetContactCampaignStatsTransacAttributes        `json:"transacAttributes,omitempty"`
	Delivered         []GetExtendedContactDetailsStatisticsMessagesSent `json:"delivered,omitempty"`
}
