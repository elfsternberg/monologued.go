package rfc1288

import (
	"fmt"
	"path/filepath"
	"runtime"
	"reflect"
	"testing"
)

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestGood_List(t *testing.T) {
	res, req := ParseRfc1288Request("/W")
	assert(t, res == nil, "Expected result to be nil.")
	assert(t, req.Type == UserList, "Expected type to be Userlist")
}

func TestGood_ListWSpaces(t *testing.T) {
	res, req := ParseRfc1288Request("/W                           ");
	assert(t, res == nil, "Expected result to be nil.")
	assert(t, req.Type == UserList, "Expected type to be Userlist")
}

func TestBad_Start(t *testing.T) {
	res, _ := ParseRfc1288Request("")
	assert(t, res != nil, "Expected result to be BadProtocol.")
}

func TestBad_Start1(t *testing.T) {
	res, _ := ParseRfc1288Request("/")
	assert(t, res != nil, "Expected result to be BadProtocol.")
}

func TestBad_Start2(t *testing.T) {
	res, _ := ParseRfc1288Request("/X")
	assert(t, res != nil, "Expected result to be BadProtocol.")
}

func TestGood_Name(t *testing.T) {
	res, req := ParseRfc1288Request("/W foozle")
	assert(t, res == nil, "Expected a good return")
	assert(t, req.Type == User, "Expected User as a return type")
	assert(t, *req.User == "foozle", "The user name did not match passed in value.")
}

func TestGood_NameLf(t *testing.T) {
	res, req := ParseRfc1288Request("/W foozle\n")
	assert(t, res == nil, "Expected a good return")
	assert(t, req.Type == User, "Expected User as a return type")
	assert(t, *req.User == "foozle", "The user name did not match passed in value.")
}

func TestGood_NameCr(t *testing.T) {
	res, req := ParseRfc1288Request("/W foozle\r")
	assert(t, res == nil, "Expected a good return")
	assert(t, req.Type == User, "Expected User as a return type")
	assert(t, *req.User == "foozle", "The user name did not match passed in value.")
}

func TestGood_NameCrLf(t *testing.T) {
	res, req := ParseRfc1288Request("/W foozle\r\n")
	assert(t, res == nil, "Expected a good return")
	assert(t, req.Type == User, "Expected User as a return type")
	assert(t, *req.User == "foozle", "The user name did not match passed in value.")
}

func TestGood_NameExtraSpace(t *testing.T) {
	res, req := ParseRfc1288Request("/W foozle   ")
	assert(t, res == nil, "Expected result to be nil.")
	assert(t, req.Type == User, "Expected type to be User")
	assert(t, *req.User == "foozle", "User name returned did not match")
}

func TestGood_NameWHost(t *testing.T) {
	res, req := ParseRfc1288Request("/W foozle@localhost")
	assert(t, res == nil, "Expected a good return")
	assert(t, req.Type == Remote, "Expected Remote as a return type")
}

func TestGood_NameWHostAndSpaces(t *testing.T) {
	res, req := ParseRfc1288Request("/W foozle@localhost             ")
	assert(t, res == nil, "Expected a good return")
	assert(t, req.Type == Remote, "Expected Remote as a return type")
}

func TestGood_NameWHostAndSpacesAndLowerW(t *testing.T) {
	res, req := ParseRfc1288Request("/w foozle@localhost             ")
	if res != nil {
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
	res, _ := ParseRfc1288Request("/W   foozle..   ")
	if res == nil {
		t.Error("Expected BadRequest")
	}
}

