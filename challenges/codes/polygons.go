package main

import (
	"fmt"
  "math"
)

// POINT
type Point struct {
	X, Y float64
}

// LINE
type Line struct {
	P1, P2 Point
}

// MAIN
func main() {

  // POINTS & LINES
  var points []Point
  var lines []Line

  points = []Point{
    Point{-3, 1},
    Point{2, 3},
    Point{0, 0},
    Point{-1.5, -2.5},
    // Point{-3, 1},
    // Point{2, 3},
    // Point{-3, -4},
    // Point{2, -3},
    // Point{-1.5, -2.5},
  }
  fmt.Println(points)

  collisions := areThereCollisions(points, lines)
  perimeter := getPerimeter(points)
  area := getArea(points)

  fmt.Printf("Collisions: %v \n", collisions)
  fmt.Printf("Perimeter: %v \n", perimeter)
  fmt.Printf("Area: %v \n", area)

}

// ON SEGMENT
// check if point lies on segment
func onSegment(L1 Line, P Point) bool {

  if (P.X <= math.Max(L1.P1.X, L1.P2.X) && P.X >= math.Min(L1.P1.X, L1.P2.X) &&
      P.Y <= math.Max(L1.P1.Y, L1.P2.Y) && P.Y >= math.Min(L1.P1.Y, L1.P2.Y)) {
     return true

  } else { return false }

}

// ORIENTATION
// find orientations of ordered triplets
func getOrientation(A Point, B Point, C Point) int {
  val := (B.Y-A.Y)*(C.X-B.X)-(B.X-A.X)*(C.Y-B.Y)

  // Colinear
  if val == 0 { return 0 } else

  // Anti-Clockwise
  if val < 0 { return 2 } else

  // Clockwise
  { return 1 }

}

// INTERSECTIONS
// check if two line segments intersects
func intersects(L1 Line, L2 Line) bool {
  o1 := getOrientation(L1.P1, L1.P2, L2.P1)
  o2 := getOrientation(L1.P1, L1.P2, L2.P2)
  o3 := getOrientation(L2.P1, L2.P2, L1.P1)
  o4 := getOrientation(L2.P1, L2.P2, L1.P2)

  // General case
  if o1 != o2 && o3 != o4 { return true }

  // Special cases
  if (o1 == 0 && onSegment(L1, L2.P1)) { return true } else
  if (o2 == 0 && onSegment(L1, L2.P2)) { return true } else
  if (o3 == 0 && onSegment(L2, L1.P1)) { return true } else
  if (o4 == 0 && onSegment(L2, L1.P2)) { return true } else
  { return false }

}

// COLLISIONS
// checks whether there are any collisions
func areThereCollisions(points []Point, lines []Line) bool {

  collisions := false

  if (len(points)>0) {

    // CREATE LINE ARRAY
    for key, val := range points[:len(points)-1] {
      A := val
      B := points[key+1]
      L := Line{P1: A, P2: B}
      lines = append(lines, L)
    }

    A := points[len(points)-1]
    B := points[0]
    L := Line{P1: A, P2: B}
    lines = append(lines, L)

    // CHECK COLLISIONS
    for keyO, valO := range lines[:len(lines)] {
      for _, valI := range lines[keyO+1:len(lines)] {

        if valO.P2 != valI.P1 && valO.P1 != valI.P2 {
          collisions = intersects(valO, valI)
          //fmt.Printf("%v && ", valO)
          //fmt.Printf("%v == %v \n", valI, collisions)
          if collisions == true { break }
        }

      }
      if collisions == true { break }
    }


  }

  return collisions
}

// PERIMETER
// calculates perimeter of a given array of connected points
func getPerimeter(points []Point) float64 {

  perimeter := 0.0

  if (len(points)>0) {

    for key, val := range points[:len(points)-1] {
      A := val
      B := points[key+1]
      distance := math.Sqrt(math.Pow(B.X-A.X, 2) + math.Pow(B.Y-A.Y, 2))
      perimeter += distance
    }

    A := points[len(points)-1]
    B := points[0]
    distance := math.Sqrt(math.Pow(B.X-A.X, 2) + math.Pow(B.Y-A.Y, 2))
    perimeter += distance

  }

  return perimeter
}

// AREA
// calculates the area of a given shape
func getArea(points []Point) float64 {

  area := 0.0

  if (len(points)>0) {

    left, right := 0.0, 0.0

    for key, val := range points[:len(points)-1] {
      A := val
      B := points[key+1]
      left += A.X*B.Y
      right += A.Y*B.X
    }

    A := points[len(points)-1]
    B := points[0]
    left += A.X*B.Y
    right += A.Y*B.X
    area = math.Abs(left-right) / 2

  }

  return area
}
