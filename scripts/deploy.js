const hre = require("hardhat");
// to run:
// $ npx hardhat run --network localhost scripts/deploy.js
async function main() {

    const PunkBlocks = await hre.ethers.getContractFactory("PunkBlocks");
    const blocks = await PunkBlocks.deploy();

    await blocks.deployed();

    const FP = await hre.ethers.getContractFactory("FactoryPunks");
    const fp = await FP.deploy(blocks.address);

    await fp.deployed();

    console.log(
        `Lock PunkBlocks deployed to ${blocks.address}`
    );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});