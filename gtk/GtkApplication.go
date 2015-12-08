// Application
package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/terrak/gotk3/gio"
	"github.com/terrak/gotk3/glib"
)

/*
 * GtkApplication
 */

// Application is a representation of GTK's GtkApplication.
type Application struct {
	*gio.Application
}

// native returns a pointer to the underlying GtkTreeStore.
func (v *Application) native() *C.GtkApplication {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkApplication(p)
}

func (v *Application) toApplication() *C.GtkApplication {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalApplication(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapApplication(obj), nil
}

func wrapApplication(obj *glib.Object) *Application {
	return &Application{&gio.Application{Object:obj}}
}

//IGlibConvert : special conversion for gtk.Application to gio.Application
func (v *Application) Convert(t reflect.Type) reflect.Value {
	if v == nil || strings.Compare(t.String(), "*gio.Application") != 0 {
		panic("[gtk.Application.Convert ] type not manage "+ t.String())
	}
	return reflect.ValueOf(v.Application)
}


/*
func 	 (v* Application) toGObject() *glib.C.GObject{
	return v.app.GObject
}
func (v* Application) toObject() *glib.Object{
	return v.app.Object
}
*/

// ApplicationNew is a wrapper around gtk_application_new().
/*
Creates a new GtkApplication instance.
When using GtkApplication, it is not necessary to call gtk_init() manually. It is called as soon as the application gets registered as the primary instance.
Concretely, gtk_init() is called in the default handler for the “startup” signal. Therefore, GtkApplication subclasses should chain up in their “startup” handler before using any GTK+ API.
Note that commandline arguments are not passed to gtk_init(). All GTK+ functionality that is available via commandline arguments can also be achieved by setting suitable environment variables such as G_DEBUG, so this should not be a big problem. If you absolutely must support GTK+ commandline arguments, you can explicitly call gtk_init() before creating the application instance.
If non-NULL, the application ID must be valid. See g_application_id_is_valid().
If no application ID is given then some features (most notably application uniqueness) will be disabled. A null application ID is only allowed with GTK+ 3.6 or later.
*/
func ApplicationNew(application_id string, flags gio.ApplicationFlags) (*Application, error) {
	cstr := C.CString(application_id)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_application_new((*C.gchar)(cstr), C.GApplicationFlags(flags))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	app := wrapApplication(obj)
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return app, nil
}

// AddWindow is a wrapper around gtk_application_add_window().
/*
Adds a window to application .
This call can only happen after the application has started; typically, you should add new application windows in response to the emission of the “activate” signal.
This call is equivalent to setting the “application” property of window to application .
Normally, the connection between the application and the window will remain until the window is destroyed, but you can explicitly remove it with gtk_application_remove_window().
GTK+ will keep the application running as long as it has any windows.
*/
func (v *Application) AddWindow(window IWindow) {
	var w *C.GtkWindow = nil
	if window != nil {
		w = window.toWindow()
	}
	C.gtk_application_add_window(v.native(), w)
}

// RemoveWindow is a wrapper around gtk_application_remove_window().
/*
Remove a window from application .
If window belongs to application then this call is equivalent to setting the “application” property of window to NULL.
The application may stop running as a result of a call to this function.
*/
func (v *Application) RemoveWindow(window IWindow) {
	var w *C.GtkWindow = nil
	if window != nil {
		w = window.toWindow()
	}
	C.gtk_application_remove_window(v.native(), w)
}

// GetWindows is a wrapper around gtk_application_get_windows().
/*
Gets a list of the GtkWindows associated with application .
The list is sorted by most recently focused window, such that the first element is the currently focused window. (Useful for choosing a parent for a transient window.)
The list that is returned should not be modified in any way. It will only remain valid until the next focus change or window creation or deletion.
*/

func (v *Application) GetWindows() []*Window { //FIXME
	/*
		c := C.gtk_application_get_windows(v.native())
		obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
		w := wrapWindow(obj)
		return w
	*/
	return nil
}

// GetWindowById is a wrapper around gtk_application_get_window_by_id().
/*
Returns the GtkApplicationWindow with the given ID.
*/

func (v *Application) GetWindowById(id uint) *Window {
	c := C.gtk_application_get_window_by_id(v.native(), C.guint(id))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWindow(obj)
	return w
}

// GetActiveWindow is a wrapper around gtk_application_get_active_window().
/*
Gets the “active” window for the application.
The active window is the one that was most recently focused (within the application). This window may not have the focus at the moment if another application has it — this is just the most recently-focused window within this application.
*/

func (v *Application) GetActiveWindow() *Window {
	c := C.gtk_application_get_active_window(v.native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWindow(obj)
	return w
}

// Inhibit is a wrapper around gtk_application_inhibit().
/*
Inform the session manager that certain types of actions should be inhibited. This is not guaranteed to work on all platforms and for all types of actions.
Applications should invoke this method when they begin an operation that should not be interrupted, such as creating a CD or DVD. The types of actions that may be blocked are specified by the flags parameter. When the application completes the operation it should call gtk_application_uninhibit() to remove the inhibitor. Note that an application can have multiple inhibitors, and all of the must be individually removed. Inhibitors are also cleared when the application exits.
Applications should not expect that they will always be able to block the action. In most cases, users will be given the option to force the action to take place.
Reasons should be short and to the point.
If window is given, the session manager may point the user to this window to find out more about why the action is inhibited.
*/

func (v *Application) Inhibit(window IWindow, flags ApplicationInhibitFlags, reason string) uint {
	cstr := C.CString(reason)
	defer C.free(unsafe.Pointer(cstr))
	var w *C.GtkWindow = nil
	if window != nil {
		w = window.toWindow()
	}
	c := C.gtk_application_inhibit(v.native(), w, C.GtkApplicationInhibitFlags(flags), (*C.gchar)(cstr))

	return uint(c)
}

// Uninhibit is a wrapper around gtk_application_uninhibit().
/*
Removes an inhibitor that has been established with gtk_application_inhibit(). Inhibitors are also cleared when the application exits.
*/
func (v *Application) Uninhibit(cookie uint) {
	C.gtk_application_uninhibit(v.native(), C.guint(cookie))
}

// IsInhibited is a wrapper around gtk_application_is_inhibited().
/*
Determines if any of the actions specified in flags are currently inhibited (possibly by another application).
*/
func (v *Application) IsInhibited(flags ApplicationInhibitFlags) bool {
	c := C.gtk_application_is_inhibited(v.native(), C.GtkApplicationInhibitFlags(flags))

	return gobool(c)
}

// PrefersAppMenu is a wrapper around gtk_application_prefers_app_menu().
/*
Determines if the desktop environment in which the application is running would prefer an application menu be shown.
If this function returns TRUE then the application should call gtk_application_set_app_menu() with the contents of an application menu, which will be shown by the desktop environment. If it returns FALSE then you should consider using an alternate approach, such as a menubar.
The value returned by this function is purely advisory and you are free to ignore it. If you call gtk_application_set_app_menu() even if the desktop environment doesn't support app menus, then a fallback will be provided.
Applications are similarly free not to set an app menu even if the desktop environment wants to show one. In that case, a fallback will also be created by the desktop environment (GNOME, for example, uses a menu with only a "Quit" item in it).
The value returned by this function never changes. Once it returns a particular value, it is guaranteed to always return the same value.
You may only call this function after the application has been registered and after the base startup handler has run. You're most likely to want to use this from your own startup handler. It may also make sense to consult this function while constructing UI (in activate, open or an action activation handler) in order to determine if you should show a gear menu or not.
This function will return FALSE on Mac OS and a default app menu will be created automatically with the "usual" contents of that menu typical to most Mac OS applications. If you call gtk_application_set_app_menu() anyway, then this menu will be replaced with your own.
*/

func (v *Application) PrefersAppMenu() bool {
	c := C.gtk_application_prefers_app_menu(v.native())

	return gobool(c)
}

// GetAppMenu is a wrapper around gtk_application_get_app_menu().
/*
Returns the menu model that has been set with gtk_application_set_app_menu().
*/
func (v *Application) GetAppMenu() *Menu {
	c := C.gtk_application_get_app_menu(v.native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMenu(obj)
	return m
}

// SetAppMenu is a wrapper around gtk_application_set_app_menu().
/*
Sets or unsets the application menu for application .
This can only be done in the primary instance of the application, after it has been registered. “startup” is a good place to call this.
The application menu is a single menu containing items that typically impact the application as a whole, rather than acting on a specific window or document. For example, you would expect to see “Preferences” or “Quit” in an application menu, but not “Save” or “Print”.
If supported, the application menu will be rendered by the desktop environment.
Use the base GActionMap interface to add actions, to respond to the user selecting these menu items.
*/
func (v *Application) SetAppMenu(appmenu *MenuModel) {
	C.gtk_application_set_app_menu(v.native(), appmenu.native())
}

// GetMenuBar is a wrapper around gtk_application_get_menubar().
/*
Returns the menu model that has been set with gtk_application_set_menubar().
*/
func (v *Application) GetMenuBar() *MenuModel {
	c := C.gtk_application_get_menubar(v.native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMenuModel(obj)
	return m
}

// SetMenuBar is a wrapper around gtk_application_set_menubar().
/*
Sets or unsets the menubar for windows of application .
This is a menubar in the traditional sense.
This can only be done in the primary instance of the application, after it has been registered. “startup” is a good place to call this.
Depending on the desktop environment, this may appear at the top of each window, or at the top of the screen. In some environments, if both the application menu and the menubar are set, the application menu will be presented as if it were the first item of the menubar. Other environments treat the two as completely separate — for example, the application menu may be rendered by the desktop shell while the menubar (if set) remains in each individual window.
Use the base GActionMap interface to add actions, to respond to the user selecting these menu items.
*/
func (v *Application) SetMenuBar(appmenu *MenuModel) {
	C.gtk_application_set_menubar(v.native(), appmenu.native())
}

// GetMenuById is a wrapper around gtk_application_get_menu_by_id().
/*
Returns the menu model that has been set with gtk_application_set_menubar().
*/
func (v *Application) GetMenuById(id string) *Menu {
	cstr := C.CString(id)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_application_get_menubar(v.native())
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	m := wrapMenu(obj)
	return m
}

// ListActionDescriptions is a wrapper around gtk_application_list_action_descriptions().
/*
Lists the detailed action names which have associated accelerators. See gtk_application_set_accels_for_action().
*/
func (v *Application) ListActionDescriptions() []string {
	c := C.gtk_application_list_action_descriptions(v.native())
	defer C.g_strfreev(c)
	len := int(C.g_strv_length(c))
	res := make([]string, len)
	for i := 0; i < len; i++ {
		cstr := C.get_string_array_string(c, (C.uint)(i)) //FIXME
		res[i] = C.GoString((*C.char)(cstr))
	}

	return res
}

// GetAccelsForAction is a wrapper around gtk_application_get_accels_for_action().
/*
Gets the accelerators that are currently associated with the given action.
*/
func (v *Application) GetAccelsForAction(detailed_action_name string) []string {
	cdetailed_action_name := C.CString(detailed_action_name)
	defer C.free(unsafe.Pointer(cdetailed_action_name))
	c := C.gtk_application_get_accels_for_action(v.native(), (*C.gchar)(cdetailed_action_name))
	defer C.g_strfreev(c)
	var len, i uint
	len = uint(C.g_strv_length(c))
	res := make([]string, len)
	for i = 0; i < len; i++ {
		cstr := C.get_string_array_string(c, (C.uint)(i)) //FIXME
		res[i] = C.GoString((*C.char)(cstr))
	}

	return res
}

// SetAccelsForAction is a wrapper around gtk_application_set_accels_for_action().
/*
Sets zero or more keyboard accelerators that will trigger the given action. The first item in accels will be the primary accelerator, which may be displayed in the UI.
To remove all accelerators for an action, use an empty, zero-terminated array for accels .
*/
func (v *Application) SetAccelsForAction(detailed_action_name string, actions []string) {
	cdetailed_action_name := C.CString(detailed_action_name)
	defer C.free(unsafe.Pointer(cdetailed_action_name))
	/*
		//FIXME créer cactions
		defer C.g_strfreev(cactions)
		C.gtk_application_set_accels_for_action(v.native(), cdetailed_action_name, cactions)
	*/
}

// GetActionsForAccel is a wrapper around gtk_application_get_actions_for_accel().
/*
Returns the list of actions (possibly empty) that accel maps to. Each item in the list is a detailed action name in the usual form.
This might be useful to discover if an accel already exists in order to prevent installation of a conflicting accelerator (from an accelerator editor or a plugin system, for example). Note that having more than one action per accelerator may not be a bad thing and might make sense in cases where the actions never appear in the same context.
In case there are no actions for a given accelerator, an empty array is returned. NULL is never returned.
It is a programmer error to pass an invalid accelerator string. If you are unsure, check it with gtk_accelerator_parse() first.
*/
func (v *Application) GetActionsForAccel(accel string) []string {
	caccel := C.CString(accel)
	defer C.free(unsafe.Pointer(caccel))
	c := C.gtk_application_get_actions_for_accel(v.native(), (*C.gchar)(caccel))
	defer C.g_strfreev(c)
	var len, i uint
	len = uint(C.g_strv_length(c))
	res := make([]string, len)
	for i = 0; i < len; i++ {
		cstr := C.get_string_array_string(c, (C.uint)(i)) //FIXME
		res[i] = C.GoString((*C.char)(cstr))
	}

	return res
}
