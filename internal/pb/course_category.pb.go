// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/course_category.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Blank struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blank) Reset()         { *m = Blank{} }
func (m *Blank) String() string { return proto.CompactTextString(m) }
func (*Blank) ProtoMessage()    {}
func (*Blank) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{0}
}

func (m *Blank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blank.Unmarshal(m, b)
}
func (m *Blank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blank.Marshal(b, m, deterministic)
}
func (m *Blank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blank.Merge(m, src)
}
func (m *Blank) XXX_Size() int {
	return xxx_messageInfo_Blank.Size(m)
}
func (m *Blank) XXX_DiscardUnknown() {
	xxx_messageInfo_Blank.DiscardUnknown(m)
}

var xxx_messageInfo_Blank proto.InternalMessageInfo

type Category struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{1}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateCategoryRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCategoryRequest) Reset()         { *m = CreateCategoryRequest{} }
func (m *CreateCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryRequest) ProtoMessage()    {}
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{2}
}

func (m *CreateCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryRequest.Unmarshal(m, b)
}
func (m *CreateCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryRequest.Marshal(b, m, deterministic)
}
func (m *CreateCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryRequest.Merge(m, src)
}
func (m *CreateCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryRequest.Size(m)
}
func (m *CreateCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryRequest proto.InternalMessageInfo

func (m *CreateCategoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCategoryRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CategoryList struct {
	Categories           []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CategoryList) Reset()         { *m = CategoryList{} }
func (m *CategoryList) String() string { return proto.CompactTextString(m) }
func (*CategoryList) ProtoMessage()    {}
func (*CategoryList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{3}
}

func (m *CategoryList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryList.Unmarshal(m, b)
}
func (m *CategoryList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryList.Marshal(b, m, deterministic)
}
func (m *CategoryList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryList.Merge(m, src)
}
func (m *CategoryList) XXX_Size() int {
	return xxx_messageInfo_CategoryList.Size(m)
}
func (m *CategoryList) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryList.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryList proto.InternalMessageInfo

func (m *CategoryList) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type CategoryGetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryGetRequest) Reset()         { *m = CategoryGetRequest{} }
func (m *CategoryGetRequest) String() string { return proto.CompactTextString(m) }
func (*CategoryGetRequest) ProtoMessage()    {}
func (*CategoryGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{4}
}

func (m *CategoryGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryGetRequest.Unmarshal(m, b)
}
func (m *CategoryGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryGetRequest.Marshal(b, m, deterministic)
}
func (m *CategoryGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryGetRequest.Merge(m, src)
}
func (m *CategoryGetRequest) XXX_Size() int {
	return xxx_messageInfo_CategoryGetRequest.Size(m)
}
func (m *CategoryGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryGetRequest proto.InternalMessageInfo

func (m *CategoryGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Blank)(nil), "pb.blank")
	proto.RegisterType((*Category)(nil), "pb.Category")
	proto.RegisterType((*CreateCategoryRequest)(nil), "pb.CreateCategoryRequest")
	proto.RegisterType((*CategoryList)(nil), "pb.CategoryList")
	proto.RegisterType((*CategoryGetRequest)(nil), "pb.CategoryGetRequest")
}

func init() {
	proto.RegisterFile("proto/course_category.proto", fileDescriptor_fde79cb73866aac8)
}

var fileDescriptor_fde79cb73866aac8 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbb, 0xdb, 0xff, 0x5f, 0xed, 0xa4, 0x56, 0x19, 0x54, 0xa2, 0x5e, 0xea, 0xe2, 0x21,
	0x07, 0x69, 0xa5, 0xe2, 0x49, 0x4f, 0xed, 0xa1, 0x17, 0x05, 0x49, 0x6f, 0x5e, 0x64, 0x93, 0x0c,
	0xb2, 0x98, 0x26, 0x71, 0xb3, 0x15, 0xfc, 0x9c, 0x7e, 0x21, 0xc9, 0xd2, 0xc4, 0xa4, 0x46, 0xc1,
	0xdb, 0x32, 0xf3, 0xf8, 0xcd, 0x7b, 0x8f, 0x85, 0xd3, 0x4c, 0xa7, 0x26, 0x1d, 0x87, 0xe9, 0x4a,
	0xe7, 0xf4, 0x14, 0x4a, 0x43, 0xcf, 0xa9, 0x7e, 0x1f, 0xd9, 0x29, 0xf2, 0x2c, 0x10, 0xdb, 0xf0,
	0x3f, 0x88, 0x65, 0xf2, 0x22, 0x1e, 0x60, 0x67, 0xb6, 0x5e, 0xe3, 0x00, 0xb8, 0x8a, 0x5c, 0x36,
	0x64, 0x5e, 0xcf, 0xe7, 0x2a, 0x42, 0x84, 0x7f, 0x89, 0x5c, 0x92, 0xcb, 0xed, 0xc4, 0xbe, 0x71,
	0x08, 0x4e, 0x44, 0x79, 0xa8, 0x55, 0x66, 0x54, 0x9a, 0xb8, 0x5d, 0xbb, 0xaa, 0x8f, 0xc4, 0x3d,
	0x1c, 0xce, 0x34, 0x49, 0x43, 0x25, 0xd7, 0xa7, 0xd7, 0x15, 0xe5, 0xa6, 0xc2, 0xb1, 0x9f, 0x71,
	0xfc, 0x3b, 0xee, 0x16, 0xfa, 0x25, 0xe8, 0x4e, 0xe5, 0x06, 0x2f, 0x00, 0xd6, 0x79, 0x14, 0xe5,
	0x2e, 0x1b, 0x76, 0x3d, 0x67, 0xd2, 0x1f, 0x65, 0xc1, 0xa8, 0x3a, 0x57, 0xdb, 0x8b, 0x73, 0xc0,
	0x72, 0x3e, 0x27, 0x53, 0x3a, 0xd9, 0x08, 0x3a, 0xf9, 0xe0, 0xb0, 0x57, 0xca, 0x16, 0xa4, 0xdf,
	0x54, 0x48, 0x78, 0x03, 0x83, 0x66, 0x0c, 0x3c, 0xb6, 0x57, 0xda, 0xa2, 0x9d, 0x34, 0x0c, 0x88,
	0x0e, 0xce, 0xe1, 0xa0, 0x29, 0x5c, 0x18, 0x4d, 0x72, 0xf9, 0x1b, 0x62, 0xbf, 0x8e, 0x28, 0x92,
	0x8a, 0x8e, 0xc7, 0xd0, 0x87, 0xb3, 0x36, 0xd0, 0x54, 0x45, 0x4a, 0x53, 0x58, 0x14, 0x24, 0xe3,
	0x3f, 0x18, 0xf3, 0xd8, 0x25, 0xc3, 0x31, 0x0c, 0x0a, 0xfe, 0xac, 0x6a, 0x09, 0x7b, 0x85, 0xca,
	0xfe, 0x87, 0x36, 0x1b, 0x78, 0x0d, 0xce, 0x9c, 0x4c, 0xd5, 0xc3, 0x51, 0x5d, 0xf2, 0xd5, 0xea,
	0xe6, 0xad, 0xe9, 0xee, 0xa3, 0xa3, 0x12, 0x43, 0x3a, 0x91, 0xf1, 0x38, 0x0b, 0x82, 0x2d, 0xfb,
	0xfb, 0xae, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xab, 0xd8, 0x06, 0x93, 0x9c, 0x02, 0x00, 0x00,
}
