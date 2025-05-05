# NOT WORKING RIGHT NOW

---

### [Grid](./examples/grid/main.go)

If you need more responsive layouts use a Grid which can span 12 unit across the width. Each item in the grid has a width for each breakpoint (the size of the terminal).

```go
gridView := grid.New(ctx,
    grid.Item[CustomData]{Xs: 12,
        Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.DangerDark,
            Child: text.New(ctx, "I wish I could center text! Some day...", nil),
        }),
    },
    grid.Item[CustomData]{Xs: 6,
        Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Success}),
    },
    grid.Item[CustomData]{Xs: 6,
        Item: stack.New(ctx, &stack.Options[CustomData]{
            Children: []*app.Base[CustomData]{
                text.New(ctx, "Background mess up if this text has foreground style.", nil),
                text.New(ctx, "Fix the margin to the left here. Not intentional.", nil),
                button.New(ctx, "BUTTON 1", &button.Options{Type: button.Compact}),
            },
        }),
    },
    grid.Item[CustomData]{Xs: 3,
        Item: button.New(ctx, "BUTTON 2", &button.Options{Variant: button.Danger, Type: button.Compact}),
    },
    grid.Item[CustomData]{Xs: 6,
        Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.InfoDark,
            Child: stack.New(ctx, &stack.Options[CustomData]{
                Children: []*app.Base[CustomData]{
                    text.New(ctx, "I am in a stack!", nil),
                    loader.New(ctx, loader.Meter, &loader.Options{Color: ctx.Styles.Colors.DangerDark, Text: "Loader is loading!"}),
                },
            }),
        }),
    },
    grid.Item[CustomData]{Xs: 3,
        Item: box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Warning}),
    },
)

base := app.New(ctx, app.AsRoot())
base.AddChild(gridView)
```

![Demo xs](demo_xs.gif)
![Demo](demo.gif)
