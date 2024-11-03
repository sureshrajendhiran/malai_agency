package common

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"malai_agency/backend/services/logs"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/mssola/user_agent"
	"github.com/nleeper/goment"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// StringToMap func used to convert string json to Map
func StringToMap(str string) map[string]interface{} {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		logs.Logs("(StringToMap) String json to map[string]interface{} error ", err)
		return nil
	}
	return data
}

// MapToString func used to convert map to string
func MapToString(data interface{}) string {
	datastr, err := json.Marshal(data)
	if err != nil {
		logs.Logs("(MapToString) map[string]interface{} to String json error ", err)
		return ""
	}
	return string(datastr)
}

// StringToInterface func used to convert string to interface
func StringToInterface(str string) interface{} {
	var data interface{}
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		logs.Logs("(StringToMap) String json to map[string]interface{} error ", err)
		return nil
	}
	return data
}

// JsonValid
func JsonValid(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var data interface{}
	err := decoder.Decode(&data)
	if err != nil {
		logs.Logs("(JsonValid) invalid json error", err)
		return err
	}

	return nil
}

func AssignMap(from map[string]interface{}, to map[string]interface{}) {
	for key, value := range from {
		to[key] = value
	}
	return
}

// StrToInt func used to conver string to number interger
func StrToInt(numstr string) int {
	x, err := strconv.Atoi(numstr)
	if err != nil {
		logs.Logs("(StrToInt) string to number conver error: ", err)
		return 0
	}
	return x
}

// UserAgentParse func used to parse request agent
func UserAgentParse(r *http.Request, data map[string]interface{}) {
	ua := user_agent.New(r.Header.Get("User-Agent"))
	if ua.Mobile() {
		data["device_type"] = "mobile"
	} else if ua.Bot() {
		data["device_type"] = "other"
	} else {
		data["device_type"] = "computer"
	}
	data["os_type"] = ua.OS()
	name, _ := ua.Browser()
	data["browser"] = name
	data["origin"] = r.Header.Get("Origin")
	return
}

// ArrayObjectToString func used to object to string
func ArrayObjectToString(data []interface{}, key string) string {
	str := ""
	for _, x := range data {
		obj := x.(map[string]interface{})
		if str == "" {
			str = fmt.Sprint(obj[key])
		} else {
			str = str + "," + fmt.Sprint(obj[key])
		}
	}

	return str
}

func DateNowFormat(format string) string {
	dateObj, _ := goment.New(time.Now())
	return dateObj.Format(format)
}

type HtmlToPdfFileGenerate struct {
	HtmlStr  string
	Type     string //base64,directory
	FileName string
}

func HtmlToPdfFile(data HtmlToPdfFileGenerate) (string, error) {
	pdfg, err := wkhtml.NewPDFGenerator()
	pdfg.MarginBottom.Set(1)
	pdfg.MarginTop.Set(1)
	pdfg.MarginLeft.Set(1)
	pdfg.MarginRight.Set(1)
	pdfg.PageSize.Set("A4")
	pdfg.Dpi.Set(72)
	pdfg.PageWidth.Set(210)
	pdfg.PageHeight.Set(297)
	// pdfg.Dpi.Set(339)
	if err != nil {
		logs.Logs("(bookingFileObjct) Generate Object error", err)
		return "", err
	}
	page := wkhtml.NewPageReader(strings.NewReader(data.HtmlStr))
	pdfg.AddPage(page)
	page.DisableSmartShrinking.Set(true)
	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		logs.Logs("(bookingFileObjct) Create PDF document in internal buffer error", err)
		return "", err
	}
	//Your Pdf Name
	err = pdfg.WriteFile("./" + data.FileName)
	if err != nil {
		logs.Logs("(bookingFileObjct) Your Pdf Name error", err)
		return "", err
	}
	// Create base64 format
	if data.Type == "base64" {
		file, err := os.Open(data.FileName)
		if err != nil {
			logs.Logs("(bookingFileObjct) file open error", err)
			return "", err
		}
		b, err := ioutil.ReadAll(file)
		if err != nil {
			logs.Logs("(bookingFileObjct) file read error", err)
			return "", err
		}
		encoded := base64.StdEncoding.EncodeToString(b)
		return encoded, nil
	}
	return "", nil
}

func SortValue(data []interface{}, key string) []interface{} {
	sort.Slice(data, func(i, j int) bool {
		return data[i].(map[string]interface{})[key].(string) < data[j].(map[string]interface{})[key].(string)
	})
	return data
}
