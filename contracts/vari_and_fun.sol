pragma solidity ^0.5.0;

// Test special variables and functions
contract VariAndFun {

    address public admin;

    constructor () public {
        admin = msg.sender;
    }

    /**
     * Test blockhash
     *
     * Hash of the given block, only works for 256 most recent blocks excluding current.
     */
    function getBlockHash() public view returns (bytes32) {
        bytes32 blockHash = blockhash(block.number - 1);

        return blockHash;
    }

    /**
     * Test block.number
     *
     * Current block number
     */
    function getBlockNum() public view returns (uint256) {
        return block.number;
    }

    /**
     * Test block.timestamp and now
     *
     * Current block timestamp as seconds since unix epoch
     */
    function getBlockTime() public view returns (uint256, uint256, uint256) {
        return (block.timestamp, now, block.number);
    }

    /**
     * Test msg.data
     *
     * Complete calldata
     */
    function getMsgData() public pure returns (bytes memory) {
        return msg.data;
    }

    /**
     * Test msg.sender
     *
     * Sender of the message (current call)
     */
    function getMsgSender() public view returns (address) {
        return msg.sender;
    }

    /**
     * Test msg.sig
     *
     * First four bytes of the calldata (i.e. function identifier)
     */
    function getMsgSig() public pure returns (bytes4) {
        return msg.sig;
    }

    /**
     * Test tx.origin
     *
     * Sender of the transaction (full call chain)
     */
    function getOrigin() public view returns (address) {
        return tx.origin;
    }

    /**
     * Test assert
     *
     * Throws if the condition is not met - to be used for internal errors.
     */
    function testAssert() public view returns (bool) {
        assert(msg.sender == admin);
        return true;
    }

    /**
     * Test require
     *
     * Throws if the condition is not met - to be used for errors in inputs or external components.
     */
    function testRequire() public view returns (bool) {
        require(msg.sender == admin);
        return true;
    }

    /**
     * Test revert
     *
     * Abort execution and revert state changes
     */
    function testRevert() public view returns (bool) {
        if (msg.sender != admin) {
            revert();
        }

        return true;
    }

    /**
     * Test keccak256
     *
     * Compute the Ethereum-SHA-3 (Keccak-256) hash of the (tightly packed) arguments
     */
    function testKeccak256(bytes memory _data) public pure returns (bytes32) {
        return keccak256(_data);
    }

    /**
     * Test sha256
     *
     * Compute the SHA-256 hash of the (tightly packed) arguments
     */
    function testSHA256(bytes memory _data) public pure returns (bytes32) {
        return sha256(_data);
    }

    /**
     * Test ripemd160
     *
     * Compute RIPEMD-160 hash of the (tightly packed) arguments
     */
    function testRipemd160(bytes memory _data) public pure returns (bytes20) {
        return ripemd160(_data);
    }

    /**
     * Test ecrecover
     *
     * Recover the address associated with the public key from elliptic curve signature or return zero on error.
     */
    function testEcrecover(bytes32 _operationHash, bytes memory _signature) public pure returns (address) {
        require (_signature.length == 65);

        // We need to unpack the signature, which is given as an array of 65 bytes (like eth.sign)
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := mload(add(_signature, 32))
            s := mload(add(_signature, 64))
            v := and(mload(add(_signature, 65)), 255)
        }
        if (v < 27) {
            v += 27; // Ethereum versions are 27 or 28 as opposed to 0 or 1 which is submitted by some signing libs
        }
        return ecrecover(_operationHash, v, r, s);
    }

    /**
     * Test addmod
     *
     * Compute (x + y) % k where the addition is performed with arbitrary precision and does not wrap around at 2**256.
     */
    function testAddMod(uint _x, uint _y, uint _k) public pure returns (uint) {
        return addmod(_x, _y, _k);
    }

    /**
     * Test mulmod
     *
     * Compute (x * y) % k where the multiplication is performed with arbitrary precision and does not wrap around at 2**256.
     */
    function testMulMod(uint _x, uint _y, uint _k) public pure returns (uint) {
        return mulmod(_x, _y, _k);
    }

    /**
     * Test this
     *
     * The current contract, explicitly convertible to Address.
     */
    function getThis() public view returns (address) {
        return address(this);
    }

    address create2Addr;

    /**
     * Test create2
     *
     * create2(v, p, n, s)
     * Create new contract with code mem[p…(p+n)) at address
     * keccak256(0xff . this . s . keccak256(mem[p…(p+n))) and
     * send v wei and return the new address, where 0xff is a
     * 8 byte value, this is the current contract’s address as
     * a 20 byte value and s is a big-endian 256-bit value.
     */
    function create2Deploy(bytes memory _code, uint256 _salt) public {
        address addr;
        assembly {
            addr := create2(0, add(_code, 0x20), mload(_code), _salt)
            if iszero(extcodesize(addr)) {
                revert(0, 0)
            }
        }

        create2Addr = addr;
    }

    /**
     * Get create2Addr
     */
    function getContractAddr() public view returns (address) {
        return create2Addr;
    }

    /**
     * Test extcodesize
     *
     * Size of the code at address a.
     */
    function testExist(address _addr) public view returns(bool) {
        uint256 result;

        assembly {
            result := extcodesize(_addr)
        }

        return result > 0 ;
    }
}

contract VariAndFunOrigin {

    address admin;

    /**
     * Test tx.origin
     *
     * Sender of the transaction (full call chain)
     */
    function getOrigin(address _addr) public view returns (address) {
        VariAndFun vf = VariAndFun(_addr);
        return vf.getOrigin();
    }

    /**
     * Test selfdestruct
     *
     * Destroy the current contract, sending its funds to the given Address.
     */
    function testSelfdestruct() public {
        selfdestruct(msg.sender);
    }
}