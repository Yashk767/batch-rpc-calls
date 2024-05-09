// scripts/interactWithContract.js
async function main() {
    const [owner] = await ethers.getSigners();
    const Storage = await ethers.getContractFactory("Storage");
    const storage = await Storage.attach("0x5FbDB2315678afecb367f032d93F642f64180aa3");

    // Assuming `setNumber` is a function in your Storage contract
    const setTx = await storage.setNumber(47);
    await setTx.wait();

    const number = await storage.getNumber();
    console.log(`Number: ${number}`);

    const setNameTx = await storage.setName("ABC");
    await setNameTx.wait();

    const name = await storage.getName();
    console.log(`Name: ${name}`);
}

main().catch((error) => {
    console.error(error);
    process.exit(1);
});
