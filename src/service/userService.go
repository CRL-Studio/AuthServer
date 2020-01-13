package service

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"reflect"
	"time"

	gormdao "github.com/CRL-Studio/AuthServer/src/dao/gorm"
	userdao "github.com/CRL-Studio/AuthServer/src/dao/gorm/userDao"
	redisverificationdao "github.com/CRL-Studio/AuthServer/src/dao/redis/redisVerificationDao"
	"github.com/CRL-Studio/AuthServer/src/models"
	"github.com/CRL-Studio/AuthServer/src/utils/glossary"
	"github.com/CRL-Studio/AuthServer/src/utils/hash"
	"github.com/CRL-Studio/AuthServer/src/utils/logger"
	uuid "github.com/satori/go.uuid"
)

//CreateUser is
func CreateUser(params interface{}) (result string, outputError string) {
	tx := gormdao.DB().Begin()
	log := logger.Log()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error(r)
			outputError = "Unexpected Error"
		}
	}()

	value := reflect.ValueOf(params)

	user := &models.User{}
	user.UUID = uuid.NewV4().String()
	user.Role = &models.Role{UUID: glossary.RoleMember}
	user.Account = value.FieldByName("Account").String()
	user.Password = hash.New(value.FieldByName("Password").String())
	user.Name = value.FieldByName("Name").String()
	user.Email = value.FieldByName("Email").String()
	user.Score = 0
	user.Status = glossary.StatusEnabled
	user.Verification = false
	user.CreatedBy = value.FieldByName("Account").String()

	userdao.New(tx, user)

	VerificationKey := "Ver" + value.FieldByName("Account").String()
	VerificationCode := CreateVerificationCode()

	redisverificationdao.New(
		VerificationKey,
		VerificationCode,
	)

	tx.Commit()

	SendVerificationEmail(user.Account, user.Email, VerificationCode)

	return "Success", ""
}

//UserVerificationCheck is
func UserVerificationCheck(params interface{}) (result string, outputError string) {
	tx := gormdao.DB()
	log := logger.Log()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error(r)
			outputError = "Unexpected Error"
		}
	}()

	value := reflect.ValueOf(params)

	account := value.FieldByName("Account").String()

	redisKey := "Ver" + account
	verificationcode, _ := redisverificationdao.Get(redisKey)
	if verificationcode == "" {
		log.Error("找不到" + account + "驗證碼")
		outputError = "Can't Find Verification"
		return "", outputError
	}

	fmt.Println(tx)

	return result, ""
}

//UserInfo is
func UserInfo(params interface{}) (result map[string]interface{}, outputError string) {
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	value := reflect.ValueOf(params)

	fmt.Println(value)
	fmt.Println(tx)

	return result, ""
}

//UpdateUserInfo is
func UpdateUserInfo(params interface{}) (result map[string]interface{}, outputError string) {
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()

	value := reflect.ValueOf(params)

	fmt.Println(value)
	fmt.Println(tx)

	return result, ""
}

//UpdatePassword is
func UpdatePassword(params interface{}) (result map[string]interface{}, outputError string) {
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	value := reflect.ValueOf(params)

	fmt.Println(value)
	fmt.Println(tx)

	return result, ""
}

//ResetPassword is
func ResetPassword(params interface{}) (result map[string]interface{}, outputError string) {
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	value := reflect.ValueOf(params)

	fmt.Println(value)
	fmt.Println(tx)

	return result, ""
}

// CreateVerificationCode is
func CreateVerificationCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// SendVerificationEmail is
func SendVerificationEmail(account, email, verificationcode string) {
	log := logger.Log()
	adminAccount := os.Getenv("Gmail_Account")
	adminPassword := os.Getenv("Gmail_Password")
	gmailAuth := smtp.PlainAuth("", adminAccount, adminPassword, "smtp.gmail.com")
	to := []string{email}
	msg := []byte(
		"Subject: CRL Studio會員驗證電子郵件信箱!\r\n" +
			"From: test@example.com\r\n" +
			`Content-Type: multipart/mixed; boundary="qwertyuio"` + "\r\n" +
			"\r\n" +
			"--qwertyuio\r\n" +
			"此郵件為CRL Studio自動送出。\r\n" +
			"------------------------------\r\n" +
			"暱稱：" + account + "\r\n" +
			"輸入驗證碼，即完成電子郵件信箱的登錄。\r\n" +
			"驗證碼：" + verificationcode + "\r\n" +
			"※若5分鐘內未完成登錄，上述驗證碼將會失效，敬請留意。\r\n" +
			"※若您對本郵件沒有印象，煩請刪除本郵件。\r\n" +
			"------------------------------\r\n" +
			"※請勿直接回覆至此電子郵件信箱。\r\n" +
			"\r\n" +
			"--qwertyuio--\r\n",
	)

	err := smtp.SendMail("smtp.gmail.com:587", gmailAuth, adminAccount, to, msg)
	if err != nil {
		log.Error("會員：" + account + "驗證信寄送失敗")
	}

}
