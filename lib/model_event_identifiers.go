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

// Identifies the contact associated with the event. At least one identifier is required.
type EventIdentifiers struct {
	// Email Id associated with the event
	EmailId string `json:"email_id,omitempty"`
	// SMS associated with the event
	PhoneId string `json:"phone_id,omitempty"`
	// whatsapp associated with the event
	WhatsappId string `json:"whatsapp_id,omitempty"`
	// landline_number associated with the event
	LandlineNumberId string `json:"landline_number_id,omitempty"`
	// ext_id associated with the event
	ExtId string `json:"ext_id,omitempty"`
}