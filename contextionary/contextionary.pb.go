// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contextionary.proto

package contextionary

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SearchType int32

const (
	SearchType_CLASS    SearchType = 0
	SearchType_PROPERTY SearchType = 1
)

var SearchType_name = map[int32]string{
	0: "CLASS",
	1: "PROPERTY",
}

var SearchType_value = map[string]int32{
	"CLASS":    0,
	"PROPERTY": 1,
}

func (x SearchType) String() string {
	return proto.EnumName(SearchType_name, int32(x))
}

func (SearchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{0}
}

type Kind int32

const (
	Kind_THING  Kind = 0
	Kind_ACTION Kind = 1
)

var Kind_name = map[int32]string{
	0: "THING",
	1: "ACTION",
}

var Kind_value = map[string]int32{
	"THING":  0,
	"ACTION": 1,
}

func (x Kind) String() string {
	return proto.EnumName(Kind_name, int32(x))
}

func (Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{1}
}

type Word struct {
	Word                 string   `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Word) Reset()         { *m = Word{} }
func (m *Word) String() string { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()    {}
func (*Word) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{0}
}

func (m *Word) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Word.Unmarshal(m, b)
}
func (m *Word) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Word.Marshal(b, m, deterministic)
}
func (m *Word) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Word.Merge(m, src)
}
func (m *Word) XXX_Size() int {
	return xxx_messageInfo_Word.Size(m)
}
func (m *Word) XXX_DiscardUnknown() {
	xxx_messageInfo_Word.DiscardUnknown(m)
}

var xxx_messageInfo_Word proto.InternalMessageInfo

func (m *Word) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

type WordPresent struct {
	Present              bool     `protobuf:"varint,1,opt,name=present,proto3" json:"present,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WordPresent) Reset()         { *m = WordPresent{} }
func (m *WordPresent) String() string { return proto.CompactTextString(m) }
func (*WordPresent) ProtoMessage()    {}
func (*WordPresent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{1}
}

func (m *WordPresent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WordPresent.Unmarshal(m, b)
}
func (m *WordPresent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WordPresent.Marshal(b, m, deterministic)
}
func (m *WordPresent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WordPresent.Merge(m, src)
}
func (m *WordPresent) XXX_Size() int {
	return xxx_messageInfo_WordPresent.Size(m)
}
func (m *WordPresent) XXX_DiscardUnknown() {
	xxx_messageInfo_WordPresent.DiscardUnknown(m)
}

var xxx_messageInfo_WordPresent proto.InternalMessageInfo

func (m *WordPresent) GetPresent() bool {
	if m != nil {
		return m.Present
	}
	return false
}

type Vector struct {
	Entries              []*VectorEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Vector) Reset()         { *m = Vector{} }
func (m *Vector) String() string { return proto.CompactTextString(m) }
func (*Vector) ProtoMessage()    {}
func (*Vector) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{2}
}

func (m *Vector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector.Unmarshal(m, b)
}
func (m *Vector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector.Marshal(b, m, deterministic)
}
func (m *Vector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector.Merge(m, src)
}
func (m *Vector) XXX_Size() int {
	return xxx_messageInfo_Vector.Size(m)
}
func (m *Vector) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector.DiscardUnknown(m)
}

var xxx_messageInfo_Vector proto.InternalMessageInfo

