/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type LtcUtxo struct {
	// Version of the UTXO.
	Version float64 `json:"version,omitempty"`
	Height float64 `json:"height,omitempty"`
	// Amount of UTXO in satoshis.
	Value float64 `json:"value,omitempty"`
	// Data generated by a spender which is almost always used as variables to satisfy a pubkey script.
	Script string `json:"script,omitempty"`
	// Address of the owner of the UTXO.
	Address string `json:"address,omitempty"`
	// Coinbase transaction - miner fee.
	Coinbase bool `json:"coinbase,omitempty"`
	// Transaction hash.
	Hash string `json:"hash,omitempty"`
	// Transaction index of the output.
	Index float64 `json:"index,omitempty"`
}
