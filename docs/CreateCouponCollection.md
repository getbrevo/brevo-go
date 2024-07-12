# CreateCouponCollection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the coupons collection | [default to null]
**DefaultCoupon** | **string** | Default coupons collection name | [default to null]
**ExpirationDate** | [**time.Time**](time.Time.md) | Specify an expiration date for the coupon collection in RFC3339 format. Use null to remove the expiration date. | [optional] [default to null]
**RemainingDaysAlert** | **int32** | Send a notification alert (email) when the remaining days until the expiration date are equal or fall bellow this number. Use null to disable alerts. | [optional] [default to null]
**RemainingCouponsAlert** | **int32** | Send a notification alert (email) when the remaining coupons count is equal or fall bellow this number. Use null to disable alerts. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


