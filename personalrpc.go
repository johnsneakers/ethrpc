package ethrpc

func (rpc *EthRPC) Personal_newAccount(password string) (string, error) {
	var hash string
	pwd := []string{password}
	err := rpc.call("personal_newAccount", &hash, pwd)
	return hash, err
}