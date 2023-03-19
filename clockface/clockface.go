package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCenterX     = 150
	clockCenterY     = 150

	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

type Point struct {
	X float64
	Y float64
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func hourHand(w io.Writer, t time.Time) {
	p := hourHandPoint(t)
	p = Point{p.X * hourHandLength, p.Y * hourHandLength} //Scale
	p = Point{p.X, -p.Y}                                  //flip
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}     //translate
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := minuteHandPoint(t)
	p = Point{p.X * minuteHandLength, p.Y * minuteHandLength} //Scale
	p = Point{p.X, -p.Y}                                      //flip
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         //translate
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func secondHand(w io.Writer, t time.Time) {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} //Scale
	p = Point{p.X, -p.Y}                                      //flip
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         //translate

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / float64(t.Second())))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hourInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}

func secondHandPoint(t time.Time) Point {
	radians := secondsInRadians(t)
	return angleToPoint(radians)
}

func minuteHandPoint(t time.Time) Point {
	radians := minutesInRadians(t)
	return angleToPoint(radians)
}

func hourHandPoint(t time.Time) Point {
	radians := hourInRadians(t)
	return angleToPoint(radians)
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
const svgEnd = `</svg>`
