package settings_test

import (
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/domain/settings"
)

func TestInputSettings_Validate(t *testing.T) {
	tests := []struct {
		name    string
		i       *settings.InputSettings
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("InputSettings.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAllInputs(t *testing.T) {
	tests := []struct {
		name string
		want []settings.Input
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := settings.AllInputs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllInputs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultInputSettings(t *testing.T) {
	tests := []struct {
		name string
		want settings.InputSettings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := settings.DefaultInputSettings(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultInputSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultKeyboard(t *testing.T) {
	tests := []struct {
		name string
		want map[settings.Input]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := settings.DefaultKeyboard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultKeyboard() = %v, want %v", got, tt.want)
			}
		})
	}
}
