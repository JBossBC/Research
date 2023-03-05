pragma solidity ^0.8.0;
library Address{
  function isContract(address account)internal view returns(bool){
     return account.code.length>0;
  }
function sendValue(address payable recipient,uint256 amount)internal{
    require(address(this).balance>=amount,"Address: insufficient balance");
    (bool success,)=recipient.call{value:amount}("");
    require(success,"Address: unable to send value, recipient may have reverted");
}
function functionCall(address target,bytes memory data)internal returns(bytes memory){
    return functionCallWithValue(target,data,0,"Address: low-level call failed");
}
function functionCall(address target,bytes memory data,string memory errorMessage)internal returns(bytes memory){
    return functionCallWithValue(target,data,0,errorMessage);
}
function functionCallWithValue(address target,bytes memory data,uint256 value)internal returns(bytes memory){
    return functionCallWithValue(target,data,value,"Address: low-level call with value failed");
}
function functionCallWithValue(address target,bytes memory data,uint256 value,string memory errorMessage)internal returns(bytes memory){
    require(address(this).balance>=value,"Address: insuffcient balance for call");
    (bool success,bytes memory returndata) =target.call{value: value}(data);
    return verifyCallResultFromTarget(target,success,returndata,errorMessage);
}
function functionStaticCall(address target,bytes memory data)internal view returns(bytes memory){
    return functionStaticCall(target,data,"Address: low-level static call failed");
}
function functionStaticCall(address target,bytes memory data,string memory errorMessage)internal view returns(bytes memory){
    (bool success,bytes memory returndata)=target.staticcall(data);
    return verifyCallResultFromTarget(target,success,returndata,errorMessage);
}

function functionDelegateCall(address target,bytes memory data)internal returns(bytes memory){
    return functionDelegateCall(target,data,"Address: low-level delegate call failed");
}
function functionDelegateCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        (bool success, bytes memory returndata) = target.delegatecall(data);
        return verifyCallResultFromTarget(target, success, returndata, errorMessage);
    }
function verifyCallResultFromTarget(address target,bool success,bytes memory returndata,string memory errorMessage)internal view returns(bytes memory){
    if(success){
        if (returndata.length==0){
            require(isContract(target),"Address: call to non-contract");
        }
        return returndata;
    }else{
        _revert(returndata,errorMessage);
    }
}
function verifyCallResult(bool success,bytes memory returndata,string memory errorMessage)internal pure returns(bytes memory){
    if(success){
        return returndata;
    }else{
        _revert(returndata,errorMessage);
    }
}
function _revert(bytes memory returndata,string memory errorMessage)private pure{
    if(returndata.length>0){
        assembly{
            let returndata_size:=mload(returndata);
            //TODO
            revert(add(32,returndata),returndata_size);
        }
    }else{
        revert(message);
    }
}


}