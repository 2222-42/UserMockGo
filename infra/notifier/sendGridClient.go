package notifier

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model/user"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
	"strconv"
)

type sendGridClient struct {
}

func NewActivationNotifier() infrainterface.IEmailNotifier {
	return sendGridClient{}
}

func (notifier sendGridClient) SendActivationEmail(user user.User, activation user.Activation, subjectStr string) error {
	from := mail.NewEmail("UserMockGo Admin", os.Getenv("FROM_ADDRESS"))
	subject := "[UserMockGo]" + subjectStr
	to := mail.NewEmail("UserId: "+strconv.Itoa(int(user.ID)), string(user.Email))

	plainTextContent := "UserMockGoに登録していただきありがとうございます。\n" +
		"UserMockGoのactivatorです。\n" +
		"以下のURLよりUserの有効化を完了してください。\n"
	htmlContent := "<p>UserMockGoに登録していただきありがとうございます。</p>" +
		"<p>UserMockGoのactivatorです。</p>" +
		"<p>以下のURLよりUserの有効化を完了してください。</p>" +
		"<a href='http://localhost:8080/user/activate?email=" + string(user.Email) +
		"&token=" + activation.ActivationToken + "'>アカウントの有効化</a>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return nil
}

func (notifier sendGridClient) SendCode(user user.User, code string) error {

	return nil
}
