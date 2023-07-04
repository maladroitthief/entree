package service_test

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/scene"
	"github.com/maladroitthief/entree/service"
)

func TestNewSceneService(t *testing.T) {
	type args struct {
		logger      logs.Logger
		settingsSvc service.SettingsService
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				logger:      &logger{},
				settingsSvc: &settingsService{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.NewSceneService(tt.args.logger, tt.args.settingsSvc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSceneService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_sceneService_Update(t *testing.T) {
	type fields struct {
		log         logs.Logger
		settingsSvc service.SettingsService
	}
	type args struct {
		args service.Inputs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := service.NewSceneService(tt.fields.log, tt.fields.settingsSvc)
			if err != nil {
				t.Errorf("SceneService.Update() failed to create a NewSceneService()")
				return
			}

			if err := svc.Update(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("sceneService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sceneService_GetCamera(t *testing.T) {
	type fields struct {
		log         logs.Logger
		settingsSvc service.SettingsService
	}
	tests := []struct {
		name   string
		fields fields
		want   scene.Camera
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := service.NewSceneService(tt.fields.log, tt.fields.settingsSvc)
			if err != nil {
				t.Errorf("SceneService.GetCamera() failed to create a NewSceneService()")
				return
			}
			got := svc.GetCamera()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sceneService.GetCamera() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sceneService_GetEntities(t *testing.T) {
	type fields struct {
		log         logs.Logger
		settingsSvc service.SettingsService
	}
	tests := []struct {
		name   string
		fields fields
		want   []canvas.Entity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := service.NewSceneService(tt.fields.log, tt.fields.settingsSvc)
			if err != nil {
				t.Errorf("SceneService.GetEntities() failed to create a NewSceneService()")
				return
			}
			if got := svc.GetEntities(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sceneService.GetEntities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sceneService_GetBackgroundColor(t *testing.T) {
	type fields struct {
		log         logs.Logger
		settingsSvc service.SettingsService
	}
	tests := []struct {
		name   string
		fields fields
		want   color.Color
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := service.NewSceneService(tt.fields.log, tt.fields.settingsSvc)
			if err != nil {
				t.Errorf("SceneService.GetBackgroundColor() failed to create a NewSceneService()")
				return
			}
			if got := svc.GetBackgroundColor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sceneService.GetBackgroundColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sceneService_GoTo(t *testing.T) {
	type fields struct {
		log         logs.Logger
		settingsSvc service.SettingsService
	}
	type args struct {
		s scene.Scene
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := service.NewSceneService(tt.fields.log, tt.fields.settingsSvc)
			if err != nil {
				t.Errorf("SceneService.GoTo() failed to create a NewSceneService()")
				return
			}
			if err := svc.GoTo(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("sceneService.GoTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
