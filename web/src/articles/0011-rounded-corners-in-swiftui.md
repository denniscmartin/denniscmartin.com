TITLE=Round corners of shapes in SwiftUI
DESCRIPTION=Round corners of shapes in SwiftUI
DATE=28/11/2022
+++
A SwiftUI library to round specific corners of shapes.

Code: <https://github.com/denniscmartin/dt-roundedcorners>

## Usage

```swift
struct ContentView: View {
    var body: some View {
        Rectangle()
            .roundCorners(20, corners: [.bottomLeft, .bottomRight])
    }
}
```