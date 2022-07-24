package iso20022

// Specifies additional parameters to the message or transaction.
type AdditionalParameters3 struct {

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement1Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`
}

func (a *AdditionalParameters3) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement1Code)(&value)
}

func (a *AdditionalParameters3) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters3) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters3) SetAccountServicerTransactionIdentification(value string) {
	a.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters3) SetPoolIdentification(value string) {
	a.PoolIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters3) SetCorporateActionEventIdentification(value string) {
	a.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters3) SetMarketInfrastructureTransactionIdentification(value string) {
	a.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}
