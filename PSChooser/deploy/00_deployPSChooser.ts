import { HardhatRuntimeEnvironment } from "hardhat/types"
import { DeployFunction } from "hardhat-deploy/types"
import { networkConfig, developmentChains } from "../helper-hardhat-config"

const deployPSChooser: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  // @ts-ignore
  const { getNamedAccounts, deployments, network } = hre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()
  log("----------------------------------------------------")
  log("Deploying PSChooser and waiting for confirmations...")
  const PSChooser = await deploy("PSChooser", {
    from: deployer,
    args: [],
    log: true,
    // we need to wait if on a live network so we can verify properly
    //waitConfirmations: networkConfig[network.name].blockConfirmations || 1,
  })
  log(`PSChooser at ${PSChooser.address}`)
  }

export default deployPSChooser
deployPSChooser.tags = ["all", "luckyseven"]