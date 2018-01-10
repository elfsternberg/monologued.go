package rfc1288

import(
	"errors"
	"strings"
	"bytes"
)

func is_unix_conventional(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

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

func ParseRfc1288Request(Buffer string) (error, *Rfc1288Request) {
	if pos := strings.IndexAny(Buffer, "\r\n"); pos > 0 {
		Buffer = Buffer[:pos]
	}
	
	Buflen := len(Buffer)

	if Buflen < 2 {
		return errors.New("Protocol not recognized"), nil
	}

	if Buffer[0] != '/' || (Buffer[1] != 'W' && Buffer[1] != 'w') {
		return errors.New("Protocol not recognized"), nil
	}

	if len(Buffer) == 2 {
		return nil, &Rfc1288Request{UserList, nil, nil}
	}

	index := 2
	for index < Buflen && Buffer[index] == ' ' {
		index += 1
	}

	if Buflen == index {
		return nil, &Rfc1288Request{Type: UserList, User: nil, Host: nil}
	}

	user := bytes.NewBufferString("")
	host := bytes.NewBufferString("")

	for index < Buflen && is_unix_conventional(Buffer[index]) {
		user.WriteByte(Buffer[index])
		index += 1
	}

	if index == Buflen || (index < Buflen && Buffer[index] == ' ') {
		ruser := user.String()
		return nil, &Rfc1288Request{Type: User, User: &ruser, Host: nil}
	}

	if Buffer[index] != '@' {
		return errors.New("Protocol does not meet specification"), nil
	}

	index += 1
	for index < Buflen && Buffer[index] != ' ' {
		host.WriteByte(Buffer[index])
		index += 1
	}

	ruser := user.String()
	rhost := host.String()
	return nil, &Rfc1288Request{Type: Remote, User: &ruser, Host: &rhost}
}
