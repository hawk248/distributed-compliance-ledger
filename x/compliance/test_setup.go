package compliance

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/test_constants"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestSetup struct {
	Cdc     *amino.Codec
	Ctx     sdk.Context
	App     AppModule
	Handler sdk.Handler
	Querier sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	complianceKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(complianceKey, sdk.StoreTypeIAVL, db)

	authzKey := sdk.NewKVStoreKey(authz.StoreKey)
	dbStore.MountStoreWithDB(authzKey, sdk.StoreTypeIAVL, db)

	dbStore.LoadLatestVersion()

	// Init Keepers
	complianceKeeper := NewKeeper(complianceKey, cdc)
	authzKeeper := authz.NewKeeper(authzKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: "zbl-test-chain-id"}, false, log.NewNopLogger())

	app := NewAppModule(complianceKeeper, authzKeeper)

	// Create Handler and Querier
	querier := app.NewQuerierHandler()
	handler := app.NewHandler()

	setup := TestSetup{
		Cdc:     cdc,
		Ctx:     ctx,
		App:     app,
		Handler: handler,
		Querier: querier,
	}

	return setup
}

func (setup TestSetup) Manufacturer() sdk.AccAddress {
	acc, _ := sdk.AccAddressFromBech32("cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz")
	setup.App.authzKeeper.AssignRole(setup.Ctx, acc, authz.Manufacturer)
	return acc
}

func TestMsgAddModelInfo(owner sdk.AccAddress) MsgAddModelInfo  {
	return MsgAddModelInfo{
		ID:                       test_constants.Id,
		Name:                     test_constants.Name,
		Owner:                    owner,
		Description:              test_constants.Description,
		SKU:                      test_constants.Sku,
		FirmwareVersion:          test_constants.FirmwareVersion,
		HardwareVersion:          test_constants.HardwareVersion,
		CertificateID:            test_constants.CertificateID,
		CertifiedDate:            test_constants.CertifiedDate,
		TisOrTrpTestingCompleted: false,
		Signer:                   owner,
	}
}