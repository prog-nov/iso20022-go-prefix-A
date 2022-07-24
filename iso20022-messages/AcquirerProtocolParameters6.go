package iso20022

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters6 struct {

	// Identification of the acquirer using this protocol.
	AcquirerIdentification []*GenericIdentification32 `xml:"AcqrrId"`

	// Identification of the payment application, user of the acquirer protocol.
	ApplicationIdentification []*Max35Text `xml:"ApplId,omitempty"`

	// Acquirer host configuration.
	Host []*AcquirerHostConfiguration2 `xml:"Hst,omitempty"`

	// Acquirer protocol parameters of transactions performing an online authorisation.
	OnLineTransaction *AcquirerProtocolParameters5 `xml:"OnLineTx,omitempty"`

	// Acquirer protocol parameters of transactions performing an offline authorisation.
	OffLineTransaction *AcquirerProtocolParameters5 `xml:"OffLineTx,omitempty"`

	// Configuration parameters of reconciliation exchanges.
	ReconciliationExchange *ExchangeConfiguration4 `xml:"RcncltnXchg,omitempty"`

	// Indicates the reconciliation period is assigned by the acquirer instead of the acceptor.
	ReconciliationByAcquirer *TrueFalseIndicator `xml:"RcncltnByAcqrr,omitempty"`

	// Indicates the reconciliation total amounts are computed per currency.
	TotalsPerCurrency *TrueFalseIndicator `xml:"TtlsPerCcy,omitempty"`

	// Indicates that totals in reconciliation or batch must be split per group of points of interaction and card product profiles when these informations are present in the transactions.
	SplitTotals *TrueFalseIndicator `xml:"SpltTtls,omitempty"`

	// Indicates that response messages and an AcceptorCompletionAdvice message following an authorisation exchange must contain protected or plain card data.
	CardDataVerification *TrueFalseIndicator `xml:"CardDataVrfctn,omitempty"`

	// Send a cancellation advice for offline transactions not yet captured.
	NotifyOffLineCancellation *TrueFalseIndicator `xml:"NtfyOffLineCxl,omitempty"`

	// Types of transaction to include in the batch.
	BatchTransferContent []*BatchTransactionType1Code `xml:"BtchTrfCntt,omitempty"`

	// Configuration of a message item.
	MessageItem []*MessageItemCondition1 `xml:"MsgItm,omitempty"`

	// Indicator to require protection of sensitive card data in messages.
	ProtectCardData *TrueFalseIndicator `xml:"PrtctCardData"`
}

func (a *AcquirerProtocolParameters6) AddAcquirerIdentification() *GenericIdentification32 {
	newValue := new(GenericIdentification32)
	a.AcquirerIdentification = append(a.AcquirerIdentification, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters6) AddApplicationIdentification(value string) {
	a.ApplicationIdentification = append(a.ApplicationIdentification, (*Max35Text)(&value))
}

func (a *AcquirerProtocolParameters6) AddHost() *AcquirerHostConfiguration2 {
	newValue := new(AcquirerHostConfiguration2)
	a.Host = append(a.Host, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters6) AddOnLineTransaction() *AcquirerProtocolParameters5 {
	a.OnLineTransaction = new(AcquirerProtocolParameters5)
	return a.OnLineTransaction
}

func (a *AcquirerProtocolParameters6) AddOffLineTransaction() *AcquirerProtocolParameters5 {
	a.OffLineTransaction = new(AcquirerProtocolParameters5)
	return a.OffLineTransaction
}

func (a *AcquirerProtocolParameters6) AddReconciliationExchange() *ExchangeConfiguration4 {
	a.ReconciliationExchange = new(ExchangeConfiguration4)
	return a.ReconciliationExchange
}

func (a *AcquirerProtocolParameters6) SetReconciliationByAcquirer(value string) {
	a.ReconciliationByAcquirer = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters6) SetTotalsPerCurrency(value string) {
	a.TotalsPerCurrency = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters6) SetSplitTotals(value string) {
	a.SplitTotals = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters6) SetCardDataVerification(value string) {
	a.CardDataVerification = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters6) SetNotifyOffLineCancellation(value string) {
	a.NotifyOffLineCancellation = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters6) AddBatchTransferContent(value string) {
	a.BatchTransferContent = append(a.BatchTransferContent, (*BatchTransactionType1Code)(&value))
}

func (a *AcquirerProtocolParameters6) AddMessageItem() *MessageItemCondition1 {
	newValue := new(MessageItemCondition1)
	a.MessageItem = append(a.MessageItem, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters6) SetProtectCardData(value string) {
	a.ProtectCardData = (*TrueFalseIndicator)(&value)
}
