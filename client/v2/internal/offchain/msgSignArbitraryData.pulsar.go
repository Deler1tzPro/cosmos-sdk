// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package offchain

import (
	_ "cosmossdk.io/api/amino"
	_ "cosmossdk.io/api/cosmos/msg/v1"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_MsgSignArbitraryData            protoreflect.MessageDescriptor
	fd_MsgSignArbitraryData_app_domain protoreflect.FieldDescriptor
	fd_MsgSignArbitraryData_signer     protoreflect.FieldDescriptor
	fd_MsgSignArbitraryData_data       protoreflect.FieldDescriptor
)

func init() {
	file_offchain_msgSignArbitraryData_proto_init()
	md_MsgSignArbitraryData = File_offchain_msgSignArbitraryData_proto.Messages().ByName("MsgSignArbitraryData")
	fd_MsgSignArbitraryData_app_domain = md_MsgSignArbitraryData.Fields().ByName("app_domain")
	fd_MsgSignArbitraryData_signer = md_MsgSignArbitraryData.Fields().ByName("signer")
	fd_MsgSignArbitraryData_data = md_MsgSignArbitraryData.Fields().ByName("data")
}

var _ protoreflect.Message = (*fastReflection_MsgSignArbitraryData)(nil)

type fastReflection_MsgSignArbitraryData MsgSignArbitraryData

func (x *MsgSignArbitraryData) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSignArbitraryData)(x)
}

