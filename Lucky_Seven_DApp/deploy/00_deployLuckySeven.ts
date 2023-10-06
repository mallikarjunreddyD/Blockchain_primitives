import { HardhatRuntimeEnvironment } from "hardhat/types"
import { DeployFunction } from "hardhat-deploy/types"
import { networkConfig, developmentChains } from "../helper-hardhat-config"

const deployluckyseven: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  // @ts-ignore
  const { getNamedAccounts, deployments, network } = hre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()
  log("----------------------------------------------------")
  log("Deploying luckyseven and waiting for confirmations...")
  const luckyseven = await deploy("luckyseven", {
    from: deployer,
    args: [],
    log: true,
    value: "10000000000000000000"
    // we need to wait if on a live network so we can verify properly
    //waitConfirmations: networkConfig[network.name].blockConfirmations || 1,
  })
  log(`luckyseven at ${luckyseven.address}`)
  }

export default deployluckyseven
deployluckyseven.tags = ["all", "luckyseven"]
