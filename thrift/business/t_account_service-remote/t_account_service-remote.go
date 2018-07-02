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
        "bisale/bisale-console-api/thrift/business"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TUser SelectUserByUsername(string traceId, string username)")
  fmt.Fprintln(os.Stderr, "  TUser SelectUserByUsernameAndPassword(string traceId, string username, string password)")
  fmt.Fprintln(os.Stderr, "  string GetUsernameByToken(string traceId, string token)")
  fmt.Fprintln(os.Stderr, "  bool IsTokenNotExpired(string traceId, string token)")
  fmt.Fprintln(os.Stderr, "  string GetTokenByUsername(string traceId, string username)")
  fmt.Fprintln(os.Stderr, "  TAccount SelectAccountByUserId(string traceId, i32 id)")
  fmt.Fprintln(os.Stderr, "  i32 InsertUser(string traceId, TUser user)")
  fmt.Fprintln(os.Stderr, "  void UpdateUser(string traceId, TUser user)")
  fmt.Fprintln(os.Stderr, "  void InsertGoogleAuth(string traceId, TGoogleAuth googleAuth)")
  fmt.Fprintln(os.Stderr, "  void UpdateGoogleAuth(string traceId, TGoogleAuth googleAuth)")
  fmt.Fprintln(os.Stderr, "  void InsertTradePasswordAuth(string traceId, TTradePasswordAuth tradePasswordAuth)")
  fmt.Fprintln(os.Stderr, "  void UpdateTradePasswordAuth(string traceId, TTradePasswordAuth tradePasswordAuth)")
  fmt.Fprintln(os.Stderr, "  void InsertEngine(string traceId, TEngine engine)")
  fmt.Fprintln(os.Stderr, "  void UpdateEngine(string traceId, TEngine engine)")
  fmt.Fprintln(os.Stderr, "  TKyc SelectKycByUserId(string traceId, i32 userId)")
  fmt.Fprintln(os.Stderr, "  void InsertKyc(string traceId, TKyc kyc)")
  fmt.Fprintln(os.Stderr, "  void UpdateKyc(string traceId, TKyc kyc)")
  fmt.Fprintln(os.Stderr, "  string selectUserKycStatusByUserName(string userName)")
  fmt.Fprintln(os.Stderr, "  void addCaptchaErrorCount(string user_id)")
  fmt.Fprintln(os.Stderr, "  TUser getUserInfo(string user_id)")
  fmt.Fprintln(os.Stderr, "  void lockUser(string user_id)")
  fmt.Fprintln(os.Stderr, "   getCurrencyInfo()")
  fmt.Fprintln(os.Stderr, "  void unlockUser(i32 user_id)")
  fmt.Fprintln(os.Stderr, "  TEngine getEngineInfoByUserId(string user_id)")
  fmt.Fprintln(os.Stderr, "  bool createEngineAccountByUserId(string traceId, i32 user_id, string accountInfoUrl)")
  fmt.Fprintln(os.Stderr, "  i32 updatePasswordByEmail(string email, string loginPwd)")
  fmt.Fprintln(os.Stderr, "  TEngine createEngineAccount(string user_id, string accountInfoUrl)")
  fmt.Fprintln(os.Stderr, "  TUser selectCodeUserByUsername(string traceId, string username)")
  fmt.Fprintln(os.Stderr, "  void updateLoginPassword(string traceId, string user_id, string password)")
  fmt.Fprintln(os.Stderr, "  TUserEntity checkPasswordByUserId(string traceId, string user_id, string password)")
  fmt.Fprintln(os.Stderr, "  TUserConfigStatus selectUserConfigStatusByUserId(i32 userId)")
  fmt.Fprintln(os.Stderr, "  TAccountBaseInfo selectAccountBaseInfoByUserId(i32 userId)")
  fmt.Fprintln(os.Stderr, "  void updateMobile(string traceId, i32 user_id, string mobile)")
  fmt.Fprintln(os.Stderr, "  void removeMobile(string traceId, i32 user_id)")
  fmt.Fprintln(os.Stderr, "  TUser selectUserById(string traceId, i32 user_id)")
  fmt.Fprintln(os.Stderr, "  TUser SelectTUserByUsername(string traceId, string username)")
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
  client := business.NewTAccountServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "SelectUserByUsername":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectUserByUsername requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.SelectUserByUsername(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "SelectUserByUsernameAndPassword":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "SelectUserByUsernameAndPassword requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.SelectUserByUsernameAndPassword(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "GetUsernameByToken":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetUsernameByToken requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.GetUsernameByToken(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "IsTokenNotExpired":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "IsTokenNotExpired requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.IsTokenNotExpired(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "GetTokenByUsername":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetTokenByUsername requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.GetTokenByUsername(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "SelectAccountByUserId":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectAccountByUserId requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err87 := (strconv.Atoi(flag.Arg(2)))
    if err87 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectAccountByUserId(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "InsertUser":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InsertUser requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg89 := flag.Arg(2)
    mbTrans90 := thrift.NewTMemoryBufferLen(len(arg89))
    defer mbTrans90.Close()
    _, err91 := mbTrans90.WriteString(arg89)
    if err91 != nil {
      Usage()
      return
    }
    factory92 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt93 := factory92.GetProtocol(mbTrans90)
    argvalue1 := business.NewTUser()
    err94 := argvalue1.Read(jsProt93)
    if err94 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.InsertUser(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "UpdateUser":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "UpdateUser requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg96 := flag.Arg(2)
    mbTrans97 := thrift.NewTMemoryBufferLen(len(arg96))
    defer mbTrans97.Close()
    _, err98 := mbTrans97.WriteString(arg96)
    if err98 != nil {
      Usage()
      return
    }
    factory99 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt100 := factory99.GetProtocol(mbTrans97)
    argvalue1 := business.NewTUser()
    err101 := argvalue1.Read(jsProt100)
    if err101 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.UpdateUser(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "InsertGoogleAuth":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InsertGoogleAuth requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg103 := flag.Arg(2)
    mbTrans104 := thrift.NewTMemoryBufferLen(len(arg103))
    defer mbTrans104.Close()
    _, err105 := mbTrans104.WriteString(arg103)
    if err105 != nil {
      Usage()
      return
    }
    factory106 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt107 := factory106.GetProtocol(mbTrans104)
    argvalue1 := business.NewTGoogleAuth()
    err108 := argvalue1.Read(jsProt107)
    if err108 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.InsertGoogleAuth(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "UpdateGoogleAuth":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "UpdateGoogleAuth requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg110 := flag.Arg(2)
    mbTrans111 := thrift.NewTMemoryBufferLen(len(arg110))
    defer mbTrans111.Close()
    _, err112 := mbTrans111.WriteString(arg110)
    if err112 != nil {
      Usage()
      return
    }
    factory113 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt114 := factory113.GetProtocol(mbTrans111)
    argvalue1 := business.NewTGoogleAuth()
    err115 := argvalue1.Read(jsProt114)
    if err115 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.UpdateGoogleAuth(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "InsertTradePasswordAuth":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InsertTradePasswordAuth requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg117 := flag.Arg(2)
    mbTrans118 := thrift.NewTMemoryBufferLen(len(arg117))
    defer mbTrans118.Close()
    _, err119 := mbTrans118.WriteString(arg117)
    if err119 != nil {
      Usage()
      return
    }
    factory120 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt121 := factory120.GetProtocol(mbTrans118)
    argvalue1 := business.NewTTradePasswordAuth()
    err122 := argvalue1.Read(jsProt121)
    if err122 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.InsertTradePasswordAuth(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "UpdateTradePasswordAuth":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "UpdateTradePasswordAuth requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg124 := flag.Arg(2)
    mbTrans125 := thrift.NewTMemoryBufferLen(len(arg124))
    defer mbTrans125.Close()
    _, err126 := mbTrans125.WriteString(arg124)
    if err126 != nil {
      Usage()
      return
    }
    factory127 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt128 := factory127.GetProtocol(mbTrans125)
    argvalue1 := business.NewTTradePasswordAuth()
    err129 := argvalue1.Read(jsProt128)
    if err129 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.UpdateTradePasswordAuth(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "InsertEngine":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InsertEngine requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg131 := flag.Arg(2)
    mbTrans132 := thrift.NewTMemoryBufferLen(len(arg131))
    defer mbTrans132.Close()
    _, err133 := mbTrans132.WriteString(arg131)
    if err133 != nil {
      Usage()
      return
    }
    factory134 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt135 := factory134.GetProtocol(mbTrans132)
    argvalue1 := business.NewTEngine()
    err136 := argvalue1.Read(jsProt135)
    if err136 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.InsertEngine(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "UpdateEngine":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "UpdateEngine requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg138 := flag.Arg(2)
    mbTrans139 := thrift.NewTMemoryBufferLen(len(arg138))
    defer mbTrans139.Close()
    _, err140 := mbTrans139.WriteString(arg138)
    if err140 != nil {
      Usage()
      return
    }
    factory141 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt142 := factory141.GetProtocol(mbTrans139)
    argvalue1 := business.NewTEngine()
    err143 := argvalue1.Read(jsProt142)
    if err143 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.UpdateEngine(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "SelectKycByUserId":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectKycByUserId requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err145 := (strconv.Atoi(flag.Arg(2)))
    if err145 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectKycByUserId(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "InsertKyc":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InsertKyc requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg147 := flag.Arg(2)
    mbTrans148 := thrift.NewTMemoryBufferLen(len(arg147))
    defer mbTrans148.Close()
    _, err149 := mbTrans148.WriteString(arg147)
    if err149 != nil {
      Usage()
      return
    }
    factory150 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt151 := factory150.GetProtocol(mbTrans148)
    argvalue1 := business.NewTKyc()
    err152 := argvalue1.Read(jsProt151)
    if err152 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.InsertKyc(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "UpdateKyc":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "UpdateKyc requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg154 := flag.Arg(2)
    mbTrans155 := thrift.NewTMemoryBufferLen(len(arg154))
    defer mbTrans155.Close()
    _, err156 := mbTrans155.WriteString(arg154)
    if err156 != nil {
      Usage()
      return
    }
    factory157 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt158 := factory157.GetProtocol(mbTrans155)
    argvalue1 := business.NewTKyc()
    err159 := argvalue1.Read(jsProt158)
    if err159 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.UpdateKyc(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectUserKycStatusByUserName":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SelectUserKycStatusByUserName requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.SelectUserKycStatusByUserName(context.Background(), value0))
    fmt.Print("\n")
    break
  case "addCaptchaErrorCount":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "AddCaptchaErrorCount requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.AddCaptchaErrorCount(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getUserInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetUserInfo requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetUserInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "lockUser":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "LockUser requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.LockUser(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getCurrencyInfo":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetCurrencyInfo requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetCurrencyInfo(context.Background()))
    fmt.Print("\n")
    break
  case "unlockUser":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "UnlockUser requires 1 args")
      flag.Usage()
    }
    tmp0, err164 := (strconv.Atoi(flag.Arg(1)))
    if err164 != nil {
      Usage()
      return
    }
    argvalue0 := int32(tmp0)
    value0 := argvalue0
    fmt.Print(client.UnlockUser(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getEngineInfoByUserId":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetEngineInfoByUserId requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetEngineInfoByUserId(context.Background(), value0))
    fmt.Print("\n")
    break
  case "createEngineAccountByUserId":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "CreateEngineAccountByUserId requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err167 := (strconv.Atoi(flag.Arg(2)))
    if err167 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.CreateEngineAccountByUserId(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "updatePasswordByEmail":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "UpdatePasswordByEmail requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.UpdatePasswordByEmail(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "createEngineAccount":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "CreateEngineAccount requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.CreateEngineAccount(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectCodeUserByUsername":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectCodeUserByUsername requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.SelectCodeUserByUsername(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "updateLoginPassword":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "UpdateLoginPassword requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.UpdateLoginPassword(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "checkPasswordByUserId":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "CheckPasswordByUserId requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.CheckPasswordByUserId(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "selectUserConfigStatusByUserId":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SelectUserConfigStatusByUserId requires 1 args")
      flag.Usage()
    }
    tmp0, err181 := (strconv.Atoi(flag.Arg(1)))
    if err181 != nil {
      Usage()
      return
    }
    argvalue0 := int32(tmp0)
    value0 := argvalue0
    fmt.Print(client.SelectUserConfigStatusByUserId(context.Background(), value0))
    fmt.Print("\n")
    break
  case "selectAccountBaseInfoByUserId":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SelectAccountBaseInfoByUserId requires 1 args")
      flag.Usage()
    }
    tmp0, err182 := (strconv.Atoi(flag.Arg(1)))
    if err182 != nil {
      Usage()
      return
    }
    argvalue0 := int32(tmp0)
    value0 := argvalue0
    fmt.Print(client.SelectAccountBaseInfoByUserId(context.Background(), value0))
    fmt.Print("\n")
    break
  case "updateMobile":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "UpdateMobile requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err184 := (strconv.Atoi(flag.Arg(2)))
    if err184 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.UpdateMobile(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "removeMobile":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "RemoveMobile requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err187 := (strconv.Atoi(flag.Arg(2)))
    if err187 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.RemoveMobile(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "selectUserById":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectUserById requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err189 := (strconv.Atoi(flag.Arg(2)))
    if err189 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.SelectUserById(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "SelectTUserByUsername":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SelectTUserByUsername requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.SelectTUserByUsername(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}