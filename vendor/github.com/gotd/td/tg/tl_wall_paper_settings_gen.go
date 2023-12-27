// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// WallPaperSettings represents TL type `wallPaperSettings#372efcd0`.
// Wallpaper¹ rendering information.
//
// Links:
//  1. https://core.telegram.org/api/wallpapers
//
// See https://core.telegram.org/constructor/wallPaperSettings for reference.
type WallPaperSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// For image wallpapers »¹: if set, the JPEG must be downscaled to fit in 450x450
	// square and then box-blurred with radius 12.
	//
	// Links:
	//  1) https://core.telegram.org/api/wallpapers#image-wallpapers
	Blur bool
	// If set, the background needs to be slightly moved when the device is rotated.
	Motion bool
	// Used for solid »¹, gradient »² and freeform gradient »³ fills.
	//
	// Links:
	//  1) https://core.telegram.org/api/wallpapers#solid-fill
	//  2) https://core.telegram.org/api/wallpapers#gradient-fill
	//  3) https://core.telegram.org/api/wallpapers#freeform-gradient-fill
	//
	// Use SetBackgroundColor and GetBackgroundColor helpers.
	BackgroundColor int
	// Used for gradient »¹ and freeform gradient »² fills.
	//
	// Links:
	//  1) https://core.telegram.org/api/wallpapers#gradient-fill
	//  2) https://core.telegram.org/api/wallpapers#freeform-gradient-fill
	//
	// Use SetSecondBackgroundColor and GetSecondBackgroundColor helpers.
	SecondBackgroundColor int
	// Used for freeform gradient »¹ fills.
	//
	// Links:
	//  1) https://core.telegram.org/api/wallpapers#freeform-gradient-fill
	//
	// Use SetThirdBackgroundColor and GetThirdBackgroundColor helpers.
	ThirdBackgroundColor int
	// Used for freeform gradient »¹ fills.
	//
	// Links:
	//  1) https://core.telegram.org/api/wallpapers#freeform-gradient-fill
	//
	// Use SetFourthBackgroundColor and GetFourthBackgroundColor helpers.
	FourthBackgroundColor int
	// Used for pattern wallpapers »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/wallpapers#pattern-wallpapers
	//
	// Use SetIntensity and GetIntensity helpers.
	Intensity int
	// Clockwise rotation angle of the gradient, in degrees; 0-359. Should be always
	// divisible by 45.
	//
	// Use SetRotation and GetRotation helpers.
	Rotation int
	// Emoticon field of WallPaperSettings.
	//
	// Use SetEmoticon and GetEmoticon helpers.
	Emoticon string
}

// WallPaperSettingsTypeID is TL type id of WallPaperSettings.
const WallPaperSettingsTypeID = 0x372efcd0

// Ensuring interfaces in compile-time for WallPaperSettings.
var (
	_ bin.Encoder     = &WallPaperSettings{}
	_ bin.Decoder     = &WallPaperSettings{}
	_ bin.BareEncoder = &WallPaperSettings{}
	_ bin.BareDecoder = &WallPaperSettings{}
)

func (w *WallPaperSettings) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.Blur == false) {
		return false
	}
	if !(w.Motion == false) {
		return false
	}
	if !(w.BackgroundColor == 0) {
		return false
	}
	if !(w.SecondBackgroundColor == 0) {
		return false
	}
	if !(w.ThirdBackgroundColor == 0) {
		return false
	}
	if !(w.FourthBackgroundColor == 0) {
		return false
	}
	if !(w.Intensity == 0) {
		return false
	}
	if !(w.Rotation == 0) {
		return false
	}
	if !(w.Emoticon == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WallPaperSettings) String() string {
	if w == nil {
		return "WallPaperSettings(nil)"
	}
	type Alias WallPaperSettings
	return fmt.Sprintf("WallPaperSettings%+v", Alias(*w))
}

// FillFrom fills WallPaperSettings from given interface.
func (w *WallPaperSettings) FillFrom(from interface {
	GetBlur() (value bool)
	GetMotion() (value bool)
	GetBackgroundColor() (value int, ok bool)
	GetSecondBackgroundColor() (value int, ok bool)
	GetThirdBackgroundColor() (value int, ok bool)
	GetFourthBackgroundColor() (value int, ok bool)
	GetIntensity() (value int, ok bool)
	GetRotation() (value int, ok bool)
	GetEmoticon() (value string, ok bool)
}) {
	w.Blur = from.GetBlur()
	w.Motion = from.GetMotion()
	if val, ok := from.GetBackgroundColor(); ok {
		w.BackgroundColor = val
	}

	if val, ok := from.GetSecondBackgroundColor(); ok {
		w.SecondBackgroundColor = val
	}

	if val, ok := from.GetThirdBackgroundColor(); ok {
		w.ThirdBackgroundColor = val
	}

	if val, ok := from.GetFourthBackgroundColor(); ok {
		w.FourthBackgroundColor = val
	}

	if val, ok := from.GetIntensity(); ok {
		w.Intensity = val
	}

	if val, ok := from.GetRotation(); ok {
		w.Rotation = val
	}

	if val, ok := from.GetEmoticon(); ok {
		w.Emoticon = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WallPaperSettings) TypeID() uint32 {
	return WallPaperSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*WallPaperSettings) TypeName() string {
	return "wallPaperSettings"
}

// TypeInfo returns info about TL type.
func (w *WallPaperSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "wallPaperSettings",
		ID:   WallPaperSettingsTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Blur",
			SchemaName: "blur",
			Null:       !w.Flags.Has(1),
		},
		{
			Name:       "Motion",
			SchemaName: "motion",
			Null:       !w.Flags.Has(2),
		},
		{
			Name:       "BackgroundColor",
			SchemaName: "background_color",
			Null:       !w.Flags.Has(0),
		},
		{
			Name:       "SecondBackgroundColor",
			SchemaName: "second_background_color",
			Null:       !w.Flags.Has(4),
		},
		{
			Name:       "ThirdBackgroundColor",
			SchemaName: "third_background_color",
			Null:       !w.Flags.Has(5),
		},
		{
			Name:       "FourthBackgroundColor",
			SchemaName: "fourth_background_color",
			Null:       !w.Flags.Has(6),
		},
		{
			Name:       "Intensity",
			SchemaName: "intensity",
			Null:       !w.Flags.Has(3),
		},
		{
			Name:       "Rotation",
			SchemaName: "rotation",
			Null:       !w.Flags.Has(4),
		},
		{
			Name:       "Emoticon",
			SchemaName: "emoticon",
			Null:       !w.Flags.Has(7),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (w *WallPaperSettings) SetFlags() {
	if !(w.Blur == false) {
		w.Flags.Set(1)
	}
	if !(w.Motion == false) {
		w.Flags.Set(2)
	}
	if !(w.BackgroundColor == 0) {
		w.Flags.Set(0)
	}
	if !(w.SecondBackgroundColor == 0) {
		w.Flags.Set(4)
	}
	if !(w.ThirdBackgroundColor == 0) {
		w.Flags.Set(5)
	}
	if !(w.FourthBackgroundColor == 0) {
		w.Flags.Set(6)
	}
	if !(w.Intensity == 0) {
		w.Flags.Set(3)
	}
	if !(w.Rotation == 0) {
		w.Flags.Set(4)
	}
	if !(w.Emoticon == "") {
		w.Flags.Set(7)
	}
}

// Encode implements bin.Encoder.
func (w *WallPaperSettings) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode wallPaperSettings#372efcd0 as nil")
	}
	b.PutID(WallPaperSettingsTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WallPaperSettings) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode wallPaperSettings#372efcd0 as nil")
	}
	w.SetFlags()
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode wallPaperSettings#372efcd0: field flags: %w", err)
	}
	if w.Flags.Has(0) {
		b.PutInt(w.BackgroundColor)
	}
	if w.Flags.Has(4) {
		b.PutInt(w.SecondBackgroundColor)
	}
	if w.Flags.Has(5) {
		b.PutInt(w.ThirdBackgroundColor)
	}
	if w.Flags.Has(6) {
		b.PutInt(w.FourthBackgroundColor)
	}
	if w.Flags.Has(3) {
		b.PutInt(w.Intensity)
	}
	if w.Flags.Has(4) {
		b.PutInt(w.Rotation)
	}
	if w.Flags.Has(7) {
		b.PutString(w.Emoticon)
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WallPaperSettings) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode wallPaperSettings#372efcd0 to nil")
	}
	if err := b.ConsumeID(WallPaperSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode wallPaperSettings#372efcd0: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WallPaperSettings) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode wallPaperSettings#372efcd0 to nil")
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaperSettings#372efcd0: field flags: %w", err)
		}
	}
	w.Blur = w.Flags.Has(1)
	w.Motion = w.Flags.Has(2)
	if w.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaperSettings#372efcd0: field background_color: %w", err)
		}
		w.BackgroundColor = value
	}
	if w.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaperSettings#372efcd0: field second_background_color: %w", err)
		}
		w.SecondBackgroundColor = value
	}
	if w.Flags.Has(5) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaperSettings#372efcd0: field third_background_color: %w", err)
		}
		w.ThirdBackgroundColor = value
	}
	if w.Flags.Has(6) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaperSettings#372efcd0: field fourth_background_color: %w", err)
		}
		w.FourthBackgroundColor = value
	}
	if w.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaperSettings#372efcd0: field intensity: %w", err)
		}
		w.Intensity = value
	}
	if w.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaperSettings#372efcd0: field rotation: %w", err)
		}
		w.Rotation = value
	}
	if w.Flags.Has(7) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaperSettings#372efcd0: field emoticon: %w", err)
		}
		w.Emoticon = value
	}
	return nil
}

// SetBlur sets value of Blur conditional field.
func (w *WallPaperSettings) SetBlur(value bool) {
	if value {
		w.Flags.Set(1)
		w.Blur = true
	} else {
		w.Flags.Unset(1)
		w.Blur = false
	}
}

// GetBlur returns value of Blur conditional field.
func (w *WallPaperSettings) GetBlur() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(1)
}

// SetMotion sets value of Motion conditional field.
func (w *WallPaperSettings) SetMotion(value bool) {
	if value {
		w.Flags.Set(2)
		w.Motion = true
	} else {
		w.Flags.Unset(2)
		w.Motion = false
	}
}

// GetMotion returns value of Motion conditional field.
func (w *WallPaperSettings) GetMotion() (value bool) {
	if w == nil {
		return
	}
	return w.Flags.Has(2)
}

// SetBackgroundColor sets value of BackgroundColor conditional field.
func (w *WallPaperSettings) SetBackgroundColor(value int) {
	w.Flags.Set(0)
	w.BackgroundColor = value
}

// GetBackgroundColor returns value of BackgroundColor conditional field and
// boolean which is true if field was set.
func (w *WallPaperSettings) GetBackgroundColor() (value int, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(0) {
		return value, false
	}
	return w.BackgroundColor, true
}

// SetSecondBackgroundColor sets value of SecondBackgroundColor conditional field.
func (w *WallPaperSettings) SetSecondBackgroundColor(value int) {
	w.Flags.Set(4)
	w.SecondBackgroundColor = value
}

// GetSecondBackgroundColor returns value of SecondBackgroundColor conditional field and
// boolean which is true if field was set.
func (w *WallPaperSettings) GetSecondBackgroundColor() (value int, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(4) {
		return value, false
	}
	return w.SecondBackgroundColor, true
}

// SetThirdBackgroundColor sets value of ThirdBackgroundColor conditional field.
func (w *WallPaperSettings) SetThirdBackgroundColor(value int) {
	w.Flags.Set(5)
	w.ThirdBackgroundColor = value
}

// GetThirdBackgroundColor returns value of ThirdBackgroundColor conditional field and
// boolean which is true if field was set.
func (w *WallPaperSettings) GetThirdBackgroundColor() (value int, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(5) {
		return value, false
	}
	return w.ThirdBackgroundColor, true
}

// SetFourthBackgroundColor sets value of FourthBackgroundColor conditional field.
func (w *WallPaperSettings) SetFourthBackgroundColor(value int) {
	w.Flags.Set(6)
	w.FourthBackgroundColor = value
}

// GetFourthBackgroundColor returns value of FourthBackgroundColor conditional field and
// boolean which is true if field was set.
func (w *WallPaperSettings) GetFourthBackgroundColor() (value int, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(6) {
		return value, false
	}
	return w.FourthBackgroundColor, true
}

// SetIntensity sets value of Intensity conditional field.
func (w *WallPaperSettings) SetIntensity(value int) {
	w.Flags.Set(3)
	w.Intensity = value
}

// GetIntensity returns value of Intensity conditional field and
// boolean which is true if field was set.
func (w *WallPaperSettings) GetIntensity() (value int, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(3) {
		return value, false
	}
	return w.Intensity, true
}

// SetRotation sets value of Rotation conditional field.
func (w *WallPaperSettings) SetRotation(value int) {
	w.Flags.Set(4)
	w.Rotation = value
}

// GetRotation returns value of Rotation conditional field and
// boolean which is true if field was set.
func (w *WallPaperSettings) GetRotation() (value int, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(4) {
		return value, false
	}
	return w.Rotation, true
}

// SetEmoticon sets value of Emoticon conditional field.
func (w *WallPaperSettings) SetEmoticon(value string) {
	w.Flags.Set(7)
	w.Emoticon = value
}

// GetEmoticon returns value of Emoticon conditional field and
// boolean which is true if field was set.
func (w *WallPaperSettings) GetEmoticon() (value string, ok bool) {
	if w == nil {
		return
	}
	if !w.Flags.Has(7) {
		return value, false
	}
	return w.Emoticon, true
}
