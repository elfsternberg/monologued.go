package rfc1288

import(
	"bytes"
//	"fmt"
)

func is_unix_conventional(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

var Rfc1288ErrorMsgs = []string {
	"",
	"Protocol prefix not recognized",
	"Protocol request does not meet specifications",
}

type Rfc1288ErrorCode int

const (
	Ok          Rfc1288ErrorCode = 0
	BadProtocol Rfc1288ErrorCode = 1
	BadRequest  Rfc1288ErrorCode = 2
)

type Rfc1288RequestType int

const (
	UserList Rfc1288RequestType = 0
	User     Rfc1288RequestType = 1
	Remote   Rfc1288RequestType = 2
)

type Rfc1288Request struct {
	Type Rfc1288RequestType
	User* string
	Host* string
}

func parse_rfc1288_request(Buffer string) (Rfc1288ErrorCode, *Rfc1288Request) {
	Buflen := len(Buffer)

	if Buflen < 2 {
		return BadProtocol, nil
	}

	if Buffer[0] != '/' || (Buffer[1] != 'W' && Buffer[1] != 'w') {
		return BadProtocol, nil
	}

	if len(Buffer) == 2 {
		return Ok, &Rfc1288Request{UserList, nil, nil}
	}

	index := 2
	for index < Buflen && Buffer[index] == ' ' {
		index += 1
	}

	if Buflen == index {
		return Ok, &Rfc1288Request{Type: UserList, User: nil, Host: nil}
	}

	user := bytes.NewBufferString("")
	host := bytes.NewBufferString("")

	for index < Buflen && is_unix_conventional(Buffer[index]) {
		user.WriteByte(Buffer[index])
		index += 1
	}

	if index == Buflen || (index < Buflen && Buffer[index] == ' ') {
		ruser := user.String()
		return Ok, &Rfc1288Request{Type: User, User: &ruser, Host: nil}
	}

	if Buffer[index] != '@' {
		return BadRequest, nil
	}

	index += 1
	for index < Buflen && Buffer[index] != ' ' {
		host.WriteByte(Buffer[index])
		index += 1
	}

	ruser := user.String()
	rhost := host.String()
	return Ok, &Rfc1288Request{Type: Remote, User: &ruser, Host: &rhost}
}
