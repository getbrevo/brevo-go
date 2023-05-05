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
[**GetCategories**](EcommerceApi.md#GetCategories) | **Get** /categories | Return all your categories
[**GetCategoryInfo**](EcommerceApi.md#GetCategoryInfo) | **Get** /categories/{id} | Get a category details
[**GetProductInfo**](EcommerceApi.md#GetProductInfo) | **Get** /products/{id} | Get a product&#39;s details
[**GetProducts**](EcommerceApi.md#GetProducts) | **Get** /products | Return all your products


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

[**CreateUpdateBatchProductsModel**](createUpdateBatchProductsModel.md)

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

[**CreateCategoryModel**](createCategoryModel.md)

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

[**CreateProductModel**](createProductModel.md)

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

### Return type

[**GetCategories**](getCategories.md)

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

[**GetCategoryDetails**](getCategoryDetails.md)

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

[**GetProductDetails**](getProductDetails.md)

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

### Return type

[**GetProducts**](getProducts.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

