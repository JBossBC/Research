pragma solidity ^0.8.0;
import "./logic.sol";
contract testproxy{
      uint256 private value;
      address private _logic;
      string private  _SET_VALUE="setValue(uint256)";
      constructor(address logic){
          _logic=logic;
      }
    function getValue()public view returns(uint256){
        return value;
    }
    function setValue(uint256 value)public{
        (bool success,bytes memory returndata)=_logic.delegatecall(abi.encodeWithSignature(_SET_VALUE,value));
       if (!success){
         revert("hello");
      }

    }
    function setLogic(address proxy)public{
        require(proxy!=address(0),"!!!");
        // require(proxy.code)
        _logic=proxy;
    }
    // function setValue(uint256 value)public{
    //     _logic.setValue(value);
    // }
    // function storageData()public view returns(bytes memory data){
    //     assembly{
    //        data:=sload()
    //     }
    //     return data;
    // }
}