
package main


func main() {
	println("If(true, 1, 2) =", If(true, 1, 2))
	println("If(false, 1, 2) =", If(false, 1, 2))
	println("AsmIf(true, 1, 2) =", AsmIf(true, 1, 2))
	println("AsmIf(false, 1, 2) =", AsmIf(false, 1, 2))
	println("AsmIf(false, 2, 1) =", AsmIf(false, 2, 1))
}
