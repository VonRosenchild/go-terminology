// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snomed.proto

/*
Package snomed is a generated protocol buffer package.

It is generated from these files:
	snomed.proto

It has these top-level messages:
	Concept
	Description
	Relationship
	ReferenceSetItem
	RefSetDescriptorReferenceSet
	SimpleReferenceSet
	LanguageReferenceSet
	SimpleMapReferenceSet
	ComplexMapReferenceSet
	ExtendedDescription
*/
package snomed

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Concept represents a SNOMED-CT concept.
// The RF2 release allows multiple duplicate entries per concept identifier to permit versioning.
// As such, we have a compound primary key made up of the concept identifier and the effective time.
// Only one concept with a specified identifier will be active at any time point.
// See https://confluence.ihtsdotools.org/display/DOCRELFMT/3.2.1.+Concept+File+Specification
type Concept struct {
	Id                 int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	EffectiveTime      *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=effective_time,json=effectiveTime" json:"effective_time,omitempty"`
	Active             bool                       `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
	ModuleId           int64                      `protobuf:"varint,4,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
	DefinitionStatusId int64                      `protobuf:"varint,5,opt,name=definition_status_id,json=definitionStatusId" json:"definition_status_id,omitempty"`
}

func (m *Concept) Reset()                    { *m = Concept{} }
func (m *Concept) String() string            { return proto.CompactTextString(m) }
func (*Concept) ProtoMessage()               {}
func (*Concept) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Concept) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Concept) GetEffectiveTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.EffectiveTime
	}
	return nil
}

func (m *Concept) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Concept) GetModuleId() int64 {
	if m != nil {
		return m.ModuleId
	}
	return 0
}

func (m *Concept) GetDefinitionStatusId() int64 {
	if m != nil {
		return m.DefinitionStatusId
	}
	return 0
}

// A Description holds descriptions that describe SNOMED CT concepts.
// A description is used to give meaning to a concept and provide well-understood and standard ways of referring to a concept.
// See https://confluence.ihtsdotools.org/display/DOCRELFMT/3.2.2.+Description+File+Specification
type Description struct {
	Id               int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	EffectiveTime    *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=effective_time,json=effectiveTime" json:"effective_time,omitempty"`
	Active           bool                       `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
	ModuleId         int64                      `protobuf:"varint,4,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
	ConceptId        int64                      `protobuf:"varint,5,opt,name=concept_id,json=conceptId" json:"concept_id,omitempty"`
	LanguageCode     string                     `protobuf:"bytes,6,opt,name=LanguageCode" json:"LanguageCode,omitempty"`
	TypeId           int64                      `protobuf:"varint,7,opt,name=type_id,json=typeId" json:"type_id,omitempty"`
	Term             string                     `protobuf:"bytes,8,opt,name=term" json:"term,omitempty"`
	CaseSignificance int64                      `protobuf:"varint,9,opt,name=case_significance,json=caseSignificance" json:"case_significance,omitempty"`
}

func (m *Description) Reset()                    { *m = Description{} }
func (m *Description) String() string            { return proto.CompactTextString(m) }
func (*Description) ProtoMessage()               {}
func (*Description) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Description) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Description) GetEffectiveTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.EffectiveTime
	}
	return nil
}

func (m *Description) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Description) GetModuleId() int64 {
	if m != nil {
		return m.ModuleId
	}
	return 0
}

func (m *Description) GetConceptId() int64 {
	if m != nil {
		return m.ConceptId
	}
	return 0
}

func (m *Description) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *Description) GetTypeId() int64 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *Description) GetTerm() string {
	if m != nil {
		return m.Term
	}
	return ""
}

func (m *Description) GetCaseSignificance() int64 {
	if m != nil {
		return m.CaseSignificance
	}
	return 0
}

