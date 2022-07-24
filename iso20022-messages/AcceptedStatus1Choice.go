package iso20022

// Choice between a reason or no reason for the corporate action instruction processing accepted status.
type AcceptedStatus1Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the accepted status.
	Reason []*AcceptedStatusReason1 `xml:"Rsn"`
}

func (a *AcceptedStatus1Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcceptedStatus1Choice) AddReason() *AcceptedStatusReason1 {
	newValue := new(AcceptedStatusReason1)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
