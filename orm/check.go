package orm

import (
	"database/sql"
)

type Check struct {
	// Identifiers
	Database              string   `json:"Database" sql:"Database"`
	RunOn                 NullTime `json:"Run On" sql:"Run On"`
	PrintDate             NullTime `json:"Print Date" sql:"Print Date"`
	PrintUser             string   `json:"Print User" sql:"Print User"`
	ProbablePaymentDate   NullTime `json:"Probable Payment Date" sql:"Probable Payment Date"`
	PaymentDocumentNumber string   `json:"Payment Document Number" sql:"Payment Document Number"`
	CheckIdentification   string   `json:"Identification" sql:"Identification"`
	UpperCheckNumber      string   `json:"Upper Check Number" sql:"Upper Check Number"`
	LowerCheckNumber      string   `json:"Lower Check Number" sql:"Lower Check Number"`
	ManualCheck           string   `json:"Manual Check" sql:"Manual Check"`
	PaidCheck             string   `json:"Paid Check" sql:"Paid Check"`

	// Assignments
	CompanyCodeID             string `json:"Company Code ID" sql:"Company Code ID"`
	IntercompanyCompanyCodeId string `json:"Intercompany Company Code ID" sql:"Intercompany Company Code ID"`
	BusinessAreaCode          string `json:"Business Area Code" sql:"Business Area Code"`
	GLAccountCode             string `json:"GL Account Code" sql:"GL Account Code"`
	FiscalYear                string `json:"Fiscal Year" sql:"Fiscal Year"`
	PaymentMethodCode         string `json:"Payment Method Code" sql:"Payment Method Code"`
	HouseBankCode             string `json:"House Bank Code" sql:"House Bank Code"`
	HouseBankAccountId        string `json:"House Bank Account ID" sql:"House Bank Account ID"`

	// Associated Entities
	VendorCode   string `json:"Vendor Code" sql:"Vendor Code"`
	CustomerCode string `json:"Customer Code" sql:"Customer Code"`
	PersonNumber string `json:"Person Number" sql:"Person Number"`

	// Payee Details
	PayeeCode            string `json:"Payee Code" sql:"Payee Code"`
	PayeeTitle           string `json:"Payee Title" sql:"Payee Title"`
	PayeeName1           string `json:"Payee Name 1" sql:"Payee Name 1"`
	PayeeName2           string `json:"Payee Name 2" sql:"Payee Name 2"`
	PayeeName3           string `json:"Payee Name 3" sql:"Payee Name 3"`
	PayeeName4           string `json:"Payee Name 4" sql:"Payee Name 4"`
	PayeePostalCode      string `json:"Payee Postal Code" sql:"Payee Postal Code"`
	PayeeCity            string `json:"Payee City" sql:"Payee City"`
	PayeeStreet          string `json:"Payee Street" sql:"Payee Street"`
	PayeePOBox           string `json:"Payee PO Box" sql:"Payee PO Box"`
	PayeePOBoxCity       string `json:"Payee PO Box City" sql:"Payee PO Box City"`
	PayeePOBoxPostalCode string `json:"Payee PO Box Postal Code" sql:"Payee PO Box Postal Code"`
	PayeeCountry         string `json:"Payee Country" sql:"Payee Country"`
	PayeeRegion          string `json:"Payee Region" sql:"Payee Region"`

	// Amounts
	Currency        string  `json:"Currency" sql:"Currency"`
	AmountPaid      float64 `json:"Amount Paid" sql:"Amount Paid"`
	AmountPaidInLC  float64 `json:"Amount Paid in LC" sql:"Amount Paid in LC"`
	AmountPaidInUSD float64 `json:"Amount Paid in USD" sql:"Amount Paid in USD"`

	// Void Details
	VoidReasonCode        string   `json:"Void Reason Code" sql:"Void Reason Code"`
	VoidDate              NullTime `json:"Void Date" sql:"Void Date"`
	VoidUser              string   `json:"Void User" sql:"Void User"`
	ReplacedByCheckNumber string   `json:"Replaced By Check Number" sql:"Replaced By Check Number"`
}

func (c *Check) IsICMEntity() bool { return true }

func CheckFromRows(rows *sql.Rows) (ICMEntity, error) {
	var c Check
	err := rows.Scan(
		&c.Database,
		&c.RunOn,
		&c.PrintDate,
		&c.PrintUser,
		&c.ProbablePaymentDate,
		&c.PaymentDocumentNumber,
		&c.CheckIdentification,
		&c.UpperCheckNumber,
		&c.LowerCheckNumber,
		&c.ManualCheck,
		&c.PaidCheck,
		&c.CompanyCodeID,
		&c.IntercompanyCompanyCodeId,
		&c.BusinessAreaCode,
		&c.GLAccountCode,
		&c.FiscalYear,
		&c.PaymentMethodCode,
		&c.HouseBankCode,
		&c.HouseBankAccountId,
		&c.VendorCode,
		&c.CustomerCode,
		&c.PersonNumber,
		&c.PayeeCode,
		&c.PayeeTitle,
		&c.PayeeName1,
		&c.PayeeName2,
		&c.PayeeName3,
		&c.PayeeName4,
		&c.PayeePostalCode,
		&c.PayeeCity,
		&c.PayeeStreet,
		&c.PayeePOBox,
		&c.PayeePOBoxCity,
		&c.PayeePOBoxPostalCode,
		&c.PayeeCountry,
		&c.PayeeRegion,
		&c.Currency,
		&c.AmountPaid,
		&c.AmountPaidInLC,
		&c.AmountPaidInUSD,
		&c.VoidReasonCode,
		&c.VoidDate,
		&c.VoidUser,
		&c.ReplacedByCheckNumber,
	)
	return &c, err
}
