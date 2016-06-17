
package main


import (
    "fmt"
    "time"
    "launchpad.net/mgo/bson"
)


func main() {

    var bb []byte
    testdata := map[string]string{"aadfafdffsdfafda":"bbdfafdafasfas","ccdfafdsafaf":"dddfafdfa","fafdafafafafafdaffafafasfasfafafafa":"fafafafafa","ADFDFDFDF":"FDFDFDFDSAFEURIRITRTJMIOFJSIFASIHEIRHEREJIJFDSIFMSDFDJSFSJFIS"}
    t0 := time.Now()

    for i := 0; i< 1000000; i++ {
        bb,_ = bson.Marshal(testdata)
    }
    t1 := time.Now()
    fmt.Printf("marshal takes %v to run\n",t1.Sub(t0))

    var cc bson.M
    t0 = time.Now()
    for i :=0; i<1000000;i++ {
        bson.Unmarshal(bb,&cc)
    }
    t1 = time.Now()
    fmt.Printf("unmarshal takes %v to run\n",t1.Sub(t0))
    fmt.Printf("%#v\n",bb)
    fmt.Printf("%#v\n",cc)
}
