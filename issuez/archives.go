package issuez

/*
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/lib/darwin_arm64
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/lib/darwin_amd64
#cgo linux,amd64,!musl LDFLAGS: -L${SRCDIR}/lib/linux_amd64
#cgo linux,amd64,musl LDFLAGS: -L${SRCDIR}/lib/linux_amd64_musl
#cgo linux,arm64,!musl LDFLAGS: -L${SRCDIR}/lib/linux_arm64
#cgo linux,arm64,musl LDFLAGS: -L${SRCDIR}/lib/linux_arm64_musl
#cgo linux,arm,!musl LDFLAGS: -L${SRCDIR}/lib/linux_arm
#cgo linux,arm,musl LDFLAGS: -L${SRCDIR}/lib/linux_arm_musl
*/
import "C"
