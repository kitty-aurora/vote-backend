package storage

import (
	"context"
	"log"
	"math/big"
	"sync"

	"vote-backend/vote" // abigen 生成的包

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client          *ethclient.Client
	contract        *vote.Voting
	auth            *bind.TransactOpts
	initOnce        sync.Once
	contractAddress = "0xYourContractAddressHere"
)

func initClient() {
	var err error
	client, err = ethclient.Dial("https://goerli.infura.io/v3/YOUR_INFURA_KEY")
	if err != nil {
		log.Fatal(err)
	}

	contract, err = vote.NewVoting(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	// auth 初始化：这里需要你导入钱包私钥
	// auth, _ = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
}

func GetCandidates() []vote.Candidate {
	initOnce.Do(initClient)

	candidates := []vote.Candidate{}
	for i := int64(1); i <= int64(getCandidateCount()); i++ {
		c, err := contract.GetCandidate(nil, big.NewInt(i))
		if err != nil {
			log.Println(err)
			continue
		}
		candidates = append(candidates, c)
	}
	return candidates
}

func Vote(name string) (vote.Candidate, bool) {
	initOnce.Do(initClient)

	// 查找候选人 ID
	var candidateID *big.Int
	for i := int64(1); i <= int64(getCandidateCount()); i++ {
		c, err := contract.GetCandidate(nil, big.NewInt(i))
		if err != nil {
			continue
		}
		if c.Name == name {
			candidateID = c.Id
			break
		}
	}

	if candidateID == nil {
		return vote.Candidate{}, false
	}

	// 调用链上 vote 函数（需要 auth 发送交易）
	tx, err := contract.Vote(auth, candidateID)
	if err != nil {
		log.Println(err)
		return vote.Candidate{}, false
	}
	log.Println("tx hash:", tx.Hash().Hex())

	// 返回最新数据
	c, _ := contract.GetCandidate(nil, candidateID)
	return c, true
}

func ResetVotes() {
	// 链上通常不支持 reset，除非你设计了管理员函数
	log.Println("ResetVotes not supported on blockchain")
}

func getCandidateCount() int64 {
	count, err := contract.CandidateCount(nil)
	if err != nil {
		log.Fatal(err)
	}
	return count.Int64()
}