// Relationship defines a relationship between two concepts as a type itself defined as a concept
type Relationship struct {
	Id                   int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	EffectiveTime        *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=effective_time,json=effectiveTime" json:"effective_time,omitempty"`
	Active               bool                       `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
	ModuleId             int64                      `protobuf:"varint,4,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
	SourceId             int64                      `protobuf:"varint,5,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	DestinationId        int64                      `protobuf:"varint,6,opt,name=destination_id,json=destinationId" json:"destination_id,omitempty"`
	RelationshipGroup    int64                      `protobuf:"varint,7,opt,name=relationship_group,json=relationshipGroup" json:"relationship_group,omitempty"`
	TypeId               int64                      `protobuf:"varint,8,opt,name=type_id,json=typeId" json:"type_id,omitempty"`
	CharacteristicTypeId int64                      `protobuf:"varint,9,opt,name=characteristic_type_id,json=characteristicTypeId" json:"characteristic_type_id,omitempty"`
	ModifierId           int64                      `protobuf:"varint,10,opt,name=modifier_id,json=modifierId" json:"modifier_id,omitempty"`
}

func (m *Relationship) Reset()                    { *m = Relationship{} }
func (m *Relationship) String() string            { return proto.CompactTextString(m) }
func (*Relationship) ProtoMessage()               {}
func (*Relationship) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Relationship) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Relationship) GetEffectiveTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.EffectiveTime
	}
	return nil
}

func (m *Relationship) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Relationship) GetModuleId() int64 {
	if m != nil {
		return m.ModuleId
	}
	return 0
}

func (m *Relationship) GetSourceId() int64 {
	if m != nil {
		return m.SourceId
	}
	return 0
}

func (m *Relationship) GetDestinationId() int64 {
	if m != nil {
		return m.DestinationId
	}
	return 0
}

func (m *Relationship) GetRelationshipGroup() int64 {
	if m != nil {
		return m.RelationshipGroup
	}
	return 0
}

func (m *Relationship) GetTypeId() int64 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *Relationship) GetCharacteristicTypeId() int64 {
	if m != nil {
		return m.CharacteristicTypeId
	}
	return 0
}

func (m *Relationship) GetModifierId() int64 {
	if m != nil {
		return m.ModifierId
	}
	return 0
}

// ReferenceSet support customization and enhancement of SNOMED CT content. These include representation of subsets,
// language preferences maps for or from other code systems.
// There are multiple reference set types which extend this structure
// In the specification, the referenced component ID can be a SCT identifier or a UUID which is... problematic.
// In this structure, the referenced component ID is a SCT identifier... only. For now.
// Fortunately, in concrete types of reference set ("patterns"), it is made explicit.
type ReferenceSetItem struct {
	Id                    string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	EffectiveTime         *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=effective_time,json=effectiveTime" json:"effective_time,omitempty"`
	Active                bool                       `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
	ModuleId              int64                      `protobuf:"varint,4,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
	RefsetId              int64                      `protobuf:"varint,5,opt,name=refset_id,json=refsetId" json:"refset_id,omitempty"`
	ReferencedComponentId int64                      `protobuf:"varint,6,opt,name=referenced_component_id,json=referencedComponentId" json:"referenced_component_id,omitempty"`
	// Types that are valid to be assigned to Body:
	//	*ReferenceSetItem_RefsetDescriptor
	//	*ReferenceSetItem_Simple
	//	*ReferenceSetItem_Language
	//	*ReferenceSetItem_SimpleMap
	Body isReferenceSetItem_Body `protobuf_oneof:"body"`
}

func (m *ReferenceSetItem) Reset()                    { *m = ReferenceSetItem{} }
func (m *ReferenceSetItem) String() string            { return proto.CompactTextString(m) }
func (*ReferenceSetItem) ProtoMessage()               {}
func (*ReferenceSetItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isReferenceSetItem_Body interface {
	isReferenceSetItem_Body()
}

type ReferenceSetItem_RefsetDescriptor struct {
	RefsetDescriptor *RefSetDescriptorReferenceSet `protobuf:"bytes,7,opt,name=refset_descriptor,json=refsetDescriptor,oneof"`
}
type ReferenceSetItem_Simple struct {
	Simple *SimpleReferenceSet `protobuf:"bytes,8,opt,name=simple,oneof"`
}
type ReferenceSetItem_Language struct {
	Language *LanguageReferenceSet `protobuf:"bytes,9,opt,name=language,oneof"`
}
type ReferenceSetItem_SimpleMap struct {
	SimpleMap *SimpleMapReferenceSet `protobuf:"bytes,10,opt,name=simple_map,json=simpleMap,oneof"`
}

func (*ReferenceSetItem_RefsetDescriptor) isReferenceSetItem_Body() {}
func (*ReferenceSetItem_Simple) isReferenceSetItem_Body()           {}
func (*ReferenceSetItem_Language) isReferenceSetItem_Body()         {}
func (*ReferenceSetItem_SimpleMap) isReferenceSetItem_Body()        {}

func (m *ReferenceSetItem) GetBody() isReferenceSetItem_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *ReferenceSetItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReferenceSetItem) GetEffectiveTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.EffectiveTime
	}
	return nil
}

func (m *ReferenceSetItem) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *ReferenceSetItem) GetModuleId() int64 {
	if m != nil {
		return m.ModuleId
	}
	return 0
}

func (m *ReferenceSetItem) GetRefsetId() int64 {
	if m != nil {
		return m.RefsetId
	}
	return 0
}

func (m *ReferenceSetItem) GetReferencedComponentId() int64 {
	if m != nil {
		return m.ReferencedComponentId
	}
	return 0
}

