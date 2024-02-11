+++
title = 'Experimenting with Swift networking'
date = 2022-02-02
draft = false
+++

An experiment with Swift networking.

Code: <https://github.com/denniscmartin/bazooka>

# Installation

- In Xcode go to `File` -> `Swift packages` -> `Add package dependency`

- Copy and paste https://github.com/denniscmartin/Bazooka

# Usage

```bash
import Bazooka

let bazooka = Bazooka()
bazooka.request(url: "https://someurl.come", model: MyModel.self) { response in
    print(response)
}
```

```swift
import SwiftUI
import Bazooka

struct ContentView: View {
    var body: some View {
        Text("Hello, world!")
            .onAppear {
                let bazooka = Bazooka()
                bazooka.request(url: "url", model: MyModel.self) { response in
                    print(response)
                }
            }
    }
}

struct MyModel: Codable {
    var myVar: String
}
```
