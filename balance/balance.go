package balance

import (
	"context"

	"wallet-service"

	"github.com/lovoo/goka"
)

var (
	group goka.Group = "balance"
	Table goka.Table = goka.GroupTable(group)
)

func balance(ctx goka.Context, msg interface{}) {
	var w wallet.Wallet
	if v := ctx.Value(); v != nil {
		w = v.(wallet.Wallet)
	}

	dr := msg.(*wallet.DepositRequest)
	w.Balance += dr.Amount

	ctx.SetValue(w)
}

func Run(ctx context.Context, brokers []string) func() error {
	return func() error {
		g := goka.DefineGroup(group,
			goka.Input(wallet.DepositStream, new(wallet.DepositRequestCodec), balance),
			goka.Persist(new(wallet.WalletCodec)),
		)
		p, err := goka.NewProcessor(brokers, g)
		if err != nil {
			return err
		}
		return p.Run(ctx)
	}
}
