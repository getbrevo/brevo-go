# UpdateAttribute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Value of the attribute to update. Use only if the attribute&#39;s category is &#39;calculated&#39; or &#39;global&#39; | [optional] [default to null]
**Enumeration** | [**[]UpdateAttributeEnumeration**](UpdateAttributeEnumeration.md) | List of the values and labels that the attribute can take. Use only if the attribute&#39;s category is \&quot;category\&quot;. For example, [{\&quot;value\&quot;:1, \&quot;label\&quot;:\&quot;male\&quot;}, {\&quot;value\&quot;:2, \&quot;label\&quot;:\&quot;female\&quot;}] | [optional] [default to null]
**MultiCategoryOptions** | **[]string** | Use this option to add multiple-choice attributes options only if the attribute&#39;s category is \&quot;normal\&quot;. **This option is specifically designed for updating multiple-choice attributes**. For example: **[\&quot;USA\&quot;,\&quot;INDIA\&quot;]**  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


