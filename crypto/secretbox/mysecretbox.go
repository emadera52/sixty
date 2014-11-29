/* Copyright (c) 2013 Kyle Isom <kyle@gokyle.org>
 * Extracted from https://github.com/kisom/gocrypto/secretbox/secretbox.go
 * secretbox is a wrapper around the NaCl secretbox package. It generates a
 * random nonce for each message and prepends this to the ciphertext.
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.

 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
package secretbox

import (
	"code.google.com/p/go.crypto/nacl/secretbox"
	"crypto/rand"
	"io"
)

func randBytes(size int) []byte {
	p := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, p)
	if err != nil {
		p = nil
	}
	return p
}

const (
	KeySize   = 32
	nonceSize = 24
)

func newNonce() *[nonceSize]byte {
	var nonce [nonceSize]byte
	p := randBytes(nonceSize)
	if p == nil {
		return nil
	}
	copy(nonce[:], p)
	return &nonce
}
func zero(in []byte) {
	for i := 0; i < len(in); i++ {
		in[i] = 0
	}
}

// NewKey generates a new NaCl key.
func NewKey() *[KeySize]byte {
	p := randBytes(KeySize)
	if p == nil {
		return nil
	}
	defer zero(p)
	var key [KeySize]byte
	copy(key[:], p)
	return &key
}

func Encrypt(key *[KeySize]byte, in []byte) ([]byte, bool) {
	var out = make([]byte, nonceSize)
	nonce := newNonce()
	if nonce == nil {
		return nil, false
	}

	copy(out, nonce[:])
	out = secretbox.Seal(out, in, nonce, key)
	return out, true
}

func Decrypt(key *[KeySize]byte, in []byte) ([]byte, bool) {
	if len(in) < nonceSize {
		return nil, false
	}
	var nonce [nonceSize]byte
	copy(nonce[:], in)
	return secretbox.Open(nil, in[nonceSize:], &nonce, key)
}
