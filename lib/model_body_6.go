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

import (
	"time"
)

type Body6 struct {
	// Name of task
	Name string `json:"name"`
	// Duration of task in milliseconds [1 minute = 60000 ms]
	Duration int64 `json:"duration,omitempty"`
	// Id for type of task e.g Call / Email / Meeting etc.
	TaskTypeId string `json:"taskTypeId"`
	// Task due date and time
	Date time.Time `json:"date"`
	// Notes added to a task
	Notes string `json:"notes,omitempty"`
	// Task marked as done
	Done bool `json:"done,omitempty"`
	// User id to whom task is assigned
	AssignToId string `json:"assignToId,omitempty"`
	// Contact ids for contacts linked to this task
	ContactsIds []int32 `json:"contactsIds,omitempty"`
	// Deal ids for deals a task is linked to
	DealsIds []string `json:"dealsIds,omitempty"`
	// Companies ids for companies a task is linked to
	CompaniesIds []string      `json:"companiesIds,omitempty"`
	Reminder     *TaskReminder `json:"reminder,omitempty"`
}
