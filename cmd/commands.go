package cmd

import (
	"bufio"
	"fmt"
	"github.com/tano/chess-tournament/core"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	NewTournament  string = "Create new tournament"
	Exit                  = "Exit"
	AddParticipant        = "Add new participant"
	ToMainMenu            = "Return to main menu"
)

var mainMenu = []string{
	Exit,
	NewTournament,
}

var newTournamentMenu = []string{
	ToMainMenu,
	AddParticipant,
}

var addParticipantMenu = []string{
	ToMainMenu,
}

func PrintGreeting() {
	clearScreen()
	fmt.Println("Hello, this program can help you to organize the chess tournament.")
}

func PrintFarewell() {
	fmt.Println("Thanks, buy")
}

func InteractWithUser() {
	selectMenuItem()
}

func printMenu(menu []string) {
	clearScreen()
	fmt.Println("Please choose one of the options:")
	for index, item := range menu {
		fmt.Printf("Press %d to %s\n", index, item)
	}
	fmt.Println("")
}

func selectMenuItem() {
	var done = false
	menuContent := mainMenu
	for done == false {
		printMenu(menuContent)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		itemIndex, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			continue
		}
		switch selected := menuContent[itemIndex]; selected {
		case Exit:
			done = true
		case NewTournament:
			menuContent = newTournamentMenu
		case ToMainMenu:
			menuContent = mainMenu
		case AddParticipant:
			addParticipant()
		default:
			fmt.Println("Unknown command")
		}
	}
}

func addParticipant() {
	var done = false
	for done == false {
		var participants = core.GetParticipants()
		fmt.Printf("Current participants: %v, add new participant (Y/N)?", participants)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		answer := strings.TrimSpace(text)
		if answer == "Y" {
			fmt.Printf("Enter participant name...")
			text, _ := reader.ReadString('\n')
			participant := strings.TrimSpace(text)
			core.AddParticipant(participant)
		} else {
			done = true
		}
	}
}

func clearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
