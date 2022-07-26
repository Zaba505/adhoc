// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: games/tictactoe/tictactoe.proto

package tictactoepb

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

type GameMode int32

const (
	GameMode_PVP GameMode = 0
	GameMode_PVC GameMode = 1
)

// Enum value maps for GameMode.
var (
	GameMode_name = map[int32]string{
		0: "PVP",
		1: "PVC",
	}
	GameMode_value = map[string]int32{
		"PVP": 0,
		"PVC": 1,
	}
)

func (x GameMode) Enum() *GameMode {
	p := new(GameMode)
	*p = x
	return p
}

func (x GameMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameMode) Descriptor() protoreflect.EnumDescriptor {
	return file_games_tictactoe_tictactoe_proto_enumTypes[0].Descriptor()
}

func (GameMode) Type() protoreflect.EnumType {
	return &file_games_tictactoe_tictactoe_proto_enumTypes[0]
}

func (x GameMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameMode.Descriptor instead.
func (GameMode) EnumDescriptor() ([]byte, []int) {
	return file_games_tictactoe_tictactoe_proto_rawDescGZIP(), []int{0}
}

type Piece int32

const (
	Piece_O Piece = 0
	Piece_X Piece = 1
)

// Enum value maps for Piece.
var (
	Piece_name = map[int32]string{
		0: "O",
		1: "X",
	}
	Piece_value = map[string]int32{
		"O": 0,
		"X": 1,
	}
)

func (x Piece) Enum() *Piece {
	p := new(Piece)
	*p = x
	return p
}

func (x Piece) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Piece) Descriptor() protoreflect.EnumDescriptor {
	return file_games_tictactoe_tictactoe_proto_enumTypes[1].Descriptor()
}

func (Piece) Type() protoreflect.EnumType {
	return &file_games_tictactoe_tictactoe_proto_enumTypes[1]
}

