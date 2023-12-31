// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: microgrid/iot/device.proto

package types

import (
	fmt "fmt"
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

type Device struct {
	GridId  uint64 `protobuf:"varint,1,opt,name=gridId,proto3" json:"gridId,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Lat     int32  `protobuf:"varint,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon     int32  `protobuf:"varint,4,opt,name=lon,proto3" json:"lon,omitempty"`
	Power   uint64 `protobuf:"varint,5,opt,name=power,proto3" json:"power,omitempty"`
	Voltage uint64 `protobuf:"varint,6,opt,name=voltage,proto3" json:"voltage,omitempty"`
	Others  string `protobuf:"bytes,7,opt,name=others,proto3" json:"others,omitempty"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_78dc413bb041d0f0, []int{0}
}
func (m *Device) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Device.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return m.Size()
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetGridId() uint64 {
	if m != nil {
		return m.GridId
	}
	return 0
}

func (m *Device) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Device) GetLat() int32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Device) GetLon() int32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

func (m *Device) GetPower() uint64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *Device) GetVoltage() uint64 {
	if m != nil {
		return m.Voltage
	}
	return 0
}

func (m *Device) GetOthers() string {
	if m != nil {
		return m.Others
	}
	return ""
}

func init() {
	proto.RegisterType((*Device)(nil), "microgrid.iot.Device")
}

func init() { proto.RegisterFile("microgrid/iot/device.proto", fileDescriptor_78dc413bb041d0f0) }

var fileDescriptor_78dc413bb041d0f0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x1b, 0x67, 0xda, 0xc1, 0x80, 0x20, 0x41, 0x24, 0xb8, 0x08, 0xc5, 0x55, 0x57, 0x2d,
	0xe2, 0x0d, 0xc4, 0x8d, 0xb8, 0xeb, 0xd2, 0x5d, 0xda, 0x84, 0x4e, 0xb0, 0xd3, 0x57, 0x5e, 0xdf,
	0x8c, 0x7a, 0x0b, 0x4f, 0xe1, 0x59, 0x5c, 0xce, 0xd2, 0xa5, 0xb4, 0x17, 0x91, 0xc4, 0x56, 0x77,
	0xf9, 0xbe, 0x1f, 0xfe, 0x3c, 0x7e, 0x7e, 0xb5, 0x73, 0x35, 0x42, 0x83, 0xce, 0x14, 0x0e, 0xa8,
	0x30, 0xf6, 0xe0, 0x6a, 0x9b, 0xf7, 0x08, 0x04, 0xe2, 0xec, 0x2f, 0xcb, 0x1d, 0xd0, 0xf5, 0x07,
	0xe3, 0xc9, 0x7d, 0xc8, 0xc5, 0x25, 0x4f, 0xbc, 0x7e, 0x30, 0x92, 0xa5, 0x2c, 0x5b, 0x97, 0x33,
	0x09, 0xc9, 0x37, 0xda, 0x18, 0xb4, 0xc3, 0x20, 0x4f, 0x52, 0x96, 0x9d, 0x96, 0x0b, 0x8a, 0x73,
	0xbe, 0x6a, 0x35, 0xc9, 0x55, 0xca, 0xb2, 0xb8, 0xf4, 0xcf, 0x60, 0xa0, 0x93, 0xeb, 0xd9, 0x40,
	0x27, 0x2e, 0x78, 0xdc, 0xc3, 0x8b, 0x45, 0x19, 0x87, 0xd2, 0x5f, 0xf0, 0x9d, 0x07, 0x68, 0x49,
	0x37, 0x56, 0x26, 0xc1, 0x2f, 0xe8, 0xaf, 0x00, 0xda, 0x5a, 0x1c, 0xe4, 0x26, 0x7c, 0x36, 0xd3,
	0xdd, 0xe3, 0xe7, 0xa8, 0xd8, 0x71, 0x54, 0xec, 0x7b, 0x54, 0xec, 0x7d, 0x52, 0xd1, 0x71, 0x52,
	0xd1, 0xd7, 0xa4, 0xa2, 0xa7, 0x9b, 0xc6, 0xd1, 0x76, 0x5f, 0xe5, 0x35, 0xec, 0x0a, 0x72, 0x9d,
	0xd6, 0xad, 0xd3, 0xcf, 0x95, 0xc6, 0x1e, 0xf6, 0x58, 0xfc, 0x2f, 0xf1, 0x1a, 0xb6, 0xa0, 0xb7,
	0xde, 0x0e, 0x55, 0x12, 0xb6, 0xb8, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x00, 0x06, 0xe1,
	0x29, 0x01, 0x00, 0x00,
}

func (m *Device) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Device) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Device) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Others) > 0 {
		i -= len(m.Others)
		copy(dAtA[i:], m.Others)
		i = encodeVarintDevice(dAtA, i, uint64(len(m.Others)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Voltage != 0 {
		i = encodeVarintDevice(dAtA, i, uint64(m.Voltage))
		i--
		dAtA[i] = 0x30
	}
	if m.Power != 0 {
		i = encodeVarintDevice(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x28
	}
	if m.Lon != 0 {
		i = encodeVarintDevice(dAtA, i, uint64(m.Lon))
		i--
		dAtA[i] = 0x20
	}
	if m.Lat != 0 {
		i = encodeVarintDevice(dAtA, i, uint64(m.Lat))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintDevice(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.GridId != 0 {
		i = encodeVarintDevice(dAtA, i, uint64(m.GridId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDevice(dAtA []byte, offset int, v uint64) int {
	offset -= sovDevice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Device) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GridId != 0 {
		n += 1 + sovDevice(uint64(m.GridId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.Lat != 0 {
		n += 1 + sovDevice(uint64(m.Lat))
	}
	if m.Lon != 0 {
		n += 1 + sovDevice(uint64(m.Lon))
	}
	if m.Power != 0 {
		n += 1 + sovDevice(uint64(m.Power))
	}
	if m.Voltage != 0 {
		n += 1 + sovDevice(uint64(m.Voltage))
	}
	l = len(m.Others)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	return n
}

func sovDevice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDevice(x uint64) (n int) {
	return sovDevice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Device) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDevice
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
			return fmt.Errorf("proto: Device: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Device: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GridId", wireType)
			}
			m.GridId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GridId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
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
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lat", wireType)
			}
			m.Lat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lat |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lon", wireType)
			}
			m.Lon = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lon |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voltage", wireType)
			}
			m.Voltage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Voltage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Others", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
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
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Others = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDevice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDevice
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
func skipDevice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDevice
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
					return 0, ErrIntOverflowDevice
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
					return 0, ErrIntOverflowDevice
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
				return 0, ErrInvalidLengthDevice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDevice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDevice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDevice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDevice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDevice = fmt.Errorf("proto: unexpected end of group")
)
