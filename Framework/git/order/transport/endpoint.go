package transport

import (
	"context"

	cm "tugas_6/Framework/git/order/common"
	"tugas_6/Framework/git/order/services"

	log "github.com/Sirupsen/logrus"

	"github.com/go-kit/kit/endpoint"
)

func invalidRequest() cm.Message {
	return cm.Message{
		Result: &cm.Result{
			Code:   99,
			Remark: "Invalid Request",
		},
	}
}

func TripEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		if req, ok := request.(cm.TheTrip); ok {
			return svc.TripHandler(ctx, req), nil
		}
		log.WithField("Error", request).Info("Request in in unkwon format")
		return invalidRequest(), nil
	}
}
