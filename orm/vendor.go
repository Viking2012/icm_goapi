package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type Vendor struct {
	// Identifiers
	ID             string   `json:"ICM_VENDOR_ID" sql:"ICM_VENDOR_ID"`
	Database       string   `json:"Database" sql:"Database"`
	Code           string   `json:"Vendor Code" sql:"Vendor Code"`
	Key            string   `json:"CMT Vendor Key" sql:"CMT Vendor Key"`
	Name1          string   `json:"Name 1" sql:"Name 1"`
	Name2          string   `json:"Vendor Name 2" sql:"Vendor Name 2"`
	Name3          string   `json:"Vendor Name 3" sql:"Vendor Name 3"`
	Name4          string   `json:"Vendor Name 4" sql:"Vendor Name 4"`
	GUIDCommonName string   `json:"Vendor GUID Common Name" sql:"Vendor GUID Common Name"`
	GISLegalName   string   `json:"GIS_PRLEGALNAMELOC" sql:"GIS_PRLEGALNAMELOC"`
	CustomerCode   string   `json:"Customer ID" sql:"Customer ID"`
	CreatedBy      string   `json:"Vendor Created By" sql:"Vendor Created By"`
	CreationDate   NullTime `json:"Vendor Creation Date" sql:"Vendor Creation Date"`

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
	DeletionFlag          string `json:"Vendor Deletion Flag (Central)" sql:"Vendor Deletion Flag (Central)"`
	BlockFunction         string `json:"Vendor Block Funciton" sql:"Vendor Block Funciton"`

	// BA/Division assignments
	PrimaryBU            string  `json:"Primary BU" sql:"Primary BU"`
	PrimaryBUInvoicesUsd float64 `json:"Primary BU Invoices USD" sql:"Primary BU Invoices USD"`
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
	Flags                             *VendorFlags `json:"flags" sql:"FLAGS"`
	FlagSCPRiskRating                 string       `json:"Flag: SCP Risk Rating" sql:"Flag: SCP, Risk Rating"`
	FlagSCPStatus                     string       `json:"Flag: SCP Status" sql:"Flag: SCP, Status"`
	FlagSplitInvoiceTotalInUSD        float64      `json:"Flag: Split Invoice Total in USD" sql:"Flag: Split Invoice Total in USD"`
	FlagVendorWithoutGUIDBA           string       `json:"Flag: Vendor Without GUID BA" sql:"Flag: Vendor Without GUID, BA"`
	FlagVendorWithoutGUIDBACode       string       `json:"Flag: Vendor Without GUID BA Code" sql:"Flag: Vendor Without GUID, BA Code"`
	FlagVendorWithoutGUIDCompanyCode  string       `json:"Flag: Vendor Without GUID Company Code" sql:"Flag: Vendor Without GUID, Company Code"`
	FlagVendorWithoutGUIDDivision     string       `json:"Flag: Vendor Without GUID Division" sql:"Flag: Vendor Without GUID, Division"`
	FlagVendorWithoutGUIDDivisionCode string       `json:"Flag: Vendor Without GUID Division Code" sql:"Flag: Vendor Without GUID, Division Code"`
	FlagVendorWithoutGUIDRegion       string       `json:"Flag: Vendor Without GUID Region" sql:"Flag: Vendor Without GUID, Region"`
	DescriptiveBenford                string       `json:"BenfordDescriptive" sql:"BenfordDescriptive"`
	DescriptiveSharedBank             string       `json:"SharedBankDescriptive" sql:"SharedBankDescriptive"`
	DescriptiveServiceType            string       `json:"ServiceTypeDescriptive" sql:"ServiceTypeDescriptive"`
	DescriptiveZScore                 string       `json:"Z-Score Descriptive" sql:"Z-Score Descriptive"`
	DomainListing                     string       `json:"Domain Listing" sql:"Domain Listing"`
	VendorsPerBankAccount             string       `json:"VendorsPerBankAccount" sql:"VendorsPerBankAccount"`

	// Flag Counts and Match Indicators
	FlagCount                 int64         `json:"Flags Triggered" sql:"Flags Triggered"`
	MatchDataDeviation        BoolFromFloat `json:"Data Deviation Match" sql:"Data Deviation Match"`
	FlagCountDataDeviation    int64         `json:"Data Deviation Flag Count" sql:"Data Deviation Flag Count"`
	MatchDormantVendor        BoolFromFloat `json:"Dormant Vendor Match" sql:"Dormant Vendor Match"`
	FlagCountDormant          int64         `json:"Dormant Flag Count" sql:"Dormant Flag Count"`
	MatchFictitiousVendor     BoolFromFloat `json:"Fictitious Vendor Match" sql:"Fictitious Vendor Match"`
	FlagCountFictitiousVendor int64         `json:"Fictitious Vendor Flag Count" sql:"Fictitious Vendor Flag Count"`
	MatchUniversal            BoolFromFloat `json:"Universal Match" sql:"Universal Match"`
	FlagCountUniversal        int64         `json:"Universal Flag Count" sql:"Universal Flag Count"`

	// Unassigned fields
	CreditInformationNumber               string        `json:"Vendor Credit Information Number" sql:"Vendor Credit Information Number"`
	IsIntercompany                        BoolFromFloat `json:"Is Intercompany" sql:"Is Intercompany"`
	Language                              string        `json:"Vendor Language" sql:"Vendor Language"`
	LastVendorExternalReview              NullTime      `json:"Last Vendor External Review" sql:"Last Vendor External Review"`
	ListMissingFields                     string        `json:"List Missing Fields" sql:"List Missing Fields"`
	OneTimeVendorIndicator                string        `json:"One-time Vendor Indicator" sql:"One-time Vendor Indicator"`
	PrCode6                               string        `json:"Vendor PR Code 6" sql:"Vendor PR Code 6"`
	RoundAmountRatio                      float64       `json:"Round Amount Ratio" sql:"Round Amount Ratio"`
	ShipmentStatisticsGroup               string        `json:"Vendor Shipment Statistics Group" sql:"Vendor Shipment Statistics Group"`
	SortString                            string        `json:"Vendor Sort String" sql:"Vendor Sort String"`
	TotalRoundAmountTransactions          int64         `json:"Total Round Amount Transactions" sql:"Total Round Amount Transactions"`
	TotalTransactionsTestedForRoundAmount int64         `json:"Total Transactions Tested for Round Amount" sql:"Total Transactions Tested for Round Amount"`
	TradingPartnerCompanyID               string        `json:"Company ID of trading partner" sql:"Company ID of trading partner"`
}

