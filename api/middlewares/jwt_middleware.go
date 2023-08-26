package middlewares

type JWTMiddleware struct{}

func NewJWTMiddleware() JWTMiddleware {
	return JWTMiddleware{}
}

func (m JWTMiddleware) Setup() {

}
