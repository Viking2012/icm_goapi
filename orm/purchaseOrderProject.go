package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type PurchaseOrderProject struct {
	ProjectId                          string                     `json:"Project ID" sql:"Project ID"`
	ProjectNumber                      string                     `json:"Project Number" sql:"Project Number"`
	ProjectName                        string                     `json:"Project Name" sql:"Project Name"`
	ProjectResponsiblePerson           string                     `json:"Project Responsible Person" sql:"Project Responsible Person"`
	ProjectCountry                     string                     `json:"Project Country" sql:"Project Country"`
	ProjectReleasedOn                  sql.NullTime               `json:"Project Released On" sql:"Project Released On"`
	ProjectReleaseYear                 string                     `json:"Project Release Year" sql:"Project Release Year"`
	Database                           string                     `json:"Database" sql:"Database"`
	PoNumber                           string                     `json:"PO Number" sql:"PO Number"`
	PurchasingDocumentLineItem         string                     `json:"Purchasing Document Line Item" sql:"Purchasing Document Line Item"`
	AcctAssignmentLine                 string                     `json:"Acct Assignment Line" sql:"Acct Assignment Line"`
	CompanyCodeId                      string                     `json:"Company Code ID" sql:"Company Code ID"`
	PoStatusCode                       string                     `json:"PO Status Code" sql:"PO Status Code"`
	PoStatus                           string                     `json:"PO Status" sql:"PO Status"`
	PoCreationDate                     sql.NullTime               `json:"PO Creation Date" sql:"PO Creation Date"`
	VendorNumber                       string                     `json:"Vendor Number" sql:"Vendor Number"`
	Vendor                             string                     `json:"Vendor" sql:"Vendor"`
	PurchasingDocumentDate             sql.NullTime               `json:"Purchasing Document Date" sql:"Purchasing Document Date"`
	PurchasingDocumentItemChangeDate   sql.NullTime               `json:"Purchasing Document Item Change Date" sql:"Purchasing Document Item Change Date"`
	ShortTextForPurchasingDocumentItem string                     `json:"Short Text for Purchasing Document Item" sql:"Short Text for Purchasing Document Item"`
	MaterialNumber                     string                     `json:"Material Number" sql:"Material Number"`
	MaterialGroupCode                  string                     `json:"Material Group Code" sql:"Material Group Code"`
	MaterialGroup                      string                     `json:"Material Group" sql:"Material Group"`
	AcctAssignmentCreatedOn            sql.NullTime               `json:"Acct Assignment Created On" sql:"Acct Assignment Created On"`
	Right_Proj_Pspid                   string                     `json:"Right_PROJ_PSPID" sql:"Right_PROJ_PSPID"`
	ValueInLc                          float64                    `json:"Value in LC" sql:"Value in LC"`
	ValueInUsd                         float64                    `json:"Value in USD" sql:"Value in USD"`
	InvoicedQuantity                   float64                    `json:"Invoiced Quantity" sql:"Invoiced Quantity"`
	ReceiptedQuantity                  float64                    `json:"Receipted Quantity" sql:"Receipted Quantity"`
	Firstpodatewbs                     sql.NullTime               `json:"FirstPODateWBS" sql:"FirstPODateWBS"`
	Firstpodate_No_Wbs                 sql.NullTime               `json:"FirstPODate_NO_WBS" sql:"FirstPODate_NO_WBS"`
	BusinesPartner                     string                     `json:"Busines Partner" sql:"Busines Partner"`
	IsIntercompany                     bool                       `json:"Is Intercompany" sql:"Is Intercompany"`
	Left_Left_Right_Ekko_Lifnr         string                     `json:"Left_Left_Right_EKKO_LIFNR" sql:"Left_Left_Right_EKKO_LIFNR"`
	Left_Right_Ekko_Lifnr              string                     `json:"Left_Right_EKKO_LIFNR" sql:"Left_Right_EKKO_LIFNR"`
	TltMaterialSourcingDesc            string                     `json:"TLT: Material Sourcing Desc" sql:"TLT: Material Sourcing Desc"`
	Right_Ekko_Lifnr                   string                     `json:"Right_EKKO_LIFNR" sql:"Right_EKKO_LIFNR"`
	Flags                              *PurchaseOrderProjectFlags `json:"flags" sql:"FLAGS"`
}

