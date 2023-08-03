package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type BankMaster struct {
	// Identifiers
	ID                string `json:"Bank_ID" sql:"Bank_ID|pk"`
	Database          string `json:"Database" sql:"Database"`
	BankAccount       string `json:"Bank Account" sql:"Bank Account"`
	BankCode          string `json:"Bank Code" sql:"Bank Code"`
	BankCountryKey    string `json:"Bank Country Key" sql:"Bank Country Key"`
	BankA             string `json:"BANKA" sql:"BANKA"`
	BankAccountConcat string `json:"BankAccount_Concat" sql:"BankAccount_Concat"`
	Swift             string `json:"SWIFT" sql:"SWIFT"`
	NewBankN          string `json:"New_BANKN" sql:"New_BANKN"`
	NewBankL          string `json:"New_BANKL" sql:"New_BANKL"`
	NewBankS          string `json:"New_BANKS" sql:"New_BANKS"`

	// Creation date
	CreatedBy    string   `json:"Bank Record Created By" sql:"Bank Record Created By"`
	CreationDate NullTime `json:"Bank Record Creation Date" sql:"Bank Record Creation Date"`

	// Flags
	Flags                            *BankMasterFlags `json:"flags" sql:"FLAGS"`
	FlagVendorClusterCount           int32            `json:"Flag: Vendor Cluster Count" sql:"Flag: Vendor Cluster Count"`
	FlagCountOfSecondaryBankVendors  int32            `json:"Flag: Count of Secondary Bank Vendors" sql:"Flag: Count of Secondary Bank Vendors"`
	DescriptiveSecondaryBankVendors  string           `json:"Secondary Bank Vendors" sql:"Secondary Bank Vendors"`
	DescriptiveAlternativePayeeNames string           `json:"Descriptive: Alternative Payee Names" sql:"Descriptive: Alternative Payee Names"`
	DescriptiveBenford               string           `json:"BenfordDescriptive" sql:"BenfordDescriptive"`

	// Transactional summary
	AmountTotalValueUSD float64 `json:"Sum_REGUP_WRBTR_USD" sql:"Sum_REGUP_WRBTR_USD"`

	// Unassigned Fields
	VendorsOnBankAcc    string `json:"VendorsPerBankAcc" sql:"VendorsPerBankAcc"`
	PayeeBankCountryKey string `json:"Payee Bank Country Key" sql:"Payee Bank Country Key"`
	PayeeBankAcctNumber string `json:"Payee Bank Acct Number" sql:"Payee Bank Acct Number"`
	PayeeBankNumber     string `json:"Payee Bank Number" sql:"Payee Bank Number"`
	VendorAccountGroup  string `json:"Vendor Account Group" sql:"Vendor Account Group"`
}

func (bm *BankMaster) IsICMEntity() bool { return true }

func BankMasterFromRow(rows *sql.Rows) (ICMEntity, error) {
	var bm BankMaster
	err := rows.Scan(
		&bm.ID,
		&bm.Database,
		&bm.BankAccount,
		&bm.BankCode,
		&bm.BankCountryKey,
		&bm.BankA,
		&bm.BankAccountConcat,
		&bm.Swift,
		&bm.NewBankN,
		&bm.NewBankL,
		&bm.NewBankS,
		&bm.CreatedBy,
		&bm.CreationDate,
		&bm.Flags,
		&bm.FlagVendorClusterCount,
		&bm.FlagCountOfSecondaryBankVendors,
		&bm.DescriptiveSecondaryBankVendors,
		&bm.DescriptiveAlternativePayeeNames,
		&bm.DescriptiveBenford,
		&bm.AmountTotalValueUSD,
		&bm.VendorsOnBankAcc,
		&bm.PayeeBankCountryKey,
		&bm.PayeeBankAcctNumber,
		&bm.PayeeBankNumber,
		&bm.VendorAccountGroup,
	)
	return &bm, err
}

type BankMasterFlags struct {
	FlagAlternativePayee                                        bool `json:"Flag: AlternativePayee" sql:"Flag: AlternativePayee"`
	FlagAlternativePayeeRotated                                 bool `json:"Flag: Alternative Rotated" sql:"Flag: Alternative Rotated"`
	FlagAlternativePayeesMultiple                               bool `json:"Flag: Multiple Alternative Payees" sql:"Flag: Multiple Alternative Payees"`
	FlagBankUsedForSingleDocType                                bool `json:"Flag: Bank Used for Single Doc Type" sql:"Flag: Bank Used for Single Doc Type"`
	FlagBlockedVendorBankAccountPaid                            bool `json:"Flag: Blocked Vendor Bank Account Paid" sql:"Flag: Blocked Vendor Bank Account Paid"`
	FlagDisappearingBankAccountFiOnlyInvoicesPaid               bool `json:"Flag: Disappearing Bank Account FI Only Invoices Paid" sql:"Flag: Disappearing Bank Account; FI Only Invoices Paid"`
	FlagDisappearingBankAccountFiOnlyInvoicesPaidLastMonth      bool `json:"Flag: Disappearing Bank Account FI Only Invoices Paid Last Month" sql:"Flag: Disappearing Bank Account; FI Only Invoices Paid Last Month"`
	FlagDisappearingBankAccountOnlyLimitPoInvoicesPaid          bool `json:"Flag: Disappearing Bank Account Only Limit PO Invoices Paid" sql:"Flag: Disappearing Bank Account; Only Limit PO Invoices Paid"`
	FlagDisappearingBankAccountOnlyLimitPoInvoicesPaidLastMonth bool `json:"Flag: Disappearing Bank Account Only Limit PO Invoices Paid Last Month" sql:"Flag: Disappearing Bank Account; Only Limit PO Invoices Paid Last Month"`
	FlagNotTheLargestBank                                       bool `json:"Flag: Not the Largest Bank" sql:"Flag: Not the Largest Bank"`
	FlagPaidToDisappearingAndReappearingBankAccount             bool `json:"Flag: Paid To Disappearing and Reappearing Bank Account" sql:"Flag: Paid To Disappearing and Reappearing Bank Account"`
	FlagRiskyPaymentsToDisappearingBankAccount                  bool `json:"Flag: Risky Payments to Disappearing Bank Account" sql:"Flag: Risky Payments to Disappearing Bank Account"`
	FlagRiskyPaymentsToDisappearingBankAccountThisMonth         bool `json:"Flag: Risky Payments to Disappearing Bank Account This Month" sql:"Flag: Risky Payments to Disappearing Bank Account This Month"`
	FlagTravelingBankAccount                                    bool `json:"Flag: Traveling Bank Account" sql:"Flag: Traveling Bank Account"`
	FlagTop5BenfordDeviation                                    bool `json:"Flag: Top5BenfordDeviation" sql:"Flag: Top5BenfordDeviation"`
}

func (bmf *BankMasterFlags) IsICMEntity() bool { return true }

func (bmf *BankMasterFlags) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return json.Unmarshal([]byte(v), bmf)
	case []byte:
		return json.Unmarshal(v, bmf)
	default:
		return errors.New("invalid sql return type for Vendor Flags")
	}
}
