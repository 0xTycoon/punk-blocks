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

var layers = `0,Base,
1,Cheeks,
2,Blemish,
3,Hair,
4,Beard,
5,Eyes,
6,Eyewear,
7,Nose,
8,Mouth,
9,MouthProp,
10,Earring,
11,Headgear
12,Neck`

var params2 = `0,Alien,0,9,9,,,,,,,,,,,,,10,10,10,10,10,10,10,,,,
1,Ape,0,24,24,,,,,,,,,,,,,9,9,9,9,9,9,9,,,,
2,Zombie,0,88,88,,,,,,,,,,,,,8,8,8,8,8,8,8,,,,
2,Male4,0,598,598,,,,,,,,,,,,,0,0,0,0,0,0,0,,,,
3,Male1,0,1723,1723,,,,,,,,,,,,,3,3,3,3,3,3,3,,,,
4,Male2,0,1857,1857,,,,,,,,,,,,,2,2,2,2,2,2,2,,,,
6,Male3,0,1861,1861,,,,,,,,,,,,,1,1,1,1,1,1,1,,,,
7,Female4,0,420,,420,,,,,,,,,,,,,,,,,,,4,4,4,4
8,Female1,0,1101,,1101,,,,,,,,,,,,,,,,,,,7,7,7,7
9,Female3,0,1145,,1145,,,,,,,,,,,,,,,,,,,5,5,5,5
10,Female2,0,1174,,1174,,,,,,,,,,,,,,,,,,,6,6,6,6
11,Choker,12,48,,48,,,,,,,,17,16,13,2,,,,,,,,79,79,79,79
12,Pilot Helmet,3,54,,54,,,,,,,,14,18,17,5,,,,,,,,74,74,74,74
13,Tiara,3,55,,55,,,,,,,,14,17,16,8,,,,,,,,89,89,89,89
14,Orange Side,3,68,,68,,,,,,,,21,18,22,7,,,,,,,,90,90,90,90
15,Welding Goggles,6,86,,86,,,,,,,,30,27,23,6,,,,,,,,115,115,115,115
16,Beanie,3,44,43,,,,1,15,10,15,3,,,,,36,36,36,36,36,36,36,,,,
17,Pigtails,3,94,,94,,,,,,,,22,31,33,8,,,,,,,,110,110,110,110
18,Pink With Hat,3,95,,95,,,,,,,,35,26,28,6,,,,,,,,120,120,120,120
19,Blonde Short,3,129,,129,,,,,,,,35,40,41,13,,,,,,,,100,100,100,100
20,Wild White Hair,3,136,,136,,,,,,,,40,40,40,16,,,,,,,,85,85,85,85
21,Straight Hair Blonde,3,144,,144,,,,,,,,41,46,43,14,,,,,,,,111,111,111,111
22,Wild Blonde,3,144,,144,,,,,,,,33,47,48,16,,,,,,,,82,82,82,82
23,Red Mohawk,3,147,,147,,,,,,,,43,42,45,17,,,,,,,,91,91,91,91
24,Half Shaved,3,147,,147,,,,,,,,45,42,42,18,,,,,,,,128,128,128,128
25,Blonde Bob,3,147,,147,,,,,,,,44,39,46,18,,,,,,,,121,121,121,121
26,Straight Hair Dark,3,148,,148,,,,,,,,42,44,46,16,,,,,,,,78,78,78,78
27,Straight Hair,3,151,,151,,,,,,,,42,44,45,20,,,,,,,,126,126,126,126
28,Buck Teeth,8,78,78,,,,,25,30,19,4,,,,,51,51,51,51,51,51,51,,,,
29,Dark Hair,3,157,,157,,,,,,,,52,53,39,13,,,,,,,,98,98,98,98
30,Tassle Hat,3,178,,178,,,,,,,,53,51,52,22,,,,,,,,75,75,75,75
31,Rosy Cheeks,1,128,67,60,,,1,20,23,21,3,20,18,20,2,11,11,11,11,11,11,11,127,127,127,127
32,Spots,2,124,73,51,,,,19,25,24,5,14,15,14,8,54,54,54,54,54,54,54,86,86,86,86
33,Top Hat,3,115,112,,,1,2,33,43,30,6,,,,,45,45,45,45,45,45,45,,,,
34,Clown Hair Green,3,148,85,63,,,,22,21,30,12,19,19,19,6,13,13,13,13,13,13,13,103,103,103,103
35,Purple Eye Shadow,5,262,,262,,,,,,,,72,88,76,26,,,,,,,,96,96,96,96
36,Blue Eye Shadow,5,266,,266,,,,,,,,67,101,75,23,,,,,,,,77,77,77,77
37,Silver Chain,12,156,113,43,,,,29,40,36,8,10,12,13,8,32,32,32,32,32,32,32,132,132,132,132
38,Green Eye Shadow,5,271,,271,,,,,,,,79,83,82,27,,,,,,,,125,125,125,125
39,Gold Chain,12,169,107,60,,1,1,32,32,32,11,16,20,20,4,67,67,67,67,67,67,67,117,117,117,117
40,Cowboy Hat,3,142,139,,1,2,,31,48,48,12,,,,,15,15,15,15,15,15,15,,,,
41,Medical Mask,9,175,107,65,1,,2,32,31,38,6,17,24,17,7,40,40,40,40,40,40,40,105,105,105,105
42,Big Beard,4,146,146,,,,,49,50,38,9,,,,,55,55,55,55,55,55,55,,,,
43,Vampire Hair,3,147,146,,,,1,45,39,54,8,,,,,56,56,56,56,56,56,56,,,,
44,Purple Hair,3,165,163,,,,2,46,48,47,22,,,,,27,27,27,27,27,27,27,,,,
45,Clown Nose,7,212,134,76,,,2,42,41,45,6,23,19,29,5,17,17,17,17,17,17,17,108,108,108,108
46,Fedora,3,186,184,,,1,1,64,47,58,15,,,,,59,59,59,59,59,59,59,,,,
47,Police Cap,3,203,200,,,1,2,67,47,61,25,,,,,50,50,50,50,50,50,50,,,,
48,Vape,9,272,161,110,,1,,47,53,50,11,28,30,42,10,31,31,31,31,31,31,31,119,119,119,119
49,Smile,8,238,234,,,,4,72,78,62,22,,,,,33,33,33,33,33,33,33,,,,
50,3D Glasses,6,286,205,78,,1,2,53,70,68,14,21,29,23,5,71,71,71,71,71,71,71,83,83,83,83
51,Clown Eyes Blue,5,384,108,274,,,2,35,30,38,5,86,89,75,24,72,72,72,72,72,72,72,93,93,93,93
52,Eye Mask,6,293,205,86,,1,1,53,66,67,19,29,23,29,5,70,70,70,70,70,70,70,102,102,102,102
53,Cap Forward,3,254,248,,1,2,3,72,80,71,25,,,,,24,24,24,24,24,24,24,,,,
54,Pipe,9,317,186,130,1,,,46,54,67,19,37,44,41,8,63,63,63,63,63,63,63,94,94,94,94
55,Hoodie,3,259,256,,,1,2,74,83,72,27,,,,,66,66,66,66,66,66,66,,,,
56,Front Beard Dark,4,260,255,,,,5,60,84,88,23,,,,,65,65,65,65,65,65,65,,,,
57,Frown,8,261,257,,,,4,84,67,87,19,,,,,23,23,23,23,23,23,23,,,,
58,Clown Eyes Green,5,382,136,246,,,,34,45,40,17,70,71,79,26,38,38,38,38,38,38,38,113,113,113,113
59,Handlebars,4,263,261,,,,2,68,90,84,19,,,,,47,47,47,47,47,47,47,,,,
60,Front Beard,4,273,270,,,,3,85,85,81,19,,,,,53,53,53,53,53,53,53,,,,
61,Chinstrap,4,282,280,,,,2,74,87,89,30,,,,,58,58,58,58,58,58,58,,,,
62,Luxurious Beard,4,286,281,,,,5,102,81,72,26,,,,,12,12,12,12,12,12,12,,,,
63,VR,6,332,242,88,,1,1,64,78,73,27,18,31,27,12,42,42,42,42,42,42,42,118,118,118,118
64,Mustache,4,288,286,,,,2,72,90,96,28,,,,,16,16,16,16,16,16,16,,,,
65,Normal Beard Black,4,289,286,,,,3,85,91,85,25,,,,,39,39,39,39,39,39,39,,,,
66,Normal Beard,4,292,291,,,,1,66,106,91,28,,,,,41,41,41,41,41,41,41,,,,
67,Goat,4,295,292,,,,3,86,95,85,26,,,,,25,25,25,25,25,25,25,,,,
68,Do-rag,3,300,292,,1,3,4,88,89,88,27,,,,,52,52,52,52,52,52,52,,,,
69,Cap,3,351,245,98,1,4,3,61,71,79,34,29,31,27,11,37,37,37,37,37,37,37,104,104,104,104
70,Shaved Head,3,300,298,,,,2,95,86,88,29,,,,,29,29,29,29,29,29,29,,,,
71,Peak Spike,3,303,298,,,,5,83,92,94,29,,,,,57,57,57,57,57,57,57,,,,
72,Muttonchops,4,303,303,,,,,100,83,97,23,,,,,68,68,68,68,68,68,68,,,,
73,Black Lipstick,8,617,,617,,,,,,,,171,189,195,62,,,,,,,,130,130,130,130
74,Crazy Hair,3,414,239,168,,,7,56,70,95,18,53,54,46,15,,,,,,,,95,95,95,95
75,Purple Lipstick,8,655,,655,,,,,,,,189,198,195,73,,,,,,,,107,107,107,107
76,Hot Lipstick,8,696,,696,,,,,,,,201,219,200,76,,,,,,,,76,76,76,76
77,Mohawk Dark,3,429,279,146,,,4,81,93,83,22,38,41,49,18,14,14,14,14,14,14,14,129,129,129,129
78,Headband,3,406,304,96,1,1,4,80,108,86,30,25,34,30,7,62,62,62,62,62,62,62,109,109,109,109
79,Knitted Cap,3,419,305,101,2,4,7,91,93,92,29,32,26,32,11,21,21,21,21,21,21,21,112,112,112,112
80,Mohawk Thin,3,441,285,152,,,4,87,85,94,19,44,36,55,17,35,35,35,35,35,35,35,116,116,116,116
81,Mohawk,3,441,286,153,,,2,86,88,81,31,47,47,41,18,73,73,73,73,73,73,73,122,122,122,122
82,Frumpy Hair,3,442,289,149,,,4,77,90,92,30,42,52,36,19,48,48,48,48,48,48,48,87,87,87,87
83,Wild Hair,3,447,296,144,,,7,79,90,86,41,39,48,40,17,44,44,44,44,44,44,44,95,95,95,95
84,Small Shades,6,378,372,,2,2,2,96,129,112,35,,,,,28,28,28,28,28,28,28,,,,
85,Messy Hair,3,460,294,160,,,6,94,90,81,29,40,54,47,19,64,64,64,64,64,64,64,92,92,92,92
86,Stringy Hair,3,463,292,165,,,6,83,88,92,29,48,50,48,19,69,69,69,69,69,69,69,97,97,97,97
87,Bandana,3,481,304,164,2,4,7,78,106,91,29,41,55,49,19,46,46,46,46,46,46,46,106,106,106,106
88,Eye Patch,6,461,363,92,,2,4,105,100,122,36,32,27,20,13,43,43,43,43,43,43,43,99,99,99,99
89,Classic Shades,6,501,345,154,,,3,98,108,108,31,44,42,57,11,30,30,30,30,30,30,30,101,101,101,101
90,Big Shades,6,535,372,159,,1,3,105,121,114,32,53,43,41,22,34,34,34,34,34,34,34,123,123,123,123
91,Regular Shades,6,527,393,128,1,1,4,109,120,119,45,36,45,38,9,20,20,20,20,20,20,20,81,81,81,81
92,Horned Rim Glasses,6,535,388,142,,1,4,119,113,124,32,42,41,50,9,61,61,61,61,61,61,61,131,131,131,131
93,Nerd Glasses,6,572,392,175,,2,3,107,121,133,31,48,58,54,15,19,19,19,19,19,19,19,88,88,88,88
94,Mole,2,644,357,285,,,2,105,110,111,31,90,80,81,34,26,26,26,26,26,26,26,84,84,84,84
95,Shadow Beard,4,526,516,,,,10,166,146,164,40,,,,,22,22,22,22,22,22,22,,,,
96,Cigarette,9,961,557,392,,1,11,151,191,157,58,113,120,120,39,18,18,18,18,18,18,18,114,114,114,114
97,Earring,10,2459,1498,933,3,3,22,412,475,475,136,270,315,257,91,60,60,60,60,60,60,60,124,124,124,124
98,Stogie,9,5000,5000,1000,5,5,44,700,500,500,400,600,600,440,300,133,133,133,133,133,133,133,134,134,134,134
99,Earpiece,6,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,135,135,135,135,135,135,135,136,136,136,136
100,Cig Cap,3,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,137,137,137,137,137,137,137,138,138,138,138
101,Earphones,6,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,139,139,139,130,139,139,139,140,140,140,140
102,Headphone Red,3,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,141,141,141,141,141,141,141,142,142,142,142
103,Headphone Yellow,3,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,143,143,143,143,143,143,143,144,144,144,144
104,Gas Mask,9,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,145,145,145,145,145,145,145,146,146,146,146
105,Googles,6,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,147,147,147,147,147,147,147,148,148,148,148
106,Pen,9,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,149,149,149,149,149,149,149,150,150,150,150
107,Pencil,9,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,151,151,151,151,151,151,151,152,152,152,152
108,Red Hat,3,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,153,153,153,153,153,153,153,154,154,154,154
109,Yellow Hat,3,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,155,155,155,155,155,155,155,156,156,156,156
110,White Hat,3,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,157,157,157,157,157,157,157,158,158,158,158
111,Suit,12,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,160,160,160,160,160,160,160,161,161,161,161
112,Black Suit,12,5000,5000,1000,3,3,44,700,500,500,200,400,300,440,300,162,162,162,162,162,162,162,163,163,163,163

`

var params = `0,Male4,m,0,598
1,Male3,m,0,1861
2,Male2,m,0,1857
3,Male1,m,0,1723
4,Female4,f,0,420
5,Female3,f,0,1145
6,Female2,f,0,1174
7,Female1,f,0,1101
8,Zombie,m,0,88
9,Ape,m,0,24
10,Alien,m,0,9
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
