package eternal

import "github.com/CharlesHolbrow/synk"

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
