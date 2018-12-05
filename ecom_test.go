package ecomcloud_test

import (
	"io/ioutil"
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
		statusCode, body, err := request("http://13.66.226.127/", cookie)
		validate(statusCode, body, err)
	})

	It("should find the HOME upstream on requests to /browse/home.do", func() {
		statusCode, body, err := request("http://13.66.226.127/browse/home.do", cookie)
		validate(statusCode, body, err)
	})

	It("should find the HOME upstream on requests to /browse/home.do/", func() {
		statusCode, body, err := request("http://13.66.226.127/browse/home.do/", cookie)
		validate(statusCode, body, err)
	})

	It("should match target on cookie found", func() {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://13.66.226.127/", nil)

		req.AddCookie(cookie)
		res, err := client.Do(req)
		Ω(err).ShouldNot(HaveOccurred())

		body, _ := ioutil.ReadAll(res.Body)

		Ω(res.StatusCode).Should(Equal(200))
		Ω(string(body)).Should(Equal("HOME"))
	})

})

var _ = Describe("Cookie Initialization Tests", func() {

	XIt("should match target on cookie found", func() {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://13.66.226.127/", nil)

		cookieDecoy := &http.Cookie{Name: "cookie1", Value: "123"}
		req.AddCookie(cookieDecoy)
		cookie := &http.Cookie{Name: "unknownShopperId", Value: "123"}
		req.AddCookie(cookie)
		cookieDecoy2 := &http.Cookie{Name: "cookie2", Value: "123"}
		req.AddCookie(cookieDecoy2)

		res, err := client.Do(req)
		Ω(err).ShouldNot(HaveOccurred())

		body, _ := ioutil.ReadAll(res.Body)

		Ω(res.StatusCode).Should(Equal(200))
		Ω(string(body)).Should(Equal("HOME"))
	})
})

func request(url string, cookie *http.Cookie) (int, []byte, error) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
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
