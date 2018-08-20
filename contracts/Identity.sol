pragma solidity ^0.4.22;

import { ECRecovery } from "./ECRecovery.sol";

contract IdentityFactory {
    mapping (address => address) accountToIdentity;
    
    function createIdentity(address _account, bytes32 _pubKey) public returns(bool){
        
        //Identity already exists
        if (accountToIdentity[_account] != address(0)) {
            return false;
        }
        
        address newIdentity = new Identity(_account);
        accountToIdentity[_account] = newIdentity;
        return true;
    }
    
    function getIdentityForAccount(address _account) external view returns(address) {
        return accountToIdentity[_account];   
    }
}


// borrowed heavily from 
// https://github.com/status-im/contracts/blob/73-economic-abstraction/contracts/identity/IdentityGasRelay.sol
contract Identity {
    bytes4 public constant CALL_PREFIX = bytes4(keccak256("callGasRelay(address,uint256,bytes32,uint256,uint256,address)"));
    bytes4 public constant APPROVEANDCALL_PREFIX = bytes4(keccak256("approveAndCallGasRelay(address,address,uint256,bytes32,uint256,uint256)"));

    mapping(address => uint8) public authorizedAccounts;

    uint256 public nonce;
    address gasToken;
    
    event Read();
    event Write();
    event Ping();
    event UnknownOperation();
    event Bootstrap();

    constructor(address _account) public {
        // TODO permissions are arbitrary at this point, make them useful in the future
        authorizedAccounts[_account] = 1;
    }

    function AuthorizeAccount(address _account, uint8 _permissions, bytes32 _messageHash, bytes _signedHash) external returns (bool) {
        address sender = ECRecovery.recover(_messageHash, _signedHash);
        uint8 permissions = authorizedAccounts[sender];
        require(permissions > 0, "account not authorized.");

        authorizedAccounts[_account] = _permissions;

        return false;
    }

    function RevokeAccount(address _account, bytes32 _messageHash, bytes _signedHash) external returns (bool) {
        address sender = ECRecovery.recover(_messageHash, _signedHash);
        uint8 permissions = authorizedAccounts[sender];
        require(permissions > 0, "account not authorized.");

        delete authorizedAccounts[_account];

        return true;
    }
    
    function RelayTransaction(
        address _to,
        uint256 _value,
        bytes _data,
        uint _nonce,
        uint _gasPrice,
        uint _gasLimit,
        bytes _signature
    ) external returns (bool) {
        //Verify the message sender is the owner

        bytes32 relayHash = callGasRelayHash(
            _to,
            _value,
            keccak256(_data),
            _nonce,
            _gasPrice,
            _gasLimit, 
            gasToken             
        );

        checkPermissions(relayHash, _signature);
        
        //executes transaction
        nonce++;
        bool success = _to.call.value(_value)(_data);

        return success;
    }

    function checkPermissions(bytes32 _relayHash, bytes _signature) private {
        // calculates signHash
        bytes32 signHash = getSignHash(_relayHash);
        address sender = ECRecovery.recover(signHash, _signature);
        uint8 permissions = authorizedAccounts[sender];
        require(permissions > 0, "account not authorized.");
    }

    /**
     * @notice get callHash
     * @param _to destination of call
     * @param _value call value (ether)
     * @param _dataHash call data hash
     * @param _nonce current identity nonce
     * @param _gasPrice price in SNT paid back to msg.sender for each gas unit used
     * @param _gasLimit minimal gasLimit required to execute this call
     * @param _gasToken token being used for paying `msg.sender` 
     * @return callGasRelayHash the hash to be signed by wallet
     */
    function callGasRelayHash(
        address _to,
        uint256 _value,
        bytes32 _dataHash,
        uint _nonce,
        uint256 _gasPrice,
        uint256 _gasLimit,
        address _gasToken
    )
        public 
        view 
        returns (bytes32 _callGasRelayHash) 
    {
        _callGasRelayHash = keccak256(
            address(this), 
            CALL_PREFIX, 
            _to,
            _value,
            _dataHash,
            _nonce,
            _gasPrice,
            _gasLimit,
            _gasToken
        );
    }

    
    /**
     * @notice get callHash
     * @param _to destination of call
     * @param _value call value (ether)
     * @param _dataHash call data hash
     * @param _nonce current identity nonce
     * @param _gasPrice price in SNT paid back to msg.sender for each gas unit used
     * @param _gasLimit minimal gasLimit required to execute this call
     * @param _gasToken token being used for paying `msg.sender` 
     * @return callGasRelayHash the hash to be signed by wallet
     */
    function approveAndCallGasRelayHash(
        address _baseToken,
        address _to,
        uint256 _value,
        bytes32 _dataHash,
        uint _nonce,
        uint256 _gasPrice,
        uint256 _gasLimit,
        address _gasToken
    )
        public 
        view 
        returns (bytes32 _callGasRelayHash) 
    {
        _callGasRelayHash = keccak256(
            address(this), 
            APPROVEANDCALL_PREFIX, 
            _baseToken,
            _to,
            _value,
            _dataHash,
            _nonce,
            _gasPrice,
            _gasLimit,
            _gasToken
        );
    }

    function getSignHash(
        bytes32 _hash
    )
        pure
        public
        returns(bytes32 signHash)
    {
        signHash = keccak256("\x19Ethereum Signed Message:\n32", _hash);
    }
}


//For testing verification
contract RecoveryVerifier {
    
    function verifySignature(address _account, bytes32 _pubKey, bytes32 _messageHash, bytes _signedHash) external pure returns (bool) {
        bytes32 prefixedHash = ECRecovery.toEthSignedMessageHash(_messageHash);
        return ECRecovery.recover(prefixedHash, _signedHash) == _account;
    }
    
    function verifySignatureNoPrefix(address _account, bytes32 _pubKey, bytes32 _messageHash, bytes _signedHash) external pure returns (bool) {
        return ECRecovery.recover(_messageHash, _signedHash) == _account;
    }

    function extractAccount(bytes32 _messageHash, bytes _signedHash) external pure returns (address) {
        bytes32 prefixedHash = ECRecovery.toEthSignedMessageHash(_messageHash);
        return ECRecovery.recover(prefixedHash, _signedHash);
    }

    function extractAccountNoPrefix(bytes32 _messageHash, bytes _signedHash) external pure returns (address) {
        return ECRecovery.recover(_messageHash, _signedHash);
    }
}
