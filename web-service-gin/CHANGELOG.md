# CHANGELOG

## Version 0.2

### Added

- Added best effort idiomatic unit tests

## Version 0.1

### Added

- Added In memory fake persistence
- Added Create endpoint
- Added Get endpoint
- Added Get all endpoint

#### API
```typescript
type Album = {
    id: string,
    title: string,
    artist: string,
    price: number
}
```

```http request
POST /albums

Album
```

```http request
Get /albums/{id}

{}

Album
```


```http request
Get /albums/

{}

Album[]
```
