package orm

import (
	"time"
)

type Vendor struct {
	// Identifiers
	ID             string    `json:"ICM_VENDOR_ID" sql:"ICM_VENDOR_ID"`
	Database       string    `json:"Database" sql:"Database"`
	Code           string    `json:"Vendor Code" sql:"Vendor Code"`
	Key            string    `json:"CMT Vendor Key" sql:"CMT Vendor Key"`
	Name1          string    `json:"Name 1" sql:"Name 1"`
	Name2          string    `json:"Vendor Name 2" sql:"Vendor Name 2"`
	Name3          string    `json:"Vendor Name 3" sql:"Vendor Name 3"`
	Name4          string    `json:"Vendor Name 4" sql:"Vendor Name 4"`
	GUIDCommonName string    `json:"Vendor GUID Common Name" sql:"Vendor GUID Common Name"`
	GISLegalName   string    `json:"GIS_PRLEGALNAMELOC" sql:"GIS_PRLEGALNAMELOC"`
	CustomerID     string    `json:"Customer ID" sql:"Customer ID"`
	CreatedBy      string    `json:"Vendor Created By" sql:"Vendor Created By"`
	CreationDate   time.Time `json:"Vendor Creation Date" sql:"Vendor Creation Date"`

	// Transactional summary
	AmountUSDLastMonth    float64 `json:"Amount USD Last Month" sql:"Amount USD Last Month"`
	CountLastMonth        int64   `json:"Count Last Month" sql:"Count Last Month"`
	AmountUSDCurrentYear  float64 `json:"Amount USD Current Year" sql:"Amount USD Current Year"`
	CountCurrentYear      int64   `json:"Count Current Year" sql:"Count Current Year"`
	AmountUSDCurrentYear1 float64 `json:"Amount USD Current Year -1" sql:"Amount USD Current Year -1"`
	CountCurrentYear1     int64   `json:"Count Current Year -1" sql:"Count Current Year -1"`
	AmountUSDCurrentYear2 float64 `json:"Amount USD Current Year -2" sql:"Amount USD Current Year -2"`
	CountCurrentYear2     int64   `json:"Count Current Year -2" sql:"Count Current Year -2"`

	// blocks
	BlockPayment          string `json:"Vendor Payment Block" sql:"Vendor Payment Block"`
	BlockPosting          string `json:"Central posting block" sql:"Central posting block"`
	BlockPurchase         string `json:"Central Purchase Block" sql:"Central Purchase Block"`
	DescriptiveBlockPurch string `json:"PurchBlockDescriptive" sql:"PurchBlockDescriptive"`
	DeletionFlagCentral   string `json:"Vendor Deletion Flag (Central)" sql:"Vendor Deletion Flag (Central)"`
	BlockFunction         string `json:"Vendor Block Funciton" sql:"Vendor Block Funciton"`

	// BA/Division assignments
	PrimaryBU            string  `json:"Primary BU" sql:"Primary BU"`
	PrimaryBuInvoicesUsd float64 `json:"Primary BU Invoices USD" sql:"Primary BU Invoices USD"`
	PrimaryDivision      string  `json:"Primary Division" sql:"Primary Division"`
	CompanyCodeID        string  `json:"Company Code ID" sql:"Company Code ID"`
	Organization         string  `json:"Organization" sql:"Organization"`
	Plant                string  `json:"Vendor Plant" sql:"Vendor Plant"`

	// Other Assignments
	AccountGroup           string `json:"Vendor Account Group" sql:"Vendor Account Group"`
	AccountGroupCode       string `json:"Vendor Account Group Code" sql:"Vendor Account Group Code"`
	GUIDIndustryUsageLevel string `json:"Vendor GUID Industry Usage Level" sql:"Vendor GUID Industry Usage Level"`
	GISIndustry            string `json:"GIS_PRDIUNAME" sql:"GIS_PRDIUNAME"`
	IndustryCode           string `json:"Vendor Industry" sql:"Vendor Industry"`
	GUIDChannelCode        string `json:"Vendor GUID Channel Code" sql:"Vendor GUID Channel Code"`
	GUIDChannelName        string `json:"Vendor GUID Channel Name" sql:"Vendor GUID Channel Name"`

	// Alternate Payee Details
	AlternatePayeeAllowed          string `json:"Vendor Alternative Payee Allowed" sql:"Vendor Alternative Payee Allowed"`
	AlternatePayeeCode             string `json:"Vendor Alternate Payee Code" sql:"Vendor Alternate Payee Code"`
	AlternatePayeeName             string `json:"Vendor Alternate Payee Name" sql:"Vendor Alternate Payee Name"`
	AlternatePayee                 string `json:"Vendor Alternate Payee" sql:"Vendor Alternate Payee"`
	AlternatePayeeAccountGroup     string `json:"Vendor Alternate Payee Account Group" sql:"Vendor Alternate Payee Account Group"`
	AlternatePayeeAccountGroupCode string `json:"Vendor Alternate Payee Account Group Code" sql:"Vendor Alternate Payee Account Group Code"`
	AlternatePayeeCountry          string `json:"Vendor Alternate Payee Country" sql:"Vendor Alternate Payee Country"`
	AlternatePayeeCountryCode      string `json:"Vendor Alternate Payee Country Code" sql:"Vendor Alternate Payee Country Code"`
	AlternatePayeeGUID             string `json:"Vendor Alternate Payee GUID" sql:"Vendor Alternate Payee GUID"`
	AlternatePayeeGUIDOld          string `json:"Vendor Alternate Payee Old GUID" sql:"Vendor Alternate Payee Old GUID"`

	// Address details
	POBox           string `json:"Vendor PO Box" sql:"Vendor PO Box"`
	PostalCodePOBox string `json:"Vendor PO Box Postal Code" sql:"Vendor PO Box Postal Code"`
	StreetAddress   string `json:"Address" sql:"Address"`
	City            string `json:"City" sql:"City"`
	District        string `json:"Vendor City" sql:"Vendor City"`
	Region          string `json:"Region" sql:"Region"`
	Country         string `json:"Vendor Country" sql:"Vendor Country"`
	CountryCode     string `json:"Vendor Country Code" sql:"Vendor Country Code"`
	PostalCode      string `json:"Vendor Postal Code" sql:"Vendor Postal Code"`
	GUIDCity        string `json:"Vendor GUID City" sql:"Vendor GUID City"`
	GUIDCountry     string `json:"Vendor GUID Country" sql:"Vendor GUID Country"`

	// Contact Details
	FaxNumber  string `json:"Vendor Fax Number" sql:"Vendor Fax Number"`
	Teletex    string `json:"Vendor Teletex" sql:"Vendor Teletex"`
	Telebox    string `json:"Vendor Telebox" sql:"Vendor Telebox"`
	Telex      string `json:"Vendor Telex" sql:"Vendor Telex"`
	Telephone1 string `json:"Vendor Telephone 1" sql:"Vendor Telephone 1"`
	Telephone2 string `json:"Vendor Telephone 2" sql:"Vendor Telephone 2"`
	Url        string `json:"Vendor URL" sql:"Vendor URL"`

	// Tax information
	VAT                string `json:"Vendor VAT" sql:"Vendor VAT"`
	TaxCode1           string `json:"Vendor Tax Code 1" sql:"Vendor Tax Code 1"`
	TaxCode2           string `json:"Vendor Tax Code 2" sql:"Vendor Tax Code 2"`
	TaxCode3           string `json:"Vendor Tax Code 3" sql:"Vendor Tax Code 3"`
	TaxCode4           string `json:"Vendor Tax Code 4" sql:"Vendor Tax Code 4"`
	TaxJurisdiction    string `json:"Vendor Tax Jurisdiction" sql:"Vendor Tax Jurisdiction"`
	TaxNumberType      string `json:"Vendor Tax Number Type" sql:"Vendor Tax Number Type"`
	TaxSplit           string `json:"Vendor Tax Split" sql:"Vendor Tax Split"`
	TaxBase            string `json:"Vendor Tax Base" sql:"Vendor Tax Base"`
	TransportationZone string `json:"Vendor Transporation Zone" sql:"Vendor Transporation Zone"`

	// GIS Data
	GUID         string `json:"Vendor GUID" sql:"Vendor GUID"`
	GUIDOld      string `json:"Vendor Old GUID" sql:"Vendor Old GUID"`
	GUIDDomestic string `json:"Vendor Domestic GUID" sql:"Vendor Domestic GUID"`
	GUIDParent   string `json:"GIS_PRPARENTGUID" sql:"GIS_PRPARENTGUID"`
	DUNS         string `json:"Vendor DUNS" sql:"Vendor DUNS"`
	DUNSDomestic string `json:"Vendor Domestic DUNS" sql:"Vendor Domestic DUNS"`
	DUNSParent   string `json:"GIS_PRPARENTDUNS" sql:"GIS_PRPARENTDUNS"`

	// Flags
	Flags VendorFlags `json:"FLAGS" sql:"Flags"`

	// Flag Counts and Match Indicators
	FlagCount                 int64 `json:"Flags Triggered" sql:"Flags Triggered"`
	MatchDataDeviation        bool  `json:"Data Deviation Match" sql:"Data Deviation Match"`
	FlagCountDataDeviation    int64 `json:"Data Deviation Flag Count" sql:"Data Deviation Flag Count"`
	MatchDormantVendor        bool  `json:"Dormant Vendor Match" sql:"Dormant Vendor Match"`
	FlagCountDormant          int64 `json:"Dormant Flag Count" sql:"Dormant Flag Count"`
	MatchFictitiousVendor     bool  `json:"Fictitious Vendor Match" sql:"Fictitious Vendor Match"`
	FlagCountFictitiousVendor int64 `json:"Fictitious Vendor Flag Count" sql:"Fictitious Vendor Flag Count"`
	MatchUniversal            bool  `json:"Universal Match" sql:"Universal Match"`
	FlagCountUniversal        int64 `json:"Universal Flag Count" sql:"Universal Flag Count"`

	// Unassigned fields
	CreditInformationNumber               string    `json:"Vendor Credit Information Number" sql:"Vendor Credit Information Number"`
	IsIntercompany                        bool      `json:"Is Intercompany" sql:"Is Intercompany"`
	Language                              string    `json:"Vendor Language" sql:"Vendor Language"`
	LastVendorExternalReview              time.Time `json:"Last Vendor External Review" sql:"Last Vendor External Review"`
	ListMissingFields                     string    `json:"List Missing Fields" sql:"List Missing Fields"`
	OneTimeVendorIndicator                string    `json:"One-time Vendor Indicator" sql:"One-time Vendor Indicator"`
	PrCode6                               string    `json:"Vendor PR Code 6" sql:"Vendor PR Code 6"`
	RoundAmountRatio                      float64   `json:"Round Amount Ratio" sql:"Round Amount Ratio"`
	ShipmentStatisticsGroup               string    `json:"Vendor Shipment Statistics Group" sql:"Vendor Shipment Statistics Group"`
	SortString                            string    `json:"Vendor Sort String" sql:"Vendor Sort String"`
	TotalRoundAmountTransactions          int64     `json:"Total Round Amount Transactions" sql:"Total Round Amount Transactions"`
	TotalTransactionsTestedForRoundAmount int64     `json:"Total Transactions Tested for Round Amount" sql:"Total Transactions Tested for Round Amount"`
	TradingPartnerCompanyID               string    `json:"Company ID of trading partner" sql:"Company ID of trading partner"`
}

