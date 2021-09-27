package main

import "fmt"
import "os/exec"

func main(){
  hasil1,_:=exec.Command("ipconfig").Output()
  fmt.Println(string(hasil1))
  hasil2,_:=exec.Command("tree").Output()
  fmt.Println(string(hasil2))
  hasil3,_:=exec.Command("nslookup").Output()
  fmt.Println(string(hasil3))
}
