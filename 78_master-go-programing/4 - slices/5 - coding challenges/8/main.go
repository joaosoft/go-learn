/*
Consider the following slice declaration:

 years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}

Using a slice expression and append() function
create a new slice called newYears that contains the first 3 and the last 3 elements of the slice.
newYears should be []int{2000, 2001, 2002, 2008, 2009, 2010}
*/

package main

import "fmt"

func main() {

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}

	newYears = append(years[:3], years[len(years)-3:]...)

	fmt.Printf("%#v\n", newYears)
}
