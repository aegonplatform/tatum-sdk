/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type BurnErc20Celo struct {
	// Amount of tokens to be destroyed.
	Amount string `json:"amount"`
	// Address of ERC721 token
	ContractAddress string `json:"contractAddress"`
	// Private key of sender address. Private key, or signature Id must be present.
	FromPrivateKey string `json:"fromPrivateKey"`
	// Nonce to be set to Celo transaction. If not present, last known nonce will be used.
	Nonce float64 `json:"nonce,omitempty"`
	// Currency to pay for transaction gas
	FeeCurrency string `json:"feeCurrency"`
}
