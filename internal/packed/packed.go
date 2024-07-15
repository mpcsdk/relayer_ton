package packed

import (
	"relayer_ton/internal/logic/relayer"
	"relayer_ton/internal/service"
)

func init() {
	service.RegisterRelayer(relayer.NewRelayer())
}
