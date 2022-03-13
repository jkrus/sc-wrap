package blockchain

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"sc-wrap/config"
)

type (
	// QuorumClientI describes quorum client interface.
	QuorumClientI interface {
		// Service use embedded common interface.
		Service
	}
)

type (
	quorumClient struct {
		pathIPC string
		Client  *ethclient.Client
	}
)

// NewQuorumClient return QuorumClientI interface.
func NewQuorumClient(cfg *config.Config) QuorumClientI {
	c := quorumClient{Client: nil, pathIPC: cfg.PathIPC}
	return &c
}

// Reconnect implement Service interface.
func (c *quorumClient) Reconnect() error {
	return c.Start()
}

// Start implement Service interface.
func (c *quorumClient) Start() error {
	rpcClient, err := rpc.Dial(c.pathIPC)
	qClient := ethclient.NewClient(rpcClient)
	if err != nil {
		log.Fatalf("Failed to connect to Quorum client: %v", err)
	}
	c.Client = qClient
	log.Println("Successfully connected to Quorum client.")

	return nil
}

// Stop implement Service interface.
func (c *quorumClient) Stop() error {
	log.Println("Close connection to Quorum client...")
	c.Client.Close()
	log.Println("Quorum client connection closed.")

	return nil
}
