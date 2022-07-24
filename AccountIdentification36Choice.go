package iso20022

// Choice between all accounts (GENR - General in ISO 15022) or one or more selected accounts and balance information.
type AccountIdentification36Choice struct {

	// All safekeeping accounts that own underlying financial instrument.
	ForAllAccounts *AccountIdentification10 `xml:"ForAllAccts"`

	// Selected safekeeping accounts list (and optionally balance information) to which the corporate action event applies.
	AccountsListAndBalanceDetails []*AccountAndBalance39 `xml:"AcctsListAndBalDtls"`
}

func (a *AccountIdentification36Choice) AddForAllAccounts() *AccountIdentification10 {
	a.ForAllAccounts = new(AccountIdentification10)
	return a.ForAllAccounts
}

func (a *AccountIdentification36Choice) AddAccountsListAndBalanceDetails() *AccountAndBalance39 {
	newValue := new(AccountAndBalance39)
	a.AccountsListAndBalanceDetails = append(a.AccountsListAndBalanceDetails, newValue)
	return newValue
}
