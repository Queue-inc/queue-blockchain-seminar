package main

import (
	"os"
	"queue/seraku-seminar-blockchain/api"
	"queue/seraku-seminar-blockchain/jsonstore"

	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	. "github.com/tendermint/tendermint/libs/log"
)

func main() {
	go api.ApiServer()
	initJSONStore()
}

func initJSONStore() error {
	logger := NewTMLogger(NewSyncWriter(os.Stdout))

	// Create the application
	var app types.Application

	app = jsonstore.NewJSONStoreApplication()

	// Start the listener
	srv, err := server.NewServer("tcp://127.0.0.1:26658", "socket", app)
	if err != nil {
		return err
	}
	srv.SetLogger(logger.With("module", "abci-server"))
	if err := srv.Start(); err != nil {
		return err
	}

	// Wait forever
	cmn.TrapSignal(func() {
		// Cleanup
		srv.Stop()
	})
	return nil
}
