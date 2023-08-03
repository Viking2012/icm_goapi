package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type Project struct {
	Database                              string        `json:"Database" sql:"Database"`
	ProjectId                             string        `json:"Project ID" sql:"Project ID"`
	ProjectNumber                         string        `json:"Project Number" sql:"Project Number"`
	ProjectName                           string        `json:"Project Name" sql:"Project Name"`
	ProjectObject                         string        `json:"Project Object" sql:"Project Object"`
	ProjectCreatedBy                      string        `json:"Project Created By" sql:"Project Created By"`
	ProjectCreatedOn                      NullTime      `json:"Project Created On" sql:"Project Created On"`
	ProjectChangedBy                      string        `json:"Project Changed By" sql:"Project Changed By"`
	ProjectChangedOn                      NullTime      `json:"Project Changed On" sql:"Project Changed On"`
	ProjectStatusProfileCode              string        `json:"Project Status Profile Code" sql:"Project Status Profile Code"`
	ProjectStatusProfile                  string        `json:"Project Status Profile" sql:"Project Status Profile"`
	WbsStatusProfileCode                  string        `json:"WBS Status Profile Code" sql:"WBS Status Profile Code"`
	ProjectResponsiblePerson              string        `json:"Project Responsible Person" sql:"Project Responsible Person"`
	ProjectApplicantPerson                string        `json:"Project Applicant Person" sql:"Project Applicant Person"`
	ProjectCompanyCode                    string        `json:"Project Company Code" sql:"Project Company Code"`
	ProjectBusinessArea                   string        `json:"Project Business Area" sql:"Project Business Area"`
	ProjectControllingArea                string        `json:"Project Controlling Area" sql:"Project Controlling Area"`
	ProjectProfitCenterCode               string        `json:"Project Profit Center Code" sql:"Project Profit Center Code"`
	Proj_Prctr_Org1                       string        `json:"PROJ_PRCTR_ORG1" sql:"PROJ_PRCTR_ORG1"`
	Proj_Prctr_Org1_Concat                string        `json:"PROJ_PRCTR_ORG1_Concat" sql:"PROJ_PRCTR_ORG1_Concat"`
	Proj_Prctr_Org2                       string        `json:"PROJ_PRCTR_ORG2" sql:"PROJ_PRCTR_ORG2"`
	Proj_Prctr_Org2_Concat                string        `json:"PROJ_PRCTR_ORG2_Concat" sql:"PROJ_PRCTR_ORG2_Concat"`
	Proj_Prctr_Org3                       string        `json:"PROJ_PRCTR_ORG3" sql:"PROJ_PRCTR_ORG3"`
	Proj_Prctr_Org3_Concat                string        `json:"PROJ_PRCTR_ORG3_Concat" sql:"PROJ_PRCTR_ORG3_Concat"`
	ProjectCountry                        string        `json:"Project Country" sql:"Project Country"`
	ProjectAbacusCode                     string        `json:"Project Abacus Code" sql:"Project Abacus Code"`
	ProjectCurrency                       string        `json:"Project Currency" sql:"Project Currency"`
	NetworkAssignment                     string        `json:"Network Assignment" sql:"Network Assignment"`
	ProjectPlannedStart                   NullTime      `json:"Project Planned Start" sql:"Project Planned Start"`
	ProjectPlannedEnd                     NullTime      `json:"Project Planned End" sql:"Project Planned End"`
	ProjectPlantCode                      string        `json:"Project Plant Code" sql:"Project Plant Code"`
	ProjectPlant                          string        `json:"Project Plant" sql:"Project Plant"`
	ProjectPlanningMethodForBasicDates    string        `json:"Project Planning Method for Basic Dates" sql:"Project Planning Method for Basic Dates"`
	ProjectPlanningMethodForForecastDates string        `json:"Project Planning Method for Forecast Dates" sql:"Project Planning Method for Forecast Dates"`
	Application                           string        `json:"Application" sql:"Application"`
	NetworkProfileCode                    string        `json:"Network Profile Code" sql:"Network Profile Code"`
	NetworkProfile                        string        `json:"Network Profile" sql:"Network Profile"`
	ProjectProfileCode                    string        `json:"Project Profile Code" sql:"Project Profile Code"`
	ProjectProfile                        string        `json:"Project Profile" sql:"Project Profile"`
	BudgetProfileCode                     string        `json:"Budget Profile Code" sql:"Budget Profile Code"`
	CurrentProjectNumber                  string        `json:"Current Project Number" sql:"Current Project Number"`
	ProjectStockIndicator                 string        `json:"Project Stock Indicator" sql:"Project Stock Indicator"`
	ProjectObjectClassCode                string        `json:"Project Object Class Code" sql:"Project Object Class Code"`
	ProjectObjectClass                    string        `json:"Project Object Class" sql:"Project Object Class"`
	StatisticalIndicator                  string        `json:"Statistical Indicator" sql:"Statistical Indicator"`
	DeletionFlag                          string        `json:"Deletion Indicator" sql:"Deletion Indicator"`
	InactiveIndicator                     string        `json:"Inactive Indicator" sql:"Inactive Indicator"`
	FunctionalAreaCode                    string        `json:"Functional Area Code" sql:"Functional Area Code"`
	SalesOrgCode                          string        `json:"Sales Org Code" sql:"Sales Org Code"`
	DistributionChannelCode               string        `json:"Distribution Channel Code" sql:"Distribution Channel Code"`
	DivisionCode                          string        `json:"Division Code" sql:"Division Code"`
	ProjectReleasedBy                     string        `json:"Project Released By" sql:"Project Released By"`
	ProjectReleasedOn                     NullTime      `json:"Project Released On" sql:"Project Released On"`
	Proj_Firsttecoby                      string        `json:"PROJ_FirstTecoBy" sql:"PROJ_FirstTecoBy"`
	Proj_Firsttecoon                      NullTime      `json:"PROJ_FirstTecoOn" sql:"PROJ_FirstTecoOn"`
	ProjectTecoBy                         string        `json:"Project TECO By" sql:"Project TECO By"`
	ProjectTecoOn                         NullTime      `json:"Project TECO On" sql:"Project TECO On"`
	Proj_Tecocount                        int64         `json:"PROJ_TecoCount" sql:"PROJ_TecoCount"`
	Proj_Firstclsdby                      string        `json:"PROJ_FirstClsdBy" sql:"PROJ_FirstClsdBy"`
	Proj_Firstclsdon                      NullTime      `json:"PROJ_FirstClsdOn" sql:"PROJ_FirstClsdOn"`
	ProjectClsdBy                         string        `json:"Project CLSD By" sql:"Project CLSD By"`
	ProjectClsdOn                         NullTime      `json:"Project CLSD On" sql:"Project CLSD On"`
	Proj_Clsdcount                        int64         `json:"PROJ_ClsdCount" sql:"PROJ_ClsdCount"`
	ID                                    string        `json:"ICM_ID" sql:"ICM_ID"`
	Project                               string        `json:"Project" sql:"Project"`
	ProjectReleaseYear                    string        `json:"Project Release Year" sql:"Project Release Year"`
	FirstPoYear                           string        `json:"First PO Year" sql:"First PO Year"`
	ProjectPosUsd                         float64       `json:"Project POs USD" sql:"Project POs USD"`
	FirstSalesOrderYear                   string        `json:"First Sales Order Year" sql:"First Sales Order Year"`
	ProjectSalesOrdersUsd                 float64       `json:"Project Sales Orders USD" sql:"Project Sales Orders USD"`
	Flags                                 *ProjectFlags `json:"flags" sql:"FLAGS"`
	TltProjSalesOrdersUsdString           string        `json:"TLT-PROJ Sales Orders USD STRING" sql:"TLT-PROJ Sales Orders USD STRING"`
	TltProjPurchOrdersUsdString           string        `json:"TLT- PROJ Purch Orders USD STRING" sql:"TLT- PROJ Purch Orders USD STRING"`
	ProjectFirstActivityYear              string        `json:"Project First Activity Year" sql:"Project First Activity Year"`
	ShipToProjectId                       string        `json:"SHIP TO Project ID" sql:"SHIP TO Project ID"`
	ProjectShipToLocations                string        `json:"Project Ship To Locations" sql:"Project Ship To Locations"`
}

