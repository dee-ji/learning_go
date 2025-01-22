package main

/*
#include <string.h>

// C function
int mini_calc(char *op, int a, int b) {
    if (strcmp(op, "+") == 0) {
        return a + b;
    }
    if (strcmp(op, "*") == 0) {
        return a * b;
    }
    if (strcmp(op, "-") == 0) {
        return a - b;
    }
    if (strcmp(op, "/") == 0) {
        if (b == 0) {
            return 0; // handling divide by zero
        }
        return a / b;
    }
    return 0;
}
*/
import "C"
import (
	"fmt"
)

func main() {
	// Convert Go string to a C string
	op := C.CString("+")
	//defer C.free(unsafe.Pointer(op)) // Make sure to free the C string

	// Call the C function
	result := C.mini_calc(op, 3, 4)
	fmt.Println("3 + 4 =", result)
}
