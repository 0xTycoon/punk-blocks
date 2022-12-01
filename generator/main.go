package main

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"image"
	"image/color"
	"io"
)

func main() {

	fmt.Println("dumping blocks")
	dumpBlocks()
}

func dumpBlocks() {

	var allBlocks block
	var err error
	var blocksPath = "./traits-24x24.png"
	if _, err = allBlocks.load(blocksPath); err != nil {
		fmt.Println(err)
		return
	}
	/**

			0 Base,
	        1 Cheeks,
	        2 Blemish,
	        3 Hair,
	        4 Beard,
	        5 Eyes,
	        6 Eyewear,
	        7 Nose,
	        8 Mouth,
	        8 MouthProp,
	        10 Earring,
	        11 Headgear
		    12 Neck
	*/

	params := `0,Male 1,m,0
1,Male 2,m,0
2,Male 3,m,0
3,Male 4,m,0
4,Female 1,f,0
5,Female 2,f,0
6,Female 3,f,0
7,Female 4,f,0
8,Zombie,m,0
9,Ape,m,0
10,Alien,m,0
11,Rosy Cheeks,m,1
12,Luxurious Beard,m,4
13,Clown Hair Green,m,3
14,Mohawk Dark,m,3
15,Cowboy Hat,m,11
16,Mustache,m,4
17,Clown Nose,m,7
18,Cigarette,m,8
19,Nerd Glasses,m,6
20,Regular Shades,m,6
21,Knitted Cap,m,11
22,Shadow Beard,m,4
23,Frown,m,8
24,Cap Forward,m,11
25,Goat,m,4
26,Mole,m,2
27,Purple Hair,m,3
28,Small Shades,m,6
29,Shaved Head,m,3
30,Classic Shades,m,6
31,Vape,m,8
32,Silver Chain,m,12
33,Smile,m,8
34,Big Shades,m,6
35,Mohawk Thin,m,3
36,Beanie,m,11
37,Cap,m,11
38,Clown Eyes Green,m,5
39,Normal Beard Black,m,4
40,Medical Mask,m,8
41,Normal Beard,m,4
42,VR,m,6
43,Eye Patch,m,5
44,Wild Hair,m,3
45,Top Hat,m,11
46,Bandana,m,11
47,Handlebars,m,4
48,Frumpy Hair,m,3
49,Crazy Hair,m,3
50,Police Cap,m,11
51,Buck Teeth,m,8
52,Do-rag,m,11
53,Front Beard,m,4
54,Spots,m,2
55,Big Beard,m,4
56,Vampire Hair,m,3
57,Peak Spike,m,3
58,Chinstrap,m,4
59,Fedora,m,11
60,Earring,m,10
61,Horned Rim Glasses,m,6
62,Headband,m,11
63,Pipe,m,8
64,Messy Hair,m,3
65,Front Beard Dark,m,4
66,Hoodie,m,11
67,Gold Chain,m,12
68,Muttonchops,m,4
69,Stringy Hair,m,3
70,Eye Mask,m,6
71,3D Glasses,m,6
72,Clown Eyes Blue,m,5
73,Mohawk,m,3
74,Pilot Helmet,f,11
75,Tassle Hat,f,11
76,Hot Lipstick,f,8
77,Blue Eye Shadow,f,5
78,Straight Hair Dark,f,3
79,Choker,f,12
80,Crazy Hair,f,3
81,Regular Shades,f,6
82,Wild Blonde,f,3
83,3D Glasses,f,6
84,Mole,f,2
85,Wild White Hair,f,3
86,Spots,f,2
87,Frumpy Hair,f,3
88,Nerd Glasses,f,6
89,Tiara,f,11
90,Orange Side,f,3
91,Red Mohawk,f,3
92,Messy Hair,f,3
93,Clown Eyes Blue,f,5
94,Pipe,f,8
95,Wild Hair,f,3
96,Purple Eye Shadow,f,5
97,Stringy Hair,f,3
98,Dark Hair,f,3
99,Eye Patch,f,6
100,Blonde Short,f,3
101,Classic Shades,f,6
102,Eye Mask,f,6
103,Clown Hair Green,f,3
104,Cap,f,11
105,Medical Mask,f,8
106,Bandana,f,11
107,Purple Lipstick,f,8
108,Clown Nose,f,7
109,Headband,f,11
110,Pigtails,f,3
111,Straight Hair Blonde,f,3
112,Knitted Cap,f,11
113,Clown Eyes Green,f,5
114,Cigarette,f,8
115,Welding Goggles,f,6
116,Mohawk Thin,f,3
117,Gold Chain,f,12
118,VR,f,6
119,Vape,f,8
120,Pink With Hat,f,3
121,Blonde Bob,f,3
122,Mohawk,f,3
123,Big Shades,f,6
124,Earring,f,10
125,Green Eye Shadow,f,5
126,Straight Hair,f,3
127,Rosy Cheeks,f,1
128,Half Shaved,f,3
129,Mohawk Dark,f,3
130,Black Lipstick,f,8
131,Horned Rim Glasses,f,6
132,Silver Chain,f,12
`

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

}