func (p *Project) IsICMEntity() bool { return true }

func ProjectFromRow(rows *sql.Rows) (ICMEntity, error) {
	var p Project
	err := rows.Scan(
		&p.Database,
		&p.ProjectId,
		&p.ProjectNumber,
		&p.ProjectName,
		&p.ProjectObject,
		&p.ProjectCreatedBy,
		&p.ProjectCreatedOn,
		&p.ProjectChangedBy,
		&p.ProjectChangedOn,
		&p.ProjectStatusProfileCode,
		&p.ProjectStatusProfile,
		&p.WbsStatusProfileCode,
		&p.ProjectResponsiblePerson,
		&p.ProjectApplicantPerson,
		&p.ProjectCompanyCode,
		&p.ProjectBusinessArea,
		&p.ProjectControllingArea,
		&p.ProjectProfitCenterCode,
		&p.Proj_Prctr_Org1,
		&p.Proj_Prctr_Org1_Concat,
		&p.Proj_Prctr_Org2,
		&p.Proj_Prctr_Org2_Concat,
		&p.Proj_Prctr_Org3,
		&p.Proj_Prctr_Org3_Concat,
		&p.ProjectCountry,
		&p.ProjectAbacusCode,
		&p.ProjectCurrency,
		&p.NetworkAssignment,
		&p.ProjectPlannedStart,
		&p.ProjectPlannedEnd,
		&p.ProjectPlantCode,
		&p.ProjectPlant,
		&p.ProjectPlanningMethodForBasicDates,
		&p.ProjectPlanningMethodForForecastDates,
		&p.Application,
		&p.NetworkProfileCode,
		&p.NetworkProfile,
		&p.ProjectProfileCode,
		&p.ProjectProfile,
		&p.BudgetProfileCode,
		&p.CurrentProjectNumber,
		&p.ProjectStockIndicator,
		&p.ProjectObjectClassCode,
		&p.ProjectObjectClass,
		&p.StatisticalIndicator,
		&p.DeletionFlag,
		&p.InactiveIndicator,
		&p.FunctionalAreaCode,
		&p.SalesOrgCode,
		&p.DistributionChannelCode,
		&p.DivisionCode,
		&p.ProjectReleasedBy,
		&p.ProjectReleasedOn,
		&p.Proj_Firsttecoby,
		&p.Proj_Firsttecoon,
		&p.ProjectTecoBy,
		&p.ProjectTecoOn,
		&p.Proj_Tecocount,
		&p.Proj_Firstclsdby,
		&p.Proj_Firstclsdon,
		&p.ProjectClsdBy,
		&p.ProjectClsdOn,
		&p.Proj_Clsdcount,
		&p.ID,
		&p.Project,
		&p.ProjectReleaseYear,
		&p.FirstPoYear,
		&p.ProjectPosUsd,
		&p.FirstSalesOrderYear,
		&p.ProjectSalesOrdersUsd,
		&p.Flags,
		&p.TltProjSalesOrdersUsdString,
		&p.TltProjPurchOrdersUsdString,
		&p.ProjectFirstActivityYear,
		&p.ShipToProjectId,
		&p.ProjectShipToLocations,
	)
	return &p, err
}

