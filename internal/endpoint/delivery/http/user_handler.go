package http

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo"

	"github.com/nickylogan/guestbook/internal/endpoint/usecase/user"
	"github.com/nickylogan/guestbook/internal/pkg/utils/template"
)

type UserHandler struct {
	UUseCase user.UseCase
}

// NewUserHandler will initialize the user endpoint
func NewUserHandler(e *echo.Echo, us user.UseCase) {
	handler := &UserHandler{
		UUseCase: us,
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

	userResp, err := u.UUseCase.FetchAll(ctx, filter, page)
	if err != nil {
		log.Println("[error][http][Index]", err)
		return c.HTML(http.StatusInternalServerError, err.Error())
	}

	type pageButton struct {
		Page int
		Url  string
	}

	// Build page buttons
	var pages []pageButton
	for i := userResp.Start; i <= userResp.End; i++ {
		pages = append(pages, pageButton{
			Page: i,
			Url:  buildUrl(req, i, filter),
		})
	}

	// Build prev/next url
	nextUrl := buildUrl(req, page+1, filter)
	prevUrl := buildUrl(req, page-1, filter)

	data := map[string]interface{}{
		"filter":   filter,
		"users":    userResp.Data,
		"nextPage": userResp.NextPage,
		"nextUrl":  nextUrl,
		"prevPage": userResp.PrevPage,
		"prevUrl":  prevUrl,
		"pages":    pages,
		"page":     page,
	}
	return c.Render(http.StatusOK, "index.html", data)
}

func buildUrl(req *http.Request, page int, filter string) string {
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
