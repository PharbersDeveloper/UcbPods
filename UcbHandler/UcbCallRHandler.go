package UcbHandler

import (
	"Ucb/UcbModel"
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

	var paper []UcbModel.Paper
	var paperInput UcbModel.Paperinput
	var scenario UcbModel.Scenario
	var businessInputs []UcbModel.Businessinput
	//var destConfig []UcbModel.DestConfig
	//var resourceConfig []UcbModel.ResourceConfig
	//var goodsInputs []UcbModel.Goodsinput


	req.QueryParams["account-id"] = []string{accountId}
	req.QueryParams["proposal-id"] = []string{proposalId}

	// 查询当前Paper
	_ = h.db.FindMulti(req, &UcbModel.Paper{}, &paper, -1, -1)
	currentPaperInputID := paper[0].InputIDs[len(paper[0].InputIDs)-1]

	// 查询当前PaperInput
	_ = h.db.FindOne(&UcbModel.Paperinput{ID: currentPaperInputID}, &paperInput)

	// 查询当前Scenario
	_ = h.db.FindOne(&UcbModel.Scenario{ID: paperInput.ScenarioID}, &scenario)

	// 查询当前BusinessInput
	req.QueryParams["ids"] = paperInput.BusinessinputIDs
	_ = h.db.FindMulti(req, &UcbModel.Businessinput{}, &businessInputs, -1, -1)



	// 查询当前Dest、Resource
	//var (
	//	destIds []string
	//	resources []string
	//)
	//for _, business := range businessInputs {
	//	destIds = append(destIds, business.DestConfigId)
	//	resources = append(resources, business.ResourceConfigId)
	//}
	//req.QueryParams["ids"] = destIds
	//_ = h.db.FindMulti(req, &UcbModel.DestConfig{}, &destConfig, -1, -1)
	//req.QueryParams["ids"] = resources
	//_ = h.db.FindMulti(req, &UcbModel.ResourceConfig{}, &resourceConfig, -1, -1)








	fmt.Println(paper)

	return 0
}

func (h UcbCallRHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h UcbCallRHandler) GetHandlerMethod() string {
	return h.Method
}

func getApi2goRequest(r *http.Request, header http.Header) api2go.Request{
	return api2go.Request{
		PlainRequest: r,
		Header: header,
		QueryParams: map[string][]string{},
	}
}
