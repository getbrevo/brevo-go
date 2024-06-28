# SsoTokenRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the sub-account organization | [default to null]
**Email** | **string** | User email of sub-account organization | [optional] [default to null]
**Target** | **string** | Set target after login success * automation - Redirect to Automation after login * email_campaign - Redirect to Email Campaign after login * contacts - Redirect to Contacts after login * landing_pages - Redirect to Landing Pages after login * email_transactional - Redirect to Email Transactional after login * senders - Redirect to Contacts after login * sms_campaign - Redirect to Sms Campaign after login * sms_transactional - Redirect to Sms Transactional after login  | [optional] [default to null]
**Url** | **string** | Set the full target URL after login success. The user will land directly on this target URL after login | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


