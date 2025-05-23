const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Tournament", function () {
  let tournament;
  let owner;
  let challenger1;
  let challenger2;
  let challenger3;
  
  const rootClaim = ethers.keccak256(ethers.toUtf8Bytes("Root claim"));
  const l1Head = ethers.keccak256(ethers.toUtf8Bytes("L1 head"));
  const extraData = ethers.toUtf8Bytes("Extra data");
  const BOND_AMOUNT = ethers.parseEther("0.1");
  
  beforeEach(async function () {
    [owner, challenger1, challenger2, challenger3] = await ethers.getSigners();
    
    const Tournament = await ethers.getContractFactory("Tournament");
    tournament = await Tournament.deploy(rootClaim, l1Head, extraData);
  });
  
  describe("Initialization", function () {
    it("Should set the correct initial values", async function () {
      expect(await tournament.rootClaim()).to.equal(rootClaim);
      expect(await tournament.l1Head()).to.equal(l1Head);
      expect(await tournament.gameCreator()).to.equal(owner.address);
      expect(await tournament.status()).to.equal(0); // IN_PROGRESS
      expect(await tournament.gameType()).to.equal(3); // DAVE Tournament type
    });
  });
  
  describe("Joining Tournament", function () {
    it("Should allow a challenger to join with sufficient bond", async function () {
      const claim = ethers.keccak256(ethers.toUtf8Bytes("Challenger 1 claim"));
      
      await expect(tournament.connect(challenger1).joinTournament(claim, { value: BOND_AMOUNT }))
        .to.emit(tournament, "ParticipantJoined")
        .withArgs(challenger1.address, 1, claim);
      
      // Check that the challenger was added correctly
      const node = await tournament.nodes(1);
      expect(node.participant).to.equal(challenger1.address);
      expect(node.claim).to.equal(claim);
      expect(node.hasJoined).to.be.true;
      expect(node.bondAmount).to.equal(BOND_AMOUNT);
    });
    
    it("Should revert if bond amount is insufficient", async function () {
      const claim = ethers.keccak256(ethers.toUtf8Bytes("Challenger 1 claim"));
      
      await expect(tournament.connect(challenger1).joinTournament(claim, { value: ethers.parseEther("0.05") }))
        .to.be.revertedWithCustomError(tournament, "InsufficientBond");
    });
    
    it("Should revert if participant has already joined", async function () {
      const claim1 = ethers.keccak256(ethers.toUtf8Bytes("Challenger 1 claim 1"));
      const claim2 = ethers.keccak256(ethers.toUtf8Bytes("Challenger 1 claim 2"));
      
      await tournament.connect(challenger1).joinTournament(claim1, { value: BOND_AMOUNT });
      
      await expect(tournament.connect(challenger1).joinTournament(claim2, { value: BOND_AMOUNT }))
        .to.be.revertedWithCustomError(tournament, "AlreadyParticipated");
    });
  });
  
  describe("Match Creation", function () {
    it("Should create matches when power of 2 participants join", async function () {
      // Add 3 challengers to make a total of 4 participants (including the defender)
      const claim1 = ethers.keccak256(ethers.toUtf8Bytes("Challenger 1 claim"));
      const claim2 = ethers.keccak256(ethers.toUtf8Bytes("Challenger 2 claim"));
      const claim3 = ethers.keccak256(ethers.toUtf8Bytes("Challenger 3 claim"));
      
      await tournament.connect(challenger1).joinTournament(claim1, { value: BOND_AMOUNT });
      await tournament.connect(challenger2).joinTournament(claim2, { value: BOND_AMOUNT });
      
      // This should trigger match creation (4 participants = 2^2)
      await expect(tournament.connect(challenger3).joinTournament(claim3, { value: BOND_AMOUNT }))
        .to.emit(tournament, "MatchCreated");
      
      // Check that matches were created
      const match0 = await tournament.matches(0);
      expect(match0.nodeA).to.equal(0); // Defender
      expect(match0.nodeB).to.equal(1); // Challenger 1
      expect(match0.resolved).to.be.false;
      
      const match1 = await tournament.matches(1);
      expect(match1.nodeA).to.equal(2); // Challenger 2
      expect(match1.nodeB).to.equal(3); // Challenger 3
      expect(match1.resolved).to.be.false;
    });
  });
  
  // Additional tests for match resolution and tournament completion would go here
  // but they would require time manipulation which is more complex
});
