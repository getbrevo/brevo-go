# Body8

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of deal | [default to null]
**Attributes** | [**map[string]interface{}**]| Attributes for deal creation  To assign owner of a Deal you can send attributes.deal_owner and utilize the account email or ID.  If you want to create a deal on a specific pipeline and stage you can use the following attributes &#x60;pipeline&#x60; and &#x60;deal_stage&#x60;.  Pipeline and deal_stage are ids you can fetch using this endpoint &#x60;/crm/pipeline/details/{pipelineID}&#x60;  | [optional] [default to null]
**LinkedContactsIds** | **[]int64** | Contact ids to be linked with deal | [optional] [default to null]
**LinkedCompaniesIds** | **[]string** | Company ids to be linked with deal | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


