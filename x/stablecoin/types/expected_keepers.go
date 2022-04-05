package types // noalias

import (
	pftypes "github.com/MatrixDao/matrix/x/pricefeed/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	SetAccount(sdk.Context, types.AccountI)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	// Methods imported from bank should be defined here
}

type PriceKeeper interface {
	GetCurrentPrice(sdk.Context, string) (pftypes.CurrentPrice, error)
	GetCurrentPrices(ctx sdk.Context) pftypes.CurrentPrices
	GetRawPrices(ctx sdk.Context, marketId string) pftypes.PostedPrices
	GetMarket(ctx sdk.Context, marketID string) (pftypes.Market, bool)
	GetMarkets(ctx sdk.Context) pftypes.Markets
	GetOracle(ctx sdk.Context, marketID string, address sdk.AccAddress) (sdk.AccAddress, error)
	GetOracles(ctx sdk.Context, marketID string) ([]sdk.AccAddress, error)
	SetCurrentPrices(ctx sdk.Context, marketID string) error
}
