package iso20022

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection48 struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms23 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection48) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection48) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection48) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection48) AddForeignExchangeDetails() *ForeignExchangeTerms23 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms23)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection48) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}
