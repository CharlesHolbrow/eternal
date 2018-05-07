package eternal

import "github.com/CharlesHolbrow/synk"

// cellDiff diff type for synk.Object
type cellDiff struct {
  Hue *float32 `json:"hue,omitempty"`
  X *int `json:"x,omitempty"`
  Y *int `json:"y,omitempty"`
  AudioPath *string `json:"audioPath,omitempty"`
  MapName *string `json:"mapName,omitempty"`
  Class *string `json:"class,omitempty"`
  MidiNote *int `json:"midiNote,omitempty"`
}

// State returns a fully populated diff of the unresolved state
func (o *Cell) State() interface{} {
	d := cellDiff{
    Hue: &o.Hue,
    X: &o.X,
    Y: &o.Y,
    AudioPath: &o.AudioPath,
    MapName: &o.MapName,
    Class: &o.Class,
    MidiNote: &o.MidiNote,
  }
  return d
}

// Resolve applies the current diff, then returns it
func (o *Cell) Resolve() interface{} {
  if o.diff.Hue != nil {o.Hue = *o.diff.Hue}
  if o.diff.X != nil {o.X = *o.diff.X}
  if o.diff.Y != nil {o.Y = *o.diff.Y}
  if o.diff.AudioPath != nil {o.AudioPath = *o.diff.AudioPath}
  if o.diff.MapName != nil {o.MapName = *o.diff.MapName}
  if o.diff.Class != nil {o.Class = *o.diff.Class}
  if o.diff.MidiNote != nil {o.MidiNote = *o.diff.MidiNote}
  o.V++
  diff := o.diff
  o.diff = cellDiff{}
  return diff
}

// Changed checks if struct has been changed since the last .Resolve()
func (o *Cell) Changed() bool {
  return o.diff.Hue != nil ||
		o.diff.X != nil ||
		o.diff.Y != nil ||
		o.diff.AudioPath != nil ||
		o.diff.MapName != nil ||
		o.diff.Class != nil ||
		o.diff.MidiNote != nil
}

// Diff getter
func (o *Cell) Diff() interface{} { return o.diff }
// Copy duplicates this object and returns an interface to it.
// The object's diff will be copied too, with the exception of the diffMap for
// array members. A diffMap is created automatically when we use array Element
// setters (ex SetDataElement). Copy() will create shallow copies of unresolved
// diffMaps. Usually we Resolve() after Copy() which means that our shallow copy
// will be safe to send over a channel.
func (o *Cell) Copy() synk.Object {
	n := *o
	return &n
}
// Init (ialize) all diff fields to the current values. The next call to
// Resolve() will return a diff with all the fields initialized.
func (o *Cell) Init() {
	o.diff = o.State().(cellDiff)
}
// SetHue on diff
func (o *Cell) SetHue(v float32) {
  if v != o.Hue {
    o.diff.Hue = &v
  } else {
    o.diff.Hue = nil
  }
}
// GetPrevHue Gets the previous value. Ignores diff.
func (o *Cell) GetPrevHue() float32 { return o.Hue }
// GetHue from diff. Fall back to current value if no diff
func (o *Cell) GetHue() float32 {
	if o.diff.Hue != nil {
		return *o.diff.Hue
	}
	return o.Hue
}
// GetHue. Diff method
func (o cellDiff) GetHue() *float32 { return o.Hue }
// SetX on diff
func (o *Cell) SetX(v int) {
  if v != o.X {
    o.diff.X = &v
  } else {
    o.diff.X = nil
  }
}
// GetPrevX Gets the previous value. Ignores diff.
func (o *Cell) GetPrevX() int { return o.X }
// GetX from diff. Fall back to current value if no diff
func (o *Cell) GetX() int {
	if o.diff.X != nil {
		return *o.diff.X
	}
	return o.X
}
// GetX. Diff method
func (o cellDiff) GetX() *int { return o.X }
// SetY on diff
func (o *Cell) SetY(v int) {
  if v != o.Y {
    o.diff.Y = &v
  } else {
    o.diff.Y = nil
  }
}
// GetPrevY Gets the previous value. Ignores diff.
func (o *Cell) GetPrevY() int { return o.Y }
// GetY from diff. Fall back to current value if no diff
func (o *Cell) GetY() int {
	if o.diff.Y != nil {
		return *o.diff.Y
	}
	return o.Y
}
// GetY. Diff method
func (o cellDiff) GetY() *int { return o.Y }
// SetAudioPath on diff
func (o *Cell) SetAudioPath(v string) {
  if v != o.AudioPath {
    o.diff.AudioPath = &v
  } else {
    o.diff.AudioPath = nil
  }
}
// GetPrevAudioPath Gets the previous value. Ignores diff.
func (o *Cell) GetPrevAudioPath() string { return o.AudioPath }
// GetAudioPath from diff. Fall back to current value if no diff
func (o *Cell) GetAudioPath() string {
	if o.diff.AudioPath != nil {
		return *o.diff.AudioPath
	}
	return o.AudioPath
}
// GetAudioPath. Diff method
func (o cellDiff) GetAudioPath() *string { return o.AudioPath }
// SetMapName on diff
func (o *Cell) SetMapName(v string) {
  if v != o.MapName {
    o.diff.MapName = &v
  } else {
    o.diff.MapName = nil
  }
}
// GetPrevMapName Gets the previous value. Ignores diff.
func (o *Cell) GetPrevMapName() string { return o.MapName }
// GetMapName from diff. Fall back to current value if no diff
func (o *Cell) GetMapName() string {
	if o.diff.MapName != nil {
		return *o.diff.MapName
	}
	return o.MapName
}
// GetMapName. Diff method
func (o cellDiff) GetMapName() *string { return o.MapName }
// SetClass on diff
func (o *Cell) SetClass(v string) {
  if v != o.Class {
    o.diff.Class = &v
  } else {
    o.diff.Class = nil
  }
}
// GetPrevClass Gets the previous value. Ignores diff.
func (o *Cell) GetPrevClass() string { return o.Class }
// GetClass from diff. Fall back to current value if no diff
func (o *Cell) GetClass() string {
	if o.diff.Class != nil {
		return *o.diff.Class
	}
	return o.Class
}
// GetClass. Diff method
func (o cellDiff) GetClass() *string { return o.Class }
// SetMidiNote on diff
func (o *Cell) SetMidiNote(v int) {
  if v != o.MidiNote {
    o.diff.MidiNote = &v
  } else {
    o.diff.MidiNote = nil
  }
}
// GetPrevMidiNote Gets the previous value. Ignores diff.
func (o *Cell) GetPrevMidiNote() int { return o.MidiNote }
// GetMidiNote from diff. Fall back to current value if no diff
func (o *Cell) GetMidiNote() int {
	if o.diff.MidiNote != nil {
		return *o.diff.MidiNote
	}
	return o.MidiNote
}
// GetMidiNote. Diff method
func (o cellDiff) GetMidiNote() *int { return o.MidiNote }
// noteDiff diff type for synk.Object
type noteDiff struct {
  SubKey *string `json:"subKey,omitempty"`
  Number *int `json:"number,omitempty"`
  Velocity *int `json:"velocity,omitempty"`
}

