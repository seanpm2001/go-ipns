// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ipns.proto

package ipns_pb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
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

// Deprecated: use github.com/ipfs/boxo/ipns/pb.IpnsEntry_ValidityType
type IpnsEntry_ValidityType int32

const (
	// setting an EOL says "this record is valid until..."
	//
	// Deprecated: use github.com/ipfs/boxo/ipns/pb.IpnsEntry_EOL
	IpnsEntry_EOL IpnsEntry_ValidityType = 0
)

// Deprecated: use github.com/ipfs/boxo/ipns/pb.IpnsEntry_ValidityType_name
var IpnsEntry_ValidityType_name = map[int32]string{
	0: "EOL",
}

// Deprecated: use github.com/ipfs/boxo/ipns/pb.IpnsEntry_ValidityType_value
var IpnsEntry_ValidityType_value = map[string]int32{
	"EOL": 0,
}

func (x IpnsEntry_ValidityType) Enum() *IpnsEntry_ValidityType {
	p := new(IpnsEntry_ValidityType)
	*p = x
	return p
}

func (x IpnsEntry_ValidityType) String() string {
	return proto.EnumName(IpnsEntry_ValidityType_name, int32(x))
}

func (x *IpnsEntry_ValidityType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IpnsEntry_ValidityType_value, data, "IpnsEntry_ValidityType")
	if err != nil {
		return err
	}
	*x = IpnsEntry_ValidityType(value)
	return nil
}

func (IpnsEntry_ValidityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d5b16fb32bfe8ea, []int{0, 0}
}

// Deprecated: use github.com/ipfs/boxo/ipns/pb.IpnsEntry
type IpnsEntry struct {
	Value        []byte                  `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	SignatureV1  []byte                  `protobuf:"bytes,2,opt,name=signatureV1" json:"signatureV1,omitempty"`
	ValidityType *IpnsEntry_ValidityType `protobuf:"varint,3,opt,name=validityType,enum=ipns.pb.IpnsEntry_ValidityType" json:"validityType,omitempty"`
	Validity     []byte                  `protobuf:"bytes,4,opt,name=validity" json:"validity,omitempty"`
	Sequence     *uint64                 `protobuf:"varint,5,opt,name=sequence" json:"sequence,omitempty"`
	Ttl          *uint64                 `protobuf:"varint,6,opt,name=ttl" json:"ttl,omitempty"`
	// in order for nodes to properly validate a record upon receipt, they need the public
	// key associated with it. For old RSA keys, its easiest if we just send this as part of
	// the record itself. For newer ed25519 keys, the public key can be embedded in the
	// peerID, making this field unnecessary.
	PubKey               []byte   `protobuf:"bytes,7,opt,name=pubKey" json:"pubKey,omitempty"`
	SignatureV2          []byte   `protobuf:"bytes,8,opt,name=signatureV2" json:"signatureV2,omitempty"`
	Data                 []byte   `protobuf:"bytes,9,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpnsEntry) Reset()         { *m = IpnsEntry{} }
func (m *IpnsEntry) String() string { return proto.CompactTextString(m) }
func (*IpnsEntry) ProtoMessage()    {}
func (*IpnsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5b16fb32bfe8ea, []int{0}
}
func (m *IpnsEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IpnsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IpnsEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IpnsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpnsEntry.Merge(m, src)
}
func (m *IpnsEntry) XXX_Size() int {
	return m.Size()
}
func (m *IpnsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_IpnsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_IpnsEntry proto.InternalMessageInfo

func (m *IpnsEntry) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *IpnsEntry) GetSignatureV1() []byte {
	if m != nil {
		return m.SignatureV1
	}
	return nil
}

func (m *IpnsEntry) GetValidityType() IpnsEntry_ValidityType {
	if m != nil && m.ValidityType != nil {
		return *m.ValidityType
	}
	return IpnsEntry_EOL
}

func (m *IpnsEntry) GetValidity() []byte {
	if m != nil {
		return m.Validity
	}
	return nil
}

func (m *IpnsEntry) GetSequence() uint64 {
	if m != nil && m.Sequence != nil {
		return *m.Sequence
	}
	return 0
}

func (m *IpnsEntry) GetTtl() uint64 {
	if m != nil && m.Ttl != nil {
		return *m.Ttl
	}
	return 0
}

func (m *IpnsEntry) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *IpnsEntry) GetSignatureV2() []byte {
	if m != nil {
		return m.SignatureV2
	}
	return nil
}

