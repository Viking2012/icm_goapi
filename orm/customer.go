package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

type Customer struct {
	// Identifiers
	ID             string   `json:"ICM_ID" sql:"ICM_ID"`
	Database       string   `json:"Database" sql:"Database"`
	Code           string   `json:"Customer Code" sql:"Customer Code"`
	Name1          string   `json:"Name 1" sql:"Name 1"`
	Name2          string   `json:"Name 2" sql:"Name 2"`
	Name3          string   `json:"Name 3" sql:"Name 3"`
	Name4          string   `json:"Name 4" sql:"Name 4"`
	GUIDCommonName string   `json:"ZCA_GIS_1_PRCOMMONNAME" sql:"ZCA_GIS_1_PRCOMMONNAME"`
	GISLegalName   string   `json:"ZCA_GIS_1_PRLEGALNAMELOC" sql:"ZCA_GIS_1_PRLEGALNAMELOC"`
	VendorCode     string   `json:"KNA1_LIFNR" sql:"KNA1_LIFNR"`
	CreatedBy      string   `json:"Created By" sql:"Created By"`
	CreationDate   NullTime `json:"Created On" sql:"Created On"`

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
	BlockOrder    string `json:"Order Block" sql:"Order Block"`
	BlockPosting  string `json:"Posting Block" sql:"Posting Block"`
	BlockBilling  string `json:"Billing Block" sql:"Billing Block"`
	BlockDelivery string `json:"Delivery Block" sql:"Delivery Block"`
	DeletionFlag  string `json:"Deletion Indicator" sql:"Deletion Indicator"`

	// BA/Division Assignments
	PrimaryBU       string `json:"Primary BU" sql:"Primary BU"`
	PrimaryDivision string `json:"Primary Division" sql:"Primary Division"`

	// Other Assignments
	AccountGroup           string `json:"Customer Account Group" sql:"Customer Account Group"`
	AccountGroupCode       string `json:"Customer Account Group Code" sql:"Customer Account Group Code"`
	GUIDIndustryUsageLevel string `json:"Industry Usage Level" sql:"Industry Usage Level"`
	IndustryUsageCode      string `json:"Industry Usage Code" sql:"Industry Usage Code"`
	IndustryUsage          string `json:"Industry Usage" sql:"Industry Usage"`
	GUIDChannelCode        string `json:"Channel Code" sql:"Channel Code"`
	GUIDChannelName        string `json:"Channel" sql:"Channel"`

	// Address Details
	AddressNumber string `json:"Address Number" sql:"Address Number"`
	StreetAddress string `json:"Street" sql:"Street"`
	City          string `json:"City" sql:"City"`
	Region        string `json:"Region" sql:"Region"`
	Country       string `json:"Country" sql:"Country"`
	CountryCode   string `json:"Country Code" sql:"Country Code"`
	PostalCode    string `json:"Postal Code" sql:"Postal Code"`
	GUIDCity      string `json:"ZCA_GIS_1_PRCITY1" sql:"ZCA_GIS_1_PRCITY1"`
	GUIDCountry   string `json:"ZCA_GIS_1_PRCOUNTRY1" sql:"ZCA_GIS_1_PRCOUNTRY1"`

	// Contact Details
	Telephone1 string `json:"Telephone 1" sql:"Telephone 1"`
	Fax1       string `json:"Fax 1" sql:"Fax 1"`

	// Tax information
	VAT                string `json:"VAT ID" sql:"VAT ID"`
	TaxCode1           string `json:"Tax Code 1" sql:"Tax Code 1"`
	TaxCode2           string `json:"Tax Code 2" sql:"Tax Code 2"`
	TaxCode3           string `json:"Tax Code 3" sql:"Tax Code 3"`
	TaxCode4           string `json:"Tax Code 4" sql:"Tax Code 4"`
	TaxJurisdiction    string `json:"Tax Jurisdiction" sql:"Tax Jurisdiction"`
	TransportationZone string `json:"Transportation Zone" sql:"Transportation Zone"`

	// GIS Data
	GUID         string `json:"GUID" sql:"GUID"`
	GUIDOld      string `json:"Old GUID" sql:"Old GUID"`
	GUIDDomestic string `json:"Domestic GUID" sql:"Domestic GUID"`
	GUIDParent   string `json:"Parent GUID" sql:"Parent GUID"`
	DUNS         string `json:"ZCA_GIS_1_DUNS" sql:"ZCA_GIS_1_DUNS"`
	DUNSParent   string `json:"ZCA_GIS_1_PRPARENTDUNS" sql:"ZCA_GIS_1_PRPARENTDUNS"`
	DUNSDomestic string `json:"ZCA_GIS_1_PRDOMESTICDUNS" sql:"ZCA_GIS_1_PRDOMESTICDUNS"`

	// Flags
	Flags             *CustomerFlags `json:"flags" sql:"FLAGS"`
	FlagSCPStatus     string         `json:"Flag: SCP, Status" sql:"Flag: SCP, Status"`
	FlagSCPRiskRating string         `json:"Flag: SCP, Risk Rating" sql:"Flag: SCP, Risk Rating"`

	// Unassigned
	IsActive       BoolFromFloat `json:"Is Active" sql:"Is Active"`
	IsIntercompany bool          `json:"Is Intercompany" sql:"Is Intercompany"`
	TradingPartner string        `json:"Trading Partner" sql:"Trading Partner"`
	PrCode6        string        `json:"ZCA_GIS_1_PRCODE6" sql:"ZCA_GIS_1_PRCODE6"`
	OneTimeAccount string        `json:"One-Time Account" sql:"One-Time Account"`
	SortString     string        `json:"Search Term" sql:"Search Term"`
}

