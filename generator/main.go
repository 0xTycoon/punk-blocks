package main

import (
	"bytes"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"io"
)

func main() {

	fmt.Println("dumping blocks")
	dumpBlocks()
}

func dumpBlocks() {

	var allBlocks blocks
	var err error
	var blocksPath = "./traits-24x24.png"
	if _, err = allBlocks.load(blocksPath, 140, 10); err != nil {
		fmt.Println(err)
		return
	}

	csvReader := csv.NewReader(bytes.NewReader([]byte(params)))
	records, err := csvReader.ReadAll()
	_ = records
	var str string
	declare := "Block storage "
	for i := 0; i < 133; i++ {
		hash := solsha3.SoliditySHA3(
			// types
			[]string{"string"},
			// values
			[]interface{}{
				records[i][1],
			},
		)
		key := hex.EncodeToString(hash)

		b := allBlocks.getPunkBlock(i) // 18 is cig

		b, err = optimizeImage(b)
		var buf bytes.Buffer
		bufw := io.Writer(&buf)
		err = optimalPngCompress(bufw, b)
		encoded := hex.EncodeToString(buf.Bytes())
		//encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

		str = str +
			declare + `b = blocks[bytes32(hex"` + key + `")];
b.layer = Layer(` + records[i][3] + `);
`
		if records[i][2] == "m" {
			str = str + `b.dataMale = hex"` + encoded + `";
`
		} else {
			str = str + `b.dataFemale = hex"` + encoded + `";
`
		}
		str = str + `index[nextId] = bytes32(hex"` + key + `");
nextId++;
`
		declare = ""
	}

	fmt.Println(str)
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

}
