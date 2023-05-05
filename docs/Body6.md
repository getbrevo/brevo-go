# Body6

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of task | [default to null]
**Duration** | **int64** | Duration of task in milliseconds [1 minute &#x3D; 60000 ms] | [optional] [default to null]
**TaskTypeId** | **string** | Id for type of task e.g Call / Email / Meeting etc. | [default to null]
**Date** | [**time.Time**](time.Time.md) | Task due date and time | [default to null]
**Notes** | **string** | Notes added to a task | [optional] [default to null]
**Done** | **bool** | Task marked as done | [optional] [default to null]
**AssignToId** | **string** | User id to whom task is assigned | [optional] [default to null]
**ContactsIds** | **[]int32** | Contact ids for contacts linked to this task | [optional] [default to null]
**DealsIds** | **[]string** | Deal ids for deals a task is linked to | [optional] [default to null]
**CompaniesIds** | **[]string** | Companies ids for companies a task is linked to | [optional] [default to null]
**Reminder** | [***TaskReminder**](TaskReminder.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


