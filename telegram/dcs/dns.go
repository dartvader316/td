package dcs

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
	"sort"
	"sync"

	"golang.org/x/net/dns/dnsmessage"
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/tg"
)

var dnsKey struct {
	rsa.PublicKey
	once sync.Once
	eBig *big.Int
}

func init() {
	dnsKey.once.Do(func() {
		k, err := crypto.ParseRSAPublicKeys([]byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAyr+18Rex2ohtVy8sroGPBwXD3DOoKCSpjDqYoXgCqB7ioln4eDCF
fOBUlfXUEvM/fnKCpF46VkAftlb4VuPDeQSS/ZxZYEGqHaywlroVnXHIjgqoxiAd
192xRGreuXIaUKmkwlM9JID9WS2jUsTpzQ91L8MEPLJ/4zrBwZua8W5fECwCCh2c
9G5IzzBm+otMS/YKwmR1olzRCyEkyAEjXWqBI9Ftv5eG8m0VkBzOG655WIYdyV0H
fDK/NWcvGqa0w/nriMD6mDjKOryamw0OP9QuYgMN0C9xMW9y8SmP4h92OAWodTYg
Y1hZCxdv6cs5UnW9+PWvS+WIbkh+GaWYxwIDAQAB
-----END RSA PUBLIC KEY-----`))
		if err != nil {
			panic(err)
		}

		dnsKey.PublicKey = *k[0]
		dnsKey.eBig = big.NewInt(int64(dnsKey.E))
	})
}

// parseDNSList decodes raw encrypted simple config.
//
// Notice that parseDNSList does not decode base64, user should do it manually.
func parseDNSList(input [256]byte) (tg.HelpConfigSimple, error) {
	// See https://github.com/tdlib/td/blob/master/td/telegram/ConfigManager.cpp#L148.
	x := new(big.Int).SetBytes(input[:])
	y := new(big.Int).Exp(x, dnsKey.eBig, dnsKey.N)

	dataRSA := make([]byte, 256)
	if !crypto.FillBytes(y, dataRSA) {
		return tg.HelpConfigSimple{}, xerrors.New("dataRSA has invalid size")
	}

	block, err := aes.NewCipher(dataRSA[:32])
	if err != nil {
		return tg.HelpConfigSimple{}, err
	}
	d := cipher.NewCBCDecrypter(block, dataRSA[16:32])
	dataCBC := dataRSA[32:]
	d.CryptBlocks(dataCBC, dataCBC)

	decrypted := dataCBC[:len(dataCBC)-16]
	decryptedHash := sha256.Sum256(decrypted[:])
	hash := dataCBC[len(dataCBC)-16:]

	if !bytes.Equal(decryptedHash[:16], hash) {
		return tg.HelpConfigSimple{}, xerrors.New("hash mismatch")
	}

	var cfg tg.HelpConfigSimple
	if err := cfg.Decode(&bin.Buffer{Buf: decrypted[4:]}); err != nil {
		return tg.HelpConfigSimple{}, err
	}
	return cfg, nil
}

type sortByLen []string

func (s sortByLen) Len() int {
	return len(s)
}

func (s sortByLen) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func (s sortByLen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// DNSConfig parses tg.HelpConfigSimple from TXT response.
func DNSConfig(r dnsmessage.TXTResource) (tg.HelpConfigSimple, error) {
	encoding := base64.StdEncoding
	const (
		decodedLen = 256
		encodedLen = 344
	)
	sort.Sort(sortByLen(r.TXT))

	var totalLength int
	for i := range r.TXT {
		totalLength += len(r.TXT[i])
	}
	if totalLength != encodedLen {
		return tg.HelpConfigSimple{}, xerrors.Errorf("invalid input length %d", totalLength)
	}

	var (
		encoded [encodedLen]byte
		decoded [decodedLen]byte
	)
	n := 0
	for i := range r.TXT {
		n += copy(encoded[n:], r.TXT[i])
	}

	if _, err := encoding.Decode(decoded[:], encoded[:]); err != nil {
		return tg.HelpConfigSimple{}, xerrors.Errorf("decode: %w", err)
	}

	cfg, err := parseDNSList(decoded)
	if err != nil {
		return tg.HelpConfigSimple{}, xerrors.Errorf("decrypt config: %w", err)
	}

	return cfg, nil
}
