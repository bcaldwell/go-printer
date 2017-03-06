package printer

import "testing"

func TestGreen(t *testing.T) {
	type args struct {
		in0 text
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Green(tt.args.in0); got != tt.want {
				t.Errorf("Green() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreenf(t *testing.T) {
	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greenf(tt.args.text, tt.args.a...); got != tt.want {
				t.Errorf("Greenf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRed(t *testing.T) {
	type args struct {
		in0 text
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Red(tt.args.in0); got != tt.want {
				t.Errorf("Red() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedf(t *testing.T) {
	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Redf(tt.args.text, tt.args.a...); got != tt.want {
				t.Errorf("Redf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlue(t *testing.T) {
	type args struct {
		in0 text
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Blue(tt.args.in0); got != tt.want {
				t.Errorf("Blue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBluef(t *testing.T) {
	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bluef(tt.args.text, tt.args.a...); got != tt.want {
				t.Errorf("Bluef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYellow(t *testing.T) {
	type args struct {
		in0 text
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Yellow(tt.args.in0); got != tt.want {
				t.Errorf("Yellow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYellowf(t *testing.T) {
	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Yellowf(tt.args.text, tt.args.a...); got != tt.want {
				t.Errorf("Yellowf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBold(t *testing.T) {
	type args struct {
		in0 text
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bold(tt.args.in0); got != tt.want {
				t.Errorf("Bold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoldf(t *testing.T) {
	type args struct {
		text string
		a    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Boldf(tt.args.text, tt.args.a...); got != tt.want {
				t.Errorf("Boldf() = %v, want %v", got, tt.want)
			}
		})
	}
}
