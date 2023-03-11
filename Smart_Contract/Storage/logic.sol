pragma solidity ^0.8.0;

contract logic{
    uint256 private _value;
    address private _owner;
    constructor(address owner){
        _owner=owner;
    }
    modifier isOwner(address owner){
        require(owner==_owner,"is not owner");
        _;
    }
    function setValue(uint256 value)public isOwner(msg.sender) {
        _value=value*2;
    }
}