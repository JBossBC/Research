pragma  solidity ^0.8.0;

import "./Proxy.sol";

//defend the slot conflict ,so the proxy cant include any state variable
contract BeaconProxy is Proxy,ERC1967Upgrade{
   constructor(address beacon,bytes memory data)payable{
     _upgradeBeaconToAndCall(beacon,data,false);
   }

}