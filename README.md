# entree

Game where you fight food

## architecture

```mermaid
graph LR
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
    graphicsService
    gameAdapter --> graphicsService
  end

  subgraph Domain
    action
    canvas
    ebiten --> canvas
    gameAdapter --> canvas
    sceneService --> canvas
    scene
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
