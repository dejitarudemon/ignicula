package main

import "testing"

func TestResolveAddr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		flag string
		env  string
		want string
	}{
		{name: "default", want: defaultAddr},
		{name: "env", env: "0.0.0.0", want: "0.0.0.0"},
		{name: "flag_over_env", flag: "192.168.0.1", env: "0.0.0.0", want: "192.168.0.1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := resolveAddr(tt.flag, tt.env)
			if err != nil {
				t.Fatalf("resolveAddr() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("resolveAddr() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestResolvePort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		flag    uint
		env     string
		want    uint16
		wantErr bool
	}{
		{name: "default", want: defaultPort},
		{name: "env", env: "8080", want: 8080},
		{name: "flag_over_env", flag: 9000, env: "8080", want: 9000},
		{name: "bad_env", env: "nope", wantErr: true},
		{name: "zero_env", env: "0", wantErr: true},
		{name: "flag_too_large", flag: 70000, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := resolvePort(tt.flag, tt.env)
			if tt.wantErr {
				if err == nil {
					t.Fatal("resolvePort() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("resolvePort() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("resolvePort() = %d, want %d", got, tt.want)
			}
		})
	}
}