type ProjectFlags struct {
	FlagProjectWithoutActivity             bool `json:"Flag: Project Without Activity" sql:"Flag: Project Without Activity"`
	FlagProjectWithChangeOrder             bool `json:"Flag: Project With Change Order" sql:"Flag: Project With Change Order"`
	FlagHasShipmentsOutsideProjectCountry  bool `json:"Flag: Has Shipments Outside Project Country" sql:"Flag: Has Shipments Outside Project Country"`
	FlagAllShipmentsOutsideProjectCountry  bool `json:"Flag: All Shipments Outside Project Country" sql:"Flag: All Shipments Outside Project Country"`
	FlagProjectWithoutCustomer             bool `json:"Flag: Project Without Customer" sql:"Flag: Project Without Customer"`
	FlagVendorUsedByOneCustomer            bool `json:"Flag: Vendor Used By One Customer" sql:"Flag: Vendor Used By One Customer"`
	FlagProjectWithoutVendors              bool `json:"Flag: Project Without Vendors" sql:"Flag: Project Without Vendors"`
	FlagVendorRelatedToCustomerOnProject   bool `json:"Flag: Vendor Related To Customer On Project" sql:"Flag: Vendor Related To Customer On Project"`
	FlagCloseDatesVendorCreationAndProject bool `json:"Flag: Close Dates Vendor Creation and Project" sql:"Flag: Close Dates Vendor Creation and Project"`
	FlagSeveralExternalCustomersPerProj    bool `json:"Flag: Several External Customers Per Proj" sql:"Flag: Several External Customers Per Proj"`
	FlagPoCreatedAfterProjectClosure       bool `json:"Flag: PO Created After Project Closure" sql:"Flag: PO Created After Project Closure"`
	FlagPoCreatedBeforeProjectRelease      bool `json:"Flag: PO Created Before Project Release" sql:"Flag: PO Created Before Project Release"`
	FlagPrj_Not_Break_Even                 bool `json:"Flag:Prj_Not_Break_Even" sql:"Flag:Prj_Not_Break_Even"`
	FlagHasMultiSourcedMaterials           bool `json:"Flag: Has Multi-Sourced Materials" sql:"Flag: Has Multi-Sourced Materials"`
	FlagMultiprojectLocation               bool `json:"Flag: MultiProject Location" sql:"Flag: MultiProject Location"`
	FlagMultisystemproject                 bool `json:"Flag: MultiSystemProject" sql:"Flag: MultiSystemProject"`
	ExternalCustomerProject                bool `json:"External Customer Project" sql:"External Customer Project"`
	FlagIsSalesChannel                     bool `json:"Flag: Is Sales Channel" sql:"Flag: Is Sales Channel"`
}

func (pf *ProjectFlags) IsICMEntity() bool { return true }

func (pf *ProjectFlags) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return json.Unmarshal([]byte(v), pf)
	case []byte:
		return json.Unmarshal(v, pf)
	default:
		return errors.New("invalid sql return type for Vendor Flags")
	}
}
