package orm

import (
	"database/sql"
)

type FICustomer struct {
	// Primary Keys
	Database           string       `json:"Database" sql:"Database|pk"`
	CompanyCode        string       `json:"Company Code" sql:"Company Code|pk"`
	FiscalYear         string       `json:"Fiscal Year" sql:"Fiscal Year|pk"`
	AccountingDocument string       `json:"Accounting Document" sql:"Accounting Document|pk"`
	AccountingItem     string       `json:"Accounting Item" sql:"Accounting Item|pk"`
	PostedOn           sql.NullTime `json:"Posted On" sql:"Posted On|pk"`
	Ledger             string       `json:"Ledger" sql:"Ledger|pk"`
	LedgerLine         string       `json:"Ledger Line" sql:"Ledger Line|pk"`

	CreatedOn sql.NullTime `json:"Created On" sql:"Created On"`
	UserId    string       `json:"User ID" sql:"User ID"`
	TCodeId   string       `json:"T-Code ID" sql:"T-Code ID"`
	TCode     string       `json:"T-Code" sql:"T-Code"`

	DocumentDate     sql.NullTime `json:"Document Date" sql:"Document Date"`
	ClearedOn        sql.NullTime `json:"Cleared On" sql:"Cleared On"`
	ClearingDocument string       `json:"Clearing Document" sql:"Clearing Document"`

	CustomerID               string `json:"ICM_CUSTOMER_ID" sql:"ICM_CUSTOMER_ID"`
	CustomerCode             string `json:"Customer Code" sql:"Customer Code"`
	Customer                 string `json:"Customer" sql:"Customer"`
	CustomerAccountGroupCode string `json:"Customer Account Group Code" sql:"Customer Account Group Code"`
	CustomerAccountGroup     string `json:"Customer Account Group" sql:"Customer Account Group"`
	OneTimeAccount           string `json:"One-Time Account" sql:"One-Time Account"`
	PaymentBlockCode         string `json:"Payment Block Code" sql:"Payment Block Code"`
	PaymentBlock             string `json:"Payment Block" sql:"Payment Block"`

	CompanyCodeID            string       `json:"Company Code ID" sql:"Company Code ID"`
	LocalCurrency            string       `json:"Local Currency" sql:"Local Currency"`
	COA                      string       `json:"CoA" sql:"CoA"`
	TransactionType          string       `json:"Transaction Type" sql:"Transaction Type"`
	SpecialGlIndicator       string       `json:"Special GL Indicator" sql:"Special GL Indicator"`
	Assignment               string       `json:"Assignment" sql:"Assignment"`
	Currency                 string       `json:"Currency" sql:"Currency"`
	Reference                string       `json:"Reference" sql:"Reference"`
	DocumentTypeCode         string       `json:"Document Type Code" sql:"Document Type Code"`
	DocumentType             string       `json:"Document Type" sql:"Document Type"`
	Period                   string       `json:"Period" sql:"Period"`
	TargetSpecialGlIndicator string       `json:"Target Special GL Indicator" sql:"Target Special GL Indicator"`
	DebitCreditIndicator     string       `json:"Dr/Cr Indicator" sql:"Dr/Cr Indicator"`
	TaxCode                  string       `json:"Tax Code" sql:"Tax Code"`
	TaxAmountInLC            float64      `json:"Tax Amount in LC" sql:"Tax Amount in LC"`
	TaxAmountInDC            float64      `json:"Tax Amount in DC" sql:"Tax Amount in DC"`
	BaselineDate             sql.NullTime `json:"Baseline Date" sql:"Baseline Date"`

	LedgerPostingKey         string `json:"Ledger Posting Key" sql:"Ledger Posting Key"`
	PostingKeyIsSalesRelated string `json:"Posting Key Is Sales Related" sql:"Posting Key Is Sales Related"`
	PostingKey               string `json:"Posting Key" sql:"Posting Key"`

	ItemText   string `json:"Item Text" sql:"Item Text"`
	Bsad_Aufnr string `json:"BSAD_AUFNR" sql:"BSAD_AUFNR"`
	Bsad_Saknr string `json:"BSAD_SAKNR" sql:"BSAD_SAKNR"`

	PaymentTerms string `json:"Payment Terms" sql:"Payment Terms"`
	PaymentDays1 int64  `json:"Payment Days 1" sql:"Payment Days 1"`
	PaymentDays2 int64  `json:"Payment Days 2" sql:"Payment Days 2"`
	PaymentDays3 int64  `json:"Payment Days 3" sql:"Payment Days 3"`

	DiscountPercent1   float64 `json:"Discount Percent 1" sql:"Discount Percent 1"`
	DiscountPercent2   float64 `json:"Discount Percent 2" sql:"Discount Percent 2"`
	DiscountBasisInDC  float64 `json:"Discount Basis in DC" sql:"Discount Basis in DC"`
	DiscountAmountInLC float64 `json:"Discount Amount in LC" sql:"Discount Amount in LC"`
	DiscountAmountInDC float64 `json:"Discount Amount in DC" sql:"Discount Amount in DC"`

	TaxCode1           string  `json:"Tax Code 1" sql:"Tax Code 1"`
	TaxCode1AmountInLC float64 `json:"Tax Code 1 Amount in LC" sql:"Tax Code 1 Amount in LC"`
	TaxCode1AmountInDC float64 `json:"Tax Code 1 Amount in DC" sql:"Tax Code 1 Amount in DC"`
	TaxCode2           string  `json:"Tax Code 2" sql:"Tax Code 2"`
	TaxCode2AmountInLC float64 `json:"Tax Code 2 Amount in LC" sql:"Tax Code 2 Amount in LC"`
	TaxCode2AmountInDC float64 `json:"Tax Code 2 Amount in DC" sql:"Tax Code 2 Amount in DC"`
	TaxCode3           string  `json:"Tax Code 3" sql:"Tax Code 3"`
	TaxCode3AmountInLC float64 `json:"Tax Code 3 Amount in LC" sql:"Tax Code 3 Amount in LC"`
	TaxCode3AmountInDC float64 `json:"Tax Code 3 Amount in DC" sql:"Tax Code 3 Amount in DC"`

	Bsad_Rebzg string `json:"BSAD_REBZG" sql:"BSAD_REBZG"`
	Bsad_Rebzj string `json:"BSAD_REBZJ" sql:"BSAD_REBZJ"`
	Bsad_Rebzz string `json:"BSAD_REBZZ" sql:"BSAD_REBZZ"`
	Bsad_Xanet string `json:"BSAD_XANET" sql:"BSAD_XANET"`

	PaymentMethod string `json:"Payment Method" sql:"Payment Method"`

	AddressAndBankSetIndividually string       `json:"Address and Bank Set Individually" sql:"Address and Bank Set Individually"`
	IsPaymentTransaction          string       `json:"Is Payment Transaction" sql:"Is Payment Transaction"`
	TradingPartnerCode            string       `json:"Trading Partner Code" sql:"Trading Partner Code"`
	WbsElementId                  string       `json:"WBS Element ID" sql:"WBS Element ID"`
	Bsad_Nplnr                    string       `json:"BSAD_NPLNR" sql:"BSAD_NPLNR"`
	Archive                       string       `json:"Archive" sql:"Archive"`
	DocumentStatus                string       `json:"Document Status" sql:"Document Status"`
	Bsad_Auggj                    string       `json:"BSAD_AUGGJ" sql:"BSAD_AUGGJ"`
	RecurringDocument             string       `json:"Recurring Document" sql:"Recurring Document"`
	ReverseDocumentNumber         string       `json:"Reverse Document Number" sql:"Reverse Document Number"`
	Bkpf_Stjah                    string       `json:"BKPF_STJAH" sql:"BKPF_STJAH"`
	HeaderText                    string       `json:"Header Text" sql:"Header Text"`
	AccountingDocumentStatus      string       `json:"Accounting Document Status" sql:"Accounting Document Status"`
	BusinessTransaction           string       `json:"Business Transaction" sql:"Business Transaction"`
	ReferenceTransaction          string       `json:"Reference Transaction" sql:"Reference Transaction"`
	ReferenceKey                  string       `json:"Reference Key" sql:"Reference Key"`
	ParkedBy                      string       `json:"Parked By" sql:"Parked By"`
	LedgerTransaction             string       `json:"Ledger Transaction" sql:"Ledger Transaction"`
	TransactionType2              string       `json:"Transaction Type2" sql:"Transaction Type2"`
	LedgerCurrency                string       `json:"Ledger Currency" sql:"Ledger Currency"`
	BaseUom                       string       `json:"Base UOM" sql:"Base UOM"`
	LedgerReferenceTransaction    string       `json:"Ledger Reference Transaction" sql:"Ledger Reference Transaction"`
	RecordType                    string       `json:"Record Type" sql:"Record Type"`
	Version                       string       `json:"Version" sql:"Version"`
	LogicalSystem                 string       `json:"Logical system" sql:"Logical system"`
	GlAccountCode                 string       `json:"GL Account Code" sql:"GL Account Code"`
	CostElement                   string       `json:"Cost Element" sql:"Cost Element"`
	CostCenterCode                string       `json:"Cost Center Code" sql:"Cost Center Code"`
	ProfitCenterCode              string       `json:"Profit Center Code" sql:"Profit Center Code"`
	Faglflexa_Prctr_Org1          string       `json:"FAGLFLEXA_PRCTR_ORG1" sql:"FAGLFLEXA_PRCTR_ORG1"`
	Faglflexa_Prctr_Org1_Concat   string       `json:"FAGLFLEXA_PRCTR_ORG1_Concat" sql:"FAGLFLEXA_PRCTR_ORG1_Concat"`
	Faglflexa_Prctr_Org2          string       `json:"FAGLFLEXA_PRCTR_ORG2" sql:"FAGLFLEXA_PRCTR_ORG2"`
	Faglflexa_Prctr_Org2_Concat   string       `json:"FAGLFLEXA_PRCTR_ORG2_Concat" sql:"FAGLFLEXA_PRCTR_ORG2_Concat"`
	Faglflexa_Prctr_Org3          string       `json:"FAGLFLEXA_PRCTR_ORG3" sql:"FAGLFLEXA_PRCTR_ORG3"`
	Faglflexa_Prctr_Org3_Concat   string       `json:"FAGLFLEXA_PRCTR_ORG3_Concat" sql:"FAGLFLEXA_PRCTR_ORG3_Concat"`
	ProfitCenterCountryCode       string       `json:"Profit Center Country Code" sql:"Profit Center Country Code"`
	ProfitCenterAbacusCode        string       `json:"Profit Center Abacus Code" sql:"Profit Center Abacus Code"`
	FunctionalAreaCode            string       `json:"Functional Area Code" sql:"Functional Area Code"`
	BusinessAreaCode              string       `json:"Business Area Code" sql:"Business Area Code"`
	ControllingArea               string       `json:"Controlling Area" sql:"Controlling Area"`
	Segment                       string       `json:"Segment" sql:"Segment"`
	SenderCostCenter              string       `json:"Sender Cost Center" sql:"Sender Cost Center"`
	PartnerProfitCenter           string       `json:"Partner Profit Center" sql:"Partner Profit Center"`
	PartnerFunctionalArea         string       `json:"Partner Functional Area" sql:"Partner Functional Area"`
	TradingPartnerBusinessArea    string       `json:"Trading Partner Business Area" sql:"Trading Partner Business Area"`
	TradingPartner                string       `json:"Trading Partner" sql:"Trading Partner"`
	PartnerSegment                string       `json:"Partner Segment" sql:"Partner Segment"`
	Amount                        float64      `json:"Amount" sql:"Amount"`
	AmountInLC                    float64      `json:"Amount in LC" sql:"Amount in LC"`
	AmountInUSD                   float64      `json:"Amount in USD" sql:"Amount in USD"`
	DebitCreditIndicatorLedger    string       `json:"Ledger Dr/Cr Indicator" sql:"Ledger Dr/Cr Indicator"`
	LedgerPostingPeriod           string       `json:"Ledger Posting Period" sql:"Ledger Posting Period"`
	OriginalTransactionCurrency   string       `json:"Original Transaction Currency" sql:"Original Transaction Currency"`
	LedgerLineType                string       `json:"Ledger Line Type" sql:"Ledger Line Type"`
	LedgerLineIsSplit             string       `json:"Ledger Line Is Split" sql:"Ledger Line Is Split"`
	ExchangeRateToUsd             float64      `json:"Exchange Rate to USD" sql:"Exchange Rate to USD"`
	ExchangeRateToUsd2            float64      `json:"Exchange Rate to USD2" sql:"Exchange Rate to USD2"`
	SalesDocumentNumber           string       `json:"SD Document" sql:"SD Document"`
	FinalDueDate                  sql.NullTime `json:"Final Due Date" sql:"Final Due Date"`
	SalesOrderCount               int64        `json:"SO Count" sql:"SO Count"`
	IsFiInvoice                   bool         `json:"Is FI Invoice" sql:"Is FI Invoice"`
}

