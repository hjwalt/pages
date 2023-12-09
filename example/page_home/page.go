package page_home

import (
	"html/template"
	"net/http"

	"github.com/hjwalt/routes/example"
	"github.com/hjwalt/routes/example/component_sidebar"
	"github.com/hjwalt/routes/example/component_sidebar_button"
	"github.com/hjwalt/routes/example/component_sidebar_button_list"
	"github.com/hjwalt/routes/example/component_sidebar_item"
	"github.com/hjwalt/routes/example/component_sidebar_item_header"
	"github.com/hjwalt/routes/example/component_sidebar_item_list"
	"github.com/hjwalt/routes/example/page_error_500"
	"github.com/hjwalt/routes/page"
	"github.com/hjwalt/routes/runtime_chi"
	"github.com/hjwalt/runway/runtime"
)

const (
	directory = "page_home"
	path      = "/"
)

var Html = example.Page(directory + "/page.html")

type model struct {
	Sidebar template.HTML
}

func sidebar() page.Component[example.Context] {
	sidebarTop := component_sidebar_item_list.New(
		component_sidebar_item_list.Model{},
		component_sidebar_item.Model{Icon: "dashboard", Label: "Dashboard", Link: "/", Active: true},
		component_sidebar_item.Model{Icon: "table_view", Label: "Tables", Link: "/pages/tables.html", Active: false},
		component_sidebar_item.Model{Icon: "receipt_long", Label: "Billing", Link: "/billing", Active: false},
		component_sidebar_item.Model{Icon: "view_in_ar", Label: "Virtual Reality", Link: "/pages/virtual-reality.html", Active: false},
		component_sidebar_item.Model{Icon: "format_textdirection_r_to_l", Label: "RTL", Link: "/pages/rtl.html", Active: false},
		component_sidebar_item.Model{Icon: "notifications", Label: "Notifications", Link: "/pages/notifications.html", Active: false},
		component_sidebar_item_header.New(component_sidebar_item_header.Model{Label: "Account pages"}),
		component_sidebar_item.Model{Icon: "person", Label: "Profile", Link: "/pages/profile.html", Active: false},
		component_sidebar_item.Model{Icon: "login", Label: "Sign In", Link: "/pages/sign-in.html", Active: false},
		component_sidebar_item.Model{Icon: "assignment", Label: "Sign Up", Link: "/pages/sign-up.html", Active: false},
	)
	sidebarButton := component_sidebar_button_list.New(
		component_sidebar_button_list.Model{},
		component_sidebar_button.Model{Label: "Documentation", Link: "https://www.creative-tim.com/learning-lab/bootstrap/overview/material-dashboard?ref=sidebarfree", Outlined: true},
		component_sidebar_button.Model{Label: "Upgrade to pro", Link: "https://www.creative-tim.com/product/material-dashboard-pro?ref=sidebarfree", Outlined: false},
	)
	return component_sidebar.New(
		component_sidebar.Model{},
		sidebarTop,
		sidebarButton,
	)
}

func pageHandler(c example.Context, w http.ResponseWriter, r *http.Request) (*template.Template, page.ComponentMapModel[model], error) {
	model, err := page.ComponentMapCreateModel(
		c,
		w,
		r,
		model{},
		map[string]page.Component[example.Context]{
			"sidebar": sidebar(),
		},
	)
	return Html, model, err
}

func Get() runtime.Configuration[*runtime_chi.Runtime[example.Context]] {
	return runtime_chi.WithPage(path, http.MethodGet, pageHandler, page_error_500.Error)
}
