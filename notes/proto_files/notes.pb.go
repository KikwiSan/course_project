// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: notes.proto

package notes

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

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NoteContent string `protobuf:"bytes,3,opt,name=note_content,json=noteContent,proto3" json:"note_content,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Note) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Note) GetNoteContent() string {
	if x != nil {
		return x.NoteContent
	}
	return ""
}

type GetNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNoteRequest) Reset() {
	*x = GetNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteRequest) ProtoMessage() {}

func (x *GetNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteRequest.ProtoReflect.Descriptor instead.
func (*GetNoteRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{1}
}

func (x *GetNoteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetNoteResponse) Reset() {
	*x = GetNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteResponse) ProtoMessage() {}

func (x *GetNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteResponse.ProtoReflect.Descriptor instead.
func (*GetNoteResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{2}
}

func (x *GetNoteResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note *Note `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *CreateNoteResponse) Reset() {
	*x = CreateNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteResponse) ProtoMessage() {}

func (x *CreateNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteResponse.ProtoReflect.Descriptor instead.
func (*CreateNoteResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNoteResponse) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type GetAllNotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllNotesRequest) Reset() {
	*x = GetAllNotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNotesRequest) ProtoMessage() {}

func (x *GetAllNotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNotesRequest.ProtoReflect.Descriptor instead.
func (*GetAllNotesRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{4}
}

type GetAllNotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllNotes []*Note `protobuf:"bytes,1,rep,name=allNotes,proto3" json:"allNotes,omitempty"`
}

func (x *GetAllNotesResponse) Reset() {
	*x = GetAllNotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNotesResponse) ProtoMessage() {}

func (x *GetAllNotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNotesResponse.ProtoReflect.Descriptor instead.
func (*GetAllNotesResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllNotesResponse) GetAllNotes() []*Note {
	if x != nil {
		return x.AllNotes
	}
	return nil
}

type GetNotesByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetNotesByNameRequest) Reset() {
	*x = GetNotesByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotesByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotesByNameRequest) ProtoMessage() {}

func (x *GetNotesByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotesByNameRequest.ProtoReflect.Descriptor instead.
func (*GetNotesByNameRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{6}
}

func (x *GetNotesByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetNotesByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notes []*Note `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *GetNotesByNameResponse) Reset() {
	*x = GetNotesByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotesByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotesByNameResponse) ProtoMessage() {}

func (x *GetNotesByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotesByNameResponse.ProtoReflect.Descriptor instead.
func (*GetNotesByNameResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{7}
}

func (x *GetNotesByNameResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

var File_notes_proto protoreflect.FileDescriptor

var file_notes_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x08, 0x61, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x08, 0x61, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74,
	0x65, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x32, 0x8e, 0x02, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0b,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x1a, 0x19, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x15, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x65, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65,
	0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notes_proto_rawDescOnce sync.Once
	file_notes_proto_rawDescData = file_notes_proto_rawDesc
)

func file_notes_proto_rawDescGZIP() []byte {
	file_notes_proto_rawDescOnce.Do(func() {
		file_notes_proto_rawDescData = protoimpl.X.CompressGZIP(file_notes_proto_rawDescData)
	})
	return file_notes_proto_rawDescData
}

var file_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_notes_proto_goTypes = []interface{}{
	(*Note)(nil),                   // 0: notes.Note
	(*GetNoteRequest)(nil),         // 1: notes.GetNoteRequest
	(*GetNoteResponse)(nil),        // 2: notes.GetNoteResponse
	(*CreateNoteResponse)(nil),     // 3: notes.CreateNoteResponse
	(*GetAllNotesRequest)(nil),     // 4: notes.GetAllNotesRequest
	(*GetAllNotesResponse)(nil),    // 5: notes.GetAllNotesResponse
	(*GetNotesByNameRequest)(nil),  // 6: notes.GetNotesByNameRequest
	(*GetNotesByNameResponse)(nil), // 7: notes.GetNotesByNameResponse
}
var file_notes_proto_depIdxs = []int32{
	0, // 0: notes.CreateNoteResponse.note:type_name -> notes.Note
	0, // 1: notes.GetAllNotesResponse.allNotes:type_name -> notes.Note
	0, // 2: notes.GetNotesByNameResponse.notes:type_name -> notes.Note
	0, // 3: notes.NotesService.Create:input_type -> notes.Note
	1, // 4: notes.NotesService.Get:input_type -> notes.GetNoteRequest
	6, // 5: notes.NotesService.GetNotesByName:input_type -> notes.GetNotesByNameRequest
	4, // 6: notes.NotesService.GetAll:input_type -> notes.GetAllNotesRequest
	3, // 7: notes.NotesService.Create:output_type -> notes.CreateNoteResponse
	2, // 8: notes.NotesService.Get:output_type -> notes.GetNoteResponse
	7, // 9: notes.NotesService.GetNotesByName:output_type -> notes.GetNotesByNameResponse
	5, // 10: notes.NotesService.GetAll:output_type -> notes.GetAllNotesResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_notes_proto_init() }
func file_notes_proto_init() {
	if File_notes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_notes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoteRequest); i {
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
		file_notes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoteResponse); i {
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
		file_notes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNoteResponse); i {
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
		file_notes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllNotesRequest); i {
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
		file_notes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllNotesResponse); i {
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
		file_notes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotesByNameRequest); i {
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
		file_notes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotesByNameResponse); i {
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
			RawDescriptor: file_notes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notes_proto_goTypes,
		DependencyIndexes: file_notes_proto_depIdxs,
		MessageInfos:      file_notes_proto_msgTypes,
	}.Build()
	File_notes_proto = out.File
	file_notes_proto_rawDesc = nil
	file_notes_proto_goTypes = nil
	file_notes_proto_depIdxs = nil
}
