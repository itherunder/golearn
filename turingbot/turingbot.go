package turingbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yezihack/colorlog"
)

const (
	TuringUrl       = "http://openapi.turingapi.com/openapi/api/v2"
	TuringApi       = "e72a4a68d85949f6830775b71648da07"
	TuringUserId    = "coldplay"
	PostContentType = "application/json;charset=utf-8"
)

type InputText_ struct {
	Text string `json:"text"`
}

type InputImage_ struct {
	Url string `json:"url"`
}

type Location_ struct {
	City     string `json:"city"`
	Province string `json:"province"`
	Street   string `json:"street"`
}

type SelfInfo_ struct {
	Location Location_ `json:"location"`
}

type Preception_ struct {
	InputText InputText_ `json:"inputText"`
	// InputImage InputImage_ `json:"inputImage"`
	// SelfInfo   SelfInfo_   `json:"selfInfo"`
}

type UserInfo_ struct {
	ApiKey string `json:"apiKey"`
	UserId string `json:"userId"`
}

// http://openapi.turingapi.com/openapi/api/v2
type Request struct {
	ReqType    int         `json:"reqType"`
	Preception Preception_ `json:"preception"`
	UserInfo   UserInfo_   `json:"userInfo"`
}

type Parameters_ struct {
	NearbyPlace string `json:"nearby_place"`
}

type Intent_ struct {
	Code       int         `json:"code"`
	IntentName string      `json:"intentName"`
	ActionName string      `json:"actionName"`
	Parameters Parameters_ `json:"parameters"`
}

type Values_ struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}

type Result struct {
	GroupType  int     `json:"groupType"`
	ResultType string  `json:"resultType"`
	Values     Values_ `json:"values"`
}

type Response struct {
	Intent Intent_ `json:"intent"`
}

func TuringBot() {
	colorlog.Info("Enter `EOF` to shut down:")
	inputChan := make(chan string)
	go Process(inputChan)
	for {
		var input string
		fmt.Scanln(&input)
		if input == "EOF" {
			colorlog.Info("bye")
			break
		}
		inputChan <- input
	}
}

func HandleError(err error) {
	if err != nil {
		colorlog.Error("error happended: %s", err.Error())
		panic("exit for error")
	}
}

// bot process the cmd input
func Process(inputChan <-chan string) {
	for {
		select {
		case input := <-inputChan:
			// colorlog.Info("received from command: %s", input)
			request := Request{
				ReqType:    0,
				Preception: Preception_{InputText: InputText_{Text: input}},
				UserInfo:   UserInfo_{ApiKey: TuringApi, UserId: TuringUserId},
			}
			colorlog.Info("your message: %v", request)
			requestBytes, err := json.Marshal(request)
			HandleError(err)
			res, err := http.Post(TuringUrl, PostContentType, bytes.NewBuffer([]byte(requestBytes)))
			HandleError(err)
			content, err := ioutil.ReadAll(res.Body)
			HandleError(err)
			var response Response
			err = json.Unmarshal(content, &response)
			HandleError(err)
			colorlog.Info("Turing Bot: %v", response)
		}
	}
}
