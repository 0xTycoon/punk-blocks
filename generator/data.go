package main

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
        9 MouthProp,
        10 Earring,
        11 Headgear
	    12 Neck

Male 1 - the darkest one
Male 4 - the lightest
*/

var params = `0,4,m,0,598
1,3,m,0,1861
2,2,m,0,1857
3,1,m,0,1723
4,4,f,0,420
5,3,f,0,1145
6,2,f,0,1174
7,1,f,0,1101
8,5,m,0,88
9,6,m,0,24
10,7,m,0,9
11,Rosy Cheeks,m,1,128
12,Luxurious Beard,m,4,286
13,Clown Hair Green,m,3,148
14,Mohawk Dark,m,3,429
15,Cowboy Hat,m,11,142
16,Mustache,m,4,288
17,Clown Nose,m,7,212
18,Cigarette,m,9,961
19,Nerd Glasses,m,6,572
20,Regular Shades,m,6,527
21,Knitted Cap,m,11,419
22,Shadow Beard,m,4,526
23,Frown,m,8,261
24,Cap Forward,m,11,254
25,Goat,m,4,295
26,Mole,m,2,644
27,Purple Hair,m,3,165
28,Small Shades,m,6,378
29,Shaved Head,m,3,300
30,Classic Shades,m,6,502
31,Vape,m,9,272
32,Silver Chain,m,12,156
33,Smile,m,8,238
34,Big Shades,m,6,535
35,Mohawk Thin,m,3,441
36,Beanie,m,11,44
37,Cap,m,11,351
38,Clown Eyes Green,m,5,382
39,Normal Beard Black,m,4,289
40,Medical Mask,m,9,175
41,Normal Beard,m,4,292
42,VR,m,6,332
43,Eye Patch,m,5,461
44,Wild Hair,m,3,447
45,Top Hat,m,11,115
46,Bandana,m,11,481
47,Handlebars,m,4,263
48,Frumpy Hair,m,3,442
49,Crazy Hair,m,3,414
50,Police Cap,m,11,203
51,Buck Teeth,m,8,78
52,Do-rag,m,11,300
53,Front Beard,m,4,273
54,Spots,m,2,124
55,Big Beard,m,4,146
56,Vampire Hair,m,3,147
57,Peak Spike,m,3,303
58,Chinstrap,m,4,282
59,Fedora,m,11,186
60,Earring,m,10,2459
61,Horned Rim Glasses,m,6,535
62,Headband,m,11,406
63,Pipe,m,9,317
64,Messy Hair,m,3,460
65,Front Beard Dark,m,4,260
66,Hoodie,m,11,259
67,Gold Chain,m,12,169
68,Muttonchops,m,4,303
69,Stringy Hair,m,3,463
70,Eye Mask,m,6,293
71,3D Glasses,m,6,286
72,Clown Eyes Blue,m,5,384
73,Mohawk,m,3,441
74,Pilot Helmet,f,11,54
75,Tassle Hat,f,11,178
76,Hot Lipstick,f,8,696
77,Blue Eye Shadow,f,5,266
78,Straight Hair Dark,f,3,148
79,Choker,f,12,48
80,Crazy Hair,f,3,414
81,Regular Shades,f,6,527
82,Wild Blonde,f,3,144
83,3D Glasses,f,6,286
84,Mole,f,2,644
85,Wild White Hair,f,3,136
86,Spots,f,2,124
87,Frumpy Hair,f,3,442
88,Nerd Glasses,f,6,572
89,Tiara,f,11,55
90,Orange Side,f,3,68
91,Red Mohawk,f,3,147
92,Messy Hair,f,3,460
93,Clown Eyes Blue,f,5,384
94,Pipe,f,9,317
95,Wild Hair,f,3,447
96,Purple Eye Shadow,f,5,262
97,Stringy Hair,f,3,463
98,Dark Hair,f,3,157
99,Eye Patch,f,6,461
100,Blonde Short,f,3,129
101,Classic Shades,f,6,502
102,Eye Mask,f,6,293
103,Clown Hair Green,f,3,148
104,Cap,f,11,351
105,Medical Mask,f,9,175
106,Bandana,f,11,481
107,Purple Lipstick,f,8,655
108,Clown Nose,f,7,212
109,Headband,f,11,406
110,Pigtails,f,3,94
111,Straight Hair Blonde,f,3,144
112,Knitted Cap,f,11,419
113,Clown Eyes Green,f,5,382
114,Cigarette,f,9,961
115,Welding Goggles,f,6,86
116,Mohawk Thin,f,3,441
117,Gold Chain,f,12,169
118,VR,f,6,332
119,Vape,f,9,272
120,Pink With Hat,f,3,95
121,Blonde Bob,f,3,147
122,Mohawk,f,3,441
123,Big Shades,f,6,535
124,Earring,f,10,2459
125,Green Eye Shadow,f,5,271
126,Straight Hair,f,3,151
127,Rosy Cheeks,f,1,128
128,Half Shaved,f,3,147
129,Mohawk Dark,f,3,429
130,Black Lipstick,f,8,617
131,Horned Rim Glasses,f,6,535
132,Silver Chain,f,12,156
`

