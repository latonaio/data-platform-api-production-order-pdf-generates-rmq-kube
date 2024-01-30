package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-production-order-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
	dpfm_api_output_formatter "data-platform-api-production-order-pdf-generates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-production-order-pdf-generates-rmq-kube/config"
	"encoding/json"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	lnpdf "github.com/latonaio/golang-pdf-library"
	"io/ioutil"
	"os"
	"sync"
)

func (c *DPFMAPICaller) process(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header

	for _, fn := range accepter {
		switch fn {
		case "ProductionOrder":
			func() {
				header, _ = c.GeneratePDF(input, errs, log, conf)
			}()
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header: *header,
	}

	return data
}

func (c *DPFMAPICaller) GeneratePDF(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) (*[]dpfm_api_output_formatter.Header, error) {
	var data []dpfm_api_output_formatter.Header

	for _, headerData := range input.Header {
		var err error

		outputPath := fmt.Sprintf(
			"%s/%s%s",
			conf.MountPath,
			fmt.Sprintf("%d-%d", headerData.ProductionOrder, headerData.ProductionOrderItem),
			".pdf",
		)

		dataJsonFilePath := fmt.Sprintf("%s/%s", conf.MountPath, "data.json")
		dataJsonFile, err := os.Create(
			dataJsonFilePath,
		)
		if err != nil {
			return nil, err
		}

		err = copyPDF(
			fmt.Sprintf("%s/%s", conf.MaterialPath, "layout.pdf"),
			outputPath,
		)
		if err != nil {
			return nil, err
		}

		// TODO 後でDBから取得することになりそう
		templateFileName := fmt.Sprintf("%s/%s", conf.MaterialPath, "template.json")

		pdfData := *dpfm_api_output_formatter.SetToPdfData(
			&headerData,
			&input.Header[0],
			input.Items,
		)

		pdfDataJsonBytes, err := json.Marshal(pdfData)
		if err != nil {
			return nil, err
		}

		_, err = dataJsonFile.WriteString(string(pdfDataJsonBytes))
		if err != nil {
			return nil, err
		}

		lnpdf.Builder{
			TemplatePath:   templateFileName,
			DataSourcePath: dataJsonFilePath,
		}.Build().Output(&outputPath)

		data = append(
			data,
			pdfData,
		)
	}

	return &data, nil
}

func copyPDF(srcPath, destPath string) error {
	srcFile, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destPath, srcFile, 0644)
	if err != nil {
		return err
	}

	return nil
}
