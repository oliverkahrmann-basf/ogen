// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
)

// Ref: #/components/schemas/Bird
type Bird struct {
	PetType string  `json:"petType"`
	Name    string  `json:"name"`
	Species string  `json:"species"`
	CanTalk OptBool `json:"canTalk"`
}

// GetPetType returns the value of PetType.
func (s *Bird) GetPetType() string {
	return s.PetType
}

// GetName returns the value of Name.
func (s *Bird) GetName() string {
	return s.Name
}

// GetSpecies returns the value of Species.
func (s *Bird) GetSpecies() string {
	return s.Species
}

// GetCanTalk returns the value of CanTalk.
func (s *Bird) GetCanTalk() OptBool {
	return s.CanTalk
}

// SetPetType sets the value of PetType.
func (s *Bird) SetPetType(val string) {
	s.PetType = val
}

// SetName sets the value of Name.
func (s *Bird) SetName(val string) {
	s.Name = val
}

// SetSpecies sets the value of Species.
func (s *Bird) SetSpecies(val string) {
	s.Species = val
}

// SetCanTalk sets the value of CanTalk.
func (s *Bird) SetCanTalk(val OptBool) {
	s.CanTalk = val
}

// Ref: #/components/schemas/Car
type Car struct {
	VehicleType string `json:"vehicleType"`
	Make        string `json:"make"`
	Model       string `json:"model"`
	Doors       OptInt `json:"doors"`
}

// GetVehicleType returns the value of VehicleType.
func (s *Car) GetVehicleType() string {
	return s.VehicleType
}

// GetMake returns the value of Make.
func (s *Car) GetMake() string {
	return s.Make
}

// GetModel returns the value of Model.
func (s *Car) GetModel() string {
	return s.Model
}

// GetDoors returns the value of Doors.
func (s *Car) GetDoors() OptInt {
	return s.Doors
}

// SetVehicleType sets the value of VehicleType.
func (s *Car) SetVehicleType(val string) {
	s.VehicleType = val
}

// SetMake sets the value of Make.
func (s *Car) SetMake(val string) {
	s.Make = val
}

// SetModel sets the value of Model.
func (s *Car) SetModel(val string) {
	s.Model = val
}

// SetDoors sets the value of Doors.
func (s *Car) SetDoors(val OptInt) {
	s.Doors = val
}

// Ref: #/components/schemas/Cat
type Cat struct {
	PetType    string `json:"petType"`
	Name       string `json:"name"`
	Breed      string `json:"breed"`
	MeowVolume OptInt `json:"meowVolume"`
}

// GetPetType returns the value of PetType.
func (s *Cat) GetPetType() string {
	return s.PetType
}

// GetName returns the value of Name.
func (s *Cat) GetName() string {
	return s.Name
}

// GetBreed returns the value of Breed.
func (s *Cat) GetBreed() string {
	return s.Breed
}

// GetMeowVolume returns the value of MeowVolume.
func (s *Cat) GetMeowVolume() OptInt {
	return s.MeowVolume
}

// SetPetType sets the value of PetType.
func (s *Cat) SetPetType(val string) {
	s.PetType = val
}

// SetName sets the value of Name.
func (s *Cat) SetName(val string) {
	s.Name = val
}

// SetBreed sets the value of Breed.
func (s *Cat) SetBreed(val string) {
	s.Breed = val
}

// SetMeowVolume sets the value of MeowVolume.
func (s *Cat) SetMeowVolume(val OptInt) {
	s.MeowVolume = val
}

// Ref: #/components/schemas/Dog
type Dog struct {
	PetType      string `json:"petType"`
	Name         string `json:"name"`
	Breed        string `json:"breed"`
	BarkLoudness OptInt `json:"barkLoudness"`
}

// GetPetType returns the value of PetType.
func (s *Dog) GetPetType() string {
	return s.PetType
}

// GetName returns the value of Name.
func (s *Dog) GetName() string {
	return s.Name
}

// GetBreed returns the value of Breed.
func (s *Dog) GetBreed() string {
	return s.Breed
}

