package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2,2)
	expected := 4

	if sum != expected{
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1,5)
	fmt.Println(sum)
	//Output: 6
}

func TestSum(t *testing.T) {

	numbers := [5]int{1,2,3,4,5}

	got := Sum(numbers)
	want := 15

	if got != want{
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}