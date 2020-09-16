# chargen-udp-go

## Purpose

The [**Character Generator Protocol (CHARGEN)**](https://en.wikipedia.org/wiki/Character_Generator_Protocol) is a network protocol for testing created in [1984](https://tools.ietf.org/html/rfc864). In the UDP implementation, whenever the client sends a datagram to the server it will get back another one containing a stream of 72 bytes of ASCII characters.

My goal is to learn more about UDP network programming in Go before dive into more complex protocols.

## Usage

### Running server

```
go run chargen.go
```

A typical CHARGEN UDP server listens for datagrams on [port 19](https://tools.ietf.org/html/rfc864). In this very example, I'm using `3019` so you don't have to run it as root. However, if you want to test it against a real service, change the bind/connect port using `--port` flag.

### Running client

```
nc -u -v localhost 3019
```
Hit <kbd>Enter</kbd> once session started.
