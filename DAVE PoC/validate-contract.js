const fs = require('fs');
const solc = require('solc');

// Function to find all imports in a file
function findImports(path) {
  try {
    // Handle OpenZeppelin imports
    if (path.startsWith('@openzeppelin/')) {
      const npmPath = './node_modules/' + path;
      return { contents: fs.readFileSync(npmPath, 'utf8') };
    }
    
    // Handle relative imports
    if (path.startsWith('../../')) {
      const relativePath = './packages/contracts-bedrock/' + path.substring(6);
      return { contents: fs.readFileSync(relativePath, 'utf8') };
    }
    
    return { error: 'File not found: ' + path };
  } catch (e) {
    return { error: 'Error reading file: ' + path + ', ' + e.message };
  }
}

// Read the contract file
const contractPath = './packages/contracts-bedrock/src/dispute/TournamentGame.sol';
const contractSource = fs.readFileSync(contractPath, 'utf8');

// Prepare input for solc
const input = {
  language: 'Solidity',
  sources: {
    'TournamentGame.sol': {
      content: contractSource
    }
  },
  settings: {
    outputSelection: {
      '*': {
        '*': ['abi', 'evm.bytecode']
      }
    }
  }
};

// Compile the contract
const output = JSON.parse(solc.compile(JSON.stringify(input), { import: findImports }));

// Check for errors
if (output.errors) {
  console.error('Compilation errors:');
  output.errors.forEach(error => {
    console.error(error.formattedMessage);
  });
} else {
  console.log('Contract compiled successfully!');
}
