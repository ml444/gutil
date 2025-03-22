package str

import "testing"

func TestReplaceEnvVariables(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// TODO: Add test cases.
		{
			name:  "test1",
			input: "hello, ${USER}",
			want:  "hello, cml",
		},
		{
			name:  "test2",
			input: "hello, ${USER} ${HOME}",
			want:  "hello, cml /Users/cml",
		},
		{
			name:  "test3",
			input: "hello, $USER {HOME} $PATH}",
			want:  "hello, $USER {HOME} $PATH}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceEnvVariables(tt.input); got != tt.want {
				t.Errorf("ReplaceEnvVariables() = %v, want %v", got, tt.want)
			}
		})
	}
}
