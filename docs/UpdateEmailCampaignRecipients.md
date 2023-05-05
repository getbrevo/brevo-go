# UpdateEmailCampaignRecipients

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExclusionListIds** | **[]int64** | List ids which have to be excluded from a campaign | [optional] [default to null]
**ListIds** | **[]int64** | Lists Ids to send the campaign to. Campaign should only be updated with listIds if listIds were used to create it. REQUIRED if already not present in campaign and scheduledAt is not empty | [optional] [default to null]
**SegmentIds** | **[]int64** | Mandatory if listIds are not used. Campaign should only be updated with segmentIds if segmentIds were used to create it. Segment ids to send the campaign to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


