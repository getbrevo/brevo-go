# CreateUpdateProduct

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Product ID for which you requested the details | [default to null]
**Name** | **string** | Mandatory in case of creation**. Name of the product for which you requested the details | [default to null]
**Url** | **string** | URL to the product | [optional] [default to null]
**ImageUrl** | **string** | Absolute URL to the cover image of the product | [optional] [default to null]
**Sku** | **string** | Product identifier from the shop | [optional] [default to null]
**Price** | **float32** | Price of the product | [optional] [default to null]
**Categories** | **[]string** | Category ID-s of the product | [optional] [default to null]
**ParentId** | **string** | Parent product id of the product | [optional] [default to null]
**MetaInfo** | **map[string]string** | Meta data of product such as description, vendor, producer, stock level. The size of cumulative metaInfo shall not exceed **1000 KB**. Maximum length of metaInfo object can be 10. | [optional] [default to null]
**UpdateEnabled** | **bool** | Facilitate to update the existing category in the same request (updateEnabled &#x3D; true) | [optional] [default to null]
**DeletedAt** | **string** | UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ) of the product deleted from the shop&#39;s database | [optional] [default to null]
**IsDeleted** | **bool** | product deleted from the shop&#39;s database | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


