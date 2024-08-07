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

type GetWhatsAppConfig struct {
	// Id of the WhatsApp business account
	WhatsappBusinessAccountId string `json:"whatsappBusinessAccountId,omitempty"`
	// Sending limit Information of the WhatsApp API account
	SendingLimit string `json:"sendingLimit,omitempty"`
	// Quality status of phone number associated with WhatsApp account. There are three quality ratings. example - **High (GREEN) , Medium (YELLOW) and Low(RED)**
	PhoneNumberQuality string `json:"phoneNumberQuality,omitempty"`
	// Status information related to WhatsApp Api account
	WhatsappBusinessAccountStatus string `json:"whatsappBusinessAccountStatus,omitempty"`
	// Verification status information of the Business account
	BusinessStatus string `json:"businessStatus,omitempty"`
	// Status of the name associated with WhatsApp Phone number
	PhoneNumberNameStatus string `json:"phoneNumberNameStatus,omitempty"`
}