func (m *IpnsEntry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Deprecated: use github.com/ipfs/boxo/ipns/pb.IpnsSignatureV2Checker
type IpnsSignatureV2Checker struct {
	PubKey               []byte   `protobuf:"bytes,7,opt,name=pubKey" json:"pubKey,omitempty"`
	SignatureV2          []byte   `protobuf:"bytes,8,opt,name=signatureV2" json:"signatureV2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpnsSignatureV2Checker) Reset()         { *m = IpnsSignatureV2Checker{} }
func (m *IpnsSignatureV2Checker) String() string { return proto.CompactTextString(m) }
func (*IpnsSignatureV2Checker) ProtoMessage()    {}
func (*IpnsSignatureV2Checker) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5b16fb32bfe8ea, []int{1}
}
func (m *IpnsSignatureV2Checker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IpnsSignatureV2Checker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IpnsSignatureV2Checker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IpnsSignatureV2Checker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpnsSignatureV2Checker.Merge(m, src)
}
func (m *IpnsSignatureV2Checker) XXX_Size() int {
	return m.Size()
}
func (m *IpnsSignatureV2Checker) XXX_DiscardUnknown() {
	xxx_messageInfo_IpnsSignatureV2Checker.DiscardUnknown(m)
}

var xxx_messageInfo_IpnsSignatureV2Checker proto.InternalMessageInfo

func (m *IpnsSignatureV2Checker) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *IpnsSignatureV2Checker) GetSignatureV2() []byte {
	if m != nil {
		return m.SignatureV2
	}
	return nil
}

func init() {
	proto.RegisterEnum("ipns.pb.IpnsEntry_ValidityType", IpnsEntry_ValidityType_name, IpnsEntry_ValidityType_value)
	proto.RegisterType((*IpnsEntry)(nil), "ipns.pb.IpnsEntry")
	proto.RegisterType((*IpnsSignatureV2Checker)(nil), "ipns.pb.IpnsSignatureV2Checker")
}

func init() { proto.RegisterFile("ipns.proto", fileDescriptor_4d5b16fb32bfe8ea) }

var fileDescriptor_4d5b16fb32bfe8ea = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x2c, 0xc8, 0x2b,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0xb0, 0x93, 0x94, 0x76, 0x32, 0x71, 0x71,
	0x7a, 0x16, 0xe4, 0x15, 0xbb, 0xe6, 0x95, 0x14, 0x55, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6,
	0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x41, 0x38, 0x42, 0x0a, 0x5c, 0xdc, 0xc5,
	0x99, 0xe9, 0x79, 0x89, 0x25, 0xa5, 0x45, 0xa9, 0x61, 0x86, 0x12, 0x4c, 0x60, 0x39, 0x64, 0x21,
	0x21, 0x67, 0x2e, 0x9e, 0xb2, 0xc4, 0x9c, 0xcc, 0x94, 0xcc, 0x92, 0xca, 0x90, 0xca, 0x82, 0x54,
	0x09, 0x66, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x79, 0x3d, 0xa8, 0x2d, 0x7a, 0x70, 0x1b, 0xf4, 0xc2,
	0x90, 0x94, 0x05, 0xa1, 0x68, 0x12, 0x92, 0xe2, 0xe2, 0x80, 0xf1, 0x25, 0x58, 0xc0, 0x76, 0xc0,
	0xf9, 0x20, 0xb9, 0xe2, 0xd4, 0xc2, 0xd2, 0xd4, 0xbc, 0xe4, 0x54, 0x09, 0x56, 0x05, 0x46, 0x0d,
	0x96, 0x20, 0x38, 0x5f, 0x48, 0x80, 0x8b, 0xb9, 0xa4, 0x24, 0x47, 0x82, 0x0d, 0x2c, 0x0c, 0x62,
	0x0a, 0x89, 0x71, 0xb1, 0x15, 0x94, 0x26, 0x79, 0xa7, 0x56, 0x4a, 0xb0, 0x83, 0xcd, 0x81, 0xf2,
	0x50, 0x3d, 0x62, 0x24, 0xc1, 0x81, 0xee, 0x11, 0x23, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92,
	0x44, 0x09, 0x4e, 0xb0, 0x14, 0x98, 0xad, 0x24, 0xce, 0xc5, 0x83, 0xec, 0x6a, 0x21, 0x76, 0x2e,
	0x66, 0x57, 0x7f, 0x1f, 0x01, 0x06, 0xa5, 0x20, 0x2e, 0x31, 0x90, 0xc7, 0x82, 0x11, 0xfa, 0x9d,
	0x33, 0x52, 0x93, 0xb3, 0x53, 0x8b, 0xc8, 0x77, 0x80, 0x93, 0xc4, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x18, 0xc5, 0xa5, 0xa7, 0x6f, 0x0d, 0x0a, 0xc3, 0xf8,
	0x82, 0x24, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0x85, 0x5b, 0xed, 0xbf, 0x01, 0x00, 0x00,
}

func (m *IpnsEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IpnsEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IpnsEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Data != nil {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x4a
	}
	if m.SignatureV2 != nil {
		i -= len(m.SignatureV2)
		copy(dAtA[i:], m.SignatureV2)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.SignatureV2)))
		i--
		dAtA[i] = 0x42
	}
	if m.PubKey != nil {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Ttl != nil {
		i = encodeVarintIpns(dAtA, i, uint64(*m.Ttl))
		i--
		dAtA[i] = 0x30
	}
	if m.Sequence != nil {
		i = encodeVarintIpns(dAtA, i, uint64(*m.Sequence))
		i--
		dAtA[i] = 0x28
	}
	if m.Validity != nil {
		i -= len(m.Validity)
		copy(dAtA[i:], m.Validity)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.Validity)))
		i--
		dAtA[i] = 0x22
	}
	if m.ValidityType != nil {
		i = encodeVarintIpns(dAtA, i, uint64(*m.ValidityType))
		i--
		dAtA[i] = 0x18
	}
	if m.SignatureV1 != nil {
		i -= len(m.SignatureV1)
		copy(dAtA[i:], m.SignatureV1)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.SignatureV1)))
		i--
		dAtA[i] = 0x12
	}
	if m.Value != nil {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IpnsSignatureV2Checker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IpnsSignatureV2Checker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IpnsSignatureV2Checker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SignatureV2 != nil {
		i -= len(m.SignatureV2)
		copy(dAtA[i:], m.SignatureV2)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.SignatureV2)))
		i--
		dAtA[i] = 0x42
	}
	if m.PubKey != nil {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintIpns(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}

func encodeVarintIpns(dAtA []byte, offset int, v uint64) int {
	offset -= sovIpns(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IpnsEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		l = len(m.Value)
		n += 1 + l + sovIpns(uint64(l))
	}
	if m.SignatureV1 != nil {
		l = len(m.SignatureV1)
		n += 1 + l + sovIpns(uint64(l))
	}
	if m.ValidityType != nil {
		n += 1 + sovIpns(uint64(*m.ValidityType))
	}
	if m.Validity != nil {
		l = len(m.Validity)
		n += 1 + l + sovIpns(uint64(l))
	}
	if m.Sequence != nil {
		n += 1 + sovIpns(uint64(*m.Sequence))
	}
	if m.Ttl != nil {
		n += 1 + sovIpns(uint64(*m.Ttl))
	}
	if m.PubKey != nil {
		l = len(m.PubKey)
		n += 1 + l + sovIpns(uint64(l))
	}
	if m.SignatureV2 != nil {
		l = len(m.SignatureV2)
		n += 1 + l + sovIpns(uint64(l))
	}
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovIpns(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IpnsSignatureV2Checker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PubKey != nil {
		l = len(m.PubKey)
		n += 1 + l + sovIpns(uint64(l))
	}
	if m.SignatureV2 != nil {
		l = len(m.SignatureV2)
		n += 1 + l + sovIpns(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIpns(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIpns(x uint64) (n int) {
	return sovIpns(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IpnsEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIpns
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
			return fmt.Errorf("proto: IpnsEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IpnsEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
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
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignatureV1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
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
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignatureV1 = append(m.SignatureV1[:0], dAtA[iNdEx:postIndex]...)
			if m.SignatureV1 == nil {
				m.SignatureV1 = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidityType", wireType)
			}
			var v IpnsEntry_ValidityType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= IpnsEntry_ValidityType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ValidityType = &v
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validity", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
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
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validity = append(m.Validity[:0], dAtA[iNdEx:postIndex]...)
			if m.Validity == nil {
				m.Validity = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sequence = &v
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ttl = &v
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
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
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignatureV2", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
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
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignatureV2 = append(m.SignatureV2[:0], dAtA[iNdEx:postIndex]...)
			if m.SignatureV2 == nil {
				m.SignatureV2 = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
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
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIpns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIpns
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
func (m *IpnsSignatureV2Checker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIpns
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
			return fmt.Errorf("proto: IpnsSignatureV2Checker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IpnsSignatureV2Checker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
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
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignatureV2", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpns
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
				return ErrInvalidLengthIpns
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignatureV2 = append(m.SignatureV2[:0], dAtA[iNdEx:postIndex]...)
			if m.SignatureV2 == nil {
				m.SignatureV2 = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIpns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIpns
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
func skipIpns(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIpns
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
					return 0, ErrIntOverflowIpns
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
					return 0, ErrIntOverflowIpns
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
				return 0, ErrInvalidLengthIpns
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIpns
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIpns
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	// Deprecated: use github.com/ipfs/boxo/ipns/pb.ErrInvalidLengthIpns
	ErrInvalidLengthIpns = fmt.Errorf("proto: negative length found during unmarshaling")
	// Deprecated: use github.com/ipfs/boxo/ipns/pb.ErrIntOverflowIpns
	ErrIntOverflowIpns = fmt.Errorf("proto: integer overflow")
	// Deprecated: use github.com/ipfs/boxo/ipns/pb.ErrUnexpectedEndOfGroupIpns
	ErrUnexpectedEndOfGroupIpns = fmt.Errorf("proto: unexpected end of group")
)
