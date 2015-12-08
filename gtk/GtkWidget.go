package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/terrak/gotk3/gdk"
	"github.com/terrak/gotk3/glib"
)

/*
 * GtkWidget
 */

// Widget is a representation of GTK's GtkWidget.
type Widget struct {
	glib.InitiallyUnowned
}

// IWidget is an interface type implemented by all structs
// embedding a Widget.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkWidget.
type IWidget interface {
	toWidget() *C.GtkWidget
}

// native returns a pointer to the underlying GtkWidget.
func (v *Widget) native() *C.GtkWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkWidget(p)
}

func (v *Widget) toWidget() *C.GtkWidget {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapWidget(obj), nil
}

func wrapWidget(obj *glib.Object) *Widget {
	return &Widget{glib.InitiallyUnowned{obj}}
}

// Destroy is a wrapper around gtk_widget_destroy().
func (v *Widget) Destroy() {
	C.gtk_widget_destroy(v.native())
}

// InDestruction is a wrapper around gtk_widget_in_destruction().
func (v *Widget) InDestruction() bool {
	return gobool(C.gtk_widget_in_destruction(v.native()))
}

// TODO(jrick) this may require some rethinking
/*
func (v *Widget) Destroyed(widgetPointer **Widget) {
}
*/

// Unparent is a wrapper around gtk_widget_unparent().
func (v *Widget) Unparent() {
	C.gtk_widget_unparent(v.native())
}

// Show is a wrapper around gtk_widget_show().
func (v *Widget) Show() {
	C.gtk_widget_show(v.native())
}

// Hide is a wrapper around gtk_widget_hide().
func (v *Widget) Hide() {
	C.gtk_widget_hide(v.native())
}

// GetCanFocus is a wrapper around gtk_widget_get_can_focus().
func (v *Widget) GetCanFocus() bool {
	c := C.gtk_widget_get_can_focus(v.native())
	return gobool(c)
}

// SetCanFocus is a wrapper around gtk_widget_set_can_focus().
func (v *Widget) SetCanFocus(canFocus bool) {
	C.gtk_widget_set_can_focus(v.native(), gbool(canFocus))
}

// GetMapped is a wrapper around gtk_window_get_mapped().
func (v *Widget) GetMapped() bool {
	c := C.gtk_widget_get_mapped(v.native())
	return gobool(c)
}

// SetMapped is a wrapper around gtk_widget_set_mapped().
func (v *Widget) SetMapped(mapped bool) {
	C.gtk_widget_set_can_focus(v.native(), gbool(mapped))
}

// GetRealized is a wrapper around gtk_window_get_realized().
func (v *Widget) GetRealized() bool {
	c := C.gtk_widget_get_realized(v.native())
	return gobool(c)
}

// SetRealized is a wrapper around gtk_widget_set_realized().
func (v *Widget) SetRealized(realized bool) {
	C.gtk_widget_set_realized(v.native(), gbool(realized))
}

/* DEPRECATED
// SetDoubleBuffered is a wrapper around gtk_widget_set_double_buffered().
func (v *Widget) SetDoubleBuffered(doubleBuffered bool) {
	C.gtk_widget_set_double_buffered(v.native(), gbool(doubleBuffered))
}
*/

/* DEPRECATED
// GetDoubleBuffered is a wrapper around gtk_widget_get_double_buffered().
func (v *Widget) GetDoubleBuffered() bool {
	c := C.gtk_widget_get_double_buffered(v.native())
	return gobool(c)
}
*/

// GetHasWindow is a wrapper around gtk_widget_get_has_window().
func (v *Widget) GetHasWindow() bool {
	c := C.gtk_widget_get_has_window(v.native())
	return gobool(c)
}

// SetHasWindow is a wrapper around gtk_widget_set_has_window().
func (v *Widget) SetHasWindow(hasWindow bool) {
	C.gtk_widget_set_has_window(v.native(), gbool(hasWindow))
}

// ShowNow is a wrapper around gtk_widget_show_now().
func (v *Widget) ShowNow() {
	C.gtk_widget_show_now(v.native())
}

// ShowAll is a wrapper around gtk_widget_show_all().
func (v *Widget) ShowAll() {
	C.gtk_widget_show_all(v.native())
}

// SetNoShowAll is a wrapper around gtk_widget_set_no_show_all().
func (v *Widget) SetNoShowAll(noShowAll bool) {
	C.gtk_widget_set_no_show_all(v.native(), gbool(noShowAll))
}

