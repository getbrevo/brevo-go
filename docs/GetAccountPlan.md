# GetAccountPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Displays the plan type of the user | [default to null]
**CreditsType** | **string** | This is the type of the credit, \&quot;Send Limit\&quot; is one of the possible types of credit of a user. \&quot;Send Limit\&quot; implies the total number of emails you can send to the subscribers in your account. | [default to null]
**Credits** | **float32** | Remaining credits of the user | [default to null]
**StartDate** | **string** | Date of the period from which the plan will start (only available for \&quot;subscription\&quot; and \&quot;reseller\&quot; plan type) | [optional] [default to null]
**EndDate** | **string** | Date of the period from which the plan will end (only available for \&quot;subscription\&quot; and \&quot;reseller\&quot; plan type) | [optional] [default to null]
**UserLimit** | **int32** | Only in case of reseller account. It implies the total number of child accounts you can add to your account. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


