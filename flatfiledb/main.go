package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"

	"github.com/btcsuite/btcd/database"
	"github.com/btcsuite/btcd/wire"

	_ "github.com/btcsuite/btcd/database/ffldb"
)

func main() {

	log.SetFlags(log.Lshortfile)

	pth := filepath.Join("..", "bockchain-flatfiledb-files", "blocks_ffldb")
	log.Printf("Opening flatfile database 1 %v", pth)
	// db, err := database.Create("ffldb", pth, wire.MainNet)
	db, err := database.Open("ffldb", pth, wire.MainNet)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Printf("Database opened")

	// Use the View function of the database to perform a managed read-only
	// transaction and fetch the block stored above.
	genesisHashStr := "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
	// genesisHashStr = "00000000002b8cd0faa58444df3ba2a22af2b5838c7e4a5b687444f913a575c2" // has block 70.000
	// genesisHashStr = "0000000000000e07595fca57b37fea8522e95e0f6891779cfd34d7e537524471" // has block 50.000

	var loadedBlockBytes []byte
	// err = db.Update(func(tx database.Tx) error {
	err = db.View(func(tx database.Tx) error {
		genesisHash := chaincfg.MainNetParams.GenesisHash
		genesisHashStr = genesisHash.String()
		blockHash, err := chainhash.NewHashFromStr(genesisHashStr)
		_ = blockHash
		if err != nil {
			log.Fatal(err)
		}

		blockBytes, err := tx.FetchBlock(blockHash)
		if err != nil {
			return err
		}

		// As documented, all data fetched from the database is only
		// valid during a database transaction in order to support
		// zero-copy backends.  Thus, make a copy of the data so it
		// can be used outside of the transaction.
		loadedBlockBytes = make([]byte, len(blockBytes))
		copy(loadedBlockBytes, blockBytes)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Typically at this point, the block could be deserialized via the
	// wire.MsgBlock.Deserialize function or used in its serialized form
	// depending on need.  However, for this example, just display the
	// number of serialized bytes to show it was loaded as expected.
	fmt.Printf("Serialized block size: %d bytes\n", len(loadedBlockBytes))

	bl := wire.MsgBlock{}
	bl.Deserialize(loadedBlockBytes)
	// _, _ := wire.MsgBlock.Deserialize(loadedBlockBytes)
}
