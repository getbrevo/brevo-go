# GetInboundEmailEventsByUuid

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceivedAt** | [**time.Time**](time.Time.md) | Date when email was received on SMTP relay | [optional] [default to null]
**DeliveredAt** | [**time.Time**](time.Time.md) | Date when email was delivered successfully to client’s webhook | [optional] [default to null]
**Recipient** | **string** | Recipient’s email address | [optional] [default to null]
**Sender** | **string** | Sender’s email address | [optional] [default to null]
**MessageId** | **string** | Value of the Message-ID header. This will be present only after the processing is done. | [optional] [default to null]
**Subject** | **string** | Value of the Subject header. This will be present only after the processing is done.  | [optional] [default to null]
**Attachments** | [**[]GetInboundEmailEventsByUuidAttachments**](getInboundEmailEventsByUuid_attachments.md) | List of attachments of the email. This will be present only after the processing is done. | [optional] [default to null]
**Logs** | [**[]GetInboundEmailEventsByUuidLogs**](getInboundEmailEventsByUuid_logs.md) | List of events/logs that describe the lifecycle of the email on SIB platform | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


