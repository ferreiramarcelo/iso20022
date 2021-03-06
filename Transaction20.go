package iso20022

// Identifies the details of the transaction.
type Transaction20 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails39 `xml:"TxDtls,omitempty"`

}


func (t *Transaction20) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction20) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) AddTransactionDetails() *TransactionDetails39 {
	t.TransactionDetails = new(TransactionDetails39)
	return t.TransactionDetails
}

