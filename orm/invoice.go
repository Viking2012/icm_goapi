package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

type Invoice struct {
	Database                      string        `json:"Database" sql:"Database"`
	DocumentNumber                string        `json:"Document Number" sql:"Document Number"`
	FiscalYear                    string        `json:"Fiscal Year" sql:"Fiscal Year"`
	LineItem                      string        `json:"Line item" sql:"Line item"`
	SeqNo                         string        `json:"Seq. no." sql:"Seq. no."`
	DocumentType                  string        `json:"Document Type" sql:"Document Type"`
	Rbkp_Bldat                    NullTime      `json:"RBKP_BLDAT" sql:"RBKP_BLDAT"`
	PostingDateAccountingDocument NullTime      `json:"Posting Date" sql:"Posting Date"`
	Rbkp_Usnam                    string        `json:"RBKP_USNAM" sql:"RBKP_USNAM"`
	Rbkp_Cpudt                    NullTime      `json:"RBKP_CPUDT" sql:"RBKP_CPUDT"`
	CompanyCode                   string        `json:"Company Code" sql:"Company Code"`
	Company                       string        `json:"Company" sql:"Company"`
	VendorNumber                  string        `json:"Vendor Number" sql:"Vendor Number"`
	Currency                      string        `json:"Currency" sql:"Currency"`
	Rbkp_Stblg                    string        `json:"RBKP_STBLG" sql:"RBKP_STBLG"`
	Rbkp_Stjah                    string        `json:"RBKP_STJAH" sql:"RBKP_STJAH"`
	InvoiceStatus                 string        `json:"Invoice Status" sql:"Invoice Status"`
	PurchasingDocumentNumber      string        `json:"Purchasing Document Number" sql:"Purchasing Document Number"`
	PurchasingDocumentLineItem    string        `json:"Purchasing Document Line Item" sql:"Purchasing Document Line Item"`
	MaterialCode                  string        `json:"Material Code" sql:"Material Code"`
	Material                      string        `json:"Material" sql:"Material"`
	CompanyCodeId                 string        `json:"Company Code ID" sql:"Company Code ID"`
	PlantCode                     string        `json:"Plant Code" sql:"Plant Code"`
	DebitCreditIndicator          string        `json:"Debit/Credit Indicator" sql:"Debit/Credit Indicator"`
	WbsElement                    string        `json:"WBS Element" sql:"WBS Element"`
	ProjectNumber                 string        `json:"Project Number" sql:"Project Number"`
	ProjectName                   string        `json:"PROJ_Concat" sql:"PROJ_Concat"`
	Quantity                      float64       `json:"Quantity" sql:"Quantity"`
	TaxCode                       string        `json:"Tax Code" sql:"Tax Code"`
	QuantityInOrderPriceUnit      float64       `json:"Quantity in Order Price Unit" sql:"Quantity in Order Price Unit"`
	Rbco_Nplnr                    string        `json:"RBCO_NPLNR" sql:"RBCO_NPLNR"`
	Rbco_Vbeln                    string        `json:"RBCO_VBELN" sql:"RBCO_VBELN"`
	Rbco_Vbelp                    string        `json:"RBCO_VBELP" sql:"RBCO_VBELP"`
	MaterialGroupCode             string        `json:"Material Group Code" sql:"Material Group Code"`
	MaterialGroup                 string        `json:"Material Group" sql:"Material Group"`
	AmountDcWithoutShipCharges    float64       `json:"Amount DC Without Ship Charges" sql:"Amount DC Without Ship Charges"`
	DcsShare                      float64       `json:"DCs share" sql:"DCs share"`
	Amount                        float64       `json:"Amount" sql:"Amount"`
	AmountInLc                    float64       `json:"Amount in LC" sql:"Amount in LC"`
	AmountInUsd                   float64       `json:"Amount in USD" sql:"Amount in USD"`
	IsPostedToFI                  BoolFromFloat `json:"Posted to FI" sql:"Posted to FI"`
	ProfitCenter                  string        `json:"Profit Center" sql:"Profit Center"`
	Profitcenter_Org1             string        `json:"ProfitCenter_ORG1" sql:"ProfitCenter_ORG1"`
	Profitcenter_Org1_Concat      string        `json:"ProfitCenter_ORG1_Concat" sql:"ProfitCenter_ORG1_Concat"`
	Profitcenter_Org2             string        `json:"ProfitCenter_ORG2" sql:"ProfitCenter_ORG2"`
	Profitcenter_Org2_Concat      string        `json:"ProfitCenter_ORG2_Concat" sql:"ProfitCenter_ORG2_Concat"`
	Profitcenter_Org3             string        `json:"ProfitCenter_ORG3" sql:"ProfitCenter_ORG3"`
	Profitcenter_Org3_Concat      string        `json:"ProfitCenter_ORG3_Concat" sql:"ProfitCenter_ORG3_Concat"`
	AbacusCountry                 string        `json:"Abacus Country" sql:"Abacus Country"`
	ReportingUnit                 string        `json:"Reporting Unit" sql:"Reporting Unit"`
	VendorID                      string        `json:"ICM_VENDOR_ID" sql:"ICM_VENDOR_ID"`
	DescriptiveServiceType        string        `json:"ServiceTypeDescriptive" sql:"ServiceTypeDescriptive"`
	Flags                         *InvoiceFlags `json:"flags" sql:"FLAGS"`
}

