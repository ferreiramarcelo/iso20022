package iso20022

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation12 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission17 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax21 `xml:"TaxDtls,omitempty"`

	// Specifies foreign exchange details applied to the payment of charges, taxes and commissions as a result of the transfer.
	ForeignExchangeDetails []*ForeignExchangeTerms7 `xml:"FXDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount9 `xml:"SttlmPtiesDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

}


func (r *ReceiveInformation12) SetRequestedSettlementDate(value string) {
	r.RequestedSettlementDate = (*ISODate)(&value)
}

func (r *ReceiveInformation12) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	r.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return r.EffectiveSettlementDate
}

func (r *ReceiveInformation12) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation12) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation12) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation12) AddChargeDetails() *Charge20 {
	newValue := new (Charge20)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation12) AddCommissionDetails() *Commission17 {
	newValue := new (Commission17)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation12) AddTaxDetails() *Tax21 {
	newValue := new (Tax21)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation12) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new (ForeignExchangeTerms7)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation12) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount9 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount9)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation12) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation12) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

func (r *ReceiveInformation12) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

