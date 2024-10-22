package tools

import (
	"fmt"
	"os"

	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

func UploadToHuaweiCloud(ak string, sk string, endPoint string, bucketName string, objectName string, filepath string) error {
	//Obtain an AK/SK pair using environment variables or import an AK/SK pair in other ways. Using hard coding may result in leakage.
	//Obtain an AK/SK pair on the management console. For details, see https://support.huaweicloud.com/intl/en-us/usermanual-ca/ca_01_0003.html.

	// (Optional) If you use a temporary AK/SK pair and a security token to access OBS, you are advised not to use hard coding to reduce leakage risks. You can obtain an AK/SK pair using environment variables or import an AK/SK pair in other ways.
	// securityToken := os.Getenv("SecurityToken")
	// Enter the endpoint corresponding to the bucket. CN-Hong Kong is used here as an example. Replace it with the one currently in use.

	// Create an obsClient instance.
	// If you use a temporary AK/SK pair and a security token to access OBS, use the obs.WithSecurityToken method to specify a security token when creating an instance.
	obsClient, err := obs.New(ak, sk, endPoint /*, obs.WithSecurityToken(securityToken)*/)
	if err != nil {
		fmt.Printf("Create obsClient error, errMsg: %s", err.Error())
		return err
	}
	input := &obs.PutObjectInput{}
	// Specify a bucket name.
	input.Bucket = bucketName
	// Specify the object (example/objectname as an example) to upload.
	input.Key = objectName
	fd, _ := os.Open(filepath)
	input.Body = fd
	// Upload you local file using streaming.
	output, err := obsClient.PutObject(input)
	if err == nil {
		fmt.Printf("Put object(%s) under the bucket(%s) successful!\n", input.Key, input.Bucket)
		fmt.Printf("StorageClass:%s, ETag:%s\n",
			output.StorageClass, output.ETag)
		return nil
	}
	fmt.Printf("Put object(%s) under the bucket(%s) fail!\n", input.Key, input.Bucket)
	if obsError, ok := err.(obs.ObsError); ok {
		fmt.Println("An ObsError was found, which means your request sent to OBS was rejected with an error response.")
		fmt.Println(obsError.Error())
	} else {
		fmt.Println("An Exception was found, which means the client encountered an internal problem when attempting to communicate with OBS, for example, the client was unable to access the network.")
		fmt.Println(err)
	}
	return err
}
