package orm

import "time"

type SalesOrder struct {
	Database                     string    `json:"Database" sql:"Database"`
	DocumentNumberSales          string    `json:"Sales Document" sql:"Sales Document"`
	Item                         string    `json:"Item" sql:"Item"`
	CreationDate                 time.Time `json:"Created On" sql:"Created On"`
	CreatedBy                    string    `json:"Created By" sql:"Created By"`
	ValidFrom                    time.Time `json:"Valid From" sql:"Valid From"`
	ValidTo                      time.Time `json:"Valid To" sql:"Valid To"`
	DocumentDate                 time.Time `json:"Document Date" sql:"Document Date"`
	DocumentCategoryCode         string    `json:"Document Category Code" sql:"Document Category Code"`
	DocumentCategory             string    `json:"Document Category" sql:"Document Category"`
	SalesDocumentTypeCode        string    `json:"Sales Document Type Code" sql:"Sales Document Type Code"`
	SalesDocumentType            string    `json:"Sales Document Type" sql:"Sales Document Type"`
	OrderReasonCode              string    `json:"Order Reason Code" sql:"Order Reason Code"`
	BlockDelivery                string    `json:"Delivery Block" sql:"Delivery Block"`
	BlockBilling                 string    `json:"Billing Block" sql:"Billing Block"`
	Vbak_Netwr                   float64   `json:"VBAK_NETWR" sql:"VBAK_NETWR"`
	Currency                     string    `json:"Currency" sql:"Currency"`
	LocalCurrency                string    `json:"Local Currency" sql:"Local Currency"`
	SalesOrganizationCode        string    `json:"Sales Organization Code" sql:"Sales Organization Code"`
	SalesOrganization            string    `json:"Sales Organization" sql:"Sales Organization"`
	DistributionChannelCode      string    `json:"Distribution Channel Code" sql:"Distribution Channel Code"`
	DistributionChannel          string    `json:"Distribution Channel" sql:"Distribution Channel"`
	DivisionCode                 string    `json:"Division Code" sql:"Division Code"`
	SalesGroup                   string    `json:"Sales Group" sql:"Sales Group"`
	SalesOffice                  string    `json:"Sales Office" sql:"Sales Office"`
	BusinessAreaCode             string    `json:"Business Area Code" sql:"Business Area Code"`
	DocumentCondition            string    `json:"Document Condition" sql:"Document Condition"`
	RequestedDeliveryDate        time.Time `json:"Requested Delivery Date" sql:"Requested Delivery Date"`
	PricingProcedure             string    `json:"Pricing Procedure" sql:"Pricing Procedure"`
	CustomerPoNumber             string    `json:"Customer PO Number" sql:"Customer PO Number"`
	CustomerPoType               string    `json:"Customer PO Type" sql:"Customer PO Type"`
	CustomerPoDate               time.Time `json:"Customer PO Date" sql:"Customer PO Date"`
	SoldToCode                   string    `json:"Sold To Code" sql:"Sold To Code"`
	Vbak_Kunnr_Kna1_Land1        string    `json:"VBAK_KUNNR_KNA1_LAND1" sql:"VBAK_KUNNR_KNA1_LAND1"`
	Vbak_Kunnr_Kna1_Name1        string    `json:"VBAK_KUNNR_KNA1_NAME1" sql:"VBAK_KUNNR_KNA1_NAME1"`
	Vbak_Kunnr_Kna1_Ktokd        string    `json:"VBAK_KUNNR_KNA1_KTOKD" sql:"VBAK_KUNNR_KNA1_KTOKD"`
	Vbak_Kunnr_Kna1_Txt30        string    `json:"VBAK_KUNNR_KNA1_TXT30" sql:"VBAK_KUNNR_KNA1_TXT30"`
	CostCenter                   string    `json:"Cost Center" sql:"Cost Center"`
	ControllingAreaCode          string    `json:"Controlling Area Code" sql:"Controlling Area Code"`
	Project                      string    `json:"Project" sql:"Project"`
	CreditControlArea            string    `json:"Credit Control Area" sql:"Credit Control Area"`
	CreditAccount                string    `json:"Credit Account" sql:"Credit Account"`
	ObjectNumber                 string    `json:"Object Number" sql:"Object Number"`
	CompanyCodeToBeBilled        string    `json:"Company Code To Be Billed" sql:"Company Code To Be Billed"`
	Vbak_Bukrs_Vf_Concat         string    `json:"VBAK_BUKRS_VF_Concat" sql:"VBAK_BUKRS_VF_Concat"`
	Confirmed                    string    `json:"Confirmed" sql:"Confirmed"`
	DeliveryStatus               string    `json:"Delivery Status" sql:"Delivery Status"`
	BillingStatus                string    `json:"Billing Status" sql:"Billing Status"`
	OverallStatus                string    `json:"Overall Status" sql:"Overall Status"`
	CreditCheckCode              string    `json:"Credit Check Code" sql:"Credit Check Code"`
	MaterialCode                 string    `json:"Material Code" sql:"Material Code"`
	Material                     string    `json:"Material" sql:"Material"`
	MaterialGroupCode            string    `json:"Material Group Code" sql:"Material Group Code"`
	MaterialGroup                string    `json:"Material Group" sql:"Material Group"`
	ItemText                     string    `json:"Item Text" sql:"Item Text"`
	ItemCategoryCode             string    `json:"Item Category Code" sql:"Item Category Code"`
	ItemTypeCode                 string    `json:"Item Type Code" sql:"Item Type Code"`
	RelevantForDelivery          string    `json:"Relevant for Delivery" sql:"Relevant for Delivery"`
	RelevantForBilling           string    `json:"Relevant for Billing" sql:"Relevant for Billing"`
	Vbap_Uepos                   string    `json:"VBAP_UEPOS" sql:"VBAP_UEPOS"`
	RejectionReasonCode          string    `json:"Rejection Reason Code" sql:"Rejection Reason Code"`
	ProductHierarchyCode         string    `json:"Product Hierarchy Code" sql:"Product Hierarchy Code"`
	TargetQuantity               float64   `json:"Target Quantity" sql:"Target Quantity"`
	TargetUom                    string    `json:"Target UOM" sql:"Target UOM"`
	UOMBased                     string    `json:"Base UOM" sql:"Base UOM"`
	CustomerMaterialNumber       string    `json:"Customer Material Number" sql:"Customer Material Number"`
	UnlimitedOverdeliveryAllowed string    `json:"Unlimited Overdelivery Allowed" sql:"Unlimited Overdelivery Allowed"`
	OverdeliveryTolerance        float64   `json:"Overdelivery Tolerance" sql:"Overdelivery Tolerance"`
	UnderdeliveryTolerance       float64   `json:"Underdelivery Tolerance" sql:"Underdelivery Tolerance"`
	ItemBillingBlock             string    `json:"Item Billing Block" sql:"Item Billing Block"`
	ItemDivision                 string    `json:"Item Division" sql:"Item Division"`
	ItemBusinessArea             string    `json:"Item Business Area" sql:"Item Business Area"`
	Vbap_Gsber_Concat            string    `json:"VBAP_GSBER_Concat" sql:"VBAP_GSBER_Concat"`
	NetValueDocument             float64   `json:"Net Value" sql:"Net Value"`
	NetValueItemLc               float64   `json:"Net Value in LC" sql:"Net Value in LC"`
	NetValueItemUSD              float64   `json:"Net Value in USD" sql:"Net Value in USD"`
	Quantity                     float64   `json:"Quantity" sql:"Quantity"`
	QuantityInBaseUom            float64   `json:"Quantity in Base UOM" sql:"Quantity in Base UOM"`
	UOMSales                     string    `json:"Sales UOM" sql:"Sales UOM"`
	NumeratorToBaseUom           float64   `json:"Numerator to Base UOM" sql:"Numerator to Base UOM"`
	DenominatorToBaseUom         float64   `json:"Denominator to Base UOM" sql:"Denominator to Base UOM"`
	OriginatingDocument          string    `json:"Originating Document" sql:"Originating Document"`
	OriginatingItem              string    `json:"Originating Item" sql:"Originating Item"`
	DocumentNumberReference      string    `json:"Reference Document" sql:"Reference Document"`
	ItemReference                string    `json:"Reference Item" sql:"Reference Item"`
	PlantCode                    string    `json:"Plant Code" sql:"Plant Code"`
	ItemCreatedOn                time.Time `json:"Item Created On" sql:"Item Created On"`
	ItemCreatedBy                string    `json:"Item Created By" sql:"Item Created By"`
	PricingUnit                  float64   `json:"Pricing Unit" sql:"Pricing Unit"`
	PricingUom                   string    `json:"Pricing UOM" sql:"Pricing UOM"`
	Returns                      string    `json:"Returns" sql:"Returns"`
	MaterialPricingGroupCode     string    `json:"Material Pricing Group Code" sql:"Material Pricing Group Code"`
	CommissionGroupCode          string    `json:"Commission Group Code" sql:"Commission Group Code"`
	Cost                         float64   `json:"Cost" sql:"Cost"`
	CostInLC                     float64   `json:"Cost in LC" sql:"Cost in LC"`
	CostInUSD                    float64   `json:"Cost in USD" sql:"Cost in USD"`
	ProfitCenterCode             string    `json:"Profit Center Code" sql:"Profit Center Code"`
	Vbap_Prctr_Org1              string    `json:"VBAP_PRCTR_ORG1" sql:"VBAP_PRCTR_ORG1"`
	Vbap_Prctr_Org1_Concat       string    `json:"VBAP_PRCTR_ORG1_Concat" sql:"VBAP_PRCTR_ORG1_Concat"`
	Vbap_Prctr_Org2              string    `json:"VBAP_PRCTR_ORG2" sql:"VBAP_PRCTR_ORG2"`
	Vbap_Prctr_Org2_Concat       string    `json:"VBAP_PRCTR_ORG2_Concat" sql:"VBAP_PRCTR_ORG2_Concat"`
	Vbap_Prctr_Org3              string    `json:"VBAP_PRCTR_ORG3" sql:"VBAP_PRCTR_ORG3"`
	Vbap_Prctr_Org3_Concat       string    `json:"VBAP_PRCTR_ORG3_Concat" sql:"VBAP_PRCTR_ORG3_Concat"`
	ProfitCenterCountry          string    `json:"Profit Center Country" sql:"Profit Center Country"`
	ProfitCenterAbacusCode       string    `json:"Profit Center Abacus Code" sql:"Profit Center Abacus Code"`
	Vbap_Prctr_Cepc_Land1        string    `json:"VBAP_PRCTR_CEPC_LAND1" sql:"VBAP_PRCTR_CEPC_LAND1"`
	ProfitCenter                 string    `json:"Profit Center" sql:"Profit Center"`
	WbsElementId                 string    `json:"WBS Element ID" sql:"WBS Element ID"`
	ProjectNumber                string    `json:"Project Number" sql:"Project Number"`
	ProjectName                  string    `json:"PROJ_Concat" sql:"PROJ_Concat"`
	PlanningMaterialCode         string    `json:"Planning Material Code" sql:"Planning Material Code"`
	RaKey                        string    `json:"RA Key" sql:"RA Key"`
	StatisticalValue             string    `json:"Statistical Value" sql:"Statistical Value"`
	PricingReferenceMaterialCode string    `json:"Pricing Reference Material Code" sql:"Pricing Reference Material Code"`
	MaterialPricingGroupCode2    string    `json:"Material Pricing Group Code2" sql:"Material Pricing Group Code2"`
	PromotionCode                string    `json:"Promotion Code" sql:"Promotion Code"`
	Vbap_Kostl                   string    `json:"VBAP_KOSTL" sql:"VBAP_KOSTL"`
	Vbap_Kostl_Concat            string    `json:"VBAP_KOSTL_Concat" sql:"VBAP_KOSTL_Concat"`
	ItemConfirmed                string    `json:"Item Confirmed" sql:"Item Confirmed"`
	ItemDeliveryStatus           string    `json:"Item Delivery Status" sql:"Item Delivery Status"`
	ItemOverallDeliveryStatus    string    `json:"Item Overall Delivery Status" sql:"Item Overall Delivery Status"`
	ValueSign                    int32     `json:"valueSign" sql:"valueSign"`
	ExchangeRateToLc             float64   `json:"Exchange Rate to LC" sql:"Exchange Rate to LC"`
	ExchangeRateToUsd            float64   `json:"Exchange Rate to USD" sql:"Exchange Rate to USD"`
	Icm_Customer_Id              string    `json:"ICM_CUSTOMER_ID" sql:"ICM_CUSTOMER_ID"`
}

func (so SalesOrder) GetFlags() ICMEntity {
	return nil
}
