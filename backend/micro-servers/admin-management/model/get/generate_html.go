package get

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/divan/num2words"
	"github.com/nleeper/goment"
)

var bac = `background:url(https://api.malaiagency.in/uploads/invoice_bg.png);
background-repeat: no-repeat;
background-size: 345px 220px;
background-position: center;`

func GenerateHTML(requestObj map[string]interface{}, data map[string]interface{}) string {
	htmlStr := `<div style="padding:5px;font-size:14px;` + "" + `">`
	htmlStr = htmlStr + headerInfo() + InvoiceInfo(requestObj, data)
	htmlStr = htmlStr + Items(requestObj, data) + footerInfo()
	htmlStr = htmlStr + `</div>`
	var err error
	if requestObj["operation"].(string) != "preview" {
		htmlStr, err = invoiceFileObjct(htmlStr)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("ss")
	}
	return htmlStr
}

func footerInfo() string {
	str := `
	<div style="padding:10px;font-family: Arial, sans-serif;">
	<div style="font-weight:bold;">Terms & Conditions</div>
	<div style="padding-top:10px;">
	<div style="padding-bottom:2px;">1. Only Goods are sold as per the bill</div>
	<div style="padding-bottom:2px;">2. Prices are Approved from the Customer side</div>
	<div style="padding-bottom:2px;">3. One year free service from the date of installation</div>
	</div>
	</div>
	<div style="width:100%;border-bottom:1px solid #000;border-top:1px solid #000;padding:10px;font-family: Arial, sans-serif">
	<table style="width:100%;font-weight:bold;">
	<tbody>
	<tr style="font-family: Arial, sans-serif">
	<td >
	<div style="height: 66px;"></div>
	<div style="font-weight:bold;"> Customer's Signature</div>
	</td>
	<td style="text-align:right;">
	<table style="text-align:right;width:100%;">
	<tbody >
	<tr><td style="font-weight:bold;font-family: "Times New Roman", Times, serif">For MALAI AGENCY</td></tr>
	<tr><td style="height:50px"></td></tr>
	<tr><td style="font-weight:bold;">Authorised Signatory/ Proprietor</td></tr>
	</tbody>
	</table>
	</td>
	</tr>
	</tbody>
	</table>
	</div>

	`
	return str
}

func Items(requestObj map[string]interface{}, data map[string]interface{}) string {
	dataList := make([]interface{}, 0)
	if data["item_list"] != nil {
		dataList = data["item_list"].([]interface{})
	}
	str := `
	<div style="padding:10px;font-family: Arial, sans-serif;` + bac + `;min-height:750px">
	<table style="width: 100%;border-collapse: collapse;overflow:auto;page-break-inside: avoid !important;">
	<tbody>
	<tr style="font-size:14px;">
	<td style="width:40px;background: #2f639d;color: #fff;padding:6px;font-weight: bold;border: 1px solid #000;text-align:center;">Sl.No</td>
	<td style="background: #2f639d;color: #fff;padding:6px;font-weight: bold;border: 1px solid #000;text-align:center;">Item Name</td>
	<td style="width:60px;background: #2f639d;color: #fff;padding:6px;font-weight: bold;border: 1px solid #000;text-align:center;" >HSN CODE </td>
	<td style="width:50px;background: #2f639d;color: #fff;padding:6px;font-weight: bold;border: 1px solid #000;text-align:center;" >Unit/Qty</td>
	<td style="width:100px;background: #2f639d;color: #fff;padding:6px;font-weight: bold;border: 1px solid #000;text-align:center;" >Rate per item</td>
	<td style="width:150px;background: #2f639d;color: #fff;padding:6px;font-weight: bold;border: 1px solid #000;text-align:center;" >Amount</td>
    </tr>`
	for index, i := range dataList {
		obj := i.(map[string]interface{})
		str = str + `<tr style="font-size:14px;">
		<td style="padding:6px;font-weight: bold;border: 1px solid #000;text-align:center;;height:50px">` + fmt.Sprint(index+1) + `.</td>
		<td style="padding:6px;font-weight: bold;border: 1px solid #000;text-align:left;height:50px">` + fmt.Sprint(obj["item_name"]) + `</td>
		<td style="padding:6px;font-weight: bold;border: 1px solid #000;text-align:center;" >` + fmt.Sprint(obj["hsn_code"]) + `</td>
		<td style="padding:6px;font-weight: bold;border: 1px solid #000;text-align:right;" >` + fmt.Sprint(obj["unit"]) + `</td>
		<td style="padding:6px;font-weight: bold;border: 1px solid #000;text-align:right;" >` + fmt.Sprint(obj["rate_per_item"]) + `</td>
		<td style="padding:6px;font-weight: bold;border: 1px solid #000;text-align:right;" >` + fmt.Sprint(obj["total"]) + `</td>
		</tr>`
	}

	str = str + `<tr style="font-weight:bold;">
	<td colspan="4" style="padding:3px;border: 1px solid #000;border-top:0;border-right:0;font-weight: bold;height:40px">Grand Total</td>
	<td colspan="10" style="text-align:right;padding:3px 10px;border: 1px solid #000;border-top:0;border-left:0;font-weight:bold;color:#2f639d;">` + fmt.Sprint(data["total"]) + ` INR </td>
	</tr>
	<tr  style="font-weight:bold;">
	<td colspan="2" style="padding:3px;border: 1px solid #000;border-top:0;border-right:0;font-weight: bold;height:40px">Total Amount (In Words)</td>
	<td colspan="10" style="text-align:right;padding:3px 10px;border: 1px solid #000;border-top:0;border-left:0;font-weight:bold;color:#2f639d;">` + numberToWords(int(data["total"].(float64))) + `</td>
	</tr>
	</tbody>
	</table>
	<div style="font-weight:bold;font-size:14px;text-align:center;padding:5px 10px;"><i>****  End of the Charges  ****</i></div>
	</div>
	`
	return str
}

