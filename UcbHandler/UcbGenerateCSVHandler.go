package UcbHandler

import (
	"Ucb/UcbDataStorage"
	"Ucb/Util/uuid"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

type UcbGenerateCSVHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
}

func (h UcbGenerateCSVHandler) NewGenerateCSVHandler(args ...interface{}) UcbGenerateCSVHandler {
	var m *BmMongodb.BmMongodb
	var r *BmRedis.BmRedis
	var hm string
	var md string
	var ag []string
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmMongodb" {
					m = dm.(*BmMongodb.BmMongodb)
				}
				if tm.Name() == "BmRedis" {
					r = dm.(*BmRedis.BmRedis)
				}
			}
		} else if i == 1 {
			md = arg.(string)
		} else if i == 2 {
			hm = arg.(string)
		} else if i == 3 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {
		}
	}

	return UcbGenerateCSVHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r}
}

func (h UcbGenerateCSVHandler) GenerateCSV(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	req := getApi2goRequest(r, w.Header())
	params := map[string]string{}
	res, _ := ioutil.ReadAll(r.Body)
	result := map[string]interface{}{}
	enc := json.NewEncoder(w)
	json.Unmarshal(res, &params)


	//_, err := UcbMiddleware.UcbCheckToken.CheckTokenFormFunction(w, r)
	//if err != nil {
	//	panic(fmt.Sprintf(err.Error()))
	//}

	proposalId, pok := params["proposal-id"]
	accountId, aok := params["account-id"]
	scenarioId, sok := params["scenario-id"]
	downloadType, dok := params["download-type"]

	if !pok {
		result["status"] = "error"
		result["msg"] = "生成失败，proposal-id参数缺失"
		enc.Encode(result)
		return 1
	}

	if !aok {
		result["status"] = "error"
		result["msg"] = "生成失败，account-id参数缺失"
		enc.Encode(result)
		return 1
	}

	if !sok {
		result["status"] = "error"
		result["msg"] = "生成失败，scenario-id参数缺失"
		enc.Encode(result)
		return 1
	}

	if !dok {
		result["status"] = "error"
		result["msg"] = "生成失败，download-type参数缺失"
		enc.Encode(result)
		return 1
	}

	var (
		header []string
		body [][]string
		resultMap map[string]interface{}
	)

	if downloadType == "business" {
		resultMap = h.businessOut(proposalId, accountId, scenarioId, req)
	} else if downloadType == "assessment" {
		header, body = h.assessmentOut(proposalId, accountId, scenarioId)
	}

	fmt.Println(header)
	fmt.Println(body)
	fmt.Println(resultMap)

	businessReport := resultMap["report"].(map[string]interface{})
	businessReportHeader := businessReport["header"].([]string)
	businessReportBody := businessReport["body"].([][]string)



	//content := []string{"Alex", "18510971868"}
	//header = []string{"姓名", "电话"}
	//body = [][]string{content}
	//
	var uid uuid.UUID
	uid, _ = uuid.NewRandom()
	//inputFileName := fmt.Sprint(uid.String(), ".csv")
	uid, _ = uuid.NewRandom()
	reportFileName := fmt.Sprint(uid.String(), ".csv")
	//
	//_ = generateCsvFile(inputFileName, header, body)
	_ = generateCsvFile(reportFileName, businessReportHeader, businessReportBody)
	//
	//fileNames := []string{inputFileName, reportFileName}
	//
	//result["status"] = "ok"
	//result["fileNames"] = fileNames
	//enc.Encode(result)
	return 0
}

func (h UcbGenerateCSVHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h UcbGenerateCSVHandler) GetHandlerMethod() string {
	return h.Method
}

func (h UcbGenerateCSVHandler) businessOut(proposalId, accountId, scenarioId string, req api2go.Request) map[string]interface{} {
	mdb := []BmDaemons.BmDaemon{h.db}
	scenarioStorage := UcbDataStorage.UcbScenarioStorage{}.NewScenarioStorage(mdb)
	paperStorage := UcbDataStorage.UcbPaperStorage{}.NewPaperStorage(mdb)
	//paperInputStorage := UcbDataStorage.UcbPaperinputStorage{}.NewPaperinputStorage(mdb)
	//businessInputStorage := UcbDataStorage.UcbBusinessinputStorage{}.NewBusinessinputStorage(mdb)
	destConfigStorage := UcbDataStorage.UcbDestConfigStorage{}.NewDestConfigStorage(mdb)
	hospitalConfigStorage := UcbDataStorage.UcbHospitalConfigStorage{}.NewHospitalConfigStorage(mdb)
	cityStorage := UcbDataStorage.UcbCityStorage{}.NewCityStorage(mdb)
	hospitalStorage := UcbDataStorage.UcbHospitalStorage{}.NewHospitalStorage(mdb)
	resourceConfigStorage := UcbDataStorage.UcbResourceConfigStorage{}.NewResourceConfigStorage(mdb)
	representativeConfigStorage := UcbDataStorage.UcbRepresentativeConfigStorage{}.NewRepresentativeConfigStorage(mdb)
	representativeStorage := UcbDataStorage.UcbRepresentativeStorage{}.NewRepresentativeStorage(mdb)
	//goodsInputStorage := UcbDataStorage.UcbGoodsinputStorage{}.NewGoodsinputStorage(mdb)
	goodsConfigStorage := UcbDataStorage.UcbGoodsConfigStorage{}.NewGoodsConfigStorage(mdb)
	productConfigStorage := UcbDataStorage.UcbProductConfigStorage{}.NewProductConfigStorage(mdb)
	productStorage := UcbDataStorage.UcbProductStorage{}.NewProductStorage(mdb)

	salesReportStorage := UcbDataStorage.UcbSalesReportStorage{}.NewSalesReportStorage(mdb)
	hospitalSalesReportStorage := UcbDataStorage.UcbHospitalSalesReportStorage{}.NewHospitalSalesReportStorage(mdb)

	req.QueryParams["account-id"] = []string{accountId}
	req.QueryParams["proposal-id"] = []string{proposalId}
	req.QueryParams["orderby"] = []string{"time"}

	//inputHeader := []string {"时间", "城市名称", "医院名称", "医院等级", "负责代表", "产品", "进药状态", "患者数量", "指标", "预算"}
	reportHeader := []string {"时间", "城市名称", "医院名称", "医院等级", "负责代表", "产品", "进药状态", "患者数量", "指标达成率", "销量"}

	var reportBody [][]string

	// 最新的paper
	papers := paperStorage.GetAll(req, -1, -1)
	paper := papers[len(papers) - 1]

	req.QueryParams["ids"] = paper.SalesReportIDs
	salesReports := salesReportStorage.GetAll(req, -1, -1)

	for _, salesReport := range salesReports {
		scenario, _ := scenarioStorage.GetOne(salesReport.ScenarioID)

		req.QueryParams["ids"] = salesReport.HospitalSalesReportIDs
		req.QueryParams["notEq[destConfigId]"] = []string{"-1"}
		hospitalSalesReports := hospitalSalesReportStorage.GetAll(req, -1, -1)
		if len(salesReport.PaperInputID) > 0 {

		} else {
			var reportContent []string
			for _, hospitalSalesReport :=  range hospitalSalesReports {
				destConfig, _ := destConfigStorage.GetOne(hospitalSalesReport.DestConfigID)
				hospitalConfig, _ := hospitalConfigStorage.GetOne(destConfig.DestID)
				city, _ := cityStorage.GetOne(hospitalConfig.CityID)
				hospital, _ := hospitalStorage.GetOne(hospitalConfig.HospitalID)
				resourceConfig, _ := resourceConfigStorage.GetOne(hospitalSalesReport.ResourceConfigID)
				repConfig, _ := representativeConfigStorage.GetOne(resourceConfig.ResourceID)
				rep, _ := representativeStorage.GetOne(repConfig.RepresentativeID)
				goodsConfig, _ := goodsConfigStorage.GetOne(hospitalSalesReport.GoodsConfigID)
				productConfig, _ := productConfigStorage.GetOne(goodsConfig.GoodsID)
				product, _ := productStorage.GetOne(productConfig.ProductID)
				reportContent = append(reportContent, scenario.Name, city.Name, hospital.Name, hospital.HospitalLevel,
					rep.Name, product.Name, hospitalSalesReport.DrugEntranceInfo, strconv.Itoa(hospitalSalesReport.PatientCount),
					strconv.FormatFloat(hospitalSalesReport.QuotaAchievement, 'f', -1, 32), strconv.FormatFloat(hospitalSalesReport.Sales,'f', -1, 32))
			}
			reportBody = append(reportBody, reportContent)
		}
	}


	//inputIds := paper.InputIDs[0:]
	//
	//req.QueryParams["ids"] = inputIds

	//for _, paperInput := range paperInputStorage.GetAll(req, -1, -1) {
	//	var paperInputs []map[string]interface{}
	//
	//	scenario, _ := scenarioStorage.GetOne(paperInput.ScenarioID)
	//	header = append(header, fmt.Sprint(scenario.Name, "指标"))
	//	header = append(header, fmt.Sprint(scenario.Name, "预算"))
	//
	//	req.QueryParams["ids"] = paperInput.BusinessinputIDs
	//	businessInputs := businessInputStorage.GetAll(req, -1,-1)
	//
	//	cleanQueryParams(&req)
	//
	//
	//
	//	//req.QueryParams["paperInput-id"] = []string{paperInput.ID}
	//	//salesReport := salesReportStorage.GetAll(req, -1,-1)[0]
	//	//
	//	//cleanQueryParams(&req)
	//	//req.QueryParams["ids"] = salesReport.HospitalSalesReportIDs
	//	//
	//	//hospitalSalesReports := hospitalSalesReportStorage.GetAll(req, -1, -1)
	//
	//	for _, businessInput := range businessInputs {
	//		for _, hospitalSalesReport := range hospitalSalesReports {
	//			if businessInput.DestConfigId == hospitalSalesReport.DestConfigID {
	//				cleanQueryParams(&req)
	//
	//				destConfig, _ := destConfigStorage.GetOne(businessInput.DestConfigId)
	//				hospitalConfig, _ := hospitalConfigStorage.GetOne(destConfig.DestID)
	//				city, _ := cityStorage.GetOne(hospitalConfig.CityID)
	//				hospital, _ := hospitalStorage.GetOne(hospitalConfig.HospitalID)
	//				resourceConfig, _ := resourceConfigStorage.GetOne(businessInput.ResourceConfigId)
	//				repConfig, _ := representativeConfigStorage.GetOne(resourceConfig.ResourceID)
	//				rep, _ := representativeStorage.GetOne(repConfig.RepresentativeID)
	//
	//				req.QueryParams["ids"] = businessInput.GoodsInputIds
	//
	//				for _, goodsInput := range goodsInputStorage.GetAll(req, -1, -1) {
	//					goodsConfig, _ := goodsConfigStorage.GetOne(goodsInput.GoodsConfigId)
	//					productConfig, _ := productConfigStorage.GetOne(goodsConfig.GoodsID)
	//					product, _ := productStorage.GetOne(productConfig.ProductID)
	//
	//					paperInputs = append(paperInputs, map[string]interface{}{
	//						"cityName": city.Name,
	//						"hospitalName": hospital.Name,
	//						"hospitalLevel": hospital.HospitalLevel,
	//						"reqName":	rep.Name,
	//						"productName": product.Name,
	//						"drugEntranceInfo": hospitalSalesReport.DrugEntranceInfo,
	//						"patientCount": hospitalSalesReport.PatientCount,
	//					})
	//
	//				}
	//
	//
	//			}
	//		}
	//
	//		//destIds = append(destIds, businessInput.ID)
	//	}
	//
	//	//req.QueryParams["ids"] = destIds
	//
	//
	//
	//}



	return map[string]interface{}{
		"report": map[string]interface{}{
			"header": reportHeader,
			"body": reportBody,
		},
	}

}

func (h UcbGenerateCSVHandler) assessmentOut(proposalId, accountId, scenarioId string) ([]string, [][]string) {
	var (
		header []string
		body [][]string
	)

	return header, body
}

func generateCsvFile (fileName string, header []string, body [][]string) error {
	data := [][]string{
		header,
	}
	for _, v := range body {
		data = append(data, v)
	}

	path := fmt.Sprint("./files/", fileName)

	newFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		newFile.Close()
	}()
	newFile.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码
	w := csv.NewWriter(newFile)
	err = w.WriteAll(data)
	w.Flush()
	return err
}