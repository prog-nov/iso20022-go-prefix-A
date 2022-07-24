package iso20022

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection57 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`
}

func (a *AmountAndDirection57) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection57) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}
