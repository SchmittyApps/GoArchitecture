package main

import (
	"fmt"
	"github.com/SchmittyApps/GoArchitecture/dirOne"
	"github.com/SchmittyApps/GoArchitecture/dirOne/dirThree"
	"github.com/SchmittyApps/GoArchitecture/dirTwo"
)

func main() {
	fmt.Println("Hello, World!")
	dirThree.SayHello()
	dirOne.SayGoodbye()
	dirTwo.SayShowdown()
}
