package webhookparser

import (
	"testing"
	"bytes"
	"bufio"
	"reflect"
)

var webHookLogLines = `
<html><head></head><body style=""><pre style="word-wrap: break-word; white-space: pre-wrap;">level=info response_body="" request_to="https://grotesquemoon.de" response_headers=map[Server:[nginx] X-Request-Id:[1381e8cb388db085cdc3bac457dab9bf] Date:[Tue, 07 Jul 2015 18:29:52 GMT] Content-Type:[text/html; charset=utf-8] X-Powered-By:[Phusion Passenger (mod_rails/mod_rack) 3.0.17] X-Rack-Cache:[invalidate, pass] X-Runtime:[0.034645] Connection:[keep-alive] Set-Cookie:[X-Mapping-fjhppofk=A67D55AC8119CAD031E35EC35B4BCFFD; path=/] Keep-Alive:[timeout=20] Cache-Control:[max-age=0, private, must-revalidate] Status:[200] Etag:[7215ee9c7d9dc229d2921a40e899ec5f] Vary:[Accept-Encoding] X-Ua-Compatible:[IE=Edge,chrome=1]] response_status="201"
level=info response_body="" request_to="https://severeleather.com" response_headers=map[Runscope-Message-Id:[fb814900-c6bc-4002-8007-e7e06b52abb0] Access-Control-Allow-Credentials:[true] Server:[Runscope-Gateway/1.0] Content-Type:[application/json; charset=utf-8] Connection:[keep-alive]] response_status="500"
level=info response_body="" request_to="https://woodenoyster.com.br" response_headers=map[Content-Type:[application/json; charset=utf-8] Connection:[keep-alive] Runscope-Message-Id:[fb814900-c6bc-4002-8007-e7e06b52abb0] Access-Control-Allow-Credentials:[true] Server:[Runscope-Gateway/1.0]] response_status="503"
level=info response_body="" request_to="https://solidwindshield.net.br" response_headers=map[Keep-Alive:[timeout=20] Cache-Control:[max-age=0, private, must-revalidate] Status:[200] Etag:[7215ee9c7d9dc229d2921a40e899ec5f] Vary:[Accept-Encoding] X-Ua-Compatible:[IE=Edge,chrome=1] Server:[nginx] X-Request-Id:[1381e8cb388db085cdc3bac457dab9bf] Date:[Tue, 07 Jul 2015 18:29:52 GMT] Set-Cookie:[X-Mapping-fjhppofk=A67D55AC8119CAD031E35EC35B4BCFFD; path=/] Content-Type:[text/html; charset=utf-8] X-Powered-By:[Phusion Passenger (mod_rails/mod_rack) 3.0.17] X-Rack-Cache:[invalidate, pass] X-Runtime:[0.034645] Connection:[keep-alive]] response_status="404"
`

var anyFileLogLines = `
response_headers=map[Content-Type:[application/json; charset=utf-8] Connection:[keep-alive] Runscope-Message-Id:[fb814900-c6b-4002-8007-e7e06b52abb0]
response_headers=map[Content-Type:[application/json; charset=utf-8] Connection:[keep-alive] Runscope-Message-Id:[fb814900-123-4002-8007-e7e06b52abb0]
response_headers=map[Content-Type:[application/json; charset=utf-8] Connection:[keep-alive] Runscope-Message-Id:[fb814900-555-4002-8007-e7e06b52abb0]
response_headers=map[Content-Type:[application/json; charset=utf-8] Connection:[keep-alive] Runscope-Message-Id:[fb814900-877-4002-8007-e7e06b52abb0]
`

func TestTopRequest(t *testing.T) {

	buf := bytes.NewBufferString(webHookLogLines)
	scanner := bufio.NewScanner(buf)

	urlExpected := map[string]int{"https://grotesquemoon.de": 1, "https://severeleather.com": 1, "https://woodenoyster.com.br": 1, "https://solidwindshield.net.br": 1}

	var url, _, _ = Parse(scanner)

	urlEqual := reflect.DeepEqual(urlExpected, url)

	if !urlEqual {
		t.Error("Url not Equal")
	}
}

func TestTopStatusCode(t *testing.T) {

	buf := bytes.NewBufferString(webHookLogLines)
	scanner := bufio.NewScanner(buf)

	statusExpected := map[string]int{"201": 1, "500": 1, "503": 1, "404": 1}

	var _, status, _ = Parse(scanner)

	statusEqual := reflect.DeepEqual(statusExpected, status)

	if !statusEqual {
		t.Error("Status not Equal")
	}
}

func TestNoMatchFileParsing(t *testing.T) {
	buf := bytes.NewBufferString(anyFileLogLines)
	scanner := bufio.NewScanner(buf)

	var url, status, _ = Parse(scanner)

	if len(url) > 0 {
		t.Error("Url Map Should not be > then 0")
	}

	if len(status) > 0 {
		t.Error("Status Map Should not be > then 0")
	}
}
