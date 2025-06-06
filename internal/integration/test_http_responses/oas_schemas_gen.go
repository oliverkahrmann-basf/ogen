// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

type AnyContentTypeBinaryStringSchemaDefaultDef struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s AnyContentTypeBinaryStringSchemaDefaultDef) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// AnyContentTypeBinaryStringSchemaDefaultDefStatusCodeWithHeaders wraps AnyContentTypeBinaryStringSchemaDefaultDef with status code and response headers.
type AnyContentTypeBinaryStringSchemaDefaultDefStatusCodeWithHeaders struct {
	StatusCode  int
	ContentType string
	Response    AnyContentTypeBinaryStringSchemaDefaultDef
}

// GetStatusCode returns the value of StatusCode.
func (s *AnyContentTypeBinaryStringSchemaDefaultDefStatusCodeWithHeaders) GetStatusCode() int {
	return s.StatusCode
}

// GetContentType returns the value of ContentType.
func (s *AnyContentTypeBinaryStringSchemaDefaultDefStatusCodeWithHeaders) GetContentType() string {
	return s.ContentType
}

// GetResponse returns the value of Response.
func (s *AnyContentTypeBinaryStringSchemaDefaultDefStatusCodeWithHeaders) GetResponse() AnyContentTypeBinaryStringSchemaDefaultDef {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *AnyContentTypeBinaryStringSchemaDefaultDefStatusCodeWithHeaders) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetContentType sets the value of ContentType.
func (s *AnyContentTypeBinaryStringSchemaDefaultDefStatusCodeWithHeaders) SetContentType(val string) {
	s.ContentType = val
}

// SetResponse sets the value of Response.
func (s *AnyContentTypeBinaryStringSchemaDefaultDefStatusCodeWithHeaders) SetResponse(val AnyContentTypeBinaryStringSchemaDefaultDef) {
	s.Response = val
}

type AnyContentTypeBinaryStringSchemaOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s AnyContentTypeBinaryStringSchemaOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// AnyContentTypeBinaryStringSchemaOKHeaders wraps AnyContentTypeBinaryStringSchemaOK with response headers.
type AnyContentTypeBinaryStringSchemaOKHeaders struct {
	ContentType string
	Response    AnyContentTypeBinaryStringSchemaOK
}

// GetContentType returns the value of ContentType.
func (s *AnyContentTypeBinaryStringSchemaOKHeaders) GetContentType() string {
	return s.ContentType
}

// GetResponse returns the value of Response.
func (s *AnyContentTypeBinaryStringSchemaOKHeaders) GetResponse() AnyContentTypeBinaryStringSchemaOK {
	return s.Response
}

// SetContentType sets the value of ContentType.
func (s *AnyContentTypeBinaryStringSchemaOKHeaders) SetContentType(val string) {
	s.ContentType = val
}

// SetResponse sets the value of Response.
func (s *AnyContentTypeBinaryStringSchemaOKHeaders) SetResponse(val AnyContentTypeBinaryStringSchemaOK) {
	s.Response = val
}

// Combined2XXStatusCode wraps int with StatusCode.
type Combined2XXStatusCode struct {
	StatusCode int
	Response   int
}

// GetStatusCode returns the value of StatusCode.
func (s *Combined2XXStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *Combined2XXStatusCode) GetResponse() int {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *Combined2XXStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *Combined2XXStatusCode) SetResponse(val int) {
	s.Response = val
}

func (*Combined2XXStatusCode) combinedRes() {}

// Combined5XXStatusCode wraps bool with StatusCode.
type Combined5XXStatusCode struct {
	StatusCode int
	Response   bool
}

// GetStatusCode returns the value of StatusCode.
func (s *Combined5XXStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *Combined5XXStatusCode) GetResponse() bool {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *Combined5XXStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *Combined5XXStatusCode) SetResponse(val bool) {
	s.Response = val
}

func (*Combined5XXStatusCode) combinedRes() {}

// CombinedDefStatusCode wraps []string with StatusCode.
type CombinedDefStatusCode struct {
	StatusCode int
	Response   []string
}

