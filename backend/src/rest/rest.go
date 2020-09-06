package rest

import (
	"github.com/gin-gonic/gin"
)

// HandlerInterface : 클라이언트 요청을 처리하는 핸들러 / 코드 확장성 높이는 핸들러의 모든 메서드 포함하는 인터페이스
type HandlerInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

type Handler struct {
	db dblayer
}

// RunAPI 기본 라우터
func RunAPI(address string) error {
	r := gin.Default()
	r.GET("/relativepath/to/url", func(c *gin.Context) {
		// 로직 구현
	})
	r.GET("/products", func(c *gin.Context) {
		// 클라이언트에게 상품 목록 반환
	})
	r.GET("/promos", func(c *gin.Context) {
		// 클라이언트에게 프로모션 목록 반환
	})

	// 사용자 로그인 POST 요청
	r.POST("/users/signin", func(c *gin.Context) {
		// 사용자 로그인
	})
	r.POST("/users", func(c *gin.Context) {
		// 사용자 추가
	})

	// 사용자 로그아웃 POST 요청
	r.POST("/users/:id/signout", func(c *gin.Context) {
		// 해당 ID 사용자 로그아웃
	})

	// 구매 목록 조회
	r.GET("/users/:id/orders", func(c *gin.Context) {
		// 해당 ID의 사용자의 주문내역 조회
	})
	r.POST("/users/charge", func(c *gin.Context) {
		// 신용카드 결제 처리
	})

}
