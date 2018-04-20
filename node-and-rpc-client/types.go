package main

import (
	"github.com/btcsuite/btcd/btcjson"
	"github.com/zew/blockchain-block-rpc-client/jsonclient"
)

type Tx struct {
	Id      int    `db:"id"`
	BlkHash string `db:"blk_hash"`
	// BlkTime time.Time `db:"blk_time"`
	BlkTime int64 `db:"blk_time"`

	Hash    string `db:"hash"`
	TxIndex int    `db:"tx_index"`
	Size    int    `db:"size"`
	// Inputs      []*Inputs `json:"inputs"`
	// Out         []*Out    `json:"out"`
}

func (t *Tx) FromBlockAndTx(blk *btcjson.GetBlockVerboseResult, txv *jsonclient.Transaction) {

}
func FromBlockAndTx(blk *btcjson.GetBlockVerboseResult, txv *jsonclient.Transaction) *Tx {
	t := Tx{}
	t.BlkHash = blk.Hash
	// t.BlkTime = time.Unix(blk.Time, 0).UTC()
	t.BlkTime = blk.Time

	t.Hash = txv.Hash
	t.TxIndex = txv.TxIndex
	t.Size = txv.Size
	return &t
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