// GetStatusCode returns the value of StatusCode.
func (s *CombinedDefStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *CombinedDefStatusCode) GetResponse() []string {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *CombinedDefStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *CombinedDefStatusCode) SetResponse(val []string) {
	s.Response = val
}

func (*CombinedDefStatusCode) combinedRes() {}

type CombinedOK struct {
	Ok string `json:"ok"`
}

// GetOk returns the value of Ok.
func (s *CombinedOK) GetOk() string {
	return s.Ok
}

// SetOk sets the value of Ok.
func (s *CombinedOK) SetOk(val string) {
	s.Ok = val
}

func (*CombinedOK) combinedRes() {}

type CombinedType string

const (
	CombinedType200     CombinedType = "200"
	CombinedType2XX     CombinedType = "2XX"
	CombinedType5XX     CombinedType = "5XX"
	CombinedTypeDefault CombinedType = "default"
)

// AllValues returns all CombinedType values.
func (CombinedType) AllValues() []CombinedType {
	return []CombinedType{
		CombinedType200,
		CombinedType2XX,
		CombinedType5XX,
		CombinedTypeDefault,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s CombinedType) MarshalText() ([]byte, error) {
	switch s {
	case CombinedType200:
		return []byte(s), nil
	case CombinedType2XX:
		return []byte(s), nil
	case CombinedType5XX:
		return []byte(s), nil
	case CombinedTypeDefault:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CombinedType) UnmarshalText(data []byte) error {
	switch CombinedType(data) {
	case CombinedType200:
		*s = CombinedType200
		return nil
	case CombinedType2XX:
		*s = CombinedType2XX
		return nil
	case CombinedType5XX:
		*s = CombinedType5XX
		return nil
	case CombinedTypeDefault:
		*s = CombinedTypeDefault
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/Error
type Error struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *Error) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *Error) SetError(val string) {
	s.Error = val
}

func (*Error) streamJSONRes() {}

// Headers200OK is response for Headers200 operation.
type Headers200OK struct {
	XTestHeader string
}

// GetXTestHeader returns the value of XTestHeader.
func (s *Headers200OK) GetXTestHeader() string {
	return s.XTestHeader
}

// SetXTestHeader sets the value of XTestHeader.
func (s *Headers200OK) SetXTestHeader(val string) {
	s.XTestHeader = val
}

// HeadersCombined4XX is 4XX pattern response for HeadersCombined operation.
type HeadersCombined4XX struct {
	XTestHeader string
	StatusCode  int
}

// GetXTestHeader returns the value of XTestHeader.
func (s *HeadersCombined4XX) GetXTestHeader() string {
	return s.XTestHeader
}

// GetStatusCode returns the value of StatusCode.
func (s *HeadersCombined4XX) GetStatusCode() int {
	return s.StatusCode
}

// SetXTestHeader sets the value of XTestHeader.
func (s *HeadersCombined4XX) SetXTestHeader(val string) {
	s.XTestHeader = val
}

// SetStatusCode sets the value of StatusCode.
func (s *HeadersCombined4XX) SetStatusCode(val int) {
	s.StatusCode = val
}

func (*HeadersCombined4XX) headersCombinedRes() {}

// HeadersCombinedDef is default response for HeadersCombined operation.
type HeadersCombinedDef struct {
	XTestHeader string
	StatusCode  int
}

// GetXTestHeader returns the value of XTestHeader.
func (s *HeadersCombinedDef) GetXTestHeader() string {
	return s.XTestHeader
}

// GetStatusCode returns the value of StatusCode.
func (s *HeadersCombinedDef) GetStatusCode() int {
	return s.StatusCode
}

// SetXTestHeader sets the value of XTestHeader.
func (s *HeadersCombinedDef) SetXTestHeader(val string) {
	s.XTestHeader = val
}

// SetStatusCode sets the value of StatusCode.
func (s *HeadersCombinedDef) SetStatusCode(val int) {
	s.StatusCode = val
}

func (*HeadersCombinedDef) headersCombinedRes() {}

// HeadersCombinedOK is response for HeadersCombined operation.
type HeadersCombinedOK struct {
	XTestHeader string
}

// GetXTestHeader returns the value of XTestHeader.
func (s *HeadersCombinedOK) GetXTestHeader() string {
	return s.XTestHeader
}

// SetXTestHeader sets the value of XTestHeader.
func (s *HeadersCombinedOK) SetXTestHeader(val string) {
	s.XTestHeader = val
}

func (*HeadersCombinedOK) headersCombinedRes() {}

type HeadersCombinedType string

const (
	HeadersCombinedType200     HeadersCombinedType = "200"
	HeadersCombinedTypeDefault HeadersCombinedType = "default"
	HeadersCombinedType4XX     HeadersCombinedType = "4XX"
)

// AllValues returns all HeadersCombinedType values.
func (HeadersCombinedType) AllValues() []HeadersCombinedType {
	return []HeadersCombinedType{
		HeadersCombinedType200,
		HeadersCombinedTypeDefault,
		HeadersCombinedType4XX,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s HeadersCombinedType) MarshalText() ([]byte, error) {
	switch s {
	case HeadersCombinedType200:
		return []byte(s), nil
	case HeadersCombinedTypeDefault:
		return []byte(s), nil
	case HeadersCombinedType4XX:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *HeadersCombinedType) UnmarshalText(data []byte) error {
	switch HeadersCombinedType(data) {
	case HeadersCombinedType200:
		*s = HeadersCombinedType200
		return nil
	case HeadersCombinedTypeDefault:
		*s = HeadersCombinedTypeDefault
		return nil
	case HeadersCombinedType4XX:
		*s = HeadersCombinedType4XX
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// HeadersDefaultDef is default response for HeadersDefault operation.
type HeadersDefaultDef struct {
	XTestHeader string
	StatusCode  int
}

// GetXTestHeader returns the value of XTestHeader.
func (s *HeadersDefaultDef) GetXTestHeader() string {
	return s.XTestHeader
}

// GetStatusCode returns the value of StatusCode.
func (s *HeadersDefaultDef) GetStatusCode() int {
	return s.StatusCode
}

// SetXTestHeader sets the value of XTestHeader.
func (s *HeadersDefaultDef) SetXTestHeader(val string) {
	s.XTestHeader = val
}

// SetStatusCode sets the value of StatusCode.
func (s *HeadersDefaultDef) SetStatusCode(val int) {
	s.StatusCode = val
}

// HeadersJSONOK is response for HeadersJSON operation.
type HeadersJSONOK struct {
	XJSONCustomHeader jx.Raw
	XJSONHeader       User
}

// GetXJSONCustomHeader returns the value of XJSONCustomHeader.
func (s *HeadersJSONOK) GetXJSONCustomHeader() jx.Raw {
	return s.XJSONCustomHeader
}

// GetXJSONHeader returns the value of XJSONHeader.
func (s *HeadersJSONOK) GetXJSONHeader() User {
	return s.XJSONHeader
}

// SetXJSONCustomHeader sets the value of XJSONCustomHeader.
func (s *HeadersJSONOK) SetXJSONCustomHeader(val jx.Raw) {
	s.XJSONCustomHeader = val
}

// SetXJSONHeader sets the value of XJSONHeader.
func (s *HeadersJSONOK) SetXJSONHeader(val User) {
	s.XJSONHeader = val
}

// HeadersPattern4XX is 4XX pattern response for HeadersPattern operation.
type HeadersPattern4XX struct {
	XTestHeader string
	StatusCode  int
}

// GetXTestHeader returns the value of XTestHeader.
func (s *HeadersPattern4XX) GetXTestHeader() string {
	return s.XTestHeader
}

// GetStatusCode returns the value of StatusCode.
func (s *HeadersPattern4XX) GetStatusCode() int {
	return s.StatusCode
}

// SetXTestHeader sets the value of XTestHeader.
func (s *HeadersPattern4XX) SetXTestHeader(val string) {
	s.XTestHeader = val
}

// SetStatusCode sets the value of StatusCode.
func (s *HeadersPattern4XX) SetStatusCode(val int) {
	s.StatusCode = val
}

// IntersectPatternCode2XXStatusCode wraps int with StatusCode.
type IntersectPatternCode2XXStatusCode struct {
	StatusCode int
	Response   int
}

// GetStatusCode returns the value of StatusCode.
func (s *IntersectPatternCode2XXStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *IntersectPatternCode2XXStatusCode) GetResponse() int {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *IntersectPatternCode2XXStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *IntersectPatternCode2XXStatusCode) SetResponse(val int) {
	s.Response = val
}

func (*IntersectPatternCode2XXStatusCode) intersectPatternCodeRes() {}

type IntersectPatternCodeOKApplicationJSON string

func (*IntersectPatternCodeOKApplicationJSON) intersectPatternCodeRes() {}

// NewNilInt returns new NilInt with value set to v.
func NewNilInt(v int) NilInt {
	return NilInt{
		Value: v,
	}
}

// NilInt is nullable int.
type NilInt struct {
	Value int
	Null  bool
}

// SetTo sets value to v.
func (o *NilInt) SetTo(v int) {
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o NilInt) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *NilInt) SetToNull() {
	o.Null = true
	var v int
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o NilInt) Get() (v int, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

func (*NilInt) multipleGenericResponsesRes() {}

// NewNilString returns new NilString with value set to v.
func NewNilString(v string) NilString {
	return NilString{
		Value: v,
	}
}

// NilString is nullable string.
type NilString struct {
	Value string
	Null  bool
}

// SetTo sets value to v.
func (o *NilString) SetTo(v string) {
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o NilString) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *NilString) SetToNull() {
	o.Null = true
	var v string
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o NilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

func (*NilString) multipleGenericResponsesRes() {}

type OctetStreamBinaryStringSchemaOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s OctetStreamBinaryStringSchemaOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

type OctetStreamEmptySchemaOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s OctetStreamEmptySchemaOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// OptionalHeadersOK is response for OptionalHeaders operation.
type OptionalHeadersOK struct {
	XOptional OptString
	XRequired string
}

// GetXOptional returns the value of XOptional.
func (s *OptionalHeadersOK) GetXOptional() OptString {
	return s.XOptional
}

// GetXRequired returns the value of XRequired.
func (s *OptionalHeadersOK) GetXRequired() string {
	return s.XRequired
}

// SetXOptional sets the value of XOptional.
func (s *OptionalHeadersOK) SetXOptional(val OptString) {
	s.XOptional = val
}

// SetXRequired sets the value of XRequired.
func (s *OptionalHeadersOK) SetXRequired(val string) {
	s.XRequired = val
}

type QueryData []float64

func (*QueryData) streamJSONRes() {}

type TextPlainBinaryStringSchemaOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s TextPlainBinaryStringSchemaOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// Ref: #/components/schemas/User
type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Role     UserRole `json:"role"`
	Friends  []User   `json:"friends"`
}

// GetID returns the value of ID.
func (s *User) GetID() int {
	return s.ID
}

// GetUsername returns the value of Username.
func (s *User) GetUsername() string {
	return s.Username
}

// GetRole returns the value of Role.
func (s *User) GetRole() UserRole {
	return s.Role
}

// GetFriends returns the value of Friends.
func (s *User) GetFriends() []User {
	return s.Friends
}

// SetID sets the value of ID.
func (s *User) SetID(val int) {
	s.ID = val
}

// SetUsername sets the value of Username.
func (s *User) SetUsername(val string) {
	s.Username = val
}

// SetRole sets the value of Role.
func (s *User) SetRole(val UserRole) {
	s.Role = val
}

// SetFriends sets the value of Friends.
func (s *User) SetFriends(val []User) {
	s.Friends = val
}

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
	UserRoleBot   UserRole = "bot"
)

// AllValues returns all UserRole values.
func (UserRole) AllValues() []UserRole {
	return []UserRole{
		UserRoleAdmin,
		UserRoleUser,
		UserRoleBot,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s UserRole) MarshalText() ([]byte, error) {
	switch s {
	case UserRoleAdmin:
		return []byte(s), nil
	case UserRoleUser:
		return []byte(s), nil
	case UserRoleBot:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *UserRole) UnmarshalText(data []byte) error {
	switch UserRole(data) {
	case UserRoleAdmin:
		*s = UserRoleAdmin
		return nil
	case UserRoleUser:
		*s = UserRoleUser
		return nil
	case UserRoleBot:
		*s = UserRoleBot
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}
