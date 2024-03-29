package main

import (
	"bytes"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"github.com/miguelmota/go-solidity-sha3"
	"image"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

/*
The Economic theory of our Harberger Tax says that the CEO title will always go
to the most capable and competent holder possible.
	The biggest reason why the CIGs are burned when buying the title is
bought, is to discourage CEOs who are simply motivated by money and are only
there to make CIG. Likewise, this is the reason why the CEO is not
being paid CIG, but has to pay a tax to maintain the role.
Also, burning of CIG is a method of re-distributing
profits equally to all the holders. Thus the CEO's interest is aligned
with the holders.
*/
// stogie male, stogie female, earphone fem, earphone male, employee hat male, employee hat fem, headphones male
// headphone fem, headphone red male, headphones red fem, headphones yell fem, gasmask m, gasmask f, googles m, goggles f
// pen f, pencil m, pencil f, red hat m, red hat f, white hat m, suit m
// suit f, bot m, Bot f, Killer Bot m, Killer Bot f, Green Alien m,
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
	id   int
	freq int
	name string
	cat  int
	stat map[int]int // base id => frequency count
	img  map[int]int // base id => image title id, for every stat
	sex  string
}

func (b *block) save() {

	f, err := os.Create("attr-" + b.name + "image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, b.i); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func TestMergeLayers(t *testing.T) {

	csvReader := csv.NewReader(bytes.NewReader([]byte(params2)))
	data, _ := csvReader.ReadAll()
	for r := range data {
		if data[r][2] == "11" {
			data[r][2] = "3"
		}
	}

	var b bytes.Buffer
	csvWriter := csv.NewWriter(&b)
	csvWriter.WriteAll(data)
	fmt.Println(b.String())

}

func TestMerge(t *testing.T) {
	nametocat := make(map[string]string)
	csvReader := csv.NewReader(bytes.NewReader([]byte(params)))
	records, _ := csvReader.ReadAll()
	for r := range records {
		nametocat[records[r][1]+records[r][2]] = records[r][0]
	}

	csvReader = csv.NewReader(bytes.NewReader([]byte(params2)))
	data, _ := csvReader.ReadAll()
	//sex := "m"
	for r := range data {
		if id, ok := nametocat[data[r][1]+"m"]; ok {
			data[r] = append(data[r], id, id, id, id, id, id, id)
		} else {
			data[r] = append(data[r], "", "", "", "", "", "", "")
		}

		if id, ok := nametocat[data[r][1]+"f"]; ok {
			data[r] = append(data[r], id, id, id, id)
		} else {
			data[r] = append(data[r], "", "", "", "")
		}

	}

	var b bytes.Buffer
	csvWriter := csv.NewWriter(&b)
	csvWriter.WriteAll(data)
	fmt.Println(b.String())

}

func parseIntSlice(in []string) map[int]int {
	out := make(map[int]int, len(in))
	for l := range in {
		v, _ := strconv.Atoi(in[l])
		out[l] = v
	}
	return out
}

func TestAttributes(t *testing.T) {
	var blocksPath = "./factory-traits-24x24.png"

	var err error
	if _, err = allBlocks.load(blocksPath, 20, 10); err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	csvReader := csv.NewReader(bytes.NewReader([]byte(params)))
	records, err := csvReader.ReadAll()
	attributes := make(map[int]block)
	male := make(map[int]block)
	female := make(map[int]block)
	bases := make(map[int]block)
	var fcnt, mcnt int
	var blocks map[int]block
	for i := 0; i < len(records); i++ {
		cat, _ := strconv.Atoi(records[i][3]) // layer
		id, _ := strconv.Atoi(records[i][0])
		frq, _ := strconv.Atoi(records[i][3])

		if records[i][3] == "0" {
			blocks = bases
		} else {
			blocks = attributes
		}
		blocks[i] = block{
			i:    allBlocks.getPunkBlock(id),
			id:   id,
			freq: frq,
			name: records[i][1],
			cat:  cat,
			sex:  records[i][2],
		}

		if records[i][3] != "0" && i < 133 {
			// collect non factory traits
			if records[i][2] == "m" {
				male[mcnt] = blocks[i]
				mcnt++
			} else {
				female[fcnt] = blocks[i]
				fcnt++
			}
		}

	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	c := collection{cols: 100, rows: 100}
	baseM := bases[1] // 171
	baseF := bases[5] // 175
	testAttr := make(map[int]block)
	for att := range attributes {
		if att < 133 {
			continue
		}

		if attributes[att].sex == "m" {
			testAttr = make(map[int]block)
			testAttr[attributes[att].cat] = attributes[att]
			testAttr = randumoor(male, testAttr, r1, "m")
			c.drawPunk2(baseM, testAttr)
		}

		if attributes[att].sex == "f" {
			testAttr = make(map[int]block)
			_ = r1
			testAttr[attributes[att].cat] = attributes[att]
			testAttr = randumoor(female, testAttr, r1, "f")
			c.drawPunk2(baseF, testAttr)
		}

	}
	c.save()

}

func randumoor(selection map[int]block, source map[int]block, r1 *rand.Rand, sex string) map[int]block {

	for picked := 0; picked < 8; {
		i := r1.Intn(len(selection))
		pick := selection[i]
		if _, ok := source[pick.cat]; !ok && pick.sex == sex {
			source[pick.cat] = pick
			picked++
		}

	}

	return source
}

/**
todo base and base types. A map for each base type, this map stores: a map with all layers => blocks
*/
func TestGenerator(t *testing.T) {

	var blocksPath = "./factory-traits-24x24.png"
	attributes, bases, err := loadBases(blocksPath, 20, 10)
	if err != nil {
		t.Error(err)
		return
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
	fmt.Println("all:", sum)
	sum = 0
	/*

		0. in solidity, keep a structure pre-configured as above.

		1. get random number for index

		2. using random starting index, roll a category 0. Keep rolling until one is found

		3. repeat for remaining attys

	*/
}

var allBlocks blocks

func loadBases(blocksPath string, rows int, cols int) (map[int][]block, map[int][]block, error) {

	var err error
	if _, err = allBlocks.load(blocksPath, rows, cols); err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	csvReader := csv.NewReader(bytes.NewReader([]byte(params2)))
	records, err := csvReader.ReadAll()
	attributes := make(map[int][]block)
	bases := make(map[int][]block)
	for i := 0; i < len(records); i++ {
		cat, _ := strconv.Atoi(records[i][2]) // layer
		id, _ := strconv.Atoi(records[i][0])
		frq, _ := strconv.Atoi(records[i][3])

		var blocks map[int][]block

		if records[i][3] == "0" {
			blocks = bases
		} else {
			blocks = attributes
		}
		blocks[cat] = append(blocks[cat], block{
			i:    allBlocks.getPunkBlock(id),
			id:   id,
			freq: frq,
			name: records[i][1],
			cat:  cat,
			stat: parseIntSlice(records[i][6:17]),
			img:  parseIntSlice(records[i][17:28]),
		})

	}
	return attributes, bases, err
}

type collection struct {
	i      *image.RGBA
	nextID int
	cols   int
	rows   int
}

func (c *collection) drawPunk2(base block, attributes map[int]block) {
	x := c.nextID % c.cols * 24
	y := (c.rows * c.nextID) / (c.rows * 100) * 24
	//fmt.Println("x:", x, " y:", y)
	var img *image.RGBA
	if c.i == nil {
		fmt.Println("img nil")
		c.i = image.NewRGBA(image.Rect(0, 0, c.cols*24, c.rows*24))
	}
	img = base.i.(*image.RGBA) //allBlocks.getPunkBlock(base.img[base.id]).(*image.RGBA)

	draw.Draw(c.i, image.Rect(x, y, x+24, y+24), img, img.Bounds().Min, draw.Over)
	base.save()
	fmt.Println("punkid:", c.nextID, attributes)
	for i := 0; i < 13; i++ {
		if v, ok := attributes[i]; ok {
			img = v.i.(*image.RGBA) //allBlocks.getPunkBlock(v.img[base.id]).(*image.RGBA)
			draw.Draw(c.i, image.Rect(x, y, x+24, y+24), img, img.Bounds().Min, draw.Over)
			//v.save()
		}
	}
	if c.nextID == 20 {
		//c.save()
		//os.Exit(0)
	}

	c.nextID++
}

func (c *collection) drawPunk(base block, attributes map[int]block) {
	x := c.nextID % c.cols * 24
	y := (c.rows * c.nextID) / (c.rows * 100) * 24
	//fmt.Println("x:", x, " y:", y)
	var img *image.RGBA
	if c.i == nil {
		fmt.Println("img nil")
		c.i = image.NewRGBA(image.Rect(0, 0, c.cols*24, c.rows*24))
	}
	img = allBlocks.getPunkBlock(base.img[base.id]).(*image.RGBA)

	draw.Draw(c.i, image.Rect(x, y, x+24, y+24), img, img.Bounds().Min, draw.Over)
	//base.save()
	fmt.Println("punkid:", c.nextID, attributes)
	for i := 0; i < 13; i++ {
		if v, ok := attributes[i]; ok {
			img = allBlocks.getPunkBlock(v.img[base.id]).(*image.RGBA)
			draw.Draw(c.i, image.Rect(x, y, x+24, y+24), img, img.Bounds().Min, draw.Over)
			//v.save()
		}
	}
	if c.nextID == 20 {
		//c.save()
		//os.Exit(0)
	}

	c.nextID++
}

func (c *collection) save() {

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, c.i); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func TestCollectionMaker(t *testing.T) {
	c := collection{cols: 100, rows: 100}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var blocksPath = "./traits-24x24.png"
	attributes, bases, err := loadBases(blocksPath, 16, 10)
	_ = attributes
	if err != nil {
		t.Error(err)
		return
	}
	chosenBase := pickBase(bases, r1)
	for i := 0; i < 110; i++ {
		chosenBase.save()
		c.drawPunk(chosenBase, make(map[int]block, 1))
	}
	c.save()

}

func pickCollection(
	base map[int][]block,
	attributes map[int][]block,

) (map[int]int, int) {
	var chosenAtt map[int]block
	counts := make(map[int]int)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	zbcount := 0
	c := collection{cols: 100, rows: 100}
	for i := 0; i < 10000; i++ {
		chosenBase := pickBase(base, r1)
		count := pickCount(r1)
		chosenAtt = pickPunk(attributes, chosenBase, r1, count)
		c.drawPunk(chosenBase, chosenAtt)
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
	c.save()
	return counts, zbcount
}

var attrCounts = make(map[string]int)

func pickPunk(set map[int][]block, base block, r1 *rand.Rand, desiredCount int) map[int]block {
	chosenAtt := make(map[int]block)
	//var chosenAtt []block
	if desiredCount == 0 {
		return chosenAtt
	}
	n2 := r1.Intn(10000)
	i := r1.Intn(len(set) - 1)
	catRolls := 0

	for {
		if i > len(set)-1 {
			i = 0 // wrap around the layers
		}
		if _, ok := set[i]; !ok {
			//fmt.Println("zeroooo!", i)
			i++
			continue
		}
		if _, ok := chosenAtt[i]; ok { // was there an atty picked from this layer
			i++
			continue
		}
		j := len(set[i]) - 1
		if j > 0 {
			// if the layer has more than 1 attribute
			j = r1.Intn(j) // pick a random attribute
		}

		attRolls := 0
		catRolls++
		if catRolls > len(set)*6 {
			return chosenAtt // sometimes the odds of finding a match may be impossible
		}
		//fmt.Println("cat roll:", catRolls, ", i:", i)
		for {
			if attRolls >= len(set[i]) {
				break // we rolled all attributes
			}

			//fmt.Println("attRoll, ", attRolls, "cat len", len(set[i]), " n:", set[i][j].name, " ", n2)

			n2 = r1.Intn(base.freq)
			freq := set[i][j].stat[base.id]

			if freq != 0 && freq >= n2 {
				//chosenAtt = append(chosenAtt, set[i][j])
				//fmt.Println("picked!", set[i][j])
				chosenAtt[i] = set[i][j]

				if len(chosenAtt) == desiredCount {
					//fmt.Println("choice!", chosenAtt)
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

//var counts = [8]int{8, 333, 3560, 4501, 1420, 166, 11, 1}

var counts = [11]int{0, 0, 1000, 1000, 1860, 2000, 2534, 1020, 566, 11, 9}

func pickCount(r1 *rand.Rand) int {
	i := r1.Intn(8)
	for {
		n2 := r1.Intn(10000)
		if counts[i] == 0 {
			i++
			continue
		}
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
