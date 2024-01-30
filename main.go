package main

import (
	dpfm_api_caller "data-platform-api-production-order-pdf-generates-rmq-kube/DPFM_API_Caller"
	dpfm_api_input_reader "data-platform-api-production-order-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
	dpfm_api_output_formatter "data-platform-api-production-order-pdf-generates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-production-order-pdf-generates-rmq-kube/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	lnpdf "github.com/latonaio/golang-pdf-library"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

func recovery(l *logger.Logger, err *error) {
	if e := recover(); e != nil {
		*err = fmt.Errorf("error occurred: %w", e)
		l.Error(err)
		return
	}
}

func getSessionID(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), conf.RMQ.SessionControlQueue(), conf.RMQ.QueueToSQL(), 0)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	caller := dpfm_api_caller.NewDPFMAPICaller(conf, rmq)

	for msg := range iter {
		l.Debug("received queue message")
		err = callProcess(rmq, caller, conf, msg)
		if err != nil {
			msg.Fail()
			l.Info("call process error", err)
			continue
		}
		msg.Success()
	}
}

func callProcess(rmq *rabbitmq.RabbitmqClient, caller *dpfm_api_caller.DPFMAPICaller, conf *config.Conf, msg rabbitmq.RabbitmqMessage) (err error) {
	l := logger.NewLogger()
	defer recovery(l, &err)

	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": getSessionID(msg.Data())})
	var input dpfm_api_input_reader.SDC
	var output dpfm_api_output_formatter.SDC

	err = json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		l.Error(err)
		return
	}
	err = json.Unmarshal(msg.Raw(), &output)
	if err != nil {
		l.Error(err)
		return
	}

	var errs []error

	accepter := getAccepter(&input)
	res, errs := caller.AsyncPDFCreates(accepter, &input, &output, l, conf)

	//output.MillSheetPdf = generatePdf(string(input.MillCertData.Data))

	if len(errs) != 0 {
		for _, err := range errs {
			l.Error(err)
		}
		output.APIProcessingResult = getBoolPtr(false)
		output.APIProcessingError = errs[0].Error()
		output.Message = res

		rmq.Send(conf.RMQ.QueueToResponse(), output)
		return errs[0]
	}
	output.APIProcessingResult = getBoolPtr(true)
	output.Message = res

	l.JsonParseOut(output)
	rmq.Send(conf.RMQ.QueueToResponse(), output)

	return nil
}

func getAccepter(input *dpfm_api_input_reader.SDC) []string {
	accepter := input.Accepter
	if len(input.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"ProductionOrder",
		}
	}
	return accepter
}

func getBoolPtr(b bool) *bool {
	return &b
}

func generatePdf(data string) string {
	// pdf
	pdfFile := "layout.pdf"
	//defer os.Remove(pdfFile)

	// resource json data
	dataFile := "test.json"
	//defer os.Remove(dataFile)

	createdDatafile, err := os.Create(dataFile)
	if err != nil {
		panic(err)
	}

	// write resource json data
	//os.WriteFile(dataFile, []byte(data), os.ModeAppend)
	//err = os.WriteFile(file, []byte(data), 0644)
	//if err != nil {
	//	panic(err)
	//}
	_, err = createdDatafile.WriteString(data)
	if err != nil {
		panic(err)
	}

	// build & write pdf
	lnpdf.Builder{
		TemplatePath:   "template.json",
		DataSourcePath: dataFile,
	}.Build().Output(&pdfFile)

	// encode pdf binary to base64
	file, err := os.ReadFile(pdfFile)
	if err != nil {
		panic(err)
	}
	encoded := base64.StdEncoding.EncodeToString(file)

	return encoded
}
