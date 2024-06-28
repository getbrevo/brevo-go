# CreateWhatsAppTemplate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the template | [default to null]
**Language** | **string** | Language of the template. For Example : **en** for English  | [default to null]
**Category** | **string** | Category of the template | [default to null]
**MediaUrl** | **string** | Absolute url of the media file **(no local file)** for the header. **Use this field in you want to add media in Template header and headerText is empty.** Allowed extensions for media files are: #### jpeg | png | mp4 | pdf  | [optional] [default to null]
**BodyText** | **string** | Body of the template. **Maximum allowed characters are 1024** | [default to null]
**HeaderText** | **string** | Text content of the header in the template.  **Maximum allowed characters are 45** **Use this field to add text content in template header and if mediaUrl is empty**  | [optional] [default to null]
**Source** | **string** | source of the template | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


