// package zxdg acts as a client for the xdg_shell_unstable_v6 wayland protocol.

// generated by wl-scanner
// https://github.com/dkolbly/wl-scanner
// from: https://raw.githubusercontent.com/wayland-project/wayland-protocols/master/unstable/xdg-shell/xdg-shell-unstable-v6.xml
// on 2018-02-19 14:50:40 -0600
package zxdg

import (
	"sync"

	"github.com/dkolbly/wl"
	"golang.org/x/net/context"
)

type ShellPingEvent struct {
	EventContext context.Context
	Serial       uint32
}

type ShellPingHandler interface {
	HandleShellPing(ShellPingEvent)
}

func (p *Shell) AddPingHandler(h ShellPingHandler) {
	if h != nil {
		p.mu.Lock()
		p.pingHandlers = append(p.pingHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Shell) RemovePingHandler(h ShellPingHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.pingHandlers {
		if e == h {
			p.pingHandlers = append(p.pingHandlers[:i], p.pingHandlers[i+1:]...)
			break
		}
	}
}

func (p *Shell) Dispatch(ctx context.Context, event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.pingHandlers) > 0 {
			ev := ShellPingEvent{}
			ev.EventContext = ctx
			ev.Serial = event.Uint32()
			p.mu.RLock()
			for _, h := range p.pingHandlers {
				h.HandleShellPing(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type Shell struct {
	wl.BaseProxy
	mu           sync.RWMutex
	pingHandlers []ShellPingHandler
}

func NewShell(ctx *wl.Context) *Shell {
	ret := new(Shell)
	ctx.Register(ret)
	return ret
}

// Destroy will destroy xdg_shell.
//
//
// Destroy this xdg_shell object.
//
// Destroying a bound xdg_shell object while there are surfaces
// still alive created by this xdg_shell object instance is illegal
// and will result in a protocol error.
//
func (p *Shell) Destroy() error {
	return p.Context().SendRequest(p, 0)
}

// CreatePositioner will create a positioner object.
//
//
// Create a positioner object. A positioner object is used to position
// surfaces relative to some parent surface. See the interface description
// and xdg_surface.get_popup for details.
//
func (p *Shell) CreatePositioner() (*Positioner, error) {
	ret := NewPositioner(p.Context())
	return ret, p.Context().SendRequest(p, 1, wl.Proxy(ret))
}

// GetXdgSurface will create a shell surface from a surface.
//
//
// This creates an xdg_surface for the given surface. While xdg_surface
// itself is not a role, the corresponding surface may only be assigned
// a role extending xdg_surface, such as xdg_toplevel or xdg_popup.
//
// This creates an xdg_surface for the given surface. An xdg_surface is
// used as basis to define a role to a given surface, such as xdg_toplevel
// or xdg_popup. It also manages functionality shared between xdg_surface
// based surface roles.
//
// See the documentation of xdg_surface for more details about what an
// xdg_surface is and how it is used.
//
func (p *Shell) GetXdgSurface(surface *wl.Surface) (*Surface, error) {
	ret := NewSurface(p.Context())
	return ret, p.Context().SendRequest(p, 2, wl.Proxy(ret), surface)
}

// Pong will respond to a ping event.
//
//
// A client must respond to a ping event with a pong request or
// the client may be deemed unresponsive. See xdg_shell.ping.
//
func (p *Shell) Pong(serial uint32) error {
	return p.Context().SendRequest(p, 3, serial)
}

const (
	ShellErrorRole                = 0
	ShellErrorDefunctSurfaces     = 1
	ShellErrorNotTheTopmostPopup  = 2
	ShellErrorInvalidPopupParent  = 3
	ShellErrorInvalidSurfaceState = 4
	ShellErrorInvalidPositioner   = 5
)

type Positioner struct {
	wl.BaseProxy
}

func NewPositioner(ctx *wl.Context) *Positioner {
	ret := new(Positioner)
	ctx.Register(ret)
	return ret
}

// Destroy will destroy the xdg_positioner object.
//
//
// Notify the compositor that the xdg_positioner will no longer be used.
//
func (p *Positioner) Destroy() error {
	return p.Context().SendRequest(p, 0)
}

// SetSize will set the size of the to-be positioned rectangle.
//
//
// Set the size of the surface that is to be positioned with the positioner
// object. The size is in surface-local coordinates and corresponds to the
// window geometry. See xdg_surface.set_window_geometry.
//
// If a zero or negative size is set the invalid_input error is raised.
//
func (p *Positioner) SetSize(width int32, height int32) error {
	return p.Context().SendRequest(p, 1, width, height)
}

// SetAnchorRect will set the anchor rectangle within the parent surface.
//
//
// Specify the anchor rectangle within the parent surface that the child
// surface will be placed relative to. The rectangle is relative to the
// window geometry as defined by xdg_surface.set_window_geometry of the
// parent surface. The rectangle must be at least 1x1 large.
//
// When the xdg_positioner object is used to position a child surface, the
// anchor rectangle may not extend outside the window geometry of the
// positioned child's parent surface.
//
// If a zero or negative size is set the invalid_input error is raised.
//
func (p *Positioner) SetAnchorRect(x int32, y int32, width int32, height int32) error {
	return p.Context().SendRequest(p, 2, x, y, width, height)
}

// SetAnchor will set anchor rectangle anchor edges.
//
//
// Defines a set of edges for the anchor rectangle. These are used to
// derive an anchor point that the child surface will be positioned
// relative to. If two orthogonal edges are specified (e.g. 'top' and
// 'left'), then the anchor point will be the intersection of the edges
// (e.g. the top left position of the rectangle); otherwise, the derived
// anchor point will be centered on the specified edge, or in the center of
// the anchor rectangle if no edge is specified.
//
// If two parallel anchor edges are specified (e.g. 'left' and 'right'),
// the invalid_input error is raised.
//
func (p *Positioner) SetAnchor(anchor uint32) error {
	return p.Context().SendRequest(p, 3, anchor)
}

// SetGravity will set child surface gravity.
//
//
// Defines in what direction a surface should be positioned, relative to
// the anchor point of the parent surface. If two orthogonal gravities are
// specified (e.g. 'bottom' and 'right'), then the child surface will be
// placed in the specified direction; otherwise, the child surface will be
// centered over the anchor point on any axis that had no gravity
// specified.
//
// If two parallel gravities are specified (e.g. 'left' and 'right'), the
// invalid_input error is raised.
//
func (p *Positioner) SetGravity(gravity uint32) error {
	return p.Context().SendRequest(p, 4, gravity)
}

// SetConstraintAdjustment will set the adjustment to be done when constrained.
//
//
// Specify how the window should be positioned if the originally intended
// position caused the surface to be constrained, meaning at least
// partially outside positioning boundaries set by the compositor. The
// adjustment is set by constructing a bitmask describing the adjustment to
// be made when the surface is constrained on that axis.
//
// If no bit for one axis is set, the compositor will assume that the child
// surface should not change its position on that axis when constrained.
//
// If more than one bit for one axis is set, the order of how adjustments
// are applied is specified in the corresponding adjustment descriptions.
//
// The default adjustment is none.
//
func (p *Positioner) SetConstraintAdjustment(constraint_adjustment uint32) error {
	return p.Context().SendRequest(p, 5, constraint_adjustment)
}

// SetOffset will set surface position offset.
//
//
// Specify the surface position offset relative to the position of the
// anchor on the anchor rectangle and the anchor on the surface. For
// example if the anchor of the anchor rectangle is at (x, y), the surface
// has the gravity bottom|right, and the offset is (ox, oy), the calculated
// surface position will be (x + ox, y + oy). The offset position of the
// surface is the one used for constraint testing. See
// set_constraint_adjustment.
//
// An example use case is placing a popup menu on top of a user interface
// element, while aligning the user interface element of the parent surface
// with some user interface element placed somewhere in the popup surface.
//
func (p *Positioner) SetOffset(x int32, y int32) error {
	return p.Context().SendRequest(p, 6, x, y)
}

const (
	PositionerErrorInvalidInput = 0
)

const (
	PositionerAnchorNone   = 0
	PositionerAnchorTop    = 1
	PositionerAnchorBottom = 2
	PositionerAnchorLeft   = 4
	PositionerAnchorRight  = 8
)

const (
	PositionerGravityNone   = 0
	PositionerGravityTop    = 1
	PositionerGravityBottom = 2
	PositionerGravityLeft   = 4
	PositionerGravityRight  = 8
)

const (
	PositionerConstraintAdjustmentNone    = 0
	PositionerConstraintAdjustmentSlideX  = 1
	PositionerConstraintAdjustmentSlideY  = 2
	PositionerConstraintAdjustmentFlipX   = 4
	PositionerConstraintAdjustmentFlipY   = 8
	PositionerConstraintAdjustmentResizeX = 16
	PositionerConstraintAdjustmentResizeY = 32
)

type SurfaceConfigureEvent struct {
	EventContext context.Context
	Serial       uint32
}

type SurfaceConfigureHandler interface {
	HandleSurfaceConfigure(SurfaceConfigureEvent)
}

func (p *Surface) AddConfigureHandler(h SurfaceConfigureHandler) {
	if h != nil {
		p.mu.Lock()
		p.configureHandlers = append(p.configureHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Surface) RemoveConfigureHandler(h SurfaceConfigureHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.configureHandlers {
		if e == h {
			p.configureHandlers = append(p.configureHandlers[:i], p.configureHandlers[i+1:]...)
			break
		}
	}
}

func (p *Surface) Dispatch(ctx context.Context, event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.configureHandlers) > 0 {
			ev := SurfaceConfigureEvent{}
			ev.EventContext = ctx
			ev.Serial = event.Uint32()
			p.mu.RLock()
			for _, h := range p.configureHandlers {
				h.HandleSurfaceConfigure(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type Surface struct {
	wl.BaseProxy
	mu                sync.RWMutex
	configureHandlers []SurfaceConfigureHandler
}

func NewSurface(ctx *wl.Context) *Surface {
	ret := new(Surface)
	ctx.Register(ret)
	return ret
}

// Destroy will destroy the xdg_surface.
//
//
// Destroy the xdg_surface object. An xdg_surface must only be destroyed
// after its role object has been destroyed.
//
func (p *Surface) Destroy() error {
	return p.Context().SendRequest(p, 0)
}

// GetToplevel will assign the xdg_toplevel surface role.
//
//
// This creates an xdg_toplevel object for the given xdg_surface and gives
// the associated wl_surface the xdg_toplevel role.
//
// See the documentation of xdg_toplevel for more details about what an
// xdg_toplevel is and how it is used.
//
func (p *Surface) GetToplevel() (*Toplevel, error) {
	ret := NewToplevel(p.Context())
	return ret, p.Context().SendRequest(p, 1, wl.Proxy(ret))
}

// GetPopup will assign the xdg_popup surface role.
//
//
// This creates an xdg_popup object for the given xdg_surface and gives the
// associated wl_surface the xdg_popup role.
//
// See the documentation of xdg_popup for more details about what an
// xdg_popup is and how it is used.
//
func (p *Surface) GetPopup(parent *Surface, positioner *Positioner) (*Popup, error) {
	ret := NewPopup(p.Context())
	return ret, p.Context().SendRequest(p, 2, wl.Proxy(ret), parent, positioner)
}

// SetWindowGeometry will set the new window geometry.
//
//
// The window geometry of a surface is its "visible bounds" from the
// user's perspective. Client-side decorations often have invisible
// portions like drop-shadows which should be ignored for the
// purposes of aligning, placing and constraining windows.
//
// The window geometry is double buffered, and will be applied at the
// time wl_surface.commit of the corresponding wl_surface is called.
//
// Once the window geometry of the surface is set, it is not possible to
// unset it, and it will remain the same until set_window_geometry is
// called again, even if a new subsurface or buffer is attached.
//
// If never set, the value is the full bounds of the surface,
// including any subsurfaces. This updates dynamically on every
// commit. This unset is meant for extremely simple clients.
//
// The arguments are given in the surface-local coordinate space of
// the wl_surface associated with this xdg_surface.
//
// The width and height must be greater than zero. Setting an invalid size
// will raise an error. When applied, the effective window geometry will be
// the set window geometry clamped to the bounding rectangle of the
// combined geometry of the surface of the xdg_surface and the associated
// subsurfaces.
//
func (p *Surface) SetWindowGeometry(x int32, y int32, width int32, height int32) error {
	return p.Context().SendRequest(p, 3, x, y, width, height)
}

// AckConfigure will ack a configure event.
//
//
// When a configure event is received, if a client commits the
// surface in response to the configure event, then the client
// must make an ack_configure request sometime before the commit
// request, passing along the serial of the configure event.
//
// For instance, for toplevel surfaces the compositor might use this
// information to move a surface to the top left only when the client has
// drawn itself for the maximized or fullscreen state.
//
// If the client receives multiple configure events before it
// can respond to one, it only has to ack the last configure event.
//
// A client is not required to commit immediately after sending
// an ack_configure request - it may even ack_configure several times
// before its next surface commit.
//
// A client may send multiple ack_configure requests before committing, but
// only the last request sent before a commit indicates which configure
// event the client really is responding to.
//
func (p *Surface) AckConfigure(serial uint32) error {
	return p.Context().SendRequest(p, 4, serial)
}

const (
	SurfaceErrorNotConstructed     = 1
	SurfaceErrorAlreadyConstructed = 2
	SurfaceErrorUnconfiguredBuffer = 3
)

type ToplevelConfigureEvent struct {
	EventContext context.Context
	Width        int32
	Height       int32
	States       []int32
}

type ToplevelConfigureHandler interface {
	HandleToplevelConfigure(ToplevelConfigureEvent)
}

func (p *Toplevel) AddConfigureHandler(h ToplevelConfigureHandler) {
	if h != nil {
		p.mu.Lock()
		p.configureHandlers = append(p.configureHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Toplevel) RemoveConfigureHandler(h ToplevelConfigureHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.configureHandlers {
		if e == h {
			p.configureHandlers = append(p.configureHandlers[:i], p.configureHandlers[i+1:]...)
			break
		}
	}
}

type ToplevelCloseEvent struct {
	EventContext context.Context
}

type ToplevelCloseHandler interface {
	HandleToplevelClose(ToplevelCloseEvent)
}

func (p *Toplevel) AddCloseHandler(h ToplevelCloseHandler) {
	if h != nil {
		p.mu.Lock()
		p.closeHandlers = append(p.closeHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Toplevel) RemoveCloseHandler(h ToplevelCloseHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.closeHandlers {
		if e == h {
			p.closeHandlers = append(p.closeHandlers[:i], p.closeHandlers[i+1:]...)
			break
		}
	}
}

func (p *Toplevel) Dispatch(ctx context.Context, event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.configureHandlers) > 0 {
			ev := ToplevelConfigureEvent{}
			ev.EventContext = ctx
			ev.Width = event.Int32()
			ev.Height = event.Int32()
			ev.States = event.Array()
			p.mu.RLock()
			for _, h := range p.configureHandlers {
				h.HandleToplevelConfigure(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.closeHandlers) > 0 {
			ev := ToplevelCloseEvent{}
			ev.EventContext = ctx
			p.mu.RLock()
			for _, h := range p.closeHandlers {
				h.HandleToplevelClose(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type Toplevel struct {
	wl.BaseProxy
	mu                sync.RWMutex
	configureHandlers []ToplevelConfigureHandler
	closeHandlers     []ToplevelCloseHandler
}

func NewToplevel(ctx *wl.Context) *Toplevel {
	ret := new(Toplevel)
	ctx.Register(ret)
	return ret
}

// Destroy will destroy the xdg_toplevel.
//
//
// Unmap and destroy the window. The window will be effectively
// hidden from the user's point of view, and all state like
// maximization, fullscreen, and so on, will be lost.
//
func (p *Toplevel) Destroy() error {
	return p.Context().SendRequest(p, 0)
}

// SetParent will set the parent of this surface.
//
//
// Set the "parent" of this surface. This window should be stacked
// above a parent. The parent surface must be mapped as long as this
// surface is mapped.
//
// Parent windows should be set on dialogs, toolboxes, or other
// "auxiliary" surfaces, so that the parent is raised when the dialog
// is raised.
//
func (p *Toplevel) SetParent(parent *Toplevel) error {
	return p.Context().SendRequest(p, 1, parent)
}

// SetTitle will set surface title.
//
//
// Set a short title for the surface.
//
// This string may be used to identify the surface in a task bar,
// window list, or other user interface elements provided by the
// compositor.
//
// The string must be encoded in UTF-8.
//
func (p *Toplevel) SetTitle(title string) error {
	return p.Context().SendRequest(p, 2, title)
}

// SetAppId will set application ID.
//
//
// Set an application identifier for the surface.
//
// The app ID identifies the general class of applications to which
// the surface belongs. The compositor can use this to group multiple
// surfaces together, or to determine how to launch a new application.
//
// For D-Bus activatable applications, the app ID is used as the D-Bus
// service name.
//
// The compositor shell will try to group application surfaces together
// by their app ID. As a best practice, it is suggested to select app
// ID's that match the basename of the application's .desktop file.
// For example, "org.freedesktop.FooViewer" where the .desktop file is
// "org.freedesktop.FooViewer.desktop".
//
// See the desktop-entry specification [0] for more details on
// application identifiers and how they relate to well-known D-Bus
// names and .desktop files.
//
// [0] http://standards.freedesktop.org/desktop-entry-spec/
//
func (p *Toplevel) SetAppId(app_id string) error {
	return p.Context().SendRequest(p, 3, app_id)
}

// ShowWindowMenu will show the window menu.
//
//
// Clients implementing client-side decorations might want to show
// a context menu when right-clicking on the decorations, giving the
// user a menu that they can use to maximize or minimize the window.
//
// This request asks the compositor to pop up such a window menu at
// the given position, relative to the local surface coordinates of
// the parent surface. There are no guarantees as to what menu items
// the window menu contains.
//
// This request must be used in response to some sort of user action
// like a button press, key press, or touch down event.
//
func (p *Toplevel) ShowWindowMenu(seat *wl.Seat, serial uint32, x int32, y int32) error {
	return p.Context().SendRequest(p, 4, seat, serial, x, y)
}

// Move will start an interactive move.
//
//
// Start an interactive, user-driven move of the surface.
//
// This request must be used in response to some sort of user action
// like a button press, key press, or touch down event. The passed
// serial is used to determine the type of interactive move (touch,
// pointer, etc).
//
// The server may ignore move requests depending on the state of
// the surface (e.g. fullscreen or maximized), or if the passed serial
// is no longer valid.
//
// If triggered, the surface will lose the focus of the device
// (wl_pointer, wl_touch, etc) used for the move. It is up to the
// compositor to visually indicate that the move is taking place, such as
// updating a pointer cursor, during the move. There is no guarantee
// that the device focus will return when the move is completed.
//
func (p *Toplevel) Move(seat *wl.Seat, serial uint32) error {
	return p.Context().SendRequest(p, 5, seat, serial)
}

// Resize will start an interactive resize.
//
//
// Start a user-driven, interactive resize of the surface.
//
// This request must be used in response to some sort of user action
// like a button press, key press, or touch down event. The passed
// serial is used to determine the type of interactive resize (touch,
// pointer, etc).
//
// The server may ignore resize requests depending on the state of
// the surface (e.g. fullscreen or maximized).
//
// If triggered, the client will receive configure events with the
// "resize" state enum value and the expected sizes. See the "resize"
// enum value for more details about what is required. The client
// must also acknowledge configure events using "ack_configure". After
// the resize is completed, the client will receive another "configure"
// event without the resize state.
//
// If triggered, the surface also will lose the focus of the device
// (wl_pointer, wl_touch, etc) used for the resize. It is up to the
// compositor to visually indicate that the resize is taking place,
// such as updating a pointer cursor, during the resize. There is no
// guarantee that the device focus will return when the resize is
// completed.
//
// The edges parameter specifies how the surface should be resized,
// and is one of the values of the resize_edge enum. The compositor
// may use this information to update the surface position for
// example when dragging the top left corner. The compositor may also
// use this information to adapt its behavior, e.g. choose an
// appropriate cursor image.
//
func (p *Toplevel) Resize(seat *wl.Seat, serial uint32, edges uint32) error {
	return p.Context().SendRequest(p, 6, seat, serial, edges)
}

// SetMaxSize will set the maximum size.
//
//
// Set a maximum size for the window.
//
// The client can specify a maximum size so that the compositor does
// not try to configure the window beyond this size.
//
// The width and height arguments are in window geometry coordinates.
// See xdg_surface.set_window_geometry.
//
// Values set in this way are double-buffered. They will get applied
// on the next commit.
//
// The compositor can use this information to allow or disallow
// different states like maximize or fullscreen and draw accurate
// animations.
//
// Similarly, a tiling window manager may use this information to
// place and resize client windows in a more effective way.
//
// The client should not rely on the compositor to obey the maximum
// size. The compositor may decide to ignore the values set by the
// client and request a larger size.
//
// If never set, or a value of zero in the request, means that the
// client has no expected maximum size in the given dimension.
// As a result, a client wishing to reset the maximum size
// to an unspecified state can use zero for width and height in the
// request.
//
// Requesting a maximum size to be smaller than the minimum size of
// a surface is illegal and will result in a protocol error.
//
// The width and height must be greater than or equal to zero. Using
// strictly negative values for width and height will result in a
// protocol error.
//
func (p *Toplevel) SetMaxSize(width int32, height int32) error {
	return p.Context().SendRequest(p, 7, width, height)
}

// SetMinSize will set the minimum size.
//
//
// Set a minimum size for the window.
//
// The client can specify a minimum size so that the compositor does
// not try to configure the window below this size.
//
// The width and height arguments are in window geometry coordinates.
// See xdg_surface.set_window_geometry.
//
// Values set in this way are double-buffered. They will get applied
// on the next commit.
//
// The compositor can use this information to allow or disallow
// different states like maximize or fullscreen and draw accurate
// animations.
//
// Similarly, a tiling window manager may use this information to
// place and resize client windows in a more effective way.
//
// The client should not rely on the compositor to obey the minimum
// size. The compositor may decide to ignore the values set by the
// client and request a smaller size.
//
// If never set, or a value of zero in the request, means that the
// client has no expected minimum size in the given dimension.
// As a result, a client wishing to reset the minimum size
// to an unspecified state can use zero for width and height in the
// request.
//
// Requesting a minimum size to be larger than the maximum size of
// a surface is illegal and will result in a protocol error.
//
// The width and height must be greater than or equal to zero. Using
// strictly negative values for width and height will result in a
// protocol error.
//
func (p *Toplevel) SetMinSize(width int32, height int32) error {
	return p.Context().SendRequest(p, 8, width, height)
}

// SetMaximized will maximize the window.
//
//
// Maximize the surface.
//
// After requesting that the surface should be maximized, the compositor
// will respond by emitting a configure event with the "maximized" state
// and the required window geometry. The client should then update its
// content, drawing it in a maximized state, i.e. without shadow or other
// decoration outside of the window geometry. The client must also
// acknowledge the configure when committing the new content (see
// ack_configure).
//
// It is up to the compositor to decide how and where to maximize the
// surface, for example which output and what region of the screen should
// be used.
//
// If the surface was already maximized, the compositor will still emit
// a configure event with the "maximized" state.
//
func (p *Toplevel) SetMaximized() error {
	return p.Context().SendRequest(p, 9)
}

// UnsetMaximized will unmaximize the window.
//
//
// Unmaximize the surface.
//
// After requesting that the surface should be unmaximized, the compositor
// will respond by emitting a configure event without the "maximized"
// state. If available, the compositor will include the window geometry
// dimensions the window had prior to being maximized in the configure
// request. The client must then update its content, drawing it in a
// regular state, i.e. potentially with shadow, etc. The client must also
// acknowledge the configure when committing the new content (see
// ack_configure).
//
// It is up to the compositor to position the surface after it was
// unmaximized; usually the position the surface had before maximizing, if
// applicable.
//
// If the surface was already not maximized, the compositor will still
// emit a configure event without the "maximized" state.
//
func (p *Toplevel) UnsetMaximized() error {
	return p.Context().SendRequest(p, 10)
}

// SetFullscreen will set the window as fullscreen on a monitor.
//
//
// Make the surface fullscreen.
//
// You can specify an output that you would prefer to be fullscreen.
// If this value is NULL, it's up to the compositor to choose which
// display will be used to map this surface.
//
// If the surface doesn't cover the whole output, the compositor will
// position the surface in the center of the output and compensate with
// black borders filling the rest of the output.
//
func (p *Toplevel) SetFullscreen(output *wl.Output) error {
	return p.Context().SendRequest(p, 11, output)
}

// UnsetFullscreen will .
//
//
func (p *Toplevel) UnsetFullscreen() error {
	return p.Context().SendRequest(p, 12)
}

// SetMinimized will set the window as minimized.
//
//
// Request that the compositor minimize your surface. There is no
// way to know if the surface is currently minimized, nor is there
// any way to unset minimization on this surface.
//
// If you are looking to throttle redrawing when minimized, please
// instead use the wl_surface.frame event for this, as this will
// also work with live previews on windows in Alt-Tab, Expose or
// similar compositor features.
//
func (p *Toplevel) SetMinimized() error {
	return p.Context().SendRequest(p, 13)
}

const (
	ToplevelResizeEdgeNone        = 0
	ToplevelResizeEdgeTop         = 1
	ToplevelResizeEdgeBottom      = 2
	ToplevelResizeEdgeLeft        = 4
	ToplevelResizeEdgeTopLeft     = 5
	ToplevelResizeEdgeBottomLeft  = 6
	ToplevelResizeEdgeRight       = 8
	ToplevelResizeEdgeTopRight    = 9
	ToplevelResizeEdgeBottomRight = 10
)

const (
	ToplevelStateMaximized  = 1
	ToplevelStateFullscreen = 2
	ToplevelStateResizing   = 3
	ToplevelStateActivated  = 4
)

type PopupConfigureEvent struct {
	EventContext context.Context
	X            int32
	Y            int32
	Width        int32
	Height       int32
}

type PopupConfigureHandler interface {
	HandlePopupConfigure(PopupConfigureEvent)
}

func (p *Popup) AddConfigureHandler(h PopupConfigureHandler) {
	if h != nil {
		p.mu.Lock()
		p.configureHandlers = append(p.configureHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Popup) RemoveConfigureHandler(h PopupConfigureHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.configureHandlers {
		if e == h {
			p.configureHandlers = append(p.configureHandlers[:i], p.configureHandlers[i+1:]...)
			break
		}
	}
}

type PopupPopupDoneEvent struct {
	EventContext context.Context
}

type PopupPopupDoneHandler interface {
	HandlePopupPopupDone(PopupPopupDoneEvent)
}

func (p *Popup) AddPopupDoneHandler(h PopupPopupDoneHandler) {
	if h != nil {
		p.mu.Lock()
		p.popupDoneHandlers = append(p.popupDoneHandlers, h)
		p.mu.Unlock()
	}
}

func (p *Popup) RemovePopupDoneHandler(h PopupPopupDoneHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.popupDoneHandlers {
		if e == h {
			p.popupDoneHandlers = append(p.popupDoneHandlers[:i], p.popupDoneHandlers[i+1:]...)
			break
		}
	}
}

func (p *Popup) Dispatch(ctx context.Context, event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.configureHandlers) > 0 {
			ev := PopupConfigureEvent{}
			ev.EventContext = ctx
			ev.X = event.Int32()
			ev.Y = event.Int32()
			ev.Width = event.Int32()
			ev.Height = event.Int32()
			p.mu.RLock()
			for _, h := range p.configureHandlers {
				h.HandlePopupConfigure(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.popupDoneHandlers) > 0 {
			ev := PopupPopupDoneEvent{}
			ev.EventContext = ctx
			p.mu.RLock()
			for _, h := range p.popupDoneHandlers {
				h.HandlePopupPopupDone(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type Popup struct {
	wl.BaseProxy
	mu                sync.RWMutex
	configureHandlers []PopupConfigureHandler
	popupDoneHandlers []PopupPopupDoneHandler
}

func NewPopup(ctx *wl.Context) *Popup {
	ret := new(Popup)
	ctx.Register(ret)
	return ret
}

// Destroy will remove xdg_popup interface.
//
//
// This destroys the popup. Explicitly destroying the xdg_popup
// object will also dismiss the popup, and unmap the surface.
//
// If this xdg_popup is not the "topmost" popup, a protocol error
// will be sent.
//
func (p *Popup) Destroy() error {
	return p.Context().SendRequest(p, 0)
}

// Grab will make the popup take an explicit grab.
//
//
// This request makes the created popup take an explicit grab. An explicit
// grab will be dismissed when the user dismisses the popup, or when the
// client destroys the xdg_popup. This can be done by the user clicking
// outside the surface, using the keyboard, or even locking the screen
// through closing the lid or a timeout.
//
// If the compositor denies the grab, the popup will be immediately
// dismissed.
//
// This request must be used in response to some sort of user action like a
// button press, key press, or touch down event. The serial number of the
// event should be passed as 'serial'.
//
// The parent of a grabbing popup must either be an xdg_toplevel surface or
// another xdg_popup with an explicit grab. If the parent is another
// xdg_popup it means that the popups are nested, with this popup now being
// the topmost popup.
//
// Nested popups must be destroyed in the reverse order they were created
// in, e.g. the only popup you are allowed to destroy at all times is the
// topmost one.
//
// When compositors choose to dismiss a popup, they may dismiss every
// nested grabbing popup as well. When a compositor dismisses popups, it
// will follow the same dismissing order as required from the client.
//
// The parent of a grabbing popup must either be another xdg_popup with an
// active explicit grab, or an xdg_popup or xdg_toplevel, if there are no
// explicit grabs already taken.
//
// If the topmost grabbing popup is destroyed, the grab will be returned to
// the parent of the popup, if that parent previously had an explicit grab.
//
// If the parent is a grabbing popup which has already been dismissed, this
// popup will be immediately dismissed. If the parent is a popup that did
// not take an explicit grab, an error will be raised.
//
// During a popup grab, the client owning the grab will receive pointer
// and touch events for all their surfaces as normal (similar to an
// "owner-events" grab in X11 parlance), while the top most grabbing popup
// will always have keyboard focus.
//
func (p *Popup) Grab(seat *wl.Seat, serial uint32) error {
	return p.Context().SendRequest(p, 1, seat, serial)
}

const (
	PopupErrorInvalidGrab = 0
)
