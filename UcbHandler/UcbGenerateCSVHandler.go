package UcbHandler

import (
	"Ucb/UcbDaemons/UcbXmpp"
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"Ucb/Util/uuid"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	bmconfig "github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var UcbGenerateCSV UcbGenerateCSVHandler

type UcbGenerateCSVHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
	xmpp	   *UcbXmpp.UcbXmpp
	kafka 	   *bmkafka.BmKafkaConfig
}

type csvDataStruct struct {
	Proposal   string              `json:"proposal-id"`
	Account string                 `json:"account-id"`
	Body       map[string]interface{} `json:"body"`
}

func (h UcbGenerateCSVHandler) NewGenerateCSVHandler(args ...interface{}) UcbGenerateCSVHandler {
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

	bmkafka, _ := bmkafka.GetConfigInstance()
	kafka, _ := GetConfigInstance2()
	UcbGenerateCSV = UcbGenerateCSVHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r, xmpp: x, kafka: bmkafka }

	go func() {
		topic := []string{"UCBDownLoad"}
		fmt.Println(topic)
		kafka.SubscribeTopics(topic, subscriptionGenerateFunc)
	}()
	return UcbGenerateCSV
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

	if !dok {
		result["status"] = "error"
		result["msg"] = "生成失败，download-type参数缺失"
		enc.Encode(result)
		return 1
	}

	if downloadType == "business" {
		go func() {
			h.csvDataOut(proposalId, accountId, req)
		}()
	} else if downloadType == "assessment" {
		go func() {
			h.csvDataOut(proposalId, accountId, req)
		}()
	}

	result["status"] = "ok"
	result["msg"] = "正在生成数据"
	enc.Encode(result)
	return 0
}

func (h UcbGenerateCSVHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h UcbGenerateCSVHandler) GetHandlerMethod() string {
	return h.Method
}

