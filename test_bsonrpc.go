package main

import (
    "gogobservice"
    "net/rpc"
    "bsonrpc"
)

func main() {
    
    rpc.Register(new(gogobservice.Arith))
    bsonrpc.ListenAndServe("localhost:9999")
}
