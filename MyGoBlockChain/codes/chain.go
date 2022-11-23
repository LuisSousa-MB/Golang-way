package codes

var blockGenesis Block

type Chain struct {
	Instance []Block
}

func (c *Chain) GerateGenesisBlock() {
	blockGenesis = Block{
		Index:        1,
		Nonce:        0,
		PreviousHash: "",
		Hash:         "0000000000GeNeSiS000000000000000",
		Data:         Data{},
	}
	if c.VerifyGenesisBlock(len(c.Instance)) {
		c.Instance = append(c.Instance, blockGenesis)
	}

}
func (c *Chain) VerifyGenesisBlock(len int) bool {
	if len == 0 {
		//fmt.Println("True")
		return true
	} else {
		//fmt.Println("False")

		return false
	}
}
