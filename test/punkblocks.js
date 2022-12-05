const {expect} = require("chai");
const {ContractFactory, utils, BigNumber} = require('ethers');

describe("PunkBlocks", function () {
    let PunkBlocks, blocks;
    before(async function () {
        PunkBlocks = await ethers.getContractFactory("PunkBlocks");
        blocks = await PunkBlocks.deploy();
        await blocks.deployed();
    });

    describe("FullTest", function () {

        // uncomment if running on mainnet to test svgFromPunkID
        it("Get SVG by Punk", async function () {


            svg = await blocks.svgFromPunkID(8348); // 8348
            console.log("original punk: "+svg);

        });


        it("Get SVGs", async function () {

            let attributes = ["Male 2", "Goat", "Smile", "Do-rag", "3D Glasses", "Rosy Cheeks", "Clown Eyes Green", "Pipe"];
            let svg = await blocks.svgFromNames(attributes);
            console.log(svg);

            let getKey = function(s) {
                return ethers.utils.keccak256(ethers.utils.hexlify(ethers.utils.toUtf8Bytes((s))))
            }

            attributes = [
                getKey("Male 2"),
                getKey("Goat"),
                getKey("Smile"),
                getKey("Do-rag"),
                getKey("3D Glasses"),
                getKey("Rosy Cheeks"),
                getKey("Clown Eyes Green"),
                getKey("Pipe")
            ];
            svg = await blocks.svgFromKeys(attributes);
            console.log(svg);

            svg = await blocks.svgFromIDs([10, 71, 12]);
            console.log(svg);

            // 45 is an m  only trait (Top Hat), it should not render on f
           svg = await blocks.svgFromIDs([77,5,45]);

            console.log(svg);


        });

        it("Create block", async function () {
            const fromHexString = (hexString) =>
                Uint8Array.from(hexString.match(/.{1,2}/g).map((byte) => parseInt(byte, 16)));
            expect(await blocks.registerBlock(fromHexString("89504e470d0a1a0a0000000d4948445200000018000000180403000000125920cb00000015504c5445000000000000ff00008b532c5626007237094a1201cf76e6130000000174524e530040e6d8660000004c4944415478da62a03160141414807384949414e112ca4a4a4a302946255c1c2115272517384731484914c61154c26380102e19b5343807c5390c42082d208b0419905c2d80c901040000ffff2f3c090f8ffce8ac0000000049454e44ae426082"), new Uint8Array(0), 0, "Devil 1")).to.emit(blocks, "NewBlock");

            let attributes = ["Devil 1", "Goat", "Smile", "Do-rag", "3D Glasses", "Rosy Cheeks", "Clown Eyes Green", "Pipe"];
            let svg = await blocks.svgFromNames(attributes);
            console.log(svg);

            // slot taken if you try to upload the same attr name
            await expect( blocks.registerBlock(fromHexString("89504e470d0a1a0a0000000d4948445200000018000000180403000000125920cb00000015504c5445000000000000ff00008b532c5626007237094a1201cf76e6130000000174524e530040e6d8660000004c4944415478da62a03160141414807384949414e112ca4a4a4a302946255c1c2115272517384731484914c61154c26380102e19b5343807c5390c42082d208b0419905c2d80c901040000ffff2f3c090f8ffce8ac0000000049454e44ae426082"), new Uint8Array(0), 0, "Devil 1")).to.be.revertedWith(
                "slot taken"
            );

            // we cannot add a layer 0 block with both data fields

            await expect( blocks.registerBlock(
                fromHexString("89504e470d0a1a0a0000000d4948445200000018000000180403000000125920cb00000015504c5445000000000000ff00008b532c5626007237094a1201cf76e6130000000174524e530040e6d8660000004c4944415478da62a03160141414807384949414e112ca4a4a4a302946255c1c2115272517384731484914c61154c26380102e19b5343807c5390c42082d208b0419905c2d80c901040000ffff2f3c090f8ffce8ac0000000049454e44ae426082"),
                fromHexString("89504e470d0a1a0a0000000d4948445200000018000000180403000000125920cb00000015504c5445000000000000ff00008b532c5626007237094a1201cf76e6130000000174524e530040e6d8660000004c4944415478da62a03160141414807384949414e112ca4a4a4a302946255c1c2115272517384731484914c61154c26380102e19b5343807c5390c42082d208b0419905c2d80c901040000ffff2f3c090f8ffce8ac0000000049454e44ae426082"), 0, "Devil 2")).to.be.revertedWith('layer0 cannot have both versions');

            // invalid png
            await expect( blocks.registerBlock(fromHexString("19504e470d0a1a0a0000000d4948445200000018000000180403000000125920cb00000015504c5445000000000000ff00008b532c5626007237094a1201cf76e6130000000174524e530040e6d8660000004c4944415478da62a03160141414807384949414e112ca4a4a4a302946255c1c2115272517384731484914c61154c26380102e19b5343807c5390c42082d208b0419905c2d80c901040000ffff2f3c090f8ffce8ac0000000049454e44ae426082"), new Uint8Array(0), 0, "Devil 4")).to.be.revertedWith(
                "invalid m png"
            );

        });


    })


});

