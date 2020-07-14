package bdk

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"
	"time"
)

func GetSslInfo(host string) (*x509.Certificate, error) {
	dialer := net.Dialer{Timeout: time.Second * 3}
	conn, err := tls.DialWithDialer(&dialer, "tcp", host+":443",
		&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return nil, err
	}
	defer func() { _ = conn.Close() }()
	if len(conn.ConnectionState().PeerCertificates) == 0 {
		return nil, errors.New("无证书信息")
	}
	cert := conn.ConnectionState().PeerCertificates[0]
	return cert, nil
}
