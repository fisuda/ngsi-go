/*
MIT License

Copyright (c) 2020-2021 Kazuhito Suda

This file is part of NGSI Go

https://github.com/lets-fiware/ngsi-go

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package ngsicmd

import (
	"net"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

// NetLib is ...
type NetLib interface {
	InterfaceAddrs() ([]net.Addr, error)
	ListenAndServe(addr string, handler http.Handler) error
	ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error
}

type netLib struct {
}

func (n *netLib) InterfaceAddrs() ([]net.Addr, error) {
	return net.InterfaceAddrs()
}

func (n *netLib) ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}
func (n *netLib) ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, handler)
}

func getDateTime(dateTime string) (string, error) {
	const funcName = "getDateTime"

	var err error
	if !ngsilib.IsOrionDateTime(dateTime) {
		dateTime, err = ngsilib.GetExpirationDate(dateTime)
		if err != nil {
			return "", &ngsiCmdError{funcName, 1, err.Error(), nil}
		}
	}
	return dateTime, nil
}
