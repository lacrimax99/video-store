package usecasesnotification

import "fmt"

func mainx() {

	mag1, _ := newNotification("AWS", "Template XXX", "XXX")
	mag1.SendEmail()
	fmt.Println(mag1)
}
