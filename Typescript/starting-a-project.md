# Starting a Typescript Project with Node.js
### Preparing the Environment
The first step is the preparation of the Environment, which consists of installing Node.js and NPM.
- [Installing Node.js and NPM](https://nodejs.org/en/download/package-manager/)

### Creating Typescript project with Node.js
In the project directory, we will start a default project using NPM.
```sh
$ npm init -y
```
Install Express, a Micro Framework to handle routes.
```sh
$ npm install express
```
Install the Express Type Definition Library as a Development dependency. It is not necessary that this dependency in a production environment, since in this case the code will already be converted from Typescript to Javascript.
```sh
$ npm install @types/express -D
```
Node only understands Javascript, so it is necessary to install `ts-node` to execute the code in Typescript.
```sh
$ npm install ts-node -D
```
Install Typescript as a Development dependency.
```sh
$ npm install typescript -D
```
Create and initialize the default Typescript configuration file.
```sh
$ npx tsc --init
```
Now use `npx` to run the application
```sh
$ npx ts-node {path}/{name}.ts
```
To avoid having to restart execution with each change, we can install `ts-node-dev` as a Development Dependency, which aims to "observe" the code and restart its execution automatically if there are any changes.
```sh
$ npm install ts-node-dev -D  
```

### Configuring a Script to improve usability.
In the *package.json* file where we have details of the project and its dependencies, we can configure Scripts to facilitate and improve processes, such as shortening command sets.
```json
{
  "name": "server",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "dev": "ts-node-dev src/server.ts"
  },
  "keywords": [],
  "author": "",
  "license": "ISC",
  "dependencies": {
    "express": "^4.17.1"
  },
  "devDependencies": {
    "@types/express": "^4.17.6",
    "ts-node": "^8.10.2",
    "ts-node-dev": "^1.0.0-pre.44",
    "typescript": "^3.9.3"
  }
}
```
In the scripts section, a script was called with the name: *dev* and the command: *ts-node-dev src / server.ts*.
In this way we shortened the command responsible for executing the project, to make use of the script we can now use only
```sh
$ npm run dev  
```
