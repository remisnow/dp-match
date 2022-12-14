// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common_error.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CommonError struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
	RetMsg               string   `protobuf:"bytes,2,opt,name=RetMsg,proto3" json:"RetMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonError) Reset()         { *m = CommonError{} }
func (m *CommonError) String() string { return proto.CompactTextString(m) }
func (*CommonError) ProtoMessage()    {}
func (*CommonError) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dcf5a326b03efe, []int{0}
}
func (m *CommonError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommonError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommonError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonError.Merge(m, src)
}
func (m *CommonError) XXX_Size() int {
	return m.Size()
}
func (m *CommonError) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonError.DiscardUnknown(m)
}

var xxx_messageInfo_CommonError proto.InternalMessageInfo

func (m *CommonError) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CommonError) GetRetMsg() string {
	if m != nil {
		return m.RetMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonError)(nil), "msg.CommonError")
}

func init() { proto.RegisterFile("common_error.proto", fileDescriptor_35dcf5a326b03efe) }

var fileDescriptor_35dcf5a326b03efe = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0xce, 0x2d, 0x4e, 0x57, 0xb2, 0xe7, 0xe2, 0x76, 0x06, 0x4b, 0xb9, 0x82, 0x64, 0x84, 0x24, 0xb8,
	0xd8, 0x83, 0x52, 0x4b, 0x9c, 0xf3, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x60,
	0x5c, 0x21, 0x31, 0x2e, 0xb6, 0xa0, 0xd4, 0x12, 0xdf, 0xe2, 0x74, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x28, 0xcf, 0x49, 0xea, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x8a, 0x03, 0x6c, 0x7c, 0x72, 0x7e, 0x4e, 0x12, 0x1b,
	0x98, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xf6, 0x47, 0x71, 0x7e, 0x00, 0x00, 0x00,
}

func (m *CommonError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RetMsg) > 0 {
		i -= len(m.RetMsg)
		copy(dAtA[i:], m.RetMsg)
		i = encodeVarintCommonError(dAtA, i, uint64(len(m.RetMsg)))
		i--
		dAtA[i] = 0x12
	}
	if m.RetCode != 0 {
		i = encodeVarintCommonError(dAtA, i, uint64(m.RetCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommonError(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommonError(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommonError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RetCode != 0 {
		n += 1 + sovCommonError(uint64(m.RetCode))
	}
	l = len(m.RetMsg)
	if l > 0 {
		n += 1 + l + sovCommonError(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCommonError(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommonError(x uint64) (n int) {
	return sovCommonError(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommonError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommonError
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
			return fmt.Errorf("proto: CommonError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetCode", wireType)
			}
			m.RetCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommonError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RetCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommonError
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
				return ErrInvalidLengthCommonError
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommonError
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RetMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommonError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommonError
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommonError(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommonError
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
					return 0, ErrIntOverflowCommonError
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
					return 0, ErrIntOverflowCommonError
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
				return 0, ErrInvalidLengthCommonError
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommonError
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommonError
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommonError        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommonError          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommonError = fmt.Errorf("proto: unexpected end of group")
)
