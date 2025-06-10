package sbp

import (
	"espp-mock/configs"
	"espp-mock/espp"
)

type SBPUsecase struct {
	conf *configs.Server

	ESPP *espp.IPSAdapder
}

func New(conf *configs.Server, espp *espp.IPSAdapder) *SBPUsecase {
	return &SBPUsecase{
		conf: conf,
		ESPP: espp,
	}
}

