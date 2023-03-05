

abstract contract ERC1967Upgrade{
    bytes32 private constant _ROLLBACK_SLOT = 0x4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143;
    bytes32 internal constant _IMPLEMENTATION_SLOT=0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
    event Upgraded(address indexed implementation);
    function _getImplementation()internal view returns(address){
        return StorageSlot.getAddressSlot(_IMPLEMENTATION_SLOT).value;
    }
    function _setImplementation(address newImplementation)private{
        require(Address.isContract(newImplementation),"ERC1967: new implementation is not a contract");
        StorageSlot.getAddressSlot(_IMPLEMENTATION_SLOT).value=newImplementation;
    }
    //逻辑层升级函数
    function _upgradeTo(address newImplementation)internal{
        _setImplementation(newImplementation);
        emit Upgraded(newImplementation);
    }
    //升级合约并调用函数
    function _upgradeToAndCall(address newImplementation,bytes memory data,bool forceCall)internal{
        _upgradeTo(newImplementation);
        if(data.length>0||forceCall){
            Address.functionDelegateCall(newImplementation,data);
        }
    }
    //TODO  UUPS
    function _upgradeToAndCallUUPS(address newImplementation,bytes memory data,bool forceCall)internal{
        if(StorageSlot.getBooleanSlot(_ROLLBACK_SLOT).value){
            _setImplementation(newImplementation);
        }else{
            try IERC1822Proxiable(newImplementation).proxiableUUID() returns(bytes32 slot){
                require(slot ==_IMPLEMENTATION_SLOT,"ERC1967Upgrade: unsupported proxiableUUID");
            }catch{
                revert("ERC1967Upgrade: new implementation is not UUPS");
            }
            _upgradeToAndCall(newImplementation, data, forceCall);
        }
    }

    bytes32 internal constant _ADMIN_SLOT = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
    event AdminChanged(address previousAdmin,address newAdmin);
    function _getAdmin()internal view returns(address){
        return StorageSlot.getAddressSlot(_ADMIN_SLOT).value;
    }
    function _setAdmin(address newAdmin)private{
        require(newAdmin!=address(0),"ERC1967: new admin is the zero address");
        storageSlot.getAddressSlot(_ADMIN_SLOT).value=newAdmin;
    }
    function _changeAdmin(address newAdmin)internal{
        emit AdminChanged(_getAdmin(),newAdmin);
        _setAdmin(newAdmin);
    }
    bytes32 internal constant _BEACON_SLOT = 0xa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50;

    event BeaconUpgraded(address indexed beacon);
    function _getBeacon()internal view returns (address){
        return StorageSlot.getAddressSlot(_BEACON_SLOT).value;
    }
    function _setBeacon(address newBeacon)private{
        require(Address.isContract(newBeacon),"ERC1967: new beacon is not a contract");
        require(Address.isContract(IBeacon(newBeacon).implemenation()),"ERC1967: beacon implementation is not a contract");
       StorageSlot.getAddressSlot(_BEACON_SLOT).value=newBeacon;
    }
    function _upgreadeBeaconToAndCall(address newBeacon,bytes memory data,bool forceCall)internal{
        _setBeacon(newBeacon);
        emit BeaconUpgraded(newBeacon);
        if (data.length>0 || forceCall){
            Address.functionDelegateCall(IBeacon(newBeacon).implemenation(),data);
        }
    }
}