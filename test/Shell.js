const assert = require('assert');
const Shell = artifacts.require("Shell");

contract('Shell', function(accounts) {
  let ShellContract;

  it("should deploys Shell contract", async function() {
    ShellContract = await Shell.new();
  });

  it("should add a new client", async function() {
    await ShellContract.addClient(accounts[1]);
  });

  it("should emit event for stdout", async function() {
    const stream = ['0x01', '0x02', '0x03', '0x04'];
    const result = await ShellContract.stdout(13, stream, true, {from: accounts[1]})
    assert.equal(result.logs.length, 1);
    assert.equal(result.logs[0].event, 'Stdout');
    const args = result.logs[0].args;
    assert.equal(args.clientIndex, 1);
    assert.equal(args.msgId, 13);
    assert.deepEqual(args.stream, stream);
    assert.equal(args.eof, true);
  });

  it("should emit event for stderr", async function() {
    const stream = ['0x01', '0x02', '0x03', '0x04'];
    const result = await ShellContract.stderr(14, stream, true, {from: accounts[1]})
    assert.equal(result.logs.length, 1);
    assert.equal(result.logs[0].event, 'Stderr');
    const args = result.logs[0].args;
    assert.equal(args.clientIndex, 1);
    assert.equal(args.msgId, 14);
    assert.deepEqual(args.stream, stream);
    assert.equal(args.eof, true);
  });

  it("should emit event for stdin", async function() {
    const stream = ['0x01', '0x02', '0x03', '0x04'];
    const result = await ShellContract.stdin(1, 12, stream, true, {from: accounts[0]})
    assert.equal(result.logs.length, 1);
    assert.equal(result.logs[0].event, 'Stdin');
    const args = result.logs[0].args;
    assert.equal(args.clientIndex, 1);
    assert.equal(args.msgId, 12);
    assert.deepEqual(args.stream, stream);
    assert.equal(args.eof, true);
  });
  // negative tests
  it("should not allow adding duplicated clients", async function() {
    assert.rejects(() => {
      return ShellContract.addClient(accounts[1])
    }, {reason: "already registered"} );
  });

  it("should reject unregistered client", async function() {
    const stream = ['0x01', '0x02'];
    assert.rejects(() => {
      return ShellContract.stdout(1, stream, true, {from:accounts[2]});
    }, {reason: "unknown client"} );
    assert.rejects(() => {
      return ShellContract.stderr(2, stream, true, {from:accounts[2]});
    }, {reason: "unknown client"} );
  });

  it("should reject non owner using stdin", async function() {
    const stream = ['0x01', '0x02'];
    assert.rejects(() => {
      return ShellContract.stdin(1, 3, stream, true, {from:accounts[1]});
    }, {reason: "Ownable: caller is not the owner"});
  });

  it("should prevent owner inputting unknown client index", async function() {
    const stream = ['0x01', '0x02'];
    assert.rejects(() => {
      return ShellContract.stdin(0, 4, stream, true, {from:accounts[1]});
    }, {reason: "invalid client index"});
    assert.rejects(() => {
      return ShellContract.stdin(-1, 5, stream, true, {from:accounts[1]});
    }, {reason: "invalid client index"});
    assert.rejects(() => {
      return ShellContract.stdin(2, 6, stream, true, {from:accounts[1]});
    }, {reason: "unknown client"});
  });
});
