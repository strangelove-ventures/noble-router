// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: router/ibc_forward_metadata.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

// IBCForwards are stored so incoming mints can check for a forward message
// @param source_domain_sender
// @param nonce
// @param port
// @param channel
// @param destination_receiver
// @param ack_error
type StoreIBCForwardMetadata struct {
	SourceDomainSender string              `protobuf:"bytes,1,opt,name=source_domain_sender,json=sourceDomainSender,proto3" json:"source_domain_sender,omitempty"`
	Nonce              uint64              `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Metadata           *IBCForwardMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	AckError           bool                `protobuf:"varint,4,opt,name=ack_error,json=ackError,proto3" json:"ack_error,omitempty"`
}

func (m *StoreIBCForwardMetadata) Reset()         { *m = StoreIBCForwardMetadata{} }
func (m *StoreIBCForwardMetadata) String() string { return proto.CompactTextString(m) }
func (*StoreIBCForwardMetadata) ProtoMessage()    {}
func (*StoreIBCForwardMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6b29f4e31c7ab9, []int{0}
}
func (m *StoreIBCForwardMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreIBCForwardMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreIBCForwardMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreIBCForwardMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreIBCForwardMetadata.Merge(m, src)
}
func (m *StoreIBCForwardMetadata) XXX_Size() int {
	return m.Size()
}
func (m *StoreIBCForwardMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreIBCForwardMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_StoreIBCForwardMetadata proto.InternalMessageInfo

func (m *StoreIBCForwardMetadata) GetSourceDomainSender() string {
	if m != nil {
		return m.SourceDomainSender
	}
	return ""
}

func (m *StoreIBCForwardMetadata) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *StoreIBCForwardMetadata) GetMetadata() *IBCForwardMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StoreIBCForwardMetadata) GetAckError() bool {
	if m != nil {
		return m.AckError
	}
	return false
}

// IBCForwardMetadata is the information a user includes in their depositForBurnWithMetadata data field
// TODO
// @param port
// @param channel
// @param destination_receiver
// @param memo
// @param timeout_in_nanoseconds
type IBCForwardMetadata struct {
	Port                 string `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	Channel              string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	DestinationReceiver  string `protobuf:"bytes,3,opt,name=destination_receiver,json=destinationReceiver,proto3" json:"destination_receiver,omitempty"`
	Memo                 string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo,omitempty"`
	TimeoutInNanoseconds uint64 `protobuf:"varint,5,opt,name=timeout_in_nanoseconds,json=timeoutInNanoseconds,proto3" json:"timeout_in_nanoseconds,omitempty"`
}

func (m *IBCForwardMetadata) Reset()         { *m = IBCForwardMetadata{} }
func (m *IBCForwardMetadata) String() string { return proto.CompactTextString(m) }
func (*IBCForwardMetadata) ProtoMessage()    {}
func (*IBCForwardMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6b29f4e31c7ab9, []int{1}
}
func (m *IBCForwardMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCForwardMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCForwardMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCForwardMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCForwardMetadata.Merge(m, src)
}
func (m *IBCForwardMetadata) XXX_Size() int {
	return m.Size()
}
func (m *IBCForwardMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCForwardMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_IBCForwardMetadata proto.InternalMessageInfo

func (m *IBCForwardMetadata) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *IBCForwardMetadata) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *IBCForwardMetadata) GetDestinationReceiver() string {
	if m != nil {
		return m.DestinationReceiver
	}
	return ""
}

func (m *IBCForwardMetadata) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *IBCForwardMetadata) GetTimeoutInNanoseconds() uint64 {
	if m != nil {
		return m.TimeoutInNanoseconds
	}
	return 0
}

func init() {
	proto.RegisterType((*StoreIBCForwardMetadata)(nil), "noble.router.StoreIBCForwardMetadata")
	proto.RegisterType((*IBCForwardMetadata)(nil), "noble.router.IBCForwardMetadata")
}

func init() { proto.RegisterFile("router/ibc_forward_metadata.proto", fileDescriptor_0b6b29f4e31c7ab9) }

