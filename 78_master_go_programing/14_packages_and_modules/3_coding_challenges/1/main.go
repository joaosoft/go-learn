/*
Consider this function:

// Checks if a number is prime or not
func IsPrime(num int) bool {
    for i := 2; i < int(num/2); i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}
Create a package called arithmetic in $GOPATH/packages/mypackages and paste the above function in a file of package.

Create an executable Go application, import the package arithmetic and call the function IsPrime.
*/

package main

