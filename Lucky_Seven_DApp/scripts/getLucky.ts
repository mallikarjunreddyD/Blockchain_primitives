import { network, ethers } from "hardhat"

const choice = 2;
export async function getLucky() {
    console.log("Getting lucky ................")
    const luckySeven = await ethers.getContract("luckyseven")
    const etherToSend = ethers.utils.parseEther("1"); // Sending 1 Ether

    const tx = await luckySeven.getLucky(choice,{value: etherToSend})
    const txReceipt = await tx.wait()
    const result = txReceipt.events[0].args.result;
    console.log(result.toNumber())
}

getLucky()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })