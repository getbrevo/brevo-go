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

type UpdateSmtpTemplate struct {
	// Tag of the template
	Tag string `json:"tag,omitempty"`
	Sender *UpdateSmtpTemplateSender `json:"sender,omitempty"`
	// Name of the template
	TemplateName string `json:"templateName,omitempty"`
	// Required if htmlUrl is empty. If the template is designed using Drag & Drop editor via HTML content, then the design page will not have Drag & Drop editor access for that template. Body of the message (HTML must have more than 10 characters)
	HtmlContent string `json:"htmlContent,omitempty"`
	// Required if htmlContent is empty. URL to the body of the email (HTML)
	HtmlUrl string `json:"htmlUrl,omitempty"`
	// Subject of the email
	Subject string `json:"subject,omitempty"`
	// Email on which campaign recipients will be able to reply to
	ReplyTo string `json:"replyTo,omitempty"`
	// To personalize the «To» Field. If you want to include the first name and last name of your recipient, add {FNAME} {LNAME}. These contact attributes must already exist in your Brevo account. If input parameter 'params' used please use {{contact.FNAME}} {{contact.LNAME}} for personalization
	ToField string `json:"toField,omitempty"`
	// Absolute url of the attachment (no local file). Extension allowed: xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff, rtf, bmp, cgm, css, shtml, html, htm, zip, xml, ppt, pptx, tar, ez, ics, mobi, msg, pub and eps
	AttachmentUrl string `json:"attachmentUrl,omitempty"`
	// Status of the template. isActive = false means template is inactive, isActive = true means template is active
	IsActive bool `json:"isActive,omitempty"`
}
