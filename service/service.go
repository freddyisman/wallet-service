package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"wallet-service"
	"wallet-service/above_threshold"
	"wallet-service/balance"

	"wallet-service/model/pb_wallet"

	"github.com/gorilla/mux"
	"github.com/lovoo/goka"
)

func Run(brokers []string, stream goka.Stream) {
	balanceView, err := goka.NewView(brokers, balance.Table, new(wallet.WalletCodec))
	if err != nil {
		panic(err)
	}

	aboveThresholdView, err := goka.NewView(brokers, above_threshold.Table, new(wallet.WalletInfoListCodec))
	if err != nil {
		panic(err)
	}

	go balanceView.Run(context.Background())
	go aboveThresholdView.Run(context.Background())

	emitter, err := goka.NewEmitter(brokers, stream, new(wallet.DepositRequestCodec))
	if err != nil {
		panic(err)
	}
	defer emitter.Finish()

	router := mux.NewRouter()
	router.HandleFunc("/deposit", deposit(emitter, stream)).Methods("POST")
	router.HandleFunc("/details/{wallet_id}", details(balanceView, aboveThresholdView)).Methods("GET")

	log.Printf("Listen port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func deposit(emitter *goka.Emitter, stream goka.Stream) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var dr pb_wallet.DepositRequest

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		err = json.Unmarshal(b, &dr)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		dr.CreatedAt = time.Now().Unix()

		err = emitter.EmitSync(dr.WalletID, &dr)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		log.Printf("Deposit request:\n %v\n", &dr)
		fmt.Fprintf(w, "Deposit request:\n %v\n", &dr)
	}
}

func details(balanceView, aboveThresholdView *goka.View) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		walletID := mux.Vars(r)["wallet_id"]

		balanceVal, _ := balanceView.Get(walletID)
		if balanceVal == nil {
			fmt.Fprintf(w, "%s not found!", walletID)
			return
		}
		walletResult := balanceVal.(*pb_wallet.Wallet)

		aboveThresholdVal, _ := aboveThresholdView.Get(walletID)
		if aboveThresholdVal == nil {
			fmt.Fprintf(w, "%s info not found!", walletID)
			return
		}
		walletInfoList := aboveThresholdVal.(*pb_wallet.WalletInfoList)

		aboveThreshold := walletInfoList.List[len(walletInfoList.List)-1].AboveThreshold

		response := wallet.BalanceResponse{
			WalletID:       walletResult.WalletID,
			Balance:        walletResult.Balance,
			AboveThreshold: aboveThreshold,
		}

		byteResponse, err := json.Marshal(response)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteResponse)
	}
}
