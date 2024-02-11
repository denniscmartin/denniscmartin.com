+++
title = 'A utility to stack views in Swiftui'
date = 2022-12-02
draft = true
+++

Code: <https://github.com/denniscmartin/dt-viewstack>

A SwiftUI library to stack two views like the Apple Maps app.

# Usage

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