func (m *Vector) GetEntries() []*VectorEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type VectorEntry struct {
	Entry                float32  `protobuf:"fixed32,1,opt,name=Entry,proto3" json:"Entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VectorEntry) Reset()         { *m = VectorEntry{} }
func (m *VectorEntry) String() string { return proto.CompactTextString(m) }
func (*VectorEntry) ProtoMessage()    {}
func (*VectorEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{3}
}

func (m *VectorEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VectorEntry.Unmarshal(m, b)
}
func (m *VectorEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VectorEntry.Marshal(b, m, deterministic)
}
func (m *VectorEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VectorEntry.Merge(m, src)
}
func (m *VectorEntry) XXX_Size() int {
	return xxx_messageInfo_VectorEntry.Size(m)
}
func (m *VectorEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_VectorEntry.DiscardUnknown(m)
}

var xxx_messageInfo_VectorEntry proto.InternalMessageInfo

func (m *VectorEntry) GetEntry() float32 {
	if m != nil {
		return m.Entry
	}
	return 0
}

type VectorNNParams struct {
	Vector               *Vector  `protobuf:"bytes,1,opt,name=vector,proto3" json:"vector,omitempty"`
	K                    int32    `protobuf:"varint,2,opt,name=k,proto3" json:"k,omitempty"`
	N                    int32    `protobuf:"varint,3,opt,name=n,proto3" json:"n,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VectorNNParams) Reset()         { *m = VectorNNParams{} }
func (m *VectorNNParams) String() string { return proto.CompactTextString(m) }
func (*VectorNNParams) ProtoMessage()    {}
func (*VectorNNParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{4}
}

func (m *VectorNNParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VectorNNParams.Unmarshal(m, b)
}
func (m *VectorNNParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VectorNNParams.Marshal(b, m, deterministic)
}
func (m *VectorNNParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VectorNNParams.Merge(m, src)
}
func (m *VectorNNParams) XXX_Size() int {
	return xxx_messageInfo_VectorNNParams.Size(m)
}
func (m *VectorNNParams) XXX_DiscardUnknown() {
	xxx_messageInfo_VectorNNParams.DiscardUnknown(m)
}

var xxx_messageInfo_VectorNNParams proto.InternalMessageInfo

func (m *VectorNNParams) GetVector() *Vector {
	if m != nil {
		return m.Vector
	}
	return nil
}

func (m *VectorNNParams) GetK() int32 {
	if m != nil {
		return m.K
	}
	return 0
}

func (m *VectorNNParams) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

type Corpi struct {
	Corpi                []string `protobuf:"bytes,1,rep,name=corpi,proto3" json:"corpi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Corpi) Reset()         { *m = Corpi{} }
func (m *Corpi) String() string { return proto.CompactTextString(m) }
func (*Corpi) ProtoMessage()    {}
func (*Corpi) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{5}
}

func (m *Corpi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Corpi.Unmarshal(m, b)
}
func (m *Corpi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Corpi.Marshal(b, m, deterministic)
}
func (m *Corpi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Corpi.Merge(m, src)
}
func (m *Corpi) XXX_Size() int {
	return xxx_messageInfo_Corpi.Size(m)
}
func (m *Corpi) XXX_DiscardUnknown() {
	xxx_messageInfo_Corpi.DiscardUnknown(m)
}

var xxx_messageInfo_Corpi proto.InternalMessageInfo

func (m *Corpi) GetCorpi() []string {
	if m != nil {
		return m.Corpi
	}
	return nil
}

type WordStopword struct {
	Stopword             bool     `protobuf:"varint,1,opt,name=stopword,proto3" json:"stopword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WordStopword) Reset()         { *m = WordStopword{} }
func (m *WordStopword) String() string { return proto.CompactTextString(m) }
func (*WordStopword) ProtoMessage()    {}
func (*WordStopword) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{6}
}

func (m *WordStopword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WordStopword.Unmarshal(m, b)
}
func (m *WordStopword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WordStopword.Marshal(b, m, deterministic)
}
func (m *WordStopword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WordStopword.Merge(m, src)
}
func (m *WordStopword) XXX_Size() int {
	return xxx_messageInfo_WordStopword.Size(m)
}
func (m *WordStopword) XXX_DiscardUnknown() {
	xxx_messageInfo_WordStopword.DiscardUnknown(m)
}

var xxx_messageInfo_WordStopword proto.InternalMessageInfo

func (m *WordStopword) GetStopword() bool {
	if m != nil {
		return m.Stopword
	}
	return false
}

type SimilarWordsParams struct {
	Word                 string   `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	Certainty            float32  `protobuf:"fixed32,2,opt,name=certainty,proto3" json:"certainty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimilarWordsParams) Reset()         { *m = SimilarWordsParams{} }
func (m *SimilarWordsParams) String() string { return proto.CompactTextString(m) }
func (*SimilarWordsParams) ProtoMessage()    {}
func (*SimilarWordsParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{7}
}

func (m *SimilarWordsParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimilarWordsParams.Unmarshal(m, b)
}
func (m *SimilarWordsParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimilarWordsParams.Marshal(b, m, deterministic)
}
func (m *SimilarWordsParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimilarWordsParams.Merge(m, src)
}
func (m *SimilarWordsParams) XXX_Size() int {
	return xxx_messageInfo_SimilarWordsParams.Size(m)
}
func (m *SimilarWordsParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SimilarWordsParams.DiscardUnknown(m)
}

var xxx_messageInfo_SimilarWordsParams proto.InternalMessageInfo

func (m *SimilarWordsParams) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *SimilarWordsParams) GetCertainty() float32 {
	if m != nil {
		return m.Certainty
	}
	return 0
}

type SimilarWordsResults struct {
	Words                []*Word  `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimilarWordsResults) Reset()         { *m = SimilarWordsResults{} }
func (m *SimilarWordsResults) String() string { return proto.CompactTextString(m) }
func (*SimilarWordsResults) ProtoMessage()    {}
func (*SimilarWordsResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{8}
}

func (m *SimilarWordsResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimilarWordsResults.Unmarshal(m, b)
}
func (m *SimilarWordsResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimilarWordsResults.Marshal(b, m, deterministic)
}
func (m *SimilarWordsResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimilarWordsResults.Merge(m, src)
}
func (m *SimilarWordsResults) XXX_Size() int {
	return xxx_messageInfo_SimilarWordsResults.Size(m)
}
func (m *SimilarWordsResults) XXX_DiscardUnknown() {
	xxx_messageInfo_SimilarWordsResults.DiscardUnknown(m)
}

var xxx_messageInfo_SimilarWordsResults proto.InternalMessageInfo

func (m *SimilarWordsResults) GetWords() []*Word {
	if m != nil {
		return m.Words
	}
	return nil
}

type NearestWords struct {
	Words                []string  `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
	Distances            []float32 `protobuf:"fixed32,2,rep,packed,name=distances,proto3" json:"distances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *NearestWords) Reset()         { *m = NearestWords{} }
func (m *NearestWords) String() string { return proto.CompactTextString(m) }
func (*NearestWords) ProtoMessage()    {}
func (*NearestWords) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{9}
}

func (m *NearestWords) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NearestWords.Unmarshal(m, b)
}
func (m *NearestWords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NearestWords.Marshal(b, m, deterministic)
}
func (m *NearestWords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NearestWords.Merge(m, src)
}
func (m *NearestWords) XXX_Size() int {
	return xxx_messageInfo_NearestWords.Size(m)
}
func (m *NearestWords) XXX_DiscardUnknown() {
	xxx_messageInfo_NearestWords.DiscardUnknown(m)
}

var xxx_messageInfo_NearestWords proto.InternalMessageInfo

func (m *NearestWords) GetWords() []string {
	if m != nil {
		return m.Words
	}
	return nil
}

func (m *NearestWords) GetDistances() []float32 {
	if m != nil {
		return m.Distances
	}
	return nil
}

type Keyword struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Weight               float32  `protobuf:"fixed32,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Keyword) Reset()         { *m = Keyword{} }
func (m *Keyword) String() string { return proto.CompactTextString(m) }
func (*Keyword) ProtoMessage()    {}
func (*Keyword) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{10}
}

func (m *Keyword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyword.Unmarshal(m, b)
}
func (m *Keyword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyword.Marshal(b, m, deterministic)
}
func (m *Keyword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyword.Merge(m, src)
}
func (m *Keyword) XXX_Size() int {
	return xxx_messageInfo_Keyword.Size(m)
}
func (m *Keyword) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyword.DiscardUnknown(m)
}

