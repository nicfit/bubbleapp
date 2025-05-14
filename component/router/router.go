package router

import (
	"log"
	"path"
	"strings"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/context"
)

// Route defines the structure for a single route.
type Route struct {
	Path      string  // Path segment (e.g., "users", ":id")
	Component app.FC  // Component to render for this route
	Children  []Route // Nested routes
}

// RouterProps defines the properties for the NewRouter component.
type RouterProps struct {
	Routes      []Route
	InitialPath string
	NotFound    app.FC // Component to render if no route matches
}

// --- RouterController ---

// RouterController manages routing state and navigation.
type RouterController struct {
	viewCtx     *app.Ctx
	History     []string
	Routes      []Route
	currentPath string
	notFoundFC  app.FC
}

// NewRouterController creates and initializes a new RouterController.
func NewRouterController(initialPath string, routes []Route, notFoundFC app.FC) *RouterController {
	rc := &RouterController{
		History:     make([]string, 0),
		Routes:      routes,
		currentPath: "/",
		notFoundFC:  notFoundFC,
	}
	if initialPath != "" {
		rc.currentPath = initialPath
	}
	rc.History = append(rc.History, rc.currentPath)
	return rc
}

// RegisterView allows RouterView to register its context for re-rendering.
func (rc *RouterController) RegisterView(ctx *app.Ctx) {
	rc.viewCtx = ctx
}

func (rc *RouterController) triggerRender() {
	if rc.viewCtx != nil {
		// Ensure re-render happens in the UI thread if necessary,
		// though bubbletea's model update cycle usually handles this.
		rc.viewCtx.Update()
	} else {
		log.Println("RouterController: viewCtx is nil, cannot trigger render. Ensure RouterView registers itself.")
	}
}

// Push navigates to a new path and adds it to history.
func (rc *RouterController) Push(newPath string) {
	cleanedPath := path.Clean(newPath)
	if rc.currentPath == cleanedPath {
		return
	}
	rc.currentPath = cleanedPath
	rc.History = append(rc.History, rc.currentPath)
	rc.triggerRender()
}

// Pop navigates to the previous path in history.
func (rc *RouterController) Pop() {
	if len(rc.History) <= 1 { // Can't pop the last/initial entry
		return
	}
	rc.History = rc.History[:len(rc.History)-1]
	rc.currentPath = rc.History[len(rc.History)-1]
	rc.triggerRender()
}

// Current returns the current active path.
func (rc *RouterController) Current() string {
	return rc.currentPath
}

// Replace replaces the current path in history with a new one.
func (rc *RouterController) Replace(newPath string) {
	cleanedPath := path.Clean(newPath)
	if rc.currentPath == cleanedPath && len(rc.History) > 0 {
		return
	}
	rc.currentPath = cleanedPath
	if len(rc.History) > 0 {
		rc.History[len(rc.History)-1] = rc.currentPath
	} else {
		rc.History = append(rc.History, rc.currentPath)
	}
	rc.triggerRender()
}

// ReplaceRoot clears history and navigates to a new root path.
func (rc *RouterController) ReplaceRoot(newPath string) {
	cleanedPath := path.Clean(newPath)
	rc.currentPath = cleanedPath
	rc.History = make([]string, 1)
	rc.History[0] = rc.currentPath
	rc.triggerRender()
}

// --- Contexts ---

// RouterContext holds the global RouterController instance.
// We store a pointer to RouterController in the context.
var RouterContext = context.Create(NewRouterController("/", nil, nil))

// UseRouterController is a hook to get the RouterController.
func UseRouterController(c *app.Ctx) *RouterController {
	return context.UseContext(c, RouterContext)
}

// CurrentMatchContextData holds information about the currently matched route.
type CurrentMatchContextData struct {
	MatchedRoute      *Route
	PathParams        map[string]string
	RemainingPath     string // The part of the URL not matched by this route, for Outlets
	MatchedPathPrefix string // The full path prefix matched by this route and its parents
}

// CurrentMatchContext holds data for the current route match.
// We store a pointer to CurrentMatchContextData.
var CurrentMatchContext = context.Create(CurrentMatchContextData{})

// UseCurrentMatch is a hook to get the current route match data.
func UseCurrentMatch(c *app.Ctx) CurrentMatchContextData {
	return context.UseContext(c, CurrentMatchContext)
}

// --- Router Setup ---

// NewRouter is the entry point to set up the router.
// It provides the RouterController to its children.
func NewRouter(c *app.Ctx, props RouterProps) string {
	// Create a new RouterController instance based on the provided props.
	// This instance will be specific to this Router setup.
	routerControllerInstance, _ := app.UseState(c, NewRouterController(props.InitialPath, props.Routes, props.NotFound))

	// Provide this specific routerControllerInstance to the children via RouterContext.
	return context.NewProvider(c, RouterContext, routerControllerInstance, func(providerCtx *app.Ctx) {
		// Render RouterView as the child of the provider.
		// RouterView will use UseRouterController to get the routerControllerInstance.
		c.Render(routerView, nil) // RouterView doesn't take props directly
	})
}

