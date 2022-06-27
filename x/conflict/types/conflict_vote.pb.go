// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: conflict/conflict_vote.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Provider struct {
	Account  string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Response []byte `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5ff6a0d8edaa7f1, []int{0}
}
func (m *Provider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return m.Size()
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Provider) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

type ConflictVote struct {
	Index          string            `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	VoteDeadline   int64             `protobuf:"varint,3,opt,name=voteDeadline,proto3" json:"voteDeadline,omitempty"`
	VoteStartBlock int64             `protobuf:"varint,4,opt,name=voteStartBlock,proto3" json:"voteStartBlock,omitempty"`
	VoteIsCommit   bool              `protobuf:"varint,5,opt,name=voteIsCommit,proto3" json:"voteIsCommit,omitempty"`
	ChainID        string            `protobuf:"bytes,6,opt,name=chainID,proto3" json:"chainID,omitempty"`
	ApiUrl         string            `protobuf:"bytes,7,opt,name=apiUrl,proto3" json:"apiUrl,omitempty"`
	RequestData    []byte            `protobuf:"bytes,8,opt,name=requestData,proto3" json:"requestData,omitempty"`
	RequestBlock   int64             `protobuf:"varint,9,opt,name=requestBlock,proto3" json:"requestBlock,omitempty"`
	FirstProvider  Provider          `protobuf:"bytes,10,opt,name=firstProvider,proto3" json:"firstProvider"`
	SecondProvider Provider          `protobuf:"bytes,11,opt,name=secondProvider,proto3" json:"secondProvider"`
	Voters         []string          `protobuf:"bytes,12,rep,name=voters,proto3" json:"voters,omitempty"`
	VotersHash     map[string][]byte `protobuf:"bytes,13,rep,name=votersHash,proto3" json:"votersHash" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ConflictVote) Reset()         { *m = ConflictVote{} }
func (m *ConflictVote) String() string { return proto.CompactTextString(m) }
func (*ConflictVote) ProtoMessage()    {}
func (*ConflictVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5ff6a0d8edaa7f1, []int{1}
}
func (m *ConflictVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConflictVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConflictVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConflictVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConflictVote.Merge(m, src)
}
func (m *ConflictVote) XXX_Size() int {
	return m.Size()
}
func (m *ConflictVote) XXX_DiscardUnknown() {
	xxx_messageInfo_ConflictVote.DiscardUnknown(m)
}

var xxx_messageInfo_ConflictVote proto.InternalMessageInfo

func (m *ConflictVote) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ConflictVote) GetVoteDeadline() int64 {
	if m != nil {
		return m.VoteDeadline
	}
	return 0
}

func (m *ConflictVote) GetVoteStartBlock() int64 {
	if m != nil {
		return m.VoteStartBlock
	}
	return 0
}

func (m *ConflictVote) GetVoteIsCommit() bool {
	if m != nil {
		return m.VoteIsCommit
	}
	return false
}

func (m *ConflictVote) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *ConflictVote) GetApiUrl() string {
	if m != nil {
		return m.ApiUrl
	}
	return ""
}

func (m *ConflictVote) GetRequestData() []byte {
	if m != nil {
		return m.RequestData
	}
	return nil
}

func (m *ConflictVote) GetRequestBlock() int64 {
	if m != nil {
		return m.RequestBlock
	}
	return 0
}

func (m *ConflictVote) GetFirstProvider() Provider {
	if m != nil {
		return m.FirstProvider
	}
	return Provider{}
}

func (m *ConflictVote) GetSecondProvider() Provider {
	if m != nil {
		return m.SecondProvider
	}
	return Provider{}
}

func (m *ConflictVote) GetVoters() []string {
	if m != nil {
		return m.Voters
	}
	return nil
}

func (m *ConflictVote) GetVotersHash() map[string][]byte {
	if m != nil {
		return m.VotersHash
	}
	return nil
}

func init() {
	proto.RegisterType((*Provider)(nil), "lavanet.lava.conflict.Provider")
	proto.RegisterType((*ConflictVote)(nil), "lavanet.lava.conflict.ConflictVote")
	proto.RegisterMapType((map[string][]byte)(nil), "lavanet.lava.conflict.ConflictVote.VotersHashEntry")
}

func init() { proto.RegisterFile("conflict/conflict_vote.proto", fileDescriptor_c5ff6a0d8edaa7f1) }

var fileDescriptor_c5ff6a0d8edaa7f1 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6a, 0x1b, 0x31,
	0x10, 0xb5, 0xb2, 0x8e, 0x63, 0x8f, 0x9d, 0xb4, 0x88, 0xb4, 0x08, 0x13, 0x36, 0x8b, 0x0f, 0x65,
	0x4f, 0xbb, 0x90, 0x5c, 0x4a, 0xa1, 0x50, 0x1c, 0x17, 0x1a, 0x4a, 0xa1, 0x6c, 0x69, 0xa0, 0xbd,
	0x14, 0x45, 0x56, 0x6c, 0x91, 0xb5, 0xe4, 0x4a, 0xb2, 0x89, 0xff, 0xa2, 0x9f, 0x95, 0x63, 0x8e,
	0x3d, 0x95, 0x60, 0xff, 0x48, 0xd1, 0xae, 0xd6, 0xd8, 0xa6, 0x3d, 0xe4, 0xa4, 0x79, 0x6f, 0xdf,
	0xbc, 0x99, 0x91, 0x66, 0xe1, 0x84, 0x29, 0x79, 0x93, 0x0b, 0x66, 0xd3, 0x2a, 0xf8, 0x31, 0x57,
	0x96, 0x27, 0x53, 0xad, 0xac, 0xc2, 0x2f, 0x72, 0x3a, 0xa7, 0x92, 0xdb, 0xc4, 0x9d, 0x49, 0xa5,
	0xe8, 0x1e, 0x8f, 0xd4, 0x48, 0x15, 0x8a, 0xd4, 0x45, 0xa5, 0xb8, 0xf7, 0x0e, 0x9a, 0x9f, 0xb5,
	0x9a, 0x8b, 0x21, 0xd7, 0x98, 0xc0, 0x01, 0x65, 0x4c, 0xcd, 0xa4, 0x25, 0x28, 0x42, 0x71, 0x2b,
	0xab, 0x20, 0xee, 0x42, 0x53, 0x73, 0x33, 0x55, 0xd2, 0x70, 0xb2, 0x17, 0xa1, 0xb8, 0x93, 0xad,
	0x71, 0xef, 0xb1, 0x0e, 0x9d, 0x0b, 0x5f, 0xe4, 0x4a, 0x59, 0x8e, 0x8f, 0x61, 0x5f, 0xc8, 0x21,
	0xbf, 0xf3, 0x26, 0x25, 0xc0, 0x3d, 0xe8, 0xb8, 0x1e, 0x07, 0x9c, 0x0e, 0x73, 0x21, 0x39, 0x09,
	0x22, 0x14, 0x07, 0xd9, 0x16, 0x87, 0x5f, 0xc1, 0x91, 0xc3, 0x5f, 0x2c, 0xd5, 0xb6, 0x9f, 0x2b,
	0x76, 0x4b, 0xea, 0x85, 0x6a, 0x87, 0xad, 0xbc, 0x2e, 0xcd, 0x85, 0x9a, 0x4c, 0x84, 0x25, 0xfb,
	0x11, 0x8a, 0x9b, 0xd9, 0x16, 0xe7, 0x86, 0x61, 0x63, 0x2a, 0xe4, 0xe5, 0x80, 0x34, 0xca, 0x61,
	0x3c, 0xc4, 0x2f, 0xa1, 0x41, 0xa7, 0xe2, 0xab, 0xce, 0xc9, 0x41, 0xf1, 0xc1, 0x23, 0x1c, 0x41,
	0x5b, 0xf3, 0x9f, 0x33, 0x6e, 0xec, 0x80, 0x5a, 0x4a, 0x9a, 0xc5, 0x9c, 0x9b, 0x94, 0xab, 0xeb,
	0x61, 0xd9, 0x5d, 0xab, 0x9c, 0x61, 0x93, 0xc3, 0x1f, 0xe1, 0xf0, 0x46, 0x68, 0x63, 0xab, 0x5b,
	0x25, 0x10, 0xa1, 0xb8, 0x7d, 0x76, 0x9a, 0xfc, 0xf3, 0x55, 0x92, 0x4a, 0xd6, 0xaf, 0xdf, 0xff,
	0x39, 0xad, 0x65, 0xdb, 0xb9, 0xf8, 0x13, 0x1c, 0x19, 0xce, 0x94, 0x1c, 0xae, 0xdd, 0xda, 0x4f,
	0x71, 0xdb, 0x49, 0xc6, 0x27, 0xd0, 0x70, 0x77, 0xa4, 0x0d, 0xe9, 0x44, 0x41, 0xdc, 0xf2, 0x2a,
	0xcf, 0xe1, 0x6f, 0x00, 0x65, 0xf4, 0x81, 0x9a, 0x31, 0x39, 0x8c, 0x82, 0xb8, 0x7d, 0x76, 0xfe,
	0x9f, 0x42, 0x9b, 0x0f, 0x9e, 0x5c, 0xad, 0xb3, 0xde, 0x4b, 0xab, 0x17, 0xde, 0x76, 0xc3, 0xac,
	0xfb, 0x16, 0x9e, 0xed, 0x88, 0xf0, 0x73, 0x08, 0x6e, 0xf9, 0xc2, 0xef, 0x88, 0x0b, 0xdd, 0xde,
	0xcc, 0x69, 0x3e, 0xab, 0x36, 0xac, 0x04, 0x6f, 0xf6, 0x5e, 0xa3, 0x7e, 0xff, 0x7e, 0x19, 0xa2,
	0x87, 0x65, 0x88, 0x1e, 0x97, 0x21, 0xfa, 0xb5, 0x0a, 0x6b, 0x0f, 0xab, 0xb0, 0xf6, 0x7b, 0x15,
	0xd6, 0xbe, 0xc7, 0x23, 0x61, 0xc7, 0xb3, 0xeb, 0x84, 0xa9, 0x49, 0xea, 0x3b, 0x2d, 0xce, 0xf4,
	0x6e, 0xfd, 0x6b, 0xa4, 0x76, 0x31, 0xe5, 0xe6, 0xba, 0x51, 0xec, 0xfb, 0xf9, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x35, 0xc9, 0x1e, 0xd7, 0x3c, 0x03, 0x00, 0x00,
}

func (m *Provider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Provider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Provider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Response) > 0 {
		i -= len(m.Response)
		copy(dAtA[i:], m.Response)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.Response)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConflictVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConflictVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConflictVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VotersHash) > 0 {
		for k := range m.VotersHash {
			v := m.VotersHash[k]
			baseI := i
			if len(v) > 0 {
				i -= len(v)
				copy(dAtA[i:], v)
				i = encodeVarintConflictVote(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintConflictVote(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintConflictVote(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.Voters) > 0 {
		for iNdEx := len(m.Voters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Voters[iNdEx])
			copy(dAtA[i:], m.Voters[iNdEx])
			i = encodeVarintConflictVote(dAtA, i, uint64(len(m.Voters[iNdEx])))
			i--
			dAtA[i] = 0x62
		}
	}
	{
		size, err := m.SecondProvider.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConflictVote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size, err := m.FirstProvider.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintConflictVote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.RequestBlock != 0 {
		i = encodeVarintConflictVote(dAtA, i, uint64(m.RequestBlock))
		i--
		dAtA[i] = 0x48
	}
	if len(m.RequestData) > 0 {
		i -= len(m.RequestData)
		copy(dAtA[i:], m.RequestData)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.RequestData)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ApiUrl) > 0 {
		i -= len(m.ApiUrl)
		copy(dAtA[i:], m.ApiUrl)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.ApiUrl)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x32
	}
	if m.VoteIsCommit {
		i--
		if m.VoteIsCommit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.VoteStartBlock != 0 {
		i = encodeVarintConflictVote(dAtA, i, uint64(m.VoteStartBlock))
		i--
		dAtA[i] = 0x20
	}
	if m.VoteDeadline != 0 {
		i = encodeVarintConflictVote(dAtA, i, uint64(m.VoteDeadline))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintConflictVote(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConflictVote(dAtA []byte, offset int, v uint64) int {
	offset -= sovConflictVote(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Provider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	l = len(m.Response)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	return n
}

func (m *ConflictVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	if m.VoteDeadline != 0 {
		n += 1 + sovConflictVote(uint64(m.VoteDeadline))
	}
	if m.VoteStartBlock != 0 {
		n += 1 + sovConflictVote(uint64(m.VoteStartBlock))
	}
	if m.VoteIsCommit {
		n += 2
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	l = len(m.ApiUrl)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	l = len(m.RequestData)
	if l > 0 {
		n += 1 + l + sovConflictVote(uint64(l))
	}
	if m.RequestBlock != 0 {
		n += 1 + sovConflictVote(uint64(m.RequestBlock))
	}
	l = m.FirstProvider.Size()
	n += 1 + l + sovConflictVote(uint64(l))
	l = m.SecondProvider.Size()
	n += 1 + l + sovConflictVote(uint64(l))
	if len(m.Voters) > 0 {
		for _, s := range m.Voters {
			l = len(s)
			n += 1 + l + sovConflictVote(uint64(l))
		}
	}
	if len(m.VotersHash) > 0 {
		for k, v := range m.VotersHash {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovConflictVote(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovConflictVote(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovConflictVote(uint64(mapEntrySize))
		}
	}
	return n
}

func sovConflictVote(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConflictVote(x uint64) (n int) {
	return sovConflictVote(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Provider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictVote
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Provider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Provider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Response = append(m.Response[:0], dAtA[iNdEx:postIndex]...)
			if m.Response == nil {
				m.Response = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConflictVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictVote
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConflictVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictVote
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConflictVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConflictVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteDeadline", wireType)
			}
			m.VoteDeadline = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteDeadline |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteStartBlock", wireType)
			}
			m.VoteStartBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteStartBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteIsCommit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.VoteIsCommit = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestData = append(m.RequestData[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestData == nil {
				m.RequestData = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestBlock", wireType)
			}
			m.RequestBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstProvider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FirstProvider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecondProvider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SecondProvider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voters = append(m.Voters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotersHash", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConflictVote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictVote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VotersHash == nil {
				m.VotersHash = make(map[string][]byte)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConflictVote
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConflictVote
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthConflictVote
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthConflictVote
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConflictVote
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthConflictVote
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex < 0 {
						return ErrInvalidLengthConflictVote
					}
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipConflictVote(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthConflictVote
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.VotersHash[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConflictVote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictVote
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConflictVote(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConflictVote
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConflictVote
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthConflictVote
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConflictVote
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConflictVote
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConflictVote        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConflictVote          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConflictVote = fmt.Errorf("proto: unexpected end of group")
)
