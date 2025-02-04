# Overview

This is basic, all in one, chat ui for ollama. For simplicity it is bundled with a go server and packaged into a all-in-one binary.
Simply build/download the binary and start chatting ...

# Usage

1. Start ollama
2. Start ollama-ui binary

# Development

```sh
# install dependencies
npm install
# run development web server
npm run dev
# build project
npm run build

# run tests
npm run test:unit
# lint project
npm run lint
```

To build the all in one go server binary do:

```sh
go generate
go build
```

The restulting binary is a all in one bundle (i.e. webserver + website). All you need to do is to start the binary and point your browser to http://localhost:3000

(the only other prerequisite is that ollama is running on your machine as well)
