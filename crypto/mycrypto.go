/* Copyright 2014 60plusadventures Author. All rights reserved.
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * Portions Copyright 2013 wetalk authors
 * Extracted from https://github.com/beego/wetalk/modules/utils/tools.go
 * and modified by 60plusadventures Author per the same license.
 *
 * Portions Copyright (c) 2009 The Go Authors
 * Extracted from http://code.google.com/p/go/source/browse/pbkdf2/pbkdf2.go
 * per a BSD style license
 *
 * Portions Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>
 * Copyright (C) 2014 by Daniel Kemp <twinj@github.com> Derivative work
 * Extracted from https://github.com/twinj/uuid based on LICENSE file
 * that can be viewed at that URL
 */

// Package sixty/crypto consolidates crypto funcionality used by the application
package crypto

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"github.com/emadera52/sixty/crypto/secretbox"

	"encoding/base64"
	"encoding/hex"

	"errors"
	"fmt"
	"hash"
	"net/http"
	"time"
)

//GenerateRandomString generates a random alphanumeric string
func GetRandomString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// Local func pbkdf2 is used to create an unrecoverable hash for passwords
// see http://code.google.com/p/go/source/browse/pbkdf2/pbkdf2.go?repo=crypto
func pbkdf2(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
	prf := hmac.New(h, password)
	hashLen := prf.Size()
	numBlocks := (keyLen + hashLen - 1) / hashLen

	var buf [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		// N.B.: || means concatenation, ^ means XOR
		// for each block T_i = U_1 ^ U_2 ^ ... ^ U_iter
		// U_1 = PRF(password, salt || uint(i))
		prf.Reset()
		prf.Write(salt)
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)

		// U_n = PRF(password, U_(n-1))
		for n := 2; n <= iter; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
	}
	return dk[:keyLen]
}

// EncryptPassword uses pbkdf2 to encrypt password
func EncryptPassword(rawPwd string, salt string) string {
	pwd := pbkdf2([]byte(rawPwd), []byte(salt), 8442, 50, sha256.New)
	return hex.EncodeToString(pwd)
}

// EncryptEmailAddr takes an email address as a []byte and returns a hex encoded
// encrypted string suitable for storage in the user DB table.
// If encryption fails, returns "" and a new error
func EncryptEmailAddr(em []byte) (hexEm string, err error) {
	emKey := secretbox.NewKey() // emKey is pointer to [Keysize]byte array
	encEm, ok := secretbox.Encrypt(emKey, em)
	if encEm == nil || !ok {
		err := errors.New("Error encrypting email address: ")
		return "", err
	}
	slKey := emKey[:]
	EmEnc := hex.EncodeToString(encEm) //encoded email address as hex string
	EmKey := hex.EncodeToString(slKey) // email key encoded as hex string
	return fmt.Sprintf("%s$%s", EmKey, EmEnc), nil
}

// DecryptEmailAddr takes an encrypted email addr as a string and
// returns a decrypted email address.
// If decryption fails, returns "" and a new error
func DecryptEmailAddr(encEm string) (emAddr string, err error) {
	// create byte slice of entire encoded string
	dbEmail := []byte(encEm[:])
	// for the key, create byte array (for secretbox.Decrypt) and
	// a byte slice (for hex.Decode)
	hexKey := dbEmail[:64] // hex encoded key
	emKey := [32]byte{}    // decoded hexKey will end up here
	slKey := emKey[:]      // byte slice of emKey for hex.Decode
	// for the address create slice of hex value (for hex.Decode) and
	// a byte slice buffer (for secretbox.Decrypt)
	hexEm := dbEmail[65:]                  // hex encoded, encrypted email addr
	var emBuf = make([]byte, len(hexEm)/2) // byte slice buffer for address

	// decode hex values
	hex.Decode(slKey, hexKey) // leaves decoded key value in slKey
	hex.Decode(emBuf, hexEm)  // leaves decoded, still encrypted addr in emBuf
	copy(emKey[:], slKey)     //copy decoded key to key array

	// decrypt email address using key array and emBuf []byte
	// if success return decrypted addr & nil, otherwise return "" and new error
	if slEm, _ := secretbox.Decrypt(&emKey, emBuf); slEm == nil {
		err := errors.New("Error decrypting email address: ")
		return "", err
	} else {
		return string(slEm), nil
	}
}

// CreateEncodedCookie defines a cookie to be created by ResponseWriter
// The resulting cookie persists until manually deleted or until "expire"
// HOURS have passed. "Value" is hex encoded to comply with RFC6265
func CreateEncodedCookie(rw http.ResponseWriter, name string, value string,
	expires int, httpOnly bool) {
	// double encoded as an experiment
	b64Val := base64.URLEncoding.EncodeToString([]byte(value))
	encVal := hex.EncodeToString([]byte(b64Val))
	ec := &http.Cookie{Name: name, Value: encVal,
		Expires:  time.Now().Add(time.Duration(expires) * time.Hour),
		HttpOnly: httpOnly}
	http.SetCookie(rw, ec)
}

// DecodedCookieValue returns the decoded Value of named cookie if it exists
// and can be decoded. Returns nil, nil if named cookie doesn't exist.
// Returns nil, err if there is a decoding error.
func DecodedCookieValue(rw http.ResponseWriter, req *http.Request, name string) ([]byte, error) {
	c, _ := req.Cookie(name)
	// name not found
	if c == nil {
		return nil, nil
	}

	cVal := c.Value
	if len(cVal) > 0 {
		b64Val, _ := hex.DecodeString(cVal)
		dcVal, err := base64.URLEncoding.DecodeString(string(b64Val))
		if err != nil {
			// decoding error
			return nil, err
		} else {
			// success
			return dcVal, nil
		}
	}
	// Treat 0 length Value as not found
	return nil, nil
}

// DeleteCookie marks named cookie for deletion by setting "Expires"
// to Unix epoch plus 1 second. Also sets MaxAge to -1 for modern browsers.
func DeleteCookie(rw http.ResponseWriter, name string) {
	dc := &http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(rw, dc)
}

type UUIDArray [16]byte

// NewV4UUID generates a RFC4122 version 4 UUID
// Extracted from github.com\twinj\uuid
func NewV4UUID() string {
	u := new(UUIDArray)
	// Read random values (or pseudo-randomly) into UUIDArray type.
	_, err := rand.Read(u[:16])
	if err != nil {
		panic(err) // TODO review possible errors decide what to do
	}
	// Set RFC4122 Variant value
	u[8] &= 0x3F
	u[8] |= 0x80
	// Set Version number
	u[6] &= 0x0F
	u[6] |= byte(4) << 4
	// Format as a human readable string
	b := u[:]
	return fmt.Sprintf("{%X-%X-%x-%X%X-%x}",
		b[0:4], b[4:6], b[6:8], b[8:9], b[9:10], b[10:16])
}
