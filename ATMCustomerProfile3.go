package iso20022

// Profile of the customer with the allowed services and restrictions.
type ATMCustomerProfile3 struct {

	// Reference of the customer profile.
	ProfileReference *Max35Text `xml:"PrflRef,omitempty"`

	// Identification of the customer for the bank.
	CustomerIdentification *Max35Text `xml:"CstmrId,omitempty"`

	// Description of the customer's profile in plaintext.
	ProfileDescription *Max70Text `xml:"PrflDesc,omitempty"`

	// Services allowed for the customer's profile.
	AllowedServices []*ATMService7 `xml:"AllwdSvcs,omitempty"`

}


func (a *ATMCustomerProfile3) SetProfileReference(value string) {
	a.ProfileReference = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile3) SetCustomerIdentification(value string) {
	a.CustomerIdentification = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile3) SetProfileDescription(value string) {
	a.ProfileDescription = (*Max70Text)(&value)
}

func (a *ATMCustomerProfile3) AddAllowedServices() *ATMService7 {
	newValue := new (ATMService7)
	a.AllowedServices = append(a.AllowedServices, newValue)
	return newValue
}

