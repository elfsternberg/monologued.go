package rfc1288

import "testing"

func TestGood_List(t *testing.T) {
	res, req := parse_rfc1288_request("/W")
	if res != Ok {
		t.Error("Expected a good return")
	}
	if req.Type != UserList {
		t.Error("Expected UserList as a return type")
	}
}

func TestGood_ListWSpaces(t *testing.T) {
	res, req := parse_rfc1288_request("/W                           ");
	if res != Ok {
		t.Error("Expected a good return")
	}
	if req.Type != UserList {
		t.Error("Expected UserList as a return type")
	}
}

func TestBad_Start(t *testing.T) {
	res, _ := parse_rfc1288_request("")
	if res != BadProtocol {
		t.Error("Expected BadProtocol")
	}
}

func TestBad_Start1(t *testing.T) {
	res, _ := parse_rfc1288_request("/")
	if res != BadProtocol {
		t.Error("Expected BadProtocol")
	}
}

func TestBad_Start2(t *testing.T) {
	res, _ := parse_rfc1288_request("/X")
	if res != BadProtocol {
		t.Error("Expected BadProtocol, got", res)
	}
}

func TestGood_Name(t *testing.T) {
	res, req := parse_rfc1288_request("/W foozle")
	if res != Ok {
		t.Error("Expected a good return")
	}
	if req.Type != User {
		t.Error("Expected User as a return type")
	}
	if *req.User != "foozle" {
		t.Error("The user name did not match passed in value.")
	}
}

func TestGood_NameExtraSpace(t *testing.T) {
	res, req := parse_rfc1288_request("/W foozle   ")
	if res != Ok {
		t.Error("Expected a good return")
	}
	if req.Type != User {
		t.Error("Expected User as a return type")
	}
	if *req.User != "foozle" {
		t.Error("The user name did not match passed in value.")
	}
}

func TestGood_NameWHost(t *testing.T) {
	res, req := parse_rfc1288_request("/W foozle@localhost")
	if res != Ok {
		t.Error("Expected a good return")
	}
	if req.Type != Remote {
		t.Error("Expected Remote as a return type")
	}
	if *req.User != "foozle" {
		t.Error("The user name did not match passed in value.")
	}
	if *req.Host != "localhost" {
		t.Error("The host name did not match passed in value.")
	}
}

func TestGood_NameWHostAndSpaces(t *testing.T) {
	res, req := parse_rfc1288_request("/W foozle@localhost             ")
	if res != Ok {
		t.Error("Expected a good return")
	}
	if req.Type != Remote {
		t.Error("Expected Remote as a return type")
	}
	if *req.User != "foozle" {
		t.Error("The user name did not match passed in value.")
	}
	if *req.Host != "localhost" {
		t.Error("The host name did not match passed in value.")
	}
}

func TestGood_NameWHostAndSpacesAndLowerW(t *testing.T) {
	res, req := parse_rfc1288_request("/w foozle@localhost             ")
	if res != Ok {
		t.Error("Expected a good return")
	}
	if req.Type != Remote {
		t.Error("Expected Remote as a return type")
	}
	if *req.User != "foozle" {
		t.Error("The user name did not match passed in value.")
	}
	if *req.Host != "localhost" {
		t.Error("The host name did not match passed in value.")
	}
}

func TestBad_Name(t *testing.T) {
	res, _ := parse_rfc1288_request("/W   foozle..   ")
	if res != BadRequest {
		t.Error("Expected BadRequest")
	}
}

