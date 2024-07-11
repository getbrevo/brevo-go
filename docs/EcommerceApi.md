# \EcommerceApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBatchOrder**](EcommerceApi.md#CreateBatchOrder) | **Post** /orders/status/batch | Create orders in batch
[**CreateOrder**](EcommerceApi.md#CreateOrder) | **Post** /orders/status | Managing the status of the order
[**CreateUpdateBatchCategory**](EcommerceApi.md#CreateUpdateBatchCategory) | **Post** /categories/batch | Create categories in batch
[**CreateUpdateBatchProducts**](EcommerceApi.md#CreateUpdateBatchProducts) | **Post** /products/batch | Create products in batch
[**CreateUpdateCategory**](EcommerceApi.md#CreateUpdateCategory) | **Post** /categories | Create/Update a category
[**CreateUpdateProduct**](EcommerceApi.md#CreateUpdateProduct) | **Post** /products | Create/Update a product
[**EcommerceActivatePost**](EcommerceApi.md#EcommerceActivatePost) | **Post** /ecommerce/activate | Activate the eCommerce app
[**EcommerceAttributionMetricsConversionSourceConversionSourceIdGet**](EcommerceApi.md#EcommerceAttributionMetricsConversionSourceConversionSourceIdGet) | **Get** /ecommerce/attribution/metrics/{conversionSource}/{conversionSourceId} | Get detailed attribution metrics for a single Brevo campaign
[**EcommerceAttributionMetricsGet**](EcommerceApi.md#EcommerceAttributionMetricsGet) | **Get** /ecommerce/attribution/metrics | Get attribution metrics for one or more Brevo campaigns
[**EcommerceAttributionProductsConversionSourceConversionSourceIdGet**](EcommerceApi.md#EcommerceAttributionProductsConversionSourceConversionSourceIdGet) | **Get** /ecommerce/attribution/products/{conversionSource}/{conversionSourceId} | Get attributed product sales for a single Brevo campaign
[**EcommerceConfigDisplayCurrencyGet**](EcommerceApi.md#EcommerceConfigDisplayCurrencyGet) | **Get** /ecommerce/config/displayCurrency | Get the ISO 4217 compliant display currency code for your Brevo account
[**GetCategories**](EcommerceApi.md#GetCategories) | **Get** /categories | Return all your categories
[**GetCategoryInfo**](EcommerceApi.md#GetCategoryInfo) | **Get** /categories/{id} | Get a category details
[**GetOrders**](EcommerceApi.md#GetOrders) | **Get** /orders | Get order details
[**GetProductInfo**](EcommerceApi.md#GetProductInfo) | **Get** /products/{id} | Get a product&#39;s details
[**GetProducts**](EcommerceApi.md#GetProducts) | **Get** /products | Return all your products
[**SetConfigDisplayCurrency**](EcommerceApi.md#SetConfigDisplayCurrency) | **Post** /ecommerce/config/displayCurrency | Set the ISO 4217 compliant display currency code for your Brevo account


# **CreateBatchOrder**
> CreateBatchOrder(ctx, orderBatch)
Create orders in batch

Create multiple orders at one time instead of one order at a time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderBatch** | [**OrderBatch**](OrderBatch.md)|  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrder**
> CreateOrder(ctx, order)
Managing the status of the order

Manages the transactional status of the order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **order** | [**Order**](Order.md)|  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUpdateBatchCategory**
> CreateUpdateBatchCategoryModel CreateUpdateBatchCategory(ctx, createUpdateBatchCategory)
Create categories in batch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createUpdateBatchCategory** | [**CreateUpdateBatchCategory**](CreateUpdateBatchCategory.md)| Values to create a batch of categories | 

### Return type

[**CreateUpdateBatchCategoryModel**](createUpdateBatchCategoryModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUpdateBatchProducts**
> CreateUpdateBatchProductsModel CreateUpdateBatchProducts(ctx, createUpdateBatchProducts)
Create products in batch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createUpdateBatchProducts** | [**CreateUpdateBatchProducts**](CreateUpdateBatchProducts.md)| Values to create a batch of products | 

### Return type

[**CreateUpdateBatchProductsModel**](CreateUpdateBatchProductsModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUpdateCategory**
> CreateCategoryModel CreateUpdateCategory(ctx, createUpdateCategory)
Create/Update a category

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createUpdateCategory** | [**CreateUpdateCategory**](CreateUpdateCategory.md)| Values to create/update a category | 

### Return type

[**CreateCategoryModel**](CreateCategoryModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUpdateProduct**
> CreateProductModel CreateUpdateProduct(ctx, createUpdateProduct)
Create/Update a product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createUpdateProduct** | [**CreateUpdateProduct**](CreateUpdateProduct.md)| Values to create/update a product | 

### Return type

[**CreateProductModel**](CreateProductModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EcommerceActivatePost**
> EcommerceActivatePost(ctx, )
Activate the eCommerce app

Getting access to Brevo eCommerce.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EcommerceAttributionMetricsConversionSourceConversionSourceIdGet**
> InlineResponse2006 EcommerceAttributionMetricsConversionSourceConversionSourceIdGet(ctx, conversionSource, conversionSourceId)
Get detailed attribution metrics for a single Brevo campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **conversionSource** | **string**| The Brevo campaign type for which data will be retrieved | 
  **conversionSourceId** | **float32**| The Brevo campaign id for which data will be retrieved | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EcommerceAttributionMetricsGet**
> InlineResponse2005 EcommerceAttributionMetricsGet(ctx, optional)
Get attribution metrics for one or more Brevo campaigns

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EcommerceAttributionMetricsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceAttributionMetricsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **periodFrom** | **optional.Time**| When getting metrics for a specific period, define the starting datetime in RFC3339 format | 
 **periodTo** | **optional.Time**| When getting metrics for a specific period, define the end datetime in RFC3339 format | 
 **emailCampaignId** | [**optional.Interface of []float32**](float32.md)| The email campaign id(s) to get metrics for | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EcommerceAttributionProductsConversionSourceConversionSourceIdGet**
> InlineResponse2007 EcommerceAttributionProductsConversionSourceConversionSourceIdGet(ctx, conversionSource, conversionSourceId)
Get attributed product sales for a single Brevo campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **conversionSource** | **string**| The Brevo campaign type for which data will be retrieved | 
  **conversionSourceId** | **float32**| The Brevo campaign id for which data will be retrieved | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EcommerceConfigDisplayCurrencyGet**
> InlineResponse2004 EcommerceConfigDisplayCurrencyGet(ctx, )
Get the ISO 4217 compliant display currency code for your Brevo account

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategories**
> GetCategories GetCategories(ctx, optional)
Return all your categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCategoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document in the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]
 **ids** | [**optional.Interface of []string**](string.md)| Filter by category ids | 
 **name** | **optional.String**| Filter by category name | 
 **modifiedSince** | **optional.String**| Filter (urlencoded) the categories modified after a given UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ). **Prefer to pass your timezone in date-time format for accurate result.**  | 
 **createdSince** | **optional.String**| Filter (urlencoded) the categories created after a given UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ). **Prefer to pass your timezone in date-time format for accurate result.**  | 

### Return type

[**GetCategories**](GetCategories.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategoryInfo**
> GetCategoryDetails GetCategoryInfo(ctx, id)
Get a category details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Category ID | 

### Return type

[**GetCategoryDetails**](GetCategoryDetails.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrders**
> GetOrders GetOrders(ctx, optional)
Get order details

Get all the orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document in the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]
 **modifiedSince** | **optional.String**| Filter (urlencoded) the orders modified after a given UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ). **Prefer to pass your timezone in date-time format for accurate result.**  | 
 **createdSince** | **optional.String**| Filter (urlencoded) the orders created after a given UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ). **Prefer to pass your timezone in date-time format for accurate result.**  | 

### Return type

[**GetOrders**](GetOrders.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductInfo**
> GetProductDetails GetProductInfo(ctx, id)
Get a product's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Product ID | 

### Return type

[**GetProductDetails**](GetProductDetails.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProducts**
> GetProducts GetProducts(ctx, optional)
Return all your products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetProductsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document in the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]
 **ids** | [**optional.Interface of []string**](string.md)| Filter by product ids | 
 **name** | **optional.String**| Filter by product name, minimum 3 characters should be present for search | 
 **priceLte** | **optional.Float32**| Price filter for products less than and equals to particular amount | 
 **priceGte** | **optional.Float32**| Price filter for products greater than and equals to particular amount | 
 **priceLt** | **optional.Float32**| Price filter for products less than particular amount | 
 **priceGt** | **optional.Float32**| Price filter for products greater than particular amount | 
 **priceEq** | **optional.Float32**| Price filter for products equals to particular amount | 
 **priceNe** | **optional.Float32**| Price filter for products not equals to particular amount | 
 **categories** | [**optional.Interface of []string**](string.md)| Filter by category ids | 
 **modifiedSince** | **optional.String**| Filter (urlencoded) the orders modified after a given UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ). **Prefer to pass your timezone in date-time format for accurate result.**  | 
 **createdSince** | **optional.String**| Filter (urlencoded) the orders created after a given UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ). **Prefer to pass your timezone in date-time format for accurate result.**  | 

### Return type

[**GetProducts**](GetProducts.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetConfigDisplayCurrency**
> SetConfigDisplayCurrency SetConfigDisplayCurrency(ctx, setConfigDisplayCurrency)
Set the ISO 4217 compliant display currency code for your Brevo account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **setConfigDisplayCurrency** | [**SetConfigDisplayCurrency**](SetConfigDisplayCurrency.md)| set ISO 4217 compliant display currency code payload | 

### Return type

[**SetConfigDisplayCurrency**](SetConfigDisplayCurrency.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

