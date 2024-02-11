+++
title = 'A utility to round Corners in Swiftui'
date = 2022-11-28
draft = true
+++

A SwiftUI library to round specific corners of shapes.

Code: <https://github.com/denniscmartin/dt-roundedcorners>

# Usage

```swift
struct ContentView: View {
    var body: some View {
        Rectangle()
            .roundCorners(20, corners: [.bottomLeft, .bottomRight])
    }
}
```
