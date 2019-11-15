// testutil project testutil.go
package testutil

type TestFunc interface {
	Test([]string)
}