func (m *ReferenceSetItem) GetRefsetDescriptor() *RefSetDescriptorReferenceSet {
	if x, ok := m.GetBody().(*ReferenceSetItem_RefsetDescriptor); ok {
		return x.RefsetDescriptor
	}
	return nil
}

func (m *ReferenceSetItem) GetSimple() *SimpleReferenceSet {
	if x, ok := m.GetBody().(*ReferenceSetItem_Simple); ok {
		return x.Simple
	}
	return nil
}

func (m *ReferenceSetItem) GetLanguage() *LanguageReferenceSet {
	if x, ok := m.GetBody().(*ReferenceSetItem_Language); ok {
		return x.Language
	}
	return nil
}

func (m *ReferenceSetItem) GetSimpleMap() *SimpleMapReferenceSet {
	if x, ok := m.GetBody().(*ReferenceSetItem_SimpleMap); ok {
		return x.SimpleMap
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReferenceSetItem) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReferenceSetItem_OneofMarshaler, _ReferenceSetItem_OneofUnmarshaler, _ReferenceSetItem_OneofSizer, []interface{}{
		(*ReferenceSetItem_RefsetDescriptor)(nil),
		(*ReferenceSetItem_Simple)(nil),
		(*ReferenceSetItem_Language)(nil),
		(*ReferenceSetItem_SimpleMap)(nil),
	}
}

func _ReferenceSetItem_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReferenceSetItem)
	// body
	switch x := m.Body.(type) {
	case *ReferenceSetItem_RefsetDescriptor:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RefsetDescriptor); err != nil {
			return err
		}
	case *ReferenceSetItem_Simple:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Simple); err != nil {
			return err
		}
	case *ReferenceSetItem_Language:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Language); err != nil {
			return err
		}
	case *ReferenceSetItem_SimpleMap:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SimpleMap); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ReferenceSetItem.Body has unexpected type %T", x)
	}
	return nil
}