func (v *Vendor) IsICMEntity() bool { return true }
func (v *Vendor) GetID() string     { return v.ID }

func VendorFromRow(row *sql.Rows) (ICMEntity, error) {
	var v Vendor
	err := row.Scan(
		&v.ID,
		&v.Database,
		&v.Code,
		&v.Key,
		&v.Name1,
		&v.Name2,
		&v.Name3,
		&v.Name4,
		&v.GUIDCommonName,
		&v.GISLegalName,
		&v.CustomerCode,
		&v.CreatedBy,
		&v.CreationDate,
		&v.AmountUSDLastMonth,
		&v.CountLastMonth,
		&v.AmountUSDCurrentYear,
		&v.CountCurrentYear,
		&v.AmountUSDCurrentYear1,
		&v.CountCurrentYear1,
		&v.AmountUSDCurrentYear2,
		&v.CountCurrentYear2,
		&v.BlockPayment,
		&v.BlockPosting,
		&v.BlockPurchase,
		&v.DescriptiveBlockPurch,
		&v.DeletionFlag,
		&v.BlockFunction,
		&v.PrimaryBU,
		&v.PrimaryBUInvoicesUsd,
		&v.PrimaryDivision,
		&v.CompanyCodeID,
		&v.Organization,
		&v.Plant,
		&v.AccountGroup,
		&v.AccountGroupCode,
		&v.GUIDIndustryUsageLevel,
		&v.GISIndustry,
		&v.IndustryCode,
		&v.GUIDChannelCode,
		&v.GUIDChannelName,
		&v.AlternatePayeeAllowed,
		&v.AlternatePayeeCode,
		&v.AlternatePayeeName,
		&v.AlternatePayee,
		&v.AlternatePayeeAccountGroup,
		&v.AlternatePayeeAccountGroupCode,
		&v.AlternatePayeeCountry,
		&v.AlternatePayeeCountryCode,
		&v.AlternatePayeeGUID,
		&v.AlternatePayeeGUIDOld,
		&v.POBox,
		&v.PostalCodePOBox,
		&v.StreetAddress,
		&v.City,
		&v.District,
		&v.Region,
		&v.Country,
		&v.CountryCode,
		&v.PostalCode,
		&v.GUIDCity,
		&v.GUIDCountry,
		&v.FaxNumber,
		&v.Teletex,
		&v.Telebox,
		&v.Telex,
		&v.Telephone1,
		&v.Telephone2,
		&v.Url,
		&v.VAT,
		&v.TaxCode1,
		&v.TaxCode2,
		&v.TaxCode3,
		&v.TaxCode4,
		&v.TaxJurisdiction,
		&v.TaxNumberType,
		&v.TaxSplit,
		&v.TaxBase,
		&v.TransportationZone,
		&v.GUID,
		&v.GUIDOld,
		&v.GUIDDomestic,
		&v.GUIDParent,
		&v.DUNS,
		&v.DUNSDomestic,
		&v.DUNSParent,
		&v.Flags,
		&v.FlagSCPRiskRating,
		&v.FlagSCPStatus,
		&v.FlagSplitInvoiceTotalInUSD,
		&v.FlagVendorWithoutGUIDBA,
		&v.FlagVendorWithoutGUIDBACode,
		&v.FlagVendorWithoutGUIDCompanyCode,
		&v.FlagVendorWithoutGUIDDivision,
		&v.FlagVendorWithoutGUIDDivisionCode,
		&v.FlagVendorWithoutGUIDRegion,
		&v.DescriptiveBenford,
		&v.DescriptiveSharedBank,
		&v.DescriptiveServiceType,
		&v.DescriptiveZScore,
		&v.DomainListing,
		&v.VendorsPerBankAccount,
		&v.FlagCount,
		&v.MatchDataDeviation,
		&v.FlagCountDataDeviation,
		&v.MatchDormantVendor,
		&v.FlagCountDormant,
		&v.MatchFictitiousVendor,
		&v.FlagCountFictitiousVendor,
		&v.MatchUniversal,
		&v.FlagCountUniversal,
		&v.CreditInformationNumber,
		&v.IsIntercompany,
		&v.Language,
		&v.LastVendorExternalReview,
		&v.ListMissingFields,
		&v.OneTimeVendorIndicator,
		&v.PrCode6,
		&v.RoundAmountRatio,
		&v.ShipmentStatisticsGroup,
		&v.SortString,
		&v.TotalRoundAmountTransactions,
		&v.TotalTransactionsTestedForRoundAmount,
		&v.TradingPartnerCompanyID,
	)
	return &v, err
}

