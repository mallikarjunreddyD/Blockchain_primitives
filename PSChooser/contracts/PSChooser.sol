// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract PSChooser {
   uint[] public allNumbers;
   constructor() {
     for (uint i=1;i<=56;i++) {
        allNumbers.push(i);
     }
   }
   event RandomNumber(uint number);
    
    function genNumber() public {
        uint ranIndex = uint(keccak256(abi.encodePacked(blockhash(block.number), block.timestamp, block.difficulty))) % allNumbers.length;
        uint num = allNumbers[ranIndex];
        removeFromAllNumbers(ranIndex);
        emit RandomNumber(num);
    }

    function removeFromAllNumbers(uint i) internal {
        allNumbers[i] = allNumbers[allNumbers.length - 1];
        allNumbers.pop();
    }
}
