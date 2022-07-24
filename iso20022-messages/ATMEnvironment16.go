package iso20022

// Environment of exceptions.
type ATMEnvironment16 struct {

	// Acquirer of transactions, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine9 `xml:"ATM"`

	// Customer involved in the transaction.
	Customer *ATMCustomer6 `xml:"Cstmr,omitempty"`

	// Card performing the transaction.
	Card *PaymentCard23 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment16) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment16) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment16) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment16) AddATM() *AutomatedTellerMachine9 {
	a.ATM = new(AutomatedTellerMachine9)
	return a.ATM
}

func (a *ATMEnvironment16) AddCustomer() *ATMCustomer6 {
	a.Customer = new(ATMCustomer6)
	return a.Customer
}

func (a *ATMEnvironment16) AddCard() *PaymentCard23 {
	a.Card = new(PaymentCard23)
	return a.Card
}
