const bitcore = require('bitcore-lib');
require('dotenv').config(); 

const recipientPublicKeyHex = '023a914f291f8521a3fb42665012e750bc0b055f2cb914db4d02c9ab0d18964532';
const recipientPublicKey = new bitcore.PublicKey(recipientPublicKeyHex);

// Build a Pay-to-Public-Key (P2PK) locking script
const lockingScript = new bitcore.Script.buildPublicKeyOut(recipientPublicKey);


const network =  bitcore.Networks.testnet; // Change to 'mainnet' if you're on the main Bitcoin network
const senderPrivateKeyWIF = process.env.Private_Key; // Define Private_Key in .env

const senderPrivateKey = bitcore.PrivateKey.fromWIF(senderPrivateKeyWIF);

// Create a new transaction
const transaction = new bitcore.Transaction();

const utxoTxId = '3162154b3ea312dc6ce1695856636b0dc84b5a39c555818423d2407d94a48629';
const utxoOutputIndex = 0;

transaction.from({
    txId: utxoTxId,
    outputIndex: utxoOutputIndex,
    script: new bitcore.Script(lockingScript),
    satoshis: 1000 // Amount of satoshis in the UTXO
  });

  transaction.to(recipientPublicKey,700)
  transaction.fee(300);

  transaction.sign(senderPrivateKey);

  console.log('Raw Transaction:', transaction.serialize());

  



