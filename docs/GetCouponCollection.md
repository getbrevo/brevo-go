# GetCouponCollection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the collection. | [default to null]
**Name** | **string** | The name of the collection. | [default to null]
**DefaultCoupon** | **string** | The default coupon of the collection. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Datetime on which the collection was created. | [default to null]
**TotalCoupons** | **int64** | Total number of coupons in the collection. | [default to null]
**RemainingCoupons** | **int64** | Number of coupons that have not been sent yet. | [default to null]
**ExpirationDate** | [**time.Time**](time.Time.md) | Expiration date for the coupon collection in RFC3339 format. | [optional] [default to null]
**RemainingDaysAlert** | **int32** | If present, an email notification is going to be sent the defined amount of days before the expiration date. | [optional] [default to null]
**RemainingCouponsAlert** | **int32** | If present, an email notification is going to be sent when the total number of available coupons falls below the defined threshold. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


