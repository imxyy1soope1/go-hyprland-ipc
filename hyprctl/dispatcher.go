package hyprctl

import (
	"errors"
	"fmt"
	"strconv"
)

/*  --------------------PARAM TYPE DEFINITION START--------------------  */

type Window interface {
	getWindow() string
}

type WindowWithClass struct {
	Class string
}

func (w WindowWithClass) getWindow() string {
	return w.Class
}

type WindowWithTitle struct {
	Title string
}

func (w WindowWithTitle) getWindow() string {
	return fmt.Sprintf("title:%s", w.Title)
}

type WindowWithPid struct {
	Pid uint
}

func (w WindowWithPid) getWindow() string {
	return fmt.Sprintf("pid:%d", w.Pid)
}

type WindowWithAddr struct {
	Addr string
}

func (w WindowWithAddr) getWindow() string {
	return fmt.Sprintf("address:%s", w.Addr)
}

type EmptyWindow struct {
}

func (w EmptyWindow) getWindow() string {
	return ""
}

type Workspace interface {
	getWorkspace() string
}

type WorkspaceWithId struct {
	Id int64
}

func (w WorkspaceWithId) getWorkspace() string {
	return strconv.FormatInt(w.Id, 10)
}

type WorkspaceWithRelativeId struct {
	Id     int64
	Opened bool
}

func (w WorkspaceWithRelativeId) getWorkspace() string {
	if w.Opened {
		return fmt.Sprintf("e%+d", w.Id)
	} else {
		return fmt.Sprintf("%+d", w.Id)
	}
}

type WorkspaceWithMonitorRelativeId struct {
	Id           uint64
	IncludeEmpty bool
}

func (w WorkspaceWithMonitorRelativeId) getWorkspace() string {
	if w.IncludeEmpty {
		return fmt.Sprintf("m%+d", w.Id)
	} else {
		return fmt.Sprintf("r%+d", w.Id)
	}
}

type Direction string

const (
	DirectionLeft  Direction = "l"
	DirectionRight Direction = "r"
	DirectionUp    Direction = "u"
	DirectionDown  Direction = "d"
)

type Monitor interface {
	getMonitor() string
}

type MonitorWithDirection struct {
	Direction
}

func (m MonitorWithDirection) getMonitor() string {
	return string(m.Direction)
}

type MonitorWithId struct {
	Id uint
}

func (m MonitorWithId) getMonitor() string {
	return strconv.FormatUint(uint64(m.Id), 10)
}

type MonitorWithRelativeId struct {
	Id int
}

func (m MonitorWithRelativeId) getMonitor() string {
	return fmt.Sprintf("%+d", m.Id)
}

type MonitorWithName struct {
	Name string
}

func (m MonitorWithName) getMonitor() string {
	return m.Name
}

type MonitorCurrent struct {
}

func (m MonitorCurrent) getMonitor() string {
	return "current"
}

type ResizeParams struct {
	Height, Width int
	Exact         bool
}

func (r ResizeParams) getResizeParams() (res string, err error) {
	if r.Height <= 0 || r.Width <= 0 && r.Exact {
		err = errors.New("\"Height\" and \"Width\" should be natural numbers when \"Exact\" is true")
	} else {
		res = fmt.Sprintf("%d %d", r.Height, r.Width)
	}
	if r.Exact {
		res = "exact " + res
	}
	return
}

type FloatValue struct {
	Value float64
	Exact bool
}

func (f FloatValue) parse() (res string) {
	if f.Exact {
		res = fmt.Sprintf("exact %f", f.Value)
	} else {
		res = fmt.Sprintf("%+f", f.Value)
	}
	return
}

type WorkspaceOptions string

const (
	WSOptionsAllFloat  WorkspaceOptions = "allfloat"
	WSOptionsAllPseudo WorkspaceOptions = "allpseudo"
)

type DPMSMode string

const (
	DPMSModeOn     DPMSMode = "on"
	DPMSModeOff    DPMSMode = "off"
	DPMSModeToggle DPMSMode = "toggle"
)

type Fullscreenmode string

const (
	// Fullscreen
	FullscreenmodeF Fullscreenmode = "0"
	// Maxiumize
	FullscreenmodeM Fullscreenmode = "1"
)

type MovewindowParams interface {
	getMovewindowParams() string
}

type MovewindowParamsWithDirection struct {
	Direction
}

func (m MovewindowParamsWithDirection) getMovewindowParams() string {
	return string(m.Direction)
}

