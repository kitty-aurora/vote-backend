package storage

import (
	"context"
	"log"

	"sync"

	"vote-backend/voting"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client          *ethclient.Client
	contract        *voting.Voting
	auth            *bind.TransactOpts
	initOnce        sync.Once
	contractAddress = "0xYourContractAddressHere" // TODO: 部署后替换成真实合约地址
)

// 初始化合约连接
func initContract() {
	var err error
	client, err = ethclient.Dial("https://sepolia.infura.io/v3/YOUR_INFURA_KEY") // TODO: 换成自己的节点
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// 这里简化，假设你已经有一个 account 的 auth
	// 实际上你需要用 keystore / private key 来构造 auth
	// auth, err = bind.NewTransactorWithChainID(...)
	// 省略，后续你可以补全
	// contract, err = voting.NewVoting(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}
}

// 获取候选人
func GetCandidates() []map[string]interface{} {
	initOnce.Do(initContract)

	var result []map[string]interface{}

	// 调用 getAllCandidates
	names, votes, err := contract.GetAllCandidates(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		log.Fatal(err)
	}

	for i, name := range names {
		result = append(result, map[string]interface{}{
			"name":  name,
			"votes": votes[i].Int64(),
		})
	}
	return result
}

// 投票
func Vote(name string) (map[string]interface{}, bool) {
	initOnce.Do(initContract)

	tx, err := contract.Vote(auth, name)
	if err != nil {
		log.Printf("Failed to vote: %v", err)
		return nil, false
	}

	// 返回交易哈希
	return map[string]interface{}{
		"txHash": tx.Hash().Hex(),
	}, true
}

// 重置投票
func ResetVotes() {
	//initOnce.Do(initContract)
	//
	//_, err := contract.Reset(auth)
	//if err != nil {
	//	log.Printf("Failed to reset votes: %v", err)
	//}
}
