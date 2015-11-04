Struct Default
==============
With struct default,you can initial the default value of your struct with `default` tag.

## Sample
```
package structdefault

import (
	"log"
	"testing"
)

func TestConvertByteToStruct(t *testing.T) {
	type Address struct {
		Country string `json:"country" default:"asdf"`
		City    string `json:"city" default:"beijing"`
	}
	type User struct {
		Name       string  `json:"name" default:"unknow"`
		Age        int     `json:"age" default:"3"`
		Address    Address `json:"address"`
		Installed  bool    `json:"installed"    default:"true"`
		Intime     int64   `json:"intime" default:"$timeNowUnix"`
		Intimename int64   `json:"intime" default:"$timeNowUnixNano"`
		ID         string  `json:"age" default:"$uuid"`
		Random     int64   `json:"age" default:"$range(199,200)"`
		// Random     string  `json:"age" default:"$random(1,200)"`
		// Z         map[string]string `json:"z"`
	}
	var a = User{
		Address: Address{
			City: "123",
		},
	}
	Default(&a)
	log.Println(a)
	log.Println(a.ID)
	log.Println(a.Random)
	log.Println(a.Address)
	log.Println(a.Address.City)
	log.Println(a.Address.Country)

}

```