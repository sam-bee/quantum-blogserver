# Quantum Blogserver

A simple webserver written in Go. It uses Post-Quantum Cryptography, when compiled using Go 1.23rc2. It is for hosting
static sites.

## Compiling

`make build`

## Running

`./bin/quantumblogserver run --config ./config/.env`

## Configuration

The tool can be configured with a `.env` file, or with environment variables. See `[project-root]/config/.env.example
for values.

## Result

Your static site should be being served using the `X25519Kyber768Draft00` cyphersuite. You can see that in Chrome dev
tools here:

![image info](./docs/screenshot-kyber.png)

That means you are using Post-Quantum Cryptography with Go 1.23rc2.
