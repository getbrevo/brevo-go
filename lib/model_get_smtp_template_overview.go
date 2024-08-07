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

type GetSmtpTemplateOverview struct {
	// ID of the template
	Id int64 `json:"id"`
	// Name of the template
	Name string `json:"name"`
	// Subject of the template
	Subject string `json:"subject"`
	// Status of template (true=active, false=inactive)
	IsActive bool `json:"isActive"`
	// Status of test sending for the template (true=test email has been sent, false=test email has not been sent)
	TestSent bool                           `json:"testSent"`
	Sender   *GetSmtpTemplateOverviewSender `json:"sender"`
	// Email defined as the \"Reply to\" for the template
	ReplyTo string `json:"replyTo"`
	// Customisation of the \"to\" field for the template
	ToField string `json:"toField"`
	// Tag of the template
	Tag string `json:"tag"`
	// HTML content of the template
	HtmlContent string `json:"htmlContent"`
	// Creation UTC date-time of the template (YYYY-MM-DDTHH:mm:ss.SSSZ)
	CreatedAt string `json:"createdAt"`
	// Last modification UTC date-time of the template (YYYY-MM-DDTHH:mm:ss.SSSZ)
	ModifiedAt string `json:"modifiedAt"`
	// It is true if template is a valid Double opt-in (DOI) template, otherwise it is false. This field will be available only in case of single template detail call.
	DoiTemplate bool `json:"doiTemplate,omitempty"`
}
