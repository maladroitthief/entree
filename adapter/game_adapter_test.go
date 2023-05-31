package adapter_test

import (
	"image"
	"image/color"
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/adapter"
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/sprite"
)

func TestNewGameAdapter(t *testing.T) {
	type args struct {
		log         logs.Logger
		sceneSvc    application.SceneService
		graphicsSvc application.GraphicsService
		settingsSvc application.SettingsService
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			wantErr: false,
		},
		{
			name: "no logger",
			args: args{
				log:         nil,
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			wantErr: true,
		},
		{
			name: "no scene service",
			args: args{
				log:         &logger{},
				sceneSvc:    nil,
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			wantErr: true,
		},
		{
			name: "no graphics service",
			args: args{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: nil,
				settingsSvc: &settingsService{},
			},
			wantErr: true,
		},
		{
			name: "no settings service",
			args: args{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := adapter.NewGameAdapter(tt.args.log, tt.args.sceneSvc, tt.args.graphicsSvc, tt.args.settingsSvc)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewGameAdapter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGameAdapter_Update(t *testing.T) {
	type fields struct {
		log         logs.Logger
		sceneSvc    application.SceneService
		graphicsSvc application.GraphicsService
		settingsSvc application.SettingsService
	}
	type args struct {
		args adapter.UpdateArgs
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
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			args: args{
				args: adapter.UpdateArgs{
					CursorX: 0,
					CursorY: 0,
					Inputs:  []string{""},
				},
			},
			wantErr: false,
		},
		{
			name: "bad cursor",
			fields: fields{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			args: args{
				args: adapter.UpdateArgs{
					CursorX: -1,
					CursorY: 0,
					Inputs:  []string{""},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ga, err := adapter.NewGameAdapter(
				tt.fields.log,
				tt.fields.sceneSvc,
				tt.fields.graphicsSvc,
				tt.fields.settingsSvc,
			)

			if err != nil {
				t.Errorf("GameAdapter.Update() failed to create a NewGameAdapter()")
			}

			err = ga.Update(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GameAdapter.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGameAdapter_GetEntities(t *testing.T) {
	type fields struct {
		log         logs.Logger
		sceneSvc    application.SceneService
		graphicsSvc application.GraphicsService
		settingsSvc application.SettingsService
	}
	tests := []struct {
		name   string
		fields fields
		want   []canvas.Entity
	}{
		{
			name: "base",
			fields: fields{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			want: []canvas.Entity{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ga, err := adapter.NewGameAdapter(
				tt.fields.log,
				tt.fields.sceneSvc,
				tt.fields.graphicsSvc,
				tt.fields.settingsSvc,
			)

			if err != nil {
				t.Errorf("GameAdapter.GetEntities() failed to create a NewGameAdapter()")
			}

			got := ga.GetEntities()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameAdapter.GetEntities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameAdapter_GetSpriteSheet(t *testing.T) {
	type fields struct {
		log         logs.Logger
		sceneSvc    application.SceneService
		graphicsSvc application.GraphicsService
		settingsSvc application.SettingsService
	}
	type args struct {
		sheet string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    sprite.SpriteSheet
		wantErr bool
	}{
		{
			name: "base",
			fields: fields{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			args: args{
				sheet: "",
			},
			want:    &spriteSheet{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ga, err := adapter.NewGameAdapter(
				tt.fields.log,
				tt.fields.sceneSvc,
				tt.fields.graphicsSvc,
				tt.fields.settingsSvc,
			)

			if err != nil {
				t.Errorf("GameAdapter.GetSpriteSheet() failed to create a NewGameAdapter()")
			}

			got, err := ga.GetSpriteSheet(tt.args.sheet)
			if (err != nil) != tt.wantErr {
				t.Errorf("GameAdapter.GetSpriteSheet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameAdapter.GetSpriteSheet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameAdapter_GetSpriteRectangle(t *testing.T) {
	type fields struct {
		log         logs.Logger
		sceneSvc    application.SceneService
		graphicsSvc application.GraphicsService
		settingsSvc application.SettingsService
	}
	type args struct {
		sheet  string
		sprite string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    image.Rectangle
		wantErr bool
	}{
		{
			name: "base",
			fields: fields{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			args: args{
				sheet:  "",
				sprite: "",
			},
			want:    image.Rectangle{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ga, err := adapter.NewGameAdapter(
				tt.fields.log,
				tt.fields.sceneSvc,
				tt.fields.graphicsSvc,
				tt.fields.settingsSvc,
			)

			if err != nil {
				t.Errorf("GameAdapter.GetSpriteRectangle() failed to create a NewGameAdapter()")
			}

			got, err := ga.GetSpriteRectangle(tt.args.sheet, tt.args.sprite)
			if (err != nil) != tt.wantErr {
				t.Errorf("GameAdapter.GetSpriteRectangle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameAdapter.GetSpriteRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameAdapter_Layout(t *testing.T) {
	type fields struct {
		log         logs.Logger
		sceneSvc    application.SceneService
		graphicsSvc application.GraphicsService
		settingsSvc application.SettingsService
	}
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantScreenWidth  int
		wantScreenHeight int
	}{
		{
			name: "base",
			fields: fields{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			args: args{
				width:  0,
				height: 0,
			},
			wantScreenWidth:  0,
			wantScreenHeight: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ga, err := adapter.NewGameAdapter(
				tt.fields.log,
				tt.fields.sceneSvc,
				tt.fields.graphicsSvc,
				tt.fields.settingsSvc,
			)

			if err != nil {
				t.Errorf("GameAdapter.Layout() failed to create a NewGameAdapter()")
			}

			gotScreenWidth, gotScreenHeight := ga.Layout(tt.args.width, tt.args.height)
			if gotScreenWidth != tt.wantScreenWidth {
				t.Errorf("GameAdapter.Layout() gotScreenWidth = %v, want %v", gotScreenWidth, tt.wantScreenWidth)
			}
			if gotScreenHeight != tt.wantScreenHeight {
				t.Errorf("GameAdapter.Layout() gotScreenHeight = %v, want %v", gotScreenHeight, tt.wantScreenHeight)
			}
		})
	}
}

func TestGameAdapter_GetWindowSize(t *testing.T) {
	type fields struct {
		log         logs.Logger
		sceneSvc    application.SceneService
		graphicsSvc application.GraphicsService
		settingsSvc application.SettingsService
	}
	tests := []struct {
		name             string
		fields           fields
		wantScreenWidth  int
		wantScreenHeight int
	}{
		{
			name: "base",
			fields: fields{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			wantScreenWidth:  0,
			wantScreenHeight: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ga, err := adapter.NewGameAdapter(
				tt.fields.log,
				tt.fields.sceneSvc,
				tt.fields.graphicsSvc,
				tt.fields.settingsSvc,
			)

			if err != nil {
				t.Errorf("GameAdapter.GetWindowSize() failed to create a NewGameAdapter()")
			}

			gotScreenWidth, gotScreenHeight := ga.GetWindowSize()
			if gotScreenWidth != tt.wantScreenWidth {
				t.Errorf("GameAdapter.GetWindowSize() gotScreenWidth = %v, want %v", gotScreenWidth, tt.wantScreenWidth)
			}
			if gotScreenHeight != tt.wantScreenHeight {
				t.Errorf("GameAdapter.GetWindowSize() gotScreenHeight = %v, want %v", gotScreenHeight, tt.wantScreenHeight)
			}
		})
	}
}

func TestGameAdapter_GetWindowTitle(t *testing.T) {
	type fields struct {
		log         logs.Logger
		sceneSvc    application.SceneService
		graphicsSvc application.GraphicsService
		settingsSvc application.SettingsService
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "base",
			fields: fields{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ga, err := adapter.NewGameAdapter(
				tt.fields.log,
				tt.fields.sceneSvc,
				tt.fields.graphicsSvc,
				tt.fields.settingsSvc,
			)

			if err != nil {
				t.Errorf("GameAdapter.GetWindowTitle() failed to create a NewGameAdapter()")
			}

			if got := ga.GetWindowTitle(); got != tt.want {
				t.Errorf("GameAdapter.GetWindowTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameAdapter_GetScale(t *testing.T) {
	type fields struct {
		log         logs.Logger
		sceneSvc    application.SceneService
		graphicsSvc application.GraphicsService
		settingsSvc application.SettingsService
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "base",
			fields: fields{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ga, err := adapter.NewGameAdapter(
				tt.fields.log,
				tt.fields.sceneSvc,
				tt.fields.graphicsSvc,
				tt.fields.settingsSvc,
			)

			if err != nil {
				t.Errorf("GameAdapter.GetScale() failed to create a NewGameAdapter()")
			}

			if got := ga.GetScale(); got != tt.want {
				t.Errorf("GameAdapter.GetScale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameAdapter_GetBackgroundColor(t *testing.T) {
	type fields struct {
		log         logs.Logger
		sceneSvc    application.SceneService
		graphicsSvc application.GraphicsService
		settingsSvc application.SettingsService
	}
	tests := []struct {
		name   string
		fields fields
		want   color.Color
	}{
		{
			name: "base",
			fields: fields{
				log:         &logger{},
				sceneSvc:    &sceneService{},
				graphicsSvc: &graphicsService{},
				settingsSvc: &settingsService{},
			},
			want: color.Black,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ga, err := adapter.NewGameAdapter(
				tt.fields.log,
				tt.fields.sceneSvc,
				tt.fields.graphicsSvc,
				tt.fields.settingsSvc,
			)

			if err != nil {
				t.Errorf("GameAdapter.GetBackgroundColor() failed to create a NewGameAdapter()")
			}

			if got := ga.GetBackgroundColor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameAdapter.GetBackgroundColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