// GetBarkLoudness returns the value of BarkLoudness.
func (s *Dog) GetBarkLoudness() OptInt {
	return s.BarkLoudness
}

// SetPetType sets the value of PetType.
func (s *Dog) SetPetType(val string) {
	s.PetType = val
}

// SetName sets the value of Name.
func (s *Dog) SetName(val string) {
	s.Name = val
}

// SetBreed sets the value of Breed.
func (s *Dog) SetBreed(val string) {
	s.Breed = val
}

// SetBarkLoudness sets the value of BarkLoudness.
func (s *Dog) SetBarkLoudness(val OptInt) {
	s.BarkLoudness = val
}

// Ref: #/components/schemas/EmailNotification
type EmailNotification struct {
	NotificationType string    `json:"notificationType"`
	Recipient        string    `json:"recipient"`
	Subject          string    `json:"subject"`
	Body             OptString `json:"body"`
}

// GetNotificationType returns the value of NotificationType.
func (s *EmailNotification) GetNotificationType() string {
	return s.NotificationType
}

// GetRecipient returns the value of Recipient.
func (s *EmailNotification) GetRecipient() string {
	return s.Recipient
}

// GetSubject returns the value of Subject.
func (s *EmailNotification) GetSubject() string {
	return s.Subject
}

// GetBody returns the value of Body.
func (s *EmailNotification) GetBody() OptString {
	return s.Body
}

// SetNotificationType sets the value of NotificationType.
func (s *EmailNotification) SetNotificationType(val string) {
	s.NotificationType = val
}

// SetRecipient sets the value of Recipient.
func (s *EmailNotification) SetRecipient(val string) {
	s.Recipient = val
}

// SetSubject sets the value of Subject.
func (s *EmailNotification) SetSubject(val string) {
	s.Subject = val
}

// SetBody sets the value of Body.
func (s *EmailNotification) SetBody(val OptString) {
	s.Body = val
}

// Ref: #/components/schemas/Motorcycle
type Motorcycle struct {
	VehicleType string     `json:"vehicleType"`
	Make        string     `json:"make"`
	Model       string     `json:"model"`
	EngineSize  OptFloat64 `json:"engineSize"`
}

// GetVehicleType returns the value of VehicleType.
func (s *Motorcycle) GetVehicleType() string {
	return s.VehicleType
}

// GetMake returns the value of Make.
func (s *Motorcycle) GetMake() string {
	return s.Make
}

// GetModel returns the value of Model.
func (s *Motorcycle) GetModel() string {
	return s.Model
}

// GetEngineSize returns the value of EngineSize.
func (s *Motorcycle) GetEngineSize() OptFloat64 {
	return s.EngineSize
}

// SetVehicleType sets the value of VehicleType.
func (s *Motorcycle) SetVehicleType(val string) {
	s.VehicleType = val
}

// SetMake sets the value of Make.
func (s *Motorcycle) SetMake(val string) {
	s.Make = val
}

// SetModel sets the value of Model.
func (s *Motorcycle) SetModel(val string) {
	s.Model = val
}

// SetEngineSize sets the value of EngineSize.
func (s *Motorcycle) SetEngineSize(val OptFloat64) {
	s.EngineSize = val
}

// Ref: #/components/schemas/Notification
// Notification represents sum type.
type Notification struct {
	Type              NotificationType // switch on this field
	EmailNotification EmailNotification
	SMSNotification   SMSNotification
	PushNotification  PushNotification
}

// NotificationType is oneOf type of Notification.
type NotificationType string

// Possible values for NotificationType.
const (
	NotificationEmailNotification  NotificationType = "email"
	NotificationMailNotification   NotificationType = "mail"
	NotificationSMSNotification    NotificationType = "sms"
	NotificationTextNotification   NotificationType = "text"
	NotificationMobileNotification NotificationType = "mobile"
	NotificationPushNotification   NotificationType = "push"
)

// IsEmailNotification reports whether Notification is EmailNotification.
func (s Notification) IsEmailNotification() bool {
	switch s.Type {
	case NotificationEmailNotification, NotificationMailNotification:
		return true
	default:
		return false
	}
}

