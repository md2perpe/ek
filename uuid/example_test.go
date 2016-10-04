package uuid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGenUUID() {
	fmt.Printf("UUID: %s\n", GenUUID())
}

func ExampleGenUUID4() {
	fmt.Printf("UUID v4: %s\n", GenUUID4())
}

func ExampleGenUUID5() {
	fmt.Printf("UUID v5: %s\n", GenUUID5(NsURL, "http://www.domain.com"))
}