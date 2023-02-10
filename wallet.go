package wallet

import (
	"encoding/json"

	"github.com/lovoo/goka"
)

var DepositStream goka.Stream = "deposit"

type Wallet struct {
	WalletID string
	Balance  float64
}

type WalletCodec struct{}

func (c *WalletCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *WalletCodec) Decode(data []byte) (interface{}, error) {
	var w Wallet
	err := json.Unmarshal(data, &w)
	return &w, err
}

type WalletInfo struct {
	WalletID            string
	LastDepositAmount   float64
	TwoMinuteCumulative float64
	AboveThreshold      bool
	UpdatedAt           int64
}

type WalletInfoCodec struct{}

func (c *WalletInfoCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *WalletInfoCodec) Decode(data []byte) (interface{}, error) {
	var wi WalletInfo
	err := json.Unmarshal(data, &wi)
	return &wi, err
}

type WalletInfoListCodec struct{}

func (c *WalletInfoListCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *WalletInfoListCodec) Decode(data []byte) (interface{}, error) {
	var wil []WalletInfo
	err := json.Unmarshal(data, &wil)
	return &wil, err
}

type DepositRequest struct {
	WalletID  string
	Amount    float64
	CreatedAt int64
}

type DepositRequestCodec struct{}

func (c *DepositRequestCodec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *DepositRequestCodec) Decode(data []byte) (interface{}, error) {
	var dr DepositRequest
	err := json.Unmarshal(data, &dr)
	return &dr, err
}

type BalanceResponse struct {
	WalletID       string
	Balance        float64
	AboveThreshold bool
}
