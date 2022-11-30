package convert_complementer

import (
	"context"
	dpfm_api_input_reader "convert-to-dpfm-payment-requisition-mitsubishi-from-dpfm/DPFM_API_Input_Reader"
	"convert-to-dpfm-payment-requisition-mitsubishi-from-dpfm/config"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type ConvertComplementer struct {
	ctx context.Context
	c   *config.Conf
	rmq *rabbitmq.RabbitmqClient
}

func NewConvertComplementer(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient) *ConvertComplementer {
	return &ConvertComplementer{
		ctx: ctx,
		c:   c,
		rmq: rmq,
	}
}

func (c *ConvertComplementer) ComplementHeader(data *dpfm_api_input_reader.SDC, l *logger.Logger) error {

	return nil
}
