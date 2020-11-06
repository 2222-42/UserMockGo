package notifier

import (
	"UserMockGo/config"
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model/userModel"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"strconv"
)

type sendGridClient struct {
	FromAddress string
	ApiKey      string
}

func NewActivationNotifier(config *config.NotifierConfig) infrainterface.IEmailNotifier {
	return sendGridClient{
		FromAddress: config.FromAddress,
		ApiKey:      config.SendGridApiKey,
	}
}

func (notifier sendGridClient) SendActivationEmail(user userModel.User, activation userModel.Activation, subjectStr string) error {
	from := mail.NewEmail("UserMockGo Admin", notifier.FromAddress)
	subject := "[UserMockGo]" + subjectStr
	to := mail.NewEmail("UserId: "+strconv.Itoa(int(user.ID)), string(user.Email))

	plainTextContent := "UserMockGoに登録していただきありがとうございます。\n" +
		"UserMockGoのactivatorです。\n" +
		"以下のURLよりUserの有効化を完了してください。\n"
	htmlContent := "<p>UserMockGoに登録していただきありがとうございます。</p>" +
		"<p>UserMockGoのactivatorです。</p>" +
		"<p>以下のURLよりUserの有効化を完了してください。</p>" +
		"<a href='http://localhost:8080/userModel/activate?email=" + string(user.Email) +
		"&token=" + activation.ActivationToken + "'>アカウントの有効化</a>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(notifier.ApiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return nil
}

func (notifier sendGridClient) SendCode(user userModel.User, code string) error {
	from := mail.NewEmail("UserMockGo Admin", notifier.FromAddress)
	subject := "[UserMockGo]" + "activation code"
	to := mail.NewEmail("UserId: "+strconv.Itoa(int(user.ID)), string(user.Email))

	plainTextContent := "UserMockGoにログインしていただきありがとうございます。\n" +
		"認証コードは以下の通りです。\n"
	htmlContent := "<p>UserMockGoにログインしていただきありがとうございます。</p>" +
		"<p>以認証コードは以下の通りです。</p>" +
		"<p>認証コード: " + code +
		"</p>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(notifier.ApiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return nil
}
