package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type FIVendor struct {
	Database                          string         `json:"Database" sql:"Database"`
	CompanyCodeId                     string         `json:"Company Code ID" sql:"Company Code ID"`
	CompanyCodeCountry                string         `json:"Company Code Country" sql:"Company Code Country"`
	CompanyCode                       string         `json:"Company Code" sql:"Company Code"`
	LocalCurrency                     string         `json:"Local Currency" sql:"Local Currency"`
	Coa                               string         `json:"CoA" sql:"CoA"`
	VendorNumber                      string         `json:"Vendor Number" sql:"Vendor Number"`
	Vendor                            string         `json:"Vendor" sql:"Vendor"`
	AccountGroupCode                  string         `json:"Vendor Account Group Code" sql:"Vendor Account Group Code"`
	OneTimeVendorIndicator            string         `json:"One-time Vendor Indicator" sql:"One-time Vendor Indicator"`
	AccountGroup                      string         `json:"Vendor Account Group" sql:"Vendor Account Group"`
	SpecialGlTransactionType          string         `json:"Special GL Transaction Type" sql:"Special GL Transaction Type"`
	SpecialGlIndicatorCode            string         `json:"Special GL Indicator Code" sql:"Special GL Indicator Code"`
	ClearingDate                      sql.NullTime   `json:"Clearing Date" sql:"Clearing Date"`
	ClearingDocumentNumber            string         `json:"Clearing Document Number" sql:"Clearing Document Number"`
	AssignmentNumber                  string         `json:"Assignment Number" sql:"Assignment Number"`
	FiscalYear                        string         `json:"Fiscal Year" sql:"Fiscal Year"`
	AccountingDocumentNumber          string         `json:"Accounting Document Number" sql:"Accounting Document Number"`
	AccountingDocumentLineItem        string         `json:"Accounting Document Line Item" sql:"Accounting Document Line Item"`
	PostingDateAccountingDocument     sql.NullTime   `json:"Posting Date" sql:"Posting Date"`
	DocumentDate                      sql.NullTime   `json:"Document Date" sql:"Document Date"`
	EntryDate                         sql.NullTime   `json:"Entry Date" sql:"Entry Date"`
	DocumentCurrency                  string         `json:"Document Currency" sql:"Document Currency"`
	ReferenceDocumentNumber           string         `json:"Reference Document Number" sql:"Reference Document Number"`
	DocumentTypeId                    string         `json:"Document Type ID" sql:"Document Type ID"`
	DocumentType                      string         `json:"Document Type" sql:"Document Type"`
	FiscalPeriod                      string         `json:"Fiscal Period" sql:"Fiscal Period"`
	TargetSpecialGlIndicator          string         `json:"Target Special GL Indicator" sql:"Target Special GL Indicator"`
	IndicatorDebitCredit              string         `json:"Indicator - Debit/Credit" sql:"Indicator - Debit/Credit"`
	TaxOnSalesPurchasesCode           string         `json:"Tax on Sales Purchases Code" sql:"Tax on Sales Purchases Code"`
	TaxLc                             float64        `json:"Tax LC" sql:"Tax LC"`
	Tax                               float64        `json:"Tax" sql:"Tax"`
	LineItemText                      string         `json:"Line Item text" sql:"Line Item text"`
	OrderNumber                       string         `json:"Order Number" sql:"Order Number"`
	Bsak_Saknr                        string         `json:"BSAK_SAKNR" sql:"BSAK_SAKNR"`
	BaselineDateForDueDateCalculation sql.NullTime   `json:"Baseline Date for Due Date Calculation" sql:"Baseline Date for Due Date Calculation"`
	PaymentTermsKey                   string         `json:"Payment Terms Key" sql:"Payment Terms Key"`
	CashDiscountDays1                 int64          `json:"Cash discount days 1" sql:"Cash discount days 1"`
	CashDiscountDays2                 int64          `json:"Cash discount days 2" sql:"Cash discount days 2"`
	NetPaymentTerms                   int64          `json:"Net Payment Terms" sql:"Net Payment Terms"`
	CashDiscountPercentage1           float64        `json:"Cash Discount percentage 1" sql:"Cash Discount percentage 1"`
	CashDiscountPercentage2           float64        `json:"Cash Discount percentage 2" sql:"Cash Discount percentage 2"`
	DiscountBase                      float64        `json:"Discount Base" sql:"Discount Base"`
	DiscountAmountLc                  float64        `json:"Discount Amount LC" sql:"Discount Amount LC"`
	DiscountAmount                    float64        `json:"Discount Amount" sql:"Discount Amount"`
	PaymentMethod                     string         `json:"Payment Method" sql:"Payment Method"`
	PaymentBlockId                    string         `json:"Payment Block ID" sql:"Payment Block ID"`
	PaymentBlock                      string         `json:"Payment Block" sql:"Payment Block"`
	MmInvoice                         string         `json:"MM Invoice" sql:"MM Invoice"`
	MmInvoiceYear                     string         `json:"MM Invoice Year" sql:"MM Invoice Year"`
	MmInvoiceItem                     string         `json:"MM Invoice Item" sql:"MM Invoice Item"`
	DownPaymentIndicator              string         `json:"Down Payment Indicator" sql:"Down Payment Indicator"`
	BankAndAddressSetIndividually     string         `json:"Bank and Address set Individually" sql:"Bank and Address set Individually"`
	PaymentTransaction                string         `json:"Payment Transaction" sql:"Payment Transaction"`
	TaxCodeForDistribution1           string         `json:"Tax Code for Distribution1" sql:"Tax Code for Distribution1"`
	TaxAmount1Lc                      float64        `json:"Tax Amount 1 LC" sql:"Tax Amount 1 LC"`
	TaxAmount1                        float64        `json:"Tax Amount 1" sql:"Tax Amount 1"`
	TaxCodeForDistribution2           string         `json:"Tax Code for Distribution2" sql:"Tax Code for Distribution2"`
	TaxAmount2Lc                      float64        `json:"Tax Amount 2 LC" sql:"Tax Amount 2 LC"`
	TaxAmount2                        float64        `json:"Tax Amount 2" sql:"Tax Amount 2"`
	TaxCodeForDistribution3           string         `json:"Tax Code for Distribution3" sql:"Tax Code for Distribution3"`
	TaxAmount3Lc                      float64        `json:"Tax Amount 3 LC" sql:"Tax Amount 3 LC"`
	TaxAmount3                        float64        `json:"Tax Amount 3" sql:"Tax Amount 3"`
	TradingPartner                    string         `json:"Trading Partner" sql:"Trading Partner"`
	WbsElementId                      string         `json:"WBS Element ID" sql:"WBS Element ID"`
	Network                           string         `json:"Network" sql:"Network"`
	Archive                           string         `json:"Archive" sql:"Archive"`
	DocumentStatus                    string         `json:"Document Status" sql:"Document Status"`
	ClearingYear                      string         `json:"Clearing Year" sql:"Clearing Year"`
	UserId                            string         `json:"User ID" sql:"User ID"`
	TCodeId                           string         `json:"T-Code ID" sql:"T-Code ID"`
	TCode                             string         `json:"T-Code" sql:"T-Code"`
	RecurringDocument                 string         `json:"Recurring Document" sql:"Recurring Document"`
	DocumentNumberReverse             string         `json:"Reverse Document Number" sql:"Reverse Document Number"`
	Bkpf_Stjah                        string         `json:"BKPF_STJAH" sql:"BKPF_STJAH"`
	HeaderText                        string         `json:"Header Text" sql:"Header Text"`
	AccountingDocumentStatus          string         `json:"Accounting Document Status" sql:"Accounting Document Status"`
	BusinessTransaction               string         `json:"Business Transaction" sql:"Business Transaction"`
	ReferenceTransaction              string         `json:"Reference Transaction" sql:"Reference Transaction"`
	ReferenceKey                      string         `json:"Reference Key" sql:"Reference Key"`
	ParkedBy                          string         `json:"Parked By" sql:"Parked By"`
	Ledger                            string         `json:"Ledger" sql:"Ledger"`
	LedgerLine                        string         `json:"Ledger Line" sql:"Ledger Line"`
	LedgerTransaction                 string         `json:"Ledger Transaction" sql:"Ledger Transaction"`
	TransactionType                   string         `json:"Transaction Type" sql:"Transaction Type"`
	LedgerCurrency                    string         `json:"Ledger Currency" sql:"Ledger Currency"`
	UOMBased                          string         `json:"Base UOM" sql:"Base UOM"`
	LedgerReferenceTransaction        string         `json:"Ledger Reference Transaction" sql:"Ledger Reference Transaction"`
	RecordType                        string         `json:"Record Type" sql:"Record Type"`
	Version                           string         `json:"Version" sql:"Version"`
	LogicalSystem                     string         `json:"Logical system" sql:"Logical system"`
	GLAccountCode                     string         `json:"GL Account Code" sql:"GL Account Code"`
	CostElement                       string         `json:"Cost Element" sql:"Cost Element"`
	CostCenterCode                    string         `json:"Cost Center Code" sql:"Cost Center Code"`
	ProfitCenterCode                  string         `json:"Profit Center Code" sql:"Profit Center Code"`
	Faglflexa_Prctr_Org1              string         `json:"FAGLFLEXA_PRCTR_ORG1" sql:"FAGLFLEXA_PRCTR_ORG1"`
	Faglflexa_Prctr_Org1_Concat       string         `json:"FAGLFLEXA_PRCTR_ORG1_Concat" sql:"FAGLFLEXA_PRCTR_ORG1_Concat"`
	Faglflexa_Prctr_Org2              string         `json:"FAGLFLEXA_PRCTR_ORG2" sql:"FAGLFLEXA_PRCTR_ORG2"`
	Faglflexa_Prctr_Org2_Concat       string         `json:"FAGLFLEXA_PRCTR_ORG2_Concat" sql:"FAGLFLEXA_PRCTR_ORG2_Concat"`
	Faglflexa_Prctr_Org3              string         `json:"FAGLFLEXA_PRCTR_ORG3" sql:"FAGLFLEXA_PRCTR_ORG3"`
	Faglflexa_Prctr_Org3_Concat       string         `json:"FAGLFLEXA_PRCTR_ORG3_Concat" sql:"FAGLFLEXA_PRCTR_ORG3_Concat"`
	ProfitCenterCountryCode           string         `json:"Profit Center Country Code" sql:"Profit Center Country Code"`
	ProfitCenterAbacusCode            string         `json:"Profit Center Abacus Code" sql:"Profit Center Abacus Code"`
	FunctionalAreaCode                string         `json:"Functional Area Code" sql:"Functional Area Code"`
	BusinessAreaCode                  string         `json:"Business Area Code" sql:"Business Area Code"`
	ControllingArea                   string         `json:"Controlling Area" sql:"Controlling Area"`
	Segment                           string         `json:"Segment" sql:"Segment"`
	SenderCostCenter                  string         `json:"Sender Cost Center" sql:"Sender Cost Center"`
	PartnerProfitCenter               string         `json:"Partner Profit Center" sql:"Partner Profit Center"`
	PartnerFunctionalArea             string         `json:"Partner Functional Area" sql:"Partner Functional Area"`
	TradingPartnerBusinessArea        string         `json:"Trading Partner Business Area" sql:"Trading Partner Business Area"`
	TradingPartner2                   string         `json:"Trading Partner2" sql:"Trading Partner2"`
	PartnerSegment                    string         `json:"Partner Segment" sql:"Partner Segment"`
	AmountInLc                        float64        `json:"Amount in LC" sql:"Amount in LC"`
	AmountInUsd                       float64        `json:"Amount in USD" sql:"Amount in USD"`
	Amount                            float64        `json:"Amount" sql:"Amount"`
	LedgerDrCrIndicator               string         `json:"Ledger Dr/Cr Indicator" sql:"Ledger Dr/Cr Indicator"`
	LedgerPostingPeriod               string         `json:"Ledger Posting Period" sql:"Ledger Posting Period"`
	OriginalTransactionCurrency       string         `json:"Original Transaction Currency" sql:"Original Transaction Currency"`
	LedgerPostingKey                  string         `json:"Ledger Posting Key" sql:"Ledger Posting Key"`
	PostingKeyIsSalesRelated          string         `json:"Posting Key Is Sales Related" sql:"Posting Key Is Sales Related"`
	PostingKey                        string         `json:"Posting Key" sql:"Posting Key"`
	LedgerLineType                    string         `json:"Ledger Line Type" sql:"Ledger Line Type"`
	LedgerLineIsSplit                 string         `json:"Ledger Line Is Split" sql:"Ledger Line Is Split"`
	ExchangeRateToUsd                 float64        `json:"Exchange Rate to USD" sql:"Exchange Rate to USD"`
	ExchangeRateToUsd2                float64        `json:"Exchange Rate to USD2" sql:"Exchange Rate to USD2"`
	MmDocument                        string         `json:"MM Document" sql:"MM Document"`
	MmYear                            string         `json:"MM Year" sql:"MM Year"`
	FinalDueDate                      sql.NullTime   `json:"Final Due Date" sql:"Final Due Date"`
	PoCount                           int64          `json:"PO Count" sql:"PO Count"`
	ID                                string         `json:"ICM_VENDOR_ID" sql:"ICM_VENDOR_ID"`
	IsFiInvoice                       bool           `json:"Is FI Invoice" sql:"Is FI Invoice"`
	FlagSplitTotalInUsd               float64        `json:"Flag: Split Total in USD" sql:"Flag: Split Total in USD"`
	Right_Bsak_Auggj                  string         `json:"Right_BSAK_AUGGJ" sql:"Right_BSAK_AUGGJ"`
	Right_Bkpf_Stjah                  string         `json:"Right_BKPF_STJAH" sql:"Right_BKPF_STJAH"`
	Right_Bkpf_Bktxt                  string         `json:"Right_BKPF_BKTXT" sql:"Right_BKPF_BKTXT"`
	Right_Bkpf_Ppnam                  string         `json:"Right_BKPF_PPNAM" sql:"Right_BKPF_PPNAM"`
	Right_Faglflexa_Rmvct             string         `json:"Right_FAGLFLEXA_RMVCT" sql:"Right_FAGLFLEXA_RMVCT"`
	Right_Faglflexa_Runit             string         `json:"Right_FAGLFLEXA_RUNIT" sql:"Right_FAGLFLEXA_RUNIT"`
	Right_Faglflexa_Logsys            string         `json:"Right_FAGLFLEXA_LOGSYS" sql:"Right_FAGLFLEXA_LOGSYS"`
	Right_Faglflexa_Cost_Elem         string         `json:"Right_FAGLFLEXA_COST_ELEM" sql:"Right_FAGLFLEXA_COST_ELEM"`
	Right_Faglflexa_Prctr_Org1        string         `json:"Right_FAGLFLEXA_PRCTR_ORG1" sql:"Right_FAGLFLEXA_PRCTR_ORG1"`
	Right_Faglflexa_Prctr_Org1_Concat string         `json:"Right_FAGLFLEXA_PRCTR_ORG1_Concat" sql:"Right_FAGLFLEXA_PRCTR_ORG1_Concat"`
	Right_Faglflexa_Prctr_Org2        string         `json:"Right_FAGLFLEXA_PRCTR_ORG2" sql:"Right_FAGLFLEXA_PRCTR_ORG2"`
	Right_Faglflexa_Prctr_Org2_Concat string         `json:"Right_FAGLFLEXA_PRCTR_ORG2_Concat" sql:"Right_FAGLFLEXA_PRCTR_ORG2_Concat"`
	Right_Faglflexa_Prctr_Org3        string         `json:"Right_FAGLFLEXA_PRCTR_ORG3" sql:"Right_FAGLFLEXA_PRCTR_ORG3"`
	Right_Faglflexa_Prctr_Org3_Concat string         `json:"Right_FAGLFLEXA_PRCTR_ORG3_Concat" sql:"Right_FAGLFLEXA_PRCTR_ORG3_Concat"`
	Right_Faglflexa_Segment           string         `json:"Right_FAGLFLEXA_SEGMENT" sql:"Right_FAGLFLEXA_SEGMENT"`
	Right_Faglflexa_Scntr             string         `json:"Right_FAGLFLEXA_SCNTR" sql:"Right_FAGLFLEXA_SCNTR"`
	Right_Faglflexa_Pprctr            string         `json:"Right_FAGLFLEXA_PPRCTR" sql:"Right_FAGLFLEXA_PPRCTR"`
	Right_Faglflexa_Sfarea            string         `json:"Right_FAGLFLEXA_SFAREA" sql:"Right_FAGLFLEXA_SFAREA"`
	Right_Faglflexa_Sbusa             string         `json:"Right_FAGLFLEXA_SBUSA" sql:"Right_FAGLFLEXA_SBUSA"`
	Right_Faglflexa_Rassc             string         `json:"Right_FAGLFLEXA_RASSC" sql:"Right_FAGLFLEXA_RASSC"`
	Right_Faglflexa_Psegment          string         `json:"Right_FAGLFLEXA_PSEGMENT" sql:"Right_FAGLFLEXA_PSEGMENT"`
	Right_Faglflexa_Rwcur             string         `json:"Right_FAGLFLEXA_RWCUR" sql:"Right_FAGLFLEXA_RWCUR"`
	Right_Icm_Vendor_Id               string         `json:"Right_ICM_VENDOR_ID" sql:"Right_ICM_VENDOR_ID"`
	CreationDate                      sql.NullTime   `json:"Vendor Creation Date" sql:"Vendor Creation Date"`
	DaysSinceVendorCreation           int64          `json:"Days Since Vendor Creation" sql:"Days Since Vendor Creation"`
	Avg_DaysSinceVendorCreation       float64        `json:"Avg_Days Since Vendor Creation" sql:"Avg_Days Since Vendor Creation"`
	Avgdays20                         int64          `json:"20% of AvgDays" sql:"20% of AvgDays"`
	ZScore                            float64        `json:"Z-Score" sql:"Z-Score"`
	DescriptiveBenford                string         `json:"BenfordDescriptive" sql:"BenfordDescriptive"`
	Flags                             *FIVendorFlags `json:"flags" sql:"FLAGS"`
}

