
pragma solidity ^0.8.0;

interface IBeacon{
    function implementation()external view returns(address);
}