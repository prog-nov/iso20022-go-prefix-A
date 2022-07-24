package iso20022

// Choice of format for the allegment reason.
type AllegementReason1Choice struct {

	// Specifies the reason why the instruction has been alleged.
	Code *AllegementReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has been alleged.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (a *AllegementReason1Choice) SetCode(value string) {
	a.Code = (*AllegementReason1Code)(&value)
}

func (a *AllegementReason1Choice) AddProprietary() *GenericIdentification38 {
	a.Proprietary = new(GenericIdentification38)
	return a.Proprietary
}
