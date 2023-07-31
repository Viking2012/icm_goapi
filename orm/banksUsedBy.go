package orm

import "time"

type BankUsedBy struct {
	// Identifiers
	BankID            string `json:"Bank_ID" sql:"Bank_ID"`
	Database          string `json:"Database" sql:"Database"`
	BankAccount       string `json:"Bank Account" sql:"Bank Account"`
	BankCode          string `json:"Bank Code" sql:"Bank Code"`
	BankCountryKey    string `json:"Bank Country Key" sql:"Bank Country Key"`
	BankA             string `json:"BANKA" sql:"BANKA"`
	BankAccountConcat string `json:"BankAccount_Concat" sql:"BankAccount_Concat"`
	Swift             string `json:"SWIFT" sql:"SWIFT"`
	VendorCode        string `json:"LIFNR" sql:"LIFNR"`
	Vendor            string `json:"Vendor" sql:"Vendor"`

	// Creation Date
	CreatedBy    string    `json:"Bank Record Created By" sql:"Bank Record Created By"`
	CreationDate time.Time `json:"Bank Record Creation Date" sql:"Bank Record Creation Date"`

	// Transactional summary
	AmountUSDLastMonth    float64 `json:"Amount USD Last Month" sql:"Amount USD Last Month"`
	AmountUSDCurrentYear  float64 `json:"Amount USD Current Year" sql:"Amount USD Current Year"`
	AmountUSDCurrentYear1 float64 `json:"Amount USD Current Year -1" sql:"Amount USD Current Year -1"`
	AmountUSDCurrentYear2 float64 `json:"Amount USD Current Year -2" sql:"Amount USD Current Year -2"`

	// Blocks
	BlockPosting  string `json:"Central posting block" sql:"Central posting block"`
	BlockPurchase string `json:"Central Purchase Block" sql:"Central Purchase Block"`
	DeletionFlag  string `json:"Vendor Deletion Flag (Central)" sql:"Vendor Deletion Flag (Central)"`

	// Other Assignments
	AccountGroup string `json:"Vendor Account Group" sql:"Vendor Account Group"`

	// Alternate Payee Details
	AlternatePayeeCountry      string `json:"Vendor Alternate Payee Country" sql:"Vendor Alternate Payee Country"`
	AlternatePayee             string `json:"Vendor Alternate Payee" sql:"Vendor Alternate Payee"`
	AlternatePayeeAccountGroup string `json:"Vendor Alternate Payee Account Group" sql:"Vendor Alternate Payee Account Group"`

	// Address Details
	StreetAddress string `json:"Address" sql:"Address"`
	Country       string `json:"Vendor Country" sql:"Vendor Country"`
	City          string `json:"City" sql:"City"`

	// GIS Data
	GUID       string `json:"Vendor GUID" sql:"Vendor GUID"`
	GUIDOld    string `json:"Vendor Old GUID" sql:"Vendor Old GUID"`
	GUIDParent string `json:"GIS_PRPARENTGUID" sql:"GIS_PRPARENTGUID"`

	// Flags
	Flags                      BanksUsedByFlags `json:"flags" sql:"FLAGS"`
	ListMissingFields          string           `json:"List Missing Fields" sql:"List Missing Fields"`
	DescriptivePurchBlock      string           `json:"PurchBlockDescriptive" sql:"PurchBlockDescriptive"`
	DomainListing              string           `json:"Domain Listing" sql:"Domain Listing"`
	FlagSplitInvoiceTotalInUsd float64          `json:"Flag: Split Invoice Total in USD" sql:"Flag: Split Invoice Total in USD"`
	RoundAmountRatio           float64          `json:"Round Amount Ratio" sql:"Round Amount Ratio"`

	// Unassigned fields
	IsIntercompany bool `json:"Is Intercompany" sql:"Is Intercompany"`
}

func (bub BankUsedBy) GetFlags() ICMEntity {
	return bub.Flags
}

