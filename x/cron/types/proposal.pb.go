// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/cron/v1/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// PromoteToPrivilegedContractProposal gov proposal content type to add
// "privileges" to a contract
type PromoteToPrivilegedContractProposal struct {
	// Title is a short summary
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	// Description is a human readable text
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	// Contract is the bech32 address of the smart contract
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty" yaml:"contract"`
}

func (m *PromoteToPrivilegedContractProposal) Reset()         { *m = PromoteToPrivilegedContractProposal{} }
func (m *PromoteToPrivilegedContractProposal) String() string { return proto.CompactTextString(m) }
func (*PromoteToPrivilegedContractProposal) ProtoMessage()    {}
func (*PromoteToPrivilegedContractProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3c0dd977741a5f7, []int{0}
}
func (m *PromoteToPrivilegedContractProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PromoteToPrivilegedContractProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PromoteToPrivilegedContractProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PromoteToPrivilegedContractProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromoteToPrivilegedContractProposal.Merge(m, src)
}
func (m *PromoteToPrivilegedContractProposal) XXX_Size() int {
	return m.Size()
}
func (m *PromoteToPrivilegedContractProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_PromoteToPrivilegedContractProposal.DiscardUnknown(m)
}

var xxx_messageInfo_PromoteToPrivilegedContractProposal proto.InternalMessageInfo

func (m *PromoteToPrivilegedContractProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PromoteToPrivilegedContractProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PromoteToPrivilegedContractProposal) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

// DemotePrivilegedContractProposal gov proposal content type to remove
// "privileges" from a contract
type DemotePrivilegedContractProposal struct {
	// Title is a short summary
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	// Description is a human readable text
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	// Contract is the bech32 address of the smart contract
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty" yaml:"contract"`
}

func (m *DemotePrivilegedContractProposal) Reset()         { *m = DemotePrivilegedContractProposal{} }
func (m *DemotePrivilegedContractProposal) String() string { return proto.CompactTextString(m) }
func (*DemotePrivilegedContractProposal) ProtoMessage()    {}
func (*DemotePrivilegedContractProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3c0dd977741a5f7, []int{1}
}
func (m *DemotePrivilegedContractProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DemotePrivilegedContractProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DemotePrivilegedContractProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DemotePrivilegedContractProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemotePrivilegedContractProposal.Merge(m, src)
}
func (m *DemotePrivilegedContractProposal) XXX_Size() int {
	return m.Size()
}
func (m *DemotePrivilegedContractProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_DemotePrivilegedContractProposal.DiscardUnknown(m)
}

var xxx_messageInfo_DemotePrivilegedContractProposal proto.InternalMessageInfo

func (m *DemotePrivilegedContractProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DemotePrivilegedContractProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DemotePrivilegedContractProposal) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func init() {
	proto.RegisterType((*PromoteToPrivilegedContractProposal)(nil), "publicawesome.stargaze.cron.v1.PromoteToPrivilegedContractProposal")
	proto.RegisterType((*DemotePrivilegedContractProposal)(nil), "publicawesome.stargaze.cron.v1.DemotePrivilegedContractProposal")
}

func init() { proto.RegisterFile("stargaze/cron/v1/proposal.proto", fileDescriptor_c3c0dd977741a5f7) }

