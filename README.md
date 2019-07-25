# npm-lambda-layer

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
│   ├── npm
│   └── npx
└── nodejs
    └── node_modules
```

## Article

[Lambda上でnpm installできるLayerを作った - sambaiz-net](https://www.sambaiz.net/article/233/)