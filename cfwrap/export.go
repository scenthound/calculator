package cfwrap

import ( 
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

func GetExport(region string, exportName string) (string, error) {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region)},
	))

	cf := cloudformation.New(sess)

	resp, err := cf.ListExports(&cloudformation.ListExportsInput{})

	if err != nil {
		return "", err
	}


	// A better implementation would provide pagination as well
	for _, export := range resp.Exports {
		if aws.StringValue(export.Name) == exportName {
			return aws.StringValue(export.Value), nil
		}
	}

	return "", nil
}
