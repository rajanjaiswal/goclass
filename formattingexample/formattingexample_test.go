package formattingexample

import "testing"

func TestInitFormattingExamples(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				x: 1,
				y: 2,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitFormattingExamples(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitFormattingExamples() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InitFormattingExamples() = %v, want %v", got, tt.want)
			}
		})
	}
}
