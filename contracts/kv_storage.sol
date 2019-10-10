pragma solidity ^0.4.25;

contract KvStorage {

    mapping(bytes32 => uint256) db;

    uint256 public calledCount;

    function set(bytes32 key, uint256 value) public returns (uint) {
        db[key] = value;
        calledCount++;
        return calledCount;
    }

    function get(bytes32 key) view public returns (uint){
        return db[key];
    }

    function batchSet(bytes32[] keys, uint256[] values) public {
        for (uint i = 0; i < keys.length; i++) {
            db[keys[i]] = values[i];
        }
    }

    function batchGet(bytes32[] keys) view public returns (uint[] memory){

        uint[] memory values = new uint[](keys.length);
        for (uint i = 0; i < keys.length; i ++) {
            values[i] = db[keys[i]];
        }
        return values;
    }
}

contract WrapKvStorage {
    KvStorage kv;

    constructor (KvStorage _addr) public {
        kv = _addr;
    }

    function set(bytes32 key, uint256 value) public {
        kv.set(key, value);
    }

    function get(bytes32 key) view public returns (uint){
        return kv.get(key);
    }
}
