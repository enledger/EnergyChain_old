# EnergyChain

A blockchain platform to tokenize and track energy-grid-attached appliance production, usage, availablity,
share ownership, dividend and profit-sharing payouts, and renewable resource credits generation.

<b>EnLedger EnergyChain Project</b> is a collaborative cross-industry effort to stand up a platform for the
automated tracking of energy production, usage, availablity, share ownership, dividend and profit-sharing
payouts from grid-attached devices and appliances, including tracking generation and retirement of renewable
resource credits (RRCs) and device audit/registrations. EnergyChain is meant as a platform for energy projects
from the "microgrid" to the municipal and regional level, solar / renewable installation and auditing organizations,
and government organizations to all collaborate to use blockchains to achieve positive environmental impact and
efficient energy distribution. SOLES is also building real-time markets for true pricing of energy and renewable
resource credits tokenized into blockchains in the EnergyChain Project.

# node installation instructions

### 1) install golang and set $GOPATH (compiler & runtime environment setup)

install latest golang from source is the best way... (instructions to follow)

$GOPATH can be set by the following

export $GOPATH='~/.go'

export $PATH="$PATH:$GOPATH"

set these lines in your ~/.bashrc or ~/.bash_profile file

### 2) install tendermint (Proof-of-stake BFT consensus engine setup)

go get -u github.com/tendermint/tendermint/cmd/tendermint

see : https://tendermint.com/docs , https://tendermint.com/intro   (for background on Tendermint)

cd $GOPATH//src/github.com/tendermint/tendermint

make all

make install

### 3) install Energychain ABCI app

go get -d github.com/enledger/energychain

see : https://github.com/tendermint/cosmos-sdk (for cosmos-sdk / basecoin background & tutorials)

cd $GOPATH/src/github.com/enledger/energychain

make get_vendor_deps

make install

you should have an <b>energycoin</b> binary exectuable in $GOPATH/bin when this completes

### 4) running the node, creating a recieve address, sending a transaction

in one command console window, run energychain:

> energychain

in another command console window, run tendermint:

> tendermint

finally, in a third window, you can issue commands to energychain such as:

> energychain --createaddress

> energychain --sendtoaddress [amount] [address]










