package main

import "fmt"

func main() {
  var publisher, writer, artist, title string
  var year, pageNumber uint16
  var grade float32

  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5

  fmt.Println(title, " written by ", writer, " drawn by ", artist, " and published by ", publisher, ". The year is " , year, " page: ", pageNumber, " and grade: ", grade)

  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  year = 2013
  pageNumber = 160
  grade = 9.0

  fmt.Println(title, " written by ", writer, " drawn by ", artist, " and published by ", publisher, ". The year is " , year, " page: ", pageNumber, " and grade: ", grade)
}