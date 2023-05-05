# GetCampaignOverview

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the campaign | [default to null]
**Name** | **string** | Name of the campaign | [default to null]
**Subject** | **string** | Subject of the campaign. Only available if &#x60;abTesting&#x60; flag of the campaign is &#x60;false&#x60; | [optional] [default to null]
**Type_** | **string** | Type of campaign | [default to null]
**Status** | **string** | Status of the campaign | [default to null]
**ScheduledAt** | **string** | UTC date-time on which campaign is scheduled (YYYY-MM-DDTHH:mm:ss.SSSZ) | [optional] [default to null]
**AbTesting** | **bool** | Status of A/B Test for the campaign. abTesting &#x3D; false means it is disabled, &amp; abTesting &#x3D; true means it is enabled. | [optional] [default to null]
**SubjectA** | **string** | Subject A of the ab-test campaign. Only available if &#x60;abTesting&#x60; flag of the campaign is &#x60;true&#x60; | [optional] [default to null]
**SubjectB** | **string** | Subject B of the ab-test campaign. Only available if &#x60;abTesting&#x60; flag of the campaign is &#x60;true&#x60; | [optional] [default to null]
**SplitRule** | **int32** | The size of your ab-test groups. Only available if &#x60;abTesting&#x60; flag of the campaign is &#x60;true&#x60; | [optional] [default to null]
**WinnerCriteria** | **string** | Criteria for the winning version. Only available if &#x60;abTesting&#x60; flag of the campaign is &#x60;true&#x60; | [optional] [default to null]
**WinnerDelay** | **int32** | The duration of the test in hours at the end of which the winning version will be sent. Only available if &#x60;abTesting&#x60; flag of the campaign is &#x60;true&#x60; | [optional] [default to null]
**SendAtBestTime** | **bool** | It is true if you have chosen to send your campaign at best time, otherwise it is false | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


