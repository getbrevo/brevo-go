# CreateSenderIps

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | Dedicated IP available in your account | [default to null]
**Domain** | **string** | Domain of the IP | [default to null]
**Weight** | **int64** | Weight to apply to the IP. Sum of all IP weights must be 100. Should be passed for either ALL or NONE of the IPs. If it&#39;s not passed, the sending will be equally balanced on all IPs. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


