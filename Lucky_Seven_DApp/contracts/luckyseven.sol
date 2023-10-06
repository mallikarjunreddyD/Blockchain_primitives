/**
 * A contract to demonstrate a decentralized Lucky seven game
 */

pragma solidity ^0.8.17;

contract luckyseven {
    enum choices {below_seven,seven,above_seven}
    uint public poolMoney=0;
    constructor() payable {
        poolMoney = msg.value;
    }
    event output(choices userChoice, uint result);
    function getLucky(choices choice) public payable {
        require(poolMoney >= 2*msg.value, "Not enough pool money");
        require(msg.value>0,"Atleast a wei must be sent");
        uint result = (uint(keccak256(abi.encodePacked(blockhash(block.number-uint(choice)),block.timestamp, block.difficulty))) % 14) + 1;
        if(result==7 && choice==choices.seven || 
            result<7 && choice==choices.below_seven || 
            result>7 && choice==choices.above_seven ) {
            poolMoney = poolMoney - 2*msg.value;
            payable(msg.sender).transfer(2*msg.value);
        }
        else {
            poolMoney = poolMoney + msg.value;
        }
        emit output(choice, result);
    }
}