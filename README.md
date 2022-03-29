# openvpn-file-parser

A tools to extract `CA`, `Cert` and `Key` from `.ovpn` files

[![Go Reportcard](https://goreportcard.com/badge/github.com/ezioruan/openvpn-file-parser)](https://goreportcard.com/report/github.com/ezioruan/openvpn-file-parser)


## Usage 

with Go module enabled (go 1.11+)
`
go get github.com/ezioruan/openvpn-file-parser
`


### Parse file 

```

config, err := parser.ParseFromFile("client.ovpn")
if err != nil {
    fmt.Printf("Error parse openvpn file %v", err)
    return
}
fmt.Sprintf("CA %s \n Cert %s \n Key %s \n", config.CA, config.Cert, config.Key)

```



