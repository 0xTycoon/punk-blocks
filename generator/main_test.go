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
	cat  int
}

func getDistributionData() map[string][]int {
	distribution := make(map[string][]int)
	csvReader := csv.NewReader(bytes.NewReader([]byte(distributionData)))
	records, _ := csvReader.ReadAll()
	for i := range records {
		var counts []int
		for n := 0; n < 8; n++ {
			v, _ := strconv.Atoi(records[i][n+5])
			counts = append(counts, v)
		}
		distribution[records[i][1]] = counts
	}
	return distribution
}

var distribution map[string][]int

func TestGenerator(t *testing.T) {
	var allBlocks blocks
	var err error
	var blocksPath = "./traits-24x24.png"
	if _, err = allBlocks.load(blocksPath); err != nil {
		fmt.Println(err)
		return
	}
	distribution = getDistributionData()
	csvReader := csv.NewReader(bytes.NewReader([]byte(params)))
	records, err := csvReader.ReadAll()
	male := make(map[int][]block)
	female := make(map[int][]block)
	attributes := make(map[int]block)
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
				cat,
			})
		} else if records[i][2] == "m" {
			male[cat] = append(male[cat], block{
				allBlocks.getPunkBlock(id),
				frq,
				"m",
				records[i][1],
				cat,
			})
			attributes[id] = block{
				allBlocks.getPunkBlock(id),
				frq,
				"m",
				records[i][1],
				cat,
			}
		} else {
			female[cat] = append(female[cat], block{
				allBlocks.getPunkBlock(id),
				frq,
				"f",
				records[i][1],
				cat,
			})
			attributes[id] = block{
				allBlocks.getPunkBlock(id),
				frq,
				"f",
				records[i][1],
				cat,
			}
		}

	}
	sum := 0
	total := 0
	for i := 0; i < 100; i++ {
		counts, zbcount := pickCollection(bases, female, male, attributes)
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
	attributes map[int]block,
) (map[int]int, int) {
	var chosenAtt []block
	counts := make(map[int]int)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	zbcount := 0
	for i := 0; i < 10000; i++ {
		chosenBase := pickBase(base, r1)
		count := pickCount(r1)

		chosenAtt = pickPunkToAttributes(attributes, r1, count, chosenBase)

		for moo := range chosenAtt {
			attrCounts[chosenAtt[moo].name]++
		}
		if chosenBase.name == "7" { // || chosenBase.name == "Ape" || chosenBase.name == "Alien"
			for a := range chosenAtt {
				if chosenAtt[a].name == "Beanie" {
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
	return false
	if base == "Alien" && cat == 3 {
		hairExcluded++
		return true // hair excluded from alien
	}
	return false
}

var attrCounts = make(map[string]int)

func pickPunkToAttributes(
	blocks map[int]block,
	r1 *rand.Rand,
	desiredCount int,
	base block,
) []block {
	var chosenAtt []block
	if desiredCount == 0 {
		return chosenAtt
	}
	catPicks := make(map[int]bool)
	rolls := 0
	for {
		i := r1.Intn(122) + 11
		pick := blocks[i]
		if pick.sex != base.sex {
			continue
		}

		if _, ok := catPicks[pick.cat]; ok {
			rolls++
			continue
		}
		n2 := r1.Intn(10000)
		if punkProbability(pick, base) >= n2 {
			chosenAtt = append(chosenAtt, pick)
			catPicks[pick.cat] = true // remember that this category was picked
		}
		rolls++
		if len(chosenAtt) == desiredCount {
			return chosenAtt
		}
	}
}

func punkProbability(pick block, base block) int {
	ret := 0
	i := 0
	if base.sex == "f" {
		i += 4
	}
	n, _ := strconv.Atoi(base.name)
	i += n - 1
	if v, ok := distribution[pick.name]; ok {
		return v[i]
	}
	return ret
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
	return chosenBase
}
