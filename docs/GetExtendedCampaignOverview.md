# GetExtendedCampaignOverview

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
**TestSent** | **bool** | Retrieved the status of test email sending. (true&#x3D;Test email has been sent  false&#x3D;Test email has not been sent) | [default to null]
**Header** | **string** | Header of the campaign | [default to null]
**Footer** | **string** | Footer of the campaign | [default to null]
**Sender** | [***GetExtendedCampaignOverviewSender**](getExtendedCampaignOverviewSender.md) |  | [default to null]
**ReplyTo** | **string** | Email defined as the \&quot;Reply to\&quot; of the campaign | [default to null]
**ToField** | **string** | Customisation of the \&quot;to\&quot; field of the campaign | [optional] [default to null]
**HtmlContent** | **string** | HTML content of the campaign | [default to null]
**ShareLink** | **string** | Link to share the campaign on social medias | [optional] [default to null]
**Tag** | **string** | Tag of the campaign | [optional] [default to null]
**CreatedAt** | **string** | Creation UTC date-time of the campaign (YYYY-MM-DDTHH:mm:ss.SSSZ) | [default to null]
**ModifiedAt** | **string** | UTC date-time of last modification of the campaign (YYYY-MM-DDTHH:mm:ss.SSSZ) | [default to null]
**InlineImageActivation** | **bool** | Status of inline image. inlineImageActivation &#x3D; false means image can’t be embedded, &amp; inlineImageActivation &#x3D; true means image can be embedded, in the email. | [optional] [default to null]
**MirrorActive** | **bool** | Status of mirror links in campaign. mirrorActive &#x3D; false means mirror links are deactivated, &amp; mirrorActive &#x3D; true means mirror links are activated, in the campaign | [optional] [default to null]
**Recurring** | **bool** | FOR TRIGGER ONLY ! Type of trigger campaign.recurring &#x3D; false means contact can receive the same Trigger campaign only once, &amp; recurring &#x3D; true means contact can receive the same Trigger campaign several times | [optional] [default to null]
**SentDate** | **string** | Sent UTC date-time of the campaign (YYYY-MM-DDTHH:mm:ss.SSSZ). Only available if &#39;status&#39; of the campaign is &#39;sent&#39; | [optional] [default to null]
**ReturnBounce** | **int64** | Total number of non-delivered campaigns for a particular campaign id. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


