package util

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
)

const Polygon_Node_URL = "https://polygon-mainnet.g.alchemy.com/v2/GUzL9NWX7QSpBVQJsvp15aNKYk2XsI-w"

var (
	client *ethclient.Client
)

func init() {
	dial, err := ethclient.Dial(Polygon_Node_URL)
	if err != nil {
		panic(any(errors.New("Polygon_Node_URL cant exist")))
	}
	client = dial
}
func GetClient() *ethclient.Client {
	return client
}
