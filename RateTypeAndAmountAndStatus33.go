package iso20022

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus33 struct {

	// Value expressed as a rate type.
	RateType *RateType47Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus4Choice `xml:"RateSts,omitempty"`

}


func (r *RateTypeAndAmountAndStatus33) AddRateType() *RateType47Choice {
	r.RateType = new(RateType47Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus33) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus33) AddRateStatus() *RateStatus4Choice {
	r.RateStatus = new(RateStatus4Choice)
	return r.RateStatus
}

