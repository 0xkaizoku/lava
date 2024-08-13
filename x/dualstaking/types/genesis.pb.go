// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/dualstaking/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/lavanet/lava/v2/x/fixationstore/types"
	_ "github.com/lavanet/lava/v2/x/timerstore/types"
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

// GenesisState defines the dualstaking module's genesis state.
type GenesisState struct {
	Params              Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	DelegationsFS       types.GenesisState `protobuf:"bytes,2,opt,name=delegationsFS,proto3" json:"delegationsFS"`
	DelegatorsFS        types.GenesisState `protobuf:"bytes,3,opt,name=delegatorsFS,proto3" json:"delegatorsFS"`
	DelegatorRewardList []DelegatorReward  `protobuf:"bytes,5,rep,name=delegator_reward_list,json=delegatorRewardList,proto3" json:"delegator_reward_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5bca863c53f218f, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetDelegationsFS() types.GenesisState {
	if m != nil {
		return m.DelegationsFS
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetDelegatorsFS() types.GenesisState {
	if m != nil {
		return m.DelegatorsFS
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetDelegatorRewardList() []DelegatorReward {
	if m != nil {
		return m.DelegatorRewardList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lavanet.lava.dualstaking.GenesisState")
}

func init() {
	proto.RegisterFile("lavanet/lava/dualstaking/genesis.proto", fileDescriptor_d5bca863c53f218f)
}

var fileDescriptor_d5bca863c53f218f = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0x4b, 0xf3, 0x30,
	0x1c, 0xc7, 0xdb, 0xfd, 0xe3, 0xa1, 0xdb, 0x03, 0x52, 0x15, 0xca, 0x0e, 0x75, 0x28, 0xca, 0x76,
	0x49, 0x60, 0xde, 0x3d, 0x0c, 0x51, 0x11, 0x0f, 0xb2, 0x79, 0xf2, 0x32, 0xb2, 0x35, 0xc6, 0x60,
	0xd7, 0x8c, 0xe4, 0xb7, 0x39, 0xdf, 0x85, 0x2f, 0x6b, 0xc7, 0x1d, 0x3d, 0x89, 0xb4, 0x6f, 0x44,
	0x9a, 0xc6, 0xb1, 0x08, 0xbd, 0x78, 0x4a, 0x5a, 0x3e, 0xdf, 0x4f, 0xf2, 0xcd, 0xcf, 0x3b, 0x8b,
	0xc9, 0x92, 0x24, 0x14, 0x70, 0xbe, 0xe2, 0x68, 0x41, 0x62, 0x05, 0xe4, 0x85, 0x27, 0x0c, 0x33,
	0x9a, 0x50, 0xc5, 0x15, 0x9a, 0x4b, 0x01, 0xc2, 0x0f, 0x0c, 0x87, 0xf2, 0x15, 0xed, 0x70, 0xed,
	0x03, 0x26, 0x98, 0xd0, 0x10, 0xce, 0x77, 0x05, 0xdf, 0x3e, 0x2d, 0xf5, 0xce, 0x89, 0x24, 0x33,
	0xa3, 0x6d, 0xf7, 0x2c, 0xec, 0x89, 0xaf, 0x08, 0x70, 0x91, 0x28, 0x10, 0x92, 0x6e, 0xbf, 0x0c,
	0x7a, 0x62, 0xa1, 0xc0, 0x67, 0x54, 0x16, 0x9c, 0xde, 0x1a, 0x08, 0x97, 0x1e, 0x1b, 0xd1, 0x98,
	0x32, 0x02, 0x42, 0x8e, 0x25, 0x7d, 0x25, 0x32, 0x2a, 0x02, 0xc7, 0x59, 0xc5, 0x6b, 0x5d, 0x17,
	0x4d, 0x47, 0x40, 0x80, 0xfa, 0x17, 0x5e, 0xa3, 0xb8, 0x61, 0xe0, 0x76, 0xdc, 0x6e, 0xb3, 0xdf,
	0x41, 0x65, 0xcd, 0xd1, 0xbd, 0xe6, 0x06, 0xb5, 0xf5, 0xe7, 0x91, 0x33, 0x34, 0x29, 0xff, 0xc1,
	0xfb, 0x6f, 0x8e, 0xca, 0x8b, 0x5c, 0x8d, 0x82, 0x8a, 0xd6, 0x74, 0x6d, 0x8d, 0xd5, 0x14, 0xed,
	0x5e, 0xc0, 0xe8, 0x6c, 0x89, 0x3f, 0xf4, 0x5a, 0xdb, 0x02, 0xb9, 0xb4, 0xfa, 0x27, 0xa9, 0xe5,
	0xf0, 0xa7, 0xde, 0xe1, 0xef, 0x47, 0x19, 0xc7, 0x5c, 0x41, 0x50, 0xef, 0x54, 0xbb, 0xcd, 0x7e,
	0xaf, 0xbc, 0xf8, 0xe5, 0x4f, 0x6c, 0xa8, 0x53, 0xc6, 0xbe, 0x1f, 0xd9, 0xbf, 0xef, 0xb8, 0x82,
	0xdb, 0xda, 0xbf, 0xda, 0x5e, 0x7d, 0x70, 0xb3, 0x4e, 0x43, 0x77, 0x93, 0x86, 0xee, 0x57, 0x1a,
	0xba, 0xef, 0x59, 0xe8, 0x6c, 0xb2, 0xd0, 0xf9, 0xc8, 0x42, 0xe7, 0x11, 0x31, 0x0e, 0xcf, 0x8b,
	0x09, 0x9a, 0x8a, 0x99, 0x3d, 0xbb, 0x65, 0x1f, 0xaf, 0xac, 0x01, 0xc2, 0xdb, 0x9c, 0xaa, 0x49,
	0x43, 0x8f, 0xed, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xc4, 0xa7, 0x53, 0xb8, 0x02, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DelegatorRewardList) > 0 {
		for iNdEx := len(m.DelegatorRewardList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatorRewardList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size, err := m.DelegatorsFS.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.DelegationsFS.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.DelegationsFS.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.DelegatorsFS.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.DelegatorRewardList) > 0 {
		for _, e := range m.DelegatorRewardList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationsFS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegationsFS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorsFS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegatorsFS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorRewardList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorRewardList = append(m.DelegatorRewardList, DelegatorReward{})
			if err := m.DelegatorRewardList[len(m.DelegatorRewardList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
