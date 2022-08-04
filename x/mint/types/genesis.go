package types

// this line is used by starport scaffolding # genesis/types/import

// NewGenesisState creates a new GenesisState object
func NewGenesisState(minter Minter, params Params) *GenesisState {
	return &GenesisState{
		Minter: minter,
		Params: params,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Minter: DefaultInitialMinter(),
		Params: DefaultParams(),
	}
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}

	return ValidateMinter(data.Minter)
}