type VendorFlags struct {
	Flag2WayMatchPOsOnly                       BoolFromFloat `json:"Flag: 2WayMatchPOsOnly" sql:"Flag: 2WayMatchPOsOnly"`
	FlagAllVendorInvoicesWithoutPO             BoolFromFloat `json:"Flag: All Vendor Invoices Without PO" sql:"Flag: All Vendor Invoices Without PO"`
	FlagBankAccountAddedToDormantVendor        BoolFromFloat `json:"Flag: Bank Account Added to Dormant Vendor" sql:"Flag: Bank Account Added to Dormant Vendor"`
	FlagConsInv                                BoolFromFloat `json:"Flag: ConsInvVendor" sql:"Flag: ConsInvVendor"`
	FlagCountryWithHighEnforcementActivity     BoolFromFloat `json:"Flag: Country with High Enforcement Activity" sql:"Flag: Country with High Enforcement Activity"`
	FlagCountryWithSL1CaseHistory              BoolFromFloat `json:"Flag: Country with SL1 Case History" sql:"Flag: Country with SL1 Case History"`
	FlagDormantVendor                          BoolFromFloat `json:"Flag: Dormant Vendor" sql:"Flag: Dormant Vendor"`
	FlagFIInvoice                              BoolFromFloat `json:"Flag: FI Invoice" sql:"Flag: FI Invoice"`
	FlagHadActivityLastMonth                   BoolFromFloat `json:"Flag: Had Activity Last Month" sql:"Flag: Had Activity Last Month"`
	FlagHasSplitInvoice                        BoolFromFloat `json:"Flag: Has Split Invoice" sql:"Flag: Has Split Invoice"`
	FlagIncreasingInvoiceValue                 BoolFromFloat `json:"Flag: Increasing Invoice Value" sql:"Flag: Increasing Invoice Value"`
	FlagMissingOrNonStandardDuns               BoolFromFloat `json:"Flag: Missing or Non-Standard DUNS" sql:"Flag: Missing or Non-Standard DUNS"`
	FlagMissingOrNonStandardGuid               BoolFromFloat `json:"Flag: Missing or Non-Standard GUID" sql:"Flag: Missing or Non-Standard GUID"`
	FlagMissingVendorMasterFields              BoolFromFloat `json:"Flag: Missing Vendor Master Fields" sql:"Flag: Missing Vendor Master Fields"`
	FlagNoReceiptNeeded                        BoolFromFloat `json:"Flag: NoReceiptNeeded" sql:"Flag: NoReceiptNeeded"`
	FlagOnlyExternalDomainsAreKnownFreeDomains BoolFromFloat `json:"Flag: Only External Domains are Known Free Domains" sql:"Flag: Only External Domains are Known Free Domains"`
	FlagPOBoxAddress                           BoolFromFloat `json:"Flag: PO Box Address" sql:"Flag: PO Box Address"`
	FlagPossibleGD43                           BoolFromFloat `json:"Flag: Possible GD43" sql:"Flag: Possible GD43"`
	FlagPossibleSOE                            BoolFromFloat `json:"Flag: Possible SOE" sql:"Flag: Possible SOE"`
	FlagRecentBankAccountChange                BoolFromFloat `json:"Flag: Recent Bank Account Change" sql:"Flag: Recent Bank Account Change"`
	FlagSCP                                    bool          `json:"Flag: SCP" sql:"Flag: SCP"`
	FlagServiceVendor                          BoolFromFloat `json:"Flag: ServiceVendor" sql:"Flag: ServiceVendor"`
	FlagSwiftFirstTransaction                  BoolFromFloat `json:"Flag: SwiftFirstTransaction" sql:"Flag: SwiftFirstTransaction"`
	FlagTop5BenfordDeviation                   BoolFromFloat `json:"Flag: Top5BenfordDeviation" sql:"Flag: Top5BenfordDeviation"`
	FlagTop5ZScorevendors                      BoolFromFloat `json:"Flag: Top5 Z-ScoreVendors" sql:"Flag: Top5 Z-ScoreVendors"`
	FlagUnblockedPurchBlock                    BoolFromFloat `json:"Flag: UnblockedPurchBlock" sql:"Flag: UnblockedPurchBlock"`
	FlagVendorFormerEmployee                   BoolFromFloat `json:"Flag: Vendor Former Emploee" sql:"Flag: Vendor Former Emploee"`
	FlagVendorPOSingleApproval                 BoolFromFloat `json:"Flag: VendorPOSingleApproval" sql:"Flag: VendorPOSingleApproval"`
	FlagVendorPOsNotApproved                   BoolFromFloat `json:"Flag: VendorPOsNotApproved" sql:"Flag: VendorPOsNotApproved"`
	FlagVendorWithoutGUIDActive                BoolFromFloat `json:"Flag: Vendor Without GUID Active" sql:"Flag: Vendor Without GUID, Active"`
}

func (vf *VendorFlags) IsICMEntity() bool { return true }
func (vf *VendorFlags) GetID() string     { return "" }

func (vf *VendorFlags) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return json.Unmarshal([]byte(v), vf)
	case []byte:
		return json.Unmarshal(v, vf)
	default:
		return errors.New("invalid sql return type for Vendor Flags")
	}
}
