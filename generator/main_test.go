package main

import (
	"bytes"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"github.com/miguelmota/go-solidity-sha3"
	"image"
	"math/rand"
	"sort"
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
	i     image.Image
	id    int
	freq  int
	sex   string
	name  string
	cat   int
	stats []int
	stat  map[int]int // base id => frequency count
}

func parseIntSlice(in []string) []int {
	out := make([]int, len(in))
	for l := range in {
		v, _ := strconv.Atoi(in[l])
		out[l] = v
	}
	return out
}

func parseIntSlice2(in []string) map[int]int {
	out := make(map[int]int, len(in))
	for l := range in {
		v, _ := strconv.Atoi(in[l])
		out[l] = v
	}
	return out
}

func TestMerge(t *testing.T) {

	nametocat := make(map[string]string)
	csvReader := csv.NewReader(bytes.NewReader([]byte(params)))
	records, _ := csvReader.ReadAll()
	for r := range records {
		nametocat[records[r][1]] = records[r][3]
	}
	nametocat["Alien"] = "0"
	nametocat["Ape"] = "0"
	nametocat["Zombie"] = "0"
	nametocat["Male1"] = "0"
	nametocat["Male2"] = "0"
	nametocat["Male3"] = "0"
	nametocat["Male4"] = "0"
	nametocat["Female1"] = "0"
	nametocat["Female2"] = "0"
	nametocat["Female3"] = "0"
	nametocat["Female4"] = "0"

	out := make([][]string, 0)

	csvReader = csv.NewReader(bytes.NewReader([]byte(distributionData)))
	drecords, _ := csvReader.ReadAll()
	count := 1
	for i := range drecords {
		row := make([]string, 0)

		row = append(row, strconv.Itoa(count), drecords[i][1], nametocat[drecords[i][1]])
		row = append(row, drecords[i][2]) // total pop
		row = append(row, drecords[i][4:]...)
		out = append(out, row)
		count++
	}

	sort.SliceStable(out, func(i, j int) bool {
		if out[i][2] == "0" && out[j][2] != "0" {
			return true
		}
		n1, _ := strconv.Atoi(out[i][3])
		n2, _ := strconv.Atoi(out[j][3])
		m1, _ := strconv.Atoi(out[i][4])
		m2, _ := strconv.Atoi(out[j][4])
		return n1+m1 < n2+m2
	})

	count = 0
	for l := range out {
		out[l][0] = strconv.Itoa(count)
		count++
	}

	var b bytes.Buffer
	csvWriter := csv.NewWriter(&b)
	csvWriter.WriteAll(out)
	fmt.Println(b.String())

}

/**
todo base and base types. A map for each base type, this map stores: a map with all layers => blocks
*/
func TestGenerator(t *testing.T) {
	var allBlocks blocks
	var err error
	var blocksPath = "./traits-24x24.png"
	if _, err = allBlocks.load(blocksPath); err != nil {
		fmt.Println(err)
		return
	}

	csvReader := csv.NewReader(bytes.NewReader([]byte(params2)))
	records, err := csvReader.ReadAll()
	//male := make(map[int][]block)
	//female := make(map[int][]block)
	attributes := make(map[int][]block)
	//attributes := make(map[int]block)
	bases := make(map[int][]block)

	for i := 0; i < len(records); i++ {
		cat, _ := strconv.Atoi(records[i][2]) // category
		id, _ := strconv.Atoi(records[i][0])
		frq, _ := strconv.Atoi(records[i][3])

		// new version
		freqm, _ := strconv.Atoi(records[i][4])
		freqf, _ := strconv.Atoi(records[i][5])
		freqal, _ := strconv.Atoi(records[i][6])
		freqape, _ := strconv.Atoi(records[i][7])
		freqzo, _ := strconv.Atoi(records[i][8])

		var blocks map[int][]block

		if freqm > 0 || freqal > 0 || freqape > 0 || freqzo > 0 {
			if i < 11 {
				blocks = bases
			} else {
				blocks = attributes
			}
			blocks[cat] = append(blocks[cat], block{
				i:     allBlocks.getPunkBlock(id),
				id:    id,
				freq:  frq,
				sex:   "m",
				name:  records[i][1],
				cat:   cat,
				stats: parseIntSlice(records[i][9:13]),
				stat:  parseIntSlice2(records[i][6:17]),
			})
		}
		if freqf > 0 {
			if i < 11 {
				blocks = bases
			} else {
				blocks = attributes
			}
			blocks[cat] = append(blocks[cat], block{
				i:     allBlocks.getPunkBlock(id),
				id:    id,
				freq:  freqf,
				sex:   "f",
				name:  records[i][1],
				cat:   cat,
				stats: parseIntSlice(records[i][13:17]),
				stat:  parseIntSlice2(records[i][6:17]),
			})
		}

	}
	sum := 0
	total := 0
	for i := 0; i < 1; i++ {
		counts, zbcount := pickCollection(bases, attributes)
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
	attributes map[int][]block,

) (map[int]int, int) {
	var chosenAtt []block
	counts := make(map[int]int)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	zbcount := 0
	for i := 0; i < 10000; i++ {
		chosenBase := pickBase(base, r1)
		count := pickCount(r1)
		chosenAtt = pickPunk(attributes, chosenBase, r1, count)
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

func pickPunk(set map[int][]block, base block, r1 *rand.Rand, desiredCount int) []block {

	var chosenAtt []block
	if desiredCount == 0 {
		return chosenAtt
	}
	n2 := r1.Intn(10000)
	i := r1.Intn(len(set) - 1)
	catRolls := 0
	catPicked := make(map[int]bool)
	for {
		if i > len(set)-1 {
			i = 0 // wrap around the layers
		}
		if _, ok := set[i]; !ok {
			//fmt.Println("zeroooo!", i)
			i++
			continue
		}
		if _, ok := catPicked[i]; ok {
			i++
			continue
		}

		j := len(set[i]) - 1
		if j > 0 {
			// if the category has more than 1 attribute
			j = r1.Intn(j) // pick a random attribute
		}

		attRolls := 0
		catRolls++

		//fmt.Println("cat roll:", catRolls, ", i:", i)
		for {
			if attRolls >= len(set[i]) {
				break // we rolled all attributes
			}
			n2 = r1.Intn(base.freq)
			//fmt.Println("attRoll, ", attRolls, "cat len", len(set[i]), " n:", set[i][j].name, " ", n2)

			if set[i][j].stat[base.id] >= n2 {
				chosenAtt = append(chosenAtt, set[i][j])
				fmt.Println("picked!", set[i][j])
				catPicked[i] = true
				//attRolls = 0
				//j = 0
				if len(chosenAtt) == desiredCount {
					fmt.Println("choice!", chosenAtt)
					return chosenAtt
				}
				break
			}
			j++
			attRolls++
			if attRolls == len(set[i]) {
				attRolls = 0
				break
			}
			if j >= len(set[i]) {
				j = 0 // wrap around the attributes
			}

		}

		attRolls = 0
		i++

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
