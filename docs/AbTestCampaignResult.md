# AbTestCampaignResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WinningVersion** | **string** | Winning Campaign Info. pending &#x3D; Campaign has been picked for sending and winning version is yet to be decided, tie &#x3D; A tie happened between both the versions, notAvailable &#x3D; Campaign has not yet been picked for sending. | [optional] [default to null]
**WinningCriteria** | **string** | Criteria choosen for winning version (Open/Click) | [optional] [default to null]
**WinningSubjectLine** | **string** | Subject Line of current winning version | [optional] [default to null]
**OpenRate** | **string** | Open rate for current winning version | [optional] [default to null]
**ClickRate** | **string** | Click rate for current winning version | [optional] [default to null]
**WinningVersionRate** | **string** | Open/Click rate for the winner version | [optional] [default to null]
**Statistics** | [***AbTestCampaignResultStatistics**](AbTestCampaignResultStatistics.md) |  | [optional] [default to null]
**ClickedLinks** | [***AbTestCampaignResultClickedLinks**](AbTestCampaignResultClickedLinks.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


