// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/interchain_accounts/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Query request for an interchain account address
type QueryInterchainAccountAddressRequest struct {
	// Counterparty PortID is the portID on the controller chain
	CounterpartyPortId string `protobuf:"bytes,1,opt,name=counterparty_port_id,json=counterpartyPortId,proto3" json:"counterparty_port_id,omitempty"`
}

func (m *QueryInterchainAccountAddressRequest) Reset()         { *m = QueryInterchainAccountAddressRequest{} }
func (m *QueryInterchainAccountAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryInterchainAccountAddressRequest) ProtoMessage()    {}
func (*QueryInterchainAccountAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72a16b57c3343764, []int{0}
}
func (m *QueryInterchainAccountAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInterchainAccountAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInterchainAccountAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInterchainAccountAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInterchainAccountAddressRequest.Merge(m, src)
}
func (m *QueryInterchainAccountAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryInterchainAccountAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInterchainAccountAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInterchainAccountAddressRequest proto.InternalMessageInfo

// Query response for an interchain account address
type QueryInterchainAccountAddressResponse struct {
	// The corresponding interchain account address on the host chain
	InterchainAccountAddress string `protobuf:"bytes,1,opt,name=interchain_account_address,json=interchainAccountAddress,proto3" json:"interchain_account_address,omitempty"`
}

func (m *QueryInterchainAccountAddressResponse) Reset()         { *m = QueryInterchainAccountAddressResponse{} }
func (m *QueryInterchainAccountAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryInterchainAccountAddressResponse) ProtoMessage()    {}
func (*QueryInterchainAccountAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72a16b57c3343764, []int{1}
}
func (m *QueryInterchainAccountAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInterchainAccountAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInterchainAccountAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInterchainAccountAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInterchainAccountAddressResponse.Merge(m, src)
}
func (m *QueryInterchainAccountAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryInterchainAccountAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInterchainAccountAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInterchainAccountAddressResponse proto.InternalMessageInfo

func (m *QueryInterchainAccountAddressResponse) GetInterchainAccountAddress() string {
	if m != nil {
		return m.InterchainAccountAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryInterchainAccountAddressRequest)(nil), "ibc.applications.interchain_accounts.v1.QueryInterchainAccountAddressRequest")
	proto.RegisterType((*QueryInterchainAccountAddressResponse)(nil), "ibc.applications.interchain_accounts.v1.QueryInterchainAccountAddressResponse")
}

func init() {
	proto.RegisterFile("ibc/applications/interchain_accounts/v1/query.proto", fileDescriptor_72a16b57c3343764)
}

var fileDescriptor_72a16b57c3343764 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3f, 0x4b, 0x3b, 0x31,
	0x18, 0xbe, 0x0c, 0xbf, 0x1f, 0x9a, 0xf1, 0xe8, 0x50, 0x0e, 0xb9, 0x4a, 0x51, 0x74, 0x69, 0x62,
	0x5b, 0x44, 0x10, 0x97, 0x76, 0xeb, 0xa0, 0x68, 0x47, 0x11, 0x8e, 0x5c, 0x2e, 0x5c, 0x03, 0x6d,
	0xde, 0x34, 0xc9, 0x15, 0xfa, 0x0d, 0x1c, 0xfd, 0x08, 0xfd, 0x1e, 0xce, 0x82, 0x63, 0x47, 0x47,
	0x69, 0x17, 0x3f, 0x86, 0xf4, 0xae, 0xd4, 0x82, 0x16, 0x6f, 0x70, 0x7b, 0xc9, 0xf3, 0x3e, 0x7f,
	0x78, 0x9f, 0xe0, 0xb6, 0x8c, 0x39, 0x65, 0x5a, 0x0f, 0x25, 0x67, 0x4e, 0x82, 0xb2, 0x54, 0x2a,
	0x27, 0x0c, 0x1f, 0x30, 0xa9, 0x22, 0xc6, 0x39, 0x64, 0xca, 0x59, 0x3a, 0x69, 0xd2, 0x71, 0x26,
	0xcc, 0x94, 0x68, 0x03, 0x0e, 0xfc, 0x13, 0x19, 0x73, 0xb2, 0x4d, 0x22, 0x3f, 0x90, 0xc8, 0xa4,
	0x19, 0x54, 0x52, 0x48, 0x21, 0xe7, 0xd0, 0xd5, 0x54, 0xd0, 0x83, 0x83, 0x14, 0x20, 0x1d, 0x0a,
	0xca, 0xb4, 0xa4, 0x4c, 0x29, 0x70, 0x6b, 0x91, 0x02, 0x3d, 0x2f, 0x9b, 0x68, 0x3d, 0x17, 0xb4,
	0x7a, 0x8c, 0x8f, 0xee, 0x56, 0x11, 0x7b, 0x9b, 0xe5, 0x4e, 0x81, 0x77, 0x92, 0xc4, 0x08, 0x6b,
	0xfb, 0x62, 0x9c, 0x09, 0xeb, 0xfc, 0x33, 0x5c, 0xc9, 0x9f, 0x85, 0xd1, 0xcc, 0xb8, 0x69, 0xa4,
	0xc1, 0xb8, 0x48, 0x26, 0x55, 0x74, 0x88, 0x4e, 0xf7, 0xfb, 0xfe, 0x36, 0x76, 0x0b, 0xc6, 0xf5,
	0x92, 0xcb, 0xbd, 0xc7, 0x59, 0xcd, 0xfb, 0x98, 0xd5, 0xbc, 0xba, 0xc0, 0xc7, 0xbf, 0x78, 0x58,
	0x0d, 0xca, 0x0a, 0xff, 0x0a, 0x07, 0xdf, 0x43, 0x47, 0xac, 0xd8, 0x5a, 0x5b, 0x55, 0xe5, 0x0e,
	0x95, 0xd6, 0x0b, 0xc2, 0xff, 0x72, 0x1f, 0xff, 0x19, 0xe1, 0xea, 0x2e, 0x33, 0xff, 0x9a, 0x94,
	0xac, 0x81, 0x94, 0x39, 0x4c, 0x70, 0xf3, 0x57, 0x72, 0xc5, 0x0d, 0xea, 0x5e, 0xf7, 0xe1, 0x75,
	0x11, 0xa2, 0xf9, 0x22, 0x44, 0xef, 0x8b, 0x10, 0x3d, 0x2d, 0x43, 0x6f, 0xbe, 0x0c, 0xbd, 0xb7,
	0x65, 0xe8, 0xdd, 0x77, 0x53, 0xe9, 0x06, 0x59, 0x4c, 0x38, 0x8c, 0x28, 0x07, 0x3b, 0x02, 0x4b,
	0x65, 0xcc, 0x1b, 0x29, 0xd0, 0x11, 0x24, 0xd9, 0x50, 0xd8, 0xd5, 0x07, 0xb0, 0xb4, 0x75, 0xd1,
	0xf8, 0x8a, 0xd0, 0xd8, 0x74, 0xef, 0xa6, 0x5a, 0xd8, 0xf8, 0x7f, 0xde, 0x7b, 0xfb, 0x33, 0x00,
	0x00, 0xff, 0xff, 0x2e, 0x9f, 0x3c, 0xd1, 0xc2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Query to get the address of an interchain account
	InterchainAccountAddress(ctx context.Context, in *QueryInterchainAccountAddressRequest, opts ...grpc.CallOption) (*QueryInterchainAccountAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) InterchainAccountAddress(ctx context.Context, in *QueryInterchainAccountAddressRequest, opts ...grpc.CallOption) (*QueryInterchainAccountAddressResponse, error) {
	out := new(QueryInterchainAccountAddressResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.interchain_accounts.v1.Query/InterchainAccountAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Query to get the address of an interchain account
	InterchainAccountAddress(context.Context, *QueryInterchainAccountAddressRequest) (*QueryInterchainAccountAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) InterchainAccountAddress(ctx context.Context, req *QueryInterchainAccountAddressRequest) (*QueryInterchainAccountAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InterchainAccountAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_InterchainAccountAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInterchainAccountAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).InterchainAccountAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.interchain_accounts.v1.Query/InterchainAccountAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).InterchainAccountAddress(ctx, req.(*QueryInterchainAccountAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.interchain_accounts.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InterchainAccountAddress",
			Handler:    _Query_InterchainAccountAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/interchain_accounts/v1/query.proto",
}

func (m *QueryInterchainAccountAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInterchainAccountAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInterchainAccountAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CounterpartyPortId) > 0 {
		i -= len(m.CounterpartyPortId)
		copy(dAtA[i:], m.CounterpartyPortId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CounterpartyPortId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryInterchainAccountAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInterchainAccountAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInterchainAccountAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InterchainAccountAddress) > 0 {
		i -= len(m.InterchainAccountAddress)
		copy(dAtA[i:], m.InterchainAccountAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.InterchainAccountAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryInterchainAccountAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CounterpartyPortId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryInterchainAccountAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InterchainAccountAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryInterchainAccountAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryInterchainAccountAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInterchainAccountAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CounterpartyPortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CounterpartyPortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryInterchainAccountAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryInterchainAccountAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInterchainAccountAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterchainAccountAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterchainAccountAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)