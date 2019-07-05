package try

import (
	"testing"
)

func TestRun(t *testing.T) {

	var flagtests = []struct {
		function Task
	}{
		{func() {} },
		{
			func() {
				count := 11
				count++
			}, 
		},
		{
			func(){
				sampleVar := "000"

				if sampleVar == "" {
					t.Errorf("got %v want %v", sampleVar, "")
				}
			},
		},
	}

	for i, tt := range flagtests {
		t.Run(string(i), func(t *testing.T) {

			tryFunction := tt.function

			tryStatement := Run(tryFunction)

			if tryStatement.run == nil {
				t.Errorf("got %v, want %v", tryStatement.run, tryFunction)
			}
		})
	}
}

func TestCatch(t *testing.T) {

	var flagtests = []struct {
		function Error
	}{
		{func(e interface{}) {} },
		{
			func(e interface{}) {
				count := 11
				count++
			}, 
		},
		{
			func(e interface{}){
				sampleVar := "000"

				if sampleVar == "" {
					t.Errorf("got %v want %v", sampleVar, "")
				}
			},
		},
	}

	for i, tt := range flagtests {
		t.Run(string(i), func(t *testing.T) {

			tryFunction := func() {}

			tryStatement := Run(tryFunction).Catch(tt.function)

			if tryStatement.run == nil {
				t.Errorf("got %v, want %v", tryStatement.errorHandler, tryFunction)
			}
		})
	}
}

func TestThrow(t *testing.T) {

	var flagtests = []struct {
		function Task
		errorString string
	}{
		{	func() { 
			Throw("111") 
			},
			"111",
		},
		{
			func() {
				Throw("222")
			},
			"222",
		},
		{
			func(){
				Throw("400")
			},
			"400",
		},
		{
			func(){
				Throw("800")
			},
			"800",
		},
		{
			func(){
				Throw("202")
			},
			"202",
		},
	}

	for i, tt := range flagtests {
		t.Run(string(i), func(t *testing.T) {

			defer func() {
		        if r := recover(); r != nil {
		            if r.(string) != tt.errorString {
		            	t.Errorf("got %v, want %v", r, tt.errorString)
		            }
		        }
		    }()
			
			tt.function()
		})
	}
}