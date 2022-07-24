package iso20022

// Account and holding of the next sub-level (Level 1).
type AccountSubLevel11 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-account on behalf of the account owner, that is, manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 1),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 2.
	AccountSubLevel2 []*AccountSubLevel12 `xml:"AcctSubLvl2,omitempty"`

	// Difference in holdings between the sub-account at level 1 and sub-accounts of level 2.
	AccountSubLevel2Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl2Diff,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel11) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel11) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel11) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel11) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel11) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel11) AddAccountSubLevel2() *AccountSubLevel12 {
	newValue := new(AccountSubLevel12)
	a.AccountSubLevel2 = append(a.AccountSubLevel2, newValue)
	return newValue
}

func (a *AccountSubLevel11) AddAccountSubLevel2Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	a.AccountSubLevel2Difference = append(a.AccountSubLevel2Difference, newValue)
	return newValue
}

func (a *AccountSubLevel11) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