func (fiv *FIVendor) IsICMEntity() bool { return true }

func FIVendorFromRows(rows *sql.Rows) (ICMEntity, error) {
	var v FIVendor
	err := rows.Scan(
		&v.Database,
		&v.CompanyCodeId,
		&v.CompanyCodeCountry,
		&v.CompanyCode,
		&v.LocalCurrency,
		&v.Coa,
		&v.VendorNumber,
		&v.Vendor,
		&v.AccountGroupCode,
		&v.OneTimeVendorIndicator,
		&v.AccountGroup,
		&v.SpecialGlTransactionType,
		&v.SpecialGlIndicatorCode,
		&v.ClearingDate,
		&v.ClearingDocumentNumber,
		&v.AssignmentNumber,
		&v.FiscalYear,
		&v.AccountingDocumentNumber,
		&v.AccountingDocumentLineItem,
		&v.PostingDateAccountingDocument,
		&v.DocumentDate,
		&v.EntryDate,
		&v.DocumentCurrency,
		&v.ReferenceDocumentNumber,
		&v.DocumentTypeId,
		&v.DocumentType,
		&v.FiscalPeriod,
		&v.TargetSpecialGlIndicator,
		&v.IndicatorDebitCredit,
		&v.TaxOnSalesPurchasesCode,
		&v.TaxLc,
		&v.Tax,
		&v.LineItemText,
		&v.OrderNumber,
		&v.Bsak_Saknr,
		&v.BaselineDateForDueDateCalculation,
		&v.PaymentTermsKey,
		&v.CashDiscountDays1,
		&v.CashDiscountDays2,
		&v.NetPaymentTerms,
		&v.CashDiscountPercentage1,
		&v.CashDiscountPercentage2,
		&v.DiscountBase,
		&v.DiscountAmountLc,
		&v.DiscountAmount,
		&v.PaymentMethod,
		&v.PaymentBlockId,
		&v.PaymentBlock,
		&v.MmInvoice,
		&v.MmInvoiceYear,
		&v.MmInvoiceItem,
		&v.DownPaymentIndicator,
		&v.BankAndAddressSetIndividually,
		&v.PaymentTransaction,
		&v.TaxCodeForDistribution1,
		&v.TaxAmount1Lc,
		&v.TaxAmount1,
		&v.TaxCodeForDistribution2,
		&v.TaxAmount2Lc,
		&v.TaxAmount2,
		&v.TaxCodeForDistribution3,
		&v.TaxAmount3Lc,
		&v.TaxAmount3,
		&v.TradingPartner,
		&v.WbsElementId,
		&v.Network,
		&v.Archive,
		&v.DocumentStatus,
		&v.ClearingYear,
		&v.UserId,
		&v.TCodeId,
		&v.TCode,
		&v.RecurringDocument,
		&v.DocumentNumberReverse,
		&v.Bkpf_Stjah,
		&v.HeaderText,
		&v.AccountingDocumentStatus,
		&v.BusinessTransaction,
		&v.ReferenceTransaction,
		&v.ReferenceKey,
		&v.ParkedBy,
		&v.Ledger,
		&v.LedgerLine,
		&v.LedgerTransaction,
		&v.TransactionType,
		&v.LedgerCurrency,
		&v.UOMBased,
		&v.LedgerReferenceTransaction,
		&v.RecordType,
		&v.Version,
		&v.LogicalSystem,
		&v.GLAccountCode,
		&v.CostElement,
		&v.CostCenterCode,
		&v.ProfitCenterCode,
		&v.Faglflexa_Prctr_Org1,
		&v.Faglflexa_Prctr_Org1_Concat,
		&v.Faglflexa_Prctr_Org2,
		&v.Faglflexa_Prctr_Org2_Concat,
		&v.Faglflexa_Prctr_Org3,
		&v.Faglflexa_Prctr_Org3_Concat,
		&v.ProfitCenterCountryCode,
		&v.ProfitCenterAbacusCode,
		&v.FunctionalAreaCode,
		&v.BusinessAreaCode,
		&v.ControllingArea,
		&v.Segment,
		&v.SenderCostCenter,
		&v.PartnerProfitCenter,
		&v.PartnerFunctionalArea,
		&v.TradingPartnerBusinessArea,
		&v.TradingPartner2,
		&v.PartnerSegment,
		&v.AmountInLc,
		&v.AmountInUsd,
		&v.Amount,
		&v.LedgerDrCrIndicator,
		&v.LedgerPostingPeriod,
		&v.OriginalTransactionCurrency,
		&v.LedgerPostingKey,
		&v.PostingKeyIsSalesRelated,
		&v.PostingKey,
		&v.LedgerLineType,
		&v.LedgerLineIsSplit,
		&v.ExchangeRateToUsd,
		&v.ExchangeRateToUsd2,
		&v.MmDocument,
		&v.MmYear,
		&v.FinalDueDate,
		&v.PoCount,
		&v.ID,
		&v.IsFiInvoice,
		&v.FlagSplitTotalInUsd,
		&v.Right_Bsak_Auggj,
		&v.Right_Bkpf_Stjah,
		&v.Right_Bkpf_Bktxt,
		&v.Right_Bkpf_Ppnam,
		&v.Right_Faglflexa_Rmvct,
		&v.Right_Faglflexa_Runit,
		&v.Right_Faglflexa_Logsys,
		&v.Right_Faglflexa_Cost_Elem,
		&v.Right_Faglflexa_Prctr_Org1,
		&v.Right_Faglflexa_Prctr_Org1_Concat,
		&v.Right_Faglflexa_Prctr_Org2,
		&v.Right_Faglflexa_Prctr_Org2_Concat,
		&v.Right_Faglflexa_Prctr_Org3,
		&v.Right_Faglflexa_Prctr_Org3_Concat,
		&v.Right_Faglflexa_Segment,
		&v.Right_Faglflexa_Scntr,
		&v.Right_Faglflexa_Pprctr,
		&v.Right_Faglflexa_Sfarea,
		&v.Right_Faglflexa_Sbusa,
		&v.Right_Faglflexa_Rassc,
		&v.Right_Faglflexa_Psegment,
		&v.Right_Faglflexa_Rwcur,
		&v.Right_Icm_Vendor_Id,
		&v.CreationDate,
		&v.DaysSinceVendorCreation,
		&v.Avg_DaysSinceVendorCreation,
		&v.Avgdays20,
		&v.ZScore,
		&v.DescriptiveBenford,
		&v.Flags,
	)
	return &v, err
}

