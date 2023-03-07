package main

import (
	"bytes"
	"encoding/binary"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"io"
	"strconv"
)

func main() {

	fmt.Println("dumping blocks")
	dumpBlocks()
}

type blockstruct struct {
	f     []byte
	m     []byte
	layer string
}

var myBlocks = make(map[string]*blockstruct)
var myBlockKeys = make([]string, 0)

func dumpBlocks() {

	var allBlocks blocks
	var err error
	var blocksPath = "./factory-traits-24x24.png"
	if _, err = allBlocks.load(blocksPath, 18, 10); err != nil {
		fmt.Println(err)
		return
	}

	csvReader := csv.NewReader(bytes.NewReader([]byte(params)))
	records, err := csvReader.ReadAll()
	_ = records
	var str string

	for i := 0; i < len(records); i++ {
		hash := solsha3.SoliditySHA3(
			// types
			[]string{"string"},
			// values
			[]interface{}{
				records[i][1],
			},
		)
		key := hex.EncodeToString(hash)
		var block *blockstruct
		if v, ok := myBlocks[key]; ok {
			block = v
		} else {
			block = &blockstruct{}
			myBlocks[key] = block
			myBlockKeys = append(myBlockKeys, key)
		}
		block.layer = records[i][3]

		b := allBlocks.getPunkBlock(i) // 18 is cig
		b, err = optimizeImage(b)
		var buf bytes.Buffer
		bufw := io.Writer(&buf)
		err = optimalPngCompress(bufw, b)
		//encoded := hex.EncodeToString(buf.Bytes())
		if records[i][2] == "m" {
			block.m = buf.Bytes()
		} else {
			block.f = buf.Bytes()
		}

	}

	for _, blockKey := range myBlockKeys {
		b := myBlocks[blockKey]
		//fmt.Println(blockKey)
		//var block *blockstruct
		//block = myBlocks[key]
		// need number to be 4278843142
		// FF09F706
		layer, _ := strconv.Atoi(b.layer)

		buf := new(bytes.Buffer)
		_ = binary.Write(buf, binary.LittleEndian, uint8(layer))
		err = binary.Write(buf, binary.LittleEndian, uint16(len(b.m)))
		err = binary.Write(buf, binary.LittleEndian, uint16(len(b.f)))
		bytesarr := make([]byte, 8)
		copy(bytesarr, buf.Bytes())
		little := binary.LittleEndian.Uint64(bytesarr)

		//	fmt.Println(buf, err, little)

		str = str + `
	hash = hex"` + blockKey + `";
	blocksInfo[hash] = ` + strconv.Itoa(int(little)) + `;
`
		if len(b.m) > 0 {
			str = str + `	blockL[hash] = hex"` + hex.EncodeToString(b.m) + `";
`
		}
		if len(b.f) > 0 {
			str = str + `	blockS[hash] = hex"` + hex.EncodeToString(b.f) + `";
`
		}

		str = str + `	index[nextId] = bytes32(hash);
	nextId++;
`

	}

	fmt.Println(str)

}

/*
	// create a "red" punk
	var img image.Image
	img = allBlocks.getPunkBlock(4)
	img, _ = optimizeImage(img)
	p := img.(*image.Paletted)
	p.Palette[2] = color.RGBA{255, 0, 0, 255} // change the face color
	var buf bytes.Buffer
	bufw := io.Writer(&buf)
	err = optimalPngCompress(bufw, img)
	encoded := hex.EncodeToString(buf.Bytes())
	fmt.Println(encoded)
	encoded = base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println(encoded)


*/
