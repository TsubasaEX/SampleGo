// testutil project testutil.go
package testutil

type TestFunc interface {
	Check()
	Kill()
}
