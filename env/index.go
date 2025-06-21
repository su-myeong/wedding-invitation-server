package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var UseGuestbook bool
var UseAttendance bool
var AdminPassword string
var AllowOrigin string

func init() {
	// Render 환경인지 확인 (Render는 RENDER 환경변수를 자동으로 설정함)
	if os.Getenv("RENDER") == "" {
		// Render가 아닌 로컬 개발 환경에서만 .env 파일 로드
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error: Cannot read .env file")
			panic(err.Error())
		}
	}

	// 환경 변수 설정 (Render/로컬 공통)
	UseGuestbook = os.Getenv("USE_GUESTBOOK") == "true"
	UseAttendance = os.Getenv("USE_ATTENDANCE") == "true"
	AdminPassword = os.Getenv("ADMIN_PASSWORD")
	AllowOrigin = os.Getenv("ALLOW_ORIGIN")
}