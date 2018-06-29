namespace go wallet

service Wallet {
  string execute(1: string moduleName, 2: string methodName, 3: string jsonStr);
}
