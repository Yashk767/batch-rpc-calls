require("@nomiclabs/hardhat-waffle"); // Make sure you have installed this plugin
require("@nomiclabs/hardhat-ethers");

module.exports = {
  solidity: "0.8.4", // Use the version that matches your Solidity version
  networks: {
    hardhat: {
      chainId: 1337
    },
    localhost: {
      url: "http://127.0.0.1:8545"
    }
  }
};