var fileDescriptor_0b6b29f4e31c7ab9 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbf, 0x6e, 0xd4, 0x40,
	0x10, 0xc6, 0x6f, 0xe1, 0x02, 0xe7, 0x85, 0x6a, 0xb1, 0xc0, 0x0a, 0x92, 0x65, 0x52, 0x5d, 0x13,
	0x9b, 0x03, 0x3a, 0xa8, 0xc2, 0x1f, 0x29, 0x05, 0x14, 0x8e, 0x68, 0x68, 0x56, 0xeb, 0xf5, 0x70,
	0x59, 0xe5, 0x3c, 0x73, 0x9a, 0x5d, 0x1f, 0xf0, 0x16, 0x3c, 0x10, 0x12, 0x2d, 0x65, 0x4a, 0x4a,
	0x74, 0xf7, 0x22, 0xe8, 0xd6, 0x0e, 0x44, 0x82, 0xee, 0x1b, 0xff, 0xc6, 0xfa, 0x7e, 0x1a, 0x5b,
	0x3e, 0x62, 0xea, 0x03, 0x70, 0xe5, 0x1a, 0xab, 0x3f, 0x12, 0x7f, 0x32, 0xdc, 0xea, 0x0e, 0x82,
	0x69, 0x4d, 0x30, 0xe5, 0x9a, 0x29, 0x90, 0xba, 0x8b, 0xd4, 0xac, 0xa0, 0x1c, 0x16, 0x0f, 0x73,
	0x4b, 0xbe, 0x23, 0x5f, 0x35, 0xc6, 0x43, 0xb5, 0x59, 0x34, 0x10, 0xcc, 0xa2, 0xb2, 0xe4, 0x70,
	0xd8, 0x3e, 0x4c, 0x97, 0xb4, 0xa4, 0x18, 0xab, 0x7d, 0x1a, 0x9e, 0x1e, 0x7d, 0x13, 0xf2, 0xc1,
	0x59, 0x20, 0x86, 0xd3, 0x93, 0x97, 0x6f, 0x86, 0x9a, 0xb7, 0x63, 0x8b, 0x7a, 0x2c, 0x53, 0x4f,
	0x3d, 0x5b, 0xd0, 0x2d, 0x75, 0xc6, 0xa1, 0xf6, 0x80, 0x2d, 0x70, 0x26, 0x0a, 0x31, 0x4f, 0x6a,
	0x35, 0xb0, 0x57, 0x11, 0x9d, 0x45, 0xa2, 0x52, 0x79, 0x80, 0x84, 0x16, 0xb2, 0x1b, 0x85, 0x98,
	0x4f, 0xeb, 0x61, 0x50, 0x2f, 0xe4, 0xec, 0xca, 0x3c, 0xbb, 0x59, 0x88, 0xf9, 0x9d, 0x27, 0x45,
	0x79, 0x5d, 0xbd, 0xfc, 0xb7, 0xbb, 0xfe, 0xf3, 0x86, 0x7a, 0x28, 0x13, 0x63, 0x2f, 0x34, 0x30,
	0x13, 0x67, 0xd3, 0x42, 0xcc, 0x67, 0xf5, 0xcc, 0xd8, 0x8b, 0xd7, 0xfb, 0xf9, 0xe8, 0xbb, 0x90,
	0xea, 0x3f, 0xe6, 0x4a, 0x4e, 0xd7, 0xc4, 0x61, 0x34, 0x8d, 0x59, 0x65, 0xf2, 0xb6, 0x3d, 0x37,
	0x88, 0xb0, 0x8a, 0x76, 0x49, 0x7d, 0x35, 0xaa, 0x85, 0x4c, 0x5b, 0xf0, 0xc1, 0xa1, 0x09, 0x8e,
	0x50, 0x33, 0x58, 0x70, 0x1b, 0xe0, 0xe8, 0x9a, 0xd4, 0xf7, 0xae, 0xb1, 0x7a, 0x44, 0xfb, 0x82,
	0x0e, 0x3a, 0x8a, 0x3e, 0x49, 0x1d, 0xb3, 0x7a, 0x26, 0xef, 0x07, 0xd7, 0x01, 0xf5, 0x41, 0x3b,
	0xd4, 0x68, 0x90, 0x3c, 0x58, 0xc2, 0xd6, 0x67, 0x07, 0xf1, 0x1a, 0xe9, 0x48, 0x4f, 0xf1, 0xdd,
	0x5f, 0x76, 0xf2, 0xfe, 0xc7, 0x36, 0x17, 0x97, 0xdb, 0x5c, 0xfc, 0xda, 0xe6, 0xe2, 0xeb, 0x2e,
	0x9f, 0x5c, 0xee, 0xf2, 0xc9, 0xcf, 0x5d, 0x3e, 0xf9, 0xf0, 0x7c, 0xe9, 0xc2, 0x79, 0xdf, 0x94,
	0x96, 0xba, 0xca, 0x07, 0x36, 0xb8, 0x84, 0x15, 0x6d, 0xe0, 0x78, 0x03, 0x18, 0x7a, 0x06, 0x5f,
	0xc5, 0x1b, 0x1e, 0x8f, 0xff, 0xc9, 0xe7, 0x6a, 0x0c, 0xe1, 0xcb, 0x1a, 0x7c, 0x73, 0x2b, 0x7e,
	0xde, 0xa7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x78, 0xcd, 0xf6, 0x47, 0x02, 0x00, 0x00,
}

