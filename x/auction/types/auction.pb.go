// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/auction/auction.proto

package types

import (
	fmt "fmt"
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

type Auction struct {
	Creator             string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id                  uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	InitialPrice        string `protobuf:"bytes,4,opt,name=initialPrice,proto3" json:"initialPrice,omitempty"`
	Duration            uint64 `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	CreatedAt           int64  `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	CurrentHighestBidID uint64 `protobuf:"varint,7,opt,name=currentHighestBidID,proto3" json:"currentHighestBidID,omitempty"`
	HighestBidIsPresent bool   `protobuf:"varint,8,opt,name=highestBidIsPresent,proto3" json:"highestBidIsPresent,omitempty"`
	Ended               bool   `protobuf:"varint,9,opt,name=ended,proto3" json:"ended,omitempty"`
}

func (m *Auction) Reset()         { *m = Auction{} }
func (m *Auction) String() string { return proto.CompactTextString(m) }
func (*Auction) ProtoMessage()    {}
func (*Auction) Descriptor() ([]byte, []int) {
	return fileDescriptor_028c42bd7eb07429, []int{0}
}
func (m *Auction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Auction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Auction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Auction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auction.Merge(m, src)
}
func (m *Auction) XXX_Size() int {
	return m.Size()
}
func (m *Auction) XXX_DiscardUnknown() {
	xxx_messageInfo_Auction.DiscardUnknown(m)
}

var xxx_messageInfo_Auction proto.InternalMessageInfo

func (m *Auction) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Auction) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Auction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Auction) GetInitialPrice() string {
	if m != nil {
		return m.InitialPrice
	}
	return ""
}

func (m *Auction) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Auction) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Auction) GetCurrentHighestBidID() uint64 {
	if m != nil {
		return m.CurrentHighestBidID
	}
	return 0
}

func (m *Auction) GetHighestBidIsPresent() bool {
	if m != nil {
		return m.HighestBidIsPresent
	}
	return false
}

func (m *Auction) GetEnded() bool {
	if m != nil {
		return m.Ended
	}
	return false
}

func init() {
	proto.RegisterType((*Auction)(nil), "auction.auction.Auction")
}

func init() { proto.RegisterFile("auction/auction/auction.proto", fileDescriptor_028c42bd7eb07429) }

var fileDescriptor_028c42bd7eb07429 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xb3, 0x69, 0xda, 0x24, 0x8b, 0x28, 0x8c, 0x82, 0x83, 0xe8, 0x12, 0x7a, 0xca, 0xc9,
	0x3f, 0xf8, 0x04, 0x2d, 0x1e, 0xf4, 0x56, 0x72, 0xf4, 0xb6, 0x66, 0x07, 0xbb, 0xa0, 0x9b, 0xb2,
	0xd9, 0x80, 0xbe, 0x85, 0x6f, 0xe0, 0xeb, 0x78, 0xec, 0xd1, 0xa3, 0x24, 0x2f, 0x22, 0xdd, 0x9a,
	0x16, 0x4b, 0x4f, 0x33, 0xbf, 0xf9, 0x66, 0xe6, 0x83, 0x8f, 0x5f, 0xc8, 0xa6, 0x74, 0xba, 0x32,
	0x57, 0x3b, 0xf5, 0x72, 0x61, 0x2b, 0x57, 0xc1, 0x51, 0x8f, 0x7f, 0x75, 0xfc, 0x19, 0xf2, 0x78,
	0xb2, 0xee, 0x01, 0x79, 0x5c, 0x5a, 0x92, 0xae, 0xb2, 0xc8, 0x32, 0x96, 0xa7, 0x45, 0x8f, 0x70,
	0xc8, 0x43, 0xad, 0x30, 0xcc, 0x58, 0x1e, 0x15, 0xa1, 0x56, 0x00, 0x3c, 0x32, 0xf2, 0x95, 0x70,
	0xe0, 0xd7, 0x7c, 0x0f, 0x63, 0x7e, 0xa0, 0x8d, 0x76, 0x5a, 0xbe, 0xcc, 0xac, 0x2e, 0x09, 0x23,
	0xaf, 0xfd, 0x9b, 0xc1, 0x19, 0x4f, 0x54, 0x63, 0xe5, 0xca, 0x0d, 0x87, 0xfe, 0xdb, 0x86, 0xe1,
	0x9c, 0xa7, 0xde, 0x8e, 0xd4, 0xc4, 0xe1, 0x28, 0x63, 0xf9, 0xa0, 0xd8, 0x0e, 0xe0, 0x9a, 0x1f,
	0x97, 0x8d, 0xb5, 0x64, 0xdc, 0xbd, 0x7e, 0x9e, 0x53, 0xed, 0xa6, 0x5a, 0x3d, 0xdc, 0x61, 0xec,
	0x9f, 0xec, 0x93, 0x56, 0x17, 0xf3, 0x2d, 0xd7, 0x33, 0x4b, 0x35, 0x19, 0x87, 0x49, 0xc6, 0xf2,
	0xa4, 0xd8, 0x27, 0xc1, 0x09, 0x1f, 0x92, 0x51, 0xa4, 0x30, 0xf5, 0x3b, 0x6b, 0x98, 0xde, 0x7c,
	0xb5, 0x82, 0x2d, 0x5b, 0xc1, 0x7e, 0x5a, 0xc1, 0x3e, 0x3a, 0x11, 0x2c, 0x3b, 0x11, 0x7c, 0x77,
	0x22, 0x78, 0x3c, 0xed, 0x33, 0x7e, 0xdb, 0xa4, 0xed, 0xde, 0x17, 0x54, 0x3f, 0x8d, 0x7c, 0xd8,
	0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x00, 0x17, 0x16, 0x11, 0x8d, 0x01, 0x00, 0x00,
}

func (m *Auction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Auction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Auction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ended {
		i--
		if m.Ended {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.HighestBidIsPresent {
		i--
		if m.HighestBidIsPresent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.CurrentHighestBidID != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.CurrentHighestBidID))
		i--
		dAtA[i] = 0x38
	}
	if m.CreatedAt != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x30
	}
	if m.Duration != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x28
	}
	if len(m.InitialPrice) > 0 {
		i -= len(m.InitialPrice)
		copy(dAtA[i:], m.InitialPrice)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.InitialPrice)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Auction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovAuction(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = len(m.InitialPrice)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovAuction(uint64(m.Duration))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovAuction(uint64(m.CreatedAt))
	}
	if m.CurrentHighestBidID != 0 {
		n += 1 + sovAuction(uint64(m.CurrentHighestBidID))
	}
	if m.HighestBidIsPresent {
		n += 2
	}
	if m.Ended {
		n += 2
	}
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Auction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Auction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Auction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentHighestBidID", wireType)
			}
			m.CurrentHighestBidID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentHighestBidID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HighestBidIsPresent", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
			m.HighestBidIsPresent = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ended", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
			m.Ended = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)
