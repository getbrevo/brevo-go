# GetEmailEventReportEvents

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address which generates the event | [default to null]
**Date** | **string** | UTC date-time on which the event has been generated | [default to null]
**Subject** | **string** | Subject of the event | [optional] [default to null]
**MessageId** | **string** | Message ID which generated the event | [default to null]
**Event** | **string** | Event which occurred | [default to null]
**Reason** | **string** | Reason of bounce (only available if the event is hardbounce or softbounce) | [optional] [default to null]
**Tag** | **string** | Tag of the email which generated the event | [optional] [default to null]
**Ip** | **string** | IP from which the user has opened the email or clicked on the link (only available if the event is opened or clicks) | [optional] [default to null]
**Link** | **string** | The link which is sent to the user (only available if the event is requests or opened or clicks) | [optional] [default to null]
**From** | **string** | Sender email from which the emails are sent | [optional] [default to null]
**TemplateId** | **int64** | ID of the template (only available if the email is template based) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


