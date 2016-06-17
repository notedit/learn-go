
package main


import (
    "fmt"
    "time"
    "github.com/edsrzf/tnetstring"
)


func main() {

    var bb string
    testdata := map[string]string{"aadfafdffsdfafda":"bbdfafdafasfas","ccdfafdsafaf":"dddfafdfa","fafdafafafafafafafafasfasfafafafa":"fafafafafa","ADFDFDFDF":"FDFDFDFDSAFEURIRITRTJMIOFJSIFASIHEIRHEREJIJFDSIFMSDFDJSFSJFIS"}
    t0 := time.Now()

    for i := 0; i< 1000000; i++ {
        bb,_ = tnetstring.Marshal(testdata)
    }
    t1 := time.Now()
    fmt.Printf("marshal takes %v to run\n",t1.Sub(t0))

    var cc map[string]string
    t0 = time.Now()
    for i :=0; i<1000000;i++ {
        tnetstring.Unmarshal(bb,&cc)
    }
    t1 = time.Now()
    fmt.Printf("unmarshal takes %v to run\n",t1.Sub(t0))
    fmt.Printf("%#v\n",bb)
    fmt.Printf("%#v\n",cc)
}
