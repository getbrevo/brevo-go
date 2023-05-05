# GetWhatsappEventReportEvents

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactNumber** | **string** | WhatsApp Number with country code. Example, 85264318721 | [default to null]
**Date** | **string** | UTC date-time on which the event has been generated | [default to null]
**MessageId** | **string** | Message ID which generated the event | [default to null]
**Event** | **string** | Event which occurred | [default to null]
**Reason** | **string** | Reason for the event (will be there in case of &#x60;error&#x60; and &#x60;soft-bounce&#x60; events) | [optional] [default to null]
**Body** | **string** | Text of the reply (will be there only in case of &#x60;reply&#x60; event with text) | [optional] [default to null]
**MediaUrl** | **string** | Url of the media reply (will be there only in case of &#x60;reply&#x60; event with media) | [optional] [default to null]
**SenderNumber** | **string** | WhatsApp Number with country code. Example, 85264318721 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