type MovewindowParamsWithMonitor struct {
	Monitor
}

func (m MovewindowParamsWithMonitor) getMovewindowParams() string {
	return fmt.Sprintf("mon:%s", m.Monitor.getMonitor())
}

type CycleMode string

const (
	CycleModeNext CycleMode = ""
	CycleModePrev CycleMode = "prev"
)

type Corner string

const (
	CornerBottomLeft  Corner = "0"
	CornerBottomRight Corner = "1"
	CornerTopLeft     Corner = "2"
	CornerTopRight    Corner = "3"
)

type GroupCycleMode string

const (
	GroupCycleModeBack    GroupCycleMode = "b"
	GroupCycleModeForward GroupCycleMode = "f"
)

type GroupLockMode string

const (
	GroupLockModeLock   GroupLockMode = "lock"
	GroupLockModeUnlock GroupLockMode = "unlock"
	GroupLockModeToggle GroupLockMode = "toggle"
)

/*  --------------------PARAM TYPE DEFINITION END--------------------  */

/*  --------------------DISPATCH DEFINITION START--------------------  */

func (c *Client) DispatcherExec(args []string) error {
	cmd := c.newCmd("dispatcher", "", append([]string{"exec"}, args...))
	err := cmd.err()
	return err
}

func (c *Client) DispatcherExecR(args []string) error {
	cmd := c.newCmd("dispatcher", "", append([]string{"execr"}, args...))
	err := cmd.err()
	return err
}

