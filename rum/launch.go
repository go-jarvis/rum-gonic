package rum

func AddRootRoute(e *Engine, root Operator) *Engine {
	e.Register(root)

	return e
}
