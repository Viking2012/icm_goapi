package orm

import "time"

type BillingDocument struct {
	// Document Details
	Database             string    `json:"Database" sql:"Database"`
	DocumentNumber       string    `json:"Billing Document" sql:"Billing Document"`
	BillingTypeCode      string    `json:"Billing Type Code" sql:"Billing Type Code"`
	BillingType          string    `json:"Billing Type" sql:"Billing Type"`
	BillingCategoryCode  string    `json:"Billing Category Code" sql:"Billing Category Code"`
	BillingCategory      string    `json:"Billing Category" sql:"Billing Category"`
	DocumentCategoryCode string    `json:"Document Category Code" sql:"Document Category Code"`
	DocumentCategory     string    `json:"Document Category" sql:"Document Category"`
	CreationDateDocument time.Time `json:"Created On" sql:"Created On"`
	CreatedBy            string    `json:"Created By" sql:"Created By"`
	Currency             string    `json:"Currency" sql:"Currency"`
	LocalCurrency        string    `json:"Local Currency" sql:"Local Currency"`
	NetValueDocument     float64   `json:"Net Value" sql:"Net Value"`

	// Item Details
	Item                  string    `json:"Item" sql:"Item"`
	MaterialGroupCode     string    `json:"Material Group Code" sql:"Material Group Code"`
	MaterialGroup         string    `json:"Material Group" sql:"Material Group"`
	MaterialCode          string    `json:"Material Code" sql:"Material Code"`
	Material              string    `json:"Material" sql:"Material"`
	Description           string    `json:"Description" sql:"Description"`
	CreationDateItem      time.Time `json:"VBRP_ERDAT" sql:"VBRP_ERDAT"`
	BaseUOMBilledQuantity float64   `json:"Billed Quantity in Base UOM" sql:"Billed Quantity in Base UOM"`
	TaxAmount             float64   `json:"Tax Amount" sql:"Tax Amount"`
	TaxAmountInLC         float64   `json:"Tax Amount in LC" sql:"Tax Amount in LC"`
	TaxAmountInUSD        float64   `json:"Tax Amount in USD" sql:"Tax Amount in USD"`
	Cost                  float64   `json:"Cost" sql:"Cost"`
	CostInLC              float64   `json:"Cost in LC" sql:"Cost in LC"`
	CostInUSD             float64   `json:"Cost in USD" sql:"Cost in USD"`
	NetValueItem          float64   `json:"VBRP_NETWR_Orig" sql:"VBRP_NETWR_Orig"`
	NetValueItemLc        float64   `json:"Net Value in LC" sql:"Net Value in LC"`
	NetValueItemUSD       float64   `json:"Net Value in USD" sql:"Net Value in USD"`

	// References
	DocumentNumberReference string `json:"Reference Document" sql:"Reference Document"`
	ItemReference           string `json:"Reference Item" sql:"Reference Item"`
	DocumentTypeReference   string `json:"Reference Document Type" sql:"Reference Document Type"`

	DocumentNumberOriginating string `json:"VBRP_VBELV" sql:"VBRP_VBELV"`
	ItemOriginating           string `json:"VBRP_POSNV" sql:"VBRP_POSNV"`

	DocumentNumberSales   string `json:"Sales Document" sql:"Sales Document"`
	DocumentCategorySales string `json:"VBRP_AUTYP" sql:"VBRP_AUTYP"`
	ItemSales             string `json:"Sales Item" sql:"Sales Item"`
	ItemCategorySales     string `json:"VBRP_PSTYV" sql:"VBRP_PSTYV"`

	IsReturn                   string `json:"Is Return" sql:"Is Return"`
	IsCancellation             bool   `json:"IsCancel" sql:"IsCancel"`
	IsCancelled                string `json:"Cancelled" sql:"Cancelled"`
	DocumentNumberCancellation string `json:"Cancellation Document" sql:"Cancellation Document"`
	DocumentNumberReverse      string `json:"Reverse Document Number" sql:"Reverse Document Number"`

	IsPostedToFI                   bool      `json:"Posted to FI" sql:"Posted to FI"`
	DocumentStatusAccounting       string    `json:"VBRK_RFBSK" sql:"VBRK_RFBSK"`
	DocumentNumberAccounting       string    `json:"Accounting Doc Number" sql:"Accounting Doc Number"`
	FiscalYearAccountingDocument   string    `json:"Accounting Doc Fiscal Year" sql:"Accounting Doc Fiscal Year"`
	TypeCodeAccountingDocument     string    `json:"Accounting Document Type Code" sql:"Accounting Document Type Code"`
	DocumentTypeAccountingDocument string    `json:"Accounting Document Type" sql:"Accounting Document Type"`
	PostingDateAccountingDocument  time.Time `json:"Posting Date" sql:"Posting Date"`

	ProjectDefinition string `json:"VBRP_PS_PSP_PNR" sql:"VBRP_PS_PSP_PNR"`
	ProjectNumber     string `json:"Project Number" sql:"Project Number"`
	ProjectName       string `json:"PROJ_Concat" sql:"PROJ_Concat"`

	// Document Dates
	BillingDate          time.Time `json:"Billing Date" sql:"Billing Date"`
	PricingDate          time.Time `json:"Pricing Date" sql:"Pricing Date"`
	ServicesRenderedDate time.Time `json:"Services Rendered On" sql:"Services Rendered On"`

	// BA Assignments
	Division                string `json:"VBRK_SPART" sql:"VBRK_SPART"`
	BusinessAreaCode        string `json:"Business Area Code" sql:"Business Area Code"`
	BusinessArea            string `json:"Business Area" sql:"Business Area"`
	SalesOrganizationCode   string `json:"Sales Organization Code" sql:"Sales Organization Code"`
	SalesOrganization       string `json:"Sales Organization" sql:"Sales Organization"`
	DistributionChannelCode string `json:"Distribution Channel Code" sql:"Distribution Channel Code"`
	DistributionChannel     string `json:"Distribution Channel" sql:"Distribution Channel"`
	DocumentCondition       string `json:"Document Condition" sql:"Document Condition"`
	ProfitCenterCode        string `json:"Profit Center Code" sql:"Profit Center Code"`
	ProfitCenter            string `json:"Profit Center" sql:"Profit Center"`
	ProfitCenterOrg1Code    string `json:"VBRP_PRCTR_ORG1" sql:"VBRP_PRCTR_ORG1"`
	ProfitCenterOrg1        string `json:"VBRP_PRCTR_ORG1_Concat" sql:"VBRP_PRCTR_ORG1_Concat"`
	ProfitCenterOrg2Code    string `json:"VBRP_PRCTR_ORG2" sql:"VBRP_PRCTR_ORG2"`
	ProfitCenterOrg2        string `json:"VBRP_PRCTR_ORG2_Concat" sql:"VBRP_PRCTR_ORG2_Concat"`
	ProfitCenterOrg3Code    string `json:"VBRP_PRCTR_ORG3" sql:"VBRP_PRCTR_ORG3"`
	ProfitCenterOrg3        string `json:"VBRP_PRCTR_ORG3_Concat" sql:"VBRP_PRCTR_ORG3_Concat"`
	ProfitCenterCountry     string `json:"Profit Center Country" sql:"Profit Center Country"`
	ProfitCenterLegalCode   string `json:"Profit Center Legal Code" sql:"Profit Center Legal Code"`
	DivisionCode            string `json:"Division Code" sql:"Division Code"`
	CompanyCodeId           string `json:"Company Code ID" sql:"Company Code ID"`
	CompanyCode             string `json:"Company Code" sql:"Company Code"`
	PlantCode               string `json:"Plant Code" sql:"Plant Code"`
	CostCenterCode          string `json:"Cost Center Code" sql:"Cost Center Code"`
	CostCenter              string `json:"Cost Center" sql:"Cost Center"`
	SalesGroup              string `json:"Sales Group" sql:"Sales Group"`
	SalesOfficeCode         string `json:"Sales Office Code" sql:"Sales Office Code"`
	ControllingAreaCode     string `json:"Controlling Area Code" sql:"Controlling Area Code"`
	ShippingPointCode       string `json:"Shipping Point Code" sql:"Shipping Point Code"`

	// Sold-To details
	SoldToID               string `json:"ICM_SOLD_TO_ID" sql:"ICM_SOLD_TO_ID"`
	SoldToNumber           string `json:"Sold To ID" sql:"Sold To ID"`
	SoldTo                 string `json:"Sold To" sql:"Sold To"`
	CountryKeySoldTo       string `json:"VBRK_KUNAG_KNA1_LAND1" sql:"VBRK_KUNAG_KNA1_LAND1"`
	AccountGroupCodeSoldTo string `json:"Sold To Account Group Code" sql:"Sold To Account Group Code"`
	AccountGroupSoldTo     string `json:"VBRK_KUNAG_KTOKD_Concat" sql:"VBRK_KUNAG_KTOKD_Concat"`

	// Payer details
	PayerNumber           string `json:"Payer ID" sql:"Payer ID"`
	Payer                 string `json:"Payer" sql:"Payer"`
	CountryKeyPayer       string `json:"VBRK_KUNRG_KNA1_LAND1" sql:"VBRK_KUNRG_KNA1_LAND1"`
	AccountGroupPayer     string `json:"VBRK_KUNRG_KTOKD_Concat" sql:"VBRK_KUNRG_KTOKD_Concat"`
	AccountGroupCodePayer string `json:"Payer Account Group Code" sql:"Payer Account Group Code"`

	// Document Conditions
	Incoterms1        string `json:"Incoterms 1" sql:"Incoterms 1"`
	Incoterms2        string `json:"Incoterms 2" sql:"Incoterms 2"`
	PaymentTermsCode  string `json:"Payment Terms Code" sql:"Payment Terms Code"`
	BillingPlanNumber string `json:"Billing Plan Number" sql:"Billing Plan Number"`
	BillingPlanItem   string `json:"Billing Plan Item" sql:"Billing Plan Item"`

	// Flags
	Flags BillingDocumentFlags `json:"flags" sql:"FLAGS"`

	// Unassigned fields
	ExchangeRatePosting  float64 `json:"VBRK_KURRF" sql:"VBRK_KURRF"`
	TradingPartner       string  `json:"Trading Partner" sql:"Trading Partner"`
	BilledQuantity       float64 `json:"Billed Quantity" sql:"Billed Quantity"`
	UOMSales             string  `json:"Sales UOM" sql:"Sales UOM"`
	UOMSalesNumerator    float64 `json:"VBRP_UMVKZ" sql:"VBRP_UMVKZ"`
	UOMSalesDenominator  float64 `json:"VBRP_UMVKN" sql:"VBRP_UMVKN"`
	UOMBased             string  `json:"Base UOM" sql:"Base UOM"`
	ScaleQuantity        float64 `json:"Scale Quantity" sql:"Scale Quantity"`
	ExchangeRatePricing  float64 `json:"VBRP_KURSK" sql:"VBRP_KURSK"`
	ProductHierarchyCode string  `json:"Product Hierarchy Code" sql:"Product Hierarchy Code"`
	StatisticalValue     string  `json:"Statistical Value" sql:"Statistical Value"`
	MaterialPricingGroup string  `json:"VBRP_KONDM" sql:"VBRP_KONDM"`
	DocumentNumberOrder  string  `json:"VBRP_AUFNR" sql:"VBRP_AUFNR"`
	DestinationCountry   string  `json:"Destination Country" sql:"Destination Country"`
	SalesDeal            string  `json:"VBRP_KNUMA_AG" sql:"VBRP_KNUMA_AG"`
	OrderReasonCode      string  `json:"Order Reason Code" sql:"Order Reason Code"`
	TaxCode              string  `json:"VBRP_MWSKZ" sql:"VBRP_MWSKZ"`
	ValueSign            int32   `json:"valueSign" sql:"valueSign"`
	ExchangeRateToLc     float64 `json:"Exchange Rate to LC" sql:"Exchange Rate to LC"`
	ExchangeRateToUsd    float64 `json:"Exchange Rate to USD" sql:"Exchange Rate to USD"`
}

func (bd BillingDocument) GetFlags() ICMEntity {
	return bd.Flags
}

type BillingDocumentFlags struct {
	FlagHighReturnRatio                       bool `json:"FLAG: High Return Ratio" sql:"FLAG: High Return Ratio"`
	FlagHighRelativeReturnRatioWithinDivision bool `json:"FLAG: High Relative Return Ratio within Division" sql:"FLAG: High Relative Return Ratio within Division"`
	FlagPatternOfHigherEoqReturns             bool `json:"FLAG: Pattern of Higher EoQ Returns" sql:"FLAG: Pattern of Higher EoQ Returns"`
	FlagPossibleChannelStuffing               bool `json:"FLAG: Possible Channel Stuffing" sql:"FLAG: Possible Channel Stuffing"`
}

func (bdf BillingDocumentFlags) GetFlags() ICMEntity {
	return bdf
}
