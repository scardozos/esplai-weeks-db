// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: api/weeksdb/weeks_db.proto

package weeksdb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Must be a Saturday
type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int32 `protobuf:"varint,1,opt,name=Year,proto3" json:"Year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=Month,proto3" json:"Month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=Day,proto3" json:"Day,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_weeksdb_weeks_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_api_weeksdb_weeks_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_api_weeksdb_weeks_db_proto_rawDescGZIP(), []int{0}
}

func (x *Date) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Date) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

// Must be empty
type GetStaticWeeksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStaticWeeksRequest) Reset() {
	*x = GetStaticWeeksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_weeksdb_weeks_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaticWeeksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticWeeksRequest) ProtoMessage() {}

func (x *GetStaticWeeksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_weeksdb_weeks_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticWeeksRequest.ProtoReflect.Descriptor instead.
func (*GetStaticWeeksRequest) Descriptor() ([]byte, []int) {
	return file_api_weeksdb_weeks_db_proto_rawDescGZIP(), []int{1}
}

// Returns a list with all of the weeks present in
// the database as `Date` messages
type GetStaticWeeksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaticWeeks []*Date `protobuf:"bytes,1,rep,name=static_weeks,json=staticWeeks,proto3" json:"static_weeks,omitempty"`
}

func (x *GetStaticWeeksResponse) Reset() {
	*x = GetStaticWeeksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_weeksdb_weeks_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaticWeeksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticWeeksResponse) ProtoMessage() {}

func (x *GetStaticWeeksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_weeksdb_weeks_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticWeeksResponse.ProtoReflect.Descriptor instead.
func (*GetStaticWeeksResponse) Descriptor() ([]byte, []int) {
	return file_api_weeksdb_weeks_db_proto_rawDescGZIP(), []int{2}
}

func (x *GetStaticWeeksResponse) GetStaticWeeks() []*Date {
	if x != nil {
		return x.StaticWeeks
	}
	return nil
}

// Takes a `Date` argument, which is the week to
// add to the static weeks database
type SetStaticWeekRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaticWeek *Date `protobuf:"bytes,1,opt,name=static_week,json=staticWeek,proto3" json:"static_week,omitempty"`
}

func (x *SetStaticWeekRequest) Reset() {
	*x = SetStaticWeekRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_weeksdb_weeks_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStaticWeekRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStaticWeekRequest) ProtoMessage() {}

func (x *SetStaticWeekRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_weeksdb_weeks_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStaticWeekRequest.ProtoReflect.Descriptor instead.
func (*SetStaticWeekRequest) Descriptor() ([]byte, []int) {
	return file_api_weeksdb_weeks_db_proto_rawDescGZIP(), []int{3}
}

func (x *SetStaticWeekRequest) GetStaticWeek() *Date {
	if x != nil {
		return x.StaticWeek
	}
	return nil
}

// Returns the week as a `Date` message added to the database, in case
// the RPC was successful
type SetStaticWeekResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SetWeek *Date `protobuf:"bytes,1,opt,name=set_week,json=setWeek,proto3" json:"set_week,omitempty"`
}

func (x *SetStaticWeekResponse) Reset() {
	*x = SetStaticWeekResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_weeksdb_weeks_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStaticWeekResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStaticWeekResponse) ProtoMessage() {}

func (x *SetStaticWeekResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_weeksdb_weeks_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStaticWeekResponse.ProtoReflect.Descriptor instead.
func (*SetStaticWeekResponse) Descriptor() ([]byte, []int) {
	return file_api_weeksdb_weeks_db_proto_rawDescGZIP(), []int{4}
}

func (x *SetStaticWeekResponse) GetSetWeek() *Date {
	if x != nil {
		return x.SetWeek
	}
	return nil
}

// Takes a `Date` argument, which is the week to
// remove from the static weeks database
type UnsetStaticWeekRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaticWeek *Date `protobuf:"bytes,1,opt,name=static_week,json=staticWeek,proto3" json:"static_week,omitempty"`
}

func (x *UnsetStaticWeekRequest) Reset() {
	*x = UnsetStaticWeekRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_weeksdb_weeks_db_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsetStaticWeekRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsetStaticWeekRequest) ProtoMessage() {}

func (x *UnsetStaticWeekRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_weeksdb_weeks_db_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsetStaticWeekRequest.ProtoReflect.Descriptor instead.
func (*UnsetStaticWeekRequest) Descriptor() ([]byte, []int) {
	return file_api_weeksdb_weeks_db_proto_rawDescGZIP(), []int{5}
}

func (x *UnsetStaticWeekRequest) GetStaticWeek() *Date {
	if x != nil {
		return x.StaticWeek
	}
	return nil
}

// Returns the week as a `Date` removed from the database, in case
// the RPC was successful
type UnsetStaticWeekResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnsetWeek *Date `protobuf:"bytes,1,opt,name=unset_week,json=unsetWeek,proto3" json:"unset_week,omitempty"`
}

func (x *UnsetStaticWeekResponse) Reset() {
	*x = UnsetStaticWeekResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_weeksdb_weeks_db_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsetStaticWeekResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsetStaticWeekResponse) ProtoMessage() {}

func (x *UnsetStaticWeekResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_weeksdb_weeks_db_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsetStaticWeekResponse.ProtoReflect.Descriptor instead.
func (*UnsetStaticWeekResponse) Descriptor() ([]byte, []int) {
	return file_api_weeksdb_weeks_db_proto_rawDescGZIP(), []int{6}
}

func (x *UnsetStaticWeekResponse) GetUnsetWeek() *Date {
	if x != nil {
		return x.UnsetWeek
	}
	return nil
}

// Takes a `Date` argument, which is the week to
// check whether it's present in the static weeks datbase.
type IsWeekStaticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Week *Date `protobuf:"bytes,1,opt,name=week,proto3" json:"week,omitempty"`
}

func (x *IsWeekStaticRequest) Reset() {
	*x = IsWeekStaticRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_weeksdb_weeks_db_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsWeekStaticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsWeekStaticRequest) ProtoMessage() {}

func (x *IsWeekStaticRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_weeksdb_weeks_db_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsWeekStaticRequest.ProtoReflect.Descriptor instead.
func (*IsWeekStaticRequest) Descriptor() ([]byte, []int) {
	return file_api_weeksdb_weeks_db_proto_rawDescGZIP(), []int{7}
}

func (x *IsWeekStaticRequest) GetWeek() *Date {
	if x != nil {
		return x.Week
	}
	return nil
}

// Returns a `bool` stating whether the requested week is static or not,
// and a the requested week as a `Date` message.
type IsWeekStaticResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsStatic      bool  `protobuf:"varint,1,opt,name=is_static,json=isStatic,proto3" json:"is_static,omitempty"`
	RequestedWeek *Date `protobuf:"bytes,2,opt,name=requested_week,json=requestedWeek,proto3" json:"requested_week,omitempty"`
}

func (x *IsWeekStaticResponse) Reset() {
	*x = IsWeekStaticResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_weeksdb_weeks_db_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsWeekStaticResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsWeekStaticResponse) ProtoMessage() {}

func (x *IsWeekStaticResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_weeksdb_weeks_db_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsWeekStaticResponse.ProtoReflect.Descriptor instead.
func (*IsWeekStaticResponse) Descriptor() ([]byte, []int) {
	return file_api_weeksdb_weeks_db_proto_rawDescGZIP(), []int{8}
}

func (x *IsWeekStaticResponse) GetIsStatic() bool {
	if x != nil {
		return x.IsStatic
	}
	return false
}

func (x *IsWeekStaticResponse) GetRequestedWeek() *Date {
	if x != nil {
		return x.RequestedWeek
	}
	return nil
}

var File_api_weeksdb_weeks_db_proto protoreflect.FileDescriptor

var file_api_weeksdb_weeks_db_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x2f, 0x77, 0x65,
	0x65, 0x6b, 0x73, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x77, 0x65,
	0x65, 0x6b, 0x73, 0x64, 0x62, 0x22, 0x42, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x59, 0x65, 0x61,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x61, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x44, 0x61, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x4a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57,
	0x65, 0x65, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x22, 0x46,
	0x0a, 0x14, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x77, 0x65,
	0x65, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x22, 0x41, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x07, 0x73, 0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x22, 0x48, 0x0a, 0x16, 0x55, 0x6e, 0x73,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x77, 0x65,
	0x65, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73,
	0x64, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57,
	0x65, 0x65, 0x6b, 0x22, 0x47, 0x0a, 0x17, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x0a, 0x75, 0x6e, 0x73, 0x65, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x09, 0x75, 0x6e, 0x73, 0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x22, 0x38, 0x0a, 0x13,
	0x49, 0x73, 0x57, 0x65, 0x65, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x04, 0x77, 0x65, 0x65, 0x6b, 0x22, 0x69, 0x0a, 0x14, 0x49, 0x73, 0x57, 0x65, 0x65, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x12, 0x34, 0x0a, 0x0e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x57, 0x65, 0x65,
	0x6b, 0x32, 0xd5, 0x02, 0x0a, 0x0d, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x57, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x1e, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x49, 0x73, 0x57, 0x65, 0x65, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x12, 0x1c, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62,
	0x2e, 0x49, 0x73, 0x57, 0x65, 0x65, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x49,
	0x73, 0x57, 0x65, 0x65, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x57, 0x65, 0x65, 0x6b, 0x12, 0x1d, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x53,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x53, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x1f, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62,
	0x2e, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57, 0x65, 0x65, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64,
	0x62, 0x2e, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x57, 0x65, 0x65,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x72, 0x64, 0x6f, 0x7a, 0x6f,
	0x73, 0x2f, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x69, 0x2d, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x2d, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x64, 0x62, 0x3b, 0x77, 0x65,
	0x65, 0x6b, 0x73, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_weeksdb_weeks_db_proto_rawDescOnce sync.Once
	file_api_weeksdb_weeks_db_proto_rawDescData = file_api_weeksdb_weeks_db_proto_rawDesc
)

func file_api_weeksdb_weeks_db_proto_rawDescGZIP() []byte {
	file_api_weeksdb_weeks_db_proto_rawDescOnce.Do(func() {
		file_api_weeksdb_weeks_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_weeksdb_weeks_db_proto_rawDescData)
	})
	return file_api_weeksdb_weeks_db_proto_rawDescData
}

var file_api_weeksdb_weeks_db_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_weeksdb_weeks_db_proto_goTypes = []interface{}{
	(*Date)(nil),                    // 0: weeksdb.Date
	(*GetStaticWeeksRequest)(nil),   // 1: weeksdb.GetStaticWeeksRequest
	(*GetStaticWeeksResponse)(nil),  // 2: weeksdb.GetStaticWeeksResponse
	(*SetStaticWeekRequest)(nil),    // 3: weeksdb.SetStaticWeekRequest
	(*SetStaticWeekResponse)(nil),   // 4: weeksdb.SetStaticWeekResponse
	(*UnsetStaticWeekRequest)(nil),  // 5: weeksdb.UnsetStaticWeekRequest
	(*UnsetStaticWeekResponse)(nil), // 6: weeksdb.UnsetStaticWeekResponse
	(*IsWeekStaticRequest)(nil),     // 7: weeksdb.IsWeekStaticRequest
	(*IsWeekStaticResponse)(nil),    // 8: weeksdb.IsWeekStaticResponse
}
var file_api_weeksdb_weeks_db_proto_depIdxs = []int32{
	0,  // 0: weeksdb.GetStaticWeeksResponse.static_weeks:type_name -> weeksdb.Date
	0,  // 1: weeksdb.SetStaticWeekRequest.static_week:type_name -> weeksdb.Date
	0,  // 2: weeksdb.SetStaticWeekResponse.set_week:type_name -> weeksdb.Date
	0,  // 3: weeksdb.UnsetStaticWeekRequest.static_week:type_name -> weeksdb.Date
	0,  // 4: weeksdb.UnsetStaticWeekResponse.unset_week:type_name -> weeksdb.Date
	0,  // 5: weeksdb.IsWeekStaticRequest.week:type_name -> weeksdb.Date
	0,  // 6: weeksdb.IsWeekStaticResponse.requested_week:type_name -> weeksdb.Date
	1,  // 7: weeksdb.WeeksDatabase.GetStaticWeeks:input_type -> weeksdb.GetStaticWeeksRequest
	7,  // 8: weeksdb.WeeksDatabase.IsWeekStatic:input_type -> weeksdb.IsWeekStaticRequest
	3,  // 9: weeksdb.WeeksDatabase.SetStaticWeek:input_type -> weeksdb.SetStaticWeekRequest
	5,  // 10: weeksdb.WeeksDatabase.UnsetStaticWeek:input_type -> weeksdb.UnsetStaticWeekRequest
	2,  // 11: weeksdb.WeeksDatabase.GetStaticWeeks:output_type -> weeksdb.GetStaticWeeksResponse
	8,  // 12: weeksdb.WeeksDatabase.IsWeekStatic:output_type -> weeksdb.IsWeekStaticResponse
	4,  // 13: weeksdb.WeeksDatabase.SetStaticWeek:output_type -> weeksdb.SetStaticWeekResponse
	6,  // 14: weeksdb.WeeksDatabase.UnsetStaticWeek:output_type -> weeksdb.UnsetStaticWeekResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_weeksdb_weeks_db_proto_init() }
func file_api_weeksdb_weeks_db_proto_init() {
	if File_api_weeksdb_weeks_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_weeksdb_weeks_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
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
		file_api_weeksdb_weeks_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaticWeeksRequest); i {
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
		file_api_weeksdb_weeks_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaticWeeksResponse); i {
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
		file_api_weeksdb_weeks_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStaticWeekRequest); i {
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
		file_api_weeksdb_weeks_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStaticWeekResponse); i {
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
		file_api_weeksdb_weeks_db_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsetStaticWeekRequest); i {
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
		file_api_weeksdb_weeks_db_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsetStaticWeekResponse); i {
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
		file_api_weeksdb_weeks_db_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsWeekStaticRequest); i {
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
		file_api_weeksdb_weeks_db_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsWeekStaticResponse); i {
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
			RawDescriptor: file_api_weeksdb_weeks_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_weeksdb_weeks_db_proto_goTypes,
		DependencyIndexes: file_api_weeksdb_weeks_db_proto_depIdxs,
		MessageInfos:      file_api_weeksdb_weeks_db_proto_msgTypes,
	}.Build()
	File_api_weeksdb_weeks_db_proto = out.File
	file_api_weeksdb_weeks_db_proto_rawDesc = nil
	file_api_weeksdb_weeks_db_proto_goTypes = nil
	file_api_weeksdb_weeks_db_proto_depIdxs = nil
}
