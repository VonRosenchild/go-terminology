package snomed

import (
	"fmt"
	"strconv"

	"bitbucket.org/wardle/go-snomed/verhoeff"
)

// SctID is an checksummed (Verhoeff) globally unique persistent identifier
// See https://confluence.ihtsdotools.org/display/DOCTIG/3.1.4.2.+Component+features+-+Identifiers
// The SCTID data type is 64-bit Integer which is allocated and represented in accordance with a set of rules.
// These rules enable each Identifier to refer unambiguously to a unique component.
// They also support separate partitions for allocation of Identifiers for particular types of component and
// namespaces that distinguish between different issuing organizations.
type SctID int

// IsConcept will return true if this identifier refers to a concept
// TODO: add implementation
func (id SctID) IsConcept() bool {
	pid := id.partitionIdentifier()
	return pid == "00" || pid == "10"
}

// IsDescription will return true if this identifier refers to a description.
// TODO: add implementation
func (id SctID) IsDescription() bool {
	pid := id.partitionIdentifier()
	return pid == "01" || pid == "11"
}

// IsRelationship will return true if this identifier refers to a relationship.
// TODO: add implementation
func (id SctID) isRelationship() bool {
	pid := id.partitionIdentifier()
	return pid == "02" || pid == "12"
}

// IsValid will return true if this is a valid SNOMED CT identifier
func (id SctID) IsValid() bool {
	s := strconv.Itoa(int(id))
	return verhoeff.ValidateVerhoeffString(s)
}

// partitionIdentifier returns the penultimate last digit digits
// see https://confluence.ihtsdotools.org/display/DOCRELFMT/5.5.+Partition+Identifier
// 0123456789
// xxxxxxxppc
func (id SctID) partitionIdentifier() string {
	fmt.Printf("Converting identifier: %v\n", id)
	s := strconv.Itoa(int(id))
	fmt.Printf("Converting to string: %s", s)
	l := len(s)
	return s[l-3 : l-1]
}
