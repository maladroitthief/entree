package settings_test

import (
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/domain/settings"
)

func TestSettingsDefaults(t *testing.T) {
	tests := []struct {
		name string
		want settings.Settings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := settings.SettingsDefaults(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SettingsDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}
