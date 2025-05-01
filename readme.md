# BubbleApp

> [!WARNING]
> This is work in progress and just exploration for now

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

### [Hello World!](./examples/hello-world/main.go)

```go
stack := stack.New(ctx, &stack.Options[struct{}]{
    Children: []*app.Base[struct{}]{
        text.New(ctx, "Hello World!", nil),
        divider.New(ctx),
        text.New(ctx, "Press [q] to quit.", nil),
    }},
)

base := app.New(ctx, app.AsRoot())
base.AddChild(stack)
```

![Hello world!](./examples/hello-world/demo.gif)

---

### [Multiple Views](./examples/multiple-views/main.go)

![Multiple Views](./examples/multiple-views/demo.gif)

---

### [Tabbing](./examples/tabbing/main.go)

```go
boxFill := box.New(ctx, &box.Options[CustomData]{})
addButton := button.New(ctx, "Button 1", &button.Options{Variant: button.Primary})
quitButton := button.New(ctx, "Quit App", &button.Options{Variant: button.Danger})

stack := stack.New(ctx, &stack.Options[CustomData]{
    Children: []*app.Base[CustomData]{
        text.New(ctx, "Tab through the buttons to see focus state!", nil),
        addButton,
        boxFill,
        divider.New(ctx),
        quitButton,
    }},
)

base := app.New(ctx, app.AsRoot())
base.AddChild(stack)

return model[CustomData]{
    base:         base,
    containerID:  boxFill.ID,
    addButtonID:  addButton.ID,
    quitButtonID: quitButton.ID,
}
```

```go
case button.ButtonPressMsg:
    switch msg.ID {
    case m.quitButtonID:
        return m, tea.Quit
    case m.addButtonID:
        m.base.GetChild(m.containerID).AddChild(
            text.New(m.base.Ctx, "Button pressed"),
        )
        return m, nil
    }
```

![Tabbing](./examples/tabbing/demo.gif)

---

### [Stack](./examples/stack/main.go)

```go
stack := stack.New(ctx, &stack.Options[CustomData]{
    Children: []*app.Base[CustomData]{
        box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Danger}),
        box.New(ctx, &box.Options[CustomData]{
            Child: stack.New(ctx, &stack.Options[CustomData]{
                Horizontal: true,
                Children: []*app.Base[CustomData]{
                    box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Primary}),
                    box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Secondary}),
                    box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Tertiary}),
                }},
            ),
        }),
        box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Warning}),
    }},
)

base := app.New(ctx, app.AsRoot())
base.AddChild(stack)
```

![Stack](./examples/stack/demo.gif)

---

### [Tabs](./examples/tabs/main.go)

```go
var tabsData = []tabs.TabElement[CustomData]{
	{
		Title:   "Overview",
		Content: NewOverview,
	},
	{
		Title:   "Loaders",
		Content: NewLoaders,
	},
	{
		Title:   "Scolling",
		Content: NewScrolling,
	},
}
```

```go
tabs := tabs.New(ctx, tabsData)

base := app.New(ctx, app.AsRoot())
base.AddChild(tabs)
```

![Tabs](./examples/tabs/demo.gif)

---

### [Grid](./examples/grid/main.go)

```go
gridView := grid.New(ctx,
    grid.Item[CustomData]{
        Xs: 12,
        Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.PrimaryDark,
            Child: text.New(ctx, "I wish I could center text! Some day...", nil),
        }),
    },
    grid.Item[CustomData]{
        Xs:   6,
        Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.InfoLight}),
    },
    grid.Item[CustomData]{
        Xs: 6,
        Item: stack.New(ctx, &stack.Options[CustomData]{
            Children: []*app.Base[CustomData]{
                text.New(ctx, "Background mess up if this text has foreground style.", nil),
                text.New(ctx, "Fix the margin to the left here. Not intentional.", nil),
                button.New(ctx, "BUTTON 1", nil),
            },
        }),
    },
    grid.Item[CustomData]{
        Xs:   3,
        Item: button.New(ctx, "BUTTON 2", &button.Options{Variant: button.Danger}),
    },
    grid.Item[CustomData]{
        Xs: 6,
        Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.InfoDark,
            Child: stack.New(ctx, &stack.Options[CustomData]{
                Children: []*app.Base[CustomData]{
                    text.New(ctx, "I am in a stack!", nil),
                    loader.New(ctx, loader.Meter, &loader.Options{
                        Text:  "Text style messes up bg. Fix!",
                        Color: ctx.Styles.Colors.Black,
                    }),
                },
            }),
        }),
    },
    grid.Item[CustomData]{
        Xs:   3,
        Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Success}),
    },
)

base := app.New(ctx, app.AsRoot())
base.AddChild(gridView)
```

![Grid](./examples/grid/demo.gif)
