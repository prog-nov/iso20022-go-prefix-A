package iso20022

// Context in which the transaction is performed.
type ATMContext1 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService1 `xml:"Svc,omitempty"`
}

func (a *ATMContext1) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext1) AddService() *ATMService1 {
	a.Service = new(ATMService1)
	return a.Service
}
