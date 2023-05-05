# ExportWebhooksHistory

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **int32** | Number of days in the past including today (positive integer). _Not compatible with &#39;startDate&#39; and &#39;endDate&#39;_ | [optional] [default to null]
**StartDate** | **string** | Mandatory if endDate is used. Starting date of the history (YYYY-MM-DD). Must be lower than equal to endDate | [optional] [default to null]
**EndDate** | **string** | Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD). Must be greater than equal to startDate | [optional] [default to null]
**Sort** | **string** | Sorting order of records (asc or desc) | [optional] [default to null]
**Event** | **string** | Filter the history for a specific event type | [default to null]
**NotifyURL** | **string** | Webhook URL to receive CSV file link | [default to null]
**WebhookId** | **int32** | Filter the history for a specific webhook id | [optional] [default to null]
**Email** | **string** | Filter the history for a specific email | [optional] [default to null]
**MessageId** | **int32** | Filter the history for a specific message id | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


