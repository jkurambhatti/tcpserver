### IMPLEMENT TCP SERVERS USING only `net` package of "go" std library

# run chat server

```
$ go run chatserver.go

```

# run redis like in-memory key:value database server

```
$ go run keyvaluedb.go

```

## usage

use telnet to create connection to server 

```
$ telnet localhost 8080
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
```
execute GET ,SET ,DEL commands
```
SET name jayesh
OK
GET name
jayesh
DEL name
OK
GET name
no record found
```



```
