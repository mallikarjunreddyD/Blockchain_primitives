/**
 * A contract to demonstrate a decentralized Lucky seven game
 */

pragma solidity ^0.8.17;
import "../node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol";
contract luckyseven is ERC20("luckySevens","luckySevens"){
    enum choices {below_seven,seven,above_seven}
    
    constructor() payable {
        _mint(address(this),1000000000);
    }
    function mint() public {
        _mint(msg.sender,1000);
    }
    event output(choices userChoice, uint result);
    function getLucky(choices choice, uint bet) public {
        require(balanceOf(address(this)) >= 2*bet, "Not enough pool money");
        require(bet>0,"Atleast a token must be sent");
        uint result = (uint(keccak256(abi.encodePacked(blockhash(block.number-uint(choice)),block.timestamp, block.difficulty))) % 14) + 1;
        if(result==7 && choice==choices.seven || 
            result<7 && choice==choices.below_seven || 
            result>7 && choice==choices.above_seven ) {
            _transfer(address(this), msg.sender,2*bet);
        }
        else {
            _transfer(msg.sender,address(this),2*bet);
        }
        emit output(choice, result);
    }
}