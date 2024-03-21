// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/authz/v1beta1/event.proto

package authzv1beta1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// EventGrant is emitted on Msg/Grant
//
// Since: cosmos-sdk 0.43
type EventGrant struct {
	// Msg type URL for which an authorization is granted
	MsgTypeUrl string `protobuf:"bytes,2,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
	// Granter account address
	Granter string `protobuf:"bytes,3,opt,name=granter,proto3" json:"granter,omitempty"`
	// Grantee account address
	Grantee string `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
}

func (m *EventGrant) Reset()         { *m = EventGrant{} }
func (m *EventGrant) String() string { return proto.CompactTextString(m) }
func (*EventGrant) ProtoMessage()    {}
func (*EventGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f88cbc71a8baf1f, []int{0}
}
func (m *EventGrant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventGrant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventGrant.Merge(m, src)
}
func (m *EventGrant) XXX_Size() int {
	return m.Size()
}
func (m *EventGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_EventGrant.DiscardUnknown(m)
}

var xxx_messageInfo_EventGrant proto.InternalMessageInfo

func (m *EventGrant) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func (m *EventGrant) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *EventGrant) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

// EventRevoke is emitted on Msg/Revoke
//
// Since: cosmos-sdk 0.43
type EventRevoke struct {
	// Msg type URL for which an authorization is revoked
	MsgTypeUrl string `protobuf:"bytes,2,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
	// Granter account address
	Granter string `protobuf:"bytes,3,opt,name=granter,proto3" json:"granter,omitempty"`
	// Grantee account address
	Grantee string `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
}

func (m *EventRevoke) Reset()         { *m = EventRevoke{} }
func (m *EventRevoke) String() string { return proto.CompactTextString(m) }
func (*EventRevoke) ProtoMessage()    {}
func (*EventRevoke) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f88cbc71a8baf1f, []int{1}
}
func (m *EventRevoke) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRevoke) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRevoke.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRevoke) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRevoke.Merge(m, src)
}
func (m *EventRevoke) XXX_Size() int {
	return m.Size()
}
func (m *EventRevoke) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRevoke.DiscardUnknown(m)
}

var xxx_messageInfo_EventRevoke proto.InternalMessageInfo

func (m *EventRevoke) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func (m *EventRevoke) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *EventRevoke) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

// EventPruneExpiredGrants is emitted on Msg/PruneExpiredGrants
//
// Since: x/authz 1.0.0
type EventPruneExpiredGrants struct {
	// Address of the pruner
	Pruner string `protobuf:"bytes,2,opt,name=pruner,proto3" json:"pruner,omitempty"`
}

func (m *EventPruneExpiredGrants) Reset()         { *m = EventPruneExpiredGrants{} }
func (m *EventPruneExpiredGrants) String() string { return proto.CompactTextString(m) }
func (*EventPruneExpiredGrants) ProtoMessage()    {}
func (*EventPruneExpiredGrants) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f88cbc71a8baf1f, []int{2}
}
func (m *EventPruneExpiredGrants) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPruneExpiredGrants) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPruneExpiredGrants.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPruneExpiredGrants) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPruneExpiredGrants.Merge(m, src)
}
func (m *EventPruneExpiredGrants) XXX_Size() int {
	return m.Size()
}
func (m *EventPruneExpiredGrants) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPruneExpiredGrants.DiscardUnknown(m)
}

var xxx_messageInfo_EventPruneExpiredGrants proto.InternalMessageInfo

func (m *EventPruneExpiredGrants) GetPruner() string {
	if m != nil {
		return m.Pruner
	}
	return ""
}

func init() {
	proto.RegisterType((*EventGrant)(nil), "cosmos.authz.v1beta1.EventGrant")
	proto.RegisterType((*EventRevoke)(nil), "cosmos.authz.v1beta1.EventRevoke")
	proto.RegisterType((*EventPruneExpiredGrants)(nil), "cosmos.authz.v1beta1.EventPruneExpiredGrants")
}

func init() { proto.RegisterFile("cosmos/authz/v1beta1/event.proto", fileDescriptor_1f88cbc71a8baf1f) }

