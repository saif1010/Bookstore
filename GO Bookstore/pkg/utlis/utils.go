package utils
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)
func ParseBody(r *http.Request, x interface{}){
	if body,err:=(ioutill.ReadAll(r.body)) err==nil{
		if err:=json.Unmarshal([]byte(body),x) err!=nil{
			return
		}
	}
}