// IsSMSNotification reports whether Notification is SMSNotification.
func (s Notification) IsSMSNotification() bool {
	switch s.Type {
	case NotificationSMSNotification, NotificationTextNotification:
		return true
	default:
		return false
	}
}

// IsPushNotification reports whether Notification is PushNotification.
func (s Notification) IsPushNotification() bool {
	switch s.Type {
	case NotificationMobileNotification, NotificationPushNotification:
		return true
	default:
		return false
	}
}

// SetEmailNotification sets Notification to EmailNotification.
// panics if `t` is not associated with EmailNotification
func (s *Notification) SetEmailNotification(t NotificationType, v EmailNotification) {
	s.Type = t
	s.EmailNotification = v
	if !s.IsEmailNotification() {
		panic(fmt.Errorf("invariant: %v is not EmailNotification", t))
	}
}

// GetEmailNotification returns EmailNotification and true boolean if Notification is EmailNotification.
func (s Notification) GetEmailNotification() (v EmailNotification, ok bool) {
	if !s.IsEmailNotification() {
		return v, false
	}
	return s.EmailNotification, true
}

// NewNotificationEmailNotification returns new Notification from EmailNotification.
func NewNotificationEmailNotification(v EmailNotification) Notification {
	var s Notification
	s.SetEmailNotification(NotificationEmailNotification, v)
	return s
}

// NewNotificationMailNotification returns new Notification from EmailNotification.
func NewNotificationMailNotification(v EmailNotification) Notification {
	var s Notification
	s.SetEmailNotification(NotificationMailNotification, v)
	return s
}

// SetSMSNotification sets Notification to SMSNotification.
// panics if `t` is not associated with SMSNotification
func (s *Notification) SetSMSNotification(t NotificationType, v SMSNotification) {
	s.Type = t
	s.SMSNotification = v
	if !s.IsSMSNotification() {
		panic(fmt.Errorf("invariant: %v is not SMSNotification", t))
	}
}

// GetSMSNotification returns SMSNotification and true boolean if Notification is SMSNotification.
func (s Notification) GetSMSNotification() (v SMSNotification, ok bool) {
	if !s.IsSMSNotification() {
		return v, false
	}
	return s.SMSNotification, true
}

// NewNotificationSMSNotification returns new Notification from SMSNotification.
func NewNotificationSMSNotification(v SMSNotification) Notification {
	var s Notification
	s.SetSMSNotification(NotificationSMSNotification, v)
	return s
}

// NewNotificationTextNotification returns new Notification from SMSNotification.
func NewNotificationTextNotification(v SMSNotification) Notification {
	var s Notification
	s.SetSMSNotification(NotificationTextNotification, v)
	return s
}

// SetPushNotification sets Notification to PushNotification.
// panics if `t` is not associated with PushNotification
func (s *Notification) SetPushNotification(t NotificationType, v PushNotification) {
	s.Type = t
	s.PushNotification = v
	if !s.IsPushNotification() {
		panic(fmt.Errorf("invariant: %v is not PushNotification", t))
	}
}

// GetPushNotification returns PushNotification and true boolean if Notification is PushNotification.
func (s Notification) GetPushNotification() (v PushNotification, ok bool) {
	if !s.IsPushNotification() {
		return v, false
	}
	return s.PushNotification, true
}

// NewNotificationMobileNotification returns new Notification from PushNotification.
func NewNotificationMobileNotification(v PushNotification) Notification {
	var s Notification
	s.SetPushNotification(NotificationMobileNotification, v)
	return s
}

