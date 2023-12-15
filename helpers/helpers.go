package helpers

import "fmt"

func LogError(err error) {
	if err != nil {
		fmt.Println("\n\n Error thrown is >>> \t\t", err.Error(), "\n\n ")
		panic(err)
	}
}

/** pl: for fmt.Println */
func Pl(msg ...any) {
	fmt.Println("\n", msg, "\n ")
}
