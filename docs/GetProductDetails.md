# GetProductDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Product ID for which you requested the details | [default to null]
**Name** | **string** | Name of the product for which you requested the details | [default to null]
**CreatedAt** | **string** | Creation UTC date-time of the product (YYYY-MM-DDTHH:mm:ss.SSSZ) | [default to null]
**ModifiedAt** | **string** | Last modification UTC date-time of the product (YYYY-MM-DDTHH:mm:ss.SSSZ) | [default to null]
**Url** | **string** | URL to the product | [optional] [default to null]
**ImageUrl** | **string** | Absolute URL to the cover image of the product | [optional] [default to null]
**Sku** | **string** | Product identifier from the shop | [optional] [default to null]
**Price** | **float32** | Price of the product | [optional] [default to null]
**Categories** | **[]string** | Category ID-s of the product | [optional] [default to null]
**ParentId** | **string** | Parent product id of the product | [optional] [default to null]
**S3Original** | **string** | S3 url of original image | [optional] [default to null]
**S3ThumbAnalytics** | **string** | S3 thumbnail url of original image in 120x120 dimension for analytics section | [default to null]
**MetaInfo** | [**map[string]interface{}**](interface{}.md) | Meta data of product such as description, vendor, producer, stock level, etc. | [optional] [default to null]
**S3ThumbEditor** | **string** | S3 thumbnail url of original image in 600x400 dimension for editor section | [default to null]
**IsDeleted** | **bool** | product deleted from the shop&#39;s database | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


