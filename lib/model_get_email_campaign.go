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

type GetEmailCampaign struct {
	// ID of the campaign
	Id int64 `json:"id"`
	// Name of the campaign
	Name string `json:"name"`
	// Subject of the campaign. Only available if `abTesting` flag of the campaign is `false`
	Subject string `json:"subject,omitempty"`
	// Type of campaign
	Type_ string `json:"type"`
	// Status of the campaign
	Status string `json:"status"`
	// UTC date-time on which campaign is scheduled (YYYY-MM-DDTHH:mm:ss.SSSZ)
	ScheduledAt string `json:"scheduledAt,omitempty"`
	// Status of A/B Test for the campaign. abTesting = false means it is disabled, & abTesting = true means it is enabled.
	AbTesting bool `json:"abTesting,omitempty"`
	// Subject A of the ab-test campaign. Only available if `abTesting` flag of the campaign is `true`
	SubjectA string `json:"subjectA,omitempty"`
	// Subject B of the ab-test campaign. Only available if `abTesting` flag of the campaign is `true`
	SubjectB string `json:"subjectB,omitempty"`
	// The size of your ab-test groups. Only available if `abTesting` flag of the campaign is `true`
	SplitRule int32 `json:"splitRule,omitempty"`
	// Criteria for the winning version. Only available if `abTesting` flag of the campaign is `true`
	WinnerCriteria string `json:"winnerCriteria,omitempty"`
	// The duration of the test in hours at the end of which the winning version will be sent. Only available if `abTesting` flag of the campaign is `true`
	WinnerDelay int32 `json:"winnerDelay,omitempty"`
	// It is true if you have chosen to send your campaign at best time, otherwise it is false
	SendAtBestTime bool `json:"sendAtBestTime,omitempty"`
	// Retrieved the status of test email sending. (true=Test email has been sent  false=Test email has not been sent)
	TestSent bool `json:"testSent"`
	// Header of the campaign
	Header string `json:"header"`
	// Footer of the campaign
	Footer string                             `json:"footer"`
	Sender *GetExtendedCampaignOverviewSender `json:"sender"`
	// Email defined as the \"Reply to\" of the campaign
	ReplyTo string `json:"replyTo"`
	// Customisation of the \"to\" field of the campaign
	ToField string `json:"toField,omitempty"`
	// HTML content of the campaign
	HtmlContent string `json:"htmlContent"`
	// Link to share the campaign on social medias
	ShareLink string `json:"shareLink,omitempty"`
	// Tag of the campaign
	Tag string `json:"tag,omitempty"`
	// Creation UTC date-time of the campaign (YYYY-MM-DDTHH:mm:ss.SSSZ)
	CreatedAt string `json:"createdAt"`
	// UTC date-time of last modification of the campaign (YYYY-MM-DDTHH:mm:ss.SSSZ)
	ModifiedAt string `json:"modifiedAt"`
	// Status of inline image. inlineImageActivation = false means image can’t be embedded, & inlineImageActivation = true means image can be embedded, in the email.
	InlineImageActivation bool `json:"inlineImageActivation,omitempty"`
	// Status of mirror links in campaign. mirrorActive = false means mirror links are deactivated, & mirrorActive = true means mirror links are activated, in the campaign
	MirrorActive bool `json:"mirrorActive,omitempty"`
	// FOR TRIGGER ONLY ! Type of trigger campaign.recurring = false means contact can receive the same Trigger campaign only once, & recurring = true means contact can receive the same Trigger campaign several times
	Recurring bool `json:"recurring,omitempty"`
	// Sent UTC date-time of the campaign (YYYY-MM-DDTHH:mm:ss.SSSZ). Only available if 'status' of the campaign is 'sent'
	SentDate string `json:"sentDate,omitempty"`
	// Total number of non-delivered campaigns for a particular campaign id.
	ReturnBounce int64        `json:"returnBounce,omitempty"`
	Recipients   *interface{} `json:"recipients"`
	Statistics   *interface{} `json:"statistics"`
}