var fileDescriptor_1f88cbc71a8baf1f = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x2c, 0x2d, 0xc9, 0xa8, 0xd2, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34,
	0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x81, 0xa8,
	0xd0, 0x03, 0xab, 0xd0, 0x83, 0xaa, 0x90, 0x92, 0x84, 0x88, 0xc6, 0x83, 0xd5, 0xe8, 0x43, 0x95,
	0x80, 0x39, 0x4a, 0xd3, 0x18, 0xb9, 0xb8, 0x5c, 0x41, 0x06, 0xb8, 0x17, 0x25, 0xe6, 0x95, 0x08,
	0x29, 0x70, 0xf1, 0xe4, 0x16, 0xa7, 0xc7, 0x97, 0x54, 0x16, 0xa4, 0xc6, 0x97, 0x16, 0xe5, 0x48,
	0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x71, 0xe5, 0x16, 0xa7, 0x87, 0x54, 0x16, 0xa4, 0x86, 0x16,
	0xe5, 0x08, 0x19, 0x71, 0xb1, 0xa7, 0x83, 0x94, 0xa6, 0x16, 0x49, 0x30, 0x83, 0x24, 0x9d, 0x24,
	0x2e, 0x6d, 0xd1, 0x85, 0x59, 0xeb, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x1c, 0x5c, 0x52, 0x94,
	0x99, 0x97, 0x1e, 0x04, 0x53, 0x88, 0xd0, 0x93, 0x2a, 0xc1, 0x42, 0x9c, 0x9e, 0x54, 0xa5, 0xe9,
	0x8c, 0x5c, 0xdc, 0x60, 0x87, 0x05, 0xa5, 0x96, 0xe5, 0x67, 0xa7, 0x0e, 0x22, 0x97, 0x79, 0x73,
	0x89, 0x83, 0x1d, 0x16, 0x50, 0x54, 0x9a, 0x97, 0xea, 0x5a, 0x51, 0x90, 0x59, 0x94, 0x9a, 0x02,
	0x0e, 0xbd, 0x62, 0x21, 0x03, 0x2e, 0xb6, 0x02, 0x90, 0x68, 0x11, 0xc4, 0x79, 0x78, 0x4c, 0x83,
	0xaa, 0x73, 0x7a, 0xc2, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x5c, 0x12,
	0xc9, 0xf9, 0xb9, 0x7a, 0xd8, 0xe2, 0xd3, 0x89, 0x0b, 0x6a, 0x7f, 0x7e, 0x49, 0x7e, 0x00, 0x63,
	0x94, 0x11, 0x44, 0x4d, 0x71, 0x4a, 0xb6, 0x5e, 0x66, 0xbe, 0x7e, 0x62, 0x41, 0xa6, 0x3e, 0xb6,
	0x64, 0x62, 0x0d, 0xe6, 0x41, 0x39, 0x8b, 0x98, 0x98, 0x9d, 0x1d, 0x23, 0x56, 0x31, 0x89, 0x38,
	0x43, 0xdd, 0x06, 0x36, 0x3e, 0x0c, 0x22, 0x79, 0x0a, 0x26, 0x1c, 0x03, 0x16, 0x8e, 0x81, 0x0a,
	0x3f, 0x62, 0x52, 0xc0, 0x26, 0x1c, 0xe3, 0x1e, 0xe0, 0xe4, 0x9b, 0x5a, 0x92, 0x98, 0x92, 0x58,
	0x92, 0xf8, 0x8a, 0x49, 0x0c, 0xa2, 0xc4, 0xca, 0x0a, 0xac, 0xc6, 0xca, 0x0a, 0xaa, 0x28, 0x89,
	0x0d, 0x9c, 0xda, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xa0, 0x99, 0x14, 0xc2, 0x02,
	0x00, 0x00,
}

func (m *EventGrant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventGrant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventGrant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *EventRevoke) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRevoke) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRevoke) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *EventPruneExpiredGrants) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPruneExpiredGrants) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPruneExpiredGrants) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pruner) > 0 {
		i -= len(m.Pruner)
		copy(dAtA[i:], m.Pruner)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Pruner)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventGrant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventRevoke) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventPruneExpiredGrants) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pruner)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventGrant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventGrant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventGrant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventRevoke) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventRevoke: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRevoke: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventPruneExpiredGrants) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventPruneExpiredGrants: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPruneExpiredGrants: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pruner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pruner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)