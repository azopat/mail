package main

import (
	"../common/const/mailconst"
	"bytes"
	"crypto/tls"
	b64 "encoding/base64"
	"fmt"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"html/template"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"os"
	"path"
	"time"
)

type (
	EmailParams struct {
		Subject  string
		To       string
		Template string
	}
)

func getHttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	return client
}

func sendEmail(queueMailDoc mailconst.QueueMailDoc) error {

	log.Printf("Send email for :  %s", queueMailDoc)
	log.Printf("Template file :  %s", queueMailDoc.Template)
	log.Printf(queueMailDoc.Template)

	body, err := ioutil.ReadFile(path.Join("/templates", queueMailDoc.Template+".html"))
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	tmpl, err := template.New("body").Parse(string(body))
	var doc bytes.Buffer
	if tmpl == nil {
		log.Printf("Template is null -- Error")
		return nil
	}
	err = tmpl.Execute(&doc, queueMailDoc.Params)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	to := mail.Address{"", queueMailDoc.To}

	mailjetClient := mailjet.NewMailjetClient("831374ed08d9a3f44e59e57b14cf8e5b", "0f70db0e437a9ae379766e9d316aaa59")
	var messagesInfo []mailjet.InfoMessagesV31

	// Get the invoice
	client := getHttpClient()
	log.Printf(queueMailDoc.InvoiceLink)
	if queueMailDoc.InvoiceLink != "" {
		r, err := client.Get(queueMailDoc.InvoiceLink)
		if err != nil {
			log.Printf(err.Error())
			return err
		}
		defer r.Body.Close()

		// Encode the invoice in base64
		bd, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf(err.Error())
			return err
		}
		fileB64 := b64.StdEncoding.EncodeToString(bd)

		// Create the file for writing - To override the original file by the compressed one
		filename := time.Now().Format(time.RFC850) + ".pdf"
		generatedFile := "/tmp/" + filename
		w, err := os.Create(generatedFile)
		if err != nil {
			log.Printf(err.Error())
			return err
		}
		defer w.Close()

		//n, err := io.Copy(w, r.Body)
		//if err != nil {
		//	fmt.Println(err)
		//	return err
		//}

		messagesInfo = []mailjet.InfoMessagesV31{
			mailjet.InfoMessagesV31{
				//FromEmail: queueMailDoc.FromAddress,
				//FromName:  queueMailDoc.FromName,
				From: &mailjet.RecipientV31{
					Email: queueMailDoc.FromAddress,
					Name:  queueMailDoc.FromName,
				},
				Subject:  queueMailDoc.Subject,
				TextPart: doc.String(),
				HTMLPart: doc.String(),
				//Recipients: []mailjet.Recipient{
				//	mailjet.Recipient{
				//		Email: to.String(),
				//	},
				//},
				To: &mailjet.RecipientsV31{
					mailjet.RecipientV31{
						Email: to.String(),
					},
				},
				Attachments: &mailjet.AttachmentsV31{
					mailjet.AttachmentV31{
						ContentType:   "application/pdf",
						Filename:      filename,
						Base64Content: fileB64,
					},
				},
			},
		}
	} else {
		messagesInfo = []mailjet.InfoMessagesV31{
			mailjet.InfoMessagesV31{
				//FromEmail: queueMailDoc.FromAddress,
				//FromName:  queueMailDoc.FromName,
				From: &mailjet.RecipientV31{
					Email: queueMailDoc.FromAddress,
					Name:  queueMailDoc.FromName,
				},
				Subject:  queueMailDoc.Subject,
				TextPart: doc.String(),
				HTMLPart: doc.String(),
				//Recipients: []mailjet.Recipient{
				//      mailjet.Recipient{
				//              Email: to.String(),
				//      },
				//},
				To: &mailjet.RecipientsV31{
					mailjet.RecipientV31{
						Email: to.String(),
					},
				},
			},
		}
	}

	//res, err := mailjetClient.SendMail(email)
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println("Success")
		fmt.Println(res)
	}

	return nil

}
