//GVariantType : GVariantType — introduction to the GVariant type system
package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

/*
 * GVariantType
 */

// IVariantType is an interface type implemented by VariantType and all types which embed
// an VariantType.  It is meant to be used as a type for function arguments which
// require GVariantTypes or any subclasses thereof.
type IVariantType interface {
	toGVariantType() *C.GVariantType
	toVariantType() *VariantType
}

// VariantType is a representation of GLib's GVariantType.
type VariantType struct {
	GVariantType *C.GVariantType
}

func (v *VariantType) toGVariantType() *C.GVariantType {
	if v == nil {
		return nil
	}
	return v.native()
}

func (v *VariantType) toVariantType() *VariantType {
	return v
}

// newVariantType creates a new VariantType from a GVariantType pointer.
func newVariantType(p *C.GVariantType) *VariantType {
	return &VariantType{GVariantType: p}
}

func VariantTypeFromUnsafePointer(p unsafe.Pointer) *VariantType{
	return &VariantType{ C.toGVariantType(p) }
}

// native returns a pointer to the underlying GVariantType.
func (v *VariantType) native() *C.GVariantType {
	if v == nil || v.GVariantType == nil {
		return nil
	}
	p := unsafe.Pointer(v.GVariantType)
	return C.toGVariantType(p)
}

// Native returns a pointer to the underlying GVariantType.
func (v *VariantType) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}


//#define	G_VARIANT_TYPE_BOOLEAN
//#define	G_VARIANT_TYPE_BYTE
//#define	G_VARIANT_TYPE_INT16
//#define	G_VARIANT_TYPE_UINT16
//#define	G_VARIANT_TYPE_INT32
//#define	G_VARIANT_TYPE_UINT32
//#define	G_VARIANT_TYPE_INT64
//#define	G_VARIANT_TYPE_UINT64
//#define	G_VARIANT_TYPE_HANDLE
//#define	G_VARIANT_TYPE_DOUBLE
//#define	G_VARIANT_TYPE_STRING
//#define	G_VARIANT_TYPE_OBJECT_PATH
//#define	G_VARIANT_TYPE_SIGNATURE
//#define	G_VARIANT_TYPE_VARIANT
//#define	G_VARIANT_TYPE_ANY
//#define	G_VARIANT_TYPE_BASIC
//#define	G_VARIANT_TYPE_MAYBE
//#define	G_VARIANT_TYPE_ARRAY
//#define	G_VARIANT_TYPE_TUPLE
//#define	G_VARIANT_TYPE_UNIT
//#define	G_VARIANT_TYPE_DICT_ENTRY
//#define	G_VARIANT_TYPE_DICTIONARY
//#define	G_VARIANT_TYPE_STRING_ARRAY
//#define	G_VARIANT_TYPE_OBJECT_PATH_ARRAY
//#define	G_VARIANT_TYPE_BYTESTRING
//#define	G_VARIANT_TYPE_BYTESTRING_ARRAY
//#define	G_VARIANT_TYPE_VARDICT
//#define	G_VARIANT_TYPE()
//void	g_variant_type_free ()
//GVariantType *	g_variant_type_copy ()
//GVariantType *	g_variant_type_new ()
//gboolean	g_variant_type_string_is_valid ()
//gboolean	g_variant_type_string_scan ()
//gsize	g_variant_type_get_string_length ()
//const gchar *	g_variant_type_peek_string ()
//gchar *	g_variant_type_dup_string ()
//gboolean	g_variant_type_is_definite ()
//gboolean	g_variant_type_is_container ()
//gboolean	g_variant_type_is_basic ()
//gboolean	g_variant_type_is_maybe ()
//gboolean	g_variant_type_is_array ()
//gboolean	g_variant_type_is_tuple ()
//gboolean	g_variant_type_is_dict_entry ()
//gboolean	g_variant_type_is_variant ()
//guint	g_variant_type_hash ()
//gboolean	g_variant_type_equal ()
//gboolean	g_variant_type_is_subtype_of ()
//GVariantType *	g_variant_type_new_maybe ()
//GVariantType *	g_variant_type_new_array ()
//GVariantType *	g_variant_type_new_tuple ()
//GVariantType *	g_variant_type_new_dict_entry ()
//const GVariantType *	g_variant_type_element ()
//gsize	g_variant_type_n_items ()
//const GVariantType *	g_variant_type_first ()
//const GVariantType *	g_variant_type_next ()
//const GVariantType *	g_variant_type_key ()
//const GVariantType *	g_variant_type_value ()
