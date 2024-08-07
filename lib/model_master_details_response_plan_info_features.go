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

type MasterDetailsResponsePlanInfoFeatures struct {
	// Name of the feature
	Name string `json:"name,omitempty"`
	// Unit value of the feature
	UnitValue string `json:"unitValue,omitempty"`
	// Quantity provided in the plan
	Quantity int64 `json:"quantity,omitempty"`
	// Quantity with overages provided in the plan (only applicable on ENTv2)
	QuantityWithOverages int64 `json:"quantityWithOverages,omitempty"`
	// Quantity consumed by master
	Used int64 `json:"used,omitempty"`
	// Quantity consumed by sub-organizations over the admin plan limit (only applicable on ENTv2)
	UsedOverages int64 `json:"usedOverages,omitempty"`
	// Quantity remaining in the plan
	Remaining int64 `json:"remaining,omitempty"`
}
