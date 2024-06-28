# SendSmtpEmailMessageVersions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | [**[]SendSmtpEmailTo1**](sendSmtpEmail_to_1.md) | List of email addresses and names (_optional_) of the recipients. For example, [{\&quot;name\&quot;:\&quot;Jimmy\&quot;, \&quot;email\&quot;:\&quot;jimmy98@example.com\&quot;}, {\&quot;name\&quot;:\&quot;Joe\&quot;, \&quot;email\&quot;:\&quot;joe@example.com\&quot;}] | [default to null]
**Params** | [**map[string]interface{}**](interface{}.md) | Pass the set of attributes to customize the template. For example, {\&quot;FNAME\&quot;:\&quot;Joe\&quot;, \&quot;LNAME\&quot;:\&quot;Doe\&quot;}. It&#39;s considered only if template is in New Template Language format. | [optional] [default to null]
**Bcc** | [**[]SendSmtpEmailBcc**](sendSmtpEmail_bcc.md) | List of email addresses and names (optional) of the recipients in bcc | [optional] [default to null]
**Cc** | [**[]SendSmtpEmailCc**](sendSmtpEmail_cc.md) | List of email addresses and names (optional) of the recipients in cc | [optional] [default to null]
**ReplyTo** | [***SendSmtpEmailReplyTo1**](sendSmtpEmail_replyTo_1.md) |  | [optional] [default to null]
**Subject** | **string** | Custom subject specific to message version  | [optional] [default to null]
**HtmlContent** | **string** | HTML body of the message. **Mandatory if &#39;templateId&#39; is not passed, ignored if &#39;templateId&#39; is passed**  | [optional] [default to null]
**TextContent** | **string** | Plain Text body of the message. **Ignored if &#39;templateId&#39; is passed**  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