// NewNotificationPushNotification returns new Notification from PushNotification.
func NewNotificationPushNotification(v PushNotification) Notification {
	var s Notification
	s.SetPushNotification(NotificationPushNotification, v)
	return s
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
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

// Ref: #/components/schemas/Pet
// Pet represents sum type.
type Pet struct {
	Type PetType // switch on this field
	Dog  Dog
	Cat  Cat
	Bird Bird
}

// PetType is oneOf type of Pet.
type PetType string

// Possible values for PetType.
const (
	PetCaninePet PetType = "canine"
	PetDogPet    PetType = "dog"
	PetCatPet    PetType = "cat"
	PetFelinePet PetType = "feline"
	PetAvianPet  PetType = "avian"
	PetBirdPet   PetType = "bird"
)

// IsDog reports whether Pet is Dog.
func (s Pet) IsDog() bool {
	switch s.Type {
	case PetCaninePet, PetDogPet:
		return true
	default:
		return false
	}
}

// IsCat reports whether Pet is Cat.
func (s Pet) IsCat() bool {
	switch s.Type {
	case PetCatPet, PetFelinePet:
		return true
	default:
		return false
	}
}

// IsBird reports whether Pet is Bird.
func (s Pet) IsBird() bool {
	switch s.Type {
	case PetAvianPet, PetBirdPet:
		return true
	default:
		return false
	}
}

// SetDog sets Pet to Dog.
// panics if `t` is not associated with Dog
func (s *Pet) SetDog(t PetType, v Dog) {
	s.Type = t
	s.Dog = v
	if !s.IsDog() {
		panic(fmt.Errorf("invariant: %v is not Dog", t))
	}
}

// GetDog returns Dog and true boolean if Pet is Dog.
func (s Pet) GetDog() (v Dog, ok bool) {
	if !s.IsDog() {
		return v, false
	}
	return s.Dog, true
}

// NewPetCaninePet returns new Pet from Dog.
func NewPetCaninePet(v Dog) Pet {
	var s Pet
	s.SetDog(PetCaninePet, v)
	return s
}

// NewPetDogPet returns new Pet from Dog.
func NewPetDogPet(v Dog) Pet {
	var s Pet
	s.SetDog(PetDogPet, v)
	return s
}

// SetCat sets Pet to Cat.
// panics if `t` is not associated with Cat
func (s *Pet) SetCat(t PetType, v Cat) {
	s.Type = t
	s.Cat = v
	if !s.IsCat() {
		panic(fmt.Errorf("invariant: %v is not Cat", t))
	}
}

// GetCat returns Cat and true boolean if Pet is Cat.
func (s Pet) GetCat() (v Cat, ok bool) {
	if !s.IsCat() {
		return v, false
	}
	return s.Cat, true
}

// NewPetCatPet returns new Pet from Cat.
func NewPetCatPet(v Cat) Pet {
	var s Pet
	s.SetCat(PetCatPet, v)
	return s
}

// NewPetFelinePet returns new Pet from Cat.
func NewPetFelinePet(v Cat) Pet {
	var s Pet
	s.SetCat(PetFelinePet, v)
	return s
}

// SetBird sets Pet to Bird.
// panics if `t` is not associated with Bird
func (s *Pet) SetBird(t PetType, v Bird) {
	s.Type = t
	s.Bird = v
	if !s.IsBird() {
		panic(fmt.Errorf("invariant: %v is not Bird", t))
	}
}

// GetBird returns Bird and true boolean if Pet is Bird.
func (s Pet) GetBird() (v Bird, ok bool) {
	if !s.IsBird() {
		return v, false
	}
	return s.Bird, true
}

// NewPetAvianPet returns new Pet from Bird.
func NewPetAvianPet(v Bird) Pet {
	var s Pet
	s.SetBird(PetAvianPet, v)
	return s
}

// NewPetBirdPet returns new Pet from Bird.
func NewPetBirdPet(v Bird) Pet {
	var s Pet
	s.SetBird(PetBirdPet, v)
	return s
}

// Ref: #/components/schemas/PushNotification
type PushNotification struct {
	NotificationType string    `json:"notificationType"`
	DeviceId         string    `json:"deviceId"`
	Title            string    `json:"title"`
	Body             OptString `json:"body"`
	Badge            OptInt    `json:"badge"`
}

// GetNotificationType returns the value of NotificationType.
func (s *PushNotification) GetNotificationType() string {
	return s.NotificationType
}

// GetDeviceId returns the value of DeviceId.
func (s *PushNotification) GetDeviceId() string {
	return s.DeviceId
}

// GetTitle returns the value of Title.
func (s *PushNotification) GetTitle() string {
	return s.Title
}

// GetBody returns the value of Body.
func (s *PushNotification) GetBody() OptString {
	return s.Body
}

// GetBadge returns the value of Badge.
func (s *PushNotification) GetBadge() OptInt {
	return s.Badge
}

// SetNotificationType sets the value of NotificationType.
func (s *PushNotification) SetNotificationType(val string) {
	s.NotificationType = val
}

// SetDeviceId sets the value of DeviceId.
func (s *PushNotification) SetDeviceId(val string) {
	s.DeviceId = val
}

// SetTitle sets the value of Title.
func (s *PushNotification) SetTitle(val string) {
	s.Title = val
}

// SetBody sets the value of Body.
func (s *PushNotification) SetBody(val OptString) {
	s.Body = val
}

// SetBadge sets the value of Badge.
func (s *PushNotification) SetBadge(val OptInt) {
	s.Badge = val
}

// Ref: #/components/schemas/SMSNotification
type SMSNotification struct {
	NotificationType string `json:"notificationType"`
	PhoneNumber      string `json:"phoneNumber"`
	Message          string `json:"message"`
}

// GetNotificationType returns the value of NotificationType.
func (s *SMSNotification) GetNotificationType() string {
	return s.NotificationType
}

// GetPhoneNumber returns the value of PhoneNumber.
func (s *SMSNotification) GetPhoneNumber() string {
	return s.PhoneNumber
}

// GetMessage returns the value of Message.
func (s *SMSNotification) GetMessage() string {
	return s.Message
}

// SetNotificationType sets the value of NotificationType.
func (s *SMSNotification) SetNotificationType(val string) {
	s.NotificationType = val
}

// SetPhoneNumber sets the value of PhoneNumber.
func (s *SMSNotification) SetPhoneNumber(val string) {
	s.PhoneNumber = val
}

// SetMessage sets the value of Message.
func (s *SMSNotification) SetMessage(val string) {
	s.Message = val
}

// Ref: #/components/schemas/Vehicle
// Vehicle represents sum type.
type Vehicle struct {
	Type       VehicleType // switch on this field
	Car        Car
	Motorcycle Motorcycle
}

// VehicleType is oneOf type of Vehicle.
type VehicleType string

// Possible values for VehicleType.
const (
	CarVehicle        VehicleType = "Car"
	MotorcycleVehicle VehicleType = "Motorcycle"
)

// IsCar reports whether Vehicle is Car.
func (s Vehicle) IsCar() bool { return s.Type == CarVehicle }

// IsMotorcycle reports whether Vehicle is Motorcycle.
func (s Vehicle) IsMotorcycle() bool { return s.Type == MotorcycleVehicle }

// SetCar sets Vehicle to Car.
func (s *Vehicle) SetCar(v Car) {
	s.Type = CarVehicle
	s.Car = v
}

// GetCar returns Car and true boolean if Vehicle is Car.
func (s Vehicle) GetCar() (v Car, ok bool) {
	if !s.IsCar() {
		return v, false
	}
	return s.Car, true
}

// NewCarVehicle returns new Vehicle from Car.
func NewCarVehicle(v Car) Vehicle {
	var s Vehicle
	s.SetCar(v)
	return s
}

// SetMotorcycle sets Vehicle to Motorcycle.
func (s *Vehicle) SetMotorcycle(v Motorcycle) {
	s.Type = MotorcycleVehicle
	s.Motorcycle = v
}

// GetMotorcycle returns Motorcycle and true boolean if Vehicle is Motorcycle.
func (s Vehicle) GetMotorcycle() (v Motorcycle, ok bool) {
	if !s.IsMotorcycle() {
		return v, false
	}
	return s.Motorcycle, true
}

// NewMotorcycleVehicle returns new Vehicle from Motorcycle.
func NewMotorcycleVehicle(v Motorcycle) Vehicle {
	var s Vehicle
	s.SetMotorcycle(v)
	return s
}
