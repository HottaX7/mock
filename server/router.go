package server

import (
	"espp-mock/server/sbp"
	"net/http"

	//	"github.com/go-chi/chi"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
)

type RouterBuilder struct {
	SBPHandler *sbp.Handler
}

func (b *RouterBuilder) Build() *chi.Mux {
	r := chi.NewRouter()

	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/{someID}/A01", b.SBPHandler.HandleA01)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/{transactionNumber}/C01", b.SBPHandler.HandleC01)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/C04", b.SBPHandler.HandleC04)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/C05", b.SBPHandler.HandleC05)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/C11", b.SBPHandler.HandleC11)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/C13", b.SBPHandler.HandleC13)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/C24", b.SBPHandler.HandleC24)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/{transactionNumber}/{callID}", b.SBPHandler.Log)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/{callID}", b.SBPHandler.Log)
	r.With(middleware.Logger).HandleFunc("/payment/v1/universal-payment-link/paymentdata/{id}", b.SBPHandler.HandlePaymentData)
	//	r.With(middleware.Logger).HandleFunc("/v01/C2BQRD/120000000020/B05", b.SBPHandler.HandleB05)
	r.With(middleware.Logger).HandleFunc("/api/v1/C2BQRD/120000000020/B05", b.SBPHandler.HandleB05)
	r.With(middleware.Logger).HandleFunc("/v1/C2BQRD/120000000020/B21", b.SBPHandler.HandleB21)
	//	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/120000000020/B05", b.SBPHandler.HandleB05)
	r.Post("/send-payment", b.SBPHandler.HandleSendPayment)

	return r
}
func SetupRouter(handler *sbp.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/payment/v1", func(r chi.Router) {
		r.Get("/universal-payment-link/paymentdata/{id}", handler.HandlePaymentData)
	})
	return r
}
