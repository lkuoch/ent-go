package pulid

import (
	"crypto/rand"
	"database/sql/driver"
	"fmt"
	"hash/fnv"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

// PULID is a ULID prefixed with the hash of the table name to support the node interface
type ID string

var defaultEntropySource *ulid.MonotonicEntropy

func hash64(input string) uint64 {
	hasher := fnv.New64a()
	hasher.Write([]byte(input))

	return hasher.Sum64()
}

func init() {
	// Seed the default entropy source
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}

// Generates a new ULID from the table name
// Accepts a tableName which will then be hashed (FNV) to the prefix of the ULID
func New(tableName string) ID {
	prefix := NewPrefix(tableName)
	ulid := ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource).String()

	return ID(prefix + ulid)
}

func NewPrefix(tableName string) string {
	return fmt.Sprintf("%x", hash64(tableName))
}

func ExtractPrefixHash(id ID) string {
	tuple := strings.Split(string(id), ":")
	prefix := tuple[0]

	return prefix
}

// Scan implements the Scanner interface.
func (u *ID) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("pulid: expected a value")
	}
	switch src := src.(type) {
	case string:
		*u = ID(src)
	case ID:
		*u = src
	default:
		return fmt.Errorf("pulid: unexpected type, %T", src)
	}
	return nil
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *ID) UnmarshalGQL(v interface{}) error {
	return u.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface
func (u ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(u)))
}

// Value implements the driver Valuer interface.
func (u ID) Value() (driver.Value, error) {
	return string(u), nil
}
