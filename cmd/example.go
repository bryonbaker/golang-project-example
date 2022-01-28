package main

import (
	"fmt"

	"bakerapps.net/project-example/pkg/package1"
	"bakerapps.net/project-example/pkg/package1/subpackage1"
	"bakerapps.net/project-example/pkg/package2"
)

func main() {
	fmt.Println("Hello from main()")

	package1.Package1()
	package2.Package2()
	subpackage1.SubPackage1()

}
