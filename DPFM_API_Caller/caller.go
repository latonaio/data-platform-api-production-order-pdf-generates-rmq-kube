package dpfm_api_caller

import (
	"context"
	"crypto/sha256"
	dpfm_api_input_reader "data-platform-api-production-order-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
	dpfm_api_output_formatter "data-platform-api-production-order-pdf-generates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-production-order-pdf-generates-rmq-kube/config"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"os"
	"sync"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:  context.Background(),
		conf: conf,
		rmq:  rmq,
	}
}

func generateFile(
	directoryPath string,
	fileName string,
	fileExtension string,
	docData string,
	log *logger.Logger,
) error {
	dec, err := base64.StdEncoding.DecodeString(docData)
	if err != nil {
		return err
	}

	err = createDirectory(directoryPath)
	if err != nil {
		return err
	}

	fileFullPath := fmt.Sprintf("%s/%s.%s",
		directoryPath,
		fileName,
		fileExtension,
	)

	f, err := os.Create(fileFullPath)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("failed to close file", err)
		}
	}(f)

	if _, err := f.Write(dec); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}

	return nil
}

func createDirectory(directoryPath string) error {
	err := os.MkdirAll(directoryPath, 0777)
	if err != nil {
		return err
	}
	return nil
}

func encodeToStringSha256(s string) string {
	r := sha256.Sum256([]byte(s))
	return hex.EncodeToString(r[:])
}

func (c *DPFMAPICaller) AsyncPDFCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	log *logger.Logger,
	conf *config.Conf,
) (interface{}, []error) {
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	var response interface{}

	response = c.process(nil, &mtx, input, output, accepter, &errs, log, conf)

	return response, nil
}
