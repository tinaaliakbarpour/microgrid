# microgrid
**microgrid** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

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
curl https://get.ignite.com/tinaaliakbarpour/microgrid@latest! | sudo bash
```
`tinaaliakbarpour/microgrid` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

## IoT Module
first create the chain : 
```
ignite scaffold chain microgrid --no-module
```
then we specify our module called iot:

```
ignite scaffold module iot --dep account
```

we have to specify our major objects called device and grid:

```
ignite scaffold list grid name:string center-lat:int center-lon:int  side:uint admins:array.string --no-message --module iot 

ignite scaffold map device address:string lat:int lon:int power:uint voltage:uint others:string --index grid-id:uint --no-message --module iot

```

and now we have to create our messages to implement crud and logic on this objects.

```
ignite scaffold message create-grid name center-lat:int center-lon:int side:uint --response id:uint --module iot
ignite scaffold message register-admin id:uint address:string  --response grid:Grid --module iot
ignite scaffold message delete-grid id:uint --module iot 

```

now we go through device messages:

```

ignite scaffold message create-device grid-id:uint lat:int lon:int voltage:uint power:uint others:string  --module iot

ignite scaffold message update-device-status voltage:uint power:uint others:string addres:string grid-id:uint  --module iot

ignite scaffold message delete-device grid-id:uint address:string --module iot


```

and also added a query for listing all the devices in the same grid : 

```
ignite scaffold query list-device-by-grid-id grid-id:uint --response device:Device --module iot

```

and also you can see the document [here](https://docs.google.com/document/d/11916K1DdO6JHwrsB_PRxr7jzA6shbBAGrd8yPt_ZWJs/edit?usp=sharing)


for the testing purpose you can use following commands :

```
microgridd tx iot create-grid name 369734 378492 1000 --from bob 
microgridd tx iot create-device 0 369888 368999 100 2000 heloo --from alice
microgridd tx iot update-device-status 80 1000 hello cosmos1djlnm3sq5rjt7f2wfs5g4syxwhn87k3f6pyj5v 0 --from bob
microgridd q iot list-device 
microgridd q iot list-device-by-grid-id 0 
microgridd q iot show-device 0 cosmos1g5h594gw6mhgguyc4u56yu8qt5ec8vjv06zppt
```


and you can always use --h option for further information.