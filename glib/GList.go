//GList : Doubly-Linked Lists — linked lists that can be iterated over in both directions
package glib
// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

// List is a representation of Glib's GList.
type List struct {
	Data uintptr
	Next *List
	Prev *List
}

// Append is a wrapper around g_list_append.
func (v *List) Append(data uintptr) *List {
	glist := (*C.GList)(unsafe.Pointer(v))
	glist = C.g_list_append(glist, C.gpointer(data))
	return (*List)(unsafe.Pointer(glist))
}

// Prepend is a wrapper around g_list_prepend.
func (v *List) Prepend(data uintptr) *List {
	glist := (*C.GList)(unsafe.Pointer(v))
	glist = C.g_list_prepend(glist, C.gpointer(data))
	return (*List)(unsafe.Pointer(glist))
}

// Insert is a wrapper around g_list_insert().
func (v *List) Insert(data uintptr, position int) *List {
	glist := (*C.GList)(unsafe.Pointer(v))
	glist = C.g_list_insert(glist, C.gpointer(data), C.gint(position))
	return (*List)(unsafe.Pointer(glist))
}

//GList *	g_list_insert_before ()
//GList *	g_list_insert_sorted ()
//GList *	g_list_remove ()
//GList *	g_list_remove_link ()
//GList *	g_list_delete_link ()
//GList *	g_list_remove_all ()
//void	g_list_free ()
//void	g_list_free_full ()
//GList *	g_list_alloc ()
//void	g_list_free_1 ()
//guint	g_list_length ()
//GList *	g_list_copy ()
//GList *	g_list_copy_deep ()
//GList *	g_list_reverse ()
//GList *	g_list_sort ()
//gint	(*GCompareFunc) ()
//GList *	g_list_insert_sorted_with_data ()
//GList *	g_list_sort_with_data ()
//gint	(*GCompareDataFunc) ()
//GList *	g_list_concat ()
//void	g_list_foreach ()
//void	(*GFunc) ()
//GList *	g_list_first ()
//GList *	g_list_last ()
//#define	g_list_previous()
//#define	g_list_next()
//GList *	g_list_nth ()
//gpointer	g_list_nth_data ()
//GList *	g_list_nth_prev ()
//GList *	g_list_find ()
//GList *	g_list_find_custom ()
//gint	g_list_position ()
//gint	g_list_index ()
