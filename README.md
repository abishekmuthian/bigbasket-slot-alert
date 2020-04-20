# bigbasket-slot-alert
Push notifications when BigBasket slots are available


# Why and How
https://abishekmuthian.com/bigbasket-slot-alert


## Requirements

1. API_KEY and USER_KEY from [Pushover](https://pushover.net/) for push notifications (7 days free trial available). Pushover Android or iOS app.
2. Google Chrome browser.

### Optional Requirements

1. [Lightflow Pro](https://play.google.com/store/apps/details?id=com.rageconsulting.android.lightflow) or other apps which can read out notifications like Tasker.


## Dependencies

1. [chromedp](https://github.com/chromedp/chromedp)
2. [Go wrapper for the Pushover API](https://github.com/gregdel/pushover


## Usage

Download the correct binary for your OS and Architecture from the [releases](https://github.com/heavyinfo/bigbasket-slot-alert/releases)

### Linux or macOS

Make the binary executable before running

`sudo chmod +x ./bigbasket-slot-alert`

Run the program

`./bigbasket-slot-alert API_KEY USER_KEY`

### Windows

Run the program

`./bigbasket-slot-alert.exe API_KEY USER_KEY`



## Ask for help

[@heavyinfo](https://twitter.com/heavyinfo)


## Fair Use

Please do not decrease the timeouts and do not disable automatic closure of the program to avoid stressing the BigBasket's servers.


## License

The MIT License


Copyright 2020 ABISHEK MUTHIAN

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.