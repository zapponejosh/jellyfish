# Project Jellyfish 🌊

## Installing / Developing

```shell
npm install && cd web && npm install
```

Since NPM workspaces were not working as expected we are installing dependencies in each location.

```shell
npm run dev
```

This will start the Go server and React app. Currently Storybook is not included.

```shell
cd web
npm run storybook
```

The default PORTS are:

- `3001` for the server
- `3000` for the client
- `6006` for the storybook

You can configure the server port by setting the `PORT` environment variable. Creating a `.env` file is supported. You can copy `.env.example` to `.env`.

## Building

To build the project, run:

```shell
npm run build
```

This will build the client, server and storybook.

```shell
npm start
```

In production, you have a single server serving everything.

`/api/*` is the API endpoint.  
`/storybook` is the Storybook.  
`/*` is the client.

## Tests

TypeScript, linter and prettier are checked on commit and push thanks to husky and lintstaged.

## Licensing

MIT