func (x *MsgSignArbitraryData) slowProtoReflect() protoreflect.Message {
	mi := &file_offchain_msgSignArbitraryData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSignArbitraryData_messageType fastReflection_MsgSignArbitraryData_messageType
var _ protoreflect.MessageType = fastReflection_MsgSignArbitraryData_messageType{}

type fastReflection_MsgSignArbitraryData_messageType struct{}

func (x fastReflection_MsgSignArbitraryData_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSignArbitraryData)(nil)
}
func (x fastReflection_MsgSignArbitraryData_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSignArbitraryData)
}
func (x fastReflection_MsgSignArbitraryData_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSignArbitraryData
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSignArbitraryData) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSignArbitraryData
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSignArbitraryData) Type() protoreflect.MessageType {
	return _fastReflection_MsgSignArbitraryData_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSignArbitraryData) New() protoreflect.Message {
	return new(fastReflection_MsgSignArbitraryData)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSignArbitraryData) Interface() protoreflect.ProtoMessage {
	return (*MsgSignArbitraryData)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSignArbitraryData) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.AppDomain != "" {
		value := protoreflect.ValueOfString(x.AppDomain)
		if !f(fd_MsgSignArbitraryData_app_domain, value) {
			return
		}
	}
	if x.Signer != "" {
		value := protoreflect.ValueOfString(x.Signer)
		if !f(fd_MsgSignArbitraryData_signer, value) {
			return
		}
	}
	if x.Data != "" {
		value := protoreflect.ValueOfString(x.Data)
		if !f(fd_MsgSignArbitraryData_data, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgSignArbitraryData) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "offchain.MsgSignArbitraryData.app_domain":
		return x.AppDomain != ""
	case "offchain.MsgSignArbitraryData.signer":
		return x.Signer != ""
	case "offchain.MsgSignArbitraryData.data":
		return x.Data != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: offchain.MsgSignArbitraryData"))
		}
		panic(fmt.Errorf("message offchain.MsgSignArbitraryData does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSignArbitraryData) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "offchain.MsgSignArbitraryData.app_domain":
		x.AppDomain = ""
	case "offchain.MsgSignArbitraryData.signer":
		x.Signer = ""
	case "offchain.MsgSignArbitraryData.data":
		x.Data = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: offchain.MsgSignArbitraryData"))
		}
		panic(fmt.Errorf("message offchain.MsgSignArbitraryData does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSignArbitraryData) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "offchain.MsgSignArbitraryData.app_domain":
		value := x.AppDomain
		return protoreflect.ValueOfString(value)
	case "offchain.MsgSignArbitraryData.signer":
		value := x.Signer
		return protoreflect.ValueOfString(value)
	case "offchain.MsgSignArbitraryData.data":
		value := x.Data
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: offchain.MsgSignArbitraryData"))
		}
		panic(fmt.Errorf("message offchain.MsgSignArbitraryData does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSignArbitraryData) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "offchain.MsgSignArbitraryData.app_domain":
		x.AppDomain = value.Interface().(string)
	case "offchain.MsgSignArbitraryData.signer":
		x.Signer = value.Interface().(string)
	case "offchain.MsgSignArbitraryData.data":
		x.Data = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: offchain.MsgSignArbitraryData"))
		}
		panic(fmt.Errorf("message offchain.MsgSignArbitraryData does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSignArbitraryData) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "offchain.MsgSignArbitraryData.app_domain":
		panic(fmt.Errorf("field app_domain of message offchain.MsgSignArbitraryData is not mutable"))
	case "offchain.MsgSignArbitraryData.signer":
		panic(fmt.Errorf("field signer of message offchain.MsgSignArbitraryData is not mutable"))
	case "offchain.MsgSignArbitraryData.data":
		panic(fmt.Errorf("field data of message offchain.MsgSignArbitraryData is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: offchain.MsgSignArbitraryData"))
		}
		panic(fmt.Errorf("message offchain.MsgSignArbitraryData does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSignArbitraryData) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "offchain.MsgSignArbitraryData.app_domain":
		return protoreflect.ValueOfString("")
	case "offchain.MsgSignArbitraryData.signer":
		return protoreflect.ValueOfString("")
	case "offchain.MsgSignArbitraryData.data":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: offchain.MsgSignArbitraryData"))
		}
		panic(fmt.Errorf("message offchain.MsgSignArbitraryData does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSignArbitraryData) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in offchain.MsgSignArbitraryData", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSignArbitraryData) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSignArbitraryData) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgSignArbitraryData) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSignArbitraryData) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSignArbitraryData)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.AppDomain)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Signer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Data)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgSignArbitraryData)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Data) > 0 {
			i -= len(x.Data)
			copy(dAtA[i:], x.Data)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Data)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Signer) > 0 {
			i -= len(x.Signer)
			copy(dAtA[i:], x.Signer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signer)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.AppDomain) > 0 {
			i -= len(x.AppDomain)
			copy(dAtA[i:], x.AppDomain)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AppDomain)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgSignArbitraryData)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSignArbitraryData: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSignArbitraryData: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AppDomain", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AppDomain = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Signer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Data = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: offchain/msgSignArbitraryData.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MsgSignArbitraryData defines an arbitrary, general-purpose, off-chain message
type MsgSignArbitraryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AppDomain is the application requesting off-chain message signing
	AppDomain string `protobuf:"bytes,1,opt,name=app_domain,json=appDomain,proto3" json:"app_domain,omitempty"`
	// Signer is the sdk.AccAddress of the message signer
	Signer string `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
	// Data represents the raw bytes of the content that is signed (text, json, etc)
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MsgSignArbitraryData) Reset() {
	*x = MsgSignArbitraryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offchain_msgSignArbitraryData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSignArbitraryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSignArbitraryData) ProtoMessage() {}

// Deprecated: Use MsgSignArbitraryData.ProtoReflect.Descriptor instead.
func (*MsgSignArbitraryData) Descriptor() ([]byte, []int) {
	return file_offchain_msgSignArbitraryData_proto_rawDescGZIP(), []int{0}
}

func (x *MsgSignArbitraryData) GetAppDomain() string {
	if x != nil {
		return x.AppDomain
	}
	return ""
}

func (x *MsgSignArbitraryData) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

func (x *MsgSignArbitraryData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_offchain_msgSignArbitraryData_proto protoreflect.FileDescriptor

var file_offchain_msgSignArbitraryData_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x73, 0x67, 0x53, 0x69,
	0x67, 0x6e, 0x41, 0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x1a,
	0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x14, 0x4d, 0x73, 0x67, 0x53, 0x69,
	0x67, 0x6e, 0x41, 0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x30,
	0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18,
	0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x3a, 0x2d, 0x82, 0xe7, 0xb0, 0x2a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x72, 0x8a, 0xe7, 0xb0, 0x2a, 0x1d, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x4d,
	0x73, 0x67, 0x53, 0x69, 0x67, 0x6e, 0x41, 0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x42, 0xa3, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x66, 0x66, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x42, 0x19, 0x4d, 0x73, 0x67, 0x53, 0x69, 0x67, 0x6e, 0x41, 0x72, 0x62,
	0x69, 0x74, 0x72, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0xa2, 0x02, 0x03, 0x4f, 0x58,
	0x58, 0xaa, 0x02, 0x08, 0x4f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0xca, 0x02, 0x08, 0x4f,
	0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0xe2, 0x02, 0x14, 0x4f, 0x66, 0x66, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x08, 0x4f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4a, 0xef, 0x04, 0x0a, 0x06, 0x12, 0x04,
	0x00, 0x00, 0x12, 0x01, 0x0a, 0x08, 0x0a, 0x01, 0x0c, 0x12, 0x03, 0x00, 0x00, 0x12, 0x0a, 0x08,
	0x0a, 0x01, 0x02, 0x12, 0x03, 0x02, 0x00, 0x11, 0x0a, 0x09, 0x0a, 0x02, 0x03, 0x00, 0x12, 0x03,
	0x04, 0x00, 0x23, 0x0a, 0x09, 0x0a, 0x02, 0x03, 0x01, 0x12, 0x03, 0x05, 0x00, 0x21, 0x0a, 0x09,
	0x0a, 0x02, 0x03, 0x02, 0x12, 0x03, 0x06, 0x00, 0x1b, 0x0a, 0x5b, 0x0a, 0x02, 0x04, 0x00, 0x12,
	0x04, 0x09, 0x00, 0x12, 0x01, 0x1a, 0x4f, 0x20, 0x4d, 0x73, 0x67, 0x53, 0x69, 0x67, 0x6e, 0x41,
	0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x20, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x72,
	0x79, 0x2c, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2d, 0x70, 0x75, 0x72, 0x70, 0x6f,
	0x73, 0x65, 0x2c, 0x20, 0x6f, 0x66, 0x66, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x20, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x0a, 0x0a, 0x0a, 0x0a, 0x03, 0x04, 0x00, 0x01, 0x12, 0x03, 0x09,
	0x08, 0x1c, 0x0a, 0x0a, 0x0a, 0x03, 0x04, 0x00, 0x07, 0x12, 0x03, 0x0a, 0x02, 0x42, 0x0a, 0x0e,
	0x0a, 0x07, 0x04, 0x00, 0x07, 0xf1, 0x8c, 0xa6, 0x05, 0x12, 0x03, 0x0a, 0x02, 0x42, 0x0a, 0x0a,
	0x0a, 0x03, 0x04, 0x00, 0x07, 0x12, 0x03, 0x0b, 0x02, 0x2b, 0x0a, 0x0f, 0x0a, 0x08, 0x04, 0x00,
	0x07, 0xf0, 0x8c, 0xa6, 0x05, 0x00, 0x12, 0x03, 0x0b, 0x02, 0x2b, 0x0a, 0x50, 0x0a, 0x04, 0x04,
	0x00, 0x02, 0x00, 0x12, 0x03, 0x0d, 0x02, 0x18, 0x1a, 0x43, 0x20, 0x41, 0x70, 0x70, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x20, 0x69, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x20, 0x6f, 0x66, 0x66, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x20, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x20, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x0a, 0x0a, 0x0c, 0x0a,
	0x05, 0x04, 0x00, 0x02, 0x00, 0x05, 0x12, 0x03, 0x0d, 0x02, 0x08, 0x0a, 0x0c, 0x0a, 0x05, 0x04,
	0x00, 0x02, 0x00, 0x01, 0x12, 0x03, 0x0d, 0x09, 0x13, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02,
	0x00, 0x03, 0x12, 0x03, 0x0d, 0x16, 0x17, 0x0a, 0x41, 0x0a, 0x04, 0x04, 0x00, 0x02, 0x01, 0x12,
	0x03, 0x0f, 0x02, 0x45, 0x1a, 0x34, 0x20, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x20, 0x69, 0x73,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x64, 0x6b, 0x2e, 0x41, 0x63, 0x63, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x20, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x0a, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00,
	0x02, 0x01, 0x05, 0x12, 0x03, 0x0f, 0x02, 0x08, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x01,
	0x01, 0x12, 0x03, 0x0f, 0x09, 0x0f, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x01, 0x03, 0x12,
	0x03, 0x0f, 0x12, 0x13, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x01, 0x08, 0x12, 0x03, 0x0f,
	0x14, 0x44, 0x0a, 0x0f, 0x0a, 0x08, 0x04, 0x00, 0x02, 0x01, 0x08, 0xca, 0xd6, 0x05, 0x12, 0x03,
	0x0f, 0x15, 0x43, 0x0a, 0x5c, 0x0a, 0x04, 0x04, 0x00, 0x02, 0x02, 0x12, 0x03, 0x11, 0x02, 0x12,
	0x1a, 0x4f, 0x20, 0x44, 0x61, 0x74, 0x61, 0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x61, 0x77, 0x20, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x20,
	0x74, 0x68, 0x61, 0x74, 0x20, 0x69, 0x73, 0x20, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x28,
	0x74, 0x65, 0x78, 0x74, 0x2c, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x2c, 0x20, 0x65, 0x74, 0x63, 0x29,
	0x0a, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x02, 0x05, 0x12, 0x03, 0x11, 0x02, 0x08, 0x0a,
	0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x02, 0x01, 0x12, 0x03, 0x11, 0x09, 0x0d, 0x0a, 0x0c, 0x0a,
	0x05, 0x04, 0x00, 0x02, 0x02, 0x03, 0x12, 0x03, 0x11, 0x10, 0x11, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_offchain_msgSignArbitraryData_proto_rawDescOnce sync.Once
	file_offchain_msgSignArbitraryData_proto_rawDescData = file_offchain_msgSignArbitraryData_proto_rawDesc
)

func file_offchain_msgSignArbitraryData_proto_rawDescGZIP() []byte {
	file_offchain_msgSignArbitraryData_proto_rawDescOnce.Do(func() {
		file_offchain_msgSignArbitraryData_proto_rawDescData = protoimpl.X.CompressGZIP(file_offchain_msgSignArbitraryData_proto_rawDescData)
	})
	return file_offchain_msgSignArbitraryData_proto_rawDescData
}

var file_offchain_msgSignArbitraryData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_offchain_msgSignArbitraryData_proto_goTypes = []interface{}{
	(*MsgSignArbitraryData)(nil), // 0: offchain.MsgSignArbitraryData
}
var file_offchain_msgSignArbitraryData_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_offchain_msgSignArbitraryData_proto_init() }
func file_offchain_msgSignArbitraryData_proto_init() {
	if File_offchain_msgSignArbitraryData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_offchain_msgSignArbitraryData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSignArbitraryData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_offchain_msgSignArbitraryData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_offchain_msgSignArbitraryData_proto_goTypes,
		DependencyIndexes: file_offchain_msgSignArbitraryData_proto_depIdxs,
		MessageInfos:      file_offchain_msgSignArbitraryData_proto_msgTypes,
	}.Build()
	File_offchain_msgSignArbitraryData_proto = out.File
	file_offchain_msgSignArbitraryData_proto_rawDesc = nil
	file_offchain_msgSignArbitraryData_proto_goTypes = nil
	file_offchain_msgSignArbitraryData_proto_depIdxs = nil
}