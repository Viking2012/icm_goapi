package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type SalesOrderProject struct {
	Database                              string                  `json:"Database" sql:"Database"`
	SoldToCode                            string                  `json:"Sold To Code" sql:"Sold To Code"`
	DocumentNumberSales                   string                  `json:"Sales Document" sql:"Sales Document"`
	Item                                  string                  `json:"Item" sql:"Item"`
	CreationDate                          sql.NullTime            `json:"Created On" sql:"Created On"`
	DocumentCategoryCode                  string                  `json:"Document Category Code" sql:"Document Category Code"`
	DocumentCategory                      string                  `json:"Document Category" sql:"Document Category"`
	WbsElementId                          string                  `json:"WBS Element ID" sql:"WBS Element ID"`
	NetValueItemUSD                       float64                 `json:"Net Value in USD" sql:"Net Value in USD"`
	NetValueDocument                      float64                 `json:"Net Value" sql:"Net Value"`
	ValidFrom                             sql.NullTime            `json:"Valid From" sql:"Valid From"`
	ValidTo                               sql.NullTime            `json:"Valid To" sql:"Valid To"`
	DocumentDate                          sql.NullTime            `json:"Document Date" sql:"Document Date"`
	SalesDocumentTypeCode                 string                  `json:"Sales Document Type Code" sql:"Sales Document Type Code"`
	RequestedDeliveryDate                 sql.NullTime            `json:"Requested Delivery Date" sql:"Requested Delivery Date"`
	MaterialCode                          string                  `json:"Material Code" sql:"Material Code"`
	Material                              string                  `json:"Material" sql:"Material"`
	MaterialGroupCode                     string                  `json:"Material Group Code" sql:"Material Group Code"`
	MaterialGroup                         string                  `json:"Material Group" sql:"Material Group"`
	ItemText                              string                  `json:"Item Text" sql:"Item Text"`
	ProjectNumber                         string                  `json:"Project Number" sql:"Project Number"`
	Right_Proj_Pspid                      string                  `json:"Right_PROJ_PSPID" sql:"Right_PROJ_PSPID"`
	ProjectName                           string                  `json:"Project Name" sql:"Project Name"`
	ProjectObject                         string                  `json:"Project Object" sql:"Project Object"`
	ProjectCreatedBy                      string                  `json:"Project Created By" sql:"Project Created By"`
	ProjectCreatedOn                      sql.NullTime            `json:"Project Created On" sql:"Project Created On"`
	ProjectChangedBy                      string                  `json:"Project Changed By" sql:"Project Changed By"`
	ProjectChangedOn                      sql.NullTime            `json:"Project Changed On" sql:"Project Changed On"`
	ProjectStatusProfileCode              string                  `json:"Project Status Profile Code" sql:"Project Status Profile Code"`
	ProjectStatusProfile                  string                  `json:"Project Status Profile" sql:"Project Status Profile"`
	WbsStatusProfileCode                  string                  `json:"WBS Status Profile Code" sql:"WBS Status Profile Code"`
	ProjectResponsiblePerson              string                  `json:"Project Responsible Person" sql:"Project Responsible Person"`
	ProjectApplicantPerson                string                  `json:"Project Applicant Person" sql:"Project Applicant Person"`
	ProjectCompanyCode                    string                  `json:"Project Company Code" sql:"Project Company Code"`
	ProjectBusinessArea                   string                  `json:"Project Business Area" sql:"Project Business Area"`
	ProjectControllingArea                string                  `json:"Project Controlling Area" sql:"Project Controlling Area"`
	ProjectProfitCenterCode               string                  `json:"Project Profit Center Code" sql:"Project Profit Center Code"`
	Proj_Prctr_Org1                       string                  `json:"PROJ_PRCTR_ORG1" sql:"PROJ_PRCTR_ORG1"`
	Proj_Prctr_Org1_Concat                string                  `json:"PROJ_PRCTR_ORG1_Concat" sql:"PROJ_PRCTR_ORG1_Concat"`
	Proj_Prctr_Org2                       string                  `json:"PROJ_PRCTR_ORG2" sql:"PROJ_PRCTR_ORG2"`
	Proj_Prctr_Org2_Concat                string                  `json:"PROJ_PRCTR_ORG2_Concat" sql:"PROJ_PRCTR_ORG2_Concat"`
	Proj_Prctr_Org3                       string                  `json:"PROJ_PRCTR_ORG3" sql:"PROJ_PRCTR_ORG3"`
	Proj_Prctr_Org3_Concat                string                  `json:"PROJ_PRCTR_ORG3_Concat" sql:"PROJ_PRCTR_ORG3_Concat"`
	ProjectCountry                        string                  `json:"Project Country" sql:"Project Country"`
	ProjectAbacusCode                     string                  `json:"Project Abacus Code" sql:"Project Abacus Code"`
	ProjectCurrency                       string                  `json:"Project Currency" sql:"Project Currency"`
	NetworkAssignment                     string                  `json:"Network Assignment" sql:"Network Assignment"`
	ProjectPlannedStart                   sql.NullTime            `json:"Project Planned Start" sql:"Project Planned Start"`
	ProjectPlannedEnd                     sql.NullTime            `json:"Project Planned End" sql:"Project Planned End"`
	ProjectPlantCode                      string                  `json:"Project Plant Code" sql:"Project Plant Code"`
	ProjectPlant                          string                  `json:"Project Plant" sql:"Project Plant"`
	ProjectPlanningMethodForBasicDates    string                  `json:"Project Planning Method for Basic Dates" sql:"Project Planning Method for Basic Dates"`
	ProjectPlanningMethodForForecastDates string                  `json:"Project Planning Method for Forecast Dates" sql:"Project Planning Method for Forecast Dates"`
	Application                           string                  `json:"Application" sql:"Application"`
	NetworkProfileCode                    string                  `json:"Network Profile Code" sql:"Network Profile Code"`
	NetworkProfile                        string                  `json:"Network Profile" sql:"Network Profile"`
	ProjectProfileCode                    string                  `json:"Project Profile Code" sql:"Project Profile Code"`
	ProjectProfile                        string                  `json:"Project Profile" sql:"Project Profile"`
	BudgetProfileCode                     string                  `json:"Budget Profile Code" sql:"Budget Profile Code"`
	CurrentProjectNumber                  string                  `json:"Current Project Number" sql:"Current Project Number"`
	ProjectStockIndicator                 string                  `json:"Project Stock Indicator" sql:"Project Stock Indicator"`
	ProjectObjectClassCode                string                  `json:"Project Object Class Code" sql:"Project Object Class Code"`
	ProjectObjectClass                    string                  `json:"Project Object Class" sql:"Project Object Class"`
	StatisticalIndicator                  string                  `json:"Statistical Indicator" sql:"Statistical Indicator"`
	DeletionFlag                          string                  `json:"Deletion Indicator" sql:"Deletion Indicator"`
	InactiveIndicator                     string                  `json:"Inactive Indicator" sql:"Inactive Indicator"`
	FunctionalAreaCode                    string                  `json:"Functional Area Code" sql:"Functional Area Code"`
	SalesOrgCode                          string                  `json:"Sales Org Code" sql:"Sales Org Code"`
	DistributionChannelCode               string                  `json:"Distribution Channel Code" sql:"Distribution Channel Code"`
	DivisionCode                          string                  `json:"Division Code" sql:"Division Code"`
	ProjectReleasedBy                     string                  `json:"Project Released By" sql:"Project Released By"`
	ProjectReleasedOn                     sql.NullTime            `json:"Project Released On" sql:"Project Released On"`
	Proj_Firsttecoby                      string                  `json:"PROJ_FirstTecoBy" sql:"PROJ_FirstTecoBy"`
	Proj_Firsttecoon                      sql.NullTime            `json:"PROJ_FirstTecoOn" sql:"PROJ_FirstTecoOn"`
	ProjectTecoBy                         string                  `json:"Project TECO By" sql:"Project TECO By"`
	ProjectTecoOn                         sql.NullTime            `json:"Project TECO On" sql:"Project TECO On"`
	Proj_Tecocount                        int64                   `json:"PROJ_TecoCount" sql:"PROJ_TecoCount"`
	Proj_Firstclsdby                      string                  `json:"PROJ_FirstClsdBy" sql:"PROJ_FirstClsdBy"`
	Proj_Firstclsdon                      sql.NullTime            `json:"PROJ_FirstClsdOn" sql:"PROJ_FirstClsdOn"`
	ProjectClsdBy                         string                  `json:"Project CLSD By" sql:"Project CLSD By"`
	ProjectClsdOn                         sql.NullTime            `json:"Project CLSD On" sql:"Project CLSD On"`
	Proj_Clsdcount                        int64                   `json:"PROJ_ClsdCount" sql:"PROJ_ClsdCount"`
	ID                                    string                  `json:"ICM_ID" sql:"ICM_ID"`
	ProjectReleaseYear                    string                  `json:"Project Release Year" sql:"Project Release Year"`
	Perprojectcount_Customers             int64                   `json:"PerProjectCount_Customers" sql:"PerProjectCount_Customers"`
	Maincustomer                          bool                    `json:"MainCustomer" sql:"MainCustomer"`
	AccountGroup                          string                  `json:"Customer Account Group" sql:"Customer Account Group"`
	GUID                                  string                  `json:"Vendor GUID" sql:"Vendor GUID"`
	IsIntercompany                        bool                    `json:"Is Intercompany" sql:"Is Intercompany"`
	FlagChangeOrder                       bool                    `json:"Flag: Change Order" sql:"Flag: Change Order"`
	ShipToCountryCode                     string                  `json:"Ship To Country Code" sql:"Ship To Country Code"`
	SoldToText                            string                  `json:"Sold To Text" sql:"Sold To Text"`
	Flags                                 *SalesOrderProjectFlags `json:"flags" sql:"FLAGS"`
}

