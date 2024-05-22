package brainfuck_test

import (
	"testing"

	"github.com/Laellekoenig/brainf/internal/brainfuck"
)

func testProgram(code, expected string, t *testing.T) {
	program := brainfuck.NewProgram([]byte(code))
	out := string(program.Run())
	if out != expected {
		t.Fatalf("expected output to be %s, but got %s\n", expected, out)
	}
}

func TestProgram1(t *testing.T) {
	code := ">>>>>>>>>>>>>>>>>>>>>>>>>>>>++++++++++++++++++++++++++[-<<[+<]+[>]>][<<[[-]-----<]>[>]>]<<[++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++<]>[.>]++++++++++."
	expected := "abcdefghijklmnopqrstuvwxyz\n"
	testProgram(code, expected, t)
}

func TestProgram2(t *testing.T) {
	code := "+++++++++++++++++++++++++[>++>+++>++++>+++++<<<<-]+++++++++++++++++++++++++>>+++++.>+.>-----------.------.<<<<---------------."
	expected := "Perl\n"
	testProgram(code, expected, t)
}

func TestProgram3(t *testing.T) {
	code := "+++++[>+++++[>++>++>+++>+++>++++>++++<<<<<<-]<-]+++++[>>[>]<[+.<<]>[++.>>>]<[+.<]>[-.>>]<[-.<<<]>[.>]<[+.<]<-]++++++++++."
	expected := "eL34NfeOL454KdeJ44JOdefePK55gQ67ShfTL787KegJ77JTeghfUK88iV9:XjgYL:;:KfiJ::JYfijgZK;;k[<=]lh^L=>=KgkJ==J^gklh_K>>m`?@bnicL@A@KhmJ@@JchmnidKAA\n"
	testProgram(code, expected, t)
}
