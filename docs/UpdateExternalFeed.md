# UpdateExternalFeed

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the feed | [optional] [default to null]
**Url** | **string** | URL of the feed | [optional] [default to null]
**AuthType** | **string** | Auth type of the feed:   * &#x60;basic&#x60;   * &#x60;token&#x60;   * &#x60;noAuth&#x60;  | [optional] [default to null]
**Username** | **string** | Username for authType &#x60;basic&#x60; | [optional] [default to null]
**Password** | **string** | Password for authType &#x60;basic&#x60; | [optional] [default to null]
**Token** | **string** | Token for authType &#x60;token&#x60; | [optional] [default to null]
**Headers** | [**[]GetExternalFeedByUuidHeaders**](getExternalFeedByUUIDHeaders.md) | Custom headers for the feed | [optional] [default to null]
**MaxRetries** | **int32** | Maximum number of retries on the feed url | [optional] [default to null]
**Cache** | **bool** | Toggle caching of feed url response | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


