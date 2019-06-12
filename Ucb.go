package main

import (
	"Ucb/UcbFactory"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmApiResolver"
	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"github.com/alfredyang1986/BmServiceDef/BmPodsDefine"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
)

func main() {
	version := "v0"
	prodEnv := "UCB_HOME"
	fmt.Println("UCB pods archi begins, version =", version)

	fac := UcbFactory.UcbTable{}
	var pod = BmPodsDefine.Pod{Name: "new USB", Factory: fac}
	ntmHome := os.Getenv(prodEnv)
	pod.RegisterSerFromYAML(ntmHome + "/resource/service-def.yaml")

	var bmRouter BmConfig.BmRouterConfig
	bmRouter.GenerateConfig(prodEnv)

	addr := bmRouter.Host + ":" + bmRouter.Port
	fmt.Println("Listening on ", addr)
	api := api2go.NewAPIWithResolver(version, &BmApiResolver.RequestURL{Addr: addr})
	pod.RegisterAllResource(api)

	pod.RegisterAllFunctions(version, api)
	pod.RegisterAllMiddleware(api)

	handler := api.Handler().(*httprouter.Router)
	pod.RegisterPanicHandler(handler)
	err := http.ListenAndServe(":"+bmRouter.Port, handler)
	fmt.Println(err)


}
