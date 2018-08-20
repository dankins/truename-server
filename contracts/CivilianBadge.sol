pragma solidity ^0.4.22;

contract CivilianBadge {
    
    uint nextId;
    mapping (address => uint) civilians;


    function becomeCivilian() public {
        nextId++;
        civilians[msg.sender] = nextId;
    }

    function getID(address _civilian) public view returns (uint)  {
        return civilians[_civilian];
    }
}