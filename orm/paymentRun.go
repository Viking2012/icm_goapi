package orm

import "time"

type PaymentRun struct {
	Database                      string    `json:"Database" sql:"Database"`
	RunDate                       time.Time `json:"Run Date" sql:"Run Date"`
	CheckIdentification           string    `json:"Identification" sql:"Identification"`
	OnlyProposal                  string    `json:"Only Proposal" sql:"Only Proposal"`
	PayingCompanyCode             string    `json:"Paying Company Code" sql:"Paying Company Code"`
	Reguh_Zbukr_Land1             string    `json:"REGUH_ZBUKR_LAND1" sql:"REGUH_ZBUKR_LAND1"`
	Reguh_Zbukr_Concat            string    `json:"REGUH_ZBUKR_Concat" sql:"REGUH_ZBUKR_Concat"`
	Reguh_Zbukr_Land1_Concat      string    `json:"REGUH_ZBUKR_LAND1_Concat" sql:"REGUH_ZBUKR_LAND1_Concat"`
	VendorNumber                  string    `json:"Vendor Number" sql:"Vendor Number"`
	Vendor                        string    `json:"Vendor" sql:"Vendor"`
	AccountGroupCodeVendor        string    `json:"Vendor Account Group Code" sql:"Vendor Account Group Code"`
	AccountGroupVendor            string    `json:"Vendor Account Group" sql:"Vendor Account Group"`
	Code                          string    `json:"Customer Code" sql:"Customer Code"`
	Customer                      string    `json:"Customer" sql:"Customer"`
	AccountGroupCodeCustomer      string    `json:"Customer Account Group Code" sql:"Customer Account Group Code"`
	AccountGroupCustomer          string    `json:"Customer Account Group" sql:"Customer Account Group"`
	PaymentRecipient              string    `json:"Payment Recipient" sql:"Payment Recipient"`
	PaymentDocument               string    `json:"Payment Document" sql:"Payment Document"`
	Currency                      string    `json:"Currency" sql:"Currency"`
	BusinessArea2                 string    `json:"Business Area2" sql:"Business Area2"`
	Name1                         string    `json:"Name 1" sql:"Name 1"`
	Name2                         string    `json:"Name 2" sql:"Name 2"`
	PostalCode                    string    `json:"Postal Code" sql:"Postal Code"`
	City                          string    `json:"City" sql:"City"`
	StreetAddress                 string    `json:"Street" sql:"Street"`
	PoBox                         string    `json:"PO Box" sql:"PO Box"`
	CountryKey                    string    `json:"Country Key" sql:"Country Key"`
	Reguh_Land1_Concat            string    `json:"REGUH_LAND1_Concat" sql:"REGUH_LAND1_Concat"`
	TaxNumber1                    string    `json:"Tax Number 1" sql:"Tax Number 1"`
	PayeeName1                    string    `json:"Payee Name 1" sql:"Payee Name 1"`
	PayeeName2                    string    `json:"Payee Name 2" sql:"Payee Name 2"`
	PayeeName3                    string    `json:"Payee Name 3" sql:"Payee Name 3"`
	PayeeName4                    string    `json:"Payee Name 4" sql:"Payee Name 4"`
	PayeePostalCode               string    `json:"Payee Postal Code" sql:"Payee Postal Code"`
	PayeeCity                     string    `json:"Payee City" sql:"Payee City"`
	PayeeStreet                   string    `json:"Payee Street" sql:"Payee Street"`
	PayeePOBox                    string    `json:"Payee PO Box" sql:"Payee PO Box"`
	PayeeCountryKey               string    `json:"Payee Country Key" sql:"Payee Country Key"`
	Reguh_Zland_Concat            string    `json:"REGUH_ZLAND_Concat" sql:"REGUH_ZLAND_Concat"`
	PayeeBankCountryKey           string    `json:"Payee Bank Country Key" sql:"Payee Bank Country Key"`
	Reguh_Zbnks_Concat            string    `json:"REGUH_ZBNKS_Concat" sql:"REGUH_ZBNKS_Concat"`
	PayeeBankAcctNumber           string    `json:"Payee Bank Acct Number" sql:"Payee Bank Acct Number"`
	PayeeBankNumber               string    `json:"Payee Bank Number" sql:"Payee Bank Number"`
	PayeeBankControlKey           string    `json:"Payee Bank Control Key" sql:"Payee Bank Control Key"`
	SwiftCode                     string    `json:"SWIFT Code" sql:"SWIFT Code"`
	PaymentDocumentPostingDate    time.Time `json:"Payment Document Posting Date" sql:"Payment Document Posting Date"`
	PaymentMethod                 string    `json:"Payment Method" sql:"Payment Method"`
	Reguh_Rzawe_Concat            string    `json:"REGUH_RZAWE_Concat" sql:"REGUH_RZAWE_Concat"`
	AccountId                     string    `json:"Account ID" sql:"Account ID"`
	HouseBank                     string    `json:"House Bank" sql:"House Bank"`
	AcctForBillOfExch             string    `json:"Acct for Bill of Exch" sql:"Acct for Bill of Exch"`
	HouseBankForBillOfExch        string    `json:"House Bank for Bill of Exch " sql:"House Bank for Bill of Exch "`
	PartnerBankType               string    `json:"Partner Bank Type" sql:"Partner Bank Type"`
	BankControlKey                string    `json:"Bank Control Key" sql:"Bank Control Key"`
	OurAccountNumber              string    `json:"Our Account Number" sql:"Our Account Number"`
	OurBankCountry                string    `json:"Our Bank Country" sql:"Our Bank Country"`
	Reguh_Ubnks_Concat            string    `json:"REGUH_UBNKS_Concat" sql:"REGUH_UBNKS_Concat"`
	HouseBankNumber               string    `json:"House Bank Number" sql:"House Bank Number"`
	AmountPaid                    float64   `json:"Amount Paid" sql:"Amount Paid"`
	AmountPaidInUSD               float64   `json:"Amount Paid in USD" sql:"Amount Paid in USD"`
	CashDiscountAmount            float64   `json:"Cash Discount Amount" sql:"Cash Discount Amount"`
	Reguh_Rwskt_Usd               float64   `json:"REGUH_RWSKT_USD" sql:"REGUH_RWSKT_USD"`
	ProposalChanged               string    `json:"Proposal Changed" sql:"Proposal Changed"`
	PaymentBlock                  string    `json:"Payment Block" sql:"Payment Block"`
	Reguh_Zlspr_Concat            string    `json:"REGUH_ZLSPR_Concat" sql:"REGUH_ZLSPR_Concat"`
	LostCashDisc                  float64   `json:"Lost Cash Disc" sql:"Lost Cash Disc"`
	Reguh_Skvfw_Usd               float64   `json:"REGUH_SKVFW_USD" sql:"REGUH_SKVFW_USD"`
	PersonnelNumber               string    `json:"Personnel Number" sql:"Personnel Number"`
	SubledgerGlAccount            string    `json:"Subledger GL Account" sql:"Subledger GL Account"`
	GlAccount                     string    `json:"GL Account" sql:"GL Account"`
	CompanyCodeId                 string    `json:"Company Code ID" sql:"Company Code ID"`
	DocumentNumber                string    `json:"Document Number" sql:"Document Number"`
	FiscalYear                    string    `json:"Fiscal Year" sql:"Fiscal Year"`
	LineItem                      string    `json:"Line item" sql:"Line item"`
	CurrencyRegup                 string    `json:"Currency REGUP" sql:"Currency REGUP"`
	Reference                     string    `json:"Reference" sql:"Reference"`
	DocumentType                  string    `json:"Document Type" sql:"Document Type"`
	Regup_Blart_Concat            string    `json:"REGUP_BLART_Concat" sql:"REGUP_BLART_Concat"`
	PostingDateAccountingDocument time.Time `json:"Posting Date" sql:"Posting Date"`
	DocumentDate                  time.Time `json:"Document Date" sql:"Document Date"`
	AccountType                   string    `json:"Account Type" sql:"Account Type"`
	Regup_Koart_Concat            string    `json:"REGUP_KOART_Concat" sql:"REGUP_KOART_Concat"`
	PostingKey                    string    `json:"Posting Key" sql:"Posting Key"`
	Regup_Bschl_Concat            string    `json:"REGUP_BSCHL_Concat" sql:"REGUP_BSCHL_Concat"`
	GlAccountRegup                string    `json:"GL Account REGUP" sql:"GL Account REGUP"`
	SubledgerGlAccountRegup       string    `json:"Subledger GL Account REGUP" sql:"Subledger GL Account REGUP"`
	SpecialGlInd                  string    `json:"Special GL Ind" sql:"Special GL Ind"`
	SpecialGlType                 string    `json:"Special GL Type" sql:"Special GL Type"`
	TargetSpecialGlInd            string    `json:"Target Special GL Ind" sql:"Target Special GL Ind"`
	DebitCreditIndicator          string    `json:"Dr/Cr Indicator" sql:"Dr/Cr Indicator"`
	Amount                        float64   `json:"Amount" sql:"Amount"`
	Regup_Wrbtr_Usd               float64   `json:"REGUP_WRBTR_USD" sql:"REGUP_WRBTR_USD"`
	BusinessArea                  string    `json:"Business Area" sql:"Business Area"`
	Regup_Gsber_Concat            string    `json:"REGUP_GSBER_Concat" sql:"REGUP_GSBER_Concat"`
	TaxCode                       string    `json:"Tax Code" sql:"Tax Code"`
	ItemText                      string    `json:"Item Text" sql:"Item Text"`
	BaselinePaymentDate           time.Time `json:"Baseline Payment Date" sql:"Baseline Payment Date"`
	PaymentTerms                  string    `json:"Payment Terms" sql:"Payment Terms"`
	Days1                         int32     `json:"Days 1" sql:"Days 1"`
	Days2                         int32     `json:"Days 2" sql:"Days 2"`
	DaysNet                       int32     `json:"Days Net" sql:"Days Net"`
	DiscountPercent1              float64   `json:"Discount Percent 1" sql:"Discount Percent 1"`
	DiscountPercent2              float64   `json:"Discount Percent 2" sql:"Discount Percent 2"`
	DiscountBase                  float64   `json:"Discount base" sql:"Discount base"`
	Regup_Skfbt_Usd               float64   `json:"REGUP_SKFBT_USD" sql:"REGUP_SKFBT_USD"`
	LineCashDiscountAmount        float64   `json:"Line Cash Discount Amount" sql:"Line Cash Discount Amount"`
	Regup_Wskto_Usd               float64   `json:"REGUP_WSKTO_USD" sql:"REGUP_WSKTO_USD"`
	Regup_Zbfix                   string    `json:"REGUP_ZBFIX" sql:"REGUP_ZBFIX"`
	PaymentBlockRegup             string    `json:"Payment Block REGUP" sql:"Payment Block REGUP"`
	CashDiscountDays              int32     `json:"Cash Discount Days" sql:"Cash Discount Days"`
	CashDiscountRate              float64   `json:"Cash Discount Rate" sql:"Cash Discount Rate"`
	MaximumCashDiscountRate       float64   `json:"Maximum Cash Discount Rate" sql:"Maximum Cash Discount Rate"`
	Regup_Msfbt_Usd               float64   `json:"REGUP_MSFBT_USD" sql:"REGUP_MSFBT_USD"`
	PoNumber                      string    `json:"PO Number" sql:"PO Number"`
	PoItem                        string    `json:"PO Item" sql:"PO Item"`
	TradingPartner                string    `json:"Trading Partner" sql:"Trading Partner"`
	VAT                           string    `json:"VAT ID" sql:"VAT ID"`
	Assignment                    string    `json:"Assignment" sql:"Assignment"`
	CostCenter                    string    `json:"Cost Center" sql:"Cost Center"`
	ProfitCenter                  string    `json:"Profit Center" sql:"Profit Center"`
	AddressAndBankSetIndividually string    `json:"Address and Bank Set Individually" sql:"Address and Bank Set Individually"`
	LocalCurrency                 string    `json:"Local Currency" sql:"Local Currency"`
	VendorIsIntercompany          bool      `json:"Vendor Is Intercompany" sql:"Vendor Is Intercompany"`
	CustomerIsIntercompany        bool      `json:"Customer Is Intercompany" sql:"Customer Is Intercompany"`
	ID                            string    `json:"ICM_VENDOR_ID" sql:"ICM_VENDOR_ID"`
	Icm_Customer_Id               string    `json:"ICM_CUSTOMER_ID" sql:"ICM_CUSTOMER_ID"`
}

func (p PaymentRun) GetFlags() ICMEntity {
	return nil
}
