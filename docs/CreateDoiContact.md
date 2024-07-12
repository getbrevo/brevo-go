# CreateDoiContact

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address where the confirmation email will be sent. This email address will be the identifier for all other contact attributes. | [default to null]
**Attributes** | [**map[string]interface{}**] | Pass the set of attributes and their values. These attributes must be present in your Brevo account. For eg. {&#39;FNAME&#39;:&#39;Elly&#39;, &#39;LNAME&#39;:&#39;Roger&#39;, &#39;COUNTRIES&#39;:[&#39;India&#39;,&#39;China&#39;]} | [optional] [default to null]
**IncludeListIds** | **[]int64** | Lists under user account where contact should be added | [default to null]
**ExcludeListIds** | **[]int64** | Lists under user account where contact should not be added | [optional] [default to null]
**TemplateId** | **int64** | Id of the Double opt-in (DOI) template | [default to null]
**RedirectionUrl** | **string** | URL of the web page that user will be redirected to after clicking on the double opt in URL. When editing your DOI template you can reference this URL by using the tag {{ params.DOIurl }}. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


