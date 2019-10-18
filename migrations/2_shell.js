const Shell = artifacts.require("Shell");

module.exports = function(deployer) {
  deployer.deploy(Shell);
};
