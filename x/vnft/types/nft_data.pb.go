// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vcoa/vnft/nft_data.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type NftData struct {
	Metadata JsonInput `protobuf:"bytes,1,opt,name=metadata,proto3,customtype=JsonInput" json:"metadata"`
}

func (m *NftData) Reset()         { *m = NftData{} }
func (m *NftData) String() string { return proto.CompactTextString(m) }
func (*NftData) ProtoMessage()    {}
func (*NftData) Descriptor() ([]byte, []int) {
	return fileDescriptor_819db54c383ab251, []int{0}
}
func (m *NftData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NftData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NftData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NftData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NftData.Merge(m, src)
}
func (m *NftData) XXX_Size() int {
	return m.Size()
}
func (m *NftData) XXX_DiscardUnknown() {
	xxx_messageInfo_NftData.DiscardUnknown(m)
}

var xxx_messageInfo_NftData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NftData)(nil), "vcoa.vnft.NftData")
}

func init() { proto.RegisterFile("vcoa/vnft/nft_data.proto", fileDescriptor_819db54c383ab251) }

var fileDescriptor_819db54c383ab251 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x4b, 0xce, 0x4f,
	0xd4, 0x2f, 0xcb, 0x4b, 0x2b, 0xd1, 0xcf, 0x4b, 0x2b, 0x89, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0xc9, 0xe8, 0x81, 0x64, 0xa4, 0x44, 0xd2, 0xf3, 0xd3,
	0xf3, 0xc1, 0xa2, 0xfa, 0x20, 0x16, 0x44, 0x81, 0x94, 0x18, 0x42, 0x6b, 0x41, 0x62, 0x51, 0x62,
	0x6e, 0x31, 0x44, 0x5c, 0xc9, 0x82, 0x8b, 0xdd, 0x2f, 0xad, 0xc4, 0x25, 0xb1, 0x24, 0x51, 0x48,
	0x97, 0x8b, 0x23, 0x37, 0xb5, 0x24, 0x11, 0x64, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x8f, 0x93,
	0xe0, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0x73, 0x7a, 0x15, 0xe7, 0xe7, 0x79, 0xe6, 0x15,
	0x94, 0x96, 0x04, 0xc1, 0x95, 0x38, 0x69, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x43, 0x94, 0x20, 0xd8, 0xae, 0x0a, 0x88, 0x6d, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60,
	0xdb, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x22, 0xf7, 0x71, 0x8c, 0xc2, 0x00, 0x00, 0x00,
}

func (m *NftData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NftData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Metadata.Size()
		i -= size
		if _, err := m.Metadata.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNftData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintNftData(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NftData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Metadata.Size()
	n += 1 + l + sovNftData(uint64(l))
	return n
}

func sovNftData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftData(x uint64) (n int) {
	return sovNftData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NftData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftData
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
			return fmt.Errorf("proto: NftData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NftData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftData
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
				return ErrInvalidLengthNftData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNftData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftData
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
func skipNftData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftData
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
					return 0, ErrIntOverflowNftData
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
					return 0, ErrIntOverflowNftData
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
				return 0, ErrInvalidLengthNftData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftData = fmt.Errorf("proto: unexpected end of group")
)
