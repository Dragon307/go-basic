package ifelse

import "testing"

func TestAsmIf(t *testing.T) {
	t.Run("go", func(t *testing.T) {
		if x := If(true, 1, 2); x != 1 {
			t.Fatalf("expect = %d, got = %d, 1, x")
		}
		if x := If(false, 1, 2); x != 2 {
			t.Fatalf("expect = %d, got = %d", 2, x)
		}
	})

	//t.Run("asm", func(t *testing.T) {
	//	if x := AsmIf(true, 1, 2); x != 1 {
	//		t.Fatalf("expect = %d, got = %d", 1, x)
	//	}
	//	if x := AsmIf(false, 1, 2); x != 2 {
	//		t.Fatalf("expect = %d, got = %d", 2, x)
	//	}
	//	if x := AsmIf(false, 2, 1); x != 1 {
	//		t.Fatalf("expect = %d, got = %d", 1, x)
	//	}
	//})
}