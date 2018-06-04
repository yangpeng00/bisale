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
	"bisale/bisale-console-api/thrift/inputs"
	"bisale/bisale-console-api/thrift/outputs"
        "bisale/thrift-account/thrift/account"
)

var _ = inputs.GoUnusedProtection__
var _ = outputs.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  bool Ping()")
  fmt.Fprintln(os.Stderr, "  string Version()")
  fmt.Fprintln(os.Stderr, "  LoginOutput MobileLogin(string traceId, MobileLoginInput output)")
  fmt.Fprintln(os.Stderr, "  CreateMemberOutput CreateMember(string traceId, CreateMemberInput input)")
  fmt.Fprintln(os.Stderr, "  string GenerateJWTToken(string traceId, string secretKey, i32 expired)")
  fmt.Fprintln(os.Stderr, "  bool ValidateJWT(string traceId, string tokenString, string secretKey)")
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
  client := account.NewAccountClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "Ping":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "Ping requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.Ping(context.Background()))
    fmt.Print("\n")
    break
  case "Version":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "Version requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.Version(context.Background()))
    fmt.Print("\n")
    break
  case "MobileLogin":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "MobileLogin requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg15 := flag.Arg(2)
    mbTrans16 := thrift.NewTMemoryBufferLen(len(arg15))
    defer mbTrans16.Close()
    _, err17 := mbTrans16.WriteString(arg15)
    if err17 != nil {
      Usage()
      return
    }
    factory18 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt19 := factory18.GetProtocol(mbTrans16)
    argvalue1 := inputs.NewMobileLoginInput()
    err20 := argvalue1.Read(jsProt19)
    if err20 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.MobileLogin(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "CreateMember":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "CreateMember requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg22 := flag.Arg(2)
    mbTrans23 := thrift.NewTMemoryBufferLen(len(arg22))
    defer mbTrans23.Close()
    _, err24 := mbTrans23.WriteString(arg22)
    if err24 != nil {
      Usage()
      return
    }
    factory25 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt26 := factory25.GetProtocol(mbTrans23)
    argvalue1 := inputs.NewCreateMemberInput()
    err27 := argvalue1.Read(jsProt26)
    if err27 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.CreateMember(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "GenerateJWTToken":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "GenerateJWTToken requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err30 := (strconv.Atoi(flag.Arg(3)))
    if err30 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.GenerateJWTToken(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "ValidateJWT":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "ValidateJWT requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.ValidateJWT(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
