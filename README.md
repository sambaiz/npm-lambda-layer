# lambda-npm-layer

Enable `npm` command execution in Lambda function.

```
$ make build
$ make local-run
$ open http://localhost:3000
```

```
$ tree -L 2 npm-layer
npm-layer
├── bin
│   ├── node
│   └── npm -> ../nodejs/node_modules/npm/bin/npm-cli.js
└── nodejs
    └── node_modules
```

## Article

[Lambda上でnpm installする - sambaiz-net](https://www.sambaiz.net/article/233/)