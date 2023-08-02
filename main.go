/*
Copyright © 2023 Alexander Orban <alexander.orban@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/snowflakedb/gosnowflake"
	"icm_goapi/encoding"
	"icm_goapi/orm"
	"log"
	"os"
	"strings"
)

type test struct {
	entity   orm.ICMEntity
	database string
	f        func(*sql.Rows) (orm.ICMEntity, error)
}

var tests = []test{
	{&orm.BankMaster{Flags: &orm.BankMasterFlags{}}, "DEV_FPA_LI.CMP.BANKSMASTER_FLAGGED", orm.BankMasterFromRow},
	{&orm.BankUsedBy{Flags: &orm.BanksUsedByFlags{}}, "DEV_FPA_LI.CMP.BANKSUSEDBY_FLAGGED", orm.BankUsedByFromRow},
	{&orm.BillingDocument{Flags: &orm.BillingDocumentFlags{}}, "DEV_FPA_LI.CMP.BILLINGDOCS_FLAGGED", orm.BillingDocFromRow},
	{&orm.BusinessPartner{Flags: &orm.BusinessPartnerFlags{}}, "DEV_FPA_LI.CMP.PROJECT_BPS_FLAGGED", orm.BusinessPartnerFromRow},
	{&orm.Check{}, "DEV_FPA_LI.CMP.CHECKS", orm.CheckFromRows},
	{&orm.Customer{Flags: &orm.CustomerFlags{}}, "DEV_FPA_LI.CMP.CUSTOMER_MASTER_FLAGGED", orm.CustomerFromRows},
	{&orm.FICustomer{}, "DEV_FPA_LI.CMP.CUSTOMERFI", orm.FICustomerFromRows},
	{&orm.FIVendor{Flags: &orm.FIVendorFlags{}}, "DEV_FPA_LI.CMP.VENDORFI_FLAGGED", orm.FIVendorFromRows},
	{&orm.Invoice{Flags: &orm.InvoiceFlags{}}, "DEV_FPA_LI.CMP.MMINVOICES_FLAGGED", orm.InvoiceFromRows},
	{&orm.PaymentRun{}, "DEV_FPA_LI.CMP.PAYMENTRUNS", orm.PaymentRunFromRows},
	{&orm.Project{Flags: &orm.ProjectFlags{}}, "DEV_FPA_LI.CMP.PROJECTS_FLAGGED", orm.ProjectFromRow},
	{&orm.PurchaseOrder{Flags: &orm.PurchaseOrderFlags{}}, "DEV_FPA_LI.CMP.PURCHASEORDERS_FLAGGED", orm.PurchaseOrderFromRow},
	{&orm.PurchaseOrderProject{Flags: &orm.PurchaseOrderProjectFlags{}}, "DEV_FPA_LI.CMP.PROJECTS_PURCHASEORDERS_FLAGGED", orm.PurchaseOrderProjectFromRow},
	{&orm.Requisition{}, "DEV_FPA_LI.CMP.REQUISITIONS", orm.RequisitionFromRow},
	{&orm.SalesOrder{}, "DEV_FPA_LI.CMP.SALESORDERS", orm.SalesOrderFromRow},
	{&orm.SalesOrderProject{Flags: &orm.SalesOrderProjectFlags{}}, "DEV_FPA_LI.CMP.PROJECTS_SALESORDERS_FLAGGED", orm.SalesOrderProjectFromRow},
	{&orm.Vendor{Flags: &orm.VendorFlags{}}, "DEV_FPA_LI.CMP.VENDOR_MASTER_FLAGGED", orm.VendorFromRow},
}

func main() {
	//cmd.Execute()
	dsn := GenerateSnowflakeDSN()
	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		log.Fatalf("failed to connect. %v, err: %v", dsn, err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	for _, t := range tests {
		fmt.Println(t.database)
		qry := encoding.MarshalToSelect(t.entity, t.database, false)
		err := os.WriteFile(fmt.Sprintf("./ignore/sql/%s.sql", t.database), qry, 0644)
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := db.Prepare(string(qry) + ` LIMIT 25`)
		if err != nil {
			log.Fatal(err)
		}

		var entities []orm.ICMEntity
		rows, err := stmt.Query()
		for rows.Next() {
			var e orm.ICMEntity
			var err error
			e, err = t.f(rows)
			if err != nil {
				log.Fatal(err)
			}

			entities = append(entities, e)
		}

		b, err := json.MarshalIndent(entities, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(fmt.Sprintf("./ignore/json/%s.json", t.database), b, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GenerateSnowflakeDSN() string {
	var (
		ret         = strings.Builder{}
		user string = os.Getenv("ABB_SNOWFLAKE_USER")
		pass string = os.Getenv("ABB_SNOWFLAKE_PASS")
	)

	ret.WriteString(user)
	ret.WriteString(":")
	ret.WriteString(pass)
	ret.WriteString("@")
	ret.WriteString("abb_zc.west-europe.azure")
	ret.WriteString("/")
	ret.WriteString("DEV_FPA_LI")
	ret.WriteString("/")
	ret.WriteString("SAP_AMEA")

	return ret.String()
}
