// scripts/deploy.js
async function main() {
    const [deployer] = await ethers.getSigners();
  
    console.log("Deploying contracts with the account:", deployer.address);
  
    const Storage = await ethers.getContractFactory("Storage");
    const storage = await Storage.deploy();
  
    console.log("Storage deployed to:", storage.address);
  }
  
  main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
  