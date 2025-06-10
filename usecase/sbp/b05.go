package sbp

import (
	"context"
	"espp-mock/entity"
	"time"

	"github.com/rs/zerolog/log"
)

func (u *SBPUsecase) B05(ctx context.Context, b05 *entity.B05) error {
	transactionNumber := ctx.Value("transactionNumber").(string)

	go func() {
		time.Sleep(u.conf.CallbackDelay)
		// detach context
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		b06 := entity.NewB06(transactionNumber, b05)

		err := u.ESPP.B06(newCtx, b06)
		if err != nil {
			log.Error().Err(err).Msgf("failed to send B06")
		}
	}()

	go func() {
		time.Sleep(2 * u.conf.CallbackDelay)
		// detach context
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		b21 := entity.NewB21(transactionNumber, b05)

		err := u.ESPP.B21(newCtx, b21)
		if err != nil {
			log.Error().Err(err).Msgf("failed to send B21")
		}
	}()

	return nil
}
