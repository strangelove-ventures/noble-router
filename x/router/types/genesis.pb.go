// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: router/genesis.proto

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

// GenesisState defines the router module's genesis state.
type GenesisState struct {
	Params                     Params                      `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Mints                      []Mint                      `protobuf:"bytes,3,rep,name=mints,proto3" json:"mints"`
	IbcForwards                []StoreIBCForwardMetadata   `protobuf:"bytes,4,rep,name=ibc_forwards,json=ibcForwards,proto3" json:"ibc_forwards"`
	InFlightPackets            []InFlightPacket            `protobuf:"bytes,5,rep,name=in_flight_packets,json=inFlightPackets,proto3" json:"in_flight_packets"`
	AllowedSourceDomainSenders []AllowedSourceDomainSender `protobuf:"bytes,6,rep,name=allowed_source_domain_senders,json=allowedSourceDomainSenders,proto3" json:"allowed_source_domain_senders"`
	Owner                      string                      `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d6fb1a9cb128c80, []int{0}
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

func (m *GenesisState) GetMints() []Mint {
	if m != nil {
		return m.Mints
	}
	return nil
}

func (m *GenesisState) GetIbcForwards() []StoreIBCForwardMetadata {
	if m != nil {
		return m.IbcForwards
	}
	return nil
}

func (m *GenesisState) GetInFlightPackets() []InFlightPacket {
	if m != nil {
		return m.InFlightPackets
	}
	return nil
}

func (m *GenesisState) GetAllowedSourceDomainSenders() []AllowedSourceDomainSender {
	if m != nil {
		return m.AllowedSourceDomainSenders
	}
	return nil
}

func (m *GenesisState) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "noble.router.GenesisState")
}

func init() { proto.RegisterFile("router/genesis.proto", fileDescriptor_5d6fb1a9cb128c80) }

