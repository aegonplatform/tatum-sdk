/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type TransferXlmKms struct {
	// Sender account ID
	SenderAccountId string `json:"senderAccountId"`
	// Blockchain account to send from
	FromAccount string `json:"fromAccount"`
	// Blockchain address to send assets
	Address string `json:"address"`
	// Amount to be sent, in XLM or XLM-based Asset.
	Amount string `json:"amount"`
	// Identifier of the secret associated in signing application. Secret, or signature Id must be present.
	SignatureId string `json:"signatureId"`
	// Compliance check, if withdrawal is not compliant, it will not be processed.
	Compliant bool `json:"compliant,omitempty"`
	// Short message to recipient. Usually used as an account discriminator. It can be either 28 characters long ASCII text, 64 characters long HEX string or uint64 number. When using as an account disciminator in Tatum Offchain ledger, can be in format of destination_acc|source_acc.
	Attr string `json:"attr,omitempty"`
	// Identifier of the payment, shown for created Transaction within Tatum sender account.
	PaymentId string `json:"paymentId,omitempty"`
	// Note visible to owner of withdrawing account.
	SenderNote string `json:"senderNote,omitempty"`
}