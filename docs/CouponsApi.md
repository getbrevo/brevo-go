# \CouponsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCouponCollection**](CouponsApi.md#CreateCouponCollection) | **Post** /couponCollections | Create а coupon collection
[**CreateCoupons**](CouponsApi.md#CreateCoupons) | **Post** /coupons | Create coupons for a coupon collection
[**GetCouponCollection**](CouponsApi.md#GetCouponCollection) | **Get** /couponCollections/{id} | Get a coupon collection by id
[**GetCouponCollections**](CouponsApi.md#GetCouponCollections) | **Get** /couponCollections | Get all your coupon collections
[**UpdateCouponCollection**](CouponsApi.md#UpdateCouponCollection) | **Patch** /couponCollections/{id} | Update a coupon collection by id


# **CreateCouponCollection**
> InlineResponse2013 CreateCouponCollection(ctx, createCouponCollection)
Create а coupon collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createCouponCollection** | [**CreateCouponCollection**](CreateCouponCollection.md)| Values to create a coupon collection | 

### Return type

[**InlineResponse2013**](InlineResponse2013.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCoupons**
> CreateCoupons(ctx, createCoupons)
Create coupons for a coupon collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createCoupons** | [**CreateCoupons**](CreateCoupons.md)| Values to create coupons | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCouponCollection**
> GetCouponCollection GetCouponCollection(ctx, id)
Get a coupon collection by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the collection to return | 

### Return type

[**GetCouponCollection**](GetCouponCollection.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCouponCollections**
> GetCouponCollection GetCouponCollections(ctx, optional)
Get all your coupon collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCouponCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCouponCollectionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number of documents returned per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document on the page | [default to 0]
 **sort** | **optional.String**| Sort the results by creation time in ascending/descending order | [default to desc]
 **sortBy** | **optional.String**| The field used to sort coupon collections | [default to createdAt]

### Return type

[**GetCouponCollection**](GetCouponCollection.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCouponCollection**
> InlineResponse2008 UpdateCouponCollection(ctx, id, optional)
Update a coupon collection by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the collection to update | 
 **optional** | ***UpdateCouponCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateCouponCollectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCouponCollection** | [**optional.Interface of UpdateCouponCollection**](UpdateCouponCollection.md)| Values to update the coupon collection | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

