package iso20022

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection37 struct {

	// Indicates whether the net proceeds include interest accrued on the financial instrument.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Whether the net proceeds include stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Indicates whether the net proceeds include brokerage fees for the transaction. If absent, element is not required.
	BrokerageAmountIndicator *YesNoIndicator `xml:"BrkrgAmtInd,omitempty"`

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt,omitempty"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Date and time at which the cash is at the disposal of the credit account owner, or ceases to be at the disposal of the debit account owner.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`
}

func (a *AmountAndDirection37) SetAccruedInterestIndicator(value string) {
	a.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection37) SetStampDutyIndicator(value string) {
	a.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection37) SetBrokerageAmountIndicator(value string) {
	a.BrokerageAmountIndicator = (*YesNoIndicator)(&value)
}

func (a *AmountAndDirection37) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection37) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection37) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection37) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return a.ForeignExchangeDetails
}

func (a *AmountAndDirection37) AddValueDate() *DateAndDateTimeChoice {
	a.ValueDate = new(DateAndDateTimeChoice)
	return a.ValueDate
}