func (i *Invoice) IsICMEntity() bool { return true }
func (i *Invoice) GetID() string {
	return fmt.Sprintf(
		"%s|%s|%s|%s",
		i.Database,
		i.DocumentNumber,
		i.FiscalYear,
		i.LineItem,
	)
}

func InvoiceFromRows(rows *sql.Rows) (ICMEntity, error) {
	var i Invoice
	err := rows.Scan(
		&i.Database,
		&i.DocumentNumber,
		&i.FiscalYear,
		&i.LineItem,
		&i.SeqNo,
		&i.DocumentType,
		&i.Rbkp_Bldat,
		&i.PostingDateAccountingDocument,
		&i.Rbkp_Usnam,
		&i.Rbkp_Cpudt,
		&i.CompanyCode,
		&i.Company,
		&i.VendorNumber,
		&i.Currency,
		&i.Rbkp_Stblg,
		&i.Rbkp_Stjah,
		&i.InvoiceStatus,
		&i.PurchasingDocumentNumber,
		&i.PurchasingDocumentLineItem,
		&i.MaterialCode,
		&i.Material,
		&i.CompanyCodeId,
		&i.PlantCode,
		&i.DebitCreditIndicator,
		&i.WbsElement,
		&i.ProjectNumber,
		&i.ProjectName,
		&i.Quantity,
		&i.TaxCode,
		&i.QuantityInOrderPriceUnit,
		&i.Rbco_Nplnr,
		&i.Rbco_Vbeln,
		&i.Rbco_Vbelp,
		&i.MaterialGroupCode,
		&i.MaterialGroup,
		&i.AmountDcWithoutShipCharges,
		&i.DcsShare,
		&i.Amount,
		&i.AmountInLc,
		&i.AmountInUsd,
		&i.IsPostedToFI,
		&i.ProfitCenter,
		&i.Profitcenter_Org1,
		&i.Profitcenter_Org1_Concat,
		&i.Profitcenter_Org2,
		&i.Profitcenter_Org2_Concat,
		&i.Profitcenter_Org3,
		&i.Profitcenter_Org3_Concat,
		&i.AbacusCountry,
		&i.ReportingUnit,
		&i.VendorID,
		&i.DescriptiveServiceType,
		&i.Flags,
	)
	return &i, err
}

type InvoiceFlags struct {
	FlagIncreasingInvoiceValue bool `json:"Flag: Increasing Invoice Value" sql:"Flag: Increasing Invoice Value"`
	FlagServiceVendor          bool `json:"Flag: ServiceVendor" sql:"Flag: ServiceVendor"`
}

func (i *InvoiceFlags) IsICMEntity() bool { return true }

func (i *InvoiceFlags) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		return json.Unmarshal([]byte(v), i)
	case []byte:
		return json.Unmarshal(v, i)
	default:
		return errors.New("invalid sql return type for InvoiceFlags")
	}
}
