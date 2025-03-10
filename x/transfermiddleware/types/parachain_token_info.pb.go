// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: centauri/transfermiddleware/v1beta1/parachain_token_info.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ParachainIBCTokenInfo represents information about transferable IBC tokens
// from Parachain.
type ParachainIBCTokenInfo struct {
	// ibc_denom is the denomination of the ibced token transferred from the
	// dotsama chain.
	IbcDenom string `protobuf:"bytes,1,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty" yaml:"ibc_denom"`
	// channel_id is source channel in IBC connection from centauri chain
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
	// native denom is new native minted denom in centauri chain.
	NativeDenom string `protobuf:"bytes,3,opt,name=native_denom,json=nativeDenom,proto3" json:"native_denom,omitempty" yaml:"native_denom"`
	// asset id is the id of the asset on Picasso
	AssetId string `protobuf:"bytes,4,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
}

func (m *ParachainIBCTokenInfo) Reset()         { *m = ParachainIBCTokenInfo{} }
func (m *ParachainIBCTokenInfo) String() string { return proto.CompactTextString(m) }
func (*ParachainIBCTokenInfo) ProtoMessage()    {}
func (*ParachainIBCTokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9cc8ea0d211210f, []int{0}
}
func (m *ParachainIBCTokenInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParachainIBCTokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParachainIBCTokenInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParachainIBCTokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParachainIBCTokenInfo.Merge(m, src)
}
func (m *ParachainIBCTokenInfo) XXX_Size() int {
	return m.Size()
}
func (m *ParachainIBCTokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ParachainIBCTokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ParachainIBCTokenInfo proto.InternalMessageInfo

func (m *ParachainIBCTokenInfo) GetIbcDenom() string {
	if m != nil {
		return m.IbcDenom
	}
	return ""
}

func (m *ParachainIBCTokenInfo) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ParachainIBCTokenInfo) GetNativeDenom() string {
	if m != nil {
		return m.NativeDenom
	}
	return ""
}

func (m *ParachainIBCTokenInfo) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

type RemoveParachainIBCTokenInfo struct {
	// native denom is new native minted denom in centauri chain.
	NativeDenom string `protobuf:"bytes,1,opt,name=native_denom,json=nativeDenom,proto3" json:"native_denom,omitempty" yaml:"native_denom"`
	// remove_time is the time at which the parachain token info will be removed.
	RemoveTime time.Time `protobuf:"bytes,2,opt,name=remove_time,json=removeTime,proto3,stdtime" json:"remove_time" yaml:"start_time"`
}

func (m *RemoveParachainIBCTokenInfo) Reset()         { *m = RemoveParachainIBCTokenInfo{} }
func (m *RemoveParachainIBCTokenInfo) String() string { return proto.CompactTextString(m) }
func (*RemoveParachainIBCTokenInfo) ProtoMessage()    {}
func (*RemoveParachainIBCTokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9cc8ea0d211210f, []int{1}
}
func (m *RemoveParachainIBCTokenInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveParachainIBCTokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveParachainIBCTokenInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveParachainIBCTokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveParachainIBCTokenInfo.Merge(m, src)
}
func (m *RemoveParachainIBCTokenInfo) XXX_Size() int {
	return m.Size()
}
func (m *RemoveParachainIBCTokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveParachainIBCTokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveParachainIBCTokenInfo proto.InternalMessageInfo

func (m *RemoveParachainIBCTokenInfo) GetNativeDenom() string {
	if m != nil {
		return m.NativeDenom
	}
	return ""
}

func (m *RemoveParachainIBCTokenInfo) GetRemoveTime() time.Time {
	if m != nil {
		return m.RemoveTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*ParachainIBCTokenInfo)(nil), "centauri.transfermiddleware.v1beta1.ParachainIBCTokenInfo")
	proto.RegisterType((*RemoveParachainIBCTokenInfo)(nil), "centauri.transfermiddleware.v1beta1.RemoveParachainIBCTokenInfo")
}

func init() {
	proto.RegisterFile("centauri/transfermiddleware/v1beta1/parachain_token_info.proto", fileDescriptor_b9cc8ea0d211210f)
}

var fileDescriptor_b9cc8ea0d211210f = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0xc6, 0x63, 0x40, 0x70, 0xeb, 0x22, 0x01, 0xb9, 0xad, 0xa8, 0x82, 0x48, 0x90, 0x59, 0x98,
	0x12, 0x15, 0x3a, 0x75, 0x60, 0x08, 0x2c, 0xd9, 0x50, 0xd4, 0xa9, 0x4b, 0xe4, 0x24, 0x4e, 0x6a,
	0x91, 0xd8, 0x91, 0xe3, 0x16, 0xfa, 0x16, 0x7d, 0x19, 0xde, 0xa1, 0x63, 0x47, 0xa6, 0x80, 0xda,
	0x95, 0x29, 0x4f, 0x80, 0x62, 0x37, 0x14, 0x41, 0x97, 0xbb, 0x9d, 0xa3, 0xef, 0xfc, 0xfc, 0x9d,
	0x3f, 0x86, 0xef, 0x13, 0xc2, 0x24, 0x5e, 0x0b, 0xea, 0x49, 0x81, 0x59, 0x9d, 0x11, 0x51, 0xd2,
	0x34, 0x2d, 0xc8, 0x17, 0x2c, 0x88, 0xb7, 0x99, 0xc6, 0x44, 0xe2, 0xa9, 0x57, 0x61, 0x81, 0x93,
	0x15, 0xa6, 0x2c, 0x92, 0xfc, 0x33, 0x61, 0x11, 0x65, 0x19, 0x77, 0x2b, 0xc1, 0x25, 0x37, 0x5f,
	0xf7, 0xbc, 0xfb, 0x3f, 0xef, 0x9e, 0x79, 0x6b, 0x94, 0xf3, 0x9c, 0xab, 0x7a, 0xaf, 0x8b, 0x34,
	0x6a, 0x39, 0x39, 0xe7, 0x79, 0x41, 0x3c, 0x95, 0xc5, 0xeb, 0xcc, 0x93, 0xb4, 0x24, 0xb5, 0xc4,
	0x65, 0xa5, 0x0b, 0xd0, 0x2f, 0x00, 0xc7, 0x9f, 0x7a, 0xeb, 0xc0, 0xff, 0xb0, 0xe8, 0xcc, 0x03,
	0x96, 0x71, 0x73, 0x0a, 0x07, 0x34, 0x4e, 0xa2, 0x94, 0x30, 0x5e, 0x4e, 0xc0, 0x2b, 0xf0, 0x66,
	0xe0, 0x8f, 0xda, 0xc6, 0x79, 0xba, 0xc5, 0x65, 0x31, 0x47, 0x7f, 0x24, 0x14, 0xde, 0xd0, 0x38,
	0xf9, 0xd8, 0x85, 0xe6, 0x0c, 0xc2, 0x64, 0x85, 0x19, 0x23, 0x45, 0x44, 0xd3, 0xc9, 0x3d, 0xc5,
	0x8c, 0xdb, 0xc6, 0x79, 0xa6, 0x99, 0x8b, 0x86, 0xc2, 0xc1, 0x39, 0x09, 0x52, 0x73, 0x0e, 0x1f,
	0x33, 0x2c, 0xe9, 0x86, 0x9c, 0xbd, 0xee, 0x2b, 0xee, 0x79, 0xdb, 0x38, 0xb7, 0x9a, 0xfb, 0x5b,
	0x45, 0xe1, 0x50, 0xa7, 0xda, 0xd1, 0x85, 0x37, 0xb8, 0xae, 0x89, 0xec, 0xfc, 0x1e, 0x28, 0xee,
	0xb6, 0x6d, 0x9c, 0x27, 0x9a, 0xeb, 0x15, 0x14, 0x3e, 0x52, 0x61, 0x90, 0xa2, 0x6f, 0x00, 0xbe,
	0x08, 0x49, 0xc9, 0x37, 0xe4, 0xfa, 0xd0, 0xff, 0xf6, 0x02, 0xee, 0xd0, 0xcb, 0x12, 0x0e, 0x85,
	0x7a, 0x3a, 0xea, 0x96, 0xac, 0xc6, 0x1f, 0xbe, 0xb5, 0x5c, 0x7d, 0x01, 0xb7, 0xbf, 0x80, 0xbb,
	0xe8, 0x2f, 0xe0, 0xbf, 0xdc, 0x37, 0x8e, 0x71, 0x59, 0x4f, 0x2d, 0xb1, 0x90, 0x8a, 0x45, 0xbb,
	0x1f, 0x0e, 0x08, 0xa1, 0x7e, 0xad, 0xab, 0xf7, 0x67, 0xfb, 0xa3, 0x0d, 0x0e, 0x47, 0x1b, 0xfc,
	0x3c, 0xda, 0x60, 0x77, 0xb2, 0x8d, 0xc3, 0xc9, 0x36, 0xbe, 0x9f, 0x6c, 0x63, 0x69, 0x7d, 0xbd,
	0xf6, 0xab, 0xe4, 0xb6, 0x22, 0x75, 0xfc, 0x50, 0x99, 0xbe, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0xd7, 0x2a, 0x55, 0xb9, 0x81, 0x02, 0x00, 0x00,
}

func (m *ParachainIBCTokenInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParachainIBCTokenInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParachainIBCTokenInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetId) > 0 {
		i -= len(m.AssetId)
		copy(dAtA[i:], m.AssetId)
		i = encodeVarintParachainTokenInfo(dAtA, i, uint64(len(m.AssetId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NativeDenom) > 0 {
		i -= len(m.NativeDenom)
		copy(dAtA[i:], m.NativeDenom)
		i = encodeVarintParachainTokenInfo(dAtA, i, uint64(len(m.NativeDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintParachainTokenInfo(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintParachainTokenInfo(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RemoveParachainIBCTokenInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveParachainIBCTokenInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveParachainIBCTokenInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.RemoveTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.RemoveTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParachainTokenInfo(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.NativeDenom) > 0 {
		i -= len(m.NativeDenom)
		copy(dAtA[i:], m.NativeDenom)
		i = encodeVarintParachainTokenInfo(dAtA, i, uint64(len(m.NativeDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParachainTokenInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovParachainTokenInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ParachainIBCTokenInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovParachainTokenInfo(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovParachainTokenInfo(uint64(l))
	}
	l = len(m.NativeDenom)
	if l > 0 {
		n += 1 + l + sovParachainTokenInfo(uint64(l))
	}
	l = len(m.AssetId)
	if l > 0 {
		n += 1 + l + sovParachainTokenInfo(uint64(l))
	}
	return n
}

func (m *RemoveParachainIBCTokenInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NativeDenom)
	if l > 0 {
		n += 1 + l + sovParachainTokenInfo(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.RemoveTime)
	n += 1 + l + sovParachainTokenInfo(uint64(l))
	return n
}

func sovParachainTokenInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParachainTokenInfo(x uint64) (n int) {
	return sovParachainTokenInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ParachainIBCTokenInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParachainTokenInfo
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
			return fmt.Errorf("proto: ParachainIBCTokenInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParachainIBCTokenInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParachainTokenInfo
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
				return ErrInvalidLengthParachainTokenInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParachainTokenInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParachainTokenInfo
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
				return ErrInvalidLengthParachainTokenInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParachainTokenInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParachainTokenInfo
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
				return ErrInvalidLengthParachainTokenInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParachainTokenInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NativeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParachainTokenInfo
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
				return ErrInvalidLengthParachainTokenInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParachainTokenInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParachainTokenInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParachainTokenInfo
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
func (m *RemoveParachainIBCTokenInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParachainTokenInfo
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
			return fmt.Errorf("proto: RemoveParachainIBCTokenInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveParachainIBCTokenInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParachainTokenInfo
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
				return ErrInvalidLengthParachainTokenInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParachainTokenInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NativeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoveTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParachainTokenInfo
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
				return ErrInvalidLengthParachainTokenInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParachainTokenInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.RemoveTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParachainTokenInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParachainTokenInfo
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
func skipParachainTokenInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParachainTokenInfo
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
					return 0, ErrIntOverflowParachainTokenInfo
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
					return 0, ErrIntOverflowParachainTokenInfo
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
				return 0, ErrInvalidLengthParachainTokenInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParachainTokenInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParachainTokenInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParachainTokenInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParachainTokenInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParachainTokenInfo = fmt.Errorf("proto: unexpected end of group")
)
