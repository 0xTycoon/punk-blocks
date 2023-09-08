
![PunkBlocks](title.png)

# PunkBlocks

This project deploys a contract to the Ethereum mainnet that can store 24x24 
blocks of punk attributes as png images.

Each block is numbered and indexed by its name.

A view function combines the attributes, outputting the composite image into
a single SVG image.

Anybody can register and store new punk block attributes, then use to create 
their own hybrid punks, or display the original punks. 

## Deployment

There are two contracts deployed, `punkblocks.sol` is used for storage and 
`renderblocks.sol` is used for generating the svg images. Originally,
`punkblocks.sol` did both storage and rendering, however, a bug was discovered
with the layer ordering, so the new contract rectifies this bug while building
on top of the punkblocks.sol. Both contracts are deployed to the Ethereum
mainnet.

`punkblocks.sol`   => [`0xe91eb909203c8c8cad61f86fc44edee9023bda4d`](https://etherscan.io/address/0xe91eb909203c8c8cad61f86fc44edee9023bda4d#code)

`renderblocks.sol` => [`0x829e113C94c1acb6b1b5577e714E486bb3F86593`](https://etherscan.io/address/0x829e113C94c1acb6b1b5577e714E486bb3F86593#code)

`punkblocks.sol` [seal() transaction](https://etherscan.io/tx/0xcd9922d1f7d1859a5d0a4eeccf5bed4f937eef0cc29d4fd39252fe1fb1851dbd) done.

Update: It's not possible to deploy the contracts in the current form if the chain impliments EIP-3860.
This EIP has been enabled on Ethereum in the Shangai upgrade. The above contracts were deployed
just before this upgrade. To fix the contracts, they would need to be initialized via an external function (TODO).

### Simple UI (Demo)

Here is a [simple web UI](https://0xtycoon.github.io/punk-blocks/) for generating punks from blocks, or
viewing and registering new blocks with the contract.

## Prior art, scope and purpose

Putting the CryotoPunks images on-chain is not new. In 2021, Larva Labs released 
a contract that sealed all 10,000 punks on chain, including their attribute 
lists.

"Punk Blocks" has a different scope and purpose. The difference is that 
"Punk Blocs" stores the images of the attributes, rather than their entire 
generated image. This allows the generation of entire punk images on-the-fly by 
calling the "view" methods of the Punk Blocks contract.

Another difference is that besides the original attributes, it is possible for 
anyone to add new attribute images to the contract, i.e. "Punk Blocks", and new 
unique punk images can be generated with these blocks. 
So, the scope of this contract is not to just store the original 10,000 punk 
images, but the punk images of an entire punk universe!

"Woah! Hold it... You are saying that we can now make unlimited number of punks, 
including new punks??!!" Yup. Although, these are only images, not tokens. 
This contract doesn't make any tokens, and there will always ever be 10,000 
punks.

Finally, some acknowledgments & thanks. "Punk Blocks" makes indirect use of the 
contract released by Larva Labs in 2021, to fetch a list of attributes that 
belong to a punk.

Naturally, this project would not have been possible without Larva Labs, so 
thanks to them for creating the punks & taking initiative to deploy the 
additional metadata on-chain. Also, thanks to 
[@samwilsn](https://github.com/samwilsn) who showed that it 
[is possible](https://binarycake.ca/posts/face-png/) to store png images and 
render them inside SVGs, and thanks to [@geraldb](https://github.com/geraldb) 
for providing a 
[handy sprite sheet](https://github.com/cryptopunksnotdead/punks.js)!
Finally, thanks to jeremy.eth and twitter.com/dumbnamenumbers for the review & 
feedback.

### Interface

```solidity
/**
* @dev svgFromPunkID returns the svg data as a string given a punk id
* @param _tokenID uint256 IDs a punk id, 0-9999
* @param _size the width and height of generated svg, eg. 24
* @param _orderID which order config to use when rendering, 0 is the default
*/
function svgFromPunkID(uint256 _tokenID, uint16 _x, uint16 _y, uint16 _size, uint32 _orderID) external view returns (string memory);

/**
* @dev svgFromNames returns the svg data as a string
* @param _attributeNames a list of attribute names, eg "Male 1", "Goat"
*   must have at least 1 layer 0 attribute (eg. Male, Female, Alien, Ape, Zombie)
*   e.g. ["Male 1","Goat"]
*   Where "Male 1" is a layer 0 attribute, that decides what version of
*   image to use for the higher
*   layers (dataMale or dataFemale)
* @param _size the width and height of generated svg, eg. 24
* @param _orderID which order config to use when rendering, 0 is the default
*/
function svgFromNames(string[] calldata _attributeNames, uint16 _x, uint16 _y, uint16 _size, uint32 _orderID) external view returns (string memory);

/**
* @dev svgFromKeys returns the svg data as a string
* @param _attributeKeys a list of attribute names that have been hashed,
*    eg keccak256("Male 1"), keccak256("Goat")
*    must have at least 1 layer 0 attribute (eg. keccak256("Male 1")) which
*    decides what version of image to use for the higher layers
*    (dataMale or dataFemale)
*    e.g. ["0x9039da071f773e85254cbd0f99efa70230c4c11d63fce84323db9eca8e8ef283",
*    "0xd5de5c20969a9e22f93842ca4d65bac0c0387225cee45a944a14f03f9221fd4a"]
* @param _size the width and height of generated svg, eg. 24
* @param _orderID which order config to use when rendering, 0 is the default
*/
function svgFromKeys(bytes32[] calldata _attributeKeys, uint16 _x, uint16 _y, uint16 _size, uint32 _orderID) external view returns (string memory);

/**
* @dev svgFromIDs returns the svg data as a string
*   e.g. [9,55,99]
*   One of the elements must be must be a layer 0 block.
*   This element decides what version of image to use for the higher layers
*   (dataMale or dataFemale)
* @param _ids uint256 ids of an attribute, by it's index of creation
* @param _size the width and height of generated svg, eg. 24
* @param _orderID which order config to use when rendering, 0 is the default
*/
function svgFromIDs(uint256[] calldata _ids, uint16 _x, uint16 _y, uint16 _size, uint32 _orderID) external view returns (string memory);

/**
* @dev registerBlock allows anybody to add a new block to the contract.
*   Either _dataMale or _dataFemale, or both, must contain a byte stream of a png file.
*   It's best if the png is using an 'index palette' and the lowest bit depth possible,
*   while keeping the highest compression setting.
* @param _dataL png data for the male version, 24x24
* @param _dataS png data for the female version, 24x24
* @param _layer 0 to 12, corresponding to the Layer enum type.
* @param _name the name of the trait, Camel Case. e.g. "Luxurious Beard"
*/
function registerBlock(
    bytes calldata _dataL,
    bytes calldata _dataS,
    uint8 _layer,
    string memory _name) external;

function blockS(bytes32) view external returns(bytes32);

function blockS(bytes32) view external returns(bytes32);

function blocksInfo(bytes32) view external returns(uint256);

```

### Layers

Each registered block belongs to a Layer. A layer determines what order
the block gets applied when the SVG is rendered. "Base" is applied first, 
then a "Nose" block would be applied last.

There are 13 layers in total, numbered from 0-12. The attribute names shown in 
the comments below are examples from the default attribute set. 

Note that head-top layers are separated (Hair and Hat), this allows for more trait combinations, for example, a Cap with Hair combination becomes possible.

```solidity
enum Layer {
    Base,      // 0 Base is the face. Determines if m or f version will be used to render the remaining layers
    Mouth,     // 1 (Hot Lipstick, Smile, Buck Teeth, ...)
    Cheeks,    // 2 (Rosy Cheeks)
    Blemish,   // 3 (Mole, Spots)
    Eyes,      // 4 (Clown Eyes Green, Green Eye Shadow, ...)
    Neck,      // 5 (Choker, Silver Chain, Gold Chain)
    Beard,     // 6 (Big Beard, Front Beard, Goat, ...)
    Ears,      // 7 (Earring)
    HeadTop1,  // 8 (Purple Hair, Shaved Head, Beanie, Fedora,Hoodie)
    HeadTop2,  // 9  eg. additional hat over hair (not used by LL punks)
    Eyewear,   // 10 (VR, 3D Glass, Eye Mask, Regular Shades, Welding Glasses, ...)
    MouthProp, // 11 (Medical Mask, Cigarette, ...)
    Nose       // 12 (Clown Nose)
}
```

It's possible to define your own layer ordering. Here is the default order
used by `renderblocks.sol`. This is also assumed to be the ordering of the
cryptopunks.

```
0 Base
2 Cheeks,
3 Blemish,
1 Mouth,
5 Neck,
6 Beard,
7 Earring,
8 HeadTop1,
9 HeadTop2,
4 Eyes,
11 MouthProp,
10 Eyewear,
12 Nose
```

### Adding new blocks.

New blocks can be added for any layer using the `registerBlock` function.
The block must consist of either a male attribute image data, or female data,
or both. With the exception that Layer 0 blocks **should** either contain male
attribute image data **or** female data, but **NOT both**.

If the block is a Layer 0 with *female* attribute image data present, then all 
higher blocks will be rendered with their *female* attribute image data, and 
vice-versa.

Preparing the png for upload can be tricky. It's best to use an **indexed 
palette** rather than an RGBA, using a single entry for the transparency,
with a low bit-depth. E.g. You could use a bit-depth of 2 if you have 4 colors,
or 4 if you have 16. Confused? Here is a 
[complete guide](https://optipng.sourceforge.net/pngtech/optipng.html) about 
png optimization.

There are several nifty png optimization tools available for Linux, such as
[optipng](https://www.cyberciti.biz/faq/linux-unix-optimize-lossless-png-images-with-optipng-command/)

What was used here? Well...
A simple Go program was custom-made to prepare the optimized png images. see: 
[generator](./generator/) .
It reads some blocks from a source png file and breaks them down to 24x24
images. Each image is then saved as an optimized png, then the Solidity code 
is generated with the help of a sprite sheet data in a CSV format.

### List of punk-block registrations

A list of known punk blocks. The "L" in the size column means that this block 
will only be rendered on top of sarge base types, "S" is for small base types. 

If you would like to add your blocks to the list, please submit a PR.

| Name                 | Hash                                                               | Size | Layer |
|----------------------|--------------------------------------------------------------------|:----:|------:|
| Male 1               | 0x9039da071f773e85254cbd0f99efa70230c4c11d63fce84323db9eca8e8ef283 |  L   |     0 |
| Male 2               | 0xdfcbad4edd134a08c17026fc7af40e146af242a3412600cee7c0719d0ac42d53 |  L   |     0 |
| Male 3               | 0xed94d667f893279240c415151388f335b32027819fa6a4661afaacce342f4c54 |  L   |     0 |
| Male 4               | 0x1323f587f8837b162082b8d221e381c5e015d390305ce6be8ade3ff70e70446e |  L   |     0 |
| Female 1             | 0x1bb61a688fea4953cb586baa1eadb220020829a1e284be38d2ea8fb996dd7286 |  S   |     0 |
| Female 2             | 0x47cc6a8e17679da04a479e5d29625d737670c27b21f8ccfb334e6af61bf6885a |  S   |     0 |
| Female 3             | 0x80547b534287b04dc7e9afb751db65a7515fde92b8c2394ae341e3ae0955d519 |  S   |     0 |
| Female 4             | 0xc0c9e42e9d271c94b57d055fc963197e4c62d5933e371a7449ef5d59f26be00a |  S   |     0 |
| Zombie               | 0xf41cb73ce9ba5c1f594bcdfd56e2d14e42d2ecc23f0a4863835bdd4baacd8b72 |  L   |     0 |
| Ape                  | 0xb1ea1507d58429e4dfa3f444cd2e584ba8909c931969bbfb5f1e21e2ac8b758d |  L   |     0 |
| Alien                | 0x62223f0b03d25507f52a69efbbdbcfdc7579756a7a08a95a2f0e72ada31e32b8 |  L   |     0 |
| Rosy Cheeks          | 0x047228ad95cec16eb926f7cd21ac9cc9a3288d911a6c2917a24555eac7a2c0e2 | L,S  |     2 |
| Luxurious Beard      | 0xce1f93a7afe9aad7ebb13c0add89c79d42b5e9b1272fdd1573aac99fe5d860d0 |  L   |     6 |
| Clown Hair Green     | 0xbfac272e71cad64427175cd77d774a7884f98c7901ebc4909ada29d464c8981e | L,S  |     8 |
| Mohawk Dark          | 0xa71068a671b554f75b7cc31ce4f8d63c377f276333d11989e77bc4a9205b5e42 | L,S  |     8 |
| Cowboy Hat           | 0x9a132de8409f80845eaec43154ff43d7bd61df75e52d96b4ded0b64626e4c88a |  L   |     8 |
| Mustache             | 0xfca4c5f86ef326916536dfdae74031d6960e41e10d38c624294334c3833974e2 |  L   |     6 |
| Clown Nose           | 0x4483a654781ca58fa6ba3590c74c005bce612263e17c70445d6cd167e55e900b | L,S  |    12 |
| Cigarette            | 0x1885fe71e225eade934ab7040d533bd49efc5d66e8f2d4b5aa42477ae9892ec9 | L,S  |    11 |
| Nerd Glasses         | 0x7411db1fe7a50d41767858710dc8b8432ac0c4fd26503ba78d2ed17789ce4f72 | L,S  |    10 |
| Regular Shades       | 0xdd7231e98344a83b64e1ac7a07b39d2ecc2b21128681123a9030e17a12422527 | L,S  |    10 |
| Knitted Cap          | 0x24dd0364c2b2d0e6540c7deb5a0acf9177d47737a2bf41ca29b553eb69558ef9 | L,S  |     8 |
| Shadow Beard         | 0xea5efa009543229e434689349c866e4d254811928ae8a1320abb82a36d3be53f |  L   |     6 |
| Frown                | 0x2df03e79022dc10f7539f01da354ffe10da3ef91f1e18bc7fd096db00c381de8 |  L   |     1 |
| Cap Forward          | 0xf0ac7cf8c022008e16b983f22d22dae3a15b9b5abcc635bc5c20beb4d7c91800 |  L   |     8 |
| Goat                 | 0x8580e735d58252637afd6fef159c826c5e7e6a5dcf1fe2d8398b3bf92c376d42 |  L   |     6 |
| Mole                 | 0x041bf83549434251cc54c0632896c8d3176b48d06150048c1bce6b6102c4e90c | L,S  |     3 |
| Purple Hair          | 0x591f84c8a41edd0013624b89d5e6b96cd3b0c6f1e214d4ea13a35639412f07e6 |  L   |     8 |
| Small Shades         | 0x54917cb8cff2411930ac1b1d36a674f855c6b16c8662806266734b5f718a9890 |  L   |    10 |
| Shaved Head          | 0x274ae610f9d7dec1e425c54ad990e7d265ba95c4f84683be4333542088ecb8e7 |  L   |     8 |
| Classic Shades       | 0x6a400b1508bfd84ab2f4cb067d6d74dc46f74cdae7efd8b2a2d990c9f037e426 | L,S  |    10 |
| Vape                 | 0x3e6bc8fc06a569840c9490f8122e6b7f08a7598486649b64477b548602362516 | L,S  |    11 |
| Silver Chain         | 0x2c382a7f1f32a6a2d0e9b0d378cb95e3dad70fe6909ff13888fe2a250bd10bb0 | L,S  |     5 |
| Smile                | 0x8968ce85cb55abb5d9f6f678baeeb565638b6bad5d9be0ea2e703a34f4593566 |  L   |     1 |
| Big Shades           | 0xc3075202748482832362d1b854d8274a38bf56c5ad38d418e590f46113ff10b1 | L,S  |    10 |
| Mohawk Thin          | 0x971f7c3d5d14436a3b5ef2d658445ea527464a6409bd5f9a44f3d72e30d1eba8 | L,S  |     8 |
| Beanie               | 0x1f7b5107846b1e32944ccf8aedeaa871fc859506f51e7d12d6e9ad594a4d7619 |  L   |     8 |
| Cap                  | 0xd35b2735e5fcc86991c8501996742b3b8c35772d92b69859de58ddd3559be46c | L,S  |     8 |
| Clown Eyes Green     | 0x2004722753f61acb2cefde9b14d2c01c6bcb589d749b4ea616b4e47d83fdb056 | L,S  |     4 |
| Normal Beard Black   | 0x05a5afe13f23e20e6cebabae910a492c91f4b862c2e1a5822914be79ab519bd8 |  L   |     6 |
| Medical Mask         | 0xac5194b2986dd9939aedf83029a6e0a1d7d482eb00a5dafa05fc0aaa9b616582 | L,S  |    11 |
| Normal Beard         | 0xf94798c1aedb2dce1990e0dae94c15178ddd4229aff8031c9a5b7a77743a34d4 |  L   |     6 |
| VR                   | 0x15854f7a2b735373aa76722c01e2f289d8b18cb1a70575796be435e4ce55e57a | L,S  |    10 |
| Eye Patch            | 0xd91f640608a7c1b2b750276d97d603512a02f4b84ca13c875a585b12a24320c2 | L,S  |    10 |
| Wild Hair            | 0x6bb15b5e619a28950bae0eb6a03f13daea1b430ef5ded0c5606b335f5b077cda | L,S  |     8 |
| Top Hat              | 0x7a8b4abb14bfe7b505902c23a9c4e59e5a70c7daf6e28a5f83049c13142cde5e |  L   |     8 |
| Bandana              | 0x72efa89c7645580b2d0d03f51f1a2b64a425844a5cd69f1b3bb6609a4a06e47f | L,S  |     8 |
| Handlebars           | 0xfc1c0134d4441a1d7c81368f23d7dfcdeab3776687073c12af9d268e00d6c0a8 |  L   |     6 |
| Frumpy Hair          | 0x6ced067c29d04b367c1f3cb5e7721ad5a662f5e338ee3e10c7d64d9d109ed606 | L,S  |     8 |
| Crazy Hair           | 0x66a6c35fd6db8b93449f29befe26e2e4bcb09799d56216ada0ef901c53cf439f | L,S  |     8 |
| Police Cap           | 0x85c5daead3bc85feb0d62d1f185f82fdc2627bdbc7f1f2ffed1c721c6fcc4b4d |  L   |     8 |
| Buck Teeth           | 0x3d1f5637dfc56d4147818053fdcc0c0a35886121b7e4fc1a7cff584e4bb6414f |  L   |     1 |
| Do-rag               | 0x64b53b34ebe074820dbda2f80085c52f209d5eba6c783abdae0a19950f0787ec |  L   |     8 |
| Front Beard          | 0x833ca1b7f8f2ce28f7003fb78b72e259d5a484b13477ad8212edb844217225ac |  L   |     6 |
| Spots                | 0x44c2482a71c9d39dac1cf9a7daf6de80db79735c0042846cb9d47f85ccc3ba9b | L,S  |     3 |
| Big Beard            | 0x4acd7797c5821ccc56add3739a55bcfd4e4cfd72b30274ec6c156b6c1d9185eb |  L   |     6 |
| Vampire Hair         | 0xc0ac7bb45040825a6d9a997dc99a6ec94027d27133145018c0561b880ecdb389 |  L   |     8 |
| Peak Spike           | 0xa756817780c8e400f79cdd974270d70e0cd172aa662d7cf7c9fe0b63a4a71d95 |  L   |     8 |
| Chinstrap            | 0x71c5ce05a579f7a6bbc9fb7517851ae9394c8cb6e4fcad99245ce296b6a3c541 |  L   |     6 |
| Fedora               | 0x283597377fbec1d21fb9d58af5fa0c43990b1f7c2fc6168412ceb4837d9bf86c |  L   |     8 |
| Earring              | 0xbb1f372f67259011c2e9e7346c8a03a11f260853a1fe248ddd29540219788747 | L,S  |     7 |
| Horned Rim Glasses   | 0xd5de5c20969a9e22f93842ca4d65bac0c0387225cee45a944a14f03f9221fd4a | L,S  |    10 |
| Headband             | 0xb040fea53c68833d052aa3e7c8552b04390371501b9976c938d3bd8ec66e4734 | L,S  |     8 |
| Pipe                 | 0x74ca947c09f7b62348c4f3c81b91973356ec81529d6220ff891012154ce517c7 | L,S  |    11 |
| Messy Hair           | 0x30146eda149865d57c6ae9dac707d809120563fadb039d7bca3231041bea6b2e | L,S  |     8 |
| Front Beard Dark     | 0x8394d1b7af0d52a25908dc9123cc00aa0670debcac95a76c3e9a20dd6c7e7c23 |  L   |     6 |
| Hoodie               | 0xeb787e7727b2d8d912a02d9ad4c30c964b40f4cebe754bb4d3bfb09959565c91 |  L   |     8 |
| Gold Chain           | 0x6a36bcf4268827203e8a3f374b49c1ff69b62623e234e96858ff0f2d32fbf268 | L,S  |     5 |
| Muttonchops          | 0x2f237bd68c6e318a6d0aa26172032a8a73a5e0e968ad3d74ef1178e64d209b48 |  L   |     6 |
| Stringy Hair         | 0xad07511765ae4becdc5300c486c7806cd661840b0670d0f6670e8c4014de37b0 | L,S  |     8 |
| Eye Mask             | 0x49e0947b696384a658eeca7f5746ffbdd90a5f5526f8d15e6396056b7a0dc8af | L,S  |    10 |
| 3D Glasses           | 0xc1695b389d89c71dc7afd5111f17f6540b3a28261e4d2bf5631c1484f322fc68 | L,S  |    10 |
| Clown Eyes Blue      | 0x09c36cad1064f6107d2e3bef439f87a16c8ef2e95905a827b2ce7f111dd801d7 | L,S  |     4 |
| Mohawk               | 0xeb92e34266f6fa01c275db8379f6a521f15ab6f96297fe3266df2fe6b0e1422e | L,S  |     8 |
| Pilot Helmet         | 0x1892c4c9cf47baf2c613f184114519fe8208c2bebabb732405aeac1c3031dc2b |  S   |     8 |
| Tassle Hat           | 0x250be814c80d8ca10bbef531b679392db8221a6fab289a6b5e637df663f48699 |  S   |     8 |
| Hot Lipstick         | 0xcd87356aa78c4fcb95e51f57578570d377440e347e0869cf1b4749d5a26340b5 |  S   |     1 |
| Blue Eye Shadow      | 0x4fa682c6066fcc513a0511418aa85a0037ac59a899e9491c512b63e253697a8c |  S   |     4 |
| Straight Hair Dark   | 0x36f07f03014f047728880d9f390629140a5e7c44477290695c4c1ddda356d365 |  S   |     8 |
| Choker               | 0x68107f52c261820bd73e4046eb3fb5d5a1e0926611562c07054a3b89334cef34 |  S   |     5 |
| Wild Blonde          | 0xd395cf4acda004fbc9963f85c65bf3f190c2aceb0744a535d543bc261caf6ff0 |  S   |     8 |
| Wild White Hair      | 0xbad0fc475e9d35de67c426fc37eebb7fa38141bc2135fabd5504a911e1b05540 |  S   |     8 |
| Tiara                | 0xd10bc0475e2a0eea9f6aca91e6e82c6416f894f27fc26bb0735f29b84c54a3e6 |  S   |     8 |
| Orange Side          | 0xa0a2010e841ab7b343263c98f47a16b88656913e1353d96914f5fe492511893f |  S   |     8 |
| Red Mohawk           | 0x0e6769a10f786458ca82b57684746fe8899e35f7772543acb6a8869c4ac780cd |  S   |     8 |
| Purple Eye Shadow    | 0x1004d2d00ccf8794739c7b7cbbe6048841f4c8af046b37d59e9a801a167544e2 |  S   |     4 |
| Dark Hair            | 0x629e82a55845ea763431647fcaecfb232e275a36d8427f2568377864193801cb |  S   |     8 |
| Blonde Short         | 0xcd3633a5e96d615b834e90e67029f7f9f507b832e1cb263a29685b8e25f678cf |  S   |     8 |
| Purple Lipstick      | 0xe81a9c78c0ec4339dc6772f1b9bbf406b53063f8408a91fe29f63ba1c2bc7b5a |  S   |     1 |
| Pigtails             | 0xe11278d6c191c8199a5b8bb49be7f806b837a9811195c903d844a74c4c4a704e |  S   |     8 |
| Straight Hair Blonde | 0x411ec1566affa22bd67b13a7c49ac060c018e1c806cd314cd2186118dd55e129 |  S   |     8 |
| Welding Goggles      | 0x1868a04ecae06e10c5b6dcbbed4befac1ed03dda2cf86ddbd855466cc588809f |  S   |    10 |
| Pink With Hat        | 0x3511b04ac6a3d46305172269904dc469a40f380a4e7afa8742ce6e6a44825c4a |  S   |     8 |
| Blonde Bob           | 0x2857e47dcac3b744dd7d41617ce362f1dd3ae8eb836685cc18338714205b036c |  S   |     8 |
| Green Eye Shadow     | 0x2e9a5434da70e5ea2ed439b3a33aac60bd252c92698c1ba37e9ed77f975c6cab |  S   |     4 |
| Straight Hair        | 0x8c0e60b85ff0f8be1a87b28ae066a63dcc3c02589a213b0856321a73882515f9 |  S   |     8 |
| Half Shaved          | 0xe651be5dd43261e6e9c1098ec114ab5c44e7cb07377dc674336f1b3d34428fe4 |  S   |     8 |
| Black Lipstick       | 0x1cd064e6db4e7c5180ccf5f2afe1370c6539b525fe3bea9c358f24a7cbdb50ad |  S   |     1 |
| Stogie               | 0x398534927262d4f6993396751323ddd3e8326784a8e9a4808f17b99e6693835e | L,S  |    11 |
| Earphone             | 0x3b4d5e3dd66b09dd19cc19643986e2dc15e70251b31a4e5a463ecd996f7c3dc7 | L,S  |     9 |
| Employee Cap         | 0x550aa6da33a6eca427f83a70c2510cbc3c8bdb8a1ce5e5c3a32b2262f97c4aa1 | L,S  |     9 |
| Headphones           | 0xe2f3dcf809c00413a95bf007b46163923170ba8a0fbdaba7380f5c5079fcc98c | L,S  |     7 |
| Headphones Red       | 0x975e45b489dc6726c2a27eb784068ec791a22cf46fb780ced5e6b2083f32ebc3 | L,S  |     9 |
| Headphones Yellow    | 0x421c9c08478a3dfb8a098fbef56342e7e0b53239aaa40dd2d56951cc6c178d35 | L,S  |     9 |
| Gas Mask             | 0xaffb8a29fc5ed315e2a1103abc528d4f689c8365b54b17538f96e6bcae365633 | L,S  |    11 |
| Goggles              | 0x314ff09b8866e566e22c7bf1fe4227185bc37e1167a84aaf299f5e016ca2ea7b | L,S  |    10 |
| Pen                  | 0xe5fd4286f4fc4347131889d24238df4b5ba8d8d4985cbd9cb30d447ec14cbb2f | L,S  |     7 |
| Pencil               | 0xaeae7be74009ff61e63109240ea8e00b3bd6d166bf8a7f6584f64ff75e783f09 | L,S  |    10 |
| Red Hat              | 0x1cc630fd6d4fff8ca66aacb5acdba26a0a14ce5fd8f9cb60b002a153d1582b4e | L,S  |     8 |
| Yellow Hat           | 0xbbb91da98e74857ed34286d7efaf04751ac3f4d7081d62a0aa3b09278b5ee55a | L,S  |     8 |
| White Hat            | 0x3fbda43b0bda236b4f6f6dba8b7052381641b3d92ce4b49b4a2e9be390980019 | L,S  |     8 |
| Suit                 | 0x10214dd24c8822f95b3061229664e567e7da89d1f8a408179e12bf38be2c1430 | L,S  |     5 |
| Suit Black           | 0xb52fd5c8112bb81b2c05dd854ac28867bf72fd52124cb27aee3de68a19c87812 | L,S  |     5 |
| Bot                  | 0xd7a861eff7c9242c2fc79148cdb44128460adae80afe1ba79c2d1eae290fb883 | L,S  |     0 |
| Botina               | 0x7d3615eb6acf9ca19e31084888916f38df240bce4009857da690e4681bf8d4b0 | L,S  |     0 |
| Killer Bot           | 0x18a26173165d296055f2dfd8a12afc0a3e85434dd9d3f9c3ddd1eabc37ff56bc |  L   |     0 |
| Killer Botina        | 0xb93c33f3b6e2e6aef9bd03b9ed7a064ed00f8306c06dfc93c76ae30db7a3f2b4 |  S   |     0 |
| Green Alien          | 0x9242f3766d6363a612c9e88734e9c5667f4c82e07d00b794481f5b41b97047e8 |  L   |     0 |
| Green Alienette      | 0x0c924a70f72135432a52769f20962602647a5b6528675c14bb318eaf4cbb2753 |  S   |     0 |
| Blue Ape             | 0xcd6f6379578617fc2da9c1d778e731bebaa21e9be1ed7265963ec43076d17a10 |  L   |     0 |
| Alien 2              | 0x53f8bd0b36b2d3d9abc80e02d6fe9ed6a07068216cd737604c0c36ac60f458dc |  L   |     0 |
| Alien  3             | 0xeca5ecd41019c8240974e9473044bf1a01598e7c650939425f53f561e959ec46 |  L   |     0 |            
| Alien 4              | 0x061c5772160bfea6296a0317f6eff655398285ab18dbe89497436563445eeddc |  L   |     0 |
| Alien 5              | 0x224b0f8059a7c50a19036c71e7500fd115adfd3af915c8d6d6639248c6e41283 |  L   |     0 |
| Alien 6              | 0xfb3556140e6f92df2d04796b8d8c5f6732abf43c07eb7034a90672cd4f9af372 |  L   |     0 |
| Alienette 2          | 0xe9986a150e097f2cadc995279f34846ae9786b8ce35070b152f819d7a18d7760 |  S   |     0 |
| Alienette 3          | 0x0a215113c1e36c8cf69812b89dd912e3e2f1d70ab8c7691e0439a002d772f56d |  S   |     0 |
| Alienette 4          | 0xac4fc861f4029388de1fa709cb865f504fb3198a6bf4dad71ff705a436c406c2 |  S   |     0 |
| Alienette 5          | 0xbefcd0e4ecf58c1d5e2a435bef572fca90d5fcedf6e2e3c1eb2f12b664d555a4 |  S   |     0 |
| Alienette 6          | 0x54526cc56c302d9d091979753406975ad06ca6a58c7bea1395ae25350268ab36 |  S   |     0 |
| Pink Ape             | 0xffa2b3215eb937dd3ebe2fc73a7dd3baa1f18b9906d0f69acb3ae76b99130ff7 |  L   |     0 |
| Male 5               | 0x46151bb75270ac0d6c45f21c75823f7da7a0c0281ddede44d207e1242e0a83f6 |  L   |     0 |
| Male 6               | 0xef8998f2252b6977b3cc239953db2f5fbcd066a5d454652f5107c59239265884 |  L   |     0 |
| Male 7               | 0x606da1a8306113f266975d1d05f6deed98d3b6bf84674cc69c7b1963cdc3ea86 |  L   |     0 |
| Apelinah             | miss-aligned, unused                                               |  S   |     0 |
| Apette               | 0x804b2e3828825fc709d6d2db6078f393eafdcdedceae3bdb9b36e3c81630dd5e |  S   |     0 |
| Female 5             | 0x54354de4503fcf83c4214caefd1d4814c0eaf0ce462d1783be54ff9f952ec542 |  S   |     0 |
| Female 6             | 0x8a643536421eae5a22ba595625c8ba151b3cc48f2a4f86f9671f5c186b027ceb |  S   |     0 |
| Female 7             | 0x4426d573f2858ebb8043f7fa39e34d1441d9b4fa4a8a8aa2c0ec0c78e755df0e |  S   |     0 |
| Alientina            | 0x1908d72c46a0440b2cc449de243a20ac8ab3ab9a11c096f9c5abcb6de42c99e7 |  S   |     0 |
| Zombina              | 0xcedf32c147815fdc0d5f7e785f41a33dfc773e45bbd1a9a3b5d86c264e1b8ac5 |  S   |     0 |
| ZombieApe            | 0x691d9c552cd5457793c084f8bfce824df33aa7bcff69bb398b1c50c5283700ab |  L   |     0 |
| Cigarina             | 0x44cc2bd937a1ba84d91aa4ad1c68a4019d7441276f158686ca21113d9b58c736 |  S   |     0 |
| Cyborghina 1         | 0x6ad96c1daca4b1c9f05d375a8cc7561b56dc9f8e0c47de6294d0b56e99baba9f |  S   |     0 |
| Cyborghina 2         | 0x630cf72f7f662f0e4ad0e59518468203238cfd411fb9c5b474e65247043ff6ff |  S   |     0 |
| Cyborghina 3         | 0x9c4d52ffba9e3fe6a536e1420a71503203fde6d50cc7dfd6dcffb18520ea92ac |  S   |     0 |
| Cyborghina 4         | 0xa85374c4f65c797073c8536e4d19c56b86127fd476a9b5a4b3fbf026a0a631e9 |  S   |     0 |


Notes:

- Apelinah has a miss-alignment, `Apette` fixes it
- Headphones has an incorrect layer, it should be 9 not 7. `Earbuds`
  fixes this. They were also labelled incorrectly.
- Earphone was mislabeled, use `Headset` instead. The layer should be 7.

### List of layer ordering configurations

When using renderblocks.sol

| ID | Config                       | Comment   |
|----|------------------------------|-----------|
| 0  | 0,2,3,1,5,6,7,8,9,4,11,10,12 | Default   |


### License

MIT

Note: The MIT license is for the source code only. Images registered through
this contract retain their owner's rights, whatever they may be. This contract
is a non-profit "library" project and intended to archive & preserve punk
images, so that they can become widely accessible for decentralized apps,
including marketplaces, wallets, galleries, derivative works, research,
and other uses, as implied by fair-use or granted with permission from the
copyright owner.

# Demo

Here's a never-seen-before punk rendered using Punk Blocks.

![PunkBlocks](demo1.svg)
