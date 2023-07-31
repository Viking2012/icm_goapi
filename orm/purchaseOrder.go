package orm

import "time"

type PurchaseOrder struct {
	Database                                          string             `json:"Database" sql:"Database"`
	PoNumber                                          string             `json:"PO Number" sql:"PO Number"`
	PurchasingDocumentLineItem                        string             `json:"Purchasing Document Line Item" sql:"Purchasing Document Line Item"`
	AcctAssignmentLine                                string             `json:"Acct Assignment Line" sql:"Acct Assignment Line"`
	CompanyCodeId                                     string             `json:"Company Code ID" sql:"Company Code ID"`
	LocalCurrency                                     string             `json:"Local Currency" sql:"Local Currency"`
	LocalLanguage                                     string             `json:"Local Language" sql:"Local Language"`
	ChartOfAccounts                                   string             `json:"Chart of Accounts" sql:"Chart of Accounts"`
	CompanyCode                                       string             `json:"Company Code" sql:"Company Code"`
	PurchasingDocumentCategoryCode                    string             `json:"Purchasing Document Category Code" sql:"Purchasing Document Category Code"`
	PurchasingDocumentCategory                        string             `json:"Purchasing Document Category" sql:"Purchasing Document Category"`
	PurchasingDocumentTypeCode                        string             `json:"Purchasing Document Type Code" sql:"Purchasing Document Type Code"`
	PurchasingDocumentType                            string             `json:"Purchasing Document Type" sql:"Purchasing Document Type"`
	ControlIndicatorForPurchasingDocumentType         string             `json:"Control indicator for purchasing document type" sql:"Control indicator for purchasing document type"`
	PoDeletionIndicator                               string             `json:"PO Deletion Indicator" sql:"PO Deletion Indicator"`
	PoStatusCode                                      string             `json:"PO Status Code" sql:"PO Status Code"`
	PoStatus                                          string             `json:"PO Status" sql:"PO Status"`
	PoCreationDate                                    time.Time          `json:"PO Creation Date" sql:"PO Creation Date"`
	PoCreatedBy                                       string             `json:"PO Created By" sql:"PO Created By"`
	VendorNumber                                      string             `json:"Vendor Number" sql:"Vendor Number"`
	Vendor                                            string             `json:"Vendor" sql:"Vendor"`
	CountryCode                                       string             `json:"Vendor Country Code" sql:"Vendor Country Code"`
	AccountGroupCode                                  string             `json:"Vendor Account Group Code" sql:"Vendor Account Group Code"`
	AccountGroup                                      string             `json:"Vendor Account Group" sql:"Vendor Account Group"`
	TermsOfPaymentKey                                 string             `json:"Terms of Payment Key" sql:"Terms of Payment Key"`
	CashPromptPaymentDiscount1Days                    int32              `json:"Cash Prompt Payment Discount (1) Days" sql:"Cash Prompt Payment Discount (1) Days"`
	CashPromptPaymentDiscount2Days                    int32              `json:"Cash Prompt Payment Discount (2) Days" sql:"Cash Prompt Payment Discount (2) Days"`
	CashPromptPaymentDiscount3Days                    int32              `json:"Cash Prompt Payment Discount (3) Days" sql:"Cash Prompt Payment Discount (3) Days"`
	CashDiscountPercentage1                           float64            `json:"Cash Discount Percentage 1" sql:"Cash Discount Percentage 1"`
	CashDiscountPercentage2                           float64            `json:"Cash Discount Percentage 2" sql:"Cash Discount Percentage 2"`
	PurchasingOrganizationCode                        string             `json:"Purchasing Organization Code" sql:"Purchasing Organization Code"`
	PurchasingOrganization                            string             `json:"Purchasing Organization" sql:"Purchasing Organization"`
	PurchasingGroupCode                               string             `json:"Purchasing Group Code" sql:"Purchasing Group Code"`
	PurchasingGroup                                   string             `json:"Purchasing Group" sql:"Purchasing Group"`
	PoCurrency                                        string             `json:"PO Currency" sql:"PO Currency"`
	ExchangeRate                                      float64            `json:"Exchange Rate" sql:"Exchange Rate"`
	PurchasingDocumentDate                            time.Time          `json:"Purchasing Document Date" sql:"Purchasing Document Date"`
	SupplyingVendorCode                               string             `json:"Supplying Vendor Code" sql:"Supplying Vendor Code"`
	NumberOfPrincipalPurchaseAgreementPOHeader        string             `json:"Number of Principal Purchase Agreement (PO Header)" sql:"Number of Principal Purchase Agreement (PO Header)"`
	IncotermsPart1PO                                  string             `json:"Incoterms part 1 (PO)" sql:"Incoterms part 1 (PO)"`
	IncotermsPart2PO                                  string             `json:"Incoterms part 2 (PO)" sql:"Incoterms part 2 (PO)"`
	ReleaseGroup                                      string             `json:"Release Group" sql:"Release Group"`
	ReleaseStrategy                                   string             `json:"Release Strategy" sql:"Release Strategy"`
	PurchasingDocumentReleaseIndicator                string             `json:"Purchasing Document Release Indicator" sql:"Purchasing Document Release Indicator"`
	PoReleaseStatus                                   string             `json:"PO Release Status" sql:"PO Release Status"`
	MostRecentApproverCode                            string             `json:"Most Recent Approver Code" sql:"Most Recent Approver Code"`
	MostRecentApprover                                string             `json:"Most Recent Approver" sql:"Most Recent Approver"`
	NextApproverCode                                  string             `json:"Next Approver Code" sql:"Next Approver Code"`
	NextApprover                                      string             `json:"Next Approver" sql:"Next Approver"`
	ReleaseNotYetCompletelyEffected                   string             `json:"Release Not Yet Completely Effected" sql:"Release Not Yet Completely Effected"`
	AddressId                                         string             `json:"Address ID" sql:"Address ID"`
	VatNumber                                         string             `json:"VAT Number" sql:"VAT Number"`
	DeletionIndicatorInPurchasingDocumentItem         string             `json:"Deletion Indicator in Purchasing Document Item" sql:"Deletion Indicator in Purchasing Document Item"`
	RfqStatus                                         string             `json:"RFQ status" sql:"RFQ status"`
	PurchasingDocumentItemChangeDate                  time.Time          `json:"Purchasing Document Item Change Date" sql:"Purchasing Document Item Change Date"`
	ShortTextForPurchasingDocumentItem                string             `json:"Short Text for Purchasing Document Item" sql:"Short Text for Purchasing Document Item"`
	MaterialNumber                                    string             `json:"Material Number" sql:"Material Number"`
	Ekpo_Matnr_Concat                                 string             `json:"EKPO_MATNR_Concat" sql:"EKPO_MATNR_Concat"`
	Mpn                                               string             `json:"MPN" sql:"MPN"`
	PlantCode                                         string             `json:"Plant Code" sql:"Plant Code"`
	Plant                                             string             `json:"Plant" sql:"Plant"`
	StorageLocation                                   string             `json:"Storage Location" sql:"Storage Location"`
	MaterialGroupCode                                 string             `json:"Material Group Code" sql:"Material Group Code"`
	MaterialGroup                                     string             `json:"Material Group" sql:"Material Group"`
	PoInfoRecord                                      string             `json:"PO Info Record" sql:"PO Info Record"`
	VendorMaterialNumber                              string             `json:"Vendor Material Number" sql:"Vendor Material Number"`
	PurchaseOrderUnitOfMeasure                        string             `json:"Purchase Order Unit of Measure" sql:"Purchase Order Unit of Measure"`
	OverdeliveryToleranceLimit                        float64            `json:"Overdelivery Tolerance Limit" sql:"Overdelivery Tolerance Limit"`
	IndicatorUnlimitedOverdeliveryAllowed             string             `json:"Indicator Unlimited Overdelivery Allowed" sql:"Indicator Unlimited Overdelivery Allowed"`
	UnderdeliveryToleranceLimit                       float64            `json:"Underdelivery Tolerance Limit" sql:"Underdelivery Tolerance Limit"`
	DeliveryCompletedIndicator                        string             `json:"Delivery Completed Indicator" sql:"Delivery Completed Indicator"`
	FinalInvoiceIndicator                             string             `json:"Final Invoice Indicator" sql:"Final Invoice Indicator"`
	ItemCategoryInPurchasingDocument                  string             `json:"Item Category in Purchasing Document" sql:"Item Category in Purchasing Document"`
	PoLineItemAccountAssignmentCategory               string             `json:"PO Line Item Account Assignment Category" sql:"PO Line Item Account Assignment Category"`
	ConsumptionPosting                                string             `json:"Consumption Posting" sql:"Consumption Posting"`
	DistributionIndicatorForMultipleAccountAssignment string             `json:"Distribution indicator for multiple account assignment" sql:"Distribution indicator for multiple account assignment"`
	GoodsReceiptIndicator                             string             `json:"Goods Receipt Indicator" sql:"Goods Receipt Indicator"`
	GoodsReceiptNonValuated                           string             `json:"Goods Receipt Non-Valuated" sql:"Goods Receipt Non-Valuated"`
	InvoiceReceiptIndicator                           string             `json:"Invoice Receipt Indicator" sql:"Invoice Receipt Indicator"`
	GrBasedInvoiceVerification                        string             `json:"GR-Based Invoice Verification" sql:"GR-Based Invoice Verification"`
	AcknowledgementRequired                           string             `json:"Acknowledgement Required" sql:"Acknowledgement Required"`
	PoLineAddressId                                   string             `json:"PO Line Address ID" sql:"PO Line Address ID"`
	Ekpo_Adrn2                                        string             `json:"EKPO_ADRN2" sql:"EKPO_ADRN2"`
	ConfirmationControl                               string             `json:"Confirmation Control" sql:"Confirmation Control"`
	Incoterms1                                        string             `json:"Incoterms 1" sql:"Incoterms 1"`
	Incoterms2                                        string             `json:"Incoterms 2" sql:"Incoterms 2"`
	Statistical                                       string             `json:"Statistical" sql:"Statistical"`
	PrNumber                                          string             `json:"PR Number" sql:"PR Number"`
	PrItem                                            string             `json:"PR Item" sql:"PR Item"`
	ReturnsItem                                       string             `json:"Returns Item" sql:"Returns Item"`
	FirstDeliveryDate                                 time.Time          `json:"First Delivery Date" sql:"First Delivery Date"`
	LastDeliveryDate                                  time.Time          `json:"Last Delivery Date" sql:"Last Delivery Date"`
	Eket_Slfdt_Min                                    time.Time          `json:"EKET_SLFDT_Min" sql:"EKET_SLFDT_Min"`
	Eket_Slfdt_Max                                    time.Time          `json:"EKET_SLFDT_Max" sql:"EKET_SLFDT_Max"`
	ScheduledQuantity                                 float64            `json:"Scheduled Quantity" sql:"Scheduled Quantity"`
	AcctAssignmentDeleted                             string             `json:"Acct Assignment Deleted" sql:"Acct Assignment Deleted"`
	AcctAssignmentCreatedOn                           time.Time          `json:"Acct Assignment Created On" sql:"Acct Assignment Created On"`
	Ekkn_Vproz                                        float64            `json:"EKKN_VPROZ" sql:"EKKN_VPROZ"`
	GlAccountId                                       string             `json:"GL Account ID" sql:"GL Account ID"`
	CostCenter                                        string             `json:"Cost Center" sql:"Cost Center"`
	Ekkn_Aufnr                                        string             `json:"EKKN_AUFNR" sql:"EKKN_AUFNR"`
	Ekkn_Kstrg                                        string             `json:"EKKN_KSTRG" sql:"EKKN_KSTRG"`
	WbsElementId                                      string             `json:"WBS Element ID" sql:"WBS Element ID"`
	ProjectNumber                                     string             `json:"Project Number" sql:"Project Number"`
	ProjectName                                       string             `json:"PROJ_Concat" sql:"PROJ_Concat"`
	Network                                           string             `json:"Network" sql:"Network"`
	RoutingLine                                       string             `json:"Routing Line" sql:"Routing Line"`
	RoutingNumber                                     string             `json:"Routing Number" sql:"Routing Number"`
	Ekkn_Vbeln                                        string             `json:"EKKN_VBELN" sql:"EKKN_VBELN"`
	Ekkn_Vbelp                                        string             `json:"EKKN_VBELP" sql:"EKKN_VBELP"`
	Ekkn_Anln1                                        string             `json:"EKKN_ANLN1" sql:"EKKN_ANLN1"`
	Ekkn_Anln2                                        string             `json:"EKKN_ANLN2" sql:"EKKN_ANLN2"`
	Quantity                                          float64            `json:"Quantity" sql:"Quantity"`
	Ekkn_Qty_Base                                     float64            `json:"EKKN_Qty_Base" sql:"EKKN_Qty_Base"`
	Value                                             float64            `json:"Value" sql:"Value"`
	ValueInLc                                         float64            `json:"Value in LC" sql:"Value in LC"`
	ValueInUsd                                        float64            `json:"Value in USD" sql:"Value in USD"`
	ExchangeRateToLc                                  float64            `json:"Exchange Rate to LC" sql:"Exchange Rate to LC"`
	ExchangeRateToUsd                                 float64            `json:"Exchange Rate to USD" sql:"Exchange Rate to USD"`
	AcctAssignmentRatio                               float64            `json:"Acct Assignment Ratio" sql:"Acct Assignment Ratio"`
	ProfitCenter                                      string             `json:"Profit Center" sql:"Profit Center"`
	Profitcenter_Org1                                 string             `json:"ProfitCenter_ORG1" sql:"ProfitCenter_ORG1"`
	Profitcenter_Org1_Concat                          string             `json:"ProfitCenter_ORG1_Concat" sql:"ProfitCenter_ORG1_Concat"`
	Profitcenter_Org2                                 string             `json:"ProfitCenter_ORG2" sql:"ProfitCenter_ORG2"`
	Profitcenter_Org2_Concat                          string             `json:"ProfitCenter_ORG2_Concat" sql:"ProfitCenter_ORG2_Concat"`
	Profitcenter_Org3                                 string             `json:"ProfitCenter_ORG3" sql:"ProfitCenter_ORG3"`
	Profitcenter_Org3_Concat                          string             `json:"ProfitCenter_ORG3_Concat" sql:"ProfitCenter_ORG3_Concat"`
	AbacusCountry                                     string             `json:"Abacus Country" sql:"Abacus Country"`
	ReportingUnit                                     string             `json:"Reporting Unit" sql:"Reporting Unit"`
	InvoicedQuantity                                  float64            `json:"Invoiced Quantity" sql:"Invoiced Quantity"`
	InvoicedValueInLc                                 float64            `json:"Invoiced Value in LC" sql:"Invoiced Value in LC"`
	FirstInvoicePostingDate                           time.Time          `json:"First Invoice Posting Date" sql:"First Invoice Posting Date"`
	FirstInvoiceCreatedOn                             time.Time          `json:"First Invoice Created On" sql:"First Invoice Created On"`
	LastInvoicePostingDate                            time.Time          `json:"Last Invoice Posting Date" sql:"Last Invoice Posting Date"`
	LastInvoiceCreatedOn                              time.Time          `json:"Last Invoice Created On" sql:"Last Invoice Created On"`
	FirstInvoiceDocDate                               time.Time          `json:"First Invoice Doc Date" sql:"First Invoice Doc Date"`
	LastInvoiceDocDate                                time.Time          `json:"Last Invoice Doc Date" sql:"Last Invoice Doc Date"`
	ReceiptedQuantity                                 float64            `json:"Receipted Quantity" sql:"Receipted Quantity"`
	ReceiptedValueInLc                                float64            `json:"Receipted Value in LC" sql:"Receipted Value in LC"`
	FirstReceiptPostingDate                           time.Time          `json:"First Receipt Posting Date" sql:"First Receipt Posting Date"`
	FirstReceiptCreatedOn                             time.Time          `json:"First Receipt Created On" sql:"First Receipt Created On"`
	LastReceiptPostingDate                            time.Time          `json:"Last Receipt Posting Date" sql:"Last Receipt Posting Date"`
	LastReceiptCreatedOn                              time.Time          `json:"Last Receipt Created On" sql:"Last Receipt Created On"`
	FirstReceiptDocDate                               time.Time          `json:"First Receipt Doc Date" sql:"First Receipt Doc Date"`
	LastReceiptDocDate                                time.Time          `json:"Last Receipt Doc Date" sql:"Last Receipt Doc Date"`
	ApprovalsRequired                                 int32              `json:"Approvals Required" sql:"Approvals Required"`
	ID                                                string             `json:"ICM_VENDOR_ID" sql:"ICM_VENDOR_ID"`
	AprovalDescriptive                                string             `json:"Aproval Descriptive" sql:"Aproval Descriptive"`
	TotalApprovals                                    int32              `json:"Total Approvals" sql:"Total Approvals"`
	Flags                                             PurchaseOrderFlags `json:"flags" sql:"FLAGS"`
}

func (po PurchaseOrder) GetFlags() ICMEntity {
	return po.Flags
}

type PurchaseOrderFlags struct {
	FlagNoReceiptNeeded          bool `json:"Flag: NoReceiptNeeded" sql:"Flag: NoReceiptNeeded"`
	FlagNoinvoiceneeded          bool `json:"Flag: NoInvoiceNeeded" sql:"Flag: NoInvoiceNeeded"`
	FlagNoinvoiceorreceiptneeded bool `json:"Flag: NoInvoiceOrReceiptNeeded" sql:"Flag: NoInvoiceOrReceiptNeeded"`
}

func (pof PurchaseOrderFlags) GetFlags() ICMEntity {
	return pof
}