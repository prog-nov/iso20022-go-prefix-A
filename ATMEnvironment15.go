package iso20022

// Environment of the key download.
type ATMEnvironment15 struct {

	// Acquirer of the ATM transactions, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine6 `xml:"ATM"`
}

func (a *ATMEnvironment15) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment15) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment15) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment15) AddATM() *AutomatedTellerMachine6 {
	a.ATM = new(AutomatedTellerMachine6)
	return a.ATM
}
