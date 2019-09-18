// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/raft/store/snapshot/snapshot.proto

package snapshot

import (
	fmt "fmt"
	github_com_atomix_atomix_raft_node_pkg_atomix_raft_protocol "github.com/atomix/atomix-raft-node/pkg/atomix/raft/protocol"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Snapshot descriptor
type SnapshotDescriptor struct {
	Index     github_com_atomix_atomix_raft_node_pkg_atomix_raft_protocol.Index `protobuf:"varint,1,opt,name=index,proto3,casttype=github.com/atomix/atomix-raft-node/pkg/atomix/raft/protocol.Index" json:"index,omitempty"`
	Timestamp *time.Time                                                        `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp,omitempty"`
}

func (m *SnapshotDescriptor) Reset()         { *m = SnapshotDescriptor{} }
func (m *SnapshotDescriptor) String() string { return proto.CompactTextString(m) }
func (*SnapshotDescriptor) ProtoMessage()    {}
func (*SnapshotDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4596120fca830b6, []int{0}
}
func (m *SnapshotDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotDescriptor.Merge(m, src)
}
func (m *SnapshotDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotDescriptor proto.InternalMessageInfo

func (m *SnapshotDescriptor) GetIndex() github_com_atomix_atomix_raft_node_pkg_atomix_raft_protocol.Index {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *SnapshotDescriptor) GetTimestamp() *time.Time {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*SnapshotDescriptor)(nil), "atomix.raft.SnapshotDescriptor")
}

func init() {
	proto.RegisterFile("atomix/raft/store/snapshot/snapshot.proto", fileDescriptor_c4596120fca830b6)
}

var fileDescriptor_c4596120fca830b6 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0x2c, 0xc9, 0xcf,
	0xcd, 0xac, 0xd0, 0x2f, 0x4a, 0x4c, 0x2b, 0xd1, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x2f, 0xce,
	0x4b, 0x2c, 0x28, 0xce, 0xc8, 0x2f, 0x81, 0x33, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb8,
	0x21, 0x4a, 0xf5, 0x40, 0x4a, 0xa5, 0xe4, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0x52,
	0x49, 0xa5, 0x69, 0xfa, 0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x10, 0xd5, 0x52,
	0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa6, 0x3e, 0x88, 0x05, 0x11, 0x55, 0xda, 0xc8, 0xc8, 0x25,
	0x14, 0x0c, 0x35, 0xd6, 0x25, 0xb5, 0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x24, 0xbf, 0x48, 0x28, 0x9a,
	0x8b, 0x35, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd9, 0xc9, 0xf5, 0xd7,
	0x3d, 0x79, 0xc7, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xa8, 0x1b,
	0x21, 0x94, 0x2e, 0xc8, 0x7e, 0xdd, 0xbc, 0xfc, 0x94, 0x54, 0xfd, 0x82, 0xec, 0x74, 0x7d, 0x64,
	0xf7, 0x83, 0xed, 0x48, 0xce, 0xcf, 0xd1, 0xf3, 0x04, 0x19, 0x16, 0x04, 0x31, 0x53, 0xc8, 0x8e,
	0x8b, 0x13, 0xee, 0x38, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x29, 0x3d, 0x88, 0xf3, 0xf5,
	0x60, 0xce, 0xd7, 0x0b, 0x81, 0xa9, 0x70, 0x62, 0x99, 0x70, 0x5f, 0x9e, 0x31, 0x08, 0xa1, 0xc5,
	0x49, 0x66, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0,
	0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x48,
	0x62, 0x03, 0x1b, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x61, 0xc0, 0x38, 0x49, 0x01,
	0x00, 0x00,
}

func (this *SnapshotDescriptor) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SnapshotDescriptor)
	if !ok {
		that2, ok := that.(SnapshotDescriptor)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if that1.Timestamp == nil {
		if this.Timestamp != nil {
			return false
		}
	} else if !this.Timestamp.Equal(*that1.Timestamp) {
		return false
	}
	return true
}
func (m *SnapshotDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.Timestamp):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintSnapshot(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSnapshot(dAtA []byte, offset int, v uint64) int {
	offset -= sovSnapshot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SnapshotDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovSnapshot(uint64(m.Index))
	}
	if m.Timestamp != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.Timestamp)
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}

func sovSnapshot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSnapshot(x uint64) (n int) {
	return sovSnapshot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SnapshotDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: SnapshotDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= github_com_atomix_atomix_raft_node_pkg_atomix_raft_protocol.Index(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSnapshot
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func skipSnapshot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
				return 0, ErrInvalidLengthSnapshot
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSnapshot
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSnapshot
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
				next, err := skipSnapshot(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSnapshot
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
	ErrInvalidLengthSnapshot = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnapshot   = fmt.Errorf("proto: integer overflow")
)
