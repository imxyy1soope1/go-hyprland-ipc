package event

type Event uint8

const (
	WORKSPACE Event = iota
	FOCUSEDMON
	ACTIVEWINDOW
	ACTIVEWINDOWV2
	FULLSCREEN
	MONITORREMOVED
	MONITORADDED
	CREATEWORKSPACE
	DESTROYWORKSPACE
	MOVEWORKSPACE
	ACTIVELAYOUT
	OPENWINDOW
	CLOSEWINDOW
	MOVEWINDOW
	OPENLAYER
	CLOSELAYER
	SUBMAP
	CHANGEFLOATINGMODE
	URGENT
	MINIMIZE
	SCREENCAST
	WINDOWTITLE
)

var stringToEvent = map[string]Event{
	"workspace":          WORKSPACE,
	"focusedmon":         FOCUSEDMON,
	"activewindow":       ACTIVEWINDOW,
	"activewindowv2":     ACTIVEWINDOWV2,
	"fullscreen":         FULLSCREEN,
	"monitorremoved":     MONITORREMOVED,
	"monitoradded":       MONITORADDED,
	"createworkspace":    CREATEWORKSPACE,
	"destroyworkspace":   DESTROYWORKSPACE,
	"moveworkspace":      MOVEWORKSPACE,
	"activelayout":       ACTIVELAYOUT,
	"openwindow":         OPENWINDOW,
	"closewindow":        CLOSEWINDOW,
	"openlayer":          OPENLAYER,
	"closelayer":         CLOSELAYER,
	"submap":             SUBMAP,
	"changefloatingmode": CHANGEFLOATINGMODE,
	"urgent":             URGENT,
	"minimize":           MINIMIZE,
	"screencast":         SCREENCAST,
	"windowtitle":        WINDOWTITLE,
}
