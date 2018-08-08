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
        "engine"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  i32 selectEngineAccountCountBy(TEngineAccountCheckingParams params)")
  fmt.Fprintln(os.Stderr, "   selectEngineAccountDataBy(TEngineAccountCheckingParams params)")
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
  client := engine.NewTAccountTransfersServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "selectEngineAccountCountBy":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SelectEngineAccountCountBy requires 1 args")
      flag.Usage()
    }
    arg7 := flag.Arg(1)
    mbTrans8 := thrift.NewTMemoryBufferLen(len(arg7))
    defer mbTrans8.Close()
    _, err9 := mbTrans8.WriteString(arg7)
    if err9 != nil {
      Usage()
      return
    }
    factory10 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt11 := factory10.GetProtocol(mbTrans8)
    argvalue0 := engine.NewTEngineAccountCheckingParams()
    err12 := argvalue0.Read(jsProt11)
    if err12 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.SelectEngineAccountCountBy(context.Background(), value0))
    fmt.Print("\n")
    break
  case "selectEngineAccountDataBy":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SelectEngineAccountDataBy requires 1 args")
      flag.Usage()
    }
    arg13 := flag.Arg(1)
    mbTrans14 := thrift.NewTMemoryBufferLen(len(arg13))
    defer mbTrans14.Close()
    _, err15 := mbTrans14.WriteString(arg13)
    if err15 != nil {
      Usage()
      return
    }
    factory16 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt17 := factory16.GetProtocol(mbTrans14)
    argvalue0 := engine.NewTEngineAccountCheckingParams()
    err18 := argvalue0.Read(jsProt17)
    if err18 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.SelectEngineAccountDataBy(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
