package iso20022

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus3Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the AcknowledgedAcceptedStatus.
	Reason []*AcknowledgementReason3 `xml:"Rsn,omitempty"`
}

func (a *AcknowledgedAcceptedStatus3Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus3Choice) AddReason() *AcknowledgementReason3 {
	newValue := new(AcknowledgementReason3)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