type BanksUsedByFlags struct {
	FlagPoBoxAddress                           bool `json:"Flag: PO Box Address" sql:"Flag: PO Box Address"`
	FlagMissingVendorMasterFields              bool `json:"Flag: Missing Vendor Master Fields" sql:"Flag: Missing Vendor Master Fields"`
	FlagMissingOrNonStandardDuns               bool `json:"Flag: Missing or Non-Standard DUNS" sql:"Flag: Missing or Non-Standard DUNS"`
	FlagMissingOrNonStandardGuid               bool `json:"Flag: Missing or Non-Standard GUID" sql:"Flag: Missing or Non-Standard GUID"`
	FlagVendorFormerEmployee                   bool `json:"Flag: Vendor Former Emploee" sql:"Flag: Vendor Former Emploee"`
	FlagUnblockedPurchBlock                    bool `json:"Flag: UnblockedPurchBlock" sql:"Flag: UnblockedPurchBlock"`
	FlagDormantVendor                          bool `json:"Flag: Dormant Vendor" sql:"Flag: Dormant Vendor"`
	FlagRecentBankAccountChange                bool `json:"Flag: Recent Bank Account Change" sql:"Flag: Recent Bank Account Change"`
	FlagBankAccountAddedToDormantVendor        bool `json:"Flag: Bank Account Added to Dormant Vendor" sql:"Flag: Bank Account Added to Dormant Vendor"`
	FlagOnlyExternalDomainsAreKnownFreeDomains bool `json:"Flag: Only External Domains are Known Free Domains" sql:"Flag: Only External Domains are Known Free Domains"`
	FlagServiceVendor                          bool `json:"Flag: ServiceVendor" sql:"Flag: ServiceVendor"`
	FlagHasSplitInvoice                        bool `json:"Flag: Has Split Invoice" sql:"Flag: Has Split Invoice"`
	FlagConsInvVendor                          bool `json:"Flag: ConsInvVendor" sql:"Flag: ConsInvVendor"`
	FlagAllVendorInvoicesWithoutPo             bool `json:"Flag: All Vendor Invoices Without PO" sql:"Flag: All Vendor Invoices Without PO"`
	FlagSwiftFirstTransaction                  bool `json:"Flag: SwiftFirstTransaction" sql:"Flag: SwiftFirstTransaction"`
	FlagIncreasingInvoiceValue                 bool `json:"Flag: Increasing Invoice Value" sql:"Flag: Increasing Invoice Value"`
	FlagFiInvoice                              bool `json:"Flag: FI Invoice" sql:"Flag: FI Invoice"`
	FlagTop5ZScoreVendors                      bool `json:"Flag: Top5 Z-ScoreVendors" sql:"Flag: Top5 Z-ScoreVendors"`
	FlagTop5BenfordDeviation                   bool `json:"Flag: Top5BenfordDeviation" sql:"Flag: Top5BenfordDeviation"`
	FlagVendorPOsNotApproved                   bool `json:"Flag: VendorPOsNotApproved" sql:"Flag: VendorPOsNotApproved"`
	FlagVendorPOSingleApproval                 bool `json:"Flag: VendorPOSingleApproval" sql:"Flag: VendorPOSingleApproval"`
	FlagNoReceiptNeeded                        bool `json:"Flag: NoReceiptNeeded" sql:"Flag: NoReceiptNeeded"`
	Flag2WayMatchPOsOnly                       bool `json:"Flag: 2WayMatchPOsOnly" sql:"Flag: 2WayMatchPOsOnly"`
	FlagCountryWithSl1CaseHistory              bool `json:"Flag: Country with SL1 Case History" sql:"Flag: Country with SL1 Case History"`
	FlagHadActivityLastMonth                   bool `json:"Flag: Had Activity Last Month" sql:"Flag: Had Activity Last Month"`
}

func (bubf BanksUsedByFlags) GetFlags() ICMEntity {
	return bubf
}
