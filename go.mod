module github.com/petergtz/pegomock

go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/golang/mock v1.4.4
	github.com/nxadm/tail v1.4.5 // indirect
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.3
	github.com/petergtz/vendored_package v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/sys v0.0.0-20201018121011-98379d014ca7 // indirect
	golang.org/x/tools v0.0.0-20201017001424-6003fad69a88
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

replace github.com/petergtz/vendored_package => ./vendored_package
