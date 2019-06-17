package decorator

import (
	"log"
)

//Unlike Adapter pattern, the object to be decorated is obtained by injection.
//Decorators should not alter the interface of an object.

type Object func(int) int

func LogDecorator(fn Object) Object  {
	return func(n int) int {
		log.Println("before the truth")
		result := fn(n)
		log.Println("after the truth")
		return result
	}
}