// routerView is an internal component that listens to path changes and renders the matched route.
func routerView(c *app.Ctx, rawProps app.Props) string {
	routerCtrl := UseRouterController(c)
	if routerCtrl == nil {
		log.Println("RouterView: RouterController is nil. Ensure NewRouter wraps this component.")
		return "Error: Router not initialized."
	}

	// Register this RouterView instance's context with the controller
	// so the controller can trigger re-renders on this instance.
	routerCtrl.RegisterView(c)

	// Perform route matching and render.
	// Initial match is against the full current path, with no parent prefix.
	return matchAndRender(c, routerCtrl.Routes, routerCtrl.currentPath, routerCtrl.currentPath, "", routerCtrl.notFoundFC)
}

// --- Matching Logic ---

// matchRoute attempts to match a single route definition's path against the current URL segment.
// routeDefPath: The path defined in RouteDefinition (e.g., "users", ":id").
// currentUrlSegment: The segment of the URL to match against (e.g., "/users/123/settings").
// Returns: params, isMatch, pathConsumedByThisRoute, remainingPathForChildren
func matchRoute(routeDefPath, currentUrlSegment string) (map[string]string, bool, string, string) {
	params := make(map[string]string)

	// Normalize paths: remove leading/trailing slashes for segment-wise comparison
	// but path.Clean handles this well for path.Join later.
	// For matching, we split by '/'
	cleanRouteDefPath := strings.Trim(routeDefPath, "/")
	cleanCurrentUrlSegment := strings.TrimPrefix(currentUrlSegment, "/") // Keep one leading slash if currentUrlSegment was "/"

	if cleanRouteDefPath == "" && (cleanCurrentUrlSegment == "" || currentUrlSegment == "/") { // Match for empty path (e.g. index route)
		return params, true, "", currentUrlSegment // Consumed nothing, remaining is the same (could be "/" or specific segment)
	}
	if cleanRouteDefPath == "" && currentUrlSegment != "/" && currentUrlSegment != "" { // Empty route path but non-empty URL segment
		return nil, false, "", ""
	}

	defParts := strings.Split(cleanRouteDefPath, "/")
	currentParts := strings.Split(cleanCurrentUrlSegment, "/")

	if cleanRouteDefPath != "" && cleanCurrentUrlSegment == "" && currentUrlSegment != "/" { // Non-empty route path but empty URL segment
		return nil, false, "", ""
	}

	if routeDefPath == "/" && currentUrlSegment == "/" {
		return params, true, "/", ""
	}
	if routeDefPath == "/" && currentUrlSegment != "/" {
		return nil, false, "", ""
	}

	if len(currentParts) < len(defParts) {
		return nil, false, "", ""
	}
	if len(defParts) == 1 && defParts[0] == "" && cleanRouteDefPath != "" {
		return nil, false, "", ""
	}

	pathConsumedParts := make([]string, 0, len(defParts))

	for i, defPart := range defParts {
		if i >= len(currentParts) { // Not enough segments in current URL
			return nil, false, "", ""
		}
		currentPart := currentParts[i]

		if strings.HasPrefix(defPart, ":") {
			paramName := defPart[1:]
			params[paramName] = currentPart
			pathConsumedParts = append(pathConsumedParts, currentPart)
		} else if defPart == currentPart {
			pathConsumedParts = append(pathConsumedParts, currentPart)
		} else {
			return nil, false, "", ""
		}
	}

	consumedPath := strings.Join(pathConsumedParts, "/")
	// Ensure consumedPath has a leading slash if currentUrlSegment did, and it's not empty.
	if strings.HasPrefix(currentUrlSegment, "/") && consumedPath != "" {
		consumedPath = "/" + consumedPath
	}

	remainingPath := ""
	if len(currentParts) > len(defParts) {
		remainingPath = "/" + strings.Join(currentParts[len(defParts):], "/")
	} else if len(currentParts) == len(defParts) && strings.HasSuffix(currentUrlSegment, "/") && len(defParts) > 0 {
		// If currentUrlSegment was "/foo/" and matched "/foo", remaining should be "/"
		// This is tricky. path.Clean might simplify.
		// For now, if exact match of all parts, remaining is empty unless original had trailing slash.
		// Let's assume paths are cleaned and don't typically end with slash unless it's the root "/".
	}

	// If consumedPath is empty but defParts was not (e.g. routeDefPath was "/"), consumedPath should be "/"
	if consumedPath == "" && (routeDefPath == "/" || (len(defParts) > 0 && defParts[0] == "")) {
		consumedPath = "/"
	}

	return params, true, consumedPath, remainingPath
}

type keyProps struct {
	Key string
}

