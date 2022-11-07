package auth

import (
	"context"
	"log"
	"sync"
	"time"

	v1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func init() {
	userTokenStorage = make(map[string]string)
	go clearMap()
}

var userTokenStorage map[string]string
var mutex sync.Mutex

func clearMap() {
	for {
		time.Sleep(time.Second * 500)
		mutex.Lock()
		for k := range userTokenStorage {
			delete(userTokenStorage, k)
		}
		mutex.Unlock()
	}
}
func GetK8sToken(namespace string, username string) (tokenString string) {
	mutex.Lock()
	defer mutex.Unlock()
	if t, ok := userTokenStorage[username]; ok {

		return t
	}

	tokenString = getK8sToken(namespace, username)
	userTokenStorage[username] = tokenString
	return

}

func getK8sToken(namespace string, username string) (tokenString string) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	var expiretime int64 = 1000
	ctx := context.Background()

	token, err := clientset.CoreV1().ServiceAccounts(namespace).CreateToken(ctx, username, &v1.TokenRequest{Spec: v1.TokenRequestSpec{ExpirationSeconds: &expiretime}}, metav1.CreateOptions{})
	if err != nil {
		log.Println(err)
	} else {
		tokenString = token.Status.Token
		return
	}

	return

}
