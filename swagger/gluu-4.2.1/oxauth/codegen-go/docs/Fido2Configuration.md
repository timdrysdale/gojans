# Fido2Configuration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The version of the FIDO2 U2F core protocol to which this server conforms. The value MUST be the string 1.0. | [default to null]
**Issuer** | **string** | A URI indicating the party operating the FIDO U2F server. | [default to null]
**Attestation** | [**[]Fido2ConfigurationAttestation**](FIDO2 configuration_attestation.md) |  | [default to null]
**Assertion** | [**[]Fido2ConfigurationAssertion**](FIDO2 configuration_assertion.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

