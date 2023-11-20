// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: public.proto

package public

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{0}
}
func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return m.Size()
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

func (*HealthRequest) XXX_MessageName() string {
	return "disputes.public.HealthRequest"
}

type HealthResponse struct {
	Postgres             *HealthResponse_Service `protobuf:"bytes,1,opt,name=postgres,proto3" json:"postgres"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{1}
}
func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return m.Size()
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetPostgres() *HealthResponse_Service {
	if m != nil {
		return m.Postgres
	}
	return nil
}

func (*HealthResponse) XXX_MessageName() string {
	return "disputes.public.HealthResponse"
}

type HealthResponse_Service struct {
	Status               bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse_Service) Reset()         { *m = HealthResponse_Service{} }
func (m *HealthResponse_Service) String() string { return proto.CompactTextString(m) }
func (*HealthResponse_Service) ProtoMessage()    {}
func (*HealthResponse_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{1, 0}
}
func (m *HealthResponse_Service) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthResponse_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthResponse_Service.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HealthResponse_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse_Service.Merge(m, src)
}
func (m *HealthResponse_Service) XXX_Size() int {
	return m.Size()
}
func (m *HealthResponse_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse_Service.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse_Service proto.InternalMessageInfo

func (m *HealthResponse_Service) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *HealthResponse_Service) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (*HealthResponse_Service) XXX_MessageName() string {
	return "disputes.public.HealthResponse.Service"
}
func init() {
	proto.RegisterType((*HealthRequest)(nil), "disputes.public.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "disputes.public.HealthResponse")
	proto.RegisterType((*HealthResponse_Service)(nil), "disputes.public.HealthResponse.Service")
}

func init() { proto.RegisterFile("public.proto", fileDescriptor_413a91106d7bcce8) }

var fileDescriptor_413a91106d7bcce8 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4d, 0x4a, 0x03, 0x31,
	0x18, 0x86, 0x4d, 0xc5, 0x69, 0x1b, 0x7f, 0x0a, 0x41, 0xa4, 0x0c, 0x25, 0x96, 0xd9, 0x38, 0x1b,
	0x33, 0x50, 0x17, 0xae, 0xd5, 0x8d, 0xee, 0xca, 0xb8, 0xeb, 0x6e, 0x66, 0x8c, 0x99, 0xc0, 0x38,
	0x5f, 0xcc, 0x8f, 0x07, 0xf0, 0x00, 0x6e, 0xdc, 0x78, 0xa4, 0x2e, 0x05, 0x2f, 0x20, 0x53, 0x0f,
	0x22, 0x24, 0x55, 0x50, 0x10, 0x77, 0x79, 0x9e, 0x7c, 0x2f, 0xdf, 0x9b, 0xe0, 0x1d, 0xe5, 0xca,
	0x46, 0x56, 0x4c, 0x69, 0xb0, 0x40, 0x46, 0x37, 0xd2, 0x28, 0x67, 0xb9, 0x61, 0x41, 0xc7, 0x13,
	0x01, 0x20, 0x1a, 0x9e, 0x15, 0x4a, 0x66, 0x45, 0xdb, 0x82, 0x2d, 0xac, 0x84, 0xd6, 0x84, 0xf1,
	0xf8, 0x58, 0x48, 0x5b, 0xbb, 0x92, 0x55, 0x70, 0x97, 0x09, 0x10, 0x90, 0x79, 0x5d, 0xba, 0x5b,
	0x4f, 0x1e, 0xfc, 0x29, 0x8c, 0x27, 0x23, 0xbc, 0x7b, 0xc9, 0x8b, 0xc6, 0xd6, 0x39, 0xbf, 0x77,
	0xdc, 0xd8, 0xe4, 0x09, 0xe1, 0xbd, 0x2f, 0x63, 0x14, 0xb4, 0x86, 0x93, 0x0b, 0x3c, 0x50, 0x60,
	0xac, 0xd0, 0xdc, 0x8c, 0xd1, 0x14, 0xa5, 0xdb, 0xb3, 0x23, 0xf6, 0xab, 0x14, 0xfb, 0x19, 0x61,
	0xd7, 0x5c, 0x3f, 0xc8, 0x8a, 0xe7, 0xdf, 0xc1, 0xf8, 0x14, 0xf7, 0xd7, 0x92, 0x1c, 0xe0, 0xc8,
	0xd8, 0xc2, 0x3a, 0x33, 0xee, 0x4d, 0x51, 0x3a, 0xc8, 0xd7, 0x44, 0xf6, 0xf1, 0x16, 0xd7, 0x1a,
	0xf4, 0x78, 0x73, 0x8a, 0xd2, 0x61, 0x1e, 0x60, 0x26, 0xf0, 0x70, 0xee, 0x77, 0x9c, 0xcd, 0xaf,
	0xc8, 0x02, 0x47, 0x61, 0x13, 0xa1, 0x7f, 0x56, 0xf0, 0xef, 0x88, 0x0f, 0xff, 0xa9, 0x98, 0x8c,
	0x1e, 0xdf, 0x3e, 0x9e, 0x7b, 0x43, 0xd2, 0xcf, 0x6a, 0x7f, 0x71, 0x3e, 0x59, 0x76, 0x14, 0xbd,
	0x76, 0x14, 0xbd, 0x77, 0x14, 0xbd, 0xac, 0xe8, 0xc6, 0x72, 0x45, 0xd1, 0x22, 0x0a, 0xe9, 0x32,
	0xf2, 0xff, 0x75, 0xf2, 0x19, 0x00, 0x00, 0xff, 0xff, 0xed, 0x04, 0xef, 0x82, 0x9d, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PublicAPIClient is the client API for PublicAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublicAPIClient interface {
	// Состояние сервиса
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type publicAPIClient struct {
	cc *grpc.ClientConn
}

func NewPublicAPIClient(cc *grpc.ClientConn) PublicAPIClient {
	return &publicAPIClient{cc}
}

func (c *publicAPIClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/disputes.public.PublicAPI/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicAPIServer is the server API for PublicAPI service.
type PublicAPIServer interface {
	// Состояние сервиса
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
}

// UnimplementedPublicAPIServer can be embedded to have forward compatible implementations.
type UnimplementedPublicAPIServer struct {
}

func (*UnimplementedPublicAPIServer) Health(ctx context.Context, req *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}

func RegisterPublicAPIServer(s *grpc.Server, srv PublicAPIServer) {
	s.RegisterService(&_PublicAPI_serviceDesc, srv)
}

func _PublicAPI_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicAPIServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/disputes.public.PublicAPI/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicAPIServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PublicAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "disputes.public.PublicAPI",
	HandlerType: (*PublicAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _PublicAPI_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public.proto",
}

func (m *HealthRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HealthRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *HealthResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HealthResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Postgres != nil {
		{
			size, err := m.Postgres.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPublic(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HealthResponse_Service) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthResponse_Service) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HealthResponse_Service) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintPublic(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Status {
		i--
		if m.Status {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func encodeVarintPublic(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublic(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HealthRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HealthResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Postgres != nil {
		l = m.Postgres.Size()
		n += 1 + l + sovPublic(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HealthResponse_Service) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status {
		n += 2
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovPublic(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPublic(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublic(x uint64) (n int) {
	return sovPublic(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HealthRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublic
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
			return fmt.Errorf("proto: HealthRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPublic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPublic
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
func (m *HealthResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublic
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
			return fmt.Errorf("proto: HealthResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Postgres", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublic
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
				return ErrInvalidLengthPublic
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Postgres == nil {
				m.Postgres = &HealthResponse_Service{}
			}
			if err := m.Postgres.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPublic
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
func (m *HealthResponse_Service) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublic
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
			return fmt.Errorf("proto: Service: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Service: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublic
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
			m.Status = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublic
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
				return ErrInvalidLengthPublic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPublic
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
func skipPublic(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublic
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
					return 0, ErrIntOverflowPublic
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
					return 0, ErrIntOverflowPublic
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
				return 0, ErrInvalidLengthPublic
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublic
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublic
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublic        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublic          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublic = fmt.Errorf("proto: unexpected end of group")
)
