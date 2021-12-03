package fileType

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetFileType(t *testing.T) {
	//f, err := os.Open("D:\\www\\Isheji\\365sheji-pc\\testenv.php")
	//f, err := os.Open("D:\\www\\Isheji\\365sheji-pc\\routes\\web.php")
	f, err := os.Open("D:\\www\\Isheji\\365sheji-pc\\app\\Http\\Controllers\\OrderController.php")
	//f, err := os.Open("C:\\Users\\Administrator\\Downloads\\VCG211358896239.png")
	if err != nil {
		t.Logf("open error: %v", err)
	}

	fSrc, err := ioutil.ReadAll(f)

	t.Logf("%X", fSrc[:10])
	t.Log(GetFileType(fSrc[:10]))
}
