package ecomcloud_test

import (
	"testing"

	. "github.com/khurlbut/ecomcloud"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEcomcloud(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ecomcloud Suite")
}

var _ = BeforeSuite(func() {
	InitCloud()
	KubeCtlApply("./yaml/ecom-namespace.yaml")
	KubeCtlApply("./yaml/fakes/pages/home.yaml")
	KubeCtlApply("./yaml/fakes/services/user-segmenter.yaml")
	KubeCtlApply("./yaml/app/networking/service-home.yaml")
	KubeCtlApply("./yaml/app/networking/service-user-segmenter.yaml")
	KubeCtlApply("./yaml/app/networking/virtual-service-home.yaml")
	KubeCtlApply("./yaml/app/networking/gateway-home.yaml")
})

var _ = AfterSuite(func() {
})
