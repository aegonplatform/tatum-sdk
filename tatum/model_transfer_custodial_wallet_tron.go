/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type TransferCustodialWalletTron struct {
	// The blockchain to work with
	Chain string `json:"chain"`
	// The gas pump address that transfers the asset
	CustodialAddress string `json:"custodialAddress"`
	// The blockchain address that receives the asset
	Recipient string `json:"recipient"`
	// The type of the asset to transfer. Set <code>0</code> for fungible tokens (ERC-20 or equivalent), <code>1</code> for NFTs (ERC-721 or equivalent), or <code>3</code> for native blockchain currencies.
	ContractType float64 `json:"contractType"`
	// (Only if the asset is a fungible token or NFT) The address of the token to transfer. Do not use if the asset is a native blockchain currency.
	TokenAddress string `json:"tokenAddress,omitempty"`
	// (Only if the asset is a fungible token or native blockchain currency) The amount of the asset to transfer. Do not use if the asset is an NFT.
	Amount string `json:"amount,omitempty"`
	// (Only if the asset is an NFT) The ID of the token to transfer. Do not use if the asset is a fungible token or native blockchain currency.
	TokenId string `json:"tokenId,omitempty"`
	// The private key of the blockchain address that owns the gas pump address (\"master address\")
	FromPrivateKey string `json:"fromPrivateKey"`
	// The maximum amount to be paid as the gas fee (in TRX)
	FeeLimit float64 `json:"feeLimit"`
}