// matchAndRender recursively matches routes and renders the component.
// fullUrl: The complete current URL from the RouterController.
// pathSegmentToMatch: The part of the fullUrl that this level is trying to match.
// accumulatedParentPrefix: The path prefix matched by parent routes (e.g., "/admin/users").
func matchAndRender(
	c *app.Ctx,
	routes []Route,
	fullUrl string,
	pathSegmentToMatch string,
	accumulatedParentPrefix string,
	notFoundFC app.FC,
) string {
	normalizedPathSegmentToMatch := path.Clean(pathSegmentToMatch)
	if normalizedPathSegmentToMatch == "." { // path.Clean can return "." for empty or "/"
		normalizedPathSegmentToMatch = "/"
	}

	for _, route := range routes {
		routeCopy := route // Important for pointer safety if taking &route later

		params, matched, pathConsumed, remainingPathForChildren := matchRoute(route.Path, normalizedPathSegmentToMatch)

		if matched {
			// Create the specific match data for this matched route.
			newMatchData := CurrentMatchContextData{
				MatchedRoute:      &routeCopy,
				PathParams:        params,
				RemainingPath:     remainingPathForChildren,
				MatchedPathPrefix: path.Join(accumulatedParentPrefix, pathConsumed),
			}

			// Provide this specific newMatchData to the matched component and its children (e.g., Outlet)
			// via CurrentMatchContext.
			return context.NewProvider(c, CurrentMatchContext, newMatchData, func(c *app.Ctx) {
				if routeCopy.Component != nil {
					ps := keyProps{
						Key: routeCopy.Path,
					}
					routeCopy.Component(c, ps)
				} else {
					// If no component, but has children, it's a layout/group route.
					// An Outlet component should be used explicitly within the parent's render flow
					// if children are meant to be rendered.
					// This router doesn't implicitly render an Outlet.
					log.Printf("Route matched (%s) but has no component. If it has children, ensure an <Outlet /> is used in its layout if it were a layout component, or provide a component.", newMatchData.MatchedPathPrefix)
					panic("Route matched but no component provided.")
				}
			})
		}
	}

	if notFoundFC != nil {
		return notFoundFC(c, nil)
	}
	return "route not found" // Or a default "Not Found" message
}

// --- Outlet ---

func NewOutlet(c *app.Ctx) string {
	routerCtrl := UseRouterController(c)

	return c.Render(outlet, outletProps{Key: routerCtrl.currentPath})
}

type outletProps struct {
	Key string
}

// outlet is a component that renders the matched child route.
func outlet(c *app.Ctx, _ app.Props) string {
	currentMatch := UseCurrentMatch(c)
	routerCtrl := UseRouterController(c) // For NotFoundFC and full URL if needed

	if currentMatch.MatchedRoute == nil {
		log.Println("Outlet: No current match context found. Outlet must be a child of a matched route component.")
		return ""
	}
	if len(currentMatch.MatchedRoute.Children) == 0 {
		return "" // No children to render
	}

	return matchAndRender(
		c,
		currentMatch.MatchedRoute.Children,
		routerCtrl.Current(),
		currentMatch.RemainingPath,
		currentMatch.MatchedPathPrefix,
		routerCtrl.notFoundFC,
	)
}

// --- Navigation Components/Functions ---

// LinkProps defines properties for the Link component.
type LinkProps struct {
	To       string
	Children app.FC // Content of the link
	// TODO: Add 'replace bool' option
}

// Link creates a navigational link.
func Link(c *app.Ctx, props LinkProps) string {
	routerCtrl := UseRouterController(c)
	if routerCtrl == nil {
		log.Println("Link: RouterController not found.")
		return "Error: Link cannot function."
	}

	app.UseAction(c, func(_ string) {
		routerCtrl.Push(props.To)
	})

	// Basic clickable text. Styling would be applied by parent or via props.
	// For now, just render children or the "To" path as text.
	// This assumes the UI layer can make this clickable and trigger actionID.
	// In a bubbletea context, this might be a styled string that, when selected,
	// posts a tea.Msg that eventually calls the action.
	// For simplicity, we assume the app.FC children know how to be clickable
	// or this is wrapped by a clickable primitive.

	var content string
	if props.Children != nil {
		content = c.Render(props.Children, nil)
	} else {
		content = props.To // Default link text
	}

	// This is a conceptual representation. Actual clickable behavior
	// depends on how `app.Ctx.UseAction` integrates with the UI elements.
	// We'll assume for now that rendering the content is enough, and
	// the parent context handles making it interactive.
	// A more complete solution might involve returning a specific component type
	// that the rendering engine knows how to make interactive.
	// E.g., return c.RenderClickable(content, actionID)
	return content // Placeholder: actual rendering of a clickable element needed
}

// NavigateOptions provides options for programmatic navigation.
type NavigateOptions struct {
	Replace bool
	// TODO: State any // For history state, if ever implemented
}

// NavigateOption defines a function that modifies NavigateOptions.
type NavigateOption func(*NavigateOptions)

// WithReplace is an option for Navigate to replace the current history entry.
func WithReplace(replace bool) NavigateOption {
	return func(o *NavigateOptions) {
		o.Replace = replace
	}
}

// Navigate allows programmatic navigation.
func Navigate(c *app.Ctx, to string, opts ...NavigateOption) {
	routerCtrl := UseRouterController(c)
	if routerCtrl == nil {
		log.Println("Navigate: RouterController not found.")
		return
	}

	options := NavigateOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	if options.Replace {
		routerCtrl.Replace(to)
	} else {
		routerCtrl.Push(to)
	}
}
