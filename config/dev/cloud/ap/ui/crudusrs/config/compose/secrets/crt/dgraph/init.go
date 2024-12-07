package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"

	"os"

	"time"
)

func main() {
	_, _, err := certsetup()
	if err != nil {
		panic(err)
	}
}
func certsetup() (*tls.Config, *tls.Config, error) {
	ca := genCA()
	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}
	saveKeys("key", caPrivKey)
	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return nil, nil, err
	}
	caPEM, _ := os.Create("ca.pem")
	pem.Encode(caPEM, &pem.Block{Type: "CERTIFICATE", Bytes: caBytes})
	defer caPEM.Close()
	caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey)})
	cert := genServerCertDB()
	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}
	saveKeys("server", certPrivKey)
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, &certPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return nil, nil, err
	}
	saveCert(certBytes)
	certPEM, _ := os.Create("cert.pem")
	pem.Encode(certPEM, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	defer certPEM.Close()
	certPrivKeyPEM, _ := os.Create("id_cert.pem")
	pem.Encode(certPrivKeyPEM, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(certPrivKey)})
	defer certPrivKeyPEM.Close()
	certPemRead, _ := os.ReadFile("cert.pem")
	certPrivKeyPEMRead, _ := os.ReadFile("id_cert.pem")
	serverCert, err := tls.X509KeyPair(certPemRead, certPrivKeyPEMRead)
	if err != nil {
		return nil, nil, err
	}
	serverTLSConf := &tls.Config{Certificates: []tls.Certificate{serverCert}}
	caPEMRead, _ := os.ReadFile("ca.pem")
	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(caPEMRead)
	clientTLSConf := &tls.Config{RootCAs: certpool}
	return serverTLSConf, clientTLSConf, err
}
func genCA() *x509.Certificate {
	return &x509.Certificate{SerialNumber: big.NewInt(2021), Subject: pkix.Name{Country: []string{"US"}, Organization: []string{"AppPathway"}, OrganizationalUnit: []string{"Executive"}, Locality: []string{"North Attleboro"}, Province: []string{"Massachusetts"}}, NotBefore: time.Now(), NotAfter: time.Now().AddDate(10, 0, 0), IsCA: true, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}, KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, BasicConstraintsValid: true}
}

func genServerCertDB() *x509.Certificate {
	return &x509.Certificate{SerialNumber: big.NewInt(2021), Subject: pkix.Name{
		Country:            []string{"US"},
		Organization:       []string{"AppPathway"},
		OrganizationalUnit: []string{"Executive"},
		Locality:           []string{"North Attleboro"},
		Province:           []string{"Massachusetts"}},
		DNSNames: []string{
			"alpha1",
			"apppathwayserver",
			"is",
			"intentsysd",
			"intentd",
			"schedulerd",
			"smsd",
		},
		IPAddresses: []net.IP{net.IPv4(10, 0, 0, 42)},
		NotBefore:   time.Now(), NotAfter: time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
}
func saveKeys(name string, privatekey *rsa.PrivateKey) {
	pkey := x509.MarshalPKCS1PrivateKey(privatekey)
	ioutil.WriteFile(name+"_private.key", pkey, 0777)
	fmt.Println("private key saved to private.key")
	publickey := &privatekey.PublicKey
	pubkey, _ := x509.MarshalPKIXPublicKey(publickey)
	ioutil.WriteFile(name+"_public.key", pubkey, 0777)
	fmt.Println("public key saved to public.key")
}
func saveCert(cert []byte) {
	ioutil.WriteFile("cert.pem", cert, 0777)
	fmt.Println("certificate saved to cert.pem")
}
