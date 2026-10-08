package main

import "testing"

func TestCompact(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"object", "{\n  \"b\": 1,\n  \"a\": [1, 2, 3]\n}", `{"b":1,"a":[1,2,3]}`, false},
		{"spaces inside strings kept", `{ "k" : "a  b" }`, `{"k":"a  b"}`, false},
		{"number format kept", `[ 1.50, 1e3 ]`, `[1.50,1e3]`, false},
		{"invalid", `{"a": }`, "", true},
		{"empty", ``, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compact([]byte(tt.in))
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tt.wantErr)
			}
			if string(got) != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}