// id, trait, total, male, female, male4, male 3, male 2, male 1, female 4, female 3, female2, female 1
var distributionData = `1,Cigarette,961,557,392,151,191,157,58,113,120,120,39
2,Medical Mask,175,107,65,32,31,38,6,17,24,17,7
3,Nerd Glasses,572,392,175,107,121,133,31,48,58,54,15
4,Orange Side,68,,68,,,,,21,18,22,7
5,Dark Hair,157,,157,,,,,52,53,39,13
6,Red Mohawk,147,,147,,,,,43,42,45,17
7,Muttonchops,303,303,,100,83,97,23,,,,
9,Buck Teeth,78,78,,25,30,19,4,,,,
10,Small Shades,378,372,,96,129,112,35,,,,
11,Cap,351,245,98,61,71,79,34,29,31,27,11
12,Eye Patch,461,363,92,105,100,122,36,32,27,20,13
13,Regular Shades,527,393,128,109,120,119,45,36,45,38,9
14,Silver Chain,156,113,43,29,40,36,8,10,12,13,8
15,Mohawk Dark,429,279,146,81,93,83,22,38,41,49,18
16,Pilot Helmet,54,,54,,,,,14,18,17,5
17,Blonde Short,129,,129,,,,,35,40,41,13
18,Straight Hair Blonde,144,,144,,,,,41,46,43,14
19,Wild Blonde,144,,144,,,,,33,47,48,16
22,Top Hat,115,112,,33,43,30,6,,,,
23,Rosy Cheeks,128,67,60,20,23,21,3,20,18,20,2
24,Normal Beard,292,291,,66,106,91,28,,,,
25,Smile,238,234,,72,78,62,22,,,,
26,Messy Hair,460,294,160,94,90,81,29,40,54,47,19
27,Hoodie,259,256,,74,83,72,27,,,,
28,Stringy Hair,463,292,165,83,88,92,29,48,50,48,19
29,Wild White Hair,136,,136,,,,,40,40,40,16
30,Half Shaved,147,,147,,,,,45,42,42,18
32,Big Beard,146,146,,49,50,38,9,,,,
33,Knitted Cap,419,305,101,91,93,92,29,32,26,32,11
34,Purple Hair,165,163,,46,48,47,22,,,,
36,Frumpy Hair,442,289,149,77,90,92,30,42,52,36,19
38,Mole,644,357,285,105,110,111,31,90,80,81,34
39,Earring,2459,1498,933,412,475,475,136,270,315,257,91
40,VR,332,242,88,64,78,73,27,18,31,27,12
41,Eye Mask,293,205,86,53,66,67,19,29,23,29,5
42,Welding Goggles,86,,86,,,,,30,27,23,6
43,Pigtails,94,,94,,,,,22,31,33,8
44,Pink With Hat,95,,95,,,,,35,26,28,6
46,Gold Chain,169,107,60,32,32,32,11,16,20,20,4
47,Choker,48,,48,,,,,17,16,13,2
48,Normal Beard Black,289,286,,85,91,85,25,,,,
49,Shaved Head,300,298,,95,86,88,29,,,,
50,Purple Lipstick,655,,655,,,,,189,198,195,73
51,Big Shades,535,372,159,105,121,114,32,53,43,41,22
52,Vampire Hair,147,146,,45,39,54,8,,,,
53,Peak Spike,303,298,,83,92,94,29,,,,
54,Blue Eye Shadow,266,,266,,,,,67,101,75,23
55,Pipe,317,186,130,46,54,67,19,37,44,41,8
57,Fedora,186,184,,64,47,58,15,,,,
58,Spots,124,73,51,19,25,24,5,14,15,14,8
59,Chinstrap,282,280,,74,87,89,30,,,,
60,Mustache,288,286,,72,90,96,28,,,,
61,Blonde Bob,147,,147,,,,,44,39,46,18
62,Horned Rim Glasses,535,388,142,119,113,124,32,42,41,50,9
63,Straight Hair Dark,148,,148,,,,,42,44,46,16
65,Luxurious Beard,286,281,,102,81,72,26,,,,
66,Front Beard,273,270,,85,85,81,19,,,,
67,Clown Nose,212,134,76,42,41,45,6,23,19,29,5
68,Do-rag,300,292,,88,89,88,27,,,,
69,Wild Hair,447,296,144,79,90,86,41,39,48,40,17
70,Police Cap,203,200,,67,47,61,25,,,,
71,Mohawk,441,286,153,86,88,81,31,47,47,41,18
72,Mohawk Thin,441,285,152,87,85,94,19,44,36,55,17
73,Tiara,55,,55,,,,,14,17,16,8
75,Frown,261,257,,84,67,87,19,,,,
76,Goat,295,292,,86,95,85,26,,,,
77,Vape,272,161,110,47,53,50,11,28,30,42,10
78,Clown Eyes Blue,384,108,274,35,30,38,5,86,89,75,24
79,Handlebars,263,261,,68,90,84,19,,,,
80,Clown Eyes Green,382,136,246,34,45,40,17,70,71,79,26
81,Classic Shades,501,345,154,98,108,108,31,44,42,57,11
82,Cowboy Hat,142,139,,31,48,48,12,,,,
83,Bandana,481,304,164,78,106,91,29,41,55,49,19
84,Front Beard Dark,260,255,,60,84,88,23,,,,
85,Straight Hair,151,,151,,,,,42,44,45,20
86,Shadow Beard,526,516,,166,146,164,40,,,,
87,Beanie,44,43,,15,10,15,3,,,,
88,Green Eye Shadow,271,,271,,,,,79,83,82,27
90,Crazy hair,414,239,168,56,70,95,18,53,54,46,15
91,Tassle Hat,178,,178,,,,,53,51,52,22
92,Headband,406,304,96,80,108,86,30,25,34,30,7
93,Black Lipstick,617,,617,,,,,171,189,195,62
94,3D Glasses,286,205,78,53,70,68,14,21,29,23,5
95,Clown Hair Green,148,85,63,22,21,30,12,19,19,19,6
96,Hot Lipstick,696,,696,,,,,201,219,200,76
97,Cap Forward,254,248,,72,80,71,25,,,,
98,Purple Eye Shadow,262,,262,,,,,72,88,76,26
`