func (c *Customer) IsICMEntity() bool { return true }
func (c *Customer) GetID() string     { return fmt.Sprintf("%s|%s", c.Database, c.Code) }

func CustomerFromRows(rows *sql.Rows) (ICMEntity, error) {
	var c Customer
	err := rows.Scan(
		&c.ID,
		&c.Database,
		&c.Code,
		&c.Name1,
		&c.Name2,
		&c.Name3,
		&c.Name4,
		&c.GUIDCommonName,
		&c.GISLegalName,
		&c.VendorCode,
		&c.CreatedBy,
		&c.CreationDate,
		&c.AmountUSDLastMonth,
		&c.CountLastMonth,
		&c.AmountUSDCurrentYear,
		&c.CountCurrentYear,
		&c.AmountUSDCurrentYear1,
		&c.CountCurrentYear1,
		&c.AmountUSDCurrentYear2,
		&c.CountCurrentYear2,
		&c.BlockOrder,
		&c.BlockPosting,
		&c.BlockBilling,
		&c.BlockDelivery,
		&c.DeletionFlag,
		&c.PrimaryBU,
		&c.PrimaryDivision,
		&c.AccountGroup,
		&c.AccountGroupCode,
		&c.GUIDIndustryUsageLevel,
		&c.IndustryUsageCode,
		&c.IndustryUsage,
		&c.GUIDChannelCode,
		&c.GUIDChannelName,
		&c.AddressNumber,
		&c.StreetAddress,
		&c.City,
		&c.Region,
		&c.Country,
		&c.CountryCode,
		&c.PostalCode,
		&c.GUIDCity,
		&c.GUIDCountry,
		&c.Telephone1,
		&c.Fax1,
		&c.VAT,
		&c.TaxCode1,
		&c.TaxCode2,
		&c.TaxCode3,
		&c.TaxCode4,
		&c.TaxJurisdiction,
		&c.TransportationZone,
		&c.GUID,
		&c.GUIDOld,
		&c.GUIDDomestic,
		&c.GUIDParent,
		&c.DUNS,
		&c.DUNSParent,
		&c.DUNSDomestic,
		&c.Flags,
		&c.FlagSCPStatus,
		&c.FlagSCPRiskRating,
		&c.IsActive,
		&c.IsIntercompany,
		&c.TradingPartner,
		&c.PrCode6,
		&c.OneTimeAccount,
		&c.SortString,
	)
	return &c, err
}

type CustomerFlags struct {
	FlagPossibleSOE                           bool `json:"Flag: Possible SOE" sql:"Flag: Possible SOE"`
	FlagPossibleGD43                          bool `json:"Flag: Possible GD43" sql:"Flag: Possible GD43"`
	FlagHighReturnRatio                       bool `json:"FLAG: High Return Ratio" sql:"FLAG: High Return Ratio"`
	FlagHighRelativeReturnRatioWithinDivision bool `json:"FLAG: High Relative Return Ratio within Division" sql:"FLAG: High Relative Return Ratio within Division"`
	FlagPatternOfHigherEOQReturns             bool `json:"FLAG: Pattern of Higher EoQ Returns" sql:"FLAG: Pattern of Higher EoQ Returns"`
	FlagPossibleChannelStuffing               bool `json:"FLAG: Possible Channel Stuffing" sql:"FLAG: Possible Channel Stuffing"`
	FlagCMPaidToCustomer                      bool `json:"Flag: CM Paid to Customer" sql:"Flag: CM Paid to Customer"`
	FlagManualCreditMemo                      bool `json:"Flag: Manual Credit Memo" sql:"Flag: Manual Credit Memo"`
	FlagNoCmRequest                           bool `json:"Flag: No CM Request" sql:"Flag: No CM Request"`
	FlagHadActivityLastMonth                  bool `json:"Flag: Had Activity Last Month" sql:"Flag: Had Activity Last Month"`
	FlagCountryWithSl1CaseHistory             bool `json:"Flag: Country with SL1 Case History" sql:"Flag: Country with SL1 Case History"`
	FlagCountryWithHighEnforcementActivity    bool `json:"Flag: Country with High Enforcement Activity" sql:"Flag: Country with High Enforcement Activity"`
	FlagSCP                                   bool `json:"Flag: SCP" sql:"Flag: SCP"`
	FlagNotInSalesforceTPM                    bool `json:"Flag: Not In SalesForceTPM" sql:"Flag: Not In SalesForceTPM"`
	FlagNoGUID                                bool `json:"Flag: No GUID" sql:"Flag: No GUID"`
}

func (cf *CustomerFlags) IsICMEntity() bool { return true }
func (cf *CustomerFlags) GetID() string     { return "" }

func (cf *CustomerFlags) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return json.Unmarshal([]byte(v), cf)
	case []byte:
		return json.Unmarshal(v, cf)
	default:
		return errors.New("invalid sql return type for Vendor Flags")
	}
}
