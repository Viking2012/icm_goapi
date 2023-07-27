package orm

import "time"

type BankMaster struct {
	// Identifiers
	ID                string `json:"Bank_ID"`
	Database          string `json:"Database"`
	BankAccount       string `json:"Bank Account"`
	BankCode          string `json:"Bank Code"`
	BankCountryKey    string `json:"Bank Country Key"`
	BankA             string `json:"BANKA"`
	BankAccountConcat string `json:"BankAccount_Concat"`
	Swift             string `json:"SWIFT"`
	NewBankN          string `json:"New_BANKN"`
	NewBankL          string `json:"New_BANKL"`
	NewBankS          string `json:"New_BANKS"`

	// Creation date
	CreatedBy    string    `json:"Bank Record Created By"`
	CreationDate time.Time `json:"Bank Record Creation Date"`

	// Flags
	FlagAlternativePayee                                        bool         `json:"Flag: AlternativePayee"`
	FlagAlternativePayeeRotated                                 bool         `json:"Flag: Alternative Rotated"`
	FlagAlternativePayeesMultiple                               bool         `json:"Flag: Multiple Alternative Payees"`
	DescriptiveAlternativePayeeNames                            string       `json:"Descriptive: Alternative Payee Names"`
	FlagBankUsedForSingleDocType                                bool         `json:"Flag: Bank Used for Single Doc Type"`
	FlagBlockedVendorBankAccountPaid                            bool         `json:"Flag: Blocked Vendor Bank Account Paid"`
	FlagCountOfSecondaryBankVendors                             IntFromFloat `json:"Flag: Count of Secondary Bank Vendors"`
	DescriptiveSecondaryBankVendors                             string       `json:"Secondary Bank Vendors"`
	FlagDisappearingBankAccountFiOnlyInvoicesPaid               bool         `json:"Flag: Disappearing Bank Account; FI Only Invoices Paid"`
	FlagDisappearingBankAccountFiOnlyInvoicesPaidLastMonth      bool         `json:"Flag: Disappearing Bank Account; FI Only Invoices Paid Last Month"`
	FlagDisappearingBankAccountOnlyLimitPoInvoicesPaid          bool         `json:"Flag: Disappearing Bank Account; Only Limit PO Invoices Paid"`
	FlagDisappearingBankAccountOnlyLimitPoInvoicesPaidLastMonth bool         `json:"Flag: Disappearing Bank Account; Only Limit PO Invoices Paid Last Month"`
	FlagNotTheLargestBank                                       bool         `json:"Flag: Not the Largest Bank"`
	FlagPaidToDisappearingAndReappearingBankAccount             bool         `json:"Flag: Paid To Disappearing and Reappearing Bank Account"`
	FlagRiskyPaymentsToDisappearingBankAccount                  bool         `json:"Flag: Risky Payments to Disappearing Bank Account"`
	FlagRiskyPaymentsToDisappearingBankAccountThisMonth         bool         `json:"Flag: Risky Payments to Disappearing Bank Account This Month"`
	FlagTravelingBankAccount                                    bool         `json:"Flag: Traveling Bank Account"`
	FlagTop5BenfordDeviation                                    bool         `json:"Flag: Top5BenfordDeviation"`
	DescriptiveBenford                                          string       `json:"BenfordDescriptive"`
	FlagVendorClusterCount                                      IntFromFloat `json:"Flag: Vendor Cluster Count"`

	// Transactional summary
	AmountTotalValueUSD float64 `json:"Sum_REGUP_WRBTR_USD"`

	// Unassigned Fields
	VendorsOnBankAcc    string `json:"VendorsPerBankAcc"`
	PayeeBankCountryKey string `json:"Payee Bank Country Key"`
	PayeeBankAcctNumber string `json:"Payee Bank Acct Number"`
	PayeeBankNumber     string `json:"Payee Bank Number"`
	VendorAccountGroup  string `json:"Vendor Account Group"`
}

type BankMasterFlags struct {
	FlagAlternativePayee                                        bool `json:"Flag: AlternativePayee"`
	FlagAlternativePayeeRotated                                 bool `json:"Flag: Alternative Rotated"`
	FlagAlternativePayeesMultiple                               bool `json:"Flag: Multiple Alternative Payees"`
	FlagBankUsedForSingleDocType                                bool `json:"Flag: Bank Used for Single Doc Type"`
	FlagBlockedVendorBankAccountPaid                            bool `json:"Flag: Blocked Vendor Bank Account Paid"`
	FlagDisappearingBankAccountFiOnlyInvoicesPaid               bool `json:"Flag: Disappearing Bank Account; FI Only Invoices Paid"`
	FlagDisappearingBankAccountFiOnlyInvoicesPaidLastMonth      bool `json:"Flag: Disappearing Bank Account; FI Only Invoices Paid Last Month"`
	FlagDisappearingBankAccountOnlyLimitPoInvoicesPaid          bool `json:"Flag: Disappearing Bank Account; Only Limit PO Invoices Paid"`
	FlagDisappearingBankAccountOnlyLimitPoInvoicesPaidLastMonth bool `json:"Flag: Disappearing Bank Account; Only Limit PO Invoices Paid Last Month"`
	FlagNotTheLargestBank                                       bool `json:"Flag: Not the Largest Bank"`
	FlagPaidToDisappearingAndReappearingBankAccount             bool `json:"Flag: Paid To Disappearing and Reappearing Bank Account"`
	FlagRiskyPaymentsToDisappearingBankAccount                  bool `json:"Flag: Risky Payments to Disappearing Bank Account"`
	FlagRiskyPaymentsToDisappearingBankAccountThisMonth         bool `json:"Flag: Risky Payments to Disappearing Bank Account This Month"`
	FlagTravelingBankAccount                                    bool `json:"Flag: Traveling Bank Account"`
	FlagTop5BenfordDeviation                                    bool `json:"Flag: Top5BenfordDeviation"`
}
