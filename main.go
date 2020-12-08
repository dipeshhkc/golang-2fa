package main

import (
	"bufio"
	"bytes"
	"crypto"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/sec51/twofactor"
)

var (
	email  = "dipesh.kc@wesionary.team"
	issuer = "wesionaryTEAM"
)

func getUserInput() string {
	//Get User Inputs(token)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)

}

func main() {
	//Instantiate totp
	otp, err := twofactor.NewTOTP(email, issuer, crypto.SHA1, 6)
	if err != nil {
		log.Print("Error during OTP generation: ", err)
	}
	// Generate QR bytes
	qrBytes, err := otp.QR()
	if err != nil {
		log.Print("Error during QR generation: ", err)
	}
	//Generate PNG from QR bytes
	img, _, _ := image.Decode(bytes.NewReader(qrBytes))
	out, err := os.Create("./QRImg.png")
	err = png.Encode(out, img)

	for {
		log.Print("Enter OTP Token")
		text := getUserInput()
		//Validate the token
		if err = otp.Validate(text); err != nil {
			log.Print("Error during Validaion: ", err)

		} else {

			log.Print("Validation Success")
			break
		}
	}

}
