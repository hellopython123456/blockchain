package main
import (
        "encoding/binary"
        "io"
        "bytes"
)

func IntToByte(num int64)[]byte{
    var  buffer bytes.Buffer
    err := binary.Write(&buffer,binary.BigEndian,num)
    CheckErr("pos intobyte",err)
    return buffer.bytes()
}
func CheckErr(pos string, err error){
    if err != nil{
        fmt.Println("pos err",pos, err)
        os.Exit()
    }
}
