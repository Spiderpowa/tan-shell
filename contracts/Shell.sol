pragma solidity ^0.5.0;

import 'openzeppelin-solidity/contracts/ownership/Ownable.sol';

contract Shell is Ownable {
    mapping(address => int) private _clients; 
    int private _clientCount;

    event ClientAdded(int indexed clientIndex, address indexed client);
    event Stdout(int indexed clientIndex, int indexed msgId, byte[] stream, bool eof);
    event Stderr(int indexed clientIndex, int indexed msgId, byte[] stream, bool eof);
    event Stdin(int indexed clientIndex, int indexed msgId, byte[] stream, bool eof);

    function addClient(address client) public onlyOwner {
        require(_clients[client]==0, "already registered");
        _clientCount += 1;
        _clients[client] = _clientCount;
        emit ClientAdded(_clientCount, client);
    }

    function stdout(int msgId, byte[] memory stream, bool eof) public {
        int idx = _clients[msg.sender];
        require(idx != 0, "unknown client");
        emit Stdout(idx, msgId, stream, eof);
    }

    function stderr(int msgId, byte[] memory stream, bool eof) public {
        int idx = _clients[msg.sender];
        require(idx != 0, "unknown client");
        emit Stderr(idx, msgId, stream, eof);
    }

    function stdin(int clientIndex, int msgId, byte[] memory stream, bool eof) public onlyOwner {
        require(clientIndex >= 0, "invalid client index");
        require(clientIndex <= _clientCount, "unknown client");
        emit Stdin(clientIndex, msgId, stream, eof);
    }
}
