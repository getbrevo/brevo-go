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

type GetWhatsappEventReportEvents struct {
	// WhatsApp Number with country code. Example, 85264318721
	ContactNumber string `json:"contactNumber"`
	// UTC date-time on which the event has been generated
	Date string `json:"date"`
	// Message ID which generated the event
	MessageId string `json:"messageId"`
	// Event which occurred
	Event string `json:"event"`
	// Reason for the event (will be there in case of `error` and `soft-bounce` events)
	Reason string `json:"reason,omitempty"`
	// Text of the reply (will be there only in case of `reply` event with text)
	Body string `json:"body,omitempty"`
	// Url of the media reply (will be there only in case of `reply` event with media)
	MediaUrl string `json:"mediaUrl,omitempty"`
	// WhatsApp Number with country code. Example, 85264318721
	SenderNumber string `json:"senderNumber"`
}
