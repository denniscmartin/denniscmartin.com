+++
title = 'Stock Charts, my first open source project'
date = 2021-04-28
thumbnail = 'thumbnail-stock-charts.png'
draft = false
+++

*This project is now deprecated in favor of the [Apple's native framework](https://developer.apple.com/documentation/charts). At the time of archiving this project, it has accumulated 98 stars and 11 forks on Github.*   

StockCharts is a library to create intertactive charts in SwiftUI.

![StockCharts screenshot](stock-charts-1.png)

## Installation

- In Xcode go to `File` -> `Swift packages` -> `Add package dependency`
- Copy and paste <https://github.com/denniscmartin/stock-charts.git>

I've created a demo app using StockCharts called Trades. Check out the code
in my [Github](https://github.com/denniscmartin/trades-demo)

### Usage

```swift
import StockCharts
```

#### Line chart

```swift
let lineChartController = LineChartController(prices: [Double])
LineChartView(lineChartController: lineChartController)
```

You can customise the line chart with LineChartController

```swift
LineChartController(
    prices: [Double],
    dates: [String]?, // format: yy-MM-dd
    hours: [String]?, // has to correspond to dates
    labelColor: Color,
    indicatorPointColor: Color,
    showingIndicatorLineColor: Color,
    flatTrendLineColor: Color,
    uptrendLineColor: Color,
    downtrendLineColor: Color,
    dragGesture: Bool
)
```

To enable the drag gesture set `dragGesture` to `true` in the `LineChartController`

```swift
LineChartView(
    lineChartController:
        LineChartController(
            prices: [Double],
            dragGesture: true
        )
)
```

#### Capsule chart

```swift
CapsuleChartView(percentageOfWidth: CGFloat)
// percentageOfWidth: must be 0 <= x <= 1

```

```swift
import SwiftUI
import StockCharts

struct ContentView: View {
    var body: some View {
        RoundedRectangle(cornerRadius: 25)
            .frame(width: 400, height: 120)
            .foregroundColor(.white)
            .shadow(color: Color(.gray).opacity(0.15), radius: 10)
            .overlay(
                VStack(alignment: .leading) {
                    Text("Dennis Concepcion")
                        .font(.title3)
                        .fontWeight(.semibold)

                    Text("Random guy")

                    CapsuleChartView(percentageOfWidth: 0.6, style: CapsuleChartStyle(capsuleColor: Color.blue))
                        .padding(.top)
                }
                .padding()
            )
    }
}
```