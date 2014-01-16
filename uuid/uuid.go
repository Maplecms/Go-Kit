package uuid

import (
	"crypto/rand"
	"encoding/hex"
_	"encoding"
	"regexp"
	"strings"
)

// Unique Identifiers intend to conform to RFC4122. UUID Version 4.
type UniqueID struct {
	bytes []byte
}

const (
	UniqueIDVersion = 4
)

var (
	UniqueIDRegex = regexp.MustCompile("[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[4][0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}")
)

// NewSafe returns a UniqueID or an error if it cannot read random bytes from the system.
func NewSafe() (UniqueID, error) {
	uuid := make([]byte, 16)
	if _, err := rand.Read(uuid); err != nil {
		return nil, err
	}
	uuid[6] = uuid[6] & 0x0F + 0x40
	uuid[8] = uuid[8] & 0x3F + 0x80
	return UniqueID{uuid}, nil
}

// New returns a UniqueID. It panics if it cannot read random bytes from the system.
func New() UniqueID {
	uuid, err := New()
	if err != nil {
		panic(err)
	}
	return uuid
}

// UniqueID implements the Stringer interface
func (self UniqueID) String() string {
	encode := hex.EncodeToString
	bytes := self.bytes
	return encode(bytes[0:4]) + "-" + encode(bytes[4:6]) + "-" + encode(bytes[6:8]) + "-" +
		encode(bytes[8:10]) + "-" + encode(bytes[10:16])
}

// UniqueID implements the TextMarshaler interface
func (self UniqueID) MarshalText() (text []byte, err error) {
	return []byte(self.String()), nil
}

// UniqueID implements the TextUnmarshaler interface
func (self *UniqueID) UnmarshalText(text []byte) error {
	hexString := strings.Replace(string(text), "-", "", -1)
	hexDigits, err := hex.DecodeString(hexString)
	if err != nil {
		return err
	}
	self.bytes = make([]byte, 16)
	copy(self.bytes, hexDigits[:16])
	return nil
}
