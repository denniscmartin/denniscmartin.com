TITLE=Photopicker wrapper in SwiftUI
DESCRIPTION=Photopicker wrapper in SwiftUI
DATE=25/11/2022
+++
A generic implementation of a photo picker in SwiftUI using PhotosUI.

Code: <https://github.com/denniscmartin/dt-photopicker>

## Usage

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