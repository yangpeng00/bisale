// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "context"
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "user"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "   selectUserByConditions(TUserParams params)")
  fmt.Fprintln(os.Stderr, "  i32 selectUserCountByConditions(TUserParams params)")
  fmt.Fprintln(os.Stderr, "  bool updateUserStatusByUserId(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "  TUserInfo selectUserBaseInfoByUserId(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "  bool resetGoogleCode(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "  TUser selectUserById(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "  TGoogleStatusResult selectUserGoogleStatus(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "  i32 selectSlaveAllUserCount(string traceId)")
  fmt.Fprintln(os.Stderr, "   selectSlaveRegisterCountDay(string traceId, i32 days)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := user.NewTUserServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "selectUserByConditions":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SelectUserByConditions requires 1 args")
      flag.Usage()
    }
    arg22 := flag.Arg(1)
    mbTrans23 := thrift.NewTMemoryBufferLen(len(arg22))
    defer mbTrans23.Close()
    _, err24 := mbTrans23.WriteString(arg22)
    if err24 != nil {
      Usage()
      return
    }
    factory25 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt26 := factory25.GetProtocol(mbTrans23)
    argvalue0 := user.NewTUserParams()
    err27 := argvalue0.Read(jsProt26)
    if err27 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.SelectUserByConditions(context.Background(), value0))
    fmt.Print("\n")
    break
  case "selectUserCountByConditions":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SelectUserCountByConditions requires 1 args")
      flag.Usage()
    }
    arg28 := flag.Arg(1)
    mbTrans29 := thrift.NewTMemoryBufferLen(len(arg28))
    defer mbTrans29.Close()
    _, err30 := mbTrans29.WriteString(arg28)
    if err30 != nil {
      Usage()
      return
    }
    factory31 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt32 := factory31.GetProtocol(mbTrans29)
    argvalue0 := user.NewTUserParams()
    err33 := argvalue0.Read(jsProt32)
    if err33 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.SelectUserCountByConditions(context.Background(), value0))
    fmt.Print("\n")
    break
  case "updateUserStatusByUserId":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "UpdateUserStatusByUserId requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err35 := (strconv.Atoi(flag.Arg(2)))
    if err35 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.UpdateUserStatusByUserId(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectUserBaseInfoByUserId":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectUserBaseInfoByUserId requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err37 := (strconv.Atoi(flag.Arg(2)))
    if err37 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectUserBaseInfoByUserId(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "resetGoogleCode":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ResetGoogleCode requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err39 := (strconv.Atoi(flag.Arg(2)))
    if err39 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.ResetGoogleCode(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectUserById":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectUserById requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err41 := (strconv.Atoi(flag.Arg(2)))
    if err41 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectUserById(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectUserGoogleStatus":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectUserGoogleStatus requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err43 := (strconv.Atoi(flag.Arg(2)))
    if err43 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectUserGoogleStatus(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectSlaveAllUserCount":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SelectSlaveAllUserCount requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.SelectSlaveAllUserCount(context.Background(), value0))
    fmt.Print("\n")
    break
  case "selectSlaveRegisterCountDay":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectSlaveRegisterCountDay requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err46 := (strconv.Atoi(flag.Arg(2)))
    if err46 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectSlaveRegisterCountDay(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
