package service_test

import (
	"fmt"
	"testing"

	"github.com/pakawatkung/go-unit-test/service"
)

func TestCheckgrade(t *testing.T) {

	type testCase struct {
		name  string
		score int
		want  string
	}

	mytest := []testCase{
		{name: "Grade A", score: 80, want: "A"},
		{name: "Grade B", score: 70, want: "B"},
		{name: "Grade C", score: 60, want: "C"},
		{name: "Grade D", score: 50, want: "D"},
		{name: "Grade F", score: 30, want: "F"},
	}

	for _, test := range mytest {
		t.Run(test.name, func(t *testing.T) {
			grade := service.CheckGrade(test.score)
			want := test.want

			if grade != want {
				t.Errorf("got %v want %v", grade, want)
			}
		})
	}

}


func BenchmarkCheckGrade(b *testing.B) {
	for i:=0; i<b.N; i++ {
		service.CheckGrade(80)
	}
}

func ExampleCheckGrade(){
	grade := service.CheckGrade(80)
	fmt.Println(grade)
	// Output : A
}