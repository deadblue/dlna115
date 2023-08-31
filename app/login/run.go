package login

import (
	"bytes"
	"fmt"
	_ "image/png"
	"log"
	"os"
	"strings"

	"github.com/deadblue/elevengo"
	"github.com/deadblue/qrascii"
)

func (c *Command) Run() (err error) {
	agent := elevengo.Default()
	session := &elevengo.QrcodeSession{}
	if err = agent.QrcodeStart(session, c.Platform); err != nil {
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
			err = agent.QrcodeStart(session, c.Platform)
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

	var file *os.File
	if c.SaveFile != "" {
		file, _ = os.OpenFile(c.SaveFile, os.O_WRONLY|os.O_CREATE, 0644)
	}
	useStdout := file == nil
	if useStdout {
		file = os.Stdout
		println("Please save follow cookies to a file.\n")
	}
	defer file.Close()

	fmt.Fprintf(file, "UID=%s\n", cred.UID)
	fmt.Fprintf(file, "CID=%s\n", cred.CID)
	fmt.Fprintf(file, "SEID=%s\n", cred.SEID)

	if useStdout {
		println()
	} else {
		fmt.Printf("Cookies saved at: %s\n\n", c.SaveFile)
	}
	return
}
