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
	i    image.Image
	freq int
	sex  string
	name string
}

// 3,
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
	male := make(map[int][]block)
	female := make(map[int][]block)
	base := make(map[int][]block)

	for i := 0; i < 133; i++ {
		cat, _ := strconv.Atoi(records[i][3])
		id, _ := strconv.Atoi(records[i][0])
		frq, _ := strconv.Atoi(records[i][4])

		if i < 11 {
			base[0] = append(base[cat], block{
				allBlocks.getPunkBlock(id),
				frq,
				records[i][2],
				records[i][1],
			})
		} else if records[i][2] == "m" {
			male[cat] = append(male[cat], block{
				allBlocks.getPunkBlock(id),
				frq,
				"m",
				records[i][1],
			})
		} else {
			female[cat] = append(female[cat], block{
				allBlocks.getPunkBlock(id),
				frq,
				"f",
				records[i][1],
			})

		}

	}
	var chosenAtt []block
	counts := make(map[int]int)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 10000; i++ {
		chosenBase := pickBase(base, r1)
		if chosenBase.sex == "f" {
			chosenAtt = pickPunk(female, r1)
		} else {
			chosenAtt = pickPunk(male, r1)
		}
		println(len(chosenAtt))
		if v, ok := counts[len(chosenAtt)]; ok {
			counts[len(chosenAtt)] = v + 1
		} else {
			counts[len(chosenAtt)] = 1
		}

	}

	fmt.Println("counts are:", counts)

	/*

		0. in solidity, keep a structure pre-configured as above.

		1. get random number for index

		2. using random starting index, roll a category 0. Keep rolling until one is found

		3. repeat for remaining attys

	*/
}

func pickPunk(set map[int][]block, r1 *rand.Rand) []block {

	var chosenAtt []block

	n2 := r1.Intn(10000)
	i := (n2 % len(set)) + 1
	catRolls := 0
	for {
		if _, ok := set[i]; !ok {
			fmt.Println("zeroooo!")
			i++
			continue
		}
		j := n2 % len(set[i])
		j = r1.Intn(len(set[i]))
		attRolls := 0
		catRolls++

		fmt.Println("cat roll:", catRolls, ", i:", i)
		for {
			n2 = r1.Intn(10000)
			fmt.Println("attRoll, ", attRolls, "cat len", len(set[i]), " n:", set[i][j].name, " ", n2)

			if set[i][j].freq >= n2 {
				chosenAtt = append(chosenAtt, set[i][j])
				fmt.Println("picked!", set[i][j])
				attRolls = 0
				j = 0
				break
			}
			j++
			attRolls++
			if attRolls >= len(set[i]) {
				break
			}

			if j >= len(set[i]) {
				j = 0
			}
		}
		if catRolls >= len(set) {
			fmt.Println("max catRolls, ", catRolls)
			catRolls = 0
			attRolls = 0
			break
		}
		i++

		if i > 12 {
			i = 1 // wrap around the categories
		}
	}
	return chosenAtt
}

func pickBase(base map[int][]block, r1 *rand.Rand) block {
	var chosenBase block
	n := r1.Intn(10000)
	i := n % len(base[0]) // random index (i) to start from

	fmt.Println("i is:", i)
	for {
		n2 := r1.Intn(10000)
		if base[0][i].freq >= n2 {
			chosenBase = base[0][i]
			break
		}

		i++
		if i >= len(base[0]) {
			i = 0
		}
	}
	fmt.Println("chosen base:", chosenBase)
	return chosenBase
}