func (pop *PurchaseOrderProject) GetFlags() ICMEntity {
	return pop.Flags
}

func PurchaseOrderProjectFromRow(rows *sql.Rows) (ICMEntity, error) {
	var pop PurchaseOrderProject
	err := rows.Scan(
		&pop.ProjectId,
		&pop.ProjectNumber,
		&pop.ProjectName,
		&pop.ProjectResponsiblePerson,
		&pop.ProjectCountry,
		&pop.ProjectReleasedOn,
		&pop.ProjectReleaseYear,
		&pop.Database,
		&pop.PoNumber,
		&pop.PurchasingDocumentLineItem,
		&pop.AcctAssignmentLine,
		&pop.CompanyCodeId,
		&pop.PoStatusCode,
		&pop.PoStatus,
		&pop.PoCreationDate,
		&pop.VendorNumber,
		&pop.Vendor,
		&pop.PurchasingDocumentDate,
		&pop.PurchasingDocumentItemChangeDate,
		&pop.ShortTextForPurchasingDocumentItem,
		&pop.MaterialNumber,
		&pop.MaterialGroupCode,
		&pop.MaterialGroup,
		&pop.AcctAssignmentCreatedOn,
		&pop.Right_Proj_Pspid,
		&pop.ValueInLc,
		&pop.ValueInUsd,
		&pop.InvoicedQuantity,
		&pop.ReceiptedQuantity,
		&pop.Firstpodatewbs,
		&pop.Firstpodate_No_Wbs,
		&pop.BusinesPartner,
		&pop.IsIntercompany,
		&pop.Left_Left_Right_Ekko_Lifnr,
		&pop.Left_Right_Ekko_Lifnr,
		&pop.TltMaterialSourcingDesc,
		&pop.Right_Ekko_Lifnr,
		&pop.Flags,
	)
	return &pop, err
}

type PurchaseOrderProjectFlags struct {
	Wbs_Vendoronly                     bool `json:"WBS_VendorOnly" sql:"WBS_VendorOnly"`
	Mixwbsandnone                      bool `json:"MixWBSandNone" sql:"MixWBSandNone"`
	FlagVendorOnProjectWithoutCustomer bool `json:"Flag: Vendor On Project Without Customer" sql:"Flag: Vendor On Project Without Customer"`
	FlagPoCreatedAfterProjectClosure   bool `json:"Flag: PO Created After Project Closure" sql:"Flag: PO Created After Project Closure"`
	FlagPoCreatedBeforeProjectRelease  bool `json:"Flag: PO Created Before Project Release" sql:"Flag: PO Created Before Project Release"`
	FlagEquallySourced                 bool `json:"Flag: Equally Sourced" sql:"Flag: Equally Sourced"`
	FlagNearlySoleSource               bool `json:"Flag: Nearly Sole Source (>95%)" sql:"Flag: Nearly Sole Source (>95%)"`
	FlagOnePrimarySource               bool `json:"Flag: One Primary Source" sql:"Flag: One Primary Source"`
	FlagOtherSourceDivisionPattern     bool `json:"Flag: Other Source Division Pattern" sql:"Flag: Other Source Division Pattern"`
	FlagAllPosIssuedInSameWeek         bool `json:"Flag: All POs Issued In Same Week" sql:"Flag: All POs Issued In Same Week"`
	FlagAllPosIssuedWithinShortPeriod  bool `json:"Flag: All POs Issued Within Short Period" sql:"Flag: All POs Issued Within Short Period"`
	FlagOtherPoIssuancePattern         bool `json:"Flag: Other PO Issuance Pattern" sql:"Flag: Other PO Issuance Pattern"`
	FlagHasMultiSourcedMaterials       bool `json:"Flag: Has Multi-Sourced Materials" sql:"Flag: Has Multi-Sourced Materials"`
}

func (popf *PurchaseOrderProjectFlags) GetFlags() ICMEntity {
	return popf
}
func (popf *PurchaseOrderProjectFlags) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return json.Unmarshal([]byte(v), popf)
	case []byte:
		return json.Unmarshal(v, popf)
	default:
		return errors.New("invalid sql return type for Vendor Flags")
	}
}
