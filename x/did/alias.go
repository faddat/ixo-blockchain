package did

import (
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
)

const (
	ModuleName   = types.ModuleName
	//QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
)

type (
	Keeper       = keeper.Keeper
	//GenesisState = types.GenesisState

	Did           = exported.Did
	DidCredential = types.DidCredential
	//DidDoc        = exported.DidDoc
	//IxoDid        = exported.IxoDid

	//MsgAddDid        = types.MsgAddDid
	//MsgAddCredential = types.MsgAddCredential

	//MsgServer = types.MsgServer
)

var (
	// function aliases
	NewKeeper     = keeper.NewKeeper
	//NewQuerier    = keeper.NewQuerier
	//RegisterCodec = types.RegisterCodec
	//RegisterLegacyAminoCodec = types.RegisterLegacyAminoCodec
	//RegisterInterfaces = types.RegisterInterfaces

	//NewGenesisState     = types.NewGenesisState
	//DefaultGenesisState = types.DefaultGenesisState
	//ValidateGenesis     = types.ValidateGenesis
	//RegisterQueryHandlerClient = types.RegisterQueryHandlerClient
	//NewQueryClient      = types.NewQueryClient

	//VerifyKeyToAddr = exported.VerifyKeyToAddr

	IsValidDid      = types.IsValidDid
	//IsValidPubKey   = types.IsValidPubKey
	UnmarshalIxoDid = types.UnmarshalIxoDid

	// variable aliases
	//ModuleCdc = types.ModuleCdc

	ErrInvalidDid        = types.ErrInvalidDid
	//ErrInvalidPubKey     = types.ErrInvalidPubKey
	//ErrDidPubKeyMismatch = types.ErrDidPubKeyMismatch
)