package lexicon

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func Digest(value string) string {
	sum := sha256.Sum256([]byte(value))
	return hex.EncodeToString(sum[:])
}
func RevisionDigest(value string) string {
	return Digest(value + " ")
}
func KeyOf(prefix string, parts ...string) Key {
	return Key(prefix + "-" + Digest(strings.Join(parts, "|"))[:14])
}
func Normalize(value string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(value)), " ")
}
