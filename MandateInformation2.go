package iso20022

// Information that serves as a basis to debit an account.
type MandateInformation2 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation1 `xml:"Tp,omitempty"`

	// Set of elements used to provide details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences1 `xml:"Ocrncs,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification32 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification32 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification32 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

	// Set of elements used to provide information to identify the underlying documents associated with the mandate.
	ReferredDocument *ReferredDocumentInformation3 `xml:"RfrdDoc,omitempty"`

}


func (m *MandateInformation2) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateInformation2) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *MandateInformation2) AddType() *MandateTypeInformation1 {
	m.Type = new(MandateTypeInformation1)
	return m.Type
}

func (m *MandateInformation2) AddOccurrences() *MandateOccurrences1 {
	m.Occurrences = new(MandateOccurrences1)
	return m.Occurrences
}

func (m *MandateInformation2) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MandateInformation2) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MandateInformation2) AddCreditorSchemeIdentification() *PartyIdentification32 {
	m.CreditorSchemeIdentification = new(PartyIdentification32)
	return m.CreditorSchemeIdentification
}

func (m *MandateInformation2) AddCreditor() *PartyIdentification32 {
	m.Creditor = new(PartyIdentification32)
	return m.Creditor
}

func (m *MandateInformation2) AddCreditorAccount() *CashAccount16 {
	m.CreditorAccount = new(CashAccount16)
	return m.CreditorAccount
}

func (m *MandateInformation2) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return m.CreditorAgent
}

func (m *MandateInformation2) AddUltimateCreditor() *PartyIdentification32 {
	m.UltimateCreditor = new(PartyIdentification32)
	return m.UltimateCreditor
}

func (m *MandateInformation2) AddDebtor() *PartyIdentification32 {
	m.Debtor = new(PartyIdentification32)
	return m.Debtor
}

func (m *MandateInformation2) AddDebtorAccount() *CashAccount16 {
	m.DebtorAccount = new(CashAccount16)
	return m.DebtorAccount
}

func (m *MandateInformation2) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return m.DebtorAgent
}

func (m *MandateInformation2) AddUltimateDebtor() *PartyIdentification32 {
	m.UltimateDebtor = new(PartyIdentification32)
	return m.UltimateDebtor
}

func (m *MandateInformation2) AddReferredDocument() *ReferredDocumentInformation3 {
	m.ReferredDocument = new(ReferredDocumentInformation3)
	return m.ReferredDocument
}

