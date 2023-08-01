# Generational Index

## Intent

Allow an entity to query it's components without sacrificing performance

## Motivation

The Motivation section describes an example problem that we will be applying the pattern to.

One method of handling entities in game design is a entity component system (ECS). Using composition entities can be made of only the components they need rather than dumping all of that functionality into one object. It does invite new problems though, one of which is how do we store and retrieve components?

### Array

```go
type Entity struct {
  PhysicsId   int
  AnimationId int
  StateId     int
}

physics := []PhysicsComponents{}
animations := []AnimationComponents{}
states := []StateComponents{}
```

An array initial seems like the natural fit. We don't need to worry about sorting or searching since the entity should already have the index of whatever components it is using. However, deleting components from the array is `O(n)` and we now run the risk of the entity not knowing if the component ID it has is still valid.

### Hash table

```go
type Entity struct {
  PhysicsId   int
  AnimationId int
  StateId     int
}

physics := map[int]PhysicsComponents{}
animations := map[int]AnimationComponents{}
states := map[int]StateComponents{}
```

Hash tables solve the issue of deleting components being expensive. In addition, most implementations of hash tables include some form of a `key?()` method that the Entity could use to check if a component still exists. The downside? Data locality





## Pattern

The Pattern section distills the essence of the pattern out of the previous example.

