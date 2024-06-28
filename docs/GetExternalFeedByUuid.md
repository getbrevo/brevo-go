# GetExternalFeedByUuid

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the feed | [default to null]
**Name** | **string** | Name of the feed | [default to null]
**Url** | **string** | URL of the feed | [default to null]
**AuthType** | **string** | Auth type of the feed: * &#x60;basic&#x60; * &#x60;token&#x60; * &#x60;noAuth&#x60;  | [default to null]
**Username** | **string** | Username for authType &#x60;basic&#x60; | [optional] [default to null]
**Password** | **string** | Password for authType &#x60;basic&#x60; | [optional] [default to null]
**Token** | **string** | Token for authType &#x60;token&#x60; | [optional] [default to null]
**Headers** | [**[]GetExternalFeedByUuidHeaders**](getExternalFeedByUUID_headers.md) | Custom headers for the feed | [default to null]
**MaxRetries** | **int32** | Maximum number of retries on the feed url | [default to null]
**Cache** | **bool** | Toggle caching of feed url response | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Datetime on which the feed was created | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) | Datetime on which the feed was modified | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


