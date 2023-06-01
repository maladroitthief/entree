package application_test

import (
	"testing"

	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/settings"
)

func TestNewSettingsService(t *testing.T) {
	type args struct {
		logger logs.Logger
		repo   settings.Repository
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				logger: &logger{},
				repo:   &settingsRepository{},
			},
			wantErr: false,
		},
		{
			name: "nil logger",
			args: args{
				logger: nil,
				repo:   &settingsRepository{},
			},
			wantErr: true,
		},
		{
			name: "nil repo",
			args: args{
				logger: &logger{},
				repo:   nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := application.NewSettingsService(tt.args.logger, tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSettingsService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_settingsService_Update(t *testing.T) {
	type fields struct {
		repo settings.Repository
		log  logs.Logger
	}
	type args struct {
		args application.Inputs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "base",
			fields: fields{
				repo: &settingsRepository{},
				log:  &logger{},
			},
			args: args{
				args: application.Inputs{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewSettingsService(tt.fields.log, tt.fields.repo)
			if err != nil {
				t.Errorf("SettingsService.Update() failed to create a NewSettingsService()")
				return
			}

			if err := svc.Update(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("settingsService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_settingsService_IsAny(t *testing.T) {
	type fields struct {
		repo settings.Repository
		log  logs.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "base",
			fields: fields{
				repo: &settingsRepository{},
				log:  &logger{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewSettingsService(tt.fields.log, tt.fields.repo)
			if err != nil {
				t.Errorf("SettingsService.IsAny() failed to create a NewSettingsService()")
				return
			}

			if got := svc.IsAny(); got != tt.want {
				t.Errorf("settingsService.IsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_settingsService_IsPressed(t *testing.T) {
	type fields struct {
		repo settings.Repository
		log  logs.Logger
	}
	type args struct {
		i settings.Input
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "base",
			fields: fields{
				repo: &settingsRepository{},
				log:  &logger{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewSettingsService(tt.fields.log, tt.fields.repo)
			if err != nil {
				t.Errorf("SettingsService.IsPressed() failed to create a NewSettingsService()")
				return
			}

			if got := svc.IsPressed(tt.args.i); got != tt.want {
				t.Errorf("settingsService.IsPressed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_settingsService_IsJustPressed(t *testing.T) {
	type fields struct {
		repo settings.Repository
		log  logs.Logger
	}
	type args struct {
		i settings.Input
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "base",
			fields: fields{
				repo: &settingsRepository{},
				log:  &logger{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewSettingsService(tt.fields.log, tt.fields.repo)
			if err != nil {
				t.Errorf("SettingsService.IsJustPressed() failed to create a NewSettingsService()")
				return
			}

			if got := svc.IsJustPressed(tt.args.i); got != tt.want {
				t.Errorf("settingsService.IsJustPressed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_settingsService_GetCursor(t *testing.T) {
	type fields struct {
		repo settings.Repository
		log  logs.Logger
	}
	tests := []struct {
		name   string
		fields fields
		wantX  int
		wantY  int
	}{
		{
			name: "base",
			fields: fields{
				repo: &settingsRepository{},
				log:  &logger{},
			},
			wantX: 0,
			wantY: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewSettingsService(tt.fields.log, tt.fields.repo)
			if err != nil {
				t.Errorf("SettingsService.GetCursor() failed to create a NewSettingsService()")
				return
			}

			gotX, gotY := svc.GetCursor()
			if gotX != tt.wantX {
				t.Errorf("settingsService.GetCursor() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("settingsService.GetCursor() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func Test_settingsService_CurrentInputs(t *testing.T) {
	type fields struct {
		repo settings.Repository
		log  logs.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   []settings.Input
	}{
		{
			name: "base",
			fields: fields{
				repo: &settingsRepository{},
				log:  &logger{},
			},
			want: []settings.Input{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewSettingsService(tt.fields.log, tt.fields.repo)
			if err != nil {
				t.Errorf("SettingsService.CurrentInputs() failed to create a NewSettingsService()")
				return
			}

			got := svc.CurrentInputs()
			if len(got) != len(tt.want) {
				t.Errorf("settingsService.CurrentInputs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_settingsService_GetWindowHeight(t *testing.T) {
	type fields struct {
		repo settings.Repository
		log  logs.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "base",
			fields: fields{
				repo: &settingsRepository{},
				log:  &logger{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewSettingsService(tt.fields.log, tt.fields.repo)
			if err != nil {
				t.Errorf("SettingsService.GetWindowHeight() failed to create a NewSettingsService()")
				return
			}

			if got := svc.GetWindowHeight(); got != tt.want {
				t.Errorf("settingsService.GetWindowHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_settingsService_GetWindowWidth(t *testing.T) {
	type fields struct {
		repo settings.Repository
		log  logs.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "base",
			fields: fields{
				repo: &settingsRepository{},
				log:  &logger{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewSettingsService(tt.fields.log, tt.fields.repo)
			if err != nil {
				t.Errorf("SettingsService.GetWindowWidth() failed to create a NewSettingsService()")
				return
			}

			if got := svc.GetWindowWidth(); got != tt.want {
				t.Errorf("settingsService.GetWindowWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_settingsService_GetWindowTitle(t *testing.T) {
	type fields struct {
		repo settings.Repository
		log  logs.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "base",
			fields: fields{
				repo: &settingsRepository{},
				log:  &logger{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewSettingsService(tt.fields.log, tt.fields.repo)
			if err != nil {
				t.Errorf("SettingsService.GetWindowTitle() failed to create a NewSettingsService()")
				return
			}

			if got := svc.GetWindowTitle(); got != tt.want {
				t.Errorf("settingsService.GetWindowTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_settingsService_GetScale(t *testing.T) {
	type fields struct {
		repo settings.Repository
		log  logs.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "base",
			fields: fields{
				repo: &settingsRepository{},
				log:  &logger{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := application.NewSettingsService(tt.fields.log, tt.fields.repo)
			if err != nil {
				t.Errorf("SettingsService.GetScale() failed to create a NewSettingsService()")
				return
			}

			if got := svc.GetScale(); got != tt.want {
				t.Errorf("settingsService.GetScale() = %v, want %v", got, tt.want)
			}
		})
	}
}
