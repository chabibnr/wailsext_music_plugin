package music_plugin

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

// ---------------- Plugin Setup ----------------
// This is the main plugin struct. It can be named anything you like.
// It must implement the application.Plugin interface.
// Both the Init() and Shutdown() methods are called synchronously when the app starts and stops.

type Config struct {
	// Add any configuration options here
}

type Plugin struct {
	config *Config
	app    *application.App
}

func NewPlugin(config *Config) *Plugin {
	return &Plugin{
		config: config,
	}
}

// Shutdown is called when the app is shutting down
// You can use this to clean up any resources you have allocated
func (p *Plugin) Shutdown() {}

// Name returns the name of the plugin.
// You should use the go module format e.g. github.com/myuser/myplugin
func (p *Plugin) Name() string {
	return "github.com/myuser/example_plugin"
}

// Init is called when the app is starting up. You can use this to
// initialise any resources you need. You can also access the application
// instance via the app property.
func (p *Plugin) Init(app *application.App) error {
	p.app = app
	return nil
}

// Exported returns a list of exported methods that can be called from the frontend
func (p *Plugin) CallableByJS() []string {
	return []string{
		"Greet",
	}
}

// InjectJS returns any JS that should be injected into the frontend
func (p *Plugin) InjectJS() string {
	return ""
}

// ---------------- Plugin Methods ----------------
// Plugin methods are just normal Go methods. You can add as many as you like.
// The only requirement is that they are exported (start with a capital letter).
// You can also return any type that is JSON serializable.
// Any methods that you want to be callable from the frontend must be returned by the
// CallableByJS() method above.
// See https://golang.org/pkg/encoding/json/#Marshal for more information.

func (p *Plugin) Greet(name string) string {
	return "Hello " + name
}
