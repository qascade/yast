/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/

//This file will contain all the choices a player can have while interacting with tui
package tui

var (
	playerChoice       string
	targetChoice       string
	chosenResultMagnet string
)

func SetPlayerChoice(choice string) {
	playerChoice = choice
}

func GetPlayerChoice() string {
	return playerChoice
}

func SetMagnetChoice(magnet string) {
	chosenResultMagnet = magnet
}

func GetMagnetFromListModel() string {
	return chosenResultMagnet
}
