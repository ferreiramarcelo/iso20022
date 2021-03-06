package iso20022

// Specifies the type of modification to be applied on a data set.
type ModificationScope17 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification *GenericIdentification44 `xml:"OthrId"`

}


func (m *ModificationScope17) SetModificationScopeIndication(value string) {
	m.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (m *ModificationScope17) AddOtherIdentification() *GenericIdentification44 {
	m.OtherIdentification = new(GenericIdentification44)
	return m.OtherIdentification
}

