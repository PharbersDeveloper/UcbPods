package UcbHandler

import (
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
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

func (h UcbCallRHandler) CallRCalculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	params := map[string]string{}
	res, _ := ioutil.ReadAll(r.Body)
	result := map[string]interface{}{}
	enc := json.NewEncoder(w)
	json.Unmarshal(res, &params)

	tc := time.After(time.Second * 2)

	proposalId, pok := params["proposal-id"]
	accountId, aok := params["account-id"]

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


	// 拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	resource := fmt.Sprint(h.Args[0], "/", h.Args[1], "/",proposalId, "/", accountId)
	mergeURL := strings.Join([]string{scheme, resource}, "")

	fmt.Println(mergeURL)

	// 转发
	client := &http.Client{}
	req, _ := http.NewRequest("GET", mergeURL, nil)
	for k, v := range r.Header {
		req.Header.Add(k, v[0])
	}
	response, err := client.Do(req)

	if err != nil {
		return 1
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return 1
	}

	rCalcResultBody := map[string]string{}
	json.Unmarshal(body, &rCalcResultBody)

	resultBody, sok := rCalcResultBody["status"]



	if sok && resultBody == "Success" {
		result["status"] = "Success"
		result["msg"] = "计算成功"
		<-tc
		enc.Encode(result)
	} else {
		//result["status"] = "Error"
		//result["msg"] = "计算失败"
		//enc.Encode(result)
		<-tc
		panic("计算失败")
	}
	return 0
}

func (h UcbCallRHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h UcbCallRHandler) GetHandlerMethod() string {
	return h.Method
}
