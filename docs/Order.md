# Order

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the order. | [default to null]
**CreatedAt** | **string** | Event occurrence UTC date-time (YYYY-MM-DDTHH:mm:ssZ), when order is actually created. | [default to null]
**UpdatedAt** | **string** | Event updated UTC date-time (YYYY-MM-DDTHH:mm:ssZ), when the status of the order is actually changed/updated. | [default to null]
**Status** | **string** | State of the order. | [default to null]
**Amount** | **float32** | Total amount of the order, including all shipping expenses, tax and the price of items. | [default to null]
**Products** | [**[]OrderProducts**](orderProducts.md) |  | [default to null]
**Email** | **string** | Email of the contact, Mandatory if \&quot;phone\&quot; field is not passed in \&quot;billing\&quot; parameter. | [optional] [default to null]
**Billing** | [***OrderBilling**](orderBilling.md) |  | [optional] [default to null]
**Coupons** | **[]string** | Coupons applied to the order. Stored case insensitive. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