// GetNoShowAll is a wrapper around gtk_widget_get_no_show_all().
func (v *Widget) GetNoShowAll() bool {
	c := C.gtk_widget_get_no_show_all(v.native())
	return gobool(c)
}

// Map is a wrapper around gtk_widget_map().
func (v *Widget) Map() {
	C.gtk_widget_map(v.native())
}

// Unmap is a wrapper around gtk_widget_unmap().
func (v *Widget) Unmap() {
	C.gtk_widget_unmap(v.native())
}

//void gtk_widget_realize(GtkWidget *widget);
//void gtk_widget_unrealize(GtkWidget *widget);
//void gtk_widget_draw(GtkWidget *widget, cairo_t *cr);
//void gtk_widget_queue_resize(GtkWidget *widget);
//void gtk_widget_queue_resize_no_redraw(GtkWidget *widget);
//GdkFrameClock *gtk_widget_get_frame_clock(GtkWidget *widget);
//guint gtk_widget_add_tick_callback (GtkWidget *widget,
//                                    GtkTickCallback callback,
//                                    gpointer user_data,
//                                    GDestroyNotify notify);
//void gtk_widget_remove_tick_callback(GtkWidget *widget, guint id);

// TODO(jrick) GtkAllocation
/*
func (v *Widget) SizeAllocate() {
}
*/

// TODO(jrick) GtkAccelGroup GdkModifierType GtkAccelFlags
/*
func (v *Widget) AddAccelerator() {
}
*/

// TODO(jrick) GtkAccelGroup GdkModifierType
/*
func (v *Widget) RemoveAccelerator() {
}
*/

// TODO(jrick) GtkAccelGroup
/*
func (v *Widget) SetAccelPath() {
}
*/

// TODO(jrick) GList
/*
func (v *Widget) ListAccelClosures() {
}
*/

// GetAllocatedWidth() is a wrapper around gtk_widget_get_allocated_width().
func (v *Widget) GetAllocatedWidth() int {
	return int(C.gtk_widget_get_allocated_width(v.native()))
}

// GetAllocatedHeight() is a wrapper around gtk_widget_get_allocated_height().
func (v *Widget) GetAllocatedHeight() int {
	return int(C.gtk_widget_get_allocated_height(v.native()))
}

//gboolean gtk_widget_can_activate_accel(GtkWidget *widget, guint signal_id);
func (v *Widget) CanActivateAccel(signal_id uint) bool {
	c := C.gtk_widget_can_activate_accel(v.native(), C.guint(signal_id))
	return gobool(c)
}

// Event() is a wrapper around gtk_widget_event().
func (v *Widget) Event(event *gdk.Event) bool {
	c := C.gtk_widget_event(v.native(),
		(*C.GdkEvent)(unsafe.Pointer(event.Native())))
	return gobool(c)
}

// Activate() is a wrapper around gtk_widget_activate().
func (v *Widget) Activate() bool {
	return gobool(C.gtk_widget_activate(v.native()))
}

/* DEPRECATED
// Reparent() is a wrapper around gtk_widget_reparent().
func (v *Widget) Reparent(newParent IWidget) {
	C.gtk_widget_reparent(v.native(), newParent.toWidget())
}
*/

// TODO(jrick) GdkRectangle
/*
func (v *Widget) Intersect() {
}
*/

// IsFocus() is a wrapper around gtk_widget_is_focus().
func (v *Widget) IsFocus() bool {
	return gobool(C.gtk_widget_is_focus(v.native()))
}

// GrabFocus() is a wrapper around gtk_widget_grab_focus().
func (v *Widget) GrabFocus() {
	C.gtk_widget_grab_focus(v.native())
}

// GrabDefault() is a wrapper around gtk_widget_grab_default().
func (v *Widget) GrabDefault() {
	C.gtk_widget_grab_default(v.native())
}

// SetName() is a wrapper around gtk_widget_set_name().
func (v *Widget) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_name(v.native(), (*C.gchar)(cstr))
}

