package internal

import (
	"encoding/json"
	"fmt"
)

// KomorebiEvent represents the entire event structure from komorebi
type KomorebiEvent struct {
	Event EventInfo     `json:"event"`
	State KomorebiState `json:"state"`
}

// EventInfo represents the event part of the komorebi event
type EventInfo struct {
	Type    string          `json:"type"`
	Content json.RawMessage `json:"content"`
}

// KomorebiState represents the state of the window manager
type KomorebiState struct {
	Monitors    Monitors       `json:"monitors"`
	MonitorMap  map[string]int `json:"monitor_usr_idx_map"`
	ResizeDelta int            `json:"resize_delta"`
	// We don't need all fields from the state, just focusing on what we care about
}

// Monitors contains all monitor elements
type Monitors struct {
	Elements []Monitor `json:"elements"`
	Focused  int       `json:"focused"`
}

// Monitor represents a physical monitor
type Monitor struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Workspaces Workspaces `json:"workspaces"`
}

// Workspaces contains all workspace elements
type Workspaces struct {
	Elements []Workspace `json:"elements"`
	Focused  int         `json:"focused"`
}

// Workspace represents a workspace on a monitor
type Workspace struct {
	Name       string     `json:"name"`
	Containers Containers `json:"containers"`
}

// Containers contains all container elements
type Containers struct {
	Elements []Container `json:"elements"`
	Focused  int         `json:"focused"`
}

// Container represents a container within a workspace
type Container struct {
	ID      string  `json:"id"`
	Windows Windows `json:"windows"`
}

// Windows contains all window elements
type Windows struct {
	Elements []Window `json:"elements"`
	Focused  int      `json:"focused"`
}

// Window represents a window within a container
type Window struct {
	Hwnd  int    `json:"hwnd"`
	Title string `json:"title"`
	Exe   string `json:"exe"`
	Class string `json:"class"`
}

// PertinentState is our simplified representation of the state
type PertinentState struct {
	MonitorWorkspaces map[string][]WorkspaceInfo
	ActiveMonitor     string
	ActiveWorkspace   string
}

// WorkspaceInfo contains the essential workspace information we care about
type WorkspaceInfo struct {
	Name           string
	ContainerCount int
	WindowCount    int
	ContainerInfos []ContainerInfo
}

// ContainerInfo contains essential container information
type ContainerInfo struct {
	ID          string
	WindowCount int
	Windows     []string // Window titles
}

// StateTracker maintains the current state
type StateTracker struct {
	CurrentState   KomorebiState
	PertinentState PertinentState
}

// NewStateTracker creates a new state tracker
func NewStateTracker() *StateTracker {
	return &StateTracker{
		PertinentState: PertinentState{
			MonitorWorkspaces: make(map[string][]WorkspaceInfo),
		},
	}
}

// ProcessEvent processes a raw JSON event and updates the state
func (st *StateTracker) ProcessEvent(eventJSON string) error {
	var event KomorebiEvent
	err := json.Unmarshal([]byte(eventJSON), &event)
	if err != nil {
		return fmt.Errorf("error unmarshaling event: %v", err)
	}

	// Update our state
	st.CurrentState = event.State

	// Update the pertinent state
	st.UpdatePertinentState()

	return nil
}

// UpdatePertinentState updates our simplified view of the state
func (st *StateTracker) UpdatePertinentState() {
	pertinent := PertinentState{
		MonitorWorkspaces: make(map[string][]WorkspaceInfo),
	}

	// Find the focused monitor and workspace
	if len(st.CurrentState.Monitors.Elements) > 0 {
		focusedMonitorIdx := st.CurrentState.Monitors.Focused
		if focusedMonitorIdx < len(st.CurrentState.Monitors.Elements) {
			focusedMonitor := st.CurrentState.Monitors.Elements[focusedMonitorIdx]
			pertinent.ActiveMonitor = focusedMonitor.Name

			if len(focusedMonitor.Workspaces.Elements) > 0 {
				focusedWorkspaceIdx := focusedMonitor.Workspaces.Focused
				if focusedWorkspaceIdx < len(focusedMonitor.Workspaces.Elements) {
					pertinent.ActiveWorkspace = focusedMonitor.Workspaces.Elements[focusedWorkspaceIdx].Name
				}
			}
		}
	}

	// Process each monitor
	for _, monitor := range st.CurrentState.Monitors.Elements {
		var workspaceInfos []WorkspaceInfo

		for _, workspace := range monitor.Workspaces.Elements {
			workspaceInfo := WorkspaceInfo{
				Name:           workspace.Name,
				ContainerCount: len(workspace.Containers.Elements),
				ContainerInfos: make([]ContainerInfo, 0),
			}

			totalWindows := 0

			// Process each container in the workspace
			for _, container := range workspace.Containers.Elements {
				containerInfo := ContainerInfo{
					ID:          container.ID,
					WindowCount: len(container.Windows.Elements),
					Windows:     make([]string, 0),
				}

				totalWindows += containerInfo.WindowCount

				// Get window titles
				for _, window := range container.Windows.Elements {
					containerInfo.Windows = append(containerInfo.Windows, window.Title)
				}

				workspaceInfo.ContainerInfos = append(workspaceInfo.ContainerInfos, containerInfo)
			}

			workspaceInfo.WindowCount = totalWindows
			workspaceInfos = append(workspaceInfos, workspaceInfo)
		}

		pertinent.MonitorWorkspaces[monitor.Name] = workspaceInfos
	}

	st.PertinentState = pertinent
}

// PrintSummary prints a summary of the current state
func (st *StateTracker) PrintSummary() {
	fmt.Println("\n=== Komorebi State Summary ===")
	fmt.Printf("Active Monitor: %s, Active Workspace: %s\n",
		st.PertinentState.ActiveMonitor,
		st.PertinentState.ActiveWorkspace)

	for monitorName, workspaces := range st.PertinentState.MonitorWorkspaces {
		fmt.Printf("\nMonitor: %s\n", monitorName)

		for _, workspace := range workspaces {
			fmt.Printf("  Workspace: %s (Containers: %d, Windows: %d)\n",
				workspace.Name,
				workspace.ContainerCount,
				workspace.WindowCount)

			for i, container := range workspace.ContainerInfos {
				fmt.Printf("    Container %d: %d windows\n", i+1, container.WindowCount)

				for j, window := range container.Windows {
					fmt.Printf("      Window %d: %s\n", j+1, window)
				}
			}
		}
	}

	fmt.Println("===============================")
}
