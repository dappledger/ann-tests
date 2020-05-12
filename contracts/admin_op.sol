pragma solidity ^0.5.0;

contract Admin {
    function changenode(bytes memory txdata) public {

        bytes memory data = abi.encodePacked(msg.sender, txdata);
        uint size = data.length + 32;

        assembly {
            let success := call(not(0), 0xfe, 0, data, size, 0, 0)
            if iszero(success){
                revert(0, 0)
            }
        }
    }
}