type FIVendorFlags struct {
	FlagIsSplit                bool          `json:"Flag: Is Split" sql:"Flag: Is Split"`
	FlagIsRoundAmount          bool          `json:"Flag: Is Round Amount" sql:"Flag: Is Round Amount"`
	ConsInvVendor              BoolFromFloat `json:"ConsInvVendor" sql:"ConsInvVendor"`
	FlagConsInv                bool          `json:"Flag: ConsInvVendor" sql:"Flag: ConsInvVendor"`
	FlagInvoice_Without_Po     bool          `json:"Flag: Invoice_without_PO" sql:"Flag: Invoice_without_PO"`
	FlagSwiftFirstTransaction  bool          `json:"Flag: SwiftFirstTransaction" sql:"Flag: SwiftFirstTransaction"`
	FlagIncreasingInvoiceValue bool          `json:"Flag: Increasing Invoice Value" sql:"Flag: Increasing Invoice Value"`
	FlagTop5ZScoreVendors      bool          `json:"Flag: Top5 Z-ScoreVendors" sql:"Flag: Top5 Z-ScoreVendors"`
	FlagTop5BenfordDeviation   bool          `json:"Flag: Top5BenfordDeviation" sql:"Flag: Top5BenfordDeviation"`
}

func (fivf *FIVendorFlags) IsICMEntity() bool { return true }
func (fivf *FIVendorFlags) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return json.Unmarshal([]byte(v), fivf)
	case []byte:
		return json.Unmarshal(v, fivf)
	default:
		return errors.New("invalid sql return type for Vendor Flags")
	}
}
