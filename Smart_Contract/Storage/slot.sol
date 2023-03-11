pragma solidity ^0.8.0;

import "./ShortStrings.sol";

contract Slot{
    struct addressSlot{
        address value;
    }
    mapping(uint256=> uint256)private mapLocation;
    uint256  private location; 
    addressSlot private   test;
    constructor(){
        mapLocation[1]=1;
        test.value=msg.sender;

    } 
   function testSlotLocation()public view returns(string memory){
       bytes memory temp ="0x133322222";
       bytes32 value= bytes32(uint256(bytes32(temp)) | temp.length);
     uint256 len =value.length;
       string memory result =new string(32);
        assembly{
         mstore(result,len)
         mstore(add(result,0x20),value)
     }
    //  result = "xiyang";
     return result;
   }
   function testtoString()public view returns(string memory){
        return ShortStrings.toString(ShortStrings.toShortString("123"));
         
   }
   function testStringMmeory()public pure returns(string memory){
       string memory s = "xxxxx";
       bytes32  bytess="xxxxx";
       string memory one =new string(32);
       string memory two =new string(32);
       assembly{
          let  pointer:=mload(s)
           mstore(one,pointer)
           mstore(add(one,0x20), bytess)
       }
       return one;
   }
}