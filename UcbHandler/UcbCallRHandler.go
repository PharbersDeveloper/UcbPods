package UcbHandler

import (
	"Ucb/UcbDataStorage"
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"io/ioutil"
	"net/http"
	"reflect"
)

type UcbCallRHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
}

type calcStruct struct {
	Header     map[string]string      `json:"header"`
	Account    string                 `json:"account"`
	Proposal   string                 `json:"proposal"`
	PaperInput string                 `json:"paperInput"`
	Scenario   map[string]interface{} `json:"scenario"`
	Body       map[string]interface{} `json:"body"`
}

func (h UcbCallRHandler) NewCallRHandler(args ...interface{}) UcbCallRHandler {
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

	return UcbCallRHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r }
}

//rr api2go.Request
func (h UcbCallRHandler) CallRCalculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	mdb := []BmDaemons.BmDaemon{h.db}
	w.Header().Add("Content-Type", "application/json")
	req := getApi2goRequest(r, w.Header())
	params := map[string]string{}
	res, _ := ioutil.ReadAll(r.Body)
	result := map[string]interface{}{}
	enc := json.NewEncoder(w)
	json.Unmarshal(res, &params)

	proposalId, pok := params["proposal-id"]
	accountId, aok := params["account-id"]
	scenarioId, sok := params["scenario-id"]

	fmt.Println(proposalId)
	fmt.Println(accountId)
	fmt.Println(scenarioId)

	if !pok {
		result["status"] = "error"
		result["msg"] = "计算失败，proposal-id参数缺失"
		enc.Encode(result)
		return 1
	}

	if !aok {
		result["status"] = "error"
		result["msg"] = "计算失败，account-id参数缺失"
		enc.Encode(result)
		return 1
	}

	if !sok {
		result["status"] = "error"
		result["msg"] = "计算失败，scenario-id参数缺失"
		enc.Encode(result)
		return 1
	}

	req.QueryParams["account-id"] = []string{accountId}
	req.QueryParams["proposal-id"] = []string{proposalId}
	req.QueryParams["orderby"] = []string{"time"}

	var (
		inputs []map[string]interface{}
		histories []map[string]interface{}
	)

	scenarioStorage := UcbDataStorage.UcbScenarioStorage{}.NewScenarioStorage(mdb)
	paperStorage := UcbDataStorage.UcbPaperStorage{}.NewPaperStorage(mdb)
	paperInputStorage := UcbDataStorage.UcbPaperinputStorage{}.NewPaperinputStorage(mdb)
	businessInputStorage := UcbDataStorage.UcbBusinessinputStorage{}.NewBusinessinputStorage(mdb)
	destConfigStorage := UcbDataStorage.UcbDestConfigStorage{}.NewDestConfigStorage(mdb)
	hospitalConfigStorage := UcbDataStorage.UcbHospitalConfigStorage{}.NewHospitalConfigStorage(mdb)
	cityStorage := UcbDataStorage.UcbCityStorage{}.NewCityStorage(mdb)
	hospitalStorage := UcbDataStorage.UcbHospitalStorage{}.NewHospitalStorage(mdb)
	resourceConfigStorage := UcbDataStorage.UcbResourceConfigStorage{}.NewResourceConfigStorage(mdb)
	representativeConfigStorage := UcbDataStorage.UcbRepresentativeConfigStorage{}.NewRepresentativeConfigStorage(mdb)
	representativeStorage := UcbDataStorage.UcbRepresentativeStorage{}.NewRepresentativeStorage(mdb)
	goodsInputStorage := UcbDataStorage.UcbGoodsinputStorage{}.NewGoodsinputStorage(mdb)
	goodsConfigStorage := UcbDataStorage.UcbGoodsConfigStorage{}.NewGoodsConfigStorage(mdb)
	productConfigStorage := UcbDataStorage.UcbProductConfigStorage{}.NewProductConfigStorage(mdb)
	productStorage := UcbDataStorage.UcbProductStorage{}.NewProductStorage(mdb)
	salesConfigStorage := UcbDataStorage.UcbSalesConfigStorage{}.NewSalesConfigStorage(mdb)

	salesReportStorage := UcbDataStorage.UcbSalesReportStorage{}.NewSalesReportStorage(mdb)
	hospitalSalesReport := UcbDataStorage.UcbHospitalSalesReportStorage{}.NewHospitalSalesReportStorage(mdb)

	// 最新的paper
	paper := paperStorage.GetAll(req, -1, -1)[0]

	// 最新的输入
	paperInput, _ := paperInputStorage.GetOne(paper.InputIDs[len(paper.InputIDs) - 1])

	cleanQueryParams(&req)

	// 最新的BusinessInputs Inputs
	req.QueryParams["ids"] = paperInput.BusinessinputIDs
	businessInputs := businessInputStorage.GetAll(req, -1,-1)
	for _, businessInput := range businessInputs {
		var products []interface{}
		product := map[string]interface{}{}
		hospitalMap := map[string]interface{}{}
		representativeMap := map[string]interface{}{}

		destConfig, _ := destConfigStorage.GetOne(businessInput.DestConfigId)
		hospitalConfig, _ := hospitalConfigStorage.GetOne(destConfig.DestID)
		city, _ := cityStorage.GetOne(hospitalConfig.CityID)
		hospital, _ := hospitalStorage.GetOne(hospitalConfig.HospitalID)
		resourceConfig, _ := resourceConfigStorage.GetOne(businessInput.ResourceConfigId)
		representativeConfig, _ := representativeConfigStorage.GetOne(resourceConfig.ResourceID)
		representative, _ := representativeStorage.GetOne(representativeConfig.RepresentativeID)

		cleanQueryParams(&req)
		req.QueryParams["ids"] = businessInput.GoodsInputIds
		goodsInputs := goodsInputStorage.GetAll(req, -1, -1)
		for _, goodsInput := range goodsInputs {
			goodsConfig, _ := goodsConfigStorage.GetOne(goodsInput.GoodsConfigId)
			productConfig, _ := productConfigStorage.GetOne(goodsConfig.GoodsID)
			productModel, _ := productStorage.GetOne(productConfig.ProductID)
			cleanQueryParams(&req)
			req.QueryParams["scenario-id"] = []string{paperInput.ScenarioID}
			req.QueryParams["dest-config-id"] = []string{destConfig.ID}
			req.QueryParams["goods-config-id"] = []string{goodsConfig.ID}

			salesConfig := salesConfigStorage.GetAll(req, -1, -1)[0]

			product["product-id"] = goodsInput.GoodsConfigId
			product["product-name"] = productModel.Name
			product["sales-target"] = goodsInput.SalesTarget
			product["budget"] = goodsInput.Budget
			product["potential"] = salesConfig.Potential
			product["patient-count"] = salesConfig.PatientCount
			products = append(products, product)
		}

		representativeMap["representative-id"] = resourceConfig.ID
		representativeMap["representative-name"] = representative.Name

		hospitalMap["city-id"] = city.ID
		hospitalMap["city-name"] = city.Name
		hospitalMap["hospital-id"] = destConfig.ID
		hospitalMap["hospital-name"] = hospital.Name
		hospitalMap["hospital-level"] = hospital.HospitalLevel
		hospitalMap["representative"] = representativeMap
		hospitalMap["products"] = products
		hospitalMap["meeting-places"] = businessInput.MeetingPlaces
		hospitalMap["visit-time"] = businessInput.VisitTime

		inputs = append(inputs, hospitalMap)
	}

	// 上4周期的医院销售报告
	cleanQueryParams(&req)
	req.QueryParams["ids"] = paper.SalesReportIDs[len(paper.SalesReportIDs) - 4:]
	for _, salesReport := range salesReportStorage.GetAll(req, -1, -1) {
		hospitalMap := map[string]interface{}{}
		scenarioMap := map[string]interface{}{}

		scenario, _ := scenarioStorage.GetOne(salesReport.ScenarioID)
		scenarioMap["name"] = scenario.Name
		scenarioMap["phase"] = scenario.Phase

		req.QueryParams["ids"] = salesReport.HospitalSalesReportIDs
		for _, hospitalSalesReport := range hospitalSalesReport.GetAll(req, -1, -1) {
			destConfig, _ := destConfigStorage.GetOne(hospitalSalesReport.DestConfigID)
			hospitalConfig, _ := hospitalConfigStorage.GetOne(destConfig.DestID)
			hospital, _ := hospitalStorage.GetOne(hospitalConfig.HospitalID)

			resourceConfig, _ := resourceConfigStorage.GetOne(hospitalSalesReport.ResourceConfigID)
			representativeConfig, _ := representativeConfigStorage.GetOne(resourceConfig.ResourceID)
			representative, _ := representativeStorage.GetOne(representativeConfig.RepresentativeID)

			goodsConfig, _  := goodsConfigStorage.GetOne(hospitalSalesReport.GoodsConfigID)
			productConfig, _ := productConfigStorage.GetOne(goodsConfig.GoodsID)
			product, _ := productStorage.GetOne(productConfig.ProductID)

			hospitalMap["scenario"] = scenarioMap
			hospitalMap["hospital-name"] = hospital.Name
			hospitalMap["representative-name"] = representative.Name
			hospitalMap["product-name"] = product.Name

			hospitalMap["sales"] = hospitalSalesReport.Sales
			hospitalMap["sales-quota"] = hospitalSalesReport.SalesQuota
			hospitalMap["ytd-sales"] = hospitalSalesReport.YTDSales
			hospitalMap["drug-entrance-info"] = hospitalSalesReport.DrugEntranceInfo
			histories = append(histories, hospitalMap)
		}
	}


	header := map[string]string{}
	scenario := map[string]interface{}{}
	body := map[string]interface{}{}

	header["application"] = "ucb"
	header["contentType"] = "json"

	sm, _ := scenarioStorage.GetOne(scenarioId)
	scenario["id"] = sm.ID
	scenario["name"] = sm.Name
	scenario["phase"] = sm.Phase

	body["inputs"] = inputs
	body["histories"] = histories

	cs := &calcStruct {
		Header:     header,
		Account:    accountId,
		Proposal:   proposalId,
		PaperInput: paperInput.ID,
		Scenario:   scenario,
		Body:       body,
	}

	c, _ := json.MarshalIndent(cs, "", " ")
	fmt.Println(string(c))

	return 0
}

func (h UcbCallRHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h UcbCallRHandler) GetHandlerMethod() string {
	return h.Method
}

func cleanQueryParams(r *api2go.Request) {
	r.QueryParams = map[string][]string{}
}

func getApi2goRequest(r *http.Request, header http.Header) api2go.Request{
	return api2go.Request{
		PlainRequest: r,
		Header: header,
		QueryParams: map[string][]string{},
	}
}
