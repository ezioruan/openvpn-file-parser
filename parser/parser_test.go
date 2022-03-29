package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	FilePath string
	CA       string
	Cert     string
	Key      string
}

func TestParseFromFile(t *testing.T) {

	testCases := []TestCase{
		{"../test-data/test.ovpn",
			`-----BEGIN CERTIFICATE-----
MIID+jCCAuKgAwIBAgIBADANBgkqhkiG9w0BAQsFADB8MSMwIQYDVQQDDBp2cG4y
OTc1OTY1Nzkuc29mdGV0aGVyLm5ldDEjMCEGA1UECgwadnBuMjk3NTk2NTc5LnNv
ZnRldGhlci5uZXQxIzAhBgNVBAsMGnZwbjI5NzU5NjU3OS5zb2Z0ZXRoZXIubmV0
MQswCQYDVQQGEwJVUzAeFw0yMTAyMjQwOTAzMTVaFw0zNzEyMzEwOTAzMTVaMHwx
IzAhBgNVBAMMGnZwbjI5NzU5NjU3OS5zb2Z0ZXRoZXIubmV0MSMwIQYDVQQKDBp2
cG4yOTc1OTY1Nzkuc29mdGV0aGVyLm5ldDEjMCEGA1UECwwadnBuMjk3NTk2NTc5
LnNvZnRldGhlci5uZXQxCzAJBgNVBAYTAlVTMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAsN0+zDhpRK5Q3A0P4kvgGpzFJ2JRyRTcD1MiNIl5oBpKigkQ
KnjFJV3XbxOcNjzlvG7GStW3ht3b3pvzvjzjbhzhOJMftn7MBnZoOQxvy0U5NVK7
U5aHY/gCW92ShGV9QPCuthwck2KDsWeE7UEt9fAClc35ubzwGDZTaRU0cRa4r6Av
xxPwvcZR0M0/xQg+dpIjrNksylLdf4wxOL6AhTFa5u6eTFsTdxNSXvS5RfYjXQdd
nBwOLmvU65kPFNeqVMkRTwpHWtIzWNvd4g62f3Mg2p43dmJnyK1tO8vLmo7t7Rdt
DWQ+dw6cJtZRvil3nsvrA3qTZilJmGZigstHAQIDAQABo4GGMIGDMA8GA1UdEwEB
/wQFMAMBAf8wCwYDVR0PBAQDAgH2MGMGA1UdJQRcMFoGCCsGAQUFBwMBBggrBgEF
BQcDAgYIKwYBBQUHAwMGCCsGAQUFBwMEBggrBgEFBQcDBQYIKwYBBQUHAwYGCCsG
AQUFBwMHBggrBgEFBQcDCAYIKwYBBQUHAwkwDQYJKoZIhvcNAQELBQADggEBAJQm
EIvwAeqExlJvVCtXPxM0Ji2+OGdlncgY2v5EUqVSSx/w0vuCuHce/+XyJDaSdQMI
Ew4HY+MAau8T8Q9JLlhkJ9qFPOqGo/8A5RxhXZbtfllxVJTENDYMM+0ptJ3N1WcW
59xx6dHktrPOn7gNO4YpNEm+vofxlir1q2QcxTm49DEBTC+3LGr51kld79z7O2Qx
fGsnXVBlDxHi9Ucr+ek98fBhqTzJct29dnrhQLCIqeJcUos2GsYm58TLBsLLl8Ti
F+O5FGg1ck4h6brB1ukp6l0vDEyd9E7rbgZbIwOfN+8jSHjV8YC3A7IyD9RcHde5
r5a5Hb3cohvzu/WI9No=
-----END CERTIFICATE-----`, `-----BEGIN CERTIFICATE-----
test cert
-----END CERTIFICATE-----`, `-----BEGIN RSA PRIVATE KEY-----
test key
-----END RSA PRIVATE KEY-----`},
	}

	for _, testCase := range testCases {
		config, err := ParseFromFile(testCase.FilePath)
		if err != nil {
			t.Errorf("ParseFromFile error %v", err)
		}
		assert.NotEmpty(t, config.CA)
		assert.Equal(t, config.CA, testCase.CA)
		assert.Equal(t, config.Cert, testCase.Cert)
		assert.NotEmpty(t, config.Cert)
		assert.Equal(t, config.Key, testCase.Key)
		assert.NotEmpty(t, config.Key)

	}

}