var xxx_messageInfo_Keyword proto.InternalMessageInfo

func (m *Keyword) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *Keyword) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type SchemaSearchParams struct {
	SearchType           SearchType `protobuf:"varint,1,opt,name=searchType,proto3,enum=contextionary.SearchType" json:"searchType,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Keywords             []*Keyword `protobuf:"bytes,3,rep,name=keywords,proto3" json:"keywords,omitempty"`
	Kind                 Kind       `protobuf:"varint,4,opt,name=kind,proto3,enum=contextionary.Kind" json:"kind,omitempty"`
	Certainty            float32    `protobuf:"fixed32,5,opt,name=certainty,proto3" json:"certainty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SchemaSearchParams) Reset()         { *m = SchemaSearchParams{} }
func (m *SchemaSearchParams) String() string { return proto.CompactTextString(m) }
func (*SchemaSearchParams) ProtoMessage()    {}
func (*SchemaSearchParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{11}
}

func (m *SchemaSearchParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchemaSearchParams.Unmarshal(m, b)
}
func (m *SchemaSearchParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchemaSearchParams.Marshal(b, m, deterministic)
}
func (m *SchemaSearchParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaSearchParams.Merge(m, src)
}
func (m *SchemaSearchParams) XXX_Size() int {
	return xxx_messageInfo_SchemaSearchParams.Size(m)
}
func (m *SchemaSearchParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaSearchParams.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaSearchParams proto.InternalMessageInfo

func (m *SchemaSearchParams) GetSearchType() SearchType {
	if m != nil {
		return m.SearchType
	}
	return SearchType_CLASS
}

func (m *SchemaSearchParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SchemaSearchParams) GetKeywords() []*Keyword {
	if m != nil {
		return m.Keywords
	}
	return nil
}

func (m *SchemaSearchParams) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_THING
}

func (m *SchemaSearchParams) GetCertainty() float32 {
	if m != nil {
		return m.Certainty
	}
	return 0
}

type SchemaSearchResults struct {
	Type                 SearchType            `protobuf:"varint,1,opt,name=type,proto3,enum=contextionary.SearchType" json:"type,omitempty"`
	Results              []*SchemaSearchResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SchemaSearchResults) Reset()         { *m = SchemaSearchResults{} }
func (m *SchemaSearchResults) String() string { return proto.CompactTextString(m) }
func (*SchemaSearchResults) ProtoMessage()    {}
func (*SchemaSearchResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{12}
}

func (m *SchemaSearchResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchemaSearchResults.Unmarshal(m, b)
}
func (m *SchemaSearchResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchemaSearchResults.Marshal(b, m, deterministic)
}
func (m *SchemaSearchResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaSearchResults.Merge(m, src)
}
func (m *SchemaSearchResults) XXX_Size() int {
	return xxx_messageInfo_SchemaSearchResults.Size(m)
}
func (m *SchemaSearchResults) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaSearchResults.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaSearchResults proto.InternalMessageInfo

func (m *SchemaSearchResults) GetType() SearchType {
	if m != nil {
		return m.Type
	}
	return SearchType_CLASS
}

func (m *SchemaSearchResults) GetResults() []*SchemaSearchResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type SchemaSearchResult struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind                 Kind     `protobuf:"varint,2,opt,name=kind,proto3,enum=contextionary.Kind" json:"kind,omitempty"`
	Certainty            float32  `protobuf:"fixed32,3,opt,name=certainty,proto3" json:"certainty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchemaSearchResult) Reset()         { *m = SchemaSearchResult{} }
func (m *SchemaSearchResult) String() string { return proto.CompactTextString(m) }
func (*SchemaSearchResult) ProtoMessage()    {}
func (*SchemaSearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6af9fd695f521f0, []int{13}
}

func (m *SchemaSearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchemaSearchResult.Unmarshal(m, b)
}
func (m *SchemaSearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchemaSearchResult.Marshal(b, m, deterministic)
}
func (m *SchemaSearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaSearchResult.Merge(m, src)
}
func (m *SchemaSearchResult) XXX_Size() int {
	return xxx_messageInfo_SchemaSearchResult.Size(m)
}
func (m *SchemaSearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaSearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaSearchResult proto.InternalMessageInfo

func (m *SchemaSearchResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SchemaSearchResult) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_THING
}

func (m *SchemaSearchResult) GetCertainty() float32 {
	if m != nil {
		return m.Certainty
	}
	return 0
}

func init() {
	proto.RegisterEnum("contextionary.SearchType", SearchType_name, SearchType_value)
	proto.RegisterEnum("contextionary.Kind", Kind_name, Kind_value)
	proto.RegisterType((*Word)(nil), "contextionary.Word")
	proto.RegisterType((*WordPresent)(nil), "contextionary.WordPresent")
	proto.RegisterType((*Vector)(nil), "contextionary.Vector")
	proto.RegisterType((*VectorEntry)(nil), "contextionary.VectorEntry")
	proto.RegisterType((*VectorNNParams)(nil), "contextionary.VectorNNParams")
	proto.RegisterType((*Corpi)(nil), "contextionary.Corpi")
	proto.RegisterType((*WordStopword)(nil), "contextionary.WordStopword")
	proto.RegisterType((*SimilarWordsParams)(nil), "contextionary.SimilarWordsParams")
	proto.RegisterType((*SimilarWordsResults)(nil), "contextionary.SimilarWordsResults")
	proto.RegisterType((*NearestWords)(nil), "contextionary.NearestWords")
	proto.RegisterType((*Keyword)(nil), "contextionary.Keyword")
	proto.RegisterType((*SchemaSearchParams)(nil), "contextionary.SchemaSearchParams")
	proto.RegisterType((*SchemaSearchResults)(nil), "contextionary.SchemaSearchResults")
	proto.RegisterType((*SchemaSearchResult)(nil), "contextionary.SchemaSearchResult")
}

func init() { proto.RegisterFile("contextionary.proto", fileDescriptor_e6af9fd695f521f0) }

var fileDescriptor_e6af9fd695f521f0 = []byte{
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xb5, 0x73, 0xcf, 0x24, 0xa9, 0xa2, 0x49, 0xa9, 0x4c, 0xa0, 0x52, 0x58, 0x84, 0x5a, 0x2a,
	0xb5, 0x0f, 0x81, 0x17, 0x54, 0x71, 0x69, 0x43, 0x5b, 0xaa, 0xa2, 0x34, 0x5a, 0x47, 0x54, 0x88,
	0x27, 0x93, 0x2c, 0xc4, 0x4a, 0x63, 0x47, 0xbb, 0x0b, 0x25, 0x8f, 0x7c, 0x14, 0xbf, 0xc3, 0xb7,
	0xa0, 0x5d, 0xaf, 0x53, 0xc7, 0x31, 0xb7, 0xb7, 0x99, 0xf1, 0x99, 0xcb, 0x39, 0x63, 0x8f, 0xa1,
	0x35, 0x0a, 0x03, 0xc9, 0xbe, 0x49, 0x3f, 0x0c, 0x3c, 0xbe, 0x38, 0x98, 0xf3, 0x50, 0x86, 0xd8,
	0x58, 0x09, 0x92, 0x36, 0x14, 0xae, 0x42, 0x3e, 0x46, 0x84, 0xc2, 0x4d, 0xc8, 0xc7, 0x8e, 0xdd,
	0xb1, 0x77, 0xab, 0x54, 0xdb, 0x64, 0x07, 0x6a, 0xea, 0xd9, 0x80, 0x33, 0xc1, 0x02, 0x89, 0x0e,
	0x94, 0xe7, 0x91, 0xa9, 0x51, 0x15, 0x1a, 0xbb, 0xe4, 0x05, 0x94, 0xde, 0xb1, 0x91, 0x0c, 0x39,
	0x3e, 0x85, 0x32, 0x0b, 0x24, 0xf7, 0x99, 0x70, 0xec, 0x4e, 0x7e, 0xb7, 0xd6, 0x6d, 0x1f, 0xac,
	0x0e, 0x11, 0xe1, 0x4e, 0x02, 0xc9, 0x17, 0x34, 0x86, 0x92, 0x87, 0x50, 0x4b, 0xc4, 0x71, 0x13,
	0x8a, 0xda, 0xd0, 0x6d, 0x72, 0x34, 0x72, 0xc8, 0x07, 0xd8, 0x88, 0x40, 0xfd, 0xfe, 0xc0, 0xe3,
	0xde, 0x4c, 0xe0, 0x3e, 0x94, 0xbe, 0xea, 0x88, 0x06, 0xd6, 0xba, 0x77, 0x32, 0x7b, 0x51, 0x03,
	0xc2, 0x3a, 0xd8, 0x53, 0x27, 0xd7, 0xb1, 0x77, 0x8b, 0xd4, 0x9e, 0x2a, 0x2f, 0x70, 0xf2, 0x91,
	0x17, 0x90, 0x6d, 0x28, 0xf6, 0x42, 0x3e, 0xf7, 0x55, 0xef, 0x91, 0x32, 0xf4, 0xf8, 0x55, 0x1a,
	0x39, 0x64, 0x0f, 0xea, 0x4a, 0x09, 0x57, 0x86, 0x73, 0xa5, 0x0c, 0xb6, 0xa1, 0x22, 0x8c, 0x6d,
	0xb4, 0x58, 0xfa, 0xe4, 0x14, 0xd0, 0xf5, 0x67, 0xfe, 0xb5, 0xc7, 0x55, 0x8a, 0x30, 0xb3, 0x66,
	0xe8, 0x8b, 0xf7, 0xa1, 0x3a, 0x62, 0x5c, 0x7a, 0x7e, 0x20, 0x17, 0x7a, 0xb0, 0x1c, 0xbd, 0x0d,
	0x90, 0x57, 0xd0, 0x4a, 0xd6, 0xa1, 0x4c, 0x7c, 0xb9, 0x96, 0x02, 0x1f, 0x43, 0x51, 0x25, 0xc7,
	0xfa, 0xb6, 0x52, 0x9c, 0x15, 0x96, 0x46, 0x08, 0x72, 0x0c, 0xf5, 0x3e, 0xf3, 0x38, 0x13, 0x52,
	0x57, 0x50, 0xdc, 0x6e, 0x53, 0xab, 0x06, 0xa5, 0xa6, 0x18, 0xfb, 0x42, 0x7a, 0xc1, 0x88, 0x09,
	0x27, 0xd7, 0xc9, 0xab, 0x29, 0x96, 0x01, 0x72, 0x08, 0xe5, 0x0b, 0xb6, 0xd0, 0xe3, 0x3a, 0x50,
	0x9e, 0x46, 0xa6, 0x61, 0x11, 0xbb, 0xb8, 0x05, 0xa5, 0x1b, 0xe6, 0x7f, 0x9e, 0x48, 0xc3, 0xc2,
	0x78, 0xe4, 0xa7, 0x0d, 0xe8, 0x8e, 0x26, 0x6c, 0xe6, 0xb9, 0xcc, 0xe3, 0xa3, 0x89, 0xd1, 0xe2,
	0x19, 0x80, 0xd0, 0xfe, 0x70, 0x31, 0x67, 0xba, 0xd6, 0x46, 0xf7, 0x6e, 0x8a, 0x87, 0xbb, 0x04,
	0xd0, 0x04, 0x58, 0xc9, 0x18, 0x78, 0x33, 0xa6, 0xfb, 0x54, 0xa9, 0xb6, 0xb1, 0x0b, 0x15, 0x33,
	0x88, 0x70, 0xf2, 0x5a, 0x94, 0xad, 0x54, 0x31, 0xc3, 0x80, 0x2e, 0x71, 0xb8, 0x03, 0x85, 0xa9,
	0x1f, 0x8c, 0x9d, 0x82, 0x6e, 0x9e, 0x16, 0xf1, 0xc2, 0x0f, 0xc6, 0x54, 0x03, 0x56, 0x77, 0x54,
	0x4c, 0xef, 0xe8, 0xbb, 0x0d, 0xad, 0x24, 0xc1, 0x78, 0x49, 0xfb, 0x50, 0x90, 0xff, 0xc4, 0x4d,
	0xc3, 0xf0, 0x10, 0xca, 0x3c, 0xca, 0xd4, 0x0b, 0xa8, 0x75, 0x1f, 0xa4, 0x33, 0xd6, 0x7a, 0xd0,
	0x38, 0x83, 0x84, 0xab, 0x1a, 0x47, 0x8f, 0x97, 0x42, 0xd9, 0x09, 0xa1, 0x62, 0xd2, 0xb9, 0xff,
	0x22, 0x9d, 0x4f, 0x91, 0xde, 0x7b, 0x04, 0x70, 0xcb, 0x00, 0xab, 0x50, 0xec, 0xbd, 0x3d, 0x72,
	0xdd, 0xa6, 0x85, 0x75, 0xa8, 0x0c, 0xe8, 0xe5, 0xe0, 0x84, 0x0e, 0xdf, 0x37, 0xed, 0xbd, 0x6d,
	0x28, 0xa8, 0x92, 0x0a, 0x30, 0x7c, 0x73, 0xde, 0x3f, 0x6b, 0x5a, 0x08, 0x50, 0x3a, 0xea, 0x0d,
	0xcf, 0x2f, 0xfb, 0x4d, 0xbb, 0xfb, 0xa3, 0x00, 0x8d, 0x5e, 0x72, 0x00, 0x7c, 0x0d, 0x1b, 0xe7,
	0x62, 0xe5, 0x33, 0xcb, 0x7a, 0xb9, 0xdb, 0xf7, 0x32, 0x82, 0x71, 0x06, 0xb1, 0xf0, 0x18, 0x1a,
	0x51, 0x95, 0xf8, 0x6c, 0x65, 0x16, 0x69, 0x67, 0x04, 0x4d, 0x02, 0xb1, 0xf0, 0x0a, 0xea, 0x49,
	0x49, 0xf1, 0x4f, 0xeb, 0x88, 0xde, 0xe9, 0x36, 0xf9, 0xeb, 0xc6, 0x04, 0xb1, 0x70, 0x0a, 0x1d,
	0xd7, 0xfb, 0xc4, 0xce, 0x98, 0x4c, 0x7e, 0xda, 0x57, 0xbe, 0x9c, 0xf4, 0x62, 0x79, 0xd7, 0x9b,
	0xad, 0x1d, 0x93, 0xf5, 0x66, 0xeb, 0x77, 0x82, 0x58, 0xf8, 0x1c, 0x1a, 0xd1, 0x05, 0x3c, 0x0d,
	0xf5, 0xa3, 0x6c, 0x25, 0xb2, 0x8f, 0x26, 0xb1, 0xf0, 0x65, 0x7c, 0x6f, 0x4f, 0x43, 0x6e, 0x6e,
	0x63, 0x0a, 0xaa, 0xa3, 0xbf, 0x2f, 0x30, 0x84, 0xcd, 0xe4, 0xf9, 0x39, 0x5e, 0x98, 0x7f, 0xc4,
	0x76, 0x66, 0x42, 0x7c, 0xd5, 0xd7, 0xf6, 0x9b, 0xac, 0x41, 0xac, 0x8f, 0x25, 0xfd, 0x1b, 0x7b,
	0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x92, 0x39, 0xcf, 0x27, 0xdd, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContextionaryClient is the client API for Contextionary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContextionaryClient interface {
	IsWordStopword(ctx context.Context, in *Word, opts ...grpc.CallOption) (*WordStopword, error)
	IsWordPresent(ctx context.Context, in *Word, opts ...grpc.CallOption) (*WordPresent, error)
	SchemaSearch(ctx context.Context, in *SchemaSearchParams, opts ...grpc.CallOption) (*SchemaSearchResults, error)
	SafeGetSimilarWordsWithCertainty(ctx context.Context, in *SimilarWordsParams, opts ...grpc.CallOption) (*SimilarWordsResults, error)
	VectorForWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Vector, error)
	VectorForCorpi(ctx context.Context, in *Corpi, opts ...grpc.CallOption) (*Vector, error)
	NearestWordsByVector(ctx context.Context, in *VectorNNParams, opts ...grpc.CallOption) (*NearestWords, error)
}

type contextionaryClient struct {
	cc *grpc.ClientConn
}

func NewContextionaryClient(cc *grpc.ClientConn) ContextionaryClient {
	return &contextionaryClient{cc}
}

func (c *contextionaryClient) IsWordStopword(ctx context.Context, in *Word, opts ...grpc.CallOption) (*WordStopword, error) {
	out := new(WordStopword)
	err := c.cc.Invoke(ctx, "/contextionary.Contextionary/IsWordStopword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextionaryClient) IsWordPresent(ctx context.Context, in *Word, opts ...grpc.CallOption) (*WordPresent, error) {
	out := new(WordPresent)
	err := c.cc.Invoke(ctx, "/contextionary.Contextionary/IsWordPresent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextionaryClient) SchemaSearch(ctx context.Context, in *SchemaSearchParams, opts ...grpc.CallOption) (*SchemaSearchResults, error) {
	out := new(SchemaSearchResults)
	err := c.cc.Invoke(ctx, "/contextionary.Contextionary/SchemaSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextionaryClient) SafeGetSimilarWordsWithCertainty(ctx context.Context, in *SimilarWordsParams, opts ...grpc.CallOption) (*SimilarWordsResults, error) {
	out := new(SimilarWordsResults)
	err := c.cc.Invoke(ctx, "/contextionary.Contextionary/SafeGetSimilarWordsWithCertainty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextionaryClient) VectorForWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Vector, error) {
	out := new(Vector)
	err := c.cc.Invoke(ctx, "/contextionary.Contextionary/VectorForWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextionaryClient) VectorForCorpi(ctx context.Context, in *Corpi, opts ...grpc.CallOption) (*Vector, error) {
	out := new(Vector)
	err := c.cc.Invoke(ctx, "/contextionary.Contextionary/VectorForCorpi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextionaryClient) NearestWordsByVector(ctx context.Context, in *VectorNNParams, opts ...grpc.CallOption) (*NearestWords, error) {
	out := new(NearestWords)
	err := c.cc.Invoke(ctx, "/contextionary.Contextionary/NearestWordsByVector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContextionaryServer is the server API for Contextionary service.
type ContextionaryServer interface {
	IsWordStopword(context.Context, *Word) (*WordStopword, error)
	IsWordPresent(context.Context, *Word) (*WordPresent, error)
	SchemaSearch(context.Context, *SchemaSearchParams) (*SchemaSearchResults, error)
	SafeGetSimilarWordsWithCertainty(context.Context, *SimilarWordsParams) (*SimilarWordsResults, error)
	VectorForWord(context.Context, *Word) (*Vector, error)
	VectorForCorpi(context.Context, *Corpi) (*Vector, error)
	NearestWordsByVector(context.Context, *VectorNNParams) (*NearestWords, error)
}

func RegisterContextionaryServer(s *grpc.Server, srv ContextionaryServer) {
	s.RegisterService(&_Contextionary_serviceDesc, srv)
}

func _Contextionary_IsWordStopword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextionaryServer).IsWordStopword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contextionary.Contextionary/IsWordStopword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextionaryServer).IsWordStopword(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contextionary_IsWordPresent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextionaryServer).IsWordPresent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contextionary.Contextionary/IsWordPresent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextionaryServer).IsWordPresent(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contextionary_SchemaSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaSearchParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextionaryServer).SchemaSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contextionary.Contextionary/SchemaSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextionaryServer).SchemaSearch(ctx, req.(*SchemaSearchParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contextionary_SafeGetSimilarWordsWithCertainty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimilarWordsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextionaryServer).SafeGetSimilarWordsWithCertainty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contextionary.Contextionary/SafeGetSimilarWordsWithCertainty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextionaryServer).SafeGetSimilarWordsWithCertainty(ctx, req.(*SimilarWordsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contextionary_VectorForWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextionaryServer).VectorForWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contextionary.Contextionary/VectorForWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextionaryServer).VectorForWord(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contextionary_VectorForCorpi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Corpi)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextionaryServer).VectorForCorpi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contextionary.Contextionary/VectorForCorpi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextionaryServer).VectorForCorpi(ctx, req.(*Corpi))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contextionary_NearestWordsByVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VectorNNParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextionaryServer).NearestWordsByVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contextionary.Contextionary/NearestWordsByVector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextionaryServer).NearestWordsByVector(ctx, req.(*VectorNNParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contextionary_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contextionary.Contextionary",
	HandlerType: (*ContextionaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsWordStopword",
			Handler:    _Contextionary_IsWordStopword_Handler,
		},
		{
			MethodName: "IsWordPresent",
			Handler:    _Contextionary_IsWordPresent_Handler,
		},
		{
			MethodName: "SchemaSearch",
			Handler:    _Contextionary_SchemaSearch_Handler,
		},
		{
			MethodName: "SafeGetSimilarWordsWithCertainty",
			Handler:    _Contextionary_SafeGetSimilarWordsWithCertainty_Handler,
		},
		{
			MethodName: "VectorForWord",
			Handler:    _Contextionary_VectorForWord_Handler,
		},
		{
			MethodName: "VectorForCorpi",
			Handler:    _Contextionary_VectorForCorpi_Handler,
		},
		{
			MethodName: "NearestWordsByVector",
			Handler:    _Contextionary_NearestWordsByVector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contextionary.proto",
}
