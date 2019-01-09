package migration

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/sapcc/kubernikus/pkg/apis/kubernikus/v1"
	"github.com/sapcc/kubernikus/pkg/client/openstack"
	"github.com/sapcc/kubernikus/pkg/util"
)

func AddAggregationLayerCertificates(rawKluster []byte, kluster *v1.Kluster, client kubernetes.Interface, openstackFactory openstack.SharedOpenstackClientFactory) (err error) {
	apiSecret, err := client.CoreV1().Secrets(kluster.Namespace).Get(kluster.GetName(), metav1.GetOptions{})
	if err != nil {
		return err
	}

	secret, err := v1.NewSecret(apiSecret)
	if err != nil {
		return err
	}

	if secret.AggregationAggregatorCertificate != "" && secret.AggregationCACertificate != "" && secret.AggregationAggregatorPrivateKey != "" && secret.AggregationCAPrivateKey != "" {
		return nil
	}

	factory := util.NewCertificateFactory(kluster, &secret.Certificates, "cluster.local")

	if err := factory.Ensure(); err != nil {
		return err
	}

	secretData, err := secret.ToData()
	if err != nil {
		return err
	}

	apiSecret.Data = secretData
	_, err = client.CoreV1().Secrets(kluster.Namespace).Update(apiSecret)

	return err
}
