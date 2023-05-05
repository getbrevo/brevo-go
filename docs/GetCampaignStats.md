# GetCampaignStats

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListId** | **int64** | List Id of email campaign (only in case of get email campaign(s)(not for global stats)) | [optional] [default to null]
**UniqueClicks** | **int64** | Number of unique clicks for the campaign | [default to null]
**Clickers** | **int64** | Number of total clicks for the campaign | [default to null]
**Complaints** | **int64** | Number of complaints (Spam reports) for the campaign | [default to null]
**Delivered** | **int64** | Number of delivered emails for the campaign | [default to null]
**Sent** | **int64** | Number of sent emails for the campaign | [default to null]
**SoftBounces** | **int64** | Number of softbounce for the campaign | [default to null]
**HardBounces** | **int64** | Number of harbounce for the campaign | [default to null]
**UniqueViews** | **int64** | Number of unique openings for the campaign | [default to null]
**TrackableViews** | **int64** | Recipients without any privacy protection option enabled in their email client | [default to null]
**TrackableViewsRate** | **float32** | Rate of recipients without any privacy protection option enabled in their email client | [optional] [default to null]
**EstimatedViews** | **int64** | Rate of recipients without any privacy protection option enabled in their email client, applied to all delivered emails | [optional] [default to null]
**Unsubscriptions** | **int64** | Number of unsubscription for the campaign | [default to null]
**Viewed** | **int64** | Number of openings for the campaign | [default to null]
**Deferred** | **int64** | Number of deferred emails for the campaign | [optional] [default to null]
**ReturnBounce** | **int64** | Total number of non-delivered campaigns for a particular campaign id. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