// State returns a fully populated diff of the unresolved state
func (o *Note) State() interface{} {
	d := noteDiff{
    SubKey: &o.SubKey,
    Number: &o.Number,
    Velocity: &o.Velocity,
  }
  return d
}

// Resolve applies the current diff, then returns it
func (o *Note) Resolve() interface{} {
  if o.diff.SubKey != nil {o.SubKey = *o.diff.SubKey}
  if o.diff.Number != nil {o.Number = *o.diff.Number}
  if o.diff.Velocity != nil {o.Velocity = *o.diff.Velocity}
  o.V++
  diff := o.diff
  o.diff = noteDiff{}
  return diff
}

// Changed checks if struct has been changed since the last .Resolve()
func (o *Note) Changed() bool {
  return o.diff.SubKey != nil ||
		o.diff.Number != nil ||
		o.diff.Velocity != nil
}

// Diff getter
func (o *Note) Diff() interface{} { return o.diff }
// Copy duplicates this object and returns an interface to it.
// The object's diff will be copied too, with the exception of the diffMap for
// array members. A diffMap is created automatically when we use array Element
// setters (ex SetDataElement). Copy() will create shallow copies of unresolved
// diffMaps. Usually we Resolve() after Copy() which means that our shallow copy
// will be safe to send over a channel.
func (o *Note) Copy() synk.Object {
	n := *o
	return &n
}
// Init (ialize) all diff fields to the current values. The next call to
// Resolve() will return a diff with all the fields initialized.
func (o *Note) Init() {
	o.diff = o.State().(noteDiff)
}
// SetSubKey on diff
func (o *Note) SetSubKey(v string) {
  if v != o.SubKey {
    o.diff.SubKey = &v
  } else {
    o.diff.SubKey = nil
  }
}
// GetPrevSubKey Gets the previous value. Ignores diff.
func (o *Note) GetPrevSubKey() string { return o.SubKey }
// GetSubKey from diff. Fall back to current value if no diff
func (o *Note) GetSubKey() string {
	if o.diff.SubKey != nil {
		return *o.diff.SubKey
	}
	return o.SubKey
}
// GetSubKey. Diff method
func (o noteDiff) GetSubKey() *string { return o.SubKey }
// SetNumber on diff
func (o *Note) SetNumber(v int) {
  if v != o.Number {
    o.diff.Number = &v
  } else {
    o.diff.Number = nil
  }
}
// GetPrevNumber Gets the previous value. Ignores diff.
func (o *Note) GetPrevNumber() int { return o.Number }
// GetNumber from diff. Fall back to current value if no diff
func (o *Note) GetNumber() int {
	if o.diff.Number != nil {
		return *o.diff.Number
	}
	return o.Number
}
// GetNumber. Diff method
func (o noteDiff) GetNumber() *int { return o.Number }
// SetVelocity on diff
func (o *Note) SetVelocity(v int) {
  if v != o.Velocity {
    o.diff.Velocity = &v
  } else {
    o.diff.Velocity = nil
  }
}
// GetPrevVelocity Gets the previous value. Ignores diff.
func (o *Note) GetPrevVelocity() int { return o.Velocity }
// GetVelocity from diff. Fall back to current value if no diff
func (o *Note) GetVelocity() int {
	if o.diff.Velocity != nil {
		return *o.diff.Velocity
	}
	return o.Velocity
}
// GetVelocity. Diff method
func (o noteDiff) GetVelocity() *int { return o.Velocity }
