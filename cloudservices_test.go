package ecomcloud_test

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/khurlbut/ecomcloud"
)

var _ = PDescribe("Ecom", func() {
	BeforeEach(func() {
	})

	It("should create the ecom namespace", func() {
		err := KubeCtlApply("./yaml/ecom-namespace.yaml")
		Ω(err).Should(BeNil())

		err = KubeCtlGetNamespace("ecom")
		Ω(err).Should(BeNil())

		err = KubeCtlDelete("./yaml/ecom-namespace.yaml")
		Ω(err).Should(BeNil())
	})

	Specify("that kubectl appy returns an error if yaml file not found", func() {
		err := KubeCtlApply("does_not_exist.yaml")
		Ω(err).ShouldNot(BeNil())
	})

	PSpecify("that kubectl get namespace returns an error if the namespace is not found", func() {
		err := KubeCtlGetNamespace("ecom")
		log.Println(err)
		Ω(err).ShouldNot(BeNil())
	})

})