func (sop *SalesOrderProject) GetFlags() ICMEntity {
	return sop.Flags
}

func SalesOrderProjectFromRow(rows *sql.Rows) (ICMEntity, error) {
	var sop SalesOrderProject
	err := rows.Scan(
		&sop.Database,
		&sop.SoldToCode,
		&sop.DocumentNumberSales,
		&sop.Item,
		&sop.CreationDate,
		&sop.DocumentCategoryCode,
		&sop.DocumentCategory,
		&sop.WbsElementId,
		&sop.NetValueItemUSD,
		&sop.NetValueDocument,
		&sop.ValidFrom,
		&sop.ValidTo,
		&sop.DocumentDate,
		&sop.SalesDocumentTypeCode,
		&sop.RequestedDeliveryDate,
		&sop.MaterialCode,
		&sop.Material,
		&sop.MaterialGroupCode,
		&sop.MaterialGroup,
		&sop.ItemText,
		&sop.ProjectNumber,
		&sop.Right_Proj_Pspid,
		&sop.ProjectName,
		&sop.ProjectObject,
		&sop.ProjectCreatedBy,
		&sop.ProjectCreatedOn,
		&sop.ProjectChangedBy,
		&sop.ProjectChangedOn,
		&sop.ProjectStatusProfileCode,
		&sop.ProjectStatusProfile,
		&sop.WbsStatusProfileCode,
		&sop.ProjectResponsiblePerson,
		&sop.ProjectApplicantPerson,
		&sop.ProjectCompanyCode,
		&sop.ProjectBusinessArea,
		&sop.ProjectControllingArea,
		&sop.ProjectProfitCenterCode,
		&sop.Proj_Prctr_Org1,
		&sop.Proj_Prctr_Org1_Concat,
		&sop.Proj_Prctr_Org2,
		&sop.Proj_Prctr_Org2_Concat,
		&sop.Proj_Prctr_Org3,
		&sop.Proj_Prctr_Org3_Concat,
		&sop.ProjectCountry,
		&sop.ProjectAbacusCode,
		&sop.ProjectCurrency,
		&sop.NetworkAssignment,
		&sop.ProjectPlannedStart,
		&sop.ProjectPlannedEnd,
		&sop.ProjectPlantCode,
		&sop.ProjectPlant,
		&sop.ProjectPlanningMethodForBasicDates,
		&sop.ProjectPlanningMethodForForecastDates,
		&sop.Application,
		&sop.NetworkProfileCode,
		&sop.NetworkProfile,
		&sop.ProjectProfileCode,
		&sop.ProjectProfile,
		&sop.BudgetProfileCode,
		&sop.CurrentProjectNumber,
		&sop.ProjectStockIndicator,
		&sop.ProjectObjectClassCode,
		&sop.ProjectObjectClass,
		&sop.StatisticalIndicator,
		&sop.DeletionFlag,
		&sop.InactiveIndicator,
		&sop.FunctionalAreaCode,
		&sop.SalesOrgCode,
		&sop.DistributionChannelCode,
		&sop.DivisionCode,
		&sop.ProjectReleasedBy,
		&sop.ProjectReleasedOn,
		&sop.Proj_Firsttecoby,
		&sop.Proj_Firsttecoon,
		&sop.ProjectTecoBy,
		&sop.ProjectTecoOn,
		&sop.Proj_Tecocount,
		&sop.Proj_Firstclsdby,
		&sop.Proj_Firstclsdon,
		&sop.ProjectClsdBy,
		&sop.ProjectClsdOn,
		&sop.Proj_Clsdcount,
		&sop.ID,
		&sop.ProjectReleaseYear,
		&sop.Perprojectcount_Customers,
		&sop.Maincustomer,
		&sop.AccountGroup,
		&sop.GUID,
		&sop.IsIntercompany,
		&sop.FlagChangeOrder,
		&sop.ShipToCountryCode,
		&sop.SoldToText,
		&sop.Flags,
	)
	return &sop, err
}

type SalesOrderProjectFlags struct {
	FlagMultipleShipToPerOrder          bool `json:"Flag: Multiple Ship TO Per Order" sql:"Flag: Multiple Ship TO Per Order"`
	FlagNoShipToCountry                 bool `json:"Flag: No Ship To Country" sql:"Flag: No Ship To Country"`
	FlagShipsOutsideProjectCountry      bool `json:"Flag: Ships Outside Project Country" sql:"Flag: Ships Outside Project Country"`
	FlagVendorUsedByOneCustomer         bool `json:"Flag: Vendor Used By One Customer" sql:"Flag: Vendor Used By One Customer"`
	FlagCustomerOnProjectWithoutVendors bool `json:"Flag: Customer On Project Without Vendors" sql:"Flag: Customer On Project Without Vendors"`
}

func (sopf *SalesOrderProjectFlags) GetFlags() ICMEntity {
	return sopf
}

func (sopf *SalesOrderProjectFlags) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return json.Unmarshal([]byte(v), sopf)
	case []byte:
		return json.Unmarshal(v, sopf)
	default:
		return errors.New("invalid sql return type for Vendor Flags")
	}
}
