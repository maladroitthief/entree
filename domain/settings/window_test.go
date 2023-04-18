package settings_test

import (
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/domain/settings"
)

func TestDefaultWindowSettings(t *testing.T) {
	tests := []struct {
		name string
		want settings.WindowSettings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := settings.DefaultWindowSettings(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultWindowSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindowSettings_Validate(t *testing.T) {
	tests := []struct {
		name    string
		w       *settings.WindowSettings
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("WindowSettings.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