var fileDescriptor_c3c0dd977741a5f7 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x92, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x17, 0x45, 0xd1, 0x2a, 0x28, 0x55, 0x64, 0xee, 0x90, 0x8e, 0x0a, 0xe2, 0x65, 0x09,
	0x65, 0x17, 0xd9, 0x71, 0x7a, 0x15, 0xc6, 0xf0, 0xe4, 0x45, 0xd2, 0x2c, 0xd4, 0x40, 0xdb, 0xaf,
	0x24, 0x59, 0x74, 0x3e, 0x85, 0x0f, 0xe3, 0x43, 0x78, 0x1c, 0x82, 0xb0, 0xd3, 0x90, 0xed, 0x0d,
	0xf6, 0x04, 0xd2, 0xa6, 0x1b, 0x7b, 0x06, 0x6f, 0x49, 0xfe, 0xbf, 0x8f, 0x2f, 0x7f, 0xf8, 0x79,
	0x81, 0x36, 0x4c, 0x25, 0xec, 0x5d, 0x50, 0xae, 0x20, 0xa7, 0x36, 0xa2, 0x85, 0x82, 0x02, 0x34,
	0x4b, 0x49, 0xa1, 0xc0, 0x80, 0x8f, 0x8b, 0x71, 0x9c, 0x4a, 0xce, 0x5e, 0x85, 0x86, 0x4c, 0x90,
	0x35, 0x4e, 0x4a, 0x9c, 0xd8, 0xa8, 0x75, 0x9e, 0x40, 0x02, 0x15, 0x4a, 0xcb, 0x93, 0x9b, 0x6a,
	0x5d, 0x72, 0xd0, 0x19, 0xe8, 0x67, 0x17, 0xb8, 0x8b, 0x8b, 0xc2, 0x19, 0xf2, 0xae, 0x06, 0x0a,
	0x32, 0x30, 0xe2, 0x11, 0x06, 0x4a, 0x5a, 0x99, 0x8a, 0x44, 0x8c, 0xee, 0x20, 0x37, 0x8a, 0x71,
	0x33, 0xa8, 0xd7, 0xfb, 0xd7, 0xde, 0x9e, 0x91, 0x26, 0x15, 0x4d, 0xd4, 0x46, 0x37, 0x87, 0xfd,
	0xd3, 0xd5, 0x3c, 0x38, 0x9e, 0xb0, 0x2c, 0xed, 0x85, 0xd5, 0x73, 0x38, 0x74, 0xb1, 0x7f, 0xeb,
	0x1d, 0x8d, 0x84, 0xe6, 0x4a, 0x16, 0x46, 0x42, 0xde, 0xdc, 0xa9, 0xe8, 0x8b, 0xd5, 0x3c, 0xf0,
	0x1d, 0xbd, 0x15, 0x86, 0xc3, 0x6d, 0xd4, 0xa7, 0xde, 0x01, 0xaf, 0xb7, 0x36, 0x77, 0xab, 0xb1,
	0xb3, 0xd5, 0x3c, 0x38, 0x71, 0x63, 0xeb, 0x24, 0x1c, 0x6e, 0xa0, 0x1e, 0xfe, 0xfe, 0xec, 0xb4,
	0xea, 0x32, 0x09, 0x58, 0x62, 0xa3, 0x58, 0x18, 0x16, 0x91, 0xf2, 0xef, 0x22, 0x37, 0xe1, 0x0f,
	0xf2, 0xda, 0xf7, 0xa2, 0x6c, 0xf6, 0xaf, 0x7a, 0xf5, 0x1f, 0xbe, 0x16, 0x18, 0x4d, 0x17, 0x18,
	0xfd, 0x2e, 0x30, 0xfa, 0x58, 0xe2, 0xc6, 0x74, 0x89, 0x1b, 0xb3, 0x25, 0x6e, 0x3c, 0x75, 0x13,
	0x69, 0x5e, 0xc6, 0x31, 0xe1, 0x90, 0x51, 0x27, 0x4a, 0xa7, 0x36, 0x85, 0x6e, 0xc4, 0xb2, 0x51,
	0x44, 0xdf, 0x9c, 0x5e, 0x66, 0x52, 0x08, 0x1d, 0xef, 0x57, 0x22, 0x74, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x06, 0x16, 0x12, 0xef, 0x7c, 0x02, 0x00, 0x00,
}

func (m *PromoteToPrivilegedContractProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PromoteToPrivilegedContractProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PromoteToPrivilegedContractProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DemotePrivilegedContractProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DemotePrivilegedContractProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DemotePrivilegedContractProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PromoteToPrivilegedContractProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *DemotePrivilegedContractProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PromoteToPrivilegedContractProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: PromoteToPrivilegedContractProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PromoteToPrivilegedContractProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *DemotePrivilegedContractProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: DemotePrivilegedContractProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DemotePrivilegedContractProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
