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

type UpdateEmailCampaign struct {
	// Tag of the campaign
	Tag    string                     `json:"tag,omitempty"`
	Sender *UpdateEmailCampaignSender `json:"sender,omitempty"`
	// Name of the campaign
	Name string `json:"name,omitempty"`
	// Body of the message (HTML version). If the campaign is designed using Drag & Drop editor via HTML content, then the design page will not have Drag & Drop editor access for that campaign. REQUIRED if htmlUrl is empty
	HtmlContent string `json:"htmlContent,omitempty"`
	// Url which contents the body of the email message. REQUIRED if htmlContent is empty
	HtmlUrl string `json:"htmlUrl,omitempty"`
	// UTC date-time on which the campaign has to run (YYYY-MM-DDTHH:mm:ss.SSSZ). Prefer to pass your timezone in date-time format for accurate result. If sendAtBestTime is set to true, your campaign will be sent according to the date passed (ignoring the time part).
	ScheduledAt string `json:"scheduledAt,omitempty"`
	// Subject of the campaign
	Subject string `json:"subject,omitempty"`
	// Preview text or preheader of the email campaign
	PreviewText string `json:"previewText,omitempty"`
	// Email on which campaign recipients will be able to reply to
	ReplyTo string `json:"replyTo,omitempty"`
	// To personalize the «To» Field. If you want to include the first name and last name of your recipient, add {FNAME} {LNAME}. These contact attributes must already exist in your Brevo account. If input parameter 'params' used please use {{contact.FNAME}} {{contact.LNAME}} for personalization
	ToField    string                         `json:"toField,omitempty"`
	Recipients *UpdateEmailCampaignRecipients `json:"recipients,omitempty"`
	// Absolute url of the attachment (no local file). Extension allowed: xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff, rtf, bmp, cgm, css, shtml, html, htm, zip, xml, ppt, pptx, tar, ez, ics, mobi, msg, pub and eps
	AttachmentUrl string `json:"attachmentUrl,omitempty"`
	// Status of inline image. inlineImageActivation = false means image can’t be embedded, & inlineImageActivation = true means image can be embedded, in the email. You cannot send a campaign of more than 4MB with images embedded in the email. Campaigns with the images embedded in the email must be sent to less than 5000 contacts.
	InlineImageActivation bool `json:"inlineImageActivation,omitempty"`
	// Status of mirror links in campaign. mirrorActive = false means mirror links are deactivated, & mirrorActive = true means mirror links are activated, in the campaign
	MirrorActive bool `json:"mirrorActive,omitempty"`
	// FOR TRIGGER ONLY ! Type of trigger campaign.recurring = false means contact can receive the same Trigger campaign only once, & recurring = true means contact can receive the same Trigger campaign several times
	Recurring bool `json:"recurring,omitempty"`
	// Footer of the email campaign
	Footer string `json:"footer,omitempty"`
	// Header of the email campaign
	Header string `json:"header,omitempty"`
	// Customize the utm_campaign value. If this field is empty, the campaign name will be used. Only alphanumeric characters and spaces are allowed
	UtmCampaign string `json:"utmCampaign,omitempty"`
	// Pass the set of attributes to customize the type 'classic' campaign. For example, {\"FNAME\":\"Joe\", \"LNAME\":\"Doe\"}. The 'params' field will get updated, only if the campaign is in New Template Language, else ignored. The New Template Language is dependent on the values of 'subject', 'htmlContent/htmlUrl', 'sender.name' & 'toField'
	Params map[string]interface{} `json:"params,omitempty"`
	// Set this to true if you want to send your campaign at best time. Note:- if true, warmup ip will be disabled.
	SendAtBestTime bool `json:"sendAtBestTime,omitempty"`
	// Status of A/B Test. abTesting = false means it is disabled, & abTesting = true means it is enabled. 'subjectA', 'subjectB', 'splitRule', 'winnerCriteria' & 'winnerDelay' will be considered if abTesting is set to true. 'subject' if passed is ignored.  Can be set to true only if 'sendAtBestTime' is 'false'. You will be able to set up two subject lines for your campaign and send them to a random sample of your total recipients. Half of the test group will receive version A, and the other half will receive version B
	AbTesting bool `json:"abTesting,omitempty"`
	// Subject A of the campaign. Considered if abTesting = true. subjectA & subjectB should have unique value
	SubjectA string `json:"subjectA,omitempty"`
	// Subject B of the campaign. Considered if abTesting = true. subjectA & subjectB should have unique value
	SubjectB string `json:"subjectB,omitempty"`
	// Add the size of your test groups. Considered if abTesting = true. We'll send version A and B to a random sample of recipients, and then the winning version to everyone else
	SplitRule int64 `json:"splitRule,omitempty"`
	// Choose the metrics that will determinate the winning version. Considered if 'splitRule' >= 1 and < 50. If splitRule = 50, 'winnerCriteria' is ignored if passed or alreday exist in record
	WinnerCriteria string `json:"winnerCriteria,omitempty"`
	// Choose the duration of the test in hours. Maximum is 7 days, pass 24*7 = 168 hours. The winning version will be sent at the end of the test. Considered if 'splitRule' >= 1 and < 50. If splitRule = 50, 'winnerDelay' is ignored if passed or alreday exist in record
	WinnerDelay int64 `json:"winnerDelay,omitempty"`
	// Available for dedicated ip clients. Set this to true if you wish to warm up your ip.
	IpWarmupEnable bool `json:"ipWarmupEnable,omitempty"`
	// Set an initial quota greater than 1 for warming up your ip. We recommend you set a value of 3000.
	InitialQuota int64 `json:"initialQuota,omitempty"`
	// Set a percentage increase rate for warming up your ip. We recommend you set the increase rate to 30% per day. If you want to send the same number of emails every day, set the daily increase value to 0%.
	IncreaseRate int64 `json:"increaseRate,omitempty"`
	// Enter an unsubscription page id. The page id is a 24 digit alphanumeric id that can be found in the URL when editing the page.
	UnsubscriptionPageId string `json:"unsubscriptionPageId,omitempty"`
	// Mandatory if templateId is used containing the {{ update_profile }} tag. Enter an update profile form id. The form id is a 24 digit alphanumeric id that can be found in the URL when editing the form.
	UpdateFormId string `json:"updateFormId,omitempty"`
}
