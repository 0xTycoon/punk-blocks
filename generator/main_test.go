package main

import (
	"bytes"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"github.com/miguelmota/go-solidity-sha3"
	"image"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestThis(t *testing.T) {
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"string"},

		// values
		[]interface{}{
			"Hello",
		},
	)

	fmt.Println(hex.EncodeToString(hash))
	//dumpBlocks()
}

type block struct {
	i     image.Image
	trait string
}

func TestGenerator(t *testing.T) {
	var allBlocks blocks
	var err error
	var blocksPath = "./traits-24x24.png"
	if _, err = allBlocks.load(blocksPath); err != nil {
		fmt.Println(err)
		return
	}
	csvReader := csv.NewReader(bytes.NewReader([]byte(params)))
	records, err := csvReader.ReadAll()
	m := make(map[int][]block)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	r1.Intn(10000)
	for i := 0; i < 133; i++ {
		cat, _ := strconv.Atoi(records[i][3])
		id, _ := strconv.Atoi(records[i][0])
		m[cat] = append(m[cat], block{
			allBlocks.getPunkBlock(id),
			records[i][2]},
		)
	}

	for {

	}

	/*

		0. in solidity, keep a structure pre-configured as above.

		1. get random number for index

		2. using random starting index, roll a category 0. Keep rolling until one is found

		3. repeat for remaining attys

	*/
}
