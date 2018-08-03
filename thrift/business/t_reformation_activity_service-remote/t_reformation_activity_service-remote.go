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
        "bisale/bisale-console-api/thrift/reformationgh"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  void InsertRelation(string traceId, TParticipantRelation participantRelation)")
  fmt.Fprintln(os.Stderr, "   SelectTop10Account(string traceId)")
  fmt.Fprintln(os.Stderr, "   SelectInviteeListByUserId(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "  void EnableParticipant(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "  TUsername selectRealNameByUsername(string traceId, string username)")
  fmt.Fprintln(os.Stderr, "   selectInviteRecordByUserId(string traceId, i32 userId, i32 page, i32 pageSize)")
  fmt.Fprintln(os.Stderr, "  InviteNum selectInviteNumByUserId(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "   selectRewardRecordByUserId(string traceId, i32 userId, i32 page, i32 pageSize)")
  fmt.Fprintln(os.Stderr, "   selectRewardAmountByUserId(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "  i32 selectTotalRewardRecords(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "  i32 selectTotalInviteRecords(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "   selectInviters(string traceId, i32 invitee_id)")
  fmt.Fprintln(os.Stderr, "  TCandyParameter getCandyParameter(string traceId)")
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
  client := reformationgh.NewTReformationActivityServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "InsertRelation":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InsertRelation requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg35 := flag.Arg(2)
    mbTrans36 := thrift.NewTMemoryBufferLen(len(arg35))
    defer mbTrans36.Close()
    _, err37 := mbTrans36.WriteString(arg35)
    if err37 != nil {
      Usage()
      return
    }
    factory38 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt39 := factory38.GetProtocol(mbTrans36)
    argvalue1 := reformationgh.NewTParticipantRelation()
    err40 := argvalue1.Read(jsProt39)
    if err40 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.InsertRelation(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "SelectTop10Account":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SelectTop10Account requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.SelectTop10Account(context.Background(), value0))
    fmt.Print("\n")
    break
  case "SelectInviteeListByUserId":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectInviteeListByUserId requires 2 args")
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
    fmt.Print(client.SelectInviteeListByUserId(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "EnableParticipant":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "EnableParticipant requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err45 := (strconv.Atoi(flag.Arg(2)))
    if err45 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.EnableParticipant(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectRealNameByUsername":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectRealNameByUsername requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.SelectRealNameByUsername(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectInviteRecordByUserId":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "SelectInviteRecordByUserId requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err49 := (strconv.Atoi(flag.Arg(2)))
    if err49 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err50 := (strconv.Atoi(flag.Arg(3)))
    if err50 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    tmp3, err51 := (strconv.Atoi(flag.Arg(4)))
    if err51 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    fmt.Print(client.SelectInviteRecordByUserId(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "selectInviteNumByUserId":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectInviteNumByUserId requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err53 := (strconv.Atoi(flag.Arg(2)))
    if err53 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectInviteNumByUserId(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectRewardRecordByUserId":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "SelectRewardRecordByUserId requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err55 := (strconv.Atoi(flag.Arg(2)))
    if err55 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err56 := (strconv.Atoi(flag.Arg(3)))
    if err56 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    tmp3, err57 := (strconv.Atoi(flag.Arg(4)))
    if err57 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    fmt.Print(client.SelectRewardRecordByUserId(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "selectRewardAmountByUserId":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectRewardAmountByUserId requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err59 := (strconv.Atoi(flag.Arg(2)))
    if err59 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectRewardAmountByUserId(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectTotalRewardRecords":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectTotalRewardRecords requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err61 := (strconv.Atoi(flag.Arg(2)))
    if err61 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectTotalRewardRecords(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectTotalInviteRecords":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectTotalInviteRecords requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err63 := (strconv.Atoi(flag.Arg(2)))
    if err63 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectTotalInviteRecords(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectInviters":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectInviters requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err65 := (strconv.Atoi(flag.Arg(2)))
    if err65 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectInviters(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "getCandyParameter":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCandyParameter requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetCandyParameter(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
