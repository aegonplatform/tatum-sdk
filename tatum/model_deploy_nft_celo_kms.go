/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type DeployNftCeloKms struct {
	// Chain to work with.
	Chain string `json:"chain"`
	// Name of the NFT token
	Name string `json:"name"`
	// Symbol of the NFT token
	Symbol string `json:"symbol"`
	// If signatureId is mnemonic-based, this is the index to the specific address from that mnemonic.
	Index float64 `json:"index,omitempty"`
	// Identifier of the private key associated in signing application. Private key, or signature Id must be present.
	SignatureId string `json:"signatureId"`
	Nonce float64 `json:"nonce,omitempty"`
	// True if the contract is provenance percentage royalty type. False by default. <a href=\"https://github.com/tatumio/smart-contracts\" target=\"_blank\">Details and sources avaiable here.</a>
	Provenance bool `json:"provenance,omitempty"`
	// True if the contract is fixed price royalty type. False by default. <a href=\"https://github.com/tatumio/smart-contracts\" target=\"_blank\">Details and sources avaiable here.</a>
	Cashback bool `json:"cashback,omitempty"`
	// True if the contract is publicMint type. False by default.
	PublicMint bool `json:"publicMint,omitempty"`
	// Currency to pay for transaction gas
	FeeCurrency string `json:"feeCurrency"`
}
