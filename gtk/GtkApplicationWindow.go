// ApplicationWindow
package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/terrak/gotk3/glib"
)

/*
 * GtkApplicationWindow
 */

// ApplicationWindow is a representation of GTK's GtkApplicationWindow.
type ApplicationWindow struct {
	Window
}

// native returns a pointer to the underlying GtkApplicationWindow.
func (v *ApplicationWindow) native() *C.GtkApplicationWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkApplicationWindow(p)
}

func (v *ApplicationWindow) toApplicationWindow() *C.GtkApplicationWindow {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalApplicationWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapApplicationWindow(obj), nil
}

func wrapApplicationWindow(obj *glib.Object) *ApplicationWindow {
	return &ApplicationWindow{Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}
}

//GtkWidget *
//gtk_application_window_new (GtkApplication *application);

//Creates a new GtkApplicationWindow.
func ApplicationWindowNew(app *Application) (*ApplicationWindow, error) {
	c := C.gtk_application_window_new(app.native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	appw := wrapApplicationWindow(obj)
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return appw, nil
}

//void
//gtk_application_window_set_show_menubar
//                               (GtkApplicationWindow *window,
//                                gboolean show_menubar);

//Sets whether the window will display a menubar for the app menu and menubar as needed.
func (v *ApplicationWindow) SetShowMenubar(show_menubar bool) {
	C.gtk_application_window_set_show_menubar(v.native(), gbool(show_menubar))
}

//gboolean
//gtk_application_window_get_show_menubar
//                               (GtkApplicationWindow *window);

//Returns whether the window will display a menubar for the app menu and menubar as needed.
func (v *ApplicationWindow) GetShowMenubar() bool {
	c := C.gtk_application_window_get_show_menubar(v.native())
	return gobool(c)
}

//guint
//gtk_application_window_get_id (GtkApplicationWindow *window);

//Returns the unique ID of the window. If the window has not yet been added to a GtkApplication, returns 0.
func (v *ApplicationWindow) GetId() uint {
	c := C.gtk_application_window_get_id(v.native())
	return uint(c)
}
