# GetSmtpTemplateOverview

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the template | [default to null]
**Name** | **string** | Name of the template | [default to null]
**Subject** | **string** | Subject of the template | [default to null]
**IsActive** | **bool** | Status of template (true&#x3D;active, false&#x3D;inactive) | [default to null]
**TestSent** | **bool** | Status of test sending for the template (true&#x3D;test email has been sent, false&#x3D;test email has not been sent) | [default to null]
**Sender** | [***GetSmtpTemplateOverviewSender**](GetSmtpTemplateOverviewSender.md) |  | [default to null]
**ReplyTo** | **string** | Email defined as the \&quot;Reply to\&quot; for the template | [default to null]
**ToField** | **string** | Customisation of the \&quot;to\&quot; field for the template | [default to null]
**Tag** | **string** | Tag of the template | [default to null]
**HtmlContent** | **string** | HTML content of the template | [default to null]
**CreatedAt** | **string** | Creation UTC date-time of the template (YYYY-MM-DDTHH:mm:ss.SSSZ) | [default to null]
**ModifiedAt** | **string** | Last modification UTC date-time of the template (YYYY-MM-DDTHH:mm:ss.SSSZ) | [default to null]
**DoiTemplate** | **bool** | It is true if template is a valid Double opt-in (DOI) template, otherwise it is false. This field will be available only in case of single template detail call. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


