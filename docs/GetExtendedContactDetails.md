# GetExtendedContactDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address of the contact for which you requested the details | [optional] [default to null]
**Id** | **int64** | ID of the contact for which you requested the details | [default to null]
**EmailBlacklisted** | **bool** | Blacklist status for email campaigns (true&#x3D;blacklisted, false&#x3D;not blacklisted) | [default to null]
**SmsBlacklisted** | **bool** | Blacklist status for SMS campaigns (true&#x3D;blacklisted, false&#x3D;not blacklisted) | [default to null]
**CreatedAt** | **string** | Creation UTC date-time of the contact (YYYY-MM-DDTHH:mm:ss.SSSZ) | [default to null]
**ModifiedAt** | **string** | Last modification UTC date-time of the contact (YYYY-MM-DDTHH:mm:ss.SSSZ) | [default to null]
**ListIds** | **[]int64** |  | [default to null]
**ListUnsubscribed** | **[]int64** |  | [optional] [default to null]
**Attributes** | [**map[string]interface{}**] | Set of attributes of the contact | [default to null]
**Statistics** | [***GetExtendedContactDetailsStatistics**](GetExtendedContactDetailsStatistics.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


