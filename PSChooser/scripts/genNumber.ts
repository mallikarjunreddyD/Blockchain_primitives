import { network, ethers } from "hardhat"


export async function genNumber() {
    console.log("Getting Random number  ................")
    const luckySeven = await ethers.getContract("PSChooser")
    
    for(let i=0;i<56;i++){
    const tx = await luckySeven.genNumber()
    const txReceipt = await tx.wait()
    const result = txReceipt.events[0].args.number;
    console.log(result.toNumber())
    }
}

genNumber()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })