package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "GtkMisc.go.h"
import "C"
import (
	"unsafe"

	"github.com/terrak/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums

		// Objects/Interfaces

		{glib.Type(C.gtk_misc_get_type()), marshalMisc},

		// Boxed

	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * GtkMisc
 */

// Misc is a representation of GTK's GtkMisc.
type Misc struct {
	Widget
}

// native returns a pointer to the underlying GtkMisc.
func (v *Misc) native() *C.GtkMisc {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMisc(p)
}

func marshalMisc(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMisc(obj), nil
}

func wrapMisc(obj *glib.Object) *Misc {
	return &Misc{Widget{glib.InitiallyUnowned{obj}}}
}

/*DEPRECATED
// GetAlignment is a wrapper around gtk_misc_get_alignment().
func (v *Misc) GetAlignment() (xAlign, yAlign float32) {
	var x, y C.gfloat
	C.gtk_misc_get_alignment(v.native(), &x, &y)
	return float32(x), float32(y)
}

// SetAlignment is a wrapper around gtk_misc_set_alignment().
func (v *Misc) SetAlignment(xAlign, yAlign float32) {
	C.gtk_misc_set_alignment(v.native(), C.gfloat(xAlign), C.gfloat(yAlign))
}

// GetPadding is a wrapper around gtk_misc_get_padding().
func (v *Misc) GetPadding() (xpad, ypad int) {
	var x, y C.gint
	C.gtk_misc_get_padding(v.native(), &x, &y)
	return int(x), int(y)
}

// SetPadding is a wrapper around gtk_misc_set_padding().
func (v *Misc) SetPadding(xPad, yPad int) {
	C.gtk_misc_set_padding(v.native(), C.gint(xPad), C.gint(yPad))
}
*/
