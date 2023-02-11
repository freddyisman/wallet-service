package above_threshold

import (
	"context"

	"wallet-service"
	"wallet-service/model/pb_wallet"
	"wallet-service/topic_init"

	"github.com/lovoo/goka"
)

const (
	twoMinute             = 120
	twoMinuteDepositLimit = 10000
)

var (
	group goka.Group = "above_threshold"
	Table goka.Table = goka.GroupTable(group)
)

func aboveThreshold(ctx goka.Context, msg interface{}) {
	wil := new(pb_wallet.WalletInfoList)
	if v := ctx.Value(); v != nil {
		wil = v.(*pb_wallet.WalletInfoList)
	}

	dr := msg.(*pb_wallet.DepositRequest)

	if len(wil.List) == 0 {
		newWalletInfo := pb_wallet.WalletInfo{
			WalletID:            dr.WalletID,
			LastDepositAmount:   dr.Amount,
			TwoMinuteCumulative: dr.Amount,
			AboveThreshold:      dr.Amount > twoMinuteDepositLimit,
			CreatedAt:           dr.CreatedAt,
		}

		wil.List = append(wil.List, &newWalletInfo)
		ctx.SetValue(wil)
		return
	}

	lastWalletInfo := wil.List[len(wil.List)-1]

	if !(lastWalletInfo.AboveThreshold) {
		start_idx := 0
		var expiredCumulative float64 = 0.0

		for _, wi := range wil.List {
			if (dr.CreatedAt - wi.CreatedAt) < twoMinute {
				break
			}
			expiredCumulative += wi.LastDepositAmount
			start_idx += 1
		}

		newTwoMinuteCumulative := lastWalletInfo.TwoMinuteCumulative + dr.Amount - expiredCumulative
		newWalletInfo := pb_wallet.WalletInfo{
			WalletID:            dr.WalletID,
			LastDepositAmount:   dr.Amount,
			TwoMinuteCumulative: newTwoMinuteCumulative,
			AboveThreshold:      newTwoMinuteCumulative > twoMinuteDepositLimit,
			CreatedAt:           dr.CreatedAt,
		}

		wil.List = append(wil.List[start_idx:], &newWalletInfo)

		ctx.SetValue(wil)
	}
}

func PrepareTopics(brokers []string) {
	topic_init.EnsureTableExists(string(Table), brokers)
	topic_init.EnsureStreamExists(string(wallet.DepositStream), brokers)
}

func Run(ctx context.Context, brokers []string) func() error {
	return func() error {
		g := goka.DefineGroup(group,
			goka.Input(wallet.DepositStream, new(wallet.DepositRequestCodec), aboveThreshold),
			goka.Persist(new(wallet.WalletInfoListCodec)),
		)
		p, err := goka.NewProcessor(brokers, g)
		if err != nil {
			return err
		}
		return p.Run(ctx)
	}
}
