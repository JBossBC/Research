abstract contract Proxy{
    //代理调用函数,calldatacopy(X,y,z)将calldata中y位置开始复制zbytes到memory中的x位置
    //returndatacopy(x,y,z)一样的道理,只是说，对于eth源码来看，returndata是单独存在的一个字段,这里的意思就是从y开始的returndata里面复制z长度大小给memory中
    //x位置开始的
    function _delegate(address implementation)internal virtual{
        assembly{
            calldatacopy(0,0,calldatasize())
            let result :=delegatecall(gas(),implementation,0,calldatasize(),0,0)
            returndatacopy(0,0,returndatasize())
            switch result
             case 0{
                revert(0,returndatasize())
             }
             default{
                return(0,returndatasize())
             }
        }
    }

    function _implementation()internal view virtual returns(address);

    function _fallback()internal virtual{
        _beforeFallback();
        _delegate(_implementation());
    }
    fallback()external payable virtual{
        _fallback();
    }
    function _beforeFallback()internal virtual{}
    
}