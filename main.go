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

func main() {
	//Instantiate totp
	otp, err := twofactor.NewTOTP("kcdipesh429@gmail.com", "dipesh", crypto.SHA1, 6)
	if err != nil {
		log.Print("Error during OTP generation", err)
	}
	// Generate QR bytes
	qrBytes, err := otp.QR()
	if err != nil {
		log.Print("Error during QR generation", err)
	}
	//Generate PNG from QR bytes
	img, _, _ := image.Decode(bytes.NewReader(qrBytes))
	out, err := os.Create("./QRImg.png")
	err = png.Encode(out, img)

	//Get User Inputs(token)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	//Validate the token
	if err = otp.Validate(text); err != nil {
		log.Print("Error during Validaion", err)
	} else {

		log.Print("Validation Success")
	}

}
