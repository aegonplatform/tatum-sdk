/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type CreateAccountXpub struct {
	// Account currency. Supported values are BTC, BNB, LTC, DOGE, BCH, ETH, XLM, XRP, TRON, BSC, SOL, MATIC, ALGO, KCS, EGLD, CELO, KLAY, XDC, FLOW, Tatum virtual currencies started with the \"VC_\" prefix (this includes fiat currencies), USDT, WBTC, LEO, LINK, GMC, UNI, FREE, MKR, USDC, BAT, TUSD, BUSD, PAX, PAXG, MMY, XCON, USDT_TRON, BETH, BUSD, BBTC, BADA, WBNB, BDOT, BXRP, BLTC, BBCH, CAKE, BUSD_BSC, ERC-20, BEP-20 or TRC-10/20 custom tokens registered on the Tatum platform, XLM or XRP assets created via the Tatum platform. ERC-20 tokens and BEP-20 tokens do not have testnet blockchains, therefore you cannot use them in a non-production environment. You can emulate a testnet environment by <a href=\"https://apidoc.tatum.io/tag/Blockchain-operations#operation/registerErc20Token\" target=\"_blank\">registering a custom ERC-20 or BEP-20 token</a> on the Tatum  platform and then minting some tokens from the token's address using the <a href=\"https://erc20faucet.com/\" target=\"_blank\">Ethereum ERC-20 Token Faucet</a>. 
	Currency string `json:"currency"`
	// Extended public key to generate addresses from.
	Xpub string `json:"xpub"`
	Customer *CustomerRegistration `json:"customer,omitempty"`
	// Enable compliant checks. If this is enabled, it is impossible to create account if compliant check fails.
	Compliant bool `json:"compliant,omitempty"`
	// For bookkeeping to distinct account purpose.
	AccountCode string `json:"accountCode,omitempty"`
	// All transaction will be accounted in this currency for all accounts. Currency can be overridden per account level. If not set, customer accountingCurrency is used or EUR by default. ISO-4217
	AccountingCurrency string `json:"accountingCurrency,omitempty"`
	// Account number from external system.
	AccountNumber string `json:"accountNumber,omitempty"`
}
