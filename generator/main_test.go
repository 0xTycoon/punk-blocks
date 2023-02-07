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

func TestTrial(t *testing.T) {

	picked := 0
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	i := r1.Intn(8)
	for {
		n2 := r1.Intn(10000)
		if 696 >= n2 {
			picked++
		}
		i++
		if i == 10000 {
			break
		}
	}
	fmt.Println(picked)

}

type block struct {
	i    image.Image
	freq int
	sex  string
	name string
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
	male := make(map[int][]block)
	female := make(map[int][]block)
	bases := make(map[int][]block)

	for i := 0; i < 133; i++ {
		cat, _ := strconv.Atoi(records[i][3])
		id, _ := strconv.Atoi(records[i][0])
		frq, _ := strconv.Atoi(records[i][4])

		if i < 11 {
			bases[0] = append(bases[cat], block{
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
	sum := 0
	total := 0
	for i := 0; i < 1; i++ {
		counts, zbcount := pickCollection(bases, female, male)
		if zbcount > 0 {
			total += 1
		}

		fmt.Println("Alien beanie cnt", zbcount)
		fmt.Println("counts are:", counts)
		fmt.Println(attrCounts)
		for k, _ := range attrCounts {
			sum += attrCounts[k]
		}
		attrCounts = make(map[string]int)
		//25,984
	}
	fmt.Println("Alien beanie total:", total)
	fmt.Println("hairExcluded:", hairExcluded)
	fmt.Println("all:", sum)
	sum = 0
	/*

		0. in solidity, keep a structure pre-configured as above.

		1. get random number for index

		2. using random starting index, roll a category 0. Keep rolling until one is found

		3. repeat for remaining attys

	*/
}

func pickCollection(
	base map[int][]block,
	female map[int][]block,
	male map[int][]block,
) (map[int]int, int) {
	var chosenAtt []block
	counts := make(map[int]int)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	zbcount := 0
	for i := 0; i < 10000; i++ {
		chosenBase := pickBase(base, r1)
		count := pickCount(r1)
		if chosenBase.sex == "f" {
			attrCounts["m"]++
			chosenAtt = pickPunkToAttributes2(female, r1, count, chosenBase.name)
		} else {
			attrCounts["f"]++
			chosenAtt = pickPunkToAttributes2(male, r1, count, chosenBase.name)
		}
		for moo := range chosenAtt {
			attrCounts[chosenAtt[moo].name]++
		}
		if chosenBase.name == "Alien" { // || chosenBase.name == "Ape" || chosenBase.name == "Alien"
			for a := range chosenAtt {
				if chosenAtt[a].name == "Cowboy Hat" {
					zbcount++
				}
			}
		}
		if v, ok := counts[len(chosenAtt)]; ok {
			counts[len(chosenAtt)] = v + 1
		} else {
			counts[len(chosenAtt)] = 1
		}

	}
	return counts, zbcount
}

var hairExcluded int

func isExcluded(base string, cat int) bool {
	if base == "Alien" && cat == 3 {
		hairExcluded++
		return true // hair excluded from alien
	}
	return false
}

var attrCounts = make(map[string]int)

func pickPunkToAttributes2(
	set map[int][]block,
	r1 *rand.Rand,
	desiredCount int,
	base string) []block {
	var chosenAtt []block
	if desiredCount == 0 {
		return chosenAtt
	}
	for {
		i := r1.Intn(len(set)) + 1
		if _, ok := set[i]; !ok {
			//fmt.Println("zeroooo!") // eg. f has no facial hair
			continue
		}
		j := r1.Intn(len(set[i]))
		n2 := r1.Intn(10000)
		if set[i][j].freq >= n2 {
			chosenAtt = append(chosenAtt, set[i][j])
			if _, ok := attrCounts[set[i][j].name+strconv.Itoa(set[i][j].freq)]; ok {
				//	attrCounts[set[i][j].name+strconv.Itoa(set[i][j].freq)]++
			} else {
				//	attrCounts[set[i][j].name+strconv.Itoa(set[i][j].freq)] = 1
			}
		}
		if len(chosenAtt) == desiredCount {
			return chosenAtt
		}
	}
	//return chosenAtt
}

func pickPunkToAttributes(
	set map[int][]block,
	r1 *rand.Rand,
	desiredCount int,
	base string) []block {

	var chosenAtt []block
	if desiredCount == 0 {
		return chosenAtt
	}
	n2 := r1.Intn(10000)

	i := r1.Intn(len(set)) + 1
	catRolls := 0
	for {
		if _, ok := set[i]; !ok {
			//fmt.Println("zeroooo!")
			i++
			continue
		}

		j := r1.Intn(len(set[i]))
		attRolls := 0
		catRolls++

		//fmt.Println("cat roll:", catRolls, ", i:", i)
		for {
			n2 = r1.Intn(10000)
			//fmt.Println("attRoll, ", attRolls, "cat len", len(set[i]), " n:", set[i][j].name, " ", n2)

			if set[i][j].freq >= n2 && !isExcluded(base, i) {
				chosenAtt = append(chosenAtt, set[i][j])
				if _, ok := attrCounts[set[i][j].name+strconv.Itoa(set[i][j].freq)]; ok {
					attrCounts[set[i][j].name+strconv.Itoa(set[i][j].freq)]++
				} else {
					attrCounts[set[i][j].name+strconv.Itoa(set[i][j].freq)] = 1
				}
				//fmt.Println("picked!", set[i][j])
				if len(chosenAtt) == desiredCount {
					return chosenAtt
				}

				attRolls = 0
				j = 0
				break // finish rolling this category
			}
			j++
			attRolls++
			if attRolls >= len(set[i]) {
				break
				//i = 1
			}

			if j >= len(set[i]) {
				j = 0
			}
		}
		i++
		if i > 12 {
			i = 1 // wrap around the categories
		}
	}
}

var counts = [8]int{8, 333, 3560, 4501, 1420, 166, 11, 1}

func pickCount(r1 *rand.Rand) int {
	i := r1.Intn(8)
	for {
		n2 := r1.Intn(10000)
		if counts[i] >= n2 {
			return i
		}
		i++
		if i == len(counts) {
			i = 0
		}
	}
}

func pickBase(base map[int][]block, r1 *rand.Rand) block {
	var chosenBase block
	i := r1.Intn(len(base[0])) // random index (i) to start from
	//fmt.Println("i is:", i)
	for {
		n2 := r1.Intn(10000)
		if base[0][i].freq >= n2 {
			chosenBase = base[0][i]
			break
		}
		i++
		if i == len(base[0]) {
			i = 0
		}
	}
	//fmt.Println("chosen base:", chosenBase)
	return chosenBase
}