var fileDescriptor_5d6fb1a9cb128c80 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0xab, 0xd3, 0x40,
	0x14, 0x85, 0x13, 0xfa, 0x5a, 0x31, 0x2d, 0xc8, 0x8b, 0x59, 0x84, 0xe0, 0x8b, 0x55, 0x10, 0xeb,
	0xe2, 0x25, 0x50, 0x97, 0xae, 0xac, 0xf2, 0xe4, 0x2d, 0x2a, 0xa5, 0xc1, 0x8d, 0x9b, 0x30, 0x49,
	0x6e, 0xd3, 0xc1, 0x64, 0x26, 0xcc, 0x4c, 0x5a, 0xfd, 0x17, 0xae, 0xfd, 0x45, 0x5d, 0x76, 0xe9,
	0x4a, 0xa4, 0xfd, 0x23, 0xd2, 0x99, 0x5b, 0x68, 0x04, 0xdd, 0x4d, 0xee, 0xf9, 0xce, 0xc9, 0x9d,
	0x93, 0x38, 0x9e, 0xe0, 0xad, 0x02, 0x11, 0x97, 0xc0, 0x40, 0x52, 0x19, 0x35, 0x82, 0x2b, 0xee,
	0x8e, 0x18, 0xcf, 0x2a, 0x88, 0x8c, 0x16, 0x78, 0x25, 0x2f, 0xb9, 0x16, 0xe2, 0xd3, 0xc9, 0x30,
	0xc1, 0x33, 0x74, 0xd2, 0x2c, 0x4f, 0x57, 0x5c, 0x6c, 0x89, 0x28, 0xd2, 0x1a, 0x14, 0x29, 0x88,
	0x22, 0x88, 0xdc, 0x9c, 0x11, 0x96, 0xae, 0x2a, 0x5a, 0xae, 0x55, 0xda, 0x90, 0xfc, 0x0b, 0x28,
	0x94, 0xaf, 0x51, 0xae, 0x29, 0x3b, 0x8f, 0x1e, 0xe3, 0xa8, 0x21, 0x82, 0xd4, 0xb8, 0x4d, 0xf0,
	0x0a, 0x87, 0xa4, 0xaa, 0xf8, 0x16, 0x8a, 0x54, 0xf2, 0x56, 0xe4, 0x90, 0x16, 0xbc, 0x26, 0x94,
	0xa5, 0x12, 0x58, 0x01, 0xc2, 0xa0, 0xcf, 0x7f, 0xf4, 0x9c, 0xd1, 0x07, 0x73, 0x95, 0x44, 0x11,
	0x05, 0xee, 0xd4, 0x19, 0x98, 0x2c, 0xdf, 0x1e, 0xdb, 0x93, 0xe1, 0xd4, 0x8b, 0x2e, 0xaf, 0x16,
	0x2d, 0xb4, 0x36, 0xbb, 0xda, 0xfd, 0x7a, 0x6a, 0x2d, 0x91, 0x74, 0x23, 0xa7, 0x7f, 0x5a, 0x49,
	0xfa, 0xbd, 0x71, 0x6f, 0x32, 0x9c, 0xba, 0x5d, 0xcb, 0x9c, 0x32, 0x85, 0x06, 0x83, 0xb9, 0x1f,
	0x9d, 0xd1, 0x45, 0x09, 0xd2, 0xbf, 0xd2, 0xb6, 0x17, 0x5d, 0x5b, 0xa2, 0xb8, 0x80, 0xfb, 0xd9,
	0xbb, 0x3b, 0x43, 0xcd, 0xb1, 0x29, 0x4c, 0x1a, 0xd2, 0x2c, 0x47, 0xe5, 0x94, 0x77, 0xfd, 0x77,
	0x63, 0xd2, 0xef, 0xeb, 0xd0, 0x27, 0xdd, 0xd0, 0x7b, 0x76, 0xa7, 0xa9, 0x85, 0x86, 0x30, 0xeb,
	0x11, 0xed, 0x4c, 0xa5, 0xdb, 0x38, 0x37, 0xff, 0xab, 0x4e, 0xfa, 0x03, 0x9d, 0xfd, 0xb2, 0x9b,
	0xfd, 0xd6, 0x58, 0x12, 0xed, 0x78, 0xaf, 0x0d, 0x89, 0xe6, 0xf1, 0x35, 0x01, 0xf9, 0x17, 0x20,
	0x5d, 0xcf, 0xe9, 0xf3, 0x2d, 0x03, 0xe1, 0x3f, 0x18, 0xdb, 0x93, 0x87, 0x4b, 0xf3, 0x30, 0xfb,
	0xb4, 0x3b, 0x84, 0xf6, 0xfe, 0x10, 0xda, 0xbf, 0x0f, 0xa1, 0xfd, 0xfd, 0x18, 0x5a, 0xfb, 0x63,
	0x68, 0xfd, 0x3c, 0x86, 0xd6, 0xe7, 0x37, 0x25, 0x55, 0xeb, 0x36, 0x8b, 0x72, 0x5e, 0xc7, 0x52,
	0x09, 0xc2, 0x4a, 0xa8, 0xf8, 0x06, 0x6e, 0x37, 0xc0, 0x54, 0x2b, 0x40, 0xc6, 0x7a, 0xb3, 0x5b,
	0xfc, 0x0f, 0xbe, 0xc6, 0x78, 0x50, 0xdf, 0x1a, 0x90, 0xd9, 0x40, 0x7f, 0xfa, 0xd7, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xe2, 0xf1, 0xcd, 0x55, 0xcb, 0x02, 0x00, 0x00,
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
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.AllowedSourceDomainSenders) > 0 {
		for iNdEx := len(m.AllowedSourceDomainSenders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllowedSourceDomainSenders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.InFlightPackets) > 0 {
		for iNdEx := len(m.InFlightPackets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InFlightPackets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.IbcForwards) > 0 {
		for iNdEx := len(m.IbcForwards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IbcForwards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Mints) > 0 {
		for iNdEx := len(m.Mints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Mints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
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
	if len(m.Mints) > 0 {
		for _, e := range m.Mints {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.IbcForwards) > 0 {
		for _, e := range m.IbcForwards {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.InFlightPackets) > 0 {
		for _, e := range m.InFlightPackets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AllowedSourceDomainSenders) > 0 {
		for _, e := range m.AllowedSourceDomainSenders {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mints", wireType)
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
			m.Mints = append(m.Mints, Mint{})
			if err := m.Mints[len(m.Mints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcForwards", wireType)
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
			m.IbcForwards = append(m.IbcForwards, StoreIBCForwardMetadata{})
			if err := m.IbcForwards[len(m.IbcForwards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InFlightPackets", wireType)
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
			m.InFlightPackets = append(m.InFlightPackets, InFlightPacket{})
			if err := m.InFlightPackets[len(m.InFlightPackets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedSourceDomainSenders", wireType)
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
			m.AllowedSourceDomainSenders = append(m.AllowedSourceDomainSenders, AllowedSourceDomainSender{})
			if err := m.AllowedSourceDomainSenders[len(m.AllowedSourceDomainSenders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
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
