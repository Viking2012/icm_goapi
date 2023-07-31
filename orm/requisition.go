package orm

import "time"

type Requisition struct {
	Database                    string    `json:"Database" sql:"Database"`
	PurchaseRequisition         string    `json:"Purchase Requisition" sql:"Purchase Requisition"`
	ItemOfRequisition           string    `json:"Item of Requisition" sql:"Item of Requisition"`
	DocumentType                string    `json:"Document Type" sql:"Document Type"`
	DocumentTypeName            string    `json:"Document Type Name" sql:"Document Type Name"`
	PurchDocCategory            string    `json:"Purch Doc Category" sql:"Purch Doc Category"`
	Bstyp_Concat                string    `json:"BSTYP_Concat" sql:"BSTYP_Concat"`
	ControlIndicator            string    `json:"Control Indicator" sql:"Control Indicator"`
	Bsakz_Concat                string    `json:"BSAKZ_Concat" sql:"BSAKZ_Concat"`
	DeletionFlag                string    `json:"Deletion Indicator" sql:"Deletion Indicator"`
	ProcessingStatus            string    `json:"Processing Status" sql:"Processing Status"`
	Statu_Concat                string    `json:"STATU_Concat" sql:"STATU_Concat"`
	CreationIndicator           string    `json:"Creation Indicator" sql:"Creation Indicator"`
	Estkz_Concat                string    `json:"ESTKZ_Concat" sql:"ESTKZ_Concat"`
	ReleaseIndicator            string    `json:"Release Indicator" sql:"Release Indicator"`
	ReleaseIndicatorName        string    `json:"Release Indicator Name" sql:"Release Indicator Name"`
	ReleaseStatus               string    `json:"Release Status" sql:"Release Status"`
	ReleaseStrategy             string    `json:"Release Strategy" sql:"Release Strategy"`
	PurchasingGroup             string    `json:"Purchasing Group" sql:"Purchasing Group"`
	PurchasingGroupName         string    `json:"Purchasing Group Name" sql:"Purchasing Group Name"`
	CreatedBy                   string    `json:"Created By" sql:"Created By"`
	CreationDate                time.Time `json:"Created On" sql:"Created On"`
	Requisitioner               string    `json:"Requisitioner" sql:"Requisitioner"`
	ShortText                   string    `json:"Short Text" sql:"Short Text"`
	MaterialCode                string    `json:"Material Code" sql:"Material Code"`
	MpnMaterial                 string    `json:"MPN Material" sql:"MPN Material"`
	Plant                       string    `json:"Plant" sql:"Plant"`
	PlantName                   string    `json:"Plant Name" sql:"Plant Name"`
	PlantPurchOrganization      string    `json:"Plant Purch Organization" sql:"Plant Purch Organization"`
	PlantPurchOrganizationName  string    `json:"Plant Purch Organization Name" sql:"Plant Purch Organization Name"`
	PlantCompanyCode            string    `json:"Plant Company Code" sql:"Plant Company Code"`
	StorageLocation             string    `json:"Storage Location" sql:"Storage Location"`
	ReqTrackingNumber           string    `json:"Req Tracking Number" sql:"Req Tracking Number"`
	MaterialGroup               string    `json:"Material Group" sql:"Material Group"`
	MaterialGroupName           string    `json:"Material Group Name" sql:"Material Group Name"`
	SupplyingPlant              string    `json:"Supplying Plant" sql:"Supplying Plant"`
	QtyBase                     float64   `json:"Qty Base" sql:"Qty Base"`
	UomBase                     string    `json:"Uom Base" sql:"Uom Base"`
	ShortageQuantity            float64   `json:"Shortage Quantity" sql:"Shortage Quantity"`
	RequisitionDate             time.Time `json:"Requisition Date" sql:"Requisition Date"`
	DeliveryDateCategory        string    `json:"Delivery Date Category" sql:"Delivery Date Category"`
	DeliveryDate                time.Time `json:"Delivery Date" sql:"Delivery Date"`
	ReleaseDate                 time.Time `json:"Release Date" sql:"Release Date"`
	GrProcessingTime            float64   `json:"GR Processing Time" sql:"GR Processing Time"`
	PriceDc                     float64   `json:"Price DC" sql:"Price DC"`
	PriceLc                     float64   `json:"Price LC" sql:"Price LC"`
	Price_Usd                   float64   `json:"Price_USD" sql:"Price_USD"`
	PriceUnit                   float64   `json:"Price Unit" sql:"Price Unit"`
	ValueInDc                   float64   `json:"Value in DC" sql:"Value in DC"`
	ValueInLc                   float64   `json:"Value in LC" sql:"Value in LC"`
	ValueInUsd                  float64   `json:"Value in USD" sql:"Value in USD"`
	ItemCategory                string    `json:"Item Category" sql:"Item Category"`
	AcctAssignmentCat           string    `json:"Acct Assignment Cat" sql:"Acct Assignment Cat"`
	AcctAssignmentCatName       string    `json:"Acct Assignment Cat Name" sql:"Acct Assignment Cat Name"`
	Consumption                 string    `json:"Consumption" sql:"Consumption"`
	Kzvbr_Concat                string    `json:"KZVBR_Concat" sql:"KZVBR_Concat"`
	DistribIndicator            string    `json:"Distrib Indicator" sql:"Distrib Indicator"`
	PartialInvoice              string    `json:"Partial Invoice" sql:"Partial Invoice"`
	GoodsReceipt                string    `json:"Goods Receipt" sql:"Goods Receipt"`
	GrNonValuated               string    `json:"GR Non-Valuated" sql:"GR Non-Valuated"`
	InvoiceReceipt              string    `json:"Invoice Receipt" sql:"Invoice Receipt"`
	VendorCode                  string    `json:"LIFNR" sql:"LIFNR"`
	DesiredVendorCountryKey     string    `json:"Desired Vendor Country Key" sql:"Desired Vendor Country Key"`
	VendorName1                 string    `json:"Vendor Name 1" sql:"Vendor Name 1"`
	DesiredVendorName2          string    `json:"Desired Vendor Name 2" sql:"Desired Vendor Name 2"`
	DesiredVendorCreatedBy      string    `json:"Desired Vendor Created By" sql:"Desired Vendor Created By"`
	VendorOneTimeAccount        string    `json:"Vendor One-Time Account" sql:"Vendor One-Time Account"`
	DesiredVendorTradingPartner string    `json:"Desired Vendor Trading Partner" sql:"Desired Vendor Trading Partner"`
	FixedVendor                 string    `json:"Fixed Vendor" sql:"Fixed Vendor"`
	FixedVendorCountryKey       string    `json:"Fixed Vendor Country Key" sql:"Fixed Vendor Country Key"`
	FixedVendorName1            string    `json:"Fixed Vendor Name 1" sql:"Fixed Vendor Name 1"`
	FixedVendorName2            string    `json:"Fixed Vendor Name 2" sql:"Fixed Vendor Name 2"`
	FixedVendorCreatedBy        string    `json:"Fixed Vendor Created By" sql:"Fixed Vendor Created By"`
	FixedVendorOneTimeAccount   string    `json:"Fixed Vendor One-Time Account" sql:"Fixed Vendor One-Time Account"`
	FixedVendorTrdingPartner    string    `json:"Fixed Vendor Trding Partner" sql:"Fixed Vendor Trding Partner"`
	PurchOrganization           string    `json:"Purch Organization" sql:"Purch Organization"`
	PurchOrganizationName       string    `json:"Purch Organization Name" sql:"Purch Organization Name"`
	PurchasingInfoRec           string    `json:"Purchasing Info Rec" sql:"Purchasing Info Rec"`
	Ebeln                       string    `json:"EBELN" sql:"EBELN"`
	Ebelp                       string    `json:"EBELP" sql:"EBELP"`
	PurchaseOrderDate           time.Time `json:"Purchase Order Date" sql:"Purchase Order Date"`
	QuantityOrdered             float64   `json:"Quantity Ordered" sql:"Quantity Ordered"`
	ValuationType               string    `json:"Valuation Type" sql:"Valuation Type"`
	Closed                      string    `json:"Closed" sql:"Closed"`
	Reservation                 string    `json:"Reservation" sql:"Reservation"`
	SpecialStock                string    `json:"Special Stock" sql:"Special Stock"`
	OrderUnit                   string    `json:"Order Unit" sql:"Order Unit"`
	RevisionLevel               string    `json:"Revision Level" sql:"Revision Level"`
	AdvanceProcurement          string    `json:"Advance Procurement" sql:"Advance Procurement"`
	ReleaseGroup                string    `json:"Release Group" sql:"Release Group"`
	ReleaseGroupName            string    `json:"Release Group Name" sql:"Release Group Name"`
	SubjectToRelease            string    `json:"Subject to release" sql:"Subject to release"`
	ManualAddress               string    `json:"Manual Address" sql:"Manual Address"`
	DeliveryAddress             string    `json:"Delivery Address" sql:"Delivery Address"`
	Kunnr                       string    `json:"KUNNR" sql:"KUNNR"`
	SpecStkValuation            string    `json:"Spec Stk Valuation" sql:"Spec Stk Valuation"`
	Waers                       string    `json:"WAERS" sql:"WAERS"`
	CurrencyLocal               string    `json:"Currency Local" sql:"Currency Local"`
	ExchangeRateToLc            float64   `json:"Exchange Rate to LC" sql:"Exchange Rate to LC"`
	ExchangeRateToUsd           float64   `json:"Exchange Rate to USD" sql:"Exchange Rate to USD"`
	PlannedDelivTime            float64   `json:"Planned Deliv Time" sql:"Planned Deliv Time"`
	FunctionalArea              string    `json:"Functional Area" sql:"Functional Area"`
	Grant                       string    `json:"Grant" sql:"Grant"`
	PreqProcessingState         string    `json:"PReq Processing State" sql:"PReq Processing State"`
	TotalValUponRelease         float64   `json:"Total Val Upon Release" sql:"Total Val Upon Release"`
	Version                     string    `json:"Version" sql:"Version"`
	ProcuringPlant              string    `json:"Procuring Plant" sql:"Procuring Plant"`
	Prctr                       string    `json:"PRCTR" sql:"PRCTR"`
	Isopen                      bool      `json:"isOpen" sql:"isOpen"`
	ApprovalsRequired           int32     `json:"Approvals Required" sql:"Approvals Required"`
	ID                          string    `json:"ICM_VENDOR_ID" sql:"ICM_VENDOR_ID"`
	Org1                        string    `json:"ORG1" sql:"ORG1"`
	Org1_Concat                 string    `json:"ORG1_Concat" sql:"ORG1_Concat"`
	Org2                        string    `json:"ORG2" sql:"ORG2"`
	Org2_Concat                 string    `json:"ORG2_Concat" sql:"ORG2_Concat"`
	Org3                        string    `json:"ORG3" sql:"ORG3"`
	Org3_Concat                 string    `json:"ORG3_Concat" sql:"ORG3_Concat"`
	AbacusCode                  string    `json:"Abacus Code" sql:"Abacus Code"`
}

func (r Requisition) GetFlags() ICMEntity {
	return nil
}