// GetName() is a wrapper around gtk_widget_get_name().  A non-nil
// error is returned in the case that gtk_widget_get_name returns NULL to
// differentiate between NULL and an empty string.
func (v *Widget) GetName() (string, error) {
	c := C.gtk_widget_get_name(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetSensitive is a wrapper around gtk_widget_get_sensitive().
func (v *Widget) GetSensitive() bool {
	c := C.gtk_widget_get_sensitive(v.native())
	return gobool(c)
}

// IsSensitive is a wrapper around gtk_widget_is_sensitive().
func (v *Widget) IsSensitive() bool {
	c := C.gtk_widget_is_sensitive(v.native())
	return gobool(c)
}

// SetSensitive is a wrapper around gtk_widget_set_sensitive().
func (v *Widget) SetSensitive(sensitive bool) {
	C.gtk_widget_set_sensitive(v.native(), gbool(sensitive))
}

// GetVisible is a wrapper around gtk_widget_get_visible().
func (v *Widget) GetVisible() bool {
	c := C.gtk_widget_get_visible(v.native())
	return gobool(c)
}

// SetVisible is a wrapper around gtk_widget_set_visible().
func (v *Widget) SetVisible(visible bool) {
	C.gtk_widget_set_visible(v.native(), gbool(visible))
}

// SetParent is a wrapper around gtk_widget_set_parent().
func (v *Widget) SetParent(parent IWidget) {
	C.gtk_widget_set_parent(v.native(), parent.toWidget())
}

// GetParent is a wrapper around gtk_widget_get_parent().
func (v *Widget) GetParent() (*Widget, error) {
	c := C.gtk_widget_get_parent(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// SetSizeRequest is a wrapper around gtk_widget_set_size_request().
func (v *Widget) SetSizeRequest(width, height int) {
	C.gtk_widget_set_size_request(v.native(), C.gint(width), C.gint(height))
}

// GetSizeRequest is a wrapper around gtk_widget_get_size_request().
func (v *Widget) GetSizeRequest() (width, height int) {
	var w, h C.gint
	C.gtk_widget_get_size_request(v.native(), &w, &h)
	return int(w), int(h)
}

// SetParentWindow is a wrapper around gtk_widget_set_parent_window().
func (v *Widget) SetParentWindow(parentWindow *gdk.Window) {
	C.gtk_widget_set_parent_window(v.native(),
		(*C.GdkWindow)(unsafe.Pointer(parentWindow.Native())))
}

// GetParentWindow is a wrapper around gtk_widget_get_parent_window().
func (v *Widget) GetParentWindow() (*gdk.Window, error) {
	c := C.gtk_widget_get_parent_window(v.native())
	if v == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := &gdk.Window{obj}
	w.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// SetEvents is a wrapper around gtk_widget_set_events().
func (v *Widget) SetEvents(events int) {
	C.gtk_widget_set_events(v.native(), C.gint(events))
}

// GetEvents is a wrapper around gtk_widget_get_events().
func (v *Widget) GetEvents() int {
	return int(C.gtk_widget_get_events(v.native()))
}

// AddEvents is a wrapper around gtk_widget_add_events().
func (v *Widget) AddEvents(events int) {
	C.gtk_widget_add_events(v.native(), C.gint(events))
}

// HasDefault is a wrapper around gtk_widget_has_default().
func (v *Widget) HasDefault() bool {
	c := C.gtk_widget_has_default(v.native())
	return gobool(c)
}

// HasFocus is a wrapper around gtk_widget_has_focus().
func (v *Widget) HasFocus() bool {
	c := C.gtk_widget_has_focus(v.native())
	return gobool(c)
}

// HasVisibleFocus is a wrapper around gtk_widget_has_visible_focus().
func (v *Widget) HasVisibleFocus() bool {
	c := C.gtk_widget_has_visible_focus(v.native())
	return gobool(c)
}

// HasGrab is a wrapper around gtk_widget_has_grab().
func (v *Widget) HasGrab() bool {
	c := C.gtk_widget_has_grab(v.native())
	return gobool(c)
}

// IsDrawable is a wrapper around gtk_widget_is_drawable().
func (v *Widget) IsDrawable() bool {
	c := C.gtk_widget_is_drawable(v.native())
	return gobool(c)
}

// IsToplevel is a wrapper around gtk_widget_is_toplevel().
func (v *Widget) IsToplevel() bool {
	c := C.gtk_widget_is_toplevel(v.native())
	return gobool(c)
}

// TODO(jrick) GdkEventMask
/*
func (v *Widget) SetDeviceEvents() {
}
*/

// TODO(jrick) GdkEventMask
/*
func (v *Widget) GetDeviceEvents() {
}
*/

// TODO(jrick) GdkEventMask
/*
func (v *Widget) AddDeviceEvents() {
}
*/

// SetDeviceEnabled is a wrapper around gtk_widget_set_device_enabled().
func (v *Widget) SetDeviceEnabled(device *gdk.Device, enabled bool) {
	C.gtk_widget_set_device_enabled(v.native(),
		(*C.GdkDevice)(unsafe.Pointer(device.Native())), gbool(enabled))
}

// GetDeviceEnabled is a wrapper around gtk_widget_get_device_enabled().
func (v *Widget) GetDeviceEnabled(device *gdk.Device) bool {
	c := C.gtk_widget_get_device_enabled(v.native(),
		(*C.GdkDevice)(unsafe.Pointer(device.Native())))
	return gobool(c)
}

// GetToplevel is a wrapper around gtk_widget_get_toplevel().
func (v *Widget) GetToplevel() (*Widget, error) {
	c := C.gtk_widget_get_toplevel(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// GetTooltipText is a wrapper around gtk_widget_get_tooltip_text().
// A non-nil error is returned in the case that
// gtk_widget_get_tooltip_text returns NULL to differentiate between NULL
// and an empty string.
func (v *Widget) GetTooltipText() (string, error) {
	c := C.gtk_widget_get_tooltip_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetTooltipText is a wrapper around gtk_widget_set_tooltip_text().
func (v *Widget) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

/* DEPRECATED
// OverrideFont is a wrapper around gtk_widget_override_font().
func (v *Widget) OverrideFont(description string) {
	cstr := C.CString(description)
	defer C.free(unsafe.Pointer(cstr))
	c := C.pango_font_description_from_string(cstr)
	C.gtk_widget_override_font(v.native(), c)
}
*/

// GetHAlign is a wrapper around gtk_widget_get_halign().
func (v *Widget) GetHAlign() Align {
	c := C.gtk_widget_get_halign(v.native())
	return Align(c)
}

// SetHAlign is a wrapper around gtk_widget_set_halign().
func (v *Widget) SetHAlign(align Align) {
	C.gtk_widget_set_halign(v.native(), C.GtkAlign(align))
}

// GetVAlign is a wrapper around gtk_widget_get_valign().
func (v *Widget) GetVAlign() Align {
	c := C.gtk_widget_get_valign(v.native())
	return Align(c)
}

// SetVAlign is a wrapper around gtk_widget_set_valign().
func (v *Widget) SetVAlign(align Align) {
	C.gtk_widget_set_valign(v.native(), C.GtkAlign(align))
}

// GetMarginTop is a wrapper around gtk_widget_get_margin_top().
func (v *Widget) GetMarginTop() int {
	c := C.gtk_widget_get_margin_top(v.native())
	return int(c)
}

// SetMarginTop is a wrapper around gtk_widget_set_margin_top().
func (v *Widget) SetMarginTop(margin int) {
	C.gtk_widget_set_margin_top(v.native(), C.gint(margin))
}

// GetMarginBottom is a wrapper around gtk_widget_get_margin_bottom().
func (v *Widget) GetMarginBottom() int {
	c := C.gtk_widget_get_margin_bottom(v.native())
	return int(c)
}

// SetMarginBottom is a wrapper around gtk_widget_set_margin_bottom().
func (v *Widget) SetMarginBottom(margin int) {
	C.gtk_widget_set_margin_bottom(v.native(), C.gint(margin))
}

// GetHExpand is a wrapper around gtk_widget_get_hexpand().
func (v *Widget) GetHExpand() bool {
	c := C.gtk_widget_get_hexpand(v.native())
	return gobool(c)
}

// SetHExpand is a wrapper around gtk_widget_set_hexpand().
func (v *Widget) SetHExpand(expand bool) {
	C.gtk_widget_set_hexpand(v.native(), gbool(expand))
}

// GetVExpand is a wrapper around gtk_widget_get_vexpand().
func (v *Widget) GetVExpand() bool {
	c := C.gtk_widget_get_vexpand(v.native())
	return gobool(c)
}

// SetVExpand is a wrapper around gtk_widget_set_vexpand().
func (v *Widget) SetVExpand(expand bool) {
	C.gtk_widget_set_vexpand(v.native(), gbool(expand))
}

// SetVisual is a wrapper around gtk_widget_set_visual().
func (v *Widget) SetVisual(visual *gdk.Visual) {
	C.gtk_widget_set_visual(v.native(),
		(*C.GdkVisual)(unsafe.Pointer(visual.Native())))
}

// SetAppPaintable is a wrapper around gtk_widget_set_app_paintable().
func (v *Widget) SetAppPaintable(paintable bool) {
	C.gtk_widget_set_app_paintable(v.native(), gbool(paintable))
}

// GetAppPaintable is a wrapper around gtk_widget_get_app_paintable().
func (v *Widget) GetAppPaintable() bool {
	c := C.gtk_widget_get_app_paintable(v.native())
	return gobool(c)
}

// QueueDraw is a wrapper around gtk_widget_queue_draw().
func (v *Widget) QueueDraw() {
	C.gtk_widget_queue_draw(v.native())
}

//void
//gtk_widget_init_template (GtkWidget *widget);

//Creates and initializes child widgets defined in templates. This function must be called in the instance initializer for any class which assigned itself a template using gtk_widget_class_set_template()
//It is important to call this function in the instance initializer of a GtkWidget subclass and not in GObject.constructed() or GObject.constructor() for two reasons.
//One reason is that generally derived widgets will assume that parent class composite widgets have been created in their instance initializers.
//Another reason is that when calling g_object_new() on a widget with composite templates, it’s important to build the composite widgets before the construct properties are set. Properties passed to g_object_new() should take precedence over properties set in the private template XML.
func (v *Widget) InitTemplate() {
	C.gtk_widget_init_template(v.native())
}

/* FIXME
//void
//gtk_widget_class_set_template (GtkWidgetClass *widget_class,
//                               GBytes *template_bytes);

//This should be called at class initialization time to specify the GtkBuilder XML to be used to extend a widget.
//For convenience, gtk_widget_class_set_template_from_resource() is also provided.
//Note that any class that installs templates must call gtk_widget_init_template() in the widget’s instance initializer.
func (v *Widget) SetTemplate(template_bytes []byte) {
	gtemplate_bytes := make([]C.Gbytes, len(template_bytes))
	defer C.free(unsafe.Pointer(gtemplate_bytes))
	C.gtk_widget_class_set_template(v.native(), gtemplate_bytes)
}
*/

/* FIXME
//void
//gtk_widget_class_set_template_from_resource
//                               (GtkWidgetClass *widget_class,
//                                const gchar *resource_name);

//A convenience function to call gtk_widget_class_set_template().
//Note that any class that installs templates must call gtk_widget_init_template() in the widget’s instance initializer.
func (v *Widget) SetTemplateFromResource(resource_name string) {
	cstr := C.CString(resource_name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_class_set_template_from_resource(v.native(), (*C.gchar)(cstr))
}
*/


/* FIXME
//GObject *
//gtk_widget_get_template_child (GtkWidget *widget,
//                               GType widget_type,
//                               const gchar *name);

//Fetch an object build from the template XML for widget_type in this widget instance.
//This will only report children which were previously declared with gtk_widget_class_bind_template_child_full() or one of its variants.
//This function is only meant to be called for code which is private to the widget_type which declared the child and is meant for language bindings which cannot easily make use of the GObject structure offsets.
func (v *Widget) GetTemplaceChild(widget_type GType, name string) *glib.Object {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_widget_get_template_child(v.native(), GWidgetType, (*C.gchar)(cstr))
	return Object(c)
}

//void
//gtk_widget_class_bind_template_child_full
//                               (GtkWidgetClass *widget_class,
//                                const gchar *name,
//                                gboolean internal_child,
//                                gssize struct_offset);

//Automatically assign an object declared in the class template XML to be set to a location on a freshly built instance’s private data, or alternatively accessible via gtk_widget_get_template_child().
//The struct can point either into the public instance, then you should use G_STRUCT_OFFSET(WidgetType, member) for struct_offset , or in the private struct, then you should use G_PRIVATE_OFFSET(WidgetType, member).
//An explicit strong reference will be held automatically for the duration of your instance’s life cycle, it will be released automatically when GObjectClass.dispose() runs on your instance and if a struct_offset that is != 0 is specified, then the automatic location in your instance public or private data will be set to NULL. You can however access an automated child pointer the first time your classes GObjectClass.dispose() runs, or alternatively in GtkWidgetClass.destroy().
//If internal_child is specified, GtkBuildableIface.get_internal_child() will be automatically implemented by the GtkWidget class so there is no need to implement it manually.
//The wrapper macros gtk_widget_class_bind_template_child(), gtk_widget_class_bind_template_child_internal(), gtk_widget_class_bind_template_child_private() and gtk_widget_class_bind_template_child_internal_private() might be more convenient to use.
//Note that this must be called from a composite widget classes class initializer after calling gtk_widget_class_set_template().
func (v *Widget) BindTemplateChildFull(name string, internal_child bool, struct_offset uint) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_class_bind_template_child_full(v.native(), (*C.gchar)(cstr), gbool(internal_child), (C.gssize)(struct_offset))
}

*/
