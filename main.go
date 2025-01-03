package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// SysInfo saves the basic system information
type SysInfo struct {
	Hostname string
	Uptime   uint64
	CPU      []string
	MemUsed  uint64
	MemTotal uint64
	DiskUsed uint64
	DiskTotal uint64
}

func getSystemInfo() SysInfo {
	hostname, _ := host.Info()
	uptime := hostname.Uptime

	cpuInfos, _ := cpu.Info()
	var cpuDetails []string
	for _, cpu := range cpuInfos {
		cpuDetails = append(cpuDetails, cpu.ModelName)
	}

	memInfo, _ := mem.VirtualMemory()
	diskInfo, _ := disk.Usage("/")

	return SysInfo{
		Hostname: hostname.Hostname,
		Uptime:   uptime,
		CPU:      cpuDetails,
		MemUsed:  memInfo.Used,
		MemTotal: memInfo.Total,
		DiskUsed: diskInfo.Used,
		DiskTotal: diskInfo.Total,
	}
}

func main() {
	a := app.NewWithID("com.example.sysinfo")
	a.Settings().SetTheme(theme.DarkTheme())

	w := a.NewWindow("System Information - GhostBSD")
	w.Resize(fyne.NewSize(500, 400))

	// Widgets for displaying dynamic information
	hostnameLabel := widget.NewLabel("")
	uptimeLabel := widget.NewLabel("")
	cpuLabel := widget.NewLabel("")
	memProgress := widget.NewProgressBar()
	diskProgress := widget.NewProgressBar()

	// Function to refresh system information
	refresh := func() {
		sysInfo := getSystemInfo()

		hostnameLabel.SetText(sysInfo.Hostname)
		uptimeLabel.SetText(formatUptime(sysInfo.Uptime))
		cpuLabel.SetText(cpuDetailsToString(sysInfo.CPU))

		memProgress.Value = float64(sysInfo.MemUsed) / float64(sysInfo.MemTotal)
		memProgress.TextFormatter = func() string {
			return fmt.Sprintf("%s / %s", formatBytes(sysInfo.MemUsed), formatBytes(sysInfo.MemTotal))
		}

		diskProgress.Value = float64(sysInfo.DiskUsed) / float64(sysInfo.DiskTotal)
		diskProgress.TextFormatter = func() string {
			return fmt.Sprintf("%s / %s", formatBytes(sysInfo.DiskUsed), formatBytes(sysInfo.DiskTotal))
		}
	}

	// Title and system information layout
	title := widget.NewLabelWithStyle("System Information", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	grid := container.NewVBox(
		widget.NewLabelWithStyle("Hostname", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		hostnameLabel,
		widget.NewLabelWithStyle("Uptime", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		uptimeLabel,
		widget.NewLabelWithStyle("CPU", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		cpuLabel,
		widget.NewLabelWithStyle("RAM (Memory)", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		memProgress,
		widget.NewLabelWithStyle("Storage (Disk)", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		diskProgress,
	)

	// Close button
	closeButton := widget.NewButton("Close", func() {
		w.Close()
	})

	// Main content layout with proper spacing
	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		grid,
		container.NewCenter(closeButton), // Center the Close button
	)

	w.SetContent(content)

	// Refresh data periodically
	go func() {
		for range time.Tick(time.Second * 5) {
			refresh()
		}
	}()

	// Initial refresh and show window
	refresh()
	w.ShowAndRun()
}

func cpuDetailsToString(cpu []string) string {
	return strings.Join(cpu, ", ")
}

func formatUptime(uptime uint64) string {
	hours := uptime / 3600
	minutes := (uptime % 3600) / 60
	return fmt.Sprintf("%d hours, %d minutes", hours, minutes)
}

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := unit, 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

