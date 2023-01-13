package sbi

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "sbi"})
	gin.SetMode(gin.ReleaseMode)
}

// Route is the information for every URI.
type HttpRoute struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

type HttpRoutes []HttpRoute

type HttpRouteGroup struct {
	Path   string
	Routes []HttpRoute
}

type SbiServer interface {
	Serve()
	Terminate()
}

type ServerConfig struct {
	IpAddr string `json:"ipaddr"`
	Port   int    `json:"port"`
}
type httpServer struct {
	config *ServerConfig
	server *http.Server
	wg     sync.WaitGroup
}

func NewSbiServer(config *ServerConfig, groups []HttpRouteGroup) SbiServer {
	server := &httpServer{
		config: config,
	}
	server.register(groups)
	return server
}
func (s *httpServer) Serve() {
	log.Infof("Serving http server at %s:%d", s.config.IpAddr, s.config.Port)
	go func() {
		s.server.ListenAndServe()

		/* NOTE: use below block if considering https
		if s.config.Scheme == "http" {
			if err := s.server.ListenAndServe(); err != nil {
				//log.Errorf("Http server failed to listen", err)
			}
			return
		}

		if err :=s.server.ListenAndServeTLS(s.Tls.Pem, s.Tls.Key); err != nil {
			//log.Errorf("Http server failed to listen", err)
		}
		*/

		s.wg.Add(1)
	}()
	//log.Info("http server is running")
	return
}

func (s *httpServer) Terminate() {
	if s.server == nil {
		return
		//panic(errors.New("http server has not been started"))
	}
	s.server.Close()
	s.wg.Wait()
}

// create a http server, register services and their handlers
func (s *httpServer) register(groups []HttpRouteGroup) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host",
			"Token", "X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))

	addr := fmt.Sprintf("%s:%d", s.config.IpAddr, s.config.Port)

	for _, grp := range groups {
		addHttpRoutes(router, grp.Path, grp.Routes)
	}
	s.server = &http.Server{
		Addr:    addr,
		Handler: router,
	}
	return
}

func addHttpRoutes(engine *gin.Engine, groupname string, routes []HttpRoute) *gin.RouterGroup {
	group := engine.Group(groupname)

	for _, route := range routes {
		log.Infof("register %s method on %s", route.Name, route.Pattern)
		switch route.Method {
		case http.MethodGet:
			group.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			group.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			group.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}