func (h UcbGenerateCSVHandler) csvDataOut(proposalId, accountId string, req api2go.Request) { // map[string]interface{}
	mdb := []BmDaemons.BmDaemon{h.db}
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

	salesReportStorage := UcbDataStorage.UcbSalesReportStorage{}.NewSalesReportStorage(mdb)
	hospitalSalesReportStorage := UcbDataStorage.UcbHospitalSalesReportStorage{}.NewHospitalSalesReportStorage(mdb)

	req.QueryParams["account-id"] = []string{accountId}
	req.QueryParams["proposal-id"] = []string{proposalId}
	req.QueryParams["orderby"] = []string{"time"}

	inputHeader := []string {"时间", "城市名称", "医院名称", "医院等级", "负责代表", "产品", "进药状态", "患者数量", "指标", "预算"}
	reportHeader := []string {"时间", "城市名称", "医院名称", "医院等级", "负责代表", "产品", "进药状态", "患者数量", "指标达成率", "销量"}

	var (
		inputBody [][]string
		reportBody [][]string
	)

	inputBody = [][]string{}

	// 最新的paper
	papers := paperStorage.GetAll(req, -1, -1)
	paper := papers[len(papers) - 1]

	// TODO @Alex前后端需要重新对接
	//paper, _ := paperStorage.GetOne("5d10da63421aa9bc0dce3a81")

	req.QueryParams["ids"] = paper.SalesReportIDs
	salesReports := salesReportStorage.GetAll(req, -1, -1)

	for _, salesReport := range salesReports {
		scenario, _ := scenarioStorage.GetOne(salesReport.ScenarioID)

		req.QueryParams["ids"] = salesReport.HospitalSalesReportIDs
		req.QueryParams["notEq[destConfigId]"] = []string{"-1"}
		hospitalSalesReports := hospitalSalesReportStorage.GetAll(req, -1, -1)
		fmt.Println(len(hospitalSalesReports))
		if len(salesReport.PaperInputID) > 0 {
			paperInput, _ := paperInputStorage.GetOne(salesReport.PaperInputID)

			req.QueryParams = map[string][]string{}

			req.QueryParams["ids"] = paperInput.BusinessinputIDs
			businessInputs := businessInputStorage.GetAll(req, -1, -1)
			for _, businessInput := range businessInputs {
				var (
					city UcbModel.City
					hospital UcbModel.Hospital
					rep UcbModel.Representative
				)
				destConfig, _ := destConfigStorage.GetOne(businessInput.DestConfigId)
				hospitalConfig, _ := hospitalConfigStorage.GetOne(destConfig.DestID)
				city, _ = cityStorage.GetOne(hospitalConfig.CityID)
				hospital, _ = hospitalStorage.GetOne(hospitalConfig.HospitalID)
				resourceConfig, _ := resourceConfigStorage.GetOne(businessInput.ResourceConfigId)
				repConfig, _ := representativeConfigStorage.GetOne(resourceConfig.ResourceID)
				rep, _ = representativeStorage.GetOne(repConfig.RepresentativeID)

				req.QueryParams["ids"] = businessInput.GoodsInputIds
				for _, goodsInput := range goodsInputStorage.GetAll(req, -1, -1) {
					goodsConfig, _ := goodsConfigStorage.GetOne(goodsInput.GoodsConfigId)
					productConfig, _ := productConfigStorage.GetOne(goodsConfig.GoodsID)
					product, _ := productStorage.GetOne(productConfig.ProductID)

					drugEntranceInfo := ""
					patientCount := 0

					for _, hospitalSalesReport := range hospitalSalesReports {

						if hospitalSalesReport.DestConfigID == businessInput.DestConfigId && hospitalSalesReport.GoodsConfigID == goodsInput.GoodsConfigId {
							drugEntranceInfo = hospitalSalesReport.DrugEntranceInfo
							patientCount = hospitalSalesReport.PatientCount
						}
					}
					content := []string{scenario.Name, city.Name, hospital.Name, hospital.HospitalLevel,
						rep.Name, product.Name, drugEntranceInfo,
						strconv.Itoa(patientCount),
						strconv.FormatFloat(goodsInput.Budget, 'f', -1, 32),
						strconv.FormatFloat(goodsInput.SalesTarget,'f', -1, 32)}
					inputBody = append(inputBody, content)

				}
			}
		}
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
			content := []string{scenario.Name, city.Name, hospital.Name, hospital.HospitalLevel,
				rep.Name, product.Name, hospitalSalesReport.DrugEntranceInfo,
				strconv.Itoa(hospitalSalesReport.PatientCount),
				strconv.FormatFloat(hospitalSalesReport.QuotaAchievement, 'f', -1, 32),
				strconv.FormatFloat(hospitalSalesReport.Sales,'f', -1, 32)}
			reportBody = append(reportBody, content)
		}
	}


	res := map[string]interface{}{
		"proposal-id": proposalId,
		"account-id": accountId,
		"body": map[string]interface{}{
			"input": map[string]interface{}{
				"header": inputHeader,
				"body": inputBody,
			},
			"report": map[string]interface{}{
				"header": reportHeader,
				"body": reportBody,
			},
		},
	}

	c, _ := json.Marshal(res)

	fmt.Println(string(c))


	//topic := h.kafka.Topics[0]
	//fmt.Println(topic)
	topic := "UCBDownLoad"
	h.kafka.Produce(&topic, c)
}