func (fic *FICustomer) GetFlags() ICMEntity {
	return nil
}

func FICustomerFromRows(rows *sql.Rows) (ICMEntity, error) {
	var fic FICustomer
	err := rows.Scan(
		&fic.Database,
		&fic.CompanyCode,
		&fic.FiscalYear,
		&fic.AccountingDocument,
		&fic.AccountingItem,
		&fic.PostedOn,
		&fic.Ledger,
		&fic.LedgerLine,
		&fic.CreatedOn,
		&fic.UserId,
		&fic.TCodeId,
		&fic.TCode,
		&fic.DocumentDate,
		&fic.ClearedOn,
		&fic.ClearingDocument,
		&fic.CustomerID,
		&fic.CustomerCode,
		&fic.Customer,
		&fic.CustomerAccountGroupCode,
		&fic.CustomerAccountGroup,
		&fic.OneTimeAccount,
		&fic.PaymentBlockCode,
		&fic.PaymentBlock,
		&fic.CompanyCodeID,
		&fic.LocalCurrency,
		&fic.COA,
		&fic.TransactionType,
		&fic.SpecialGlIndicator,
		&fic.Assignment,
		&fic.Currency,
		&fic.Reference,
		&fic.DocumentTypeCode,
		&fic.DocumentType,
		&fic.Period,
		&fic.TargetSpecialGlIndicator,
		&fic.DebitCreditIndicator,
		&fic.TaxCode,
		&fic.TaxAmountInLC,
		&fic.TaxAmountInDC,
		&fic.BaselineDate,
		&fic.LedgerPostingKey,
		&fic.PostingKeyIsSalesRelated,
		&fic.PostingKey,
		&fic.ItemText,
		&fic.Bsad_Aufnr,
		&fic.Bsad_Saknr,
		&fic.PaymentTerms,
		&fic.PaymentDays1,
		&fic.PaymentDays2,
		&fic.PaymentDays3,
		&fic.DiscountPercent1,
		&fic.DiscountPercent2,
		&fic.DiscountBasisInDC,
		&fic.DiscountAmountInLC,
		&fic.DiscountAmountInDC,
		&fic.TaxCode1,
		&fic.TaxCode1AmountInLC,
		&fic.TaxCode1AmountInDC,
		&fic.TaxCode2,
		&fic.TaxCode2AmountInLC,
		&fic.TaxCode2AmountInDC,
		&fic.TaxCode3,
		&fic.TaxCode3AmountInLC,
		&fic.TaxCode3AmountInDC,
		&fic.Bsad_Rebzg,
		&fic.Bsad_Rebzj,
		&fic.Bsad_Rebzz,
		&fic.Bsad_Xanet,
		&fic.PaymentMethod,
		&fic.AddressAndBankSetIndividually,
		&fic.IsPaymentTransaction,
		&fic.TradingPartnerCode,
		&fic.WbsElementId,
		&fic.Bsad_Nplnr,
		&fic.Archive,
		&fic.DocumentStatus,
		&fic.Bsad_Auggj,
		&fic.RecurringDocument,
		&fic.ReverseDocumentNumber,
		&fic.Bkpf_Stjah,
		&fic.HeaderText,
		&fic.AccountingDocumentStatus,
		&fic.BusinessTransaction,
		&fic.ReferenceTransaction,
		&fic.ReferenceKey,
		&fic.ParkedBy,
		&fic.LedgerTransaction,
		&fic.TransactionType2,
		&fic.LedgerCurrency,
		&fic.BaseUom,
		&fic.LedgerReferenceTransaction,
		&fic.RecordType,
		&fic.Version,
		&fic.LogicalSystem,
		&fic.GlAccountCode,
		&fic.CostElement,
		&fic.CostCenterCode,
		&fic.ProfitCenterCode,
		&fic.Faglflexa_Prctr_Org1,
		&fic.Faglflexa_Prctr_Org1_Concat,
		&fic.Faglflexa_Prctr_Org2,
		&fic.Faglflexa_Prctr_Org2_Concat,
		&fic.Faglflexa_Prctr_Org3,
		&fic.Faglflexa_Prctr_Org3_Concat,
		&fic.ProfitCenterCountryCode,
		&fic.ProfitCenterAbacusCode,
		&fic.FunctionalAreaCode,
		&fic.BusinessAreaCode,
		&fic.ControllingArea,
		&fic.Segment,
		&fic.SenderCostCenter,
		&fic.PartnerProfitCenter,
		&fic.PartnerFunctionalArea,
		&fic.TradingPartnerBusinessArea,
		&fic.TradingPartner,
		&fic.PartnerSegment,
		&fic.Amount,
		&fic.AmountInLC,
		&fic.AmountInUSD,
		&fic.DebitCreditIndicatorLedger,
		&fic.LedgerPostingPeriod,
		&fic.OriginalTransactionCurrency,
		&fic.LedgerLineType,
		&fic.LedgerLineIsSplit,
		&fic.ExchangeRateToUsd,
		&fic.ExchangeRateToUsd2,
		&fic.SalesDocumentNumber,
		&fic.FinalDueDate,
		&fic.SalesOrderCount,
		&fic.IsFiInvoice,
	)
	return &fic, err
}