func (x Piece) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Piece.Descriptor instead.
func (Piece) EnumDescriptor() ([]byte, []int) {
	return file_games_tictactoe_tictactoe_proto_rawDescGZIP(), []int{1}
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameMode GameMode `protobuf:"varint,1,opt,name=game_mode,json=gameMode,proto3,enum=tictactoepb.GameMode" json:"game_mode,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_tictactoe_tictactoe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_tictactoe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_tictactoe_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetGameMode() GameMode {
	if x != nil {
		return x.GameMode
	}
	return GameMode_PVP
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_tictactoe_tictactoe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_tictactoe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_tictactoe_proto_rawDescGZIP(), []int{1}
}

type ClientGameEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*ClientGameEvent_PiecePlaced_
	Event isClientGameEvent_Event `protobuf_oneof:"event"`
}

func (x *ClientGameEvent) Reset() {
	*x = ClientGameEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_tictactoe_tictactoe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientGameEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientGameEvent) ProtoMessage() {}

func (x *ClientGameEvent) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_tictactoe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientGameEvent.ProtoReflect.Descriptor instead.
func (*ClientGameEvent) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_tictactoe_proto_rawDescGZIP(), []int{2}
}

func (m *ClientGameEvent) GetEvent() isClientGameEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *ClientGameEvent) GetPiecePlaced() *ClientGameEvent_PiecePlaced {
	if x, ok := x.GetEvent().(*ClientGameEvent_PiecePlaced_); ok {
		return x.PiecePlaced
	}
	return nil
}

type isClientGameEvent_Event interface {
	isClientGameEvent_Event()
}

type ClientGameEvent_PiecePlaced_ struct {
	PiecePlaced *ClientGameEvent_PiecePlaced `protobuf:"bytes,1,opt,name=piece_placed,json=piecePlaced,proto3,oneof"`
}

func (*ClientGameEvent_PiecePlaced_) isClientGameEvent_Event() {}

type ServerGameEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*ServerGameEvent_GameStarted_
	//	*ServerGameEvent_PiecePlaced_
	//	*ServerGameEvent_PieceAssignment_
	//	*ServerGameEvent_GameOver_
	Event isServerGameEvent_Event `protobuf_oneof:"event"`
}

func (x *ServerGameEvent) Reset() {
	*x = ServerGameEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_tictactoe_tictactoe_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerGameEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerGameEvent) ProtoMessage() {}

func (x *ServerGameEvent) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_tictactoe_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerGameEvent.ProtoReflect.Descriptor instead.
func (*ServerGameEvent) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_tictactoe_proto_rawDescGZIP(), []int{3}
}

func (m *ServerGameEvent) GetEvent() isServerGameEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *ServerGameEvent) GetGameStarted() *ServerGameEvent_GameStarted {
	if x, ok := x.GetEvent().(*ServerGameEvent_GameStarted_); ok {
		return x.GameStarted
	}
	return nil
}

func (x *ServerGameEvent) GetPiecePlaced() *ServerGameEvent_PiecePlaced {
	if x, ok := x.GetEvent().(*ServerGameEvent_PiecePlaced_); ok {
		return x.PiecePlaced
	}
	return nil
}

func (x *ServerGameEvent) GetPieceAssignment() *ServerGameEvent_PieceAssignment {
	if x, ok := x.GetEvent().(*ServerGameEvent_PieceAssignment_); ok {
		return x.PieceAssignment
	}
	return nil
}

func (x *ServerGameEvent) GetGameOver() *ServerGameEvent_GameOver {
	if x, ok := x.GetEvent().(*ServerGameEvent_GameOver_); ok {
		return x.GameOver
	}
	return nil
}

type isServerGameEvent_Event interface {
	isServerGameEvent_Event()
}

type ServerGameEvent_GameStarted_ struct {
	// Game started
	GameStarted *ServerGameEvent_GameStarted `protobuf:"bytes,1,opt,name=game_started,json=gameStarted,proto3,oneof"`
}

type ServerGameEvent_PiecePlaced_ struct {
	// The other player has placed a piece on the board
	PiecePlaced *ServerGameEvent_PiecePlaced `protobuf:"bytes,2,opt,name=piece_placed,json=piecePlaced,proto3,oneof"`
}

type ServerGameEvent_PieceAssignment_ struct {
	// You have been assigned a piece by the TicTacToe service
	PieceAssignment *ServerGameEvent_PieceAssignment `protobuf:"bytes,3,opt,name=piece_assignment,json=pieceAssignment,proto3,oneof"`
}

type ServerGameEvent_GameOver_ struct {
	// Someone has won the game or it is tied
	GameOver *ServerGameEvent_GameOver `protobuf:"bytes,4,opt,name=game_over,json=gameOver,proto3,oneof"`
}

func (*ServerGameEvent_GameStarted_) isServerGameEvent_Event() {}

func (*ServerGameEvent_PiecePlaced_) isServerGameEvent_Event() {}

func (*ServerGameEvent_PieceAssignment_) isServerGameEvent_Event() {}

func (*ServerGameEvent_GameOver_) isServerGameEvent_Event() {}

type ClientGameEvent_PiecePlaced struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// row must be between 0 and 2, inclusively
	Row int64 `protobuf:"varint,1,opt,name=row,proto3" json:"row,omitempty"`
	// column must be between 0 and 2, inclusively
	Column int64 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *ClientGameEvent_PiecePlaced) Reset() {
	*x = ClientGameEvent_PiecePlaced{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_tictactoe_tictactoe_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientGameEvent_PiecePlaced) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientGameEvent_PiecePlaced) ProtoMessage() {}

func (x *ClientGameEvent_PiecePlaced) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_tictactoe_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientGameEvent_PiecePlaced.ProtoReflect.Descriptor instead.
func (*ClientGameEvent_PiecePlaced) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_tictactoe_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ClientGameEvent_PiecePlaced) GetRow() int64 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *ClientGameEvent_PiecePlaced) GetColumn() int64 {
	if x != nil {
		return x.Column
	}
	return 0
}

type ServerGameEvent_GameStarted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartingPiece Piece `protobuf:"varint,1,opt,name=starting_piece,json=startingPiece,proto3,enum=tictactoepb.Piece" json:"starting_piece,omitempty"`
}

func (x *ServerGameEvent_GameStarted) Reset() {
	*x = ServerGameEvent_GameStarted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_tictactoe_tictactoe_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerGameEvent_GameStarted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerGameEvent_GameStarted) ProtoMessage() {}

func (x *ServerGameEvent_GameStarted) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_tictactoe_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerGameEvent_GameStarted.ProtoReflect.Descriptor instead.
func (*ServerGameEvent_GameStarted) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_tictactoe_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ServerGameEvent_GameStarted) GetStartingPiece() Piece {
	if x != nil {
		return x.StartingPiece
	}
	return Piece_O
}

type ServerGameEvent_PiecePlaced struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// row must be between 0 and 2, inclusively
	Row int64 `protobuf:"varint,1,opt,name=row,proto3" json:"row,omitempty"`
	// column must be between 0 and 2, inclusively
	Column int64 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	// piece is the piece which was placed on the board
	Piece Piece `protobuf:"varint,3,opt,name=piece,proto3,enum=tictactoepb.Piece" json:"piece,omitempty"`
}

func (x *ServerGameEvent_PiecePlaced) Reset() {
	*x = ServerGameEvent_PiecePlaced{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_tictactoe_tictactoe_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerGameEvent_PiecePlaced) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerGameEvent_PiecePlaced) ProtoMessage() {}

func (x *ServerGameEvent_PiecePlaced) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_tictactoe_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerGameEvent_PiecePlaced.ProtoReflect.Descriptor instead.
func (*ServerGameEvent_PiecePlaced) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_tictactoe_proto_rawDescGZIP(), []int{3, 1}
}

func (x *ServerGameEvent_PiecePlaced) GetRow() int64 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *ServerGameEvent_PiecePlaced) GetColumn() int64 {
	if x != nil {
		return x.Column
	}
	return 0
}

func (x *ServerGameEvent_PiecePlaced) GetPiece() Piece {
	if x != nil {
		return x.Piece
	}
	return Piece_O
}

type ServerGameEvent_PieceAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Piece Piece `protobuf:"varint,1,opt,name=piece,proto3,enum=tictactoepb.Piece" json:"piece,omitempty"`
}

func (x *ServerGameEvent_PieceAssignment) Reset() {
	*x = ServerGameEvent_PieceAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_tictactoe_tictactoe_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerGameEvent_PieceAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerGameEvent_PieceAssignment) ProtoMessage() {}

func (x *ServerGameEvent_PieceAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_tictactoe_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerGameEvent_PieceAssignment.ProtoReflect.Descriptor instead.
func (*ServerGameEvent_PieceAssignment) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_tictactoe_proto_rawDescGZIP(), []int{3, 2}
}

func (x *ServerGameEvent_PieceAssignment) GetPiece() Piece {
	if x != nil {
		return x.Piece
	}
	return Piece_O
}

type ServerGameEvent_GameOver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Winner Piece `protobuf:"varint,1,opt,name=winner,proto3,enum=tictactoepb.Piece" json:"winner,omitempty"`
}

func (x *ServerGameEvent_GameOver) Reset() {
	*x = ServerGameEvent_GameOver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_tictactoe_tictactoe_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerGameEvent_GameOver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerGameEvent_GameOver) ProtoMessage() {}

func (x *ServerGameEvent_GameOver) ProtoReflect() protoreflect.Message {
	mi := &file_games_tictactoe_tictactoe_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerGameEvent_GameOver.ProtoReflect.Descriptor instead.
func (*ServerGameEvent_GameOver) Descriptor() ([]byte, []int) {
	return file_games_tictactoe_tictactoe_proto_rawDescGZIP(), []int{3, 3}
}

func (x *ServerGameEvent_GameOver) GetWinner() Piece {
	if x != nil {
		return x.Winner
	}
	return Piece_O
}

var File_games_tictactoe_tictactoe_proto protoreflect.FileDescriptor

var file_games_tictactoe_tictactoe_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f,
	0x65, 0x2f, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x22, 0x42,
	0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x0c, 0x70, 0x69, 0x65, 0x63, 0x65,
	0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x69, 0x65, 0x63,
	0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x69, 0x65, 0x63, 0x65,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x1a, 0x37, 0x0a, 0x0b, 0x50, 0x69, 0x65, 0x63, 0x65, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x42,
	0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xfb, 0x04, 0x0a, 0x0f, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x0c,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0b,
	0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x0c, 0x70,
	0x69, 0x65, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x50, 0x69, 0x65, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x70,
	0x69, 0x65, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x12, 0x59, 0x0a, 0x10, 0x70, 0x69,
	0x65, 0x63, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x69, 0x65, 0x63, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6f, 0x76,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61,
	0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x61, 0x6d,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x1a, 0x48, 0x0a, 0x0b, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0e, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x69, 0x65, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62,
	0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x50, 0x69, 0x65, 0x63, 0x65, 0x1a, 0x61, 0x0a, 0x0b, 0x50, 0x69, 0x65, 0x63, 0x65, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x28,
	0x0a, 0x05, 0x70, 0x69, 0x65, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x65, 0x63,
	0x65, 0x52, 0x05, 0x70, 0x69, 0x65, 0x63, 0x65, 0x1a, 0x3b, 0x0a, 0x0f, 0x50, 0x69, 0x65, 0x63,
	0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x70,
	0x69, 0x65, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x69, 0x63,
	0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x52, 0x05,
	0x70, 0x69, 0x65, 0x63, 0x65, 0x1a, 0x36, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65,
	0x72, 0x12, 0x2a, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x2e,
	0x50, 0x69, 0x65, 0x63, 0x65, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x42, 0x07, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2a, 0x1c, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x56, 0x50, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50,
	0x56, 0x43, 0x10, 0x01, 0x2a, 0x15, 0x0a, 0x05, 0x50, 0x69, 0x65, 0x63, 0x65, 0x12, 0x05, 0x0a,
	0x01, 0x4f, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x58, 0x10, 0x01, 0x32, 0x93, 0x01, 0x0a, 0x09,
	0x54, 0x69, 0x63, 0x54, 0x61, 0x63, 0x54, 0x6f, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x19, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x4a, 0x6f, 0x69,
	0x6e, 0x12, 0x1c, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a,
	0x1c, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x5a, 0x61, 0x62, 0x61, 0x35, 0x30, 0x35, 0x2f, 0x61, 0x64, 0x68, 0x6f, 0x63, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2f, 0x74, 0x69,
	0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_games_tictactoe_tictactoe_proto_rawDescOnce sync.Once
	file_games_tictactoe_tictactoe_proto_rawDescData = file_games_tictactoe_tictactoe_proto_rawDesc
)

func file_games_tictactoe_tictactoe_proto_rawDescGZIP() []byte {
	file_games_tictactoe_tictactoe_proto_rawDescOnce.Do(func() {
		file_games_tictactoe_tictactoe_proto_rawDescData = protoimpl.X.CompressGZIP(file_games_tictactoe_tictactoe_proto_rawDescData)
	})
	return file_games_tictactoe_tictactoe_proto_rawDescData
}

var file_games_tictactoe_tictactoe_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_games_tictactoe_tictactoe_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_games_tictactoe_tictactoe_proto_goTypes = []interface{}{
	(GameMode)(0),                           // 0: tictactoepb.GameMode
	(Piece)(0),                              // 1: tictactoepb.Piece
	(*StartRequest)(nil),                    // 2: tictactoepb.StartRequest
	(*StartResponse)(nil),                   // 3: tictactoepb.StartResponse
	(*ClientGameEvent)(nil),                 // 4: tictactoepb.ClientGameEvent
	(*ServerGameEvent)(nil),                 // 5: tictactoepb.ServerGameEvent
	(*ClientGameEvent_PiecePlaced)(nil),     // 6: tictactoepb.ClientGameEvent.PiecePlaced
	(*ServerGameEvent_GameStarted)(nil),     // 7: tictactoepb.ServerGameEvent.GameStarted
	(*ServerGameEvent_PiecePlaced)(nil),     // 8: tictactoepb.ServerGameEvent.PiecePlaced
	(*ServerGameEvent_PieceAssignment)(nil), // 9: tictactoepb.ServerGameEvent.PieceAssignment
	(*ServerGameEvent_GameOver)(nil),        // 10: tictactoepb.ServerGameEvent.GameOver
}
var file_games_tictactoe_tictactoe_proto_depIdxs = []int32{
	0,  // 0: tictactoepb.StartRequest.game_mode:type_name -> tictactoepb.GameMode
	6,  // 1: tictactoepb.ClientGameEvent.piece_placed:type_name -> tictactoepb.ClientGameEvent.PiecePlaced
	7,  // 2: tictactoepb.ServerGameEvent.game_started:type_name -> tictactoepb.ServerGameEvent.GameStarted
	8,  // 3: tictactoepb.ServerGameEvent.piece_placed:type_name -> tictactoepb.ServerGameEvent.PiecePlaced
	9,  // 4: tictactoepb.ServerGameEvent.piece_assignment:type_name -> tictactoepb.ServerGameEvent.PieceAssignment
	10, // 5: tictactoepb.ServerGameEvent.game_over:type_name -> tictactoepb.ServerGameEvent.GameOver
	1,  // 6: tictactoepb.ServerGameEvent.GameStarted.starting_piece:type_name -> tictactoepb.Piece
	1,  // 7: tictactoepb.ServerGameEvent.PiecePlaced.piece:type_name -> tictactoepb.Piece
	1,  // 8: tictactoepb.ServerGameEvent.PieceAssignment.piece:type_name -> tictactoepb.Piece
	1,  // 9: tictactoepb.ServerGameEvent.GameOver.winner:type_name -> tictactoepb.Piece
	2,  // 10: tictactoepb.TicTacToe.Start:input_type -> tictactoepb.StartRequest
	4,  // 11: tictactoepb.TicTacToe.Join:input_type -> tictactoepb.ClientGameEvent
	3,  // 12: tictactoepb.TicTacToe.Start:output_type -> tictactoepb.StartResponse
	5,  // 13: tictactoepb.TicTacToe.Join:output_type -> tictactoepb.ServerGameEvent
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_games_tictactoe_tictactoe_proto_init() }
func file_games_tictactoe_tictactoe_proto_init() {
	if File_games_tictactoe_tictactoe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_games_tictactoe_tictactoe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_games_tictactoe_tictactoe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
		file_games_tictactoe_tictactoe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientGameEvent); i {
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
		file_games_tictactoe_tictactoe_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerGameEvent); i {
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
		file_games_tictactoe_tictactoe_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientGameEvent_PiecePlaced); i {
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
		file_games_tictactoe_tictactoe_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerGameEvent_GameStarted); i {
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
		file_games_tictactoe_tictactoe_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerGameEvent_PiecePlaced); i {
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
		file_games_tictactoe_tictactoe_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerGameEvent_PieceAssignment); i {
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
		file_games_tictactoe_tictactoe_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerGameEvent_GameOver); i {
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
	file_games_tictactoe_tictactoe_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ClientGameEvent_PiecePlaced_)(nil),
	}
	file_games_tictactoe_tictactoe_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ServerGameEvent_GameStarted_)(nil),
		(*ServerGameEvent_PiecePlaced_)(nil),
		(*ServerGameEvent_PieceAssignment_)(nil),
		(*ServerGameEvent_GameOver_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_games_tictactoe_tictactoe_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_games_tictactoe_tictactoe_proto_goTypes,
		DependencyIndexes: file_games_tictactoe_tictactoe_proto_depIdxs,
		EnumInfos:         file_games_tictactoe_tictactoe_proto_enumTypes,
		MessageInfos:      file_games_tictactoe_tictactoe_proto_msgTypes,
	}.Build()
	File_games_tictactoe_tictactoe_proto = out.File
	file_games_tictactoe_tictactoe_proto_rawDesc = nil
	file_games_tictactoe_tictactoe_proto_goTypes = nil
	file_games_tictactoe_tictactoe_proto_depIdxs = nil
}
