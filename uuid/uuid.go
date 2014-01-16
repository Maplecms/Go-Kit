package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strings"
)

// Unique Identifiers conform to RFC4122. UUID Version 4.
type UniqueId struct {
	bytes []byte
}

const (
	UniqueIdVersion = 4
)

var (
	UniqueIdRegex = regexp.MustCompile("[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[4][0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}")
)

// NewSafe returns a UniqueId or an error if it cannot read random bytes from the system.
func NewSafe() (UniqueId, error) {
	uuid := make([]byte, 16)
	if _, err := rand.Read(uuid); err != nil {
		return *new(UniqueId), err
	}
	uuid[6] = uuid[6] & 0x0F + 0x40
	uuid[8] = uuid[8] & 0x3F + 0x80
	return UniqueId{uuid}, nil
}

// New returns a UniqueId. It panics if it cannot read random bytes from the system.
func New() UniqueId {
	uuid, err := NewSafe()
	if err != nil {
		panic(err)
	}
	return uuid
}

// UniqueId implements the Stringer interface
func (self UniqueId) String() string {
	encode := hex.EncodeToString
	bytes := self.bytes
	return encode(bytes[0:4]) + "-" + encode(bytes[4:6]) + "-" + encode(bytes[6:8]) + "-" +
		encode(bytes[8:10]) + "-" + encode(bytes[10:16])
}

// UniqueId implements the TextMarshaler interface
func (self UniqueId) MarshalText() (text []byte, err error) {
	return []byte(self.String()), nil
}

// UniqueId implements the TextUnmarshaler interface
func (self *UniqueId) UnmarshalText(text []byte) error {
	hexString := strings.Replace(string(text), "-", "", -1)
	hexDigits, err := hex.DecodeString(hexString)
	if err != nil {
		return err
	}
	self.bytes = make([]byte, 16)
	copy(self.bytes, hexDigits[:16])
	return nil
}