func _ReferenceSetItem_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReferenceSetItem)
	switch tag {
	case 7: // body.refset_descriptor
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RefSetDescriptorReferenceSet)
		err := b.DecodeMessage(msg)
		m.Body = &ReferenceSetItem_RefsetDescriptor{msg}
		return true, err
	case 8: // body.simple
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SimpleReferenceSet)
		err := b.DecodeMessage(msg)
		m.Body = &ReferenceSetItem_Simple{msg}
		return true, err
	case 9: // body.language
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LanguageReferenceSet)
		err := b.DecodeMessage(msg)
		m.Body = &ReferenceSetItem_Language{msg}
		return true, err
	case 10: // body.simple_map
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SimpleMapReferenceSet)
		err := b.DecodeMessage(msg)
		m.Body = &ReferenceSetItem_SimpleMap{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ReferenceSetItem_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReferenceSetItem)
	// body
	switch x := m.Body.(type) {
	case *ReferenceSetItem_RefsetDescriptor:
		s := proto.Size(x.RefsetDescriptor)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReferenceSetItem_Simple:
		s := proto.Size(x.Simple)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReferenceSetItem_Language:
		s := proto.Size(x.Language)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReferenceSetItem_SimpleMap:
		s := proto.Size(x.SimpleMap)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// RefSetDescriptorReferenceSet is a type of reference set that provides information about a different reference set
// See https://confluence.ihtsdotools.org/display/DOCRELFMT/4.2.11.+Reference+Set+Descriptor
// It provides the additional structure for a given reference set.
type RefSetDescriptorReferenceSet struct {
	AttributeDescriptionId int64  `protobuf:"varint,1,opt,name=attribute_description_id,json=attributeDescriptionId" json:"attribute_description_id,omitempty"`
	AttributeTypeId        int64  `protobuf:"varint,2,opt,name=attribute_type_id,json=attributeTypeId" json:"attribute_type_id,omitempty"`
	AttributeOrder         uint32 `protobuf:"varint,3,opt,name=attribute_order,json=attributeOrder" json:"attribute_order,omitempty"`
}

func (m *RefSetDescriptorReferenceSet) Reset()                    { *m = RefSetDescriptorReferenceSet{} }
func (m *RefSetDescriptorReferenceSet) String() string            { return proto.CompactTextString(m) }
func (*RefSetDescriptorReferenceSet) ProtoMessage()               {}
func (*RefSetDescriptorReferenceSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RefSetDescriptorReferenceSet) GetAttributeDescriptionId() int64 {
	if m != nil {
		return m.AttributeDescriptionId
	}
	return 0
}

func (m *RefSetDescriptorReferenceSet) GetAttributeTypeId() int64 {
	if m != nil {
		return m.AttributeTypeId
	}
	return 0
}

func (m *RefSetDescriptorReferenceSet) GetAttributeOrder() uint32 {
	if m != nil {
		return m.AttributeOrder
	}
	return 0
}

// SimpleReferenceSet is a simple reference set usable for defining subsets
// See https://confluence.ihtsdotools.org/display/DOCRELFMT/4.2.1.+Simple+Reference+Set
type SimpleReferenceSet struct {
}

func (m *SimpleReferenceSet) Reset()                    { *m = SimpleReferenceSet{} }
func (m *SimpleReferenceSet) String() string            { return proto.CompactTextString(m) }
func (*SimpleReferenceSet) ProtoMessage()               {}
func (*SimpleReferenceSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// LanguageReferenceSet is a A 900000000000506000 |Language type reference set| supporting the representation of
// language and dialects preferences for the use of particular descriptions.
// "The most common use case for this type of reference set is to specify the acceptable and preferred terms
// for use within a particular country or region. However, the same type of reference set can also be used to
// represent preferences for use of descriptions in a more specific context such as a clinical specialty,
// organization or department.
//
// No more than one description of a specific description type associated with a single concept may have the acceptabilityId value 900000000000548007 |Preferred|.
// Every active concept should have one preferred synonym in each language.
// This means that a language reference set should assign the acceptabilityId  900000000000548007 |Preferred|  to one  synonym (a  description with  typeId value 900000000000013009 |synonym|) associated with each concept .
// This description is the preferred term for that concept in the specified language or dialect.
// Any  description which is not referenced by an active row in the   reference set is regarded as unacceptable (i.e. not a valid  synonym in the language or  dialect ).
// If a description becomes unacceptable, the relevant language reference set member is inactivated by adding a new row with the same id, the effectiveTime of the the change and the value active=0.
// For this reason there is no requirement for an "unacceptable" value."
// See https://confluence.ihtsdotools.org/display/DOCRELFMT/4.2.4.+Language+Reference+Set
//
type LanguageReferenceSet struct {
	AcceptabilityId int64 `protobuf:"varint,1,opt,name=acceptability_id,json=acceptabilityId" json:"acceptability_id,omitempty"`
}

func (m *LanguageReferenceSet) Reset()                    { *m = LanguageReferenceSet{} }
func (m *LanguageReferenceSet) String() string            { return proto.CompactTextString(m) }
func (*LanguageReferenceSet) ProtoMessage()               {}
func (*LanguageReferenceSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LanguageReferenceSet) GetAcceptabilityId() int64 {
	if m != nil {
		return m.AcceptabilityId
	}
	return 0
}

// SimpleMapReferenceSet is a straightforward one-to-one map between SNOMED-CT concepts and another
// coding system. This is appropriate for simple maps.
// See https://confluence.ihtsdotools.org/display/DOCRELFMT/4.2.9.+Simple+Map+Reference+Set
type SimpleMapReferenceSet struct {
	MapTarget string `protobuf:"bytes,1,opt,name=map_target,json=mapTarget" json:"map_target,omitempty"`
}

func (m *SimpleMapReferenceSet) Reset()                    { *m = SimpleMapReferenceSet{} }
func (m *SimpleMapReferenceSet) String() string            { return proto.CompactTextString(m) }
func (*SimpleMapReferenceSet) ProtoMessage()               {}
func (*SimpleMapReferenceSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SimpleMapReferenceSet) GetMapTarget() string {
	if m != nil {
		return m.MapTarget
	}
	return ""
}

// ComplexMapReferenceSet represents a complex one-to-many map between SNOMED-CT and another
// coding system.
// A 447250001 |Complex map type reference set|enables representation of maps where each SNOMED
// CT concept may map to one or more codes in a target scheme.
// The type of reference set supports the general set of mapping data required to enable a
// target code to be selected at run-time from a number of alternate codes. It supports
// target code selection by accommodating the inclusion of machine readable rules and/or human readable advice.
// An 609331003 |Extended map type reference set|adds an additional field to allow categorization of maps.
type ComplexMapReferenceSet struct {
	MapGroup      int64  `protobuf:"varint,1,opt,name=map_group,json=mapGroup" json:"map_group,omitempty"`
	MapPriority   int64  `protobuf:"varint,2,opt,name=map_priority,json=mapPriority" json:"map_priority,omitempty"`
	MapRule       string `protobuf:"bytes,3,opt,name=map_rule,json=mapRule" json:"map_rule,omitempty"`
	MapAdvice     string `protobuf:"bytes,4,opt,name=map_advice,json=mapAdvice" json:"map_advice,omitempty"`
	MapTarget     string `protobuf:"bytes,5,opt,name=map_target,json=mapTarget" json:"map_target,omitempty"`
	CorrectionId  int64  `protobuf:"varint,6,opt,name=correction_id,json=correctionId" json:"correction_id,omitempty"`
	MapCategoryId int64  `protobuf:"varint,7,opt,name=map_category_id,json=mapCategoryId" json:"map_category_id,omitempty"`
}

func (m *ComplexMapReferenceSet) Reset()                    { *m = ComplexMapReferenceSet{} }
func (m *ComplexMapReferenceSet) String() string            { return proto.CompactTextString(m) }
func (*ComplexMapReferenceSet) ProtoMessage()               {}
func (*ComplexMapReferenceSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ComplexMapReferenceSet) GetMapGroup() int64 {
	if m != nil {
		return m.MapGroup
	}
	return 0
}

func (m *ComplexMapReferenceSet) GetMapPriority() int64 {
	if m != nil {
		return m.MapPriority
	}
	return 0
}

func (m *ComplexMapReferenceSet) GetMapRule() string {
	if m != nil {
		return m.MapRule
	}
	return ""
}

func (m *ComplexMapReferenceSet) GetMapAdvice() string {
	if m != nil {
		return m.MapAdvice
	}
	return ""
}

func (m *ComplexMapReferenceSet) GetMapTarget() string {
	if m != nil {
		return m.MapTarget
	}
	return ""
}

func (m *ComplexMapReferenceSet) GetCorrectionId() int64 {
	if m != nil {
		return m.CorrectionId
	}
	return 0
}

func (m *ComplexMapReferenceSet) GetMapCategoryId() int64 {
	if m != nil {
		return m.MapCategoryId
	}
	return 0
}

// ExtendedDescription represents a description together with
// sufficient additional contextual information relating to the
// description, including reference set membership as well as
// the underlying concept, the concept's relationships and
// the concept's membership of reference sets.
// It is, in essence, a denormalised relationship, useful for
// wire-exchange purposes.
// TODO: add language field
type ExtendedDescription struct {
	Description          *Description `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Concept              *Concept     `protobuf:"bytes,3,opt,name=concept" json:"concept,omitempty"`
	PreferredDescription *Description `protobuf:"bytes,4,opt,name=preferred_description,json=preferredDescription" json:"preferred_description,omitempty"`
	RecursiveParentIds   []int64      `protobuf:"varint,5,rep,packed,name=recursive_parent_ids,json=recursiveParentIds" json:"recursive_parent_ids,omitempty"`
	DirectParentIds      []int64      `protobuf:"varint,6,rep,packed,name=direct_parent_ids,json=directParentIds" json:"direct_parent_ids,omitempty"`
	ConceptRefsets       []int64      `protobuf:"varint,7,rep,packed,name=concept_refsets,json=conceptRefsets" json:"concept_refsets,omitempty"`
	DescriptionRefsets   []int64      `protobuf:"varint,8,rep,packed,name=description_refsets,json=descriptionRefsets" json:"description_refsets,omitempty"`
}

func (m *ExtendedDescription) Reset()                    { *m = ExtendedDescription{} }
func (m *ExtendedDescription) String() string            { return proto.CompactTextString(m) }
func (*ExtendedDescription) ProtoMessage()               {}
func (*ExtendedDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ExtendedDescription) GetDescription() *Description {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *ExtendedDescription) GetConcept() *Concept {
	if m != nil {
		return m.Concept
	}
	return nil
}

func (m *ExtendedDescription) GetPreferredDescription() *Description {
	if m != nil {
		return m.PreferredDescription
	}
	return nil
}

func (m *ExtendedDescription) GetRecursiveParentIds() []int64 {
	if m != nil {
		return m.RecursiveParentIds
	}
	return nil
}

func (m *ExtendedDescription) GetDirectParentIds() []int64 {
	if m != nil {
		return m.DirectParentIds
	}
	return nil
}

func (m *ExtendedDescription) GetConceptRefsets() []int64 {
	if m != nil {
		return m.ConceptRefsets
	}
	return nil
}

func (m *ExtendedDescription) GetDescriptionRefsets() []int64 {
	if m != nil {
		return m.DescriptionRefsets
	}
	return nil
}

func init() {
	proto.RegisterType((*Concept)(nil), "snomed.Concept")
	proto.RegisterType((*Description)(nil), "snomed.Description")
	proto.RegisterType((*Relationship)(nil), "snomed.Relationship")
	proto.RegisterType((*ReferenceSetItem)(nil), "snomed.ReferenceSetItem")
	proto.RegisterType((*RefSetDescriptorReferenceSet)(nil), "snomed.RefSetDescriptorReferenceSet")
	proto.RegisterType((*SimpleReferenceSet)(nil), "snomed.SimpleReferenceSet")
	proto.RegisterType((*LanguageReferenceSet)(nil), "snomed.LanguageReferenceSet")
	proto.RegisterType((*SimpleMapReferenceSet)(nil), "snomed.SimpleMapReferenceSet")
	proto.RegisterType((*ComplexMapReferenceSet)(nil), "snomed.ComplexMapReferenceSet")
	proto.RegisterType((*ExtendedDescription)(nil), "snomed.ExtendedDescription")
}

func init() { proto.RegisterFile("snomed.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 986 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0xad, 0x64, 0x5b, 0x1f, 0x43, 0xdb, 0xb2, 0xd7, 0xb2, 0xc3, 0xe6, 0x03, 0x76, 0xd9, 0x8f,
	0x38, 0x2d, 0x2a, 0xa7, 0x6a, 0x1a, 0x14, 0x3d, 0x14, 0xb0, 0x9d, 0xa2, 0x61, 0xd1, 0xa2, 0x06,
	0xe5, 0x53, 0x2f, 0xc4, 0x7a, 0x77, 0xa4, 0x2c, 0x40, 0x72, 0x89, 0xe5, 0x32, 0xb0, 0x7e, 0x54,
	0x2f, 0xfd, 0x0b, 0xfd, 0x55, 0xbd, 0xa4, 0xc5, 0x2e, 0x97, 0x22, 0xad, 0xa6, 0xbd, 0x26, 0x37,
	0xf1, 0xbd, 0x37, 0xc3, 0xe1, 0xdb, 0x99, 0x59, 0xc1, 0x76, 0x91, 0xc9, 0x14, 0xf9, 0x24, 0x57,
	0x52, 0x4b, 0xd2, 0xab, 0x9e, 0xee, 0x1f, 0x2f, 0xa4, 0x5c, 0x24, 0x78, 0x66, 0xd1, 0x9b, 0x72,
	0x7e, 0xa6, 0x45, 0x8a, 0x85, 0xa6, 0x69, 0x5e, 0x09, 0x83, 0x3f, 0x3b, 0xd0, 0xbf, 0x94, 0x19,
	0xc3, 0x5c, 0x93, 0x5d, 0xe8, 0x0a, 0xee, 0x77, 0x4e, 0x3a, 0xa7, 0x1b, 0x51, 0x57, 0x70, 0x72,
	0x0e, 0xbb, 0x38, 0x9f, 0x23, 0xd3, 0xe2, 0x35, 0xc6, 0x26, 0xd0, 0xef, 0x9e, 0x74, 0x4e, 0xbd,
	0xe9, 0xfd, 0x49, 0x95, 0x75, 0x52, 0x67, 0x9d, 0x5c, 0xd7, 0x59, 0xa3, 0x9d, 0x55, 0x84, 0xc1,
	0xc8, 0x11, 0xf4, 0xa8, 0x7d, 0xf2, 0x37, 0x4e, 0x3a, 0xa7, 0x83, 0xc8, 0x3d, 0x91, 0x07, 0x30,
	0x4c, 0x25, 0x2f, 0x13, 0x8c, 0x05, 0xf7, 0x37, 0xed, 0x1b, 0x07, 0x15, 0x10, 0x72, 0xf2, 0x14,
	0xc6, 0x1c, 0xe7, 0x22, 0x13, 0x5a, 0xc8, 0x2c, 0x2e, 0x34, 0xd5, 0x65, 0x61, 0x74, 0x5b, 0x56,
	0x47, 0x1a, 0x6e, 0x66, 0xa9, 0x90, 0x07, 0x7f, 0x74, 0xc1, 0x7b, 0x81, 0x05, 0x53, 0x22, 0x37,
	0xf8, 0x7b, 0xf3, 0x25, 0x8f, 0x00, 0x58, 0x65, 0x6e, 0x53, 0xff, 0xd0, 0x21, 0x21, 0x27, 0x01,
	0x6c, 0xff, 0x4c, 0xb3, 0x45, 0x49, 0x17, 0x78, 0x29, 0x39, 0xfa, 0xbd, 0x93, 0xce, 0xe9, 0x30,
	0xba, 0x83, 0x91, 0x7b, 0xd0, 0xd7, 0xcb, 0xdc, 0x66, 0xef, 0xdb, 0xf8, 0x9e, 0x79, 0x0c, 0x39,
	0x21, 0xb0, 0xa9, 0x51, 0xa5, 0xfe, 0xc0, 0x06, 0xd9, 0xdf, 0xe4, 0x0b, 0xd8, 0x67, 0xb4, 0xc0,
	0xb8, 0x10, 0x8b, 0x4c, 0xcc, 0x05, 0xa3, 0x19, 0x43, 0x7f, 0x68, 0xc3, 0xf6, 0x0c, 0x31, 0x6b,
	0xe1, 0xc1, 0x5f, 0x5d, 0xd8, 0x8e, 0x30, 0xa1, 0xc6, 0xb1, 0xe2, 0x95, 0xc8, 0xdf, 0x1b, 0xd7,
	0x1e, 0xc0, 0xb0, 0x90, 0xa5, 0x62, 0xd8, 0x98, 0x36, 0xa8, 0x80, 0x90, 0x93, 0x4f, 0x61, 0x97,
	0x63, 0xa1, 0x45, 0x66, 0xeb, 0x36, 0x8a, 0x9e, 0x55, 0xec, 0xb4, 0xd0, 0x90, 0x93, 0x2f, 0x81,
	0xa8, 0xd6, 0xb7, 0xc5, 0x0b, 0x25, 0xcb, 0xdc, 0x39, 0xb8, 0xdf, 0x66, 0x7e, 0x34, 0x44, 0xdb,
	0xe5, 0xc1, 0x1d, 0x97, 0x9f, 0xc1, 0x11, 0x7b, 0x45, 0x15, 0x65, 0x1a, 0x95, 0x28, 0xb4, 0x60,
	0x71, 0xad, 0xab, 0x6c, 0x1d, 0xdf, 0x65, 0xaf, 0xab, 0xa8, 0x63, 0xf0, 0x52, 0xc9, 0xc5, 0x5c,
	0xa0, 0x32, 0x52, 0xb0, 0x52, 0xa8, 0xa1, 0x90, 0x07, 0x6f, 0x36, 0x60, 0x2f, 0xc2, 0x39, 0x2a,
	0xcc, 0x18, 0xce, 0x50, 0x87, 0x1a, 0xd3, 0x96, 0xff, 0xc3, 0x77, 0xed, 0xbf, 0xc2, 0x79, 0x81,
	0xad, 0xa6, 0x1d, 0x54, 0x40, 0xc8, 0xc9, 0x73, 0xb8, 0xa7, 0xea, 0xc2, 0x79, 0xcc, 0x64, 0x9a,
	0xcb, 0x0c, 0x33, 0xdd, 0x1c, 0xc4, 0x61, 0x43, 0x5f, 0xd6, 0x6c, 0xc8, 0xc9, 0x0c, 0xf6, 0x5d,
	0x52, 0xee, 0x06, 0x55, 0x2a, 0x7b, 0x1e, 0xde, 0xf4, 0x93, 0x89, 0xdb, 0x5d, 0x11, 0xce, 0x67,
	0xa8, 0x5f, 0xac, 0xf8, 0xb6, 0x43, 0x2f, 0x3f, 0x88, 0xf6, 0xaa, 0x04, 0x0d, 0x4f, 0x9e, 0x41,
	0xaf, 0x10, 0x69, 0x9e, 0xa0, 0x3d, 0x35, 0xe3, 0x8c, 0xcb, 0x34, 0xb3, 0xe8, 0x5a, 0xbc, 0xd3,
	0x92, 0xef, 0x60, 0x90, 0xb8, 0x11, 0xb3, 0xa7, 0xe8, 0x4d, 0x1f, 0xd6, 0x71, 0xf5, 0xe8, 0xad,
	0x45, 0xae, 0xf4, 0xe4, 0x7b, 0x80, 0x2a, 0x4b, 0x9c, 0xd2, 0xdc, 0x1e, 0xac, 0x37, 0x7d, 0x74,
	0xf7, 0xad, 0xbf, 0xd0, 0x7c, 0x2d, 0x7c, 0x58, 0xd4, 0xc4, 0x45, 0x0f, 0x36, 0x6f, 0x24, 0x5f,
	0x06, 0xbf, 0x77, 0xe0, 0xe1, 0xff, 0x7d, 0x2e, 0xf9, 0x16, 0x7c, 0xaa, 0xb5, 0x12, 0x37, 0xa5,
	0xc6, 0x95, 0x65, 0xae, 0xe3, 0xab, 0x11, 0x3d, 0x5a, 0xf1, 0xad, 0xd5, 0x17, 0x72, 0xf2, 0x39,
	0xec, 0x37, 0x91, 0x75, 0xb7, 0x76, 0x6d, 0xc8, 0x68, 0x45, 0xb8, 0x46, 0x7d, 0x0c, 0x0d, 0x14,
	0x4b, 0xc5, 0x51, 0xd9, 0x46, 0xd9, 0x89, 0x76, 0x57, 0xf0, 0xaf, 0x06, 0x0d, 0xc6, 0x40, 0xfe,
	0xed, 0x69, 0x70, 0x0e, 0xe3, 0xb7, 0x39, 0x46, 0x9e, 0xc0, 0x1e, 0x65, 0x66, 0xc9, 0xd1, 0x1b,
	0x91, 0x08, 0xbd, 0x6c, 0x8a, 0x1e, 0xdd, 0xc1, 0x43, 0x1e, 0x3c, 0x87, 0xc3, 0xb7, 0xda, 0x66,
	0x76, 0x67, 0x4a, 0xf3, 0x58, 0x53, 0xb5, 0x40, 0xed, 0xa6, 0x62, 0x98, 0xd2, 0xfc, 0xda, 0x02,
	0xc1, 0x9b, 0x0e, 0x1c, 0x99, 0xfe, 0x4a, 0xf0, 0x76, 0x3d, 0xd2, 0x34, 0x37, 0xad, 0x47, 0xbe,
	0xe3, 0x9a, 0x9b, 0xba, 0x49, 0xff, 0x08, 0xb6, 0x0d, 0x99, 0x2b, 0x21, 0x95, 0xd0, 0x4b, 0x67,
	0x8c, 0x97, 0xd2, 0xfc, 0xca, 0x41, 0xe4, 0x43, 0x30, 0xf2, 0x58, 0x95, 0x49, 0x35, 0x36, 0xc3,
	0xa8, 0x9f, 0xd2, 0x3c, 0x2a, 0x13, 0xac, 0x8b, 0xa2, 0xfc, 0xb5, 0x60, 0x68, 0x07, 0xa7, 0x2a,
	0xea, 0xdc, 0x02, 0x6b, 0x35, 0x6f, 0xad, 0xd5, 0x4c, 0x3e, 0x86, 0x1d, 0x26, 0x95, 0x32, 0xf3,
	0xd9, 0x5e, 0x5d, 0xdb, 0x0d, 0x18, 0x72, 0xf2, 0x19, 0x8c, 0x4c, 0x0e, 0x46, 0x35, 0x2e, 0xa4,
	0x5a, 0x36, 0x8b, 0x7f, 0x27, 0xa5, 0xf9, 0xa5, 0x43, 0x43, 0x1e, 0xfc, 0xdd, 0x85, 0x83, 0x1f,
	0x6e, 0x35, 0x66, 0x1c, 0x79, 0xfb, 0xee, 0xfb, 0x06, 0xbc, 0x56, 0xbb, 0xd8, 0xef, 0xf7, 0xa6,
	0x07, 0x75, 0x8b, 0xb6, 0x94, 0x51, 0x5b, 0x47, 0x9e, 0x40, 0xdf, 0x5d, 0x4c, 0xf6, 0x9b, 0xbd,
	0xe9, 0xa8, 0x0e, 0x71, 0x7f, 0x0f, 0xa2, 0x9a, 0x27, 0x2f, 0xe1, 0x30, 0xb7, 0x43, 0xae, 0x90,
	0xb7, 0x5b, 0xd3, 0xfa, 0xf1, 0x1f, 0xef, 0x1a, 0xaf, 0x22, 0xda, 0xb5, 0x3e, 0x85, 0xb1, 0x42,
	0x56, 0xaa, 0xc2, 0x6c, 0xb8, 0x9c, 0xaa, 0x6a, 0x91, 0x14, 0xfe, 0xd6, 0xc9, 0x86, 0xb9, 0xe9,
	0x57, 0xdc, 0x95, 0xa5, 0x42, 0x5e, 0x98, 0xe6, 0xe6, 0xc2, 0x98, 0xd5, 0x96, 0xf7, 0xac, 0x7c,
	0x54, 0x11, 0x8d, 0xf6, 0x31, 0x8c, 0xea, 0xdb, 0xb7, 0xda, 0x1c, 0x85, 0xdf, 0xb7, 0xca, 0x5d,
	0x07, 0x47, 0x15, 0x4a, 0xce, 0xe0, 0xa0, 0x3d, 0x61, 0xb5, 0x78, 0x50, 0x55, 0xd1, 0xa2, 0x5c,
	0xc0, 0x4f, 0x9b, 0x83, 0xee, 0xde, 0xc6, 0xc5, 0x57, 0x70, 0xcc, 0x64, 0x3a, 0xc1, 0x84, 0x2b,
	0x71, 0x3b, 0x31, 0x17, 0xb0, 0xc8, 0x64, 0x22, 0x17, 0x4b, 0x67, 0x00, 0xd3, 0x17, 0xbd, 0x2b,
	0xb3, 0xa2, 0x8b, 0xdf, 0xdc, 0xff, 0xb1, 0x9b, 0x9e, 0x5d, 0xd9, 0x5f, 0xff, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xab, 0x62, 0x17, 0x0a, 0xae, 0x09, 0x00, 0x00,
}
