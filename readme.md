# carauction
**carauction** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Creating an auction (for seller)
```
Format: auctiond tx auction create-auction "auction_name" initialPrice duration --from accountName
Example: auctiond tx auction create-auction "carAuction" 100token 100 --from alpha
```

### Placing a bid (for buyer)
```
Format: auctiond tx auction place-bid auctionID bidPrice --from accoutName
Example: auctiond tx auction place-bid 0 2000token --from alice
```
Notes:
- The new bid must be higher than the previous bid
- One can only place bid between creationTime and endTime

### Finalizing an auction
```
Format: auctiond tx auction finalize-auction auctionID --from accoutName
Example: auctiond tx auction finalize-auction 0 --from alpha
```
Notes:
- Only the creator can finalize the auction
- It can only be called when currentBlockHeight >= createAt (of auction) + duration (of auction)
- It will transfer the token to the highest bidder

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/car-auction@latest! | sudo bash
```
`username/car-auction` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
