package iso20022

// Provides further details of the account notification.
type AccountNotification12 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the notification.
	//
	// Usage: The pagination of the notification is only allowed when agreed between the parties.
	NotificationPagination *Pagination `xml:"NtfctnPgntn,omitempty"`

	// Sequential number of the notification, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each notification sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the notification, as assigned by the account servicer. It is increased incrementally for each notification sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account notification is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the notification has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest3 `xml:"Intrst,omitempty"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions5 `xml:"TxsSummry,omitempty"`

	// Specifies an entry in the debit credit notification.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	//
	// Usage Rule:  In case of a Payments R-transaction the creditor / debtor referenced of the original payment initiation messages is also used for reporting of the R-transaction. The original debtor/creditor in the reporting of R-Transactions is not inverted.
	// Following elements all defined in the TransactionDetails in RelatedParties or RelatedAgents are impacted by this usage rule:
	// Creditor, UltimateCreditor, CreditorAccount, CreditorAgent, Debtor, UltimateDebtor, DebtorAccount and DebtorAgent.
	//
	Entry []*ReportEntry8 `xml:"Ntry,omitempty"`

	// Further details of the account notification.
	AdditionalNotificationInformation *Max500Text `xml:"AddtlNtfctnInf,omitempty"`
}

func (a *AccountNotification12) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification12) AddNotificationPagination() *Pagination {
	a.NotificationPagination = new(Pagination)
	return a.NotificationPagination
}

func (a *AccountNotification12) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification12) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification12) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountNotification12) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountNotification12) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountNotification12) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountNotification12) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountNotification12) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountNotification12) AddInterest() *AccountInterest3 {
	newValue := new(AccountInterest3)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountNotification12) AddTransactionsSummary() *TotalTransactions5 {
	a.TransactionsSummary = new(TotalTransactions5)
	return a.TransactionsSummary
}

func (a *AccountNotification12) AddEntry() *ReportEntry8 {
	newValue := new(ReportEntry8)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountNotification12) SetAdditionalNotificationInformation(value string) {
	a.AdditionalNotificationInformation = (*Max500Text)(&value)
}