func (c *Client) DispatcherPass(window Window) error {
	cmd := c.newCmd("dispatcher", "", []string{"pass", window.getWindow()})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherKillActive() error {
	cmd := c.newCmd("dispatcher", "", []string{"killactive"})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherCloseWindow(window Window) error {
	cmd := c.newCmd("dispatcher", "", []string{"closewindow", window.getWindow()})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherWorkspace(workspace Workspace) error {
	cmd := c.newCmd("dispatcher", "", []string{"workspace", workspace.getWorkspace()})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherMoveToWorkspace(workspace Workspace, window Window, silent bool) error {
	var command string
	if !silent {
		command = "movetoworkspace"
	} else {
		command = "movetoworkspacesilent"
	}
	var arg string
	if window.getWindow() == "" {
		arg = workspace.getWorkspace()
	} else {
		arg = workspace.getWorkspace() + "," + window.getWindow()
	}
	cmd := c.newCmd("dispatcher", "", []string{command, arg})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherToggleFloating(window Window) error {
	cmd := c.newCmd("dispatcher", "", []string{"togglefloating", window.getWindow()})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherFullscreen(mode Fullscreenmode) error {
	cmd := c.newCmd("dispatcher", "", []string{"fullscreen", string(mode)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherFakeFullscreen() error {
	cmd := c.newCmd("dispatcher", "", []string{"fakefullscreen"})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherDPMS(mode DPMSMode, monitor Monitor) error {
	cmd := c.newCmd("dispatcher", "", []string{"dpms", string(mode), monitor.getMonitor()})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherPin(window Window) error {
	cmd := c.newCmd("dispatcher", "", []string{"pin", window.getWindow()})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherMoveFocus(direction Direction) error {
	cmd := c.newCmd("dispatcher", "", []string{"movefocus", string(direction)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherMoveWindow(param MovewindowParams) error {
	cmd := c.newCmd("dispatcher", "", []string{"movewindow", param.getMovewindowParams()})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherSwapWindow(direction Direction) error {
	cmd := c.newCmd("dispatcher", "", []string{"swapwindow", string(direction)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherCenterWindow() error {
	cmd := c.newCmd("dispatcher", "", []string{"centerwindow"})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherResizeActive(param ResizeParams) error {
	paramstr, err := param.getResizeParams()
	if err != nil {
		return err
	}
	cmd := c.newCmd("dispatcher", "", []string{"resizeactive", paramstr})
	err = cmd.err()
	return err
}

func (c *Client) DispatcherMoveActive(param ResizeParams) error {
	paramstr, err := param.getResizeParams()
	if err != nil {
		return err
	}
	cmd := c.newCmd("dispatcher", "", []string{"moveactive", paramstr})
	err = cmd.err()
	return err
}

func (c *Client) DispatcherResizeWindowPixel(param ResizeParams, window Window) error {
	paramstr, err := param.getResizeParams()
	if err != nil {
		return err
	}
	cmd := c.newCmd("dispatcher", "", []string{"resizewindowpixel", fmt.Sprintf("%s,%s", paramstr, window.getWindow())})
	err = cmd.err()
	return err
}

func (c *Client) DispatcherMoveWindowPixel(param ResizeParams, window Window) error {
	paramstr, err := param.getResizeParams()
	if err != nil {
		return err
	}
	cmd := c.newCmd("dispatcher", "", []string{"movewindowpixel", fmt.Sprintf("%s,%s", paramstr, window.getWindow())})
	err = cmd.err()
	return err
}

func (c *Client) DispatcherCycleNext(mode CycleMode) error {
	cmd := c.newCmd("dispatcher", "", []string{"cyclenext", string(mode)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherSwapNext(mode CycleMode) error {
	cmd := c.newCmd("dispatcher", "", []string{"swapnext", string(mode)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherSplitRatio(value FloatValue) error {
	cmd := c.newCmd("dispatcher", "", []string{"splitratio", value.parse()})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherToggleOpaque() error {
	cmd := c.newCmd("dispatcher", "", []string{"toggleopaque"})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherMoveCursorToCorner(corner Corner) error {
	cmd := c.newCmd("dispatcher", "", []string{"movecursortocorner", string(corner)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherMoveCursor(x, y int) error {
	cmd := c.newCmd("dispatcher", "", []string{"movecursor", fmt.Sprintf("%d,%d", x, y)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherWorkspaceOpt(opt WorkspaceOptions) error {
	cmd := c.newCmd("dispatcher", "", []string{"workspaceopt", string(opt)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherRenameWorkspace(id uint64, name string) error {
	cmd := c.newCmd("dispatcher", "", []string{"renameworkspace", strconv.FormatUint(id, 10), name})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherExit() error {
	cmd := c.newCmd("dispatcher", "", []string{"exit"})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherForceRendererReload() error {
	cmd := c.newCmd("dispatcher", "", []string{"forcerendererreload"})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherMoveCurrentWorkspaceToMonitor(monitor Monitor) error {
	cmd := c.newCmd("dispatcher", "", []string{"movecurrentworkspacetomonitor", monitor.getMonitor()})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherMoveWorkspaceToMonitor(workspace Workspace, monitor Monitor) error {
	cmd := c.newCmd("dispatcher", "", []string{"moveworkspacetomonitor", workspace.getWorkspace(), monitor.getMonitor()})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherSwapActiveWorkspaces(monitor1, monitor2 Monitor) error {
	cmd := c.newCmd("dispatcher", "", []string{"swapactiveworkspaces", monitor1.getMonitor(), monitor2.getMonitor()})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherBringActiveToTop() error {
	cmd := c.newCmd("dispatcher", "", []string{"bringactivetotop"})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherToggleSpecialWorkspace(name string) error {
	cmd := c.newCmd("dispatcher", "", []string{"togglespecialworkspace", name})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherFocusUrgentOrLast() error {
	cmd := c.newCmd("dispatcher", "", []string{"focusurgentorlast"})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherToggleGroup() error {
	cmd := c.newCmd("dispatcher", "", []string{"togglegroup"})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherChangeGroupActive(mode GroupCycleMode) error {
	cmd := c.newCmd("dispatcher", "", []string{"changegroupactive", string(mode)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherFocusCurrentOrLast() error {
	cmd := c.newCmd("dispatcher", "", []string{"focuscurrentorlast"})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherLockGroups(mode GroupLockMode) error {
	cmd := c.newCmd("dispatcher", "", []string{"lockgroups", string(mode)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherLockActiveGroup(mode GroupLockMode) error {
	cmd := c.newCmd("dispatcher", "", []string{"lockactivegroup", string(mode)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherMoveIntoGroup(direction Direction) error {
	cmd := c.newCmd("dispatcher", "", []string{"moveintogroup", string(direction)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherMoveOutOfGroup() error {
	cmd := c.newCmd("dispatcher", "", []string{"moveoutofgroup"})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherMoveGroupWindow(mode GroupCycleMode) error {
	cmd := c.newCmd("dispatcher", "", []string{"movegroupwindow", string(mode)})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherGlobal(name string) error {
	cmd := c.newCmd("dispatcher", "", []string{"global", name})
	err := cmd.err()
	return err
}

func (c *Client) DispatcherSubmap(name string) error {
	cmd := c.newCmd("dispatcher", "", []string{"submap", name})
	err := cmd.err()
	return err
}

/*  --------------------DISPATCH DEFINITION END--------------------  */
