// Package jsonclient fetches data from blockchain.info
package jsonclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Address struct {
	Hash160       string         `json:"hash160"`
	Address       string         `json:"address"`
	NTx           int            `json:"n_tx"`
	TotalReceived int            `json:"total_received"`
	TotalSent     int            `json:"total_sent"`
	FinalBalance  int            `json:"final_balance"`
	Txs           []*Transaction `json:"txs"`
}

type Inputs struct {
	Sequence int      `json:"sequence"`
	Script   string   `json:"script"`
	PrevOut  *PrevOut `json:"prev_out"`
}

type PrevOut struct {
	Spent   bool   `json:"spent"`
	TxIndex int    `json:"tx_index"`
	Type    int    `json:"type"`
	Addr    string `json:"addr"`
	Value   int    `json:"value"`
	N       int    `json:"n"`
	Script  string `json:"script"`
}

type Out struct {
	Spent   bool   `json:"spent"`
	TxIndex int    `json:"tx_index"`
	Type    int    `json:"type"`
	Addr    string `json:"addr"`
	Value   int    `json:"value"`
	N       int    `json:"n"`
	Script  string `json:"script"`
}

type Transaction struct {
	Hash        string    `json:"hash"`
	Ver         int       `json:"ver"`
	VinSz       int       `json:"vin_sz"`
	VoutSz      int       `json:"vout_sz"`
	LockTime    int       `json:"lock_time"`
	Size        int       `json:"size"`
	RelayedBy   string    `json:"relayed_by"`
	BlockHeight int       `json:"block_height"`
	TxIndex     int       `json:"tx_index"`
	Inputs      []*Inputs `json:"inputs"`
	Out         []*Out    `json:"out"`
}

type Transactions struct {
	Transactions []*Transaction `json:"txs"`
}

var (
	API_ROOT = "https://blockchain.info"
)

type Client struct {
	*http.Client
}

func (c *Client) loadResponse(path string, i interface{}, formatJson bool) error {
	full_path := API_ROOT + path
	if formatJson {
		full_path = API_ROOT + path + "?format=json"
	}

	// fmt.Println("querying..." + full_path)
	rsp, e := c.Get(full_path)
	if e != nil {
		return e
	}

	defer rsp.Body.Close()

	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return e
	}
	if rsp.Status[0] != '2' {
		return fmt.Errorf("expected status 2xx, got %s: %s", rsp.Status, string(b))
	}

	return json.Unmarshal(b, &i)
}

func New() (*Client, error) {
	return &Client{Client: &http.Client{}}, nil
}

func (c *Client) GetTransaction(transaction string) (*Transaction, error) {
	rsp := &Transaction{}
	var path = "/rawtx/" + transaction
	e := c.loadResponse(path, rsp, false)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetAddress(address string) (*Address, error) {
	rsp := &Address{}
	var path = "/address/" + address
	e := c.loadResponse(path, rsp, true)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}
