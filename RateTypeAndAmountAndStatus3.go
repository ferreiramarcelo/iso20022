package iso20022

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus3 struct {

	// Value expressed as a rate type.
	RateType *RateType6Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`

}


func (r *RateTypeAndAmountAndStatus3) AddRateType() *RateType6Choice {
	r.RateType = new(RateType6Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus3) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus3) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}