func (m *StoreIBCForwardMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreIBCForwardMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoreIBCForwardMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AckError {
		i--
		if m.AckError {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIbcForwardMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Nonce != 0 {
		i = encodeVarintIbcForwardMetadata(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SourceDomainSender) > 0 {
		i -= len(m.SourceDomainSender)
		copy(dAtA[i:], m.SourceDomainSender)
		i = encodeVarintIbcForwardMetadata(dAtA, i, uint64(len(m.SourceDomainSender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IBCForwardMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCForwardMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCForwardMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutInNanoseconds != 0 {
		i = encodeVarintIbcForwardMetadata(dAtA, i, uint64(m.TimeoutInNanoseconds))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintIbcForwardMetadata(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DestinationReceiver) > 0 {
		i -= len(m.DestinationReceiver)
		copy(dAtA[i:], m.DestinationReceiver)
		i = encodeVarintIbcForwardMetadata(dAtA, i, uint64(len(m.DestinationReceiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintIbcForwardMetadata(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintIbcForwardMetadata(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIbcForwardMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovIbcForwardMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoreIBCForwardMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceDomainSender)
	if l > 0 {
		n += 1 + l + sovIbcForwardMetadata(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovIbcForwardMetadata(uint64(m.Nonce))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovIbcForwardMetadata(uint64(l))
	}
	if m.AckError {
		n += 2
	}
	return n
}

func (m *IBCForwardMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovIbcForwardMetadata(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovIbcForwardMetadata(uint64(l))
	}
	l = len(m.DestinationReceiver)
	if l > 0 {
		n += 1 + l + sovIbcForwardMetadata(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovIbcForwardMetadata(uint64(l))
	}
	if m.TimeoutInNanoseconds != 0 {
		n += 1 + sovIbcForwardMetadata(uint64(m.TimeoutInNanoseconds))
	}
	return n
}

func sovIbcForwardMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIbcForwardMetadata(x uint64) (n int) {
	return sovIbcForwardMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreIBCForwardMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbcForwardMetadata
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
			return fmt.Errorf("proto: StoreIBCForwardMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreIBCForwardMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceDomainSender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcForwardMetadata
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
				return ErrInvalidLengthIbcForwardMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbcForwardMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceDomainSender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcForwardMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcForwardMetadata
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
				return ErrInvalidLengthIbcForwardMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIbcForwardMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &IBCForwardMetadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckError", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcForwardMetadata
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
			m.AckError = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipIbcForwardMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbcForwardMetadata
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
func (m *IBCForwardMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbcForwardMetadata
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
			return fmt.Errorf("proto: IBCForwardMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCForwardMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcForwardMetadata
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
				return ErrInvalidLengthIbcForwardMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbcForwardMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcForwardMetadata
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
				return ErrInvalidLengthIbcForwardMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbcForwardMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationReceiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcForwardMetadata
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
				return ErrInvalidLengthIbcForwardMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbcForwardMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationReceiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcForwardMetadata
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
				return ErrInvalidLengthIbcForwardMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbcForwardMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutInNanoseconds", wireType)
			}
			m.TimeoutInNanoseconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbcForwardMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutInNanoseconds |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIbcForwardMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbcForwardMetadata
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
func skipIbcForwardMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIbcForwardMetadata
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
					return 0, ErrIntOverflowIbcForwardMetadata
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
					return 0, ErrIntOverflowIbcForwardMetadata
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
				return 0, ErrInvalidLengthIbcForwardMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIbcForwardMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIbcForwardMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIbcForwardMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIbcForwardMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIbcForwardMetadata = fmt.Errorf("proto: unexpected end of group")
)
