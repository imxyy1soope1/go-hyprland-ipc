package event

import "strings"

type Data interface {
	f()
}

type Workspace struct {
	WorkspaceName string
}

func (Workspace) f() {
}

func parseWorkspace(data string) Data {
	return Workspace{WorkspaceName: data}
}

type FocusedMon struct {
	MonitorName, WorkspaceName string
}

func (FocusedMon) f() {
}

func parseFocusedMon(data string) Data {
	parts := strings.Split(data, ",")
	return FocusedMon{MonitorName: parts[0], WorkspaceName: parts[1]}
}

type ActiveWindow struct {
	WindowClass, WindowTitle string
}

func (ActiveWindow) f() {
}

func parseActiveWindow(data string) Data {
	parts := strings.Split(data, ",")
	return ActiveWindow{WindowClass: parts[0], WindowTitle: parts[1]}
}

type ActiveWindowV2 struct {
	WindowAddress string
}

func (ActiveWindowV2) f() {
}

func parseActiveWindowV2(data string) Data {
	return ActiveWindowV2{WindowAddress: data}
}

type Fullscreen struct {
	Fullscreened bool
}

func (Fullscreen) f() {
}

func parseFullscreen(data string) Data {
	return Fullscreen{Fullscreened: data == "1"}
}

type MonitorRemoved struct {
	MonitorName string
}

func (MonitorRemoved) f() {
}

func parseMonitorRemoved(data string) Data {
	return MonitorRemoved{MonitorName: data}
}

type MonitorAdded struct {
	MonitorName string
}

func (MonitorAdded) f() {
}

func parseMonitorAdded(data string) Data {
	return MonitorAdded{MonitorName: data}
}

type CreateWorkspace struct {
	WorkspaceName string
}

func (CreateWorkspace) f() {
}

func parseCreateWorkspace(data string) Data {
	return CreateWorkspace{WorkspaceName: data}
}

type DestroyWorkspace struct {
	WorkspaceName string
}

func (DestroyWorkspace) f() {
}

func parseDestroyWorkspace(data string) Data {
	return DestroyWorkspace{WorkspaceName: data}
}

type MoveWorkspace struct {
	WorkspaceName, MonitorName string
}

func (MoveWorkspace) f() {
}

func parseMoveWorkspace(data string) Data {
	parts := strings.Split(data, ",")
	return MoveWorkspace{WorkspaceName: parts[0], MonitorName: parts[1]}
}

type ActiveLayout struct {
	KeyboardName, LayoutName string
}

func (ActiveLayout) f() {
}

func parseActiveLayout(data string) Data {
	parts := strings.Split(data, ",")
	return ActiveLayout{KeyboardName: parts[0], LayoutName: parts[1]}
}

type OpenWindow struct {
	WindowAddress, WorkspaceName, WindowClass, WindowTitle string
}

func (OpenWindow) f() {
}

func parseOpenWindow(data string) Data {
	parts := strings.Split(data, ",")
	return OpenWindow{
		WindowAddress: parts[0],
		WorkspaceName: parts[1],
		WindowClass:   parts[2],
		WindowTitle:   parts[3],
	}
}

type CloseWindow struct {
	WindowAddress string
}

func (CloseWindow) f() {
}

func parseCloseWindow(data string) Data {
	return CloseWindow{WindowAddress: data}
}

type MoveWindow struct {
	WindowAddress, WorkspaceName string
}

func (MoveWindow) f() {
}

func parseMoveWindow(data string) Data {
	parts := strings.Split(data, ",")
	return MoveWindow{WindowAddress: parts[0], WorkspaceName: parts[1]}
}

type OpenLayer struct {
	Namespace string
}

func (OpenLayer) f() {
}

func parseOpenLayer(data string) Data {
	return OpenLayer{Namespace: data}
}

type CloseLayer struct {
	Namespace string
}

func (CloseLayer) f() {
}

func parseCloseLayer(data string) Data {
	return CloseLayer{Namespace: data}
}

type Submap struct {
	SubmapName string
}

func (Submap) f() {
}

func parseSubmap(data string) Data {
	return Submap{SubmapName: data}
}

type ChangeFloatingMode struct {
	WindowAddress string
	Floating      bool
}

func (ChangeFloatingMode) f() {
}

func parseChangeFloatingMode(data string) Data {
	parts := strings.Split(data, ",")
	return ChangeFloatingMode{WindowAddress: parts[0], Floating: parts[1] == "1"}
}

type Urgent struct {
	WindowAddress string
}

func (Urgent) f() {
}

func parseUrgent(data string) Data {
	return Urgent{WindowAddress: data}
}

type Minimize struct {
	WindowAddress string
	Minimized     bool
}

func (Minimize) f() {
}

func parseMinimize(data string) Data {
	parts := strings.Split(data, ",")
	return Minimize{WindowAddress: parts[0], Minimized: parts[1] == "1"}
}

type Screencast struct {
	State bool
	// false for monitor, true for window
	Owner bool
}

func (Screencast) f() {
}

func parseScreencast(data string) Data {
	parts := strings.Split(data, ",")
	return Screencast{State: parts[0] == "1", Owner: parts[1] == "1"}
}

type WindowTitle struct {
	WindowAddress string
}

func (WindowTitle) f() {
}

func parseWindowTitle(data string) Data {
	return WindowTitle{WindowAddress: data}
}

var eventToParser = map[Event]func(string) Data{
	WORKSPACE:          parseWorkspace,
	FOCUSEDMON:         parseFocusedMon,
	ACTIVEWINDOW:       parseActiveWindow,
	ACTIVEWINDOWV2:     parseActiveWindowV2,
	FULLSCREEN:         parseFullscreen,
	MONITORREMOVED:     parseMonitorRemoved,
	MONITORADDED:       parseMonitorAdded,
	CREATEWORKSPACE:    parseCreateWorkspace,
	DESTROYWORKSPACE:   parseDestroyWorkspace,
	MOVEWORKSPACE:      parseMoveWorkspace,
	ACTIVELAYOUT:       parseActiveLayout,
	OPENWINDOW:         parseOpenWindow,
	CLOSEWINDOW:        parseCloseWindow,
	MOVEWINDOW:         parseMoveWindow,
	OPENLAYER:          parseOpenLayer,
	CLOSELAYER:         parseCloseLayer,
	SUBMAP:             parseSubmap,
	CHANGEFLOATINGMODE: parseChangeFloatingMode,
	URGENT:             parseUrgent,
	MINIMIZE:           parseMinimize,
	SCREENCAST:         parseScreencast,
	WINDOWTITLE:        parseWindowTitle,
}
