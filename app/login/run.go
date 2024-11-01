package login

import (
	"bytes"
	"fmt"
	_ "image/png"
	"log"
	"os"
	"strings"

	"github.com/deadblue/dlna115/pkg/credential"
	"github.com/deadblue/elevengo"
	"github.com/deadblue/qrascii"
)

func (c *Command) Run() (err error) {
	agent := elevengo.Default()
	session := &elevengo.QrcodeSession{}
	if err = agent.QrcodeStart(session); err != nil {
		log.Fatalf("Start login session failed: %s", err)
	}

	println("Please scan QRCode on your 115 App:\n")
	for done := false; !done && err == nil; {
		// Ignore error here
		qr, _ := qrascii.Parse(bytes.NewReader(session.Image))
		qrAscii := qr.String()
		println(qrAscii)

		// Poll QRCode status
		for !done && err == nil {
			done, err = agent.QrcodePoll(session)
		}

		// QRCode expired or canceled
		if err != nil {
			// Request new QRCode
			err = agent.QrcodeStart(session)
			if err == nil {
				// Replace old QRCode
				lineCount := strings.Count(qrAscii, "\n") + 1
				fmt.Printf("\033[%dA", lineCount)
			}
		}
	}

	// Login successed, export cookie
	cred := &elevengo.Credential{}
	agent.CredentialExport(cred)
	credData := credential.Encode(cred, c.secret)

	var file *os.File
	if c.saveFile != "" {
		var ferr error
		if file, ferr = os.OpenFile(c.saveFile, os.O_WRONLY|os.O_CREATE, 0644); ferr != nil {
			println("Can not open file to write: %s", c.saveFile)
		}
	}
	if file == nil {
		fmt.Printf("Please save credential to a file: \n\n%s\n\n", credData)
	} else {
		defer file.Close()
		file.WriteString(credData)
		fmt.Printf("Credential saved at: %s\n\n", c.saveFile)
	}
	return
}
