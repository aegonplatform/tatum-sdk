/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

// Additional attributes for the future type.
type CreateTradeAttr struct {
	// Time in UTC when the future will be filled.
	SealDate float64 `json:"sealDate"`
	// Percentage of the future amount which selling or buying account must have available on future creation. This amount will be blocked till future is filled or expires.
	PercentBlock *Object `json:"percentBlock,omitempty"`
	// If one of the parties dont have filled full amount of the future at the time of expiration, the party will be penalized. Penalty is sent to opposite party in exchanged crypto.
	PercentPenalty *Object `json:"percentPenalty,omitempty"`
}