func (v Vendor) GetFlags() ICMEntity {
	return v.Flags
}

type VendorFlags struct {
	Flag2WayMatchPOsOnly                       bool `json:"Flag: 2WayMatchPOsOnly" sql:"Flag: 2WayMatchPOsOnly"`
	FlagAllVendorInvoicesWithoutPO             bool `json:"Flag: All Vendor Invoices Without PO" sql:"Flag: All Vendor Invoices Without PO"`
	FlagBankAccountAddedToDormantVendor        bool `json:"Flag: Bank Account Added to Dormant Vendor" sql:"Flag: Bank Account Added to Dormant Vendor"`
	FlagConsInv                                bool `json:"Flag: ConsInvVendor" sql:"Flag: ConsInvVendor"`
	FlagCountryWithHighEnforcementActivity     bool `json:"Flag: Country with High Enforcement Activity" sql:"Flag: Country with High Enforcement Activity"`
	FlagCountryWithSL1CaseHistory              bool `json:"Flag: Country with SL1 Case History" sql:"Flag: Country with SL1 Case History"`
	FlagDormantVendor                          bool `json:"Flag: Dormant Vendor" sql:"Flag: Dormant Vendor"`
	FlagFIInvoice                              bool `json:"Flag: FI Invoice" sql:"Flag: FI Invoice"`
	FlagHadActivityLastMonth                   bool `json:"Flag: Had Activity Last Month" sql:"Flag: Had Activity Last Month"`
	FlagHasSplitInvoice                        bool `json:"Flag: Has Split Invoice" sql:"Flag: Has Split Invoice"`
	FlagIncreasingInvoiceValue                 bool `json:"Flag: Increasing Invoice Value" sql:"Flag: Increasing Invoice Value"`
	FlagMissingOrNonStandardDuns               bool `json:"Flag: Missing or Non-Standard DUNS" sql:"Flag: Missing or Non-Standard DUNS"`
	FlagMissingOrNonStandardGuid               bool `json:"Flag: Missing or Non-Standard GUID" sql:"Flag: Missing or Non-Standard GUID"`
	FlagMissingVendorMasterFields              bool `json:"Flag: Missing Vendor Master Fields" sql:"Flag: Missing Vendor Master Fields"`
	FlagNoReceiptNeeded                        bool `json:"Flag: NoReceiptNeeded" sql:"Flag: NoReceiptNeeded"`
	FlagOnlyExternalDomainsAreKnownFreeDomains bool `json:"Flag: Only External Domains are Known Free Domains" sql:"Flag: Only External Domains are Known Free Domains"`
	FlagPOBoxAddress                           bool `json:"Flag: PO Box Address" sql:"Flag: PO Box Address"`
	FlagPossibleGD43                           bool `json:"Flag: Possible GD43" sql:"Flag: Possible GD43"`
	FlagPossibleSOE                            bool `json:"Flag: Possible SOE" sql:"Flag: Possible SOE"`
	FlagRecentBankAccountChange                bool `json:"Flag: Recent Bank Account Change" sql:"Flag: Recent Bank Account Change"`
	FlagSCP                                    bool `json:"Flag: SCP" sql:"Flag: SCP"`
	FlagServiceVendor                          bool `json:"Flag: ServiceVendor" sql:"Flag: ServiceVendor"`
	FlagSwiftFirstTransaction                  bool `json:"Flag: SwiftFirstTransaction" sql:"Flag: SwiftFirstTransaction"`
	FlagTop5BenfordDeviation                   bool `json:"Flag: Top5BenfordDeviation" sql:"Flag: Top5BenfordDeviation"`
	FlagTop5ZScorevendors                      bool `json:"Flag: Top5 Z-ScoreVendors" sql:"Flag: Top5 Z-ScoreVendors"`
	FlagUnblockedPurchBlock                    bool `json:"Flag: UnblockedPurchBlock" sql:"Flag: UnblockedPurchBlock"`
	FlagVendorFormerEmployee                   bool `json:"Flag: Vendor Former Emploee" sql:"Flag: Vendor Former Emploee"`
	FlagVendorPOSingleApproval                 bool `json:"Flag: VendorPOSingleApproval" sql:"Flag: VendorPOSingleApproval"`
	FlagVendorPOsNotApproved                   bool `json:"Flag: VendorPOsNotApproved" sql:"Flag: VendorPOsNotApproved"`
	FlagVendorWithoutGUIDActive                bool `json:"Flag: Vendor Without GUID Active" sql:"Flag: Vendor Without GUID, Active"`
	FlagVendorWithoutGUIDBACode                bool `json:"Flag: Vendor Without GUID BA Code" sql:"Flag: Vendor Without GUID, BA Code"`
}

func (vf VendorFlags) GetFlags() ICMEntity {
	return vf
}
