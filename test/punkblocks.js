const {expect} = require("chai");
const {ContractFactory, utils, BigNumber} = require('ethers');

describe("PunkBlocks", function () {
    let PunkBlocks, pb, RenderBlocks, blocks;
    before(async function () {
        PunkBlocks = await ethers.getContractFactory("PunkBlocks");
        pb = await PunkBlocks.deploy();
        await pb.deployed();

        RenderBlocks = await ethers.getContractFactory("RenderBlocks");
        blocks = await RenderBlocks.deploy(pb.address);
        await blocks.deployed();

    });

    describe("FullTest", function () {

        // uncomment if running on mainnet to test svgFromPunkID
        it("Get SVG by Punk", async function () {


            svg = await blocks.svgFromPunkID(9987, 0, 0, 24, 0); // 8348
            console.log("original punk: "+svg);

        });


        it("Get SVGs", async function () {

            let attributes = ["Male 2", "Goat", "Smile", "Do-rag", "3D Glasses", "Rosy Cheeks", "Clown Eyes Green", "Pipe"];
            let svg = await blocks.svgFromNames(attributes, 0,0, 24, 0);
            console.log(svg);

            let getKey = function(s) {
                return ethers.utils.keccak256(ethers.utils.hexlify(ethers.utils.toUtf8Bytes((s))))
            }

            console.log("Vampire Hair", getKey("Vampire Hair"));

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
            svg = await blocks.svgFromKeys(attributes, 0, 0, 24, 0);
            console.log(svg);

            svg = await blocks.svgFromIDs([10, 71, 12], 0, 0, 24, 0);
            console.log(svg);

            // 45 is an m  only trait (Top Hat), it should not render on f
           svg = await blocks.svgFromIDs([77,5,45], 0, 0, 24, 0);

            console.log(svg);


        });

        it ("Test Solidity bug", async function() {
            let TestBlocking2 = await ethers.getContractFactory("TestBlocking2");
            let test = await TestBlocking2.deploy(blocks.address);
            await test.deployed();
            await test.doesntWork();

        });

        it ("Deploy Factory punks", async function() {
            let FactoryPunks = await ethers.getContractFactory("FactoryPunks");
            let fp = await FactoryPunks.deploy(blocks.address);
            await fp.deployed();

            let attributes = ["Suit Black", "Bot", "Yellow Hat", "Stogie"];
            let svg = await blocks.svgFromNames(attributes, 0, 0, 48, 0);
            console.log("Factory punk");
            console.log(svg);
            console.log("Factory end");
/*
            let info = await blocks._packInfo(2, 299, 2024);
            console.log("info:", info);
            let [layer, size1, size2] = await blocks._unpackInfo(info);
            console.log("layer, size1, size2, ", layer, size1, size2);
  */

            [layer, size1, size2] = await blocks.info("0x3fbda43b0bda236b4f6f6dba8b7052381641b3d92ce4b49b4a2e9be390980019");
            console.log("layer, size1, size2, ", layer, size1, size2);

        });

        const fromHexString = (hexString) =>
            Uint8Array.from(hexString.match(/.{1,2}/g).map((byte) => parseInt(byte, 16)));

        it("Create block", async function () {

            expect(await blocks.registerBlock(fromHexString("89504e470d0a1a0a0000000d4948445200000018000000180403000000125920cb00000015504c5445000000000000ff00008b532c5626007237094a1201cf76e6130000000174524e530040e6d8660000004c4944415478da62a03160141414807384949414e112ca4a4a4a302946255c1c2115272517384731484914c61154c26380102e19b5343807c5390c42082d208b0419905c2d80c901040000ffff2f3c090f8ffce8ac0000000049454e44ae426082"), new Uint8Array(0), 0, "Devil 1")).to.emit(blocks, "NewBlock");

            let attributes = ["Devil 1", "Goat", "Smile", "Do-rag", "3D Glasses", "Rosy Cheeks", "Clown Eyes Green", "Pipe"];
            let svg = await blocks.svgFromNames(attributes, 0, 0, 48, 0);
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
                "invalid L png"
            );

        });

        it("Grab blocks", async function () {
            let b = await blocks.getBlocks(0, 10);
            expect(b[0].length).to.equal(10);
           // console.log(b);
            expect(b[0][0]["blockL"]).to.equal(("0x89504e470d0a1a0a0000000d4948445200000018000000180403000000125920cb00000012504c5445000000000000713f1d8b532c5626007237092b4acd040000000174524e530040e6d8660000004f4944415478da62a00a1014141480b11995949414611c2165252525989490113247092747c549c945006698629092a800c264b8324674030489315a49118f3284ab9590fc23045783cc01040000ffffd8690b6ca3604b190000000049454e44ae426082"));
        });

        // registerOrderConfig
        it("test registerOrderConfig", async function () {
            await blocks.registerOrderConfig([0,2,3,1,5,6,7,8,9,4,11,10,12]);
            for (let i=0; i<13; i++) {
                console.log(await blocks.orderConfig(1, i));
            }

            let attributes = ["Female 1", "Big Shades", "Medical Mask"];
            let svg = await blocks.svgFromNames(attributes, 0, 0, 48, 0);
            console.log("Big Shades layer test");
            console.log(svg);
            console.log("END Big Shades layer test");

        });

    })


});

