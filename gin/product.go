package gin

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Product struct {
	Username    string    `json:"username" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Price       int       `json:"price" binding:"gte=0"` // TODO gte=0
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type productHandler struct {
	sync.RWMutex
	products map[string]Product
}

func (handler *productHandler) Create(c *gin.Context) {
	handler.Lock()
	defer handler.Unlock()

	// 请求参数解析
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil { // 将 Body 中的 JSON 格式数据解析到指定的 Struct 中
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 参数校验
	if _, ok := handler.products[product.Name]; ok {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("product %s already existed", product.Name)})
		return
	}
	// 业务逻辑处理
	product.CreatedAt = time.Now()
	handler.products[product.Name] = product
	log.Printf("Add product %s success", product.Name)
	// 响应客户端
	c.JSON(http.StatusOK, product)
}

func (handler *productHandler) Get(c *gin.Context) {
	handler.Lock()
	defer handler.Unlock()

	productName := c.Param("name")
	product, ok := handler.products[productName]
	if !ok {
		c.JSON(http.StatusNotFound,
			gin.H{"error": fmt.Errorf("cannot found product %s", productName)}) // TODO fmt.Errorf
		return
	}
	c.JSON(http.StatusOK, product)
}

func newProductHandler() *productHandler {
	return &productHandler{
		products: make(map[string]Product),
	}
}

func router() http.Handler {
	router := gin.Default()
	productHandler := newProductHandler()

	// 路由分组
	//v1 := router.Group("/v1", gin.BasicAuth(gin.Accounts{"user": "luca"})) // 路由统一处理
	v1 := router.Group("/v1")
	{
		productv1 := v1.Group("/products")
		{
			// 路由匹配
			productv1.POST("", productHandler.Create)
			productv1.GET(":name", productHandler.Get) // :name 精准匹配
			//productv1.GET("*name", productHandler.Get) // *name 模糊不清匹配
		}

		// 路由分组2
		// ...
	}

	// 路由分组3
	// ...

	return router
}

func ProductEntry() {
	// 一进程多服务
	var eg errgroup.Group
	eg.Go(func() error {
		// 当 HTTP 服务的访问量大时, 重启服务的时可能还有很多连接没有断开, 请求没有完成
		// 优雅关停使 HTTP 服务可以在处理完所有请求后, 正常地关闭这些连接
		err := endless.ListenAndServe(":8080", router())
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})
	eg.Go(func() error {
		dir, _ := os.Getwd()
		certFile := fmt.Sprintf("%s/gin/ssl/server.pem", dir)
		keyFile := fmt.Sprintf("%s/gin/ssl/server.key", dir)
		err := endless.ListenAndServeTLS(":8443", certFile, keyFile, router())
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	// 阻塞主线程
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
