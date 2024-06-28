# GetExtendedCampaignStats

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalStats** | [**map[string]interface{}**](interface{}.md) | Overall statistics of the campaign | [default to null]
**CampaignStats** | [**[]interface{}**](interface{}.md) | List-wise statistics of the campaign. | [default to null]
**MirrorClick** | **int64** | Number of clicks on mirror link | [default to null]
**Remaining** | **int64** | Number of remaning emails to send | [default to null]
**LinksStats** | [**map[string]interface{}**](interface{}.md) | Statistics about the number of clicks for the links | [default to null]
**StatsByDomain** | [***GetStatsByDomain**](getStatsByDomain.md) |  | [default to null]
**StatsByDevice** | [***GetStatsByDevice**](getStatsByDevice.md) | Statistics about the campaign on the basis of various devices | [default to null]
**StatsByBrowser** | [***GetStatsByBrowser**](getStatsByBrowser.md) | Statistics about the campaign on the basis of various browsers | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