func generateCsvFile (fileName string, header []string, body [][]string) error {
	data := [][]string{
		header,
	}
	for _, v := range body {
		data = append(data, v)
	}

	env := os.Getenv("DOWNLOAD")
	path := fmt.Sprint(env, fileName)

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

func subscriptionGenerateFunc(content interface{}) {
	var csvData csvDataStruct

	h := UcbGenerateCSV
	result := map[string]interface{}{}
	result["client-id"] = "5cbe7ab8f4ce4352ecb082a3"
	result["type"] = "download"
	result["time"] = strconv.FormatInt(time.Now().Unix(), 10)

	c := content.([]byte)
	err := json.Unmarshal(c, &csvData)
	result["account-id"] = csvData.Account

	businessInput := csvData.Body["input"].(map[string]interface{})
	businessInputHeader := businessInput["header"].([]interface{})
	businessInputBody := businessInput["body"].(interface{})

	businessReport := csvData.Body["report"].(map[string]interface{})
	businessReportHeader := businessReport["header"].([]interface{})
	businessReportBody := businessReport["body"].(interface{})

	var (
		uid uuid.UUID
		tempInputHeader []string
		tempInputBody [][]string
		tempReportHeader []string
		tempReportBody [][]string
	)
	uid, _ = uuid.NewRandom()
	inputFileName := fmt.Sprint(uid.String(), "_Input", ".csv")
	uid, _ = uuid.NewRandom()
	reportFileName := fmt.Sprint(uid.String(), "_SalesReport", ".csv")

	for _, v := range businessInputHeader {
		tempInputHeader = append(tempInputHeader, v.(string))
	}

	for _, v := range businessReportHeader {
		tempReportHeader = append(tempReportHeader, v.(string))
	}

	for _, v := range businessInputBody.([]interface{}) {
		var body []string
		for _, vv := range v.([]interface{}) {
			body = append(body, vv.(string))
		}
		tempInputBody = append(tempInputBody, body)
	}

	for _, v := range businessReportBody.([]interface{}) {
		var body []string
		for _, vv := range v.([]interface{}) {
			body = append(body, vv.(string))
		}
		tempReportBody = append(tempReportBody, body)
	}

	err = generateCsvFile(inputFileName, tempInputHeader, tempInputBody)
	err = generateCsvFile(reportFileName, tempReportHeader, tempReportBody)

	if err != nil {
		result["status"] = "no"
		result["msg"] = "生成文件失败"
		result["fileNames"] = []string{}
		r, _ := json.Marshal(result)
		_ = h.xmpp.SendGroupMsg(h.Args[2], string(r))
	} else {
		fileNames := []string{fmt.Sprint(h.Args[1], inputFileName), fmt.Sprint(h.Args[1], reportFileName)}
		result["status"] = "ok"
		result["msg"] = "生成文件成功"
		result["fileNames"] = fileNames
		r, _ := json.Marshal(result)
		fmt.Println(result)
		_ = h.xmpp.SendGroupMsg(h.Args[2], string(r))
	}
}


/**
 * 一坨屎，不要看，等老铁把Kafka改成多实例在改回来，现在Consumer在一个项目中不能创建多个
 */
type cfg bmkafka.BmKafkaConfig
var e error
var onceConfig sync.Once
var config *cfg
var consumer *kafka.Consumer
var onceConsumer sync.Once

// GetConsumerInstance get one KafkaConsumerInstance.
func (bkc *cfg) GetConsumerInstance() (*kafka.Consumer, error) {
	onceConsumer.Do(func() {
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": bkc.Broker,
			"broker.address.family":    "v4",
			"group.id":                 bkc.Group,
			"session.timeout.ms":       6000,
			//"auto.offset.reset":        "earliest",
			"auto.offset.reset":        "latest",
			"security.protocol":        "SSL", //默认使用SSL
			"ssl.ca.location":          bkc.CaLocation,
			"ssl.certificate.location": bkc.CaSignedLocation,
			"ssl.key.location":         bkc.SslKeyLocation,
			"ssl.key.password":         bkc.Pass,
		})

		if err != nil {
			fmt.Printf("Failed to create consumer: %s\n", err)
			e = err
		} else {
			fmt.Printf("Created Consumer %v\n", c)
			consumer = c
			e = nil
		}

		//err = c.SubscribeTopics(bkc.Topics, nil)

	})
	return consumer, e
}

func GetConfigInstance2() (*cfg, error) {
	onceConfig.Do(func() {
		configPath := os.Getenv("BM_KAFKA_CONF_HOME")
		profileItems := bmconfig.BMGetConfigMap(configPath)
		topics := make([]string, 0)
		for _, t := range profileItems["Topics"].([]interface{}) {
			topics = append(topics, t.(string))
		}
		config = &cfg {
			Broker:              profileItems["Broker"].(string),
			SchemaRepositoryUrl: profileItems["SchemaRepositoryUrl"].(string),
			Group:               profileItems["Group"].(string),
			CaLocation:          profileItems["CaLocation"].(string),
			CaSignedLocation:    profileItems["CaSignedLocation"].(string),
			SslKeyLocation:      profileItems["SslKeyLocation"].(string),
			Pass:                profileItems["Pass"].(string),
			Topics:              topics,
		}
		e = nil
	})
	return config, e
}


func (bkc *cfg) SubscribeTopics(topics []string, subscribeFunc func(interface{})) {
	if len(bkc.Topics) == 0 {
		panic("no Topics in config")
	}
	c, err := bkc.GetConsumerInstance()
	bmerror.PanicError(err)
	if len(topics) == 0 {
		err = c.SubscribeTopics(bkc.Topics, nil)
	} else {
		err = c.SubscribeTopics(topics, nil)
	}
	bmerror.PanicError(err)

	run := true
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				subscribeFunc(e.Value)
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				continue
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()
}