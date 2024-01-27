TITLE=Stack views in SwiftUI
DESCRIPTION=Stack views in SwiftUI
DATE=02/12/2022
+++
Code: <https://github.com/denniscmartin/dt-viewstack>

A SwiftUI library to stack two views like the Apple Maps app.

## Usage

```swift
GeometryReader { geoMap in
    DTViewStack(geo: geoMap) {
        Map(coordinateRegion: $region)
    } secondary: {
        Text("Hello world")
    } toolbar: {
        HStack {
            Button("Add") { }
        }
    }
}
```