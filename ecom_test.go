package ecomcloud_test

import (
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ecom Home Page Tests", func() {

	var cookie *http.Cookie

	BeforeEach(func() {
		cookie = &http.Cookie{Name: "unknownShopperId", Value: "123"}
	})

	It("should find the HOME upstream on requests to /", func() {
		statusCode, body, err := executeRequest("http://13.66.226.127/", cookie)
		validate(statusCode, body, err)
	})

	It("should find the HOME upstream on requests to /browse/home.do", func() {
		statusCode, body, err := executeRequest("http://13.66.226.127/browse/home.do", cookie)
		validate(statusCode, body, err)
	})

	It("should find the HOME upstream on requests to /browse/home.do/", func() {
		statusCode, body, err := executeRequest("http://13.66.226.127/browse/home.do/", cookie)
		validate(statusCode, body, err)
	})

	It("should go to USER-SEGMENTER on unknownShopperId cookie is not found", func() {
		statusCode, body, err := executeRequest("http://13.66.226.127/browse/home.do/", nil)
		Ω(err).ShouldNot(HaveOccurred())

		Ω(statusCode).Should(Equal(200))
		Ω(string(body)).Should(Equal("USER-SEGMENTER"))
	})

	It("should find the unknownShopperId cookie when other cookies are present", func() {
		client := &http.Client{}
		req := newRequest("GET", "http://13.66.226.127/")

		cookieDecoy := &http.Cookie{Name: "cookie1", Value: "123"}
		cookieDecoy2 := &http.Cookie{Name: "cookie2", Value: "123"}

		req.AddCookie(cookieDecoy)
		req.AddCookie(cookie)
		req.AddCookie(cookieDecoy2)

		res, err := client.Do(req)
		Ω(err).ShouldNot(HaveOccurred())

		body, _ := ioutil.ReadAll(res.Body)

		validate(res.StatusCode, body, err)
	})

})

func newRequest(method string, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("Error building request: %s", err.Error())
	}
	return req
}

func executeRequest(url string, cookie *http.Cookie) (int, []byte, error) {
	client := &http.Client{}
	req := newRequest("GET", url)
	if cookie != nil {
		req.AddCookie(cookie)
	}
	res, err := client.Do(req)

	if err != nil {
		return -1, nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return res.StatusCode, body, nil
}

func validate(statusCode int, body []byte, err error) {
	Ω(err).ShouldNot(HaveOccurred())
	Ω(string(body)).Should(Equal("HOME"))
	Ω(statusCode).Should(Equal(200))
}
