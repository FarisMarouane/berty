// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entity/device_push_identifier.proto

package entity

import (
	fmt "fmt"
	io "io"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DevicePushIdentifier struct {
	ID                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	CreatedAt            time.Time `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	UpdatedAt            time.Time `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at"`
	PushInfo             []byte    `protobuf:"bytes,4,opt,name=push_info,json=pushInfo,proto3" json:"push_info,omitempty"`
	RelayPubkey          string    `protobuf:"bytes,5,opt,name=relay_pubkey,json=relayPubkey,proto3" json:"relay_pubkey,omitempty"`
	DeviceID             string    `protobuf:"bytes,6,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DevicePushIdentifier) Reset()         { *m = DevicePushIdentifier{} }
func (m *DevicePushIdentifier) String() string { return proto.CompactTextString(m) }
func (*DevicePushIdentifier) ProtoMessage()    {}
func (*DevicePushIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_76aa55265fd1e4d6, []int{0}
}
func (m *DevicePushIdentifier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DevicePushIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DevicePushIdentifier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DevicePushIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevicePushIdentifier.Merge(m, src)
}
func (m *DevicePushIdentifier) XXX_Size() int {
	return m.Size()
}
func (m *DevicePushIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_DevicePushIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_DevicePushIdentifier proto.InternalMessageInfo

func (m *DevicePushIdentifier) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *DevicePushIdentifier) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *DevicePushIdentifier) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

func (m *DevicePushIdentifier) GetPushInfo() []byte {
	if m != nil {
		return m.PushInfo
	}
	return nil
}

func (m *DevicePushIdentifier) GetRelayPubkey() string {
	if m != nil {
		return m.RelayPubkey
	}
	return ""
}

func (m *DevicePushIdentifier) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func init() {
	proto.RegisterType((*DevicePushIdentifier)(nil), "berty.entity.DevicePushIdentifier")
}

func init() {
	proto.RegisterFile("entity/device_push_identifier.proto", fileDescriptor_76aa55265fd1e4d6)
}

var fileDescriptor_76aa55265fd1e4d6 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x4d, 0x4e, 0x83, 0x40,
	0x14, 0xee, 0xa0, 0x36, 0x30, 0x25, 0x31, 0x21, 0x8d, 0x21, 0xd5, 0x40, 0x53, 0x37, 0x98, 0x18,
	0x48, 0xea, 0xce, 0x5d, 0x6b, 0x37, 0xec, 0x1a, 0xe2, 0xca, 0x0d, 0x01, 0xe6, 0x41, 0x27, 0x2d,
	0x1d, 0x32, 0x1d, 0x4c, 0xb8, 0x85, 0x4b, 0x8f, 0xd4, 0xb8, 0xf2, 0x04, 0x68, 0xf0, 0x06, 0x9e,
	0xc0, 0xc0, 0xb4, 0xba, 0x76, 0x37, 0xf3, 0xbd, 0xef, 0x27, 0xef, 0x7d, 0xf8, 0x1a, 0xb6, 0x82,
	0x8a, 0xca, 0x23, 0xf0, 0x4c, 0x13, 0x08, 0x8b, 0x72, 0xb7, 0x0a, 0x29, 0x69, 0xc1, 0x94, 0x02,
	0x77, 0x0b, 0xce, 0x04, 0x33, 0xf4, 0x18, 0xb8, 0xa8, 0x5c, 0x49, 0x1d, 0x0d, 0x33, 0x96, 0xb1,
	0x6e, 0xe0, 0xb5, 0x2f, 0xc9, 0x19, 0xd9, 0x19, 0x63, 0xd9, 0x06, 0xbc, 0xee, 0x17, 0x97, 0xa9,
	0x27, 0x68, 0x0e, 0x3b, 0x11, 0xe5, 0x85, 0x24, 0x4c, 0xde, 0x14, 0x3c, 0x5c, 0x74, 0x29, 0xcb,
	0x72, 0xb7, 0xf2, 0x7f, 0x33, 0x8c, 0x5b, 0xac, 0x50, 0x62, 0xa2, 0x31, 0x72, 0xb4, 0xf9, 0x55,
	0x53, 0xdb, 0x8a, 0xbf, 0xf8, 0xae, 0x6d, 0x23, 0x63, 0x3c, 0xbf, 0x9f, 0x14, 0x9c, 0xe6, 0x11,
	0xaf, 0xc2, 0x35, 0x54, 0x93, 0x40, 0xa1, 0xc4, 0x78, 0xc0, 0x38, 0xe1, 0x10, 0x09, 0x20, 0x61,
	0x24, 0x4c, 0x65, 0x8c, 0x9c, 0xc1, 0x74, 0xe4, 0xca, 0x70, 0xf7, 0x18, 0xee, 0x3e, 0x1e, 0xc3,
	0xe7, 0xea, 0xbe, 0xb6, 0x7b, 0x2f, 0x1f, 0x36, 0x0a, 0xb4, 0x83, 0x6e, 0x26, 0x5a, 0x93, 0xb2,
	0x20, 0x47, 0x93, 0x93, 0xff, 0x98, 0x1c, 0x74, 0x33, 0x61, 0x5c, 0x62, 0x4d, 0x9e, 0x6b, 0x9b,
	0x32, 0xf3, 0x74, 0x8c, 0x1c, 0x3d, 0x50, 0x5b, 0xc0, 0xdf, 0xa6, 0xcc, 0x98, 0x62, 0x9d, 0xc3,
	0x26, 0xaa, 0xc2, 0xa2, 0x8c, 0xd7, 0x50, 0x99, 0x67, 0xdd, 0x7a, 0xe7, 0x4d, 0x6d, 0x0f, 0x82,
	0x16, 0x5f, 0x76, 0x70, 0x30, 0xe0, 0x7f, 0x1f, 0xe3, 0x06, 0x6b, 0x87, 0x1a, 0x28, 0x31, 0xfb,
	0x9d, 0x40, 0x6f, 0x6a, 0x5b, 0x95, 0x57, 0xf3, 0x17, 0x81, 0x2a, 0xc7, 0x3e, 0x99, 0x3b, 0xfb,
	0xc6, 0x42, 0xef, 0x8d, 0x85, 0x3e, 0x1b, 0x0b, 0xbd, 0x7e, 0x59, 0xbd, 0xa7, 0x0b, 0xd9, 0x91,
	0x80, 0x64, 0xe5, 0x25, 0x8c, 0x83, 0x27, 0xdb, 0x8a, 0xfb, 0xdd, 0x3a, 0x77, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd5, 0x02, 0x71, 0x5f, 0xe9, 0x01, 0x00, 0x00,
}

func (m *DevicePushIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevicePushIdentifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDevicePushIdentifier(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintDevicePushIdentifier(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)))
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err1 != nil {
		return 0, err1
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintDevicePushIdentifier(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)))
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err2 != nil {
		return 0, err2
	}
	i += n2
	if len(m.PushInfo) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintDevicePushIdentifier(dAtA, i, uint64(len(m.PushInfo)))
		i += copy(dAtA[i:], m.PushInfo)
	}
	if len(m.RelayPubkey) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintDevicePushIdentifier(dAtA, i, uint64(len(m.RelayPubkey)))
		i += copy(dAtA[i:], m.RelayPubkey)
	}
	if len(m.DeviceID) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintDevicePushIdentifier(dAtA, i, uint64(len(m.DeviceID)))
		i += copy(dAtA[i:], m.DeviceID)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDevicePushIdentifier(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DevicePushIdentifier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovDevicePushIdentifier(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovDevicePushIdentifier(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovDevicePushIdentifier(uint64(l))
	l = len(m.PushInfo)
	if l > 0 {
		n += 1 + l + sovDevicePushIdentifier(uint64(l))
	}
	l = len(m.RelayPubkey)
	if l > 0 {
		n += 1 + l + sovDevicePushIdentifier(uint64(l))
	}
	l = len(m.DeviceID)
	if l > 0 {
		n += 1 + l + sovDevicePushIdentifier(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDevicePushIdentifier(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDevicePushIdentifier(x uint64) (n int) {
	return sovDevicePushIdentifier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DevicePushIdentifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDevicePushIdentifier
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
			return fmt.Errorf("proto: DevicePushIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevicePushIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushIdentifier
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
				return ErrInvalidLengthDevicePushIdentifier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushIdentifier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushIdentifier
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
				return ErrInvalidLengthDevicePushIdentifier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushIdentifier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushIdentifier
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
				return ErrInvalidLengthDevicePushIdentifier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushIdentifier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushInfo", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushIdentifier
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
				return ErrInvalidLengthDevicePushIdentifier
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushIdentifier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PushInfo = append(m.PushInfo[:0], dAtA[iNdEx:postIndex]...)
			if m.PushInfo == nil {
				m.PushInfo = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayPubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushIdentifier
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
				return ErrInvalidLengthDevicePushIdentifier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushIdentifier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelayPubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushIdentifier
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
				return ErrInvalidLengthDevicePushIdentifier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushIdentifier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDevicePushIdentifier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDevicePushIdentifier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDevicePushIdentifier
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
func skipDevicePushIdentifier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDevicePushIdentifier
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
					return 0, ErrIntOverflowDevicePushIdentifier
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDevicePushIdentifier
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
				return 0, ErrInvalidLengthDevicePushIdentifier
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDevicePushIdentifier
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDevicePushIdentifier
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDevicePushIdentifier(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDevicePushIdentifier
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDevicePushIdentifier = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDevicePushIdentifier   = fmt.Errorf("proto: integer overflow")
)
