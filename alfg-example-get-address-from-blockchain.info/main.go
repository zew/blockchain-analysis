package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/alfg/blockchain" // client for data API https://blockchain.info/api/blockchain_api
	"github.com/zew/util"
)

func pf(args ...interface{}) {
	format, ok := args[0].(string)
	if ok && len(args) > 1 {
		args = args[1:]
		fmt.Printf(format+"\n", args...)
		return
	}
	if ok && len(args) == 1 {
		fmt.Printf("%v\n", args...)
		return
	}
	fmt.Print("\t")
	fmt.Print(args...)
	fmt.Print("\n")
}

func main() {

	c, e := blockchain.New()
	resp, e := c.GetAddress("162FjqU7RYdojnejCDe6zrPDUpaLcv9Hhq")
	if e != nil {
		fmt.Print(e)
	}

	maxLen := 60
	dmp := util.IndentedDump(resp)
	dmpSplit := strings.Split(dmp, "\n")
	for _, line := range dmpSplit {
		if len(line) > maxLen {
			line = line[:maxLen] + "..."
		}
		// line = strings.Replace(line, " ", "_", -1)
		if strings.HasPrefix(line, " "+strings.Repeat("\t", 4)) {
			continue
		}
		fmt.Println(line)
	}
	// fmt.Print(dmp)

	pf("Address-Hash160: %v - %v", resp.Hash160, resp.Address)
	pf("%v Txs -  Tot sent: %v   Tot recv: %v  - Final balance %v", resp.NTx, resp.TotalSent, resp.TotalReceived, resp.FinalBalance)
	fmt.Println("\tTransactions: ")
	fmt.Println("\t=========================")
	for i := range resp.Txs {
		for _, inp := range resp.Txs[i].Inputs {
			pf("\t\tInput  %v  %10.2v - spent: %v - Seq %v",
				inp.PrevOut.N,
				inp.PrevOut.Value,
				inp.PrevOut.Spent,
				inp.Sequence,
			)
		}
		for j := range resp.Txs[i].Out {
			pf("\t\tOutput %v  %10.2v - spent: %v",
				resp.Txs[i].Out[j].N,
				resp.Txs[i].Out[j].Value,
				resp.Txs[i].Out[j].Spent,
				// resp.Txs[i].Inputs[j].PrevOut.Spent,
			)
		}
		t := time.Unix(int64(resp.Txs[i].Time), int64(0))
		pf("\t%v %11v \n", t.Format(time.RFC3339Nano)[:16], resp.Txs[i].Result)
	}

}
