// TreeStore
package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/terrak/gotk3/glib"
)

/*
 * GtkTreeStore
 */

// TreeStore is a representation of GTK's GtkTreeStore.
type TreeStore struct {
	*glib.Object

	// Interfaces
	TreeModel
}

// native returns a pointer to the underlying GtkTreeStore.
func (v *TreeStore) native() *C.GtkTreeStore {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTreeStore(p)
}

func marshalTreeStore(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapTreeStore(obj), nil
}

func wrapTreeStore(obj *glib.Object) *TreeStore {
	tm := wrapTreeModel(obj)
	return &TreeStore{obj, *tm}
}

func (v *TreeStore) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return C.toGtkTreeModel(unsafe.Pointer(v.GObject))
}

// TreeStoreNew is a wrapper around gtk_tree_store_newv().
func TreeStoreNew(types ...glib.Type) (*TreeStore, error) {
	gtypes := C.alloc_types(C.int(len(types)))
	for n, val := range types {
		C.set_type(gtypes, C.int(n), C.GType(val))
	}
	defer C.g_free(C.gpointer(gtypes))
	c := C.gtk_tree_store_newv(C.gint(len(types)), gtypes)
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	ls := wrapTreeStore(obj)
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return ls, nil
}

// TODO(terrak)
/*
func (v *TreeStore) SetColumnTypes(types ...glib.Type) {
}
*/

// SetValue() is a wrapper around gtk_tree_store_set_value()
func (v *TreeStore) SetValue(iter *TreeIter, column int, value interface{}) error {
	if gv, err := glib.GValue(value); err != nil {
		return err
	} else {
		C.gtk_tree_store_set_value(v.native(), iter.native(), C.gint(column), (*C.GValue)(unsafe.Pointer(gv.Native())))
	}
	return nil
}

// Set() is a wrapper around gtk_tree_store_set_value() but provides
// a function similar to gtk_tree_store_set() in that multiple columns
// may be set by one call.  The length of columns and values slices must
// match, or Set() will return a non-nil error.
//
// As an example, a call to:
//  store.Set(iter, []int{0, 1}, []interface{}{"Foo", "Bar"})
// is functionally equivalent to calling the native C GTK function:
//  gtk_tree_store_set(store, iter, 0, "Foo", 1, "Bar", -1);
func (v *TreeStore) Set(iter *TreeIter, columns []int, values []interface{}) error {
	if len(columns) != len(values) {
		return errors.New("columns and values lengths do not match")
	}
	for i, val := range values {
		if gv, err := glib.GValue(val); err != nil {
			return err
		} else {
			C.gtk_tree_store_set_value(v.native(), iter.native(),
				C.gint(columns[i]),
				(*C.GValue)(unsafe.Pointer(gv.Native())))
		}
	}
	return nil
}

// Not implemented :
// gtk_tree_store_set_valist ()
// gtk_tree_store_set_valuesv ()

// Remove() is a wrapper around gtk_tree_store_remove().
func (v *TreeStore) Remove(iter *TreeIter) bool {
	c := C.gtk_tree_store_remove(v.native(), iter.native())
	return gobool(c)
}

// Insert() is a wrapper around gtk_tree_store_insert().
//Creates a new row at position . If parent is non-NULL, then the row will be made a child of parent . Otherwise, the row will be created at the toplevel. If position is -1 or is larger than the number of rows //at that level, then the new row will be inserted to the end of the list. iter will be changed to point to this new row. The row will be empty after this function is called. To fill in values, you need to call gtk_tree_store_set() or gtk_tree_store_set_value().
func (v *TreeStore) Insert(parent *TreeIter, position int) *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_tree_store_insert(v.native(), &ti, parent.native(), C.gint(position))
	iter := &TreeIter{ti}
	return iter
}

// InsertBefore() is a wrapper around gtk_tree_store_insert_before().
//Inserts a new row before sibling . If sibling is NULL, then the row will be appended to parent ’s children. If parent and sibling are NULL, then the row will be appended to the toplevel. If both sibling and parent are set, then parent must be the parent of sibling . When sibling is set, parent is optional.
//iter will be changed to point to this new row. The row will be empty after this function is called. To fill in values, you need to call gtk_tree_store_set() or gtk_tree_store_set_value().
func (v *TreeStore) InsertBefore(parent, sibling *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_tree_store_insert_before(v.native(), &ti, parent.native(), sibling.native())
	iter := &TreeIter{ti}
	return iter
}

