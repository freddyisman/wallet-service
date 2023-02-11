package balance

import (
	"context"

	"wallet-service"
	"wallet-service/model/pb_wallet"
	"wallet-service/topic_init"

	"github.com/lovoo/goka"
)

var (
	group goka.Group = "balance"
	Table goka.Table = goka.GroupTable(group)
)

func balance(ctx goka.Context, msg interface{}) {
	w := new(pb_wallet.Wallet)
	dr := msg.(*pb_wallet.DepositRequest)

	if v := ctx.Value(); v != nil {
		w = v.(*pb_wallet.Wallet)
	} else {
		w.WalletID = dr.WalletID
	}

	w.Balance += dr.Amount

	ctx.SetValue(w)
}

func PrepareTopics(brokers []string) {
	topic_init.EnsureTableExists(string(Table), brokers)
	topic_init.EnsureStreamExists(string(wallet.DepositStream), brokers)
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
