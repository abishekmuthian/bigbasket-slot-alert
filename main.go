package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gregdel/pushover"
)

func main() {

	// Get the API key and User Key from the arguments
	if len(os.Args) < 3 {
		fmt.Printf("usage: ./bigbasket-slot-alert api_key user_key\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	// Create a new pushover app with a token
	app := pushover.New(os.Args[1])

	// Create a new recipient
	recipient := pushover.NewRecipient(os.Args[2])

	// Create the message to send
	message := pushover.NewMessage("Possible BigBasket slot availability!")

	// Set the flags to open the browser in non-headless mode
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)

	// create executor
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	// create context
	ctx, cancelCtx := chromedp.NewContext(allocCtx)

	// String to store the text
	var res string

	// URL to open
	err := LoadandLogin(`https://bigbasket.com`, &res, ctx)
	if err != nil {
		res = ""
	}

	// Retrieved text from site message
	log.Println(strings.TrimSpace(res))

	// Process the string
	if strings.Contains(res, "slots may not be available currently") {
		log.Println("\n Slot not available... reloading in 10 minutes")
		quit := make(chan bool)
		for range time.NewTicker(10 * time.Minute).C { // 10 Minutes interval for reloading the page
			err := ReloadandCheck(&res, ctx)
			if err != nil {
				res = ""
			}

			if strings.Contains(res, "slots may not be available currently") {
				log.Println("\n Still no dice, reloading in 10 minutes")
			} else {
				log.Println("\n During reload site message is different, possible slot availability!")

				// Send the message to the recipient
				log.Println("Sending push message to the device")
				response, err := app.SendMessage(message, recipient)
				if err != nil {
					log.Panic(err)
				}

				// Print the response if you want
				log.Println(response)

				// Canceling the Context
				log.Println("\n Canceling the main browser Context")
				cancelCtx()

				// Closing the Ticker
				log.Println("\n Closing the Ticker")
				quit <- true
				close(quit)

				// Exiting the program
				os.Exit(1)
			}
		}
	}
}

func LoadandLogin(urlstr string, res *string, ctx context.Context) error {

	log.Println("Login, select location and return to the home page within 5 minutes")

	err := chromedp.Run(ctx,
		chromedp.Navigate(urlstr),
		chromedp.Sleep(5*time.Minute), // 5 Minutes to login and select the location and return to home page
	)

	if err != nil {
		log.Println("\n Error during site load")
		log.Fatal(err)
	} else {
		// create timeout context for not stalling if site message element was not found
		tctx, cancel := context.WithTimeout(ctx, 2*time.Minute) // 2 Minutes for timeout, for selecting text

		err = CheckSiteMessage(res, tctx)
		if err != nil {
			if tctx.Err() != nil {
				// we hit the timeout
				log.Println("\n During loading site_msg_label was not found, possible slot availability!")
			} else {
				// we hit some chromedp error
				log.Fatal(err)
			}
		}

		cancel()
	}
	return err
}

func ReloadandCheck(res *string, ctx context.Context) error {
	err := chromedp.Run(ctx,
		chromedp.Reload(),
	)

	if err != nil {
		log.Println("\n Error during site reload")
		log.Fatal(err)
	} else {
		log.Println("\n Reloading... and waiting for 2 minutes for the text if needed")

		// create timeout context for not stalling if site message element was not found
		tctx, cancel := context.WithTimeout(ctx, 2*time.Minute) // 2 Minutes for timeout, for selecting text

		err = CheckSiteMessage(res, tctx)
		if err != nil {
			if tctx.Err() != nil {
				// we hit the timeout
				log.Println("\n During reload site_msg_label was not found, possible slot availability!")
			} else {
				// we hit some chromedp error
				log.Fatal(err)
			}
		}

		cancel()
	}

	log.Println("\n Site reloaded")

	return err
}

func CheckSiteMessage(res *string, tctx context.Context) error {
	// Read the text from the site_message_label element
	err := chromedp.Run(tctx,
		chromedp.Text(`#site_msg_label`, res, chromedp.NodeVisible, chromedp.ByID),
	)

	log.Println("\n Checked site message")

	return err
}