// InsertAfter() is a wrapper around gtk_tree_store_insert_after().
//Inserts a new row after sibling . If sibling is NULL, then the row will be prepended to parent ’s children. If parent and sibling are NULL, then the row will be prepended to the toplevel. If both sibling and parent are set, then parent must be the parent of sibling . When sibling is set, parent is optional.
//iter will be changed to point to this new row. The row will be empty after this function is called. To fill in values, you need to call gtk_tree_store_set() or gtk_tree_store_set_value().
func (v *TreeStore) InsertAfter(parent, sibling *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_tree_store_insert_after(v.native(), &ti, parent.native(), sibling.native())
	iter := &TreeIter{ti}
	return iter
}

// Not implemented :
// gtk_tree_store_insert_with_values ()
// gtk_tree_store_insert_with_valuesv ()

// Prepend() is a wrapper around gtk_tree_store_prepend().
//Prepends a new row to tree_store . If parent is non-NULL, then it will prepend the new row before the first child of parent , otherwise it will prepend a row to the top level. iter will be changed to point to this new row. The row will be empty after this function is called. To fill in values, you need to call gtk_tree_store_set() or gtk_tree_store_set_value().
func (v *TreeStore) Prepend(parent *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_tree_store_prepend(v.native(), &ti, parent.native())
	iter := &TreeIter{ti}
	return iter
}

// Append() is a wrapper around gtk_tree_store_append().
//Appends a new row to tree_store . If parent is non-NULL, then it will append the new row after the last child of parent , otherwise it will append a row to the top level. iter will be changed to point to this new row. The row will be empty after this function is called. To fill in values, you need to call gtk_tree_store_set() or gtk_tree_store_set_value().
func (v *TreeStore) Append(parent *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_tree_store_append(v.native(), &ti, parent.native())
	iter := &TreeIter{ti}
	return iter
}

// IsAncestor() is a wrapper around gtk_tree_store_is_ancestor().
//Returns TRUE if iter is an ancestor of descendant . That is, iter is the parent (or grandparent or great-grandparent) of descendant .
func (v *TreeStore) IsAncestor(iter, descendant *TreeIter) bool {
	c := C.gtk_tree_store_is_ancestor(v.native(), iter.native(), descendant.native())
	return gobool(c)
}

// IterDepth() is a wrapper around gtk_tree_store_iter_depth().
//Returns the depth of iter . This will be 0 for anything on the root level, 1 for anything down a level, etc.
func (v *TreeStore) IterDepth(iter *TreeIter) int {
	c := C.gtk_tree_store_iter_depth(v.native(), iter.native())
	return int(c)
}

// Clear() is a wrapper around gtk_tree_store_clear().
//Removes all rows from tree_store
func (v *TreeStore) Clear() {
	C.gtk_tree_store_clear(v.native())
}

// IterIsValid() is a wrapper around gtk_tree_store_iter_is_valid().
//WARNING: This function is slow. Only use it for debugging and/or testing purposes.
//Checks if the given iter is a valid iter for this GtkTreeStore.
func (v *TreeStore) IterIsValid(iter *TreeIter) bool {
	c := C.gtk_tree_store_iter_is_valid(v.native(), iter.native())
	return gobool(c)
}

// TODO(terrak)
/*
func (v *TreeStore) Reorder(newOrder []int) {
}
*/

// Swap() is a wrapper around gtk_tree_store_swap().
//Swaps a and b in the same level of tree_store . Note that this function only works with unsorted stores.
func (v *TreeStore) Swap(a, b *TreeIter) {
	C.gtk_tree_store_swap(v.native(), a.native(), b.native())
}

// MoveBefore() is a wrapper around gtk_tree_store_move_before().
//Moves iter in tree_store to the position before position . iter and position should be in the same level. Note that this function only works with unsorted stores. If position is NULL, iter will be moved to the end of the level.
func (v *TreeStore) MoveBefore(iter, position *TreeIter) {
	C.gtk_tree_store_move_before(v.native(), iter.native(), position.native())
}

// MoveAfter() is a wrapper around gtk_tree_store_move_after().
func (v *TreeStore) MoveAfter(iter, position *TreeIter) {
	C.gtk_tree_store_move_after(v.native(), iter.native(), position.native())
}
