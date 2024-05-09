// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Storage {
    uint256 private number;
    string private name;

    function setNumber(uint256 _number) public {
        number = _number;
    }

    function getNumber() public view returns (uint256) {
        return number;
    }

    function setName(string memory _name) public {
        name = _name;
    }

    function getName() public view returns (string memory) {
        return name;
    }
}
