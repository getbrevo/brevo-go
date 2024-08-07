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

type Event struct {
	// The name of the event that occurred. This is how you will find your event in Brevo. Limited to 255 characters, alphanumerical characters and - _ only.
	EventName string `json:"event_name"`
	// Timestamp of when the event occurred (e.g. \"2024-01-24T17:39:57+01:00\"). If no value is passed, the timestamp of the event creation is used.
	EventDate   string            `json:"event_date,omitempty"`
	Identifiers *EventIdentifiers `json:"identifiers"`
	// Properties defining the state of the contact associated to this event. Useful to update contact attributes defined in your contacts database while passing the event. For example: **\"FIRSTNAME\": \"Jane\" , \"AGE\": 37**
	ContactProperties map[string]interface{} `json:"contact_properties,omitempty"`
	// Properties of the event. Top level properties and nested properties can be used to better segment contacts and personalise workflow conditions. The following field type are supported: string, number, boolean (true/false), date (Timestamp e.g. \"2024-01-24T17:39:57+01:00\"). Keys are limited to 255 characters, alphanumerical characters and - _ only. Size is limited to 50Kb.
	EventProperties map[string]interface{} `json:"event_properties,omitempty"`
}