func headerInfo() string {

	str := `

	<div style="text-align:center;">
<img src="https://api.malaiagency.in/uploads/ma_logo.png" style="width: 600px;height: 150px;">
	</div>
	
	<div style = "padding: 2px 10px;color: #000;font-size: 15px;text-align: center;font-family: Arial, sans-serif"> 	
	No:87,Odiyampattu main road, Next to SundaraLakshimi Thirumana Mahal</div>
	<div style = "font-family: Arial, sans-serif;padding: 0px 10px;color: #000;font-size: 15px;text-align: center;"> Villianur, Puducherry - 605110 </div>
	<table style="width:100%;border-bottom:1px solid #000;padding:10px;font-weight:bold;">
	<tbody>
	<tr style="font-family: Arial, sans-serif">
	<td >Manager : Sugupathy.K</td>
	<td style="text-align:right;">Phone No : 7395805783, 8098737312</td>
	</tr>
	</tbody>
	</table>
	`
	return str
}

func InvoiceInfo(requestObj map[string]interface{}, data map[string]interface{}) string {
	Type := "Invoice"
	if requestObj["type"].(string) == "invoice" {
		Type = "Invoice"
	} else {
		Type = "Quotation"
	}
	str := `
	<table style="width:100%;border-bottom:1px solid #000;padding:10px;font-family: Arial, sans-serif;">
	<tbody>
	<tr>
	<td style="width:50%;" >
	<table style="width:100%">
	<tbody style="text-align:left;">
	<tr><td >
	<div>
	<div style="font-weight:bold;">Customer Info:</div> <div style="padding-top:5px;">` + fmt.Sprint(data["customer_name"]) + `</div>
	` + fmt.Sprint(data["customer_address"]) + `
	</div>
	</td></tr>
	</tbody>
	</table>

	</td>
	<td style="width:20%">
	<div >
	<table style="width:100%;">
	<tbody style="text-align:right;">
	<tr><td style="font-weight:bold;">` + Type + ` No   :</td><td>` + fmt.Sprint(data["ref_number"]) + `</td></tr>
	<tr><td style="font-weight:bold;">` + Type + ` Date :</td><td>` + dateFormat(fmt.Sprint(data["date"])) + `</td></tr>
	</tbody>
	</table>
	</div>
	</td>
	</tr>
	</tbody>
	</table>
	<div style="padding:10px;color:#004671;font-size: 20px;text-align: center;font-weight:bold;">` + Type + `</div>
	`
	return str
}
func invoiceFileObjct(htmlStr string) (string, error) {
	pdfg, err := wkhtml.NewPDFGenerator()
	pdfg.MarginBottom.Set(3)
	pdfg.MarginTop.Set(1)
	pdfg.MarginLeft.Set(1)
	pdfg.MarginRight.Set(1)
	pdfg.PageSize.Set("A4")

	if err != nil {
		return "", err
	}
	page := wkhtml.NewPageReader(strings.NewReader(htmlStr))
	// page.HeaderHTML.Set(htmlStr)
	// page.FooterLeft.Set("Page [page] of [topage]")
	// page.FooterFontSize.Set(5)
	// page.FooterLine.Set(true)
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		return "", err
	}
	//Your Pdf Name
	err = pdfg.WriteFile("./invoice.pdf")
	if err != nil {
		return "", err
	}
	file, err := os.Open("./invoice.pdf")
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(b)
	return encoded, nil
}

func dateFormat(date string) string {
	temp := ""
	if date != "" {
		Tdate, _ := goment.New(date)
		temp = Tdate.Format("ll")

	}
	return temp
}

func numberToWords(number int) string {
	if number != 0 {
		str := num2words.Convert(number)

		return strings.Title(str) + " only"
	}
	return ""
}
