package utils

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	// TitleStyle for section titles
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("63")). // Bright blue
			MarginBottom(1)

	TaskStyle = lipgloss.NewStyle().MarginLeft(2)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("78")). // Green
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("197")). // Red
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("214")). // Orange
			Bold(true)

	InfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("246")) // Light gray

	CommandPromptStyle = lipgloss.NewStyle().
				MarginTop(1).
				MarginLeft(2).
				Foreground(lipgloss.Color("208")) // Orange-yellow

	CommandStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("39")). // A blue color
			Padding(0, 1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")) // Light gray border

	KeyStyle = lipgloss.NewStyle().Bold(true)
	// Symbols
	CheckMark   = SuccessStyle.Render("✓")
	CrossMark   = ErrorStyle.Render("✗")
	WarnMark    = WarningStyle.Render("⚠")
	InfoMark    = lipgloss.NewStyle().Foreground(lipgloss.Color("69")).Render("ℹ")  // Blue info icon
	BulletPoint = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("•") // Bullet point
	Arrow       = lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Render("➤")  // Arrow symbol
)

// --- Public Display Functions ---
// ShowTitle prints a section title.
// e.g., 🚀 Preparing Frontend Environment
func ShowTitle(title string) {
	fmt.Println(TitleStyle.Render("🚀 " + title))
}

// ShowStep prints a new step or task, typically with a leading symbol.
// e.g., • Downloading release scan-eth.zip...
func ShowStep(message string) {
	fmt.Println(TaskStyle.Render(fmt.Sprintf("%s %s", BulletPoint, message)))
}

// ShowSuccess marks a step or task as successful.
// e.g., ✓ Downloaded scan-eth.zip successfully (7.12 MB)
func ShowSuccess(message string) {
	fmt.Println(TaskStyle.Render(fmt.Sprintf("%s %s", CheckMark, SuccessStyle.Render(message))))
}

// ShowFailure marks a step or task as failed.
// e.g., ✗ Docker image build failed. See logs for details.
func ShowFailure(message string) {
	fmt.Println(TaskStyle.Render(fmt.Sprintf("%s %s", CrossMark, ErrorStyle.Render(message))))
}

// ShowWarning displays a warning message.
// e.g., ⚠ Node.js version 12.x is deprecated. Please upgrade.
func ShowWarning(message string) {
	fmt.Println(TaskStyle.Render(fmt.Sprintf("%s %s", WarnMark, WarningStyle.Render(message))))
}

// ShowInfo displays a general informational message or supplementary note.
// e.g., ℹ Frontend prepared in frontend/scan-eth
func ShowInfo(message string) {
	fmt.Println(TaskStyle.Render(fmt.Sprintf("%s %s", InfoMark, InfoStyle.Render(message))))
}

// ShowCommandSuggestion suggests a command for the user to execute next.
// e.g., 👉 To start the Docker container, run:
func ShowCommandSuggestion(message string) {
	fmt.Println(CommandPromptStyle.Render(fmt.Sprintf("👉 %s", message)))
}

// ShowCommand prints a command for the user to execute, with a border.
// e.g., cd ./frontend && docker compose up -d --build
func ShowCommand(command string) {
	fmt.Println(TaskStyle.MarginLeft(4).Render(CommandStyle.Render(command)))
}

// ShowSeparator prints a separator line for visual distinction.
func ShowSeparator() {
	fmt.Println(InfoStyle.Render(strings.Repeat("─", 30))) // Uses InfoStyle color
}

// ShowFinalMessage prints the final summary message (success or failure).
func ShowFinalMessage(isSuccess bool, message string) {
	fmt.Println() // Add a blank line for separation
	if isSuccess {
		// SuccessStyle is already bold
		fmt.Println(SuccessStyle.Render(fmt.Sprintf("🎉 %s", message)))
	} else {
		// ErrorStyle is already bold
		fmt.Println(ErrorStyle.Render(fmt.Sprintf("💥 %s", message)))
	}
}
