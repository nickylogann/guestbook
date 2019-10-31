package http

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo"

	"github.com/nickylogan/guestbook/internal/app/usecase/user"
	"github.com/nickylogan/guestbook/internal/app/usecase/visitor"
	"github.com/nickylogan/guestbook/internal/pkg/utils/template"
)

type UserHandler struct {
	UUseCase user.UseCase
	VUseCase visitor.UseCase
}

// NewUserHandler will initialize the user endpoint
func NewUserHandler(e *echo.Echo, us user.UseCase, vs visitor.UseCase) {
	handler := &UserHandler{
		UUseCase: us,
		VUseCase: vs,
	}
	e.Renderer = template.NewRenderer("./web/templates/*.html", true)
	e.GET("/", handler.Index)
}

func (u *UserHandler) Index(c echo.Context) error {
	req := c.Request()

	// Retrieve query params
	pageS := c.QueryParam("p")
	page, _ := strconv.Atoi(pageS)
	if page <= 0 {
		page = 1
	}
	filter := c.QueryParam("f")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	// Fetch users
	userResp, err := u.UUseCase.FetchAll(ctx, filter, page)
	if err != nil {
		log.Println("[error][http][Index] failed to fetch users:", err)
		return c.HTML(http.StatusInternalServerError, err.Error())
	}

	// Fetch visitors
	visitorResp, err := u.VUseCase.Visit(ctx)
	if err != nil {
		log.Println("[error][http][Index] failed to add visit:", err)
		return c.HTML(http.StatusInternalServerError, err.Error())
	}

	type pageButton struct {
		Page int
		URL  string
	}
	// Build page buttons
	var pages []pageButton
	for i := userResp.Start; i <= userResp.End; i++ {
		pages = append(pages, pageButton{
			Page: i,
			URL:  buildURL(req, i, filter),
		})
	}

	// Build prev/next url
	nextURL := buildURL(req, page+1, filter)
	prevURL := buildURL(req, page-1, filter)

	data := map[string]interface{}{
		"filter":   filter,
		"users":    userResp.Data,
		"nextPage": userResp.NextPage,
		"nextUrl":  nextURL,
		"prevPage": userResp.PrevPage,
		"prevUrl":  prevURL,
		"pages":    pages,
		"page":     page,
		"visitors": visitorResp,
	}
	return c.Render(http.StatusOK, "index.html", data)
}

func buildURL(req *http.Request, page int, filter string) string {
	q := url.Values{}

	// Add page
	q.Add("p", strconv.Itoa(page))

	// Add filter
	if filter != "" {
		q.Add("f", filter)
	}

	req.URL.RawQuery = q.Encode()
	return req.URL.String()
}
