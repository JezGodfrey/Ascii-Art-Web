# ascii-art-web

## Description

This project is a web server written in Go that allows users to input text and output the text into ascii-art.
The instructions for how to do so are included on the web page.

## Usage: how to run

Simply run the program like so to start the server:

```sh
go run .
```

Then visit `localhost:8080` on your device. On the main page, type your text into the text area. The radio buttons will change the style of the ascii-art (Standard by default). Below are screenshots from the index and ascii-art output HTML pages.


![IndexPage](/IndexPage.PNG?raw=true "Index Page")
![AsciiPage](/AsciiPage.PNG?raw=true "Ascii Page")

## Implementation

- The site generates ascii-art by making use of the AsciiArt function - the input of which is passed in via the form on the web page
- This function finds the art within txt files, one for each style - standard, shadow, thinkertoy
- The function then takes the ascii value of each character of the input and subtracts and multiplies the numerical value to find the corresponding ascii art within the txt file
- It then displays the results on a new page using the POST method
- The website itself handles 400, 404, and 500 status errors by redirecting to respective pages on the server

## Authors

This program was written by Jez Godfrey.
