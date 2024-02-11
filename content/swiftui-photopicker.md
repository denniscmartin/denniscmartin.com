+++
title = 'A photopicker wrapper in SwiftUI'
date = 2022-11-25
draft = true
+++

A generic implementation of a photo picker in SwiftUI using PhotosUI.

Code: <https://github.com/denniscmartin/dt-photopicker>

# Usage

```swift
import DTPhotoPicker

struct ContentView: View {
var body: some View {
        DTPhotoPicker { data, image in
            image
                .resizable()
                .scaledToFill()
        }
    }
}
```
