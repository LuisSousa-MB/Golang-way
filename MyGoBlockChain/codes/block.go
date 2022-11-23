package codes

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Miners int = int(runtime.NumCPU())

type Block struct {
	Index, Nonce       int
	PreviousHash, Hash string
	TimeStamp          time.Time
	Data               Data
}

func (c *Chain) Minerar() {

	// var wg sync.WaitGroup

	var hashs []string
	unlock := UnlockHash(dificuldade)
	fmt.Println(unlock)
	NewBlock.Nonce = 0
	NewBlock.TimeStamp = time.Now()
	NewBlock.PreviousHash = c.Instance[len(c.Instance)-1].Hash
	numHashs := 0

	NewBlock.Data.Transactions = TxPoll.Transactions
	fmt.Println("Block transactions:", NewBlock.Data.Transactions)
	transactions := NewBlock.Data.Transactions
	var txstrings string
	for i := 0; i < len(transactions); i++ {
		txstrings += string(transactions[i].Sender)
		txstrings += string(transactions[i].Receiver)
		txstrings += string(transactions[i].Msg)
		txstrings += string(transactions[i].TimeStamp)
		txstrings += string(transactions[i].Hash)
		txstrings += fmt.Sprintf("%f", (transactions[i].Value))

	}

	done := make(chan bool)
	hashsChan := make(chan string)
	findedHash := false
	unlockString := strings.Join(unlock, "")
	var mu sync.Mutex

	for mn := 1; mn < Miners+1; mn++ {

		go func(miner int) {

			for i := 0; i < 1; {
				mu.Lock()
				if findedHash == false {

					NewBlock.Nonce++
					passString := strconv.Itoa(NewBlock.Index) + strconv.Itoa(NewBlock.Nonce) + NewBlock.PreviousHash + NewBlock.TimeStamp.Format(time.UnixDate) + txstrings
					NewBlock.Hash = GenerateHashByte(passString)
					fmt.Println("Nonce: ", NewBlock.Nonce)
					testHash := NewBlock.Hash[:len(NewBlock.Hash)-(64-1-dificuldade)]
					numHashs++
					//fmt.Println("tam hashs: ", len(hashs))
					for _, e := range hashs {
						if e == NewBlock.Hash {
							fmt.Println("Repeat hash")
							time.Sleep(time.Second * 10)
						}
					}

					hashs = append(hashs, NewBlock.Hash)
					//fmt.Println("tam hashs: ", len(hashs))
					hashsChan <- NewBlock.Hash
					if testHash == unlockString {
						fmt.Println("Miner", miner, "\nsending passString: ", passString)
						//fmt.Println("Hash encontrado pelo minerador ", miner)
						i++
						findedHash = true
						time.Sleep(time.Second)

						done <- true
					} else {
						findedHash = false
					}
					mu.Unlock()
					runtime.Gosched()

				} else {
					break
				}

			}
		}(mn)

	}

	go func() {
		<-done
		close(hashsChan)
	}()

	for n := range hashsChan {
		fmt.Printf("minerando...\nHash: %s\n", n)

		fmt.Println("Hashs gerados: ", numHashs)
	}

}

func (c *Chain) GerarBlock(dificuldade int) {
	now := time.Now()
	c.GerateGenesisBlock()
	c.Minerar()
	duration := time.Since(now)
	fmt.Println("time elapse:", duration /* , int(duration.Hours()), " hours ", int(duration.Minutes()), " minutes ", int(duration.Seconds()), " seconds." */)
	NewBlock.Index = len(c.Instance) + 1
	//fmt.Println("Pool:", TxPoll)

	TxPoll.Transactions = TxPoll.Transactions[:0]
	c.Instance = append(c.Instance, NewBlock)
	time.Sleep(time.Second * 5)
}
func UnlockHash(dificuldade int) []string {
	var key []string
	findedGO := false
	for i := 0; i < dificuldade; i++ {
		for _, e := range key {
			if e == "Go" {
				findedGO = true
			}
		}
		if findedGO {
			key = append(key, "0")

		} else {
			key = append(key, "Go")

		}
	}
	return key
}
func ShowBlockInfo() {
	var numBlock int
	fmt.Println("Informe o Index do block para vizualizar:")
	fmt.Scan(&numBlock)
	if numBlock <= len(NewChain.Instance) {
		NewChain.PrintBlock(numBlock - 1)
	} else {
		fmt.Println("Bloco não encontrado...")
		time.Sleep(time.Second * 2)
	}
}
