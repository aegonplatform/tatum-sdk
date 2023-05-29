/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type DeployErc20OffchainMnemonicAddress struct {
	// Name of the ERC20 token - stored as a symbol on Blockchain
	Symbol string `json:"symbol"`
	// max supply of ERC20 token.
	Supply string `json:"supply"`
	// Description of the ERC20 token
	Description string `json:"description"`
	// Base pair for ERC20 token. 1 token will be equal to 1 unit of base pair. Transaction value will be calculated according to this base pair.
	BasePair string `json:"basePair"`
	// Exchange rate of the base pair. Each unit of the created curency will represent value of baseRate*1 basePair.
	BaseRate float64 `json:"baseRate,omitempty"`
	Customer *CustomerRegistration `json:"customer,omitempty"`
	// Address on Ethereum blockchain, where all initial supply will be stored. Either xpub and derivationIndex, or address must be present, not both.
	Address string `json:"address"`
	// Mnemonic to generate private key for the deploy account of ERC20, from which the gas will be paid (index will be used). If address is not present, mnemonic is used to generate xpub and index is set to 1. Either mnemonic and index or privateKey and address must be present, not both.
	Mnemonic string `json:"mnemonic"`
	// derivation index of address to pay for deployment of ERC20
	Index int32 `json:"index"`
	// Nonce to be set to Ethereum transaction. If not present, last known nonce will be used.
	Nonce float64 `json:"nonce,omitempty"`
}
