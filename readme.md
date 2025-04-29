# BubbleApp

[!WARNING]
This is work in progress and just exploration for now

An opinionated App Framework for BubbleTea. Building large BubbleTea apps can be a lot of manual work. For every state of the app you need to store that state in your model and branch accordingly.

With BubbleApp you can compose models and gain things like

- **Composable Components** -
- **Mouse support** - using [BubbleZone](https://github.com/lrstanley/bubblezone)
  - Propagate mouse events to the right component automatically
- **Focus management**
  - Tab through your entire UI tree without any code
- **Global Ticks** that all components share (for animations like loaders)
  - Adding several Spinners from Bubbles is really slow over SSH. Each have their own Ticks.

## Examples

Hello, World!
![Tabbing](./examples/hello-world/demo.gif)

Multiple Views
![Multiple Views](./examples/multiple-views/demo.gif)

Tabbing
![Tabbing](./examples/tabbing/demo.gif)

Stack
![Stack](./examples/stack/demo.gif)
