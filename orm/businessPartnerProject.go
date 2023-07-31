package orm

import "time"

type BusinessPartner struct {
	FirstPoYear                           string               `json:"First PO Year" sql:"First PO Year"`
	Database                              string               `json:"Database" sql:"Database"`
	ProjectNumber                         string               `json:"Project Number" sql:"Project Number"`
	BusinesPartner                        string               `json:"Busines Partner" sql:"Busines Partner"`
	Firstpodatewbs                        time.Time            `json:"FirstPODateWBS" sql:"FirstPODateWBS"`
	Firstpodate_No_Wbs                    time.Time            `json:"FirstPODate_NO_WBS" sql:"FirstPODate_NO_WBS"`
	Wbs_Vendoronly                        bool                 `json:"WBS_VendorOnly" sql:"WBS_VendorOnly"`
	Mixwbsandnone                         bool                 `json:"MixWBSandNone" sql:"MixWBSandNone"`
	ProjectName                           string               `json:"Project Name" sql:"Project Name"`
	Name                                  string               `json:"Name" sql:"Name"`
	Value                                 float64              `json:"Value" sql:"Value"`
	BpType                                string               `json:"BP Type" sql:"BP Type"`
	ProjectChangedOn                      time.Time            `json:"Project Changed On" sql:"Project Changed On"`
	ProjectCountry                        string               `json:"Project Country" sql:"Project Country"`
	ProjectResponsiblePerson              string               `json:"Project Responsible Person" sql:"Project Responsible Person"`
	BpName                                string               `json:"BP Name" sql:"BP Name"`
	BpCountry                             string               `json:"BP Country" sql:"BP Country"`
	BpCity                                string               `json:"BP City" sql:"BP City"`
	BpStreet                              string               `json:"BP Street" sql:"BP Street"`
	BpCreatedOn                           time.Time            `json:"BP Created On" sql:"BP Created On"`
	BpCreator                             string               `json:"BP Creator" sql:"BP Creator"`
	AccountGroup                          string               `json:"Account Group" sql:"Account Group"`
	BpOpposingSide                        string               `json:"BP Opposing Side" sql:"BP Opposing Side"`
	GUID                                  string               `json:"Vendor GUID" sql:"Vendor GUID"`
	IsIntercompany                        bool                 `json:"Is Intercompany" sql:"Is Intercompany"`
	FirstSalesOrderYear                   string               `json:"First Sales Order Year" sql:"First Sales Order Year"`
	Maincustomer                          bool                 `json:"MainCustomer" sql:"MainCustomer"`
	ProjectReleasedOn                     time.Time            `json:"Project Released On" sql:"Project Released On"`
	ProjectReleaseYear                    string               `json:"Project Release Year" sql:"Project Release Year"`
	BlockOrder                            string               `json:"Order Block" sql:"Order Block"`
	BlockDelivery                         string               `json:"Delivery Block" sql:"Delivery Block"`
	DeletionFlag                          string               `json:"Deletion Indicator" sql:"Deletion Indicator"`
	GUIDChannelName                       string               `json:"Channel" sql:"Channel"`
	ProjectFirstActivityYear              string               `json:"Project First Activity Year" sql:"Project First Activity Year"`
	FlagCustomerWithChangeOrder           bool                 `json:"Flag: Customer With Change Order" sql:"Flag: Customer With Change Order"`
	FlagAllShipmentsOutsideProjectCountry bool                 `json:"Flag: All Shipments Outside Project Country" sql:"Flag: All Shipments Outside Project Country"`
	FlagHasShipmentsOutsideProjectCountry bool                 `json:"Flag: Has Shipments Outside Project Country" sql:"Flag: Has Shipments Outside Project Country"`
	FlagVendorUsedByOneCustomer           bool                 `json:"Flag: Vendor Used By One Customer" sql:"Flag: Vendor Used By One Customer"`
	FlagCustomerOnProjectWithoutVendors   bool                 `json:"Flag: Customer On Project Without Vendors" sql:"Flag: Customer On Project Without Vendors"`
	FlagVendorOnProjectWithoutCustomer    bool                 `json:"Flag: Vendor On Project Without Customer" sql:"Flag: Vendor On Project Without Customer"`
	Right_Gis_Guid                        string               `json:"Right_GIS_GUID" sql:"Right_GIS_GUID"`
	LinkOn                                string               `json:"Link On" sql:"Link On"`
	SourceOfProbability                   string               `json:"Source Of Probability" sql:"Source Of Probability"`
	MatchProbability                      string               `json:"Match Probability" sql:"Match Probability"`
	BpProjectCount                        int32                `json:"BP Project Count" sql:"BP Project Count"`
	BpNegativeProjects                    int32                `json:"BP Negative Projects" sql:"BP Negative Projects"`
	NegativeProjectRatio                  float64              `json:"Negative Project Ratio" sql:"Negative Project Ratio"`
	TltNoBreakEvenBps                     string               `json:"TLT -  No Break Even BPs" sql:"TLT -  No Break Even BPs"`
	ShipToProjectId                       string               `json:"SHIP TO Project ID" sql:"SHIP TO Project ID"`
	ProjectShipToLocations                string               `json:"Project Ship To Locations" sql:"Project Ship To Locations"`
	HyperVendorNote1Day                   string               `json:"Hyper Vendor Note 1 Day" sql:"Hyper Vendor Note 1 Day"`
	HyperVendorNote                       string               `json:"Hyper Vendor Note" sql:"Hyper Vendor Note"`
	Right_BpType                          string               `json:"Right_BP Type" sql:"Right_BP Type"`
	Flags                                 BusinessPartnerFlags `json:"flags" sql:"FLAGS"`
}

func (bp BusinessPartner) GetFlags() ICMEntity {
	return bp.Flags
}

type BusinessPartnerFlags struct {
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
	FlagHyperActiveVendor1Day              bool `json:"Flag - Hyper Active Vendor 1 day" sql:"Flag - Hyper Active Vendor 1 day"`
	FlagHyperActiveVendor                  bool `json:"Flag - Hyper Active Vendor" sql:"Flag - Hyper Active Vendor"`
	FlagIsSalesChannel                     bool `json:"Flag: Is Sales Channel" sql:"Flag: Is Sales Channel"`
}

func (bpf BusinessPartnerFlags) GetFlags() ICMEntity {
	return bpf
}
