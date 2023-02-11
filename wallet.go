package wallet

import (
	"wallet-service/model/pb_wallet"

	"github.com/lovoo/goka"
	"google.golang.org/protobuf/proto"
)

var DepositStream goka.Stream = "deposit"

type WalletCodec struct{}

func (c *WalletCodec) Encode(value interface{}) ([]byte, error) {
	return proto.Marshal(value.(*pb_wallet.Wallet))
}

func (c *WalletCodec) Decode(data []byte) (interface{}, error) {
	var w pb_wallet.Wallet
	err := proto.Unmarshal(data, &w)
	return &w, err
}

type WalletInfoCodec struct{}

func (c *WalletInfoCodec) Encode(value interface{}) ([]byte, error) {
	return proto.Marshal(value.(*pb_wallet.WalletInfo))
}

func (c *WalletInfoCodec) Decode(data []byte) (interface{}, error) {
	var wi pb_wallet.WalletInfo
	err := proto.Unmarshal(data, &wi)
	return &wi, err
}

type WalletInfoListCodec struct{}

func (c *WalletInfoListCodec) Encode(value interface{}) ([]byte, error) {
	return proto.Marshal(value.(*pb_wallet.WalletInfoList))
}

func (c *WalletInfoListCodec) Decode(data []byte) (interface{}, error) {
	var wil pb_wallet.WalletInfoList
	err := proto.Unmarshal(data, &wil)
	return &wil, err
}

type DepositRequestCodec struct{}

func (c *DepositRequestCodec) Encode(value interface{}) ([]byte, error) {
	return proto.Marshal(value.(*pb_wallet.DepositRequest))
}

func (c *DepositRequestCodec) Decode(data []byte) (interface{}, error) {
	var dr pb_wallet.DepositRequest
	err := proto.Unmarshal(data, &dr)
	return &dr, err
}

type BalanceResponse struct {
	WalletID       string
	Balance        float64
	AboveThreshold bool
}
