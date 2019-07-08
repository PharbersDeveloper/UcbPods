package UcbHandler

import (
	"Ucb/UcbDaemons/UcbXmpp"
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

var UcbCallR UcbCallRHandler

type UcbCallRHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
	xmpp	   *UcbXmpp.UcbXmpp
	kafka 	   *bmkafka.BmKafkaConfig
}

type calcStruct struct {
	Header     map[string]string      `json:"header"`
	Account    string                 `json:"account"`
	Proposal   string                 `json:"proposal"`
	PaperInput string                 `json:"paperInput"`
	CurrentScenario   map[string]interface{} `json:"currentScenario"`
	Scenarios		[]interface{} `json:"scenarios"`
	Body       map[string]interface{} `json:"body"`
}

type resultStruct struct {
	Header     map[string]string      `json:"header"`
	Account    string                 `json:"account"`
	Proposal   string                 `json:"proposal"`
	PaperInput string                 `json:"paperInput"`
	Scenario   string 				  `json:"scenario"`
	Body       map[string]interface{} `json:"body"`
	Error 	   map[string]interface{} `json:"error"`
}

func (h UcbCallRHandler) NewCallRHandler(args ...interface{}) UcbCallRHandler {
	var m *BmMongodb.BmMongodb
	var r *BmRedis.BmRedis
	var x *UcbXmpp.UcbXmpp
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
				if tm.Name() == "UcbXmpp" {
					x = dm.(*UcbXmpp.UcbXmpp)
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

	kafka, _ := bmkafka.GetConfigInstance()
	UcbCallR = UcbCallRHandler{ Method: md, HttpMethod: hm, Args: ag, db: m, rd: r, xmpp: x, kafka: kafka }

	go func() {
		topic := []string{kafka.Topics[len(kafka.Topics) -2]}
		fmt.Println(topic)
		kafka.SubscribeTopics(topic, subscriptionFunc)
	}()

	return UcbCallR
}

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

	var (
		inputs []map[string]interface{}
		histories map[string]interface{}
		hospitals []map[string]interface{}
		goodsConfigMappings []map[string]string
		competitions []map[string]interface{}
	)

	histories = make(map[string]interface{})

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
	productReportStorage := UcbDataStorage.UcbProductSalesReportStorage{}.NewProductSalesReportStorage(mdb)
	hospitalSalesReportStorage := UcbDataStorage.UcbHospitalSalesReportStorage{}.NewHospitalSalesReportStorage(mdb)

	// 当前周期的所有产品
	req.QueryParams["scenario-id"] = []string{scenarioId}

	for _, goodsConfig := range goodsConfigStorage.GetAll(req, -1,-1) {
		productConfig, _ := productConfigStorage.GetOne(goodsConfig.GoodsID)
		product, _ := productStorage.GetOne(productConfig.ProductID)
		goodsConfigMappings = append(goodsConfigMappings, map[string]string{
			goodsConfig.ID: product.Name,
			"treatmentArea": productConfig.TreatmentArea,
		})
	}
	cleanQueryParams(&req)

	req.QueryParams["account-id"] = []string{accountId}
	req.QueryParams["proposal-id"] = []string{proposalId}
	req.QueryParams["orderby"] = []string{"time"}

	// 最新的paper
	papers := paperStorage.GetAll(req, -1, -1)
	paper := papers[len(papers) - 1]

	// 最新的输入
	paperInput, _ := paperInputStorage.GetOne(paper.InputIDs[len(paper.InputIDs) - 1])

	cleanQueryParams(&req)

	// 最新的BusinessInputs Inputs
	req.QueryParams["ids"] = paperInput.BusinessinputIDs
	businessInputs := businessInputStorage.GetAll(req, -1,-1)
	for _, businessInput := range businessInputs {
		var products []interface{}

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
			product := map[string]interface{}{}
			cleanQueryParams(&req)
			req.QueryParams["scenario-id"] = []string{paperInput.ScenarioID}
			req.QueryParams["dest-config-id"] = []string{destConfig.ID}
			req.QueryParams["goods-config-id"] = []string{goodsInput.GoodsConfigId}

			salesConfig := salesConfigStorage.GetAll(req, -1, -1)[0]

			for _, goodsConfigMapping := range goodsConfigMappings {
				if v, ok := goodsConfigMapping[goodsInput.GoodsConfigId]; ok {
					product["product-id"] = goodsInput.GoodsConfigId
					product["product-name"] = v
					product["treatment-area"] = goodsConfigMapping["treatmentArea"]
					product["sales-target"] = goodsInput.SalesTarget
					product["budget"] = goodsInput.Budget
					product["potential"] = salesConfig.Potential
					product["patient-count"] = salesConfig.PatientCount
					products = append(products, product)
				}
			}
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
		inputs = append(inputs, hospitalMap)
	}

	// 上4周期的医院销售报告
	cleanQueryParams(&req)
	req.QueryParams["ids"] = paper.SalesReportIDs[len(paper.SalesReportIDs) - 4:]
	salesReports := salesReportStorage.GetAll(req, -1, -1)
	for _, salesReport := range salesReports {

		var businessInputMapping []map[string]interface{}

		if len(salesReport.PaperInputID) > 0 {
			paperInput, _ := paperInputStorage.GetOne(salesReport.PaperInputID)
			req.QueryParams["ids"] = paperInput.BusinessinputIDs
			businessInputs := businessInputStorage.GetAll(req, -1,-1)

			for _, businessInput := range businessInputs {
				req.QueryParams["ids"] = businessInput.GoodsInputIds
				for _, goodsInput := range goodsInputStorage.GetAll(req, -1, -1) {
					businessInputMapping = append(businessInputMapping, map[string]interface{}{
						"destConfigId": businessInput.DestConfigId,
						"goodsConfigId": goodsInput.GoodsConfigId,
						"resourceConfigId": businessInput.ResourceConfigId,
						"budget": goodsInput.Budget,
					})
				}
			}
		}

		scenarioMap := map[string]interface{}{}

		scenario, _ := scenarioStorage.GetOne(salesReport.ScenarioID)
		scenarioMap["name"] = scenario.Name
		scenarioMap["phase"] = scenario.Phase

		req.QueryParams["ids"] = salesReport.HospitalSalesReportIDs
		for _, hospitalSalesReport := range hospitalSalesReportStorage.GetAll(req, -1, -1) {
			hospitalMap := map[string]interface{}{}
			if hospitalSalesReport.DestConfigID != "-1" {
				destConfig, _ := destConfigStorage.GetOne(hospitalSalesReport.DestConfigID)
				hospitalConfig, _ := hospitalConfigStorage.GetOne(destConfig.DestID)
				hospital, _ := hospitalStorage.GetOne(hospitalConfig.HospitalID)

				resourceConfig, _ := resourceConfigStorage.GetOne(hospitalSalesReport.ResourceConfigID)
				representativeConfig, _ := representativeConfigStorage.GetOne(resourceConfig.ResourceID)
				representative, _ := representativeStorage.GetOne(representativeConfig.RepresentativeID)

				goodsConfig, _  := goodsConfigStorage.GetOne(hospitalSalesReport.GoodsConfigID)
				productConfig, _ := productConfigStorage.GetOne(goodsConfig.GoodsID)
				product, _ := productStorage.GetOne(productConfig.ProductID)


				for _, businessMapGoods := range businessInputMapping {
					if businessMapGoods["destConfigId"].(string) ==  hospitalSalesReport.DestConfigID &&
						businessMapGoods["goodsConfigId"].(string) == hospitalSalesReport.GoodsConfigID &&
						businessMapGoods["resourceConfigId"].(string) == hospitalSalesReport.ResourceConfigID {

						hospitalMap["budget"] = businessMapGoods["budget"]
					}
				}

				hospitalMap["budget"] = 0
				hospitalMap["scenario"] = scenarioMap
				hospitalMap["hospital-name"] = hospital.Name
				hospitalMap["representative-name"] = representative.Name
				hospitalMap["product-name"] = product.Name

				hospitalMap["sales"] = hospitalSalesReport.Sales
				hospitalMap["sales-quota"] = hospitalSalesReport.SalesQuota
				hospitalMap["ytd-sales"] = hospitalSalesReport.YTDSales
				hospitalMap["drug-entrance-info"] = hospitalSalesReport.DrugEntranceInfo
				hospitals = append(hospitals, hospitalMap)
			}
		}

	}

	cleanQueryParams(&req)
	for _, salesReport := range salesReports[len(salesReports) - 1:] {

		req.QueryParams["ids"] = salesReport.ProductSalesReportIDs

		for _, productSalesReport := range productReportStorage.GetAll(req, -1, -1) {
			goodsConfig, _ := goodsConfigStorage.GetOne(productSalesReport.GoodsConfigID)
			productConfig, _ :=productConfigStorage.GetOne(goodsConfig.GoodsID)
			if productConfig.ProductType == 1 {
				product, _ := productStorage.GetOne(productConfig.ProductID)
				competitions = append(competitions, map[string]interface{}{
					"product-id": goodsConfig.ID,
					"product-name": product.Name,
					"treatment-name": productConfig.TreatmentArea,
					"share": productSalesReport.Share,
				})
			}
		}
	}
	histories["hospitals"] = hospitals
	histories["competitions"] = competitions

	header := map[string]string{}
	currentScenario := map[string]interface{}{}
	var scenarios []interface{}
	body := map[string]interface{}{}

	header["application"] = "ucb"
	header["contentType"] = "json"

	sm, _ := scenarioStorage.GetOne(scenarioId)
	currentScenario["id"] = sm.ID
	currentScenario["phase"] = sm.Phase

	body["inputs"] = inputs
	body["histories"] = histories

	// 查询所有的周期
	cleanQueryParams(&req)
	req.QueryParams["proposal-id"] = []string{proposalId}
	for _, v := range scenarioStorage.GetAll(req, -1, -1) {
		if v.Phase > 0 {
			scenarios = append(scenarios, map[string]interface{}{
				"id": v.ID,
				"phase" : v.Phase,
			})
		}
	}

	cs := &calcStruct {
		Header:     header,
		Account:    accountId,
		Proposal:   proposalId,
		PaperInput: paperInput.ID,
		CurrentScenario:   currentScenario,
		Scenarios:		scenarios,
		Body:       body,
	}

	c, _ := json.Marshal(cs)

	fmt.Println(string(c))


	topic := h.kafka.Topics[0]
	fmt.Println(topic)
	h.kafka.Produce(&topic, c)

	result["status"] = "ok"
	result["msg"] = "正在计算"
	enc.Encode(result)
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

func subscriptionFunc(content interface{}) {
	//c := content.([]byte)
	//fmt.Println(string(c))

	h := UcbCallR
	c := content.([]byte)
	fmt.Println(string(c))

	var result resultStruct

	err := json.Unmarshal(c, &result)

	ctx := map[string]string {
		"client-id": "5cbe7ab8f4ce4352ecb082a3",
		"type": "calc",
		"account-id": result.Account,
		"proposal-id": result.Proposal,
		"paperInput-id": result.PaperInput,
		"scenario-id": result.Scenario,
		"time": strconv.FormatInt(time.Now().UnixNano() / 1e6, 10),
	}

	if err != nil ||  result.Error != nil{
		//panic("计算失败")
		ctx["status"] = "no"
		ctx["msg"] = "计算失败"
		r, _ := json.Marshal(ctx)
		fmt.Println(string(r))
		_ = h.xmpp.SendGroupMsg(h.Args[0], string(r))
		return
	}

	if len(c) > 2 {
		var (
			result resultStruct
			hospitalSalesReport UcbModel.HospitalSalesReport
			productSalesReport UcbModel.ProductSalesReport
			representativeSalesReport UcbModel.RepresentativeSalesReport
			citySalesReport UcbModel.CitySalesReport
			scenarioResult UcbModel.ScenarioResult

			hospitalSalesReportIDs []string
			productSalesReportIDs []string
			representativeSalesReportIDs []string
			citySalesReportIDs []string
			scenarioResultIDs []string

			assessmentReportID string
		)


		err := json.Unmarshal(c, &result)
		if err != nil ||  result.Error != nil{
			//panic("计算失败")
			return
		}

		mdb := []BmDaemons.BmDaemon{h.db}

		scenarioStorage := UcbDataStorage.UcbScenarioStorage{}.NewScenarioStorage(mdb)
		hospitalSalesReportStorage := UcbDataStorage.UcbHospitalSalesReportStorage{}.NewHospitalSalesReportStorage(mdb)
		productSalesReportStorage := UcbDataStorage.UcbProductSalesReportStorage{}.NewProductSalesReportStorage(mdb)
		representativeSalesReportStorage := UcbDataStorage.UcbRepresentativeSalesReportStorage{}.NewRepresentativeSalesReportStorage(mdb)
		citySalesReportStorage := UcbDataStorage.UcbCitySalesReportStorage{}.NewCitySalesReportStorage(mdb)
		salesReportStorage := UcbDataStorage.UcbSalesReportStorage{}.NewSalesReportStorage(mdb)
		paperStorage := UcbDataStorage.UcbPaperStorage{}.NewPaperStorage(mdb)

		levelStorage := UcbDataStorage.UcbLevelStorage{}.NewLevelStorage(mdb)
		levelConfigStorage := UcbDataStorage.UcbLevelConfigStorage{}.NewLevelConfigStorage(mdb)

		scenarioResultStorage := UcbDataStorage.UcbScenarioResultStorage{}.NewScenarioResultStorage(mdb)
		simplifyResultStorage := UcbDataStorage.UcbSimplifyResultStorage{}.NewSimplifyResultStorage(mdb)
		assessmentReportStorage := UcbDataStorage.UcbAssessmentReportStorage{}.NewAssessmentReportStorage(mdb)

		req := api2go.Request{
			QueryParams: map[string][]string{},
		}

		req.QueryParams["proposal-id"] = []string{result.Proposal}
		req.QueryParams["account-id"] = []string{result.Account}
		req.QueryParams["orderby"] = []string{"time"}

		papers := paperStorage.GetAll(req, -1, -1)
		if len(papers) > 0 {
			paper := papers[len(papers) - 1]

			body := result.Body

			hospitalSalesReports := body["hospitalSalesReports"].([]interface{})
			productSalesReports := body["productSalesReports"].([]interface{})
			representativeSalesReports := body["representativeSalesReports"].([]interface{})
			citySalesReports := body["citySalesReports"].([]interface{})

			for _, v := range hospitalSalesReports {
				mapstructure.Decode(v, &hospitalSalesReport)
				hospitalSalesReportIDs = append(hospitalSalesReportIDs, hospitalSalesReportStorage.Insert(hospitalSalesReport))
			}

			for _, v := range productSalesReports {
				mapstructure.Decode(v, &productSalesReport)
				productSalesReportIDs = append(productSalesReportIDs, productSalesReportStorage.Insert(productSalesReport))
			}

			for _, v := range representativeSalesReports {
				mapstructure.Decode(v, &representativeSalesReport)
				representativeSalesReportIDs = append(representativeSalesReportIDs, representativeSalesReportStorage.Insert(representativeSalesReport))
			}

			for _, v := range citySalesReports {
				mapstructure.Decode(v, &citySalesReport)
				citySalesReportIDs = append(citySalesReportIDs, citySalesReportStorage.Insert(citySalesReport))
			}

			salesReportID := salesReportStorage.Insert(UcbModel.SalesReport{
				ScenarioID: result.Scenario,
				PaperInputID: result.PaperInput,
				Time: time.Now().UnixNano() / 1e6,
				HospitalSalesReportIDs: hospitalSalesReportIDs,
				ProductSalesReportIDs: productSalesReportIDs,
				RepresentativeSalesReportIDs: representativeSalesReportIDs,
				CitySalesReportIDs: citySalesReportIDs,
			})

			paper.SalesReportIDs = append(paper.SalesReportIDs, salesReportID)

			req.QueryParams["proposal-id"] = []string{result.Proposal}
			scenarios := scenarioStorage.GetAll(req, -1,-1)
			if s := scenarios[len(scenarios)-1]; s.ID == result.Scenario {
				simplifyReport := body["simplifyReport"].(map[string]interface{})
				level := simplifyReport["level"].(float64)
				totalQuotaAchievement := simplifyReport["total-quota-achievement"].(float64)
				scenarioResults := simplifyReport["scenarioResult"].([]interface{})

				req.QueryParams = map[string][]string{}
				req.QueryParams["code"] = []string{strconv.FormatFloat(level, 'f', -1, 32),}
				levelModels := levelStorage.GetAll(req, -1,-1)
				levelModel := levelModels[len(levelModels)-1]

				req.QueryParams["code"] = []string{"6"} // 6 => UCB 测评报告
				req.QueryParams["level-id"] = []string{levelModel.ID}
				levelConfigs := levelConfigStorage.GetAll(req, -1,-1)
				levelConfig := levelConfigs[len(levelModels) -1]

				for _, v := range scenarioResults {
					m := v.(map[string]interface{})
					mapstructure.Decode(m, &scenarioResult)
					scenarioResultIDs = append(scenarioResultIDs, scenarioResultStorage.Insert(scenarioResult))
				}

				simplifyResult := UcbModel.SimplifyResult {
					ScenarioResultsIDs: scenarioResultIDs,
					TotalQuotaAchievement: totalQuotaAchievement,
					LevelConfigID: levelConfig.ID,
				}

				simplifyResultID := simplifyResultStorage.Insert(simplifyResult)

				assessmentReport := UcbModel.AssessmentReport {
					SimplifyResultID: simplifyResultID,
					ScenarioID: result.Scenario,
					Time: time.Now().UnixNano() / 1e6,
					PaperInputID: result.PaperInput,
				}

				assessmentReportID = assessmentReportStorage.Insert(assessmentReport)

			}

			if len(assessmentReportID) > 0 {
				paper.AssessmentReportIDs = append(paper.AssessmentReportIDs, assessmentReportID)
			}


			// TODO: @Alex自己留，这面等重构
			var state int
			for _, scenario := range scenarios {
				if scenario.ID == result.Scenario {
					if paper.TotalPhase == scenario.Phase {
						state = 3
					} else {
						state = 2
					}
				}
			}
			paper.InputState = state
			// TODO: @Alex 时间有问题存在UTC转CST问题，因为服务器的都是UTC，Golang默认也是读UTC，等Bug改完后整体做转换
			paper.EndTime = time.Now().UnixNano() / 1e6

			err = paperStorage.Update(*paper)
			if err != nil {
				panic("更新Paper失败")
			}
			ctx["status"] = "ok"
			ctx["msg"] = "计算成功"

			r, _ := json.Marshal(ctx)
			fmt.Println(string(r))
			_ = h.xmpp.SendGroupMsg(h.Args[0], string(r))
		} else {
			ctx["status"] = "no"
			ctx["msg"] = "计算失败，出现异常！"

			r, _ := json.Marshal(ctx)
			fmt.Println(string(r))
			_ = h.xmpp.SendGroupMsg(h.Args[0], string(r))
		}
	}
}
