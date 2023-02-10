package above_threshold

import (
	"context"

	"wallet-service"
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
	var wil []wallet.WalletInfo
	if v := ctx.Value(); v != nil {
		wil = *(v.(*[]wallet.WalletInfo))
	}
	}

	lastWalletInfo := wil[len(wil)-1]

	if !(lastWalletInfo.AboveThreshold) {
		dr := msg.(*wallet.DepositRequest)

		start_idx := 0
		var expiredCumulative float64 = 0.0

		for i, wi := range wil {
			if (dr.CreatedAt - wi.UpdatedAt) < twoMinute {
				start_idx = i
				break
			}
			expiredCumulative += wi.LastDepositAmount
		}

		newTwoMinuteCumulative := lastWalletInfo.TwoMinuteCumulative + dr.Amount - expiredCumulative

		newWalletInfo := wallet.WalletInfo{
			WalletID:            dr.WalletID,
			LastDepositAmount:   dr.Amount,
			TwoMinuteCumulative: newTwoMinuteCumulative,
			AboveThreshold:      newTwoMinuteCumulative <= twoMinuteDepositLimit,
			UpdatedAt:           dr.CreatedAt,
		}

		wil = wil[start_idx:]
		wil = append(wil, newWalletInfo)

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
