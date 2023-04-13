# Entrée

Game where you fight food

## Architecture

### System

```mermaid
graph
  subgraph Infrastructure
    ebiten
    settingsRepository
  end

  subgraph Adapter
    gameAdapter
    ebiten --> gameAdapter
  end

  subgraph Application
    sceneService
    gameAdapter --> sceneService
    settingsService
    gameAdapter --> settingsService
    sceneService --> settingsService
    graphicsService
    gameAdapter --> graphicsService
  end

  subgraph Domain
    canvas
    canvas --> settings
    ebiten --> canvas
    gameAdapter --> canvas
    sceneService --> canvas
    scene
    scene --> canvas
    scene --> settings
    sceneService --> scene
    settings
    settingsRepository --> settings
    gameAdapter --> settings
    settingsService --> settings
    sprite
    gameAdapter --> sprite
    graphicsService --> sprite
  end
```
