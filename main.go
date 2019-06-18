package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/intel-go/fastjson"
	"github.com/osamingo/jsonrpc"
)

type (
	EchoHandler struct{}
	EchoParams  struct {
		Name string `json:"name"`
	}
	EchoResult struct {
		Message string `json:"message"`
	}

	PositionalHandler struct{}
	PositionalParams  []int
	PositionalResult  struct {
		Message []int `json:"message"`
	}

	StatusHandler struct{}
	JRPCParams  struct {
		SubId string `json:"subscriptionId,omitempty"`
		ProdCode string `json:"productCode,omitempty"`
		PriceCode string `json:"priceCode,omitempty"`
		UserId string `json:"userId,omitempty"`
		Data interface{} `json:"data"`
	}
	StatusRespData struct {
		Date string `json:"retryDate,omitempty"`
	}
	JRPCResult struct {
		Status string         `json:"status,omitempty"`
		Data   StatusRespData `json:"data,omitempty"`
	}

	ModifyHandler struct{}
	RemindHandler struct{}
	RemoveHandler struct{}
	RenewHandler struct{}
	ResumeHandler struct{}
	SuspendHandler struct{}

	AddHandler struct{}
	AddParams  struct {
		SubId string `json:"subscriptionId"`
	}
	AddData struct {
		Date string `json:"retryDate"`
	}
	AddResult  struct {
		Status string         `json:"status"`
		Data   StatusRespData `json:"data"`
	}
)
func (h StatusHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {
	var p JRPCParams
	log.Println("\n status, id - " + string(*jsonrpc.RequestID(c)) + "\n" + string(*params))
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return JRPCResult{
		Status: "active",
		Data:   StatusRespData{time.Now().Add(time.Minute).Format(time.RFC3339)},
	}, nil
}
func (h AddHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {

	var p JRPCParams
	log.Println("\n add, id - " + string(*jsonrpc.RequestID(c)) + "\n" + string(*params))
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return JRPCResult{
		Status: "active",
		Data:   StatusRespData{time.Now().Add(time.Minute).Format(time.RFC3339)},
	}, nil
}
func (h ModifyHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {

	var p JRPCParams
	log.Println("\n modify, id - " + string(*jsonrpc.RequestID(c)) + "\n" + string(*params))
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return JRPCResult{
		Status: "active",
		Data:   StatusRespData{time.Now().Add(time.Minute).Format(time.RFC3339)},
	}, nil
}
func (h RemindHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {

	var p JRPCParams
	log.Println("\n remind, id - " + string(*jsonrpc.RequestID(c)) + "\n" + string(*params))
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return JRPCResult{
		Status: "active",
		Data:   StatusRespData{time.Now().Add(time.Minute).Format(time.RFC3339)},
	}, nil
}
func (h RemoveHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {

	var p JRPCParams
	log.Println("\n remove, id - " + string(*jsonrpc.RequestID(c)) + "\n" + string(*params))
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return JRPCResult{
		Status: "active",
		Data:   StatusRespData{time.Now().Add(time.Minute).Format(time.RFC3339)},
	}, nil
}
func (h RenewHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {

	var p JRPCParams
	log.Println("\n renew, id - " + string(*jsonrpc.RequestID(c)) + "\n" + string(*params))
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return JRPCResult{
		Status: "active",
		Data:   StatusRespData{time.Now().Add(time.Minute).Format(time.RFC3339)},
	}, nil
}
func (h ResumeHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {

	var p JRPCParams
	log.Println("\n resume, id - " + string(*jsonrpc.RequestID(c)) + "\n" + string(*params))
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return JRPCResult{
		Status: "active",
		Data:   StatusRespData{time.Now().Add(time.Minute).Format(time.RFC3339)},
	}, nil
}
func (h SuspendHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {

	var p JRPCParams
	log.Println("\n suspend, id - " + string(*jsonrpc.RequestID(c)) + "\n" + string(*params))
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return JRPCResult{
		Status: "active",
		Data:   StatusRespData{time.Now().Add(time.Minute).Format(time.RFC3339)},
	}, nil
}

func (h EchoHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {

	var p EchoParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return EchoResult{
		Message: "Hello, " + p.Name,
	}, nil
}

func (h PositionalHandler) ServeJSONRPC(c context.Context, params *fastjson.RawMessage) (interface{}, *jsonrpc.Error) {

	var p PositionalParams
	if err := jsonrpc.Unmarshal(params, &p); err != nil {
		return nil, err
	}

	return PositionalResult{
		Message: p,
	}, nil
}

func main() {

	mr := jsonrpc.NewMethodRepository()

	if err := mr.RegisterMethod("resume", ResumeHandler{}, JRPCParams{}, JRPCResult{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("suspend", SuspendHandler{}, JRPCParams{}, JRPCResult{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("modify", ModifyHandler{}, JRPCParams{}, JRPCResult{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("remind", RemindHandler{}, JRPCParams{}, JRPCResult{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("renew", RenewHandler{}, JRPCParams{}, JRPCResult{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("remove", RemoveHandler{}, JRPCParams{}, JRPCResult{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("resume", ResumeHandler{}, JRPCParams{}, JRPCResult{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("status", StatusHandler{}, JRPCParams{}, JRPCResult{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("echo", EchoHandler{}, EchoParams{}, EchoResult{}); err != nil {
		log.Fatalln(err)
	}

	if err := mr.RegisterMethod("add", AddHandler{}, JRPCParams{}, JRPCResult{}); err != nil {
		log.Fatalln(err)
	}

	http.Handle("/jrpc", mr)
	http.HandleFunc("/jrpc/debug", mr.ServeDebug)

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatalln(err)
	}
}