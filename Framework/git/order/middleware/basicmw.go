package middleware

import (
	"time"

	"context"

	cm "tugas_6/Framework/git/order/common"
	"tugas_6/Framework/git/order/services"

	log "github.com/Sirupsen/logrus"
)

func BasicMiddleware() services.ServiceMiddleware {
	return func(next services.PaymentServices) services.PaymentServices {
		return BasicMiddlewareStruct{next}
	}
}

type BasicMiddlewareStruct struct {
	services.PaymentServices
}

func (mw BasicMiddlewareStruct) TripHandler(ctx context.Context, request cm.TheTrip) cm.TheTripResponse {

	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("TripHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("TripHandler begins")

	return mw.PaymentServices.TripHandler(ctx, request)

}
