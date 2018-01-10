package dotplan

import (
	"errors"
	"os"
	"os/user"
	"io/ioutil"
	"path"
)

func GetUserpath(username *string) (error, *string) {
	User, err := user.Lookup(*username)
	if err != nil {
		return err, nil
	} 
	return nil, &User.HomeDir
}

func GetUserplan(username *string) (error, *[]byte) {
	err, Path := GetUserpath(username)
	if err != nil {
		return err, nil
	}

	PlanPath := path.Join(*Path, ".plan")
	
	Plan, err := os.Stat(PlanPath)
	if err != nil {
		return err, nil
	}

	if Plan.IsDir() {
		return errors.New("Not a readable file"), nil
	}

	Data, err := ioutil.ReadFile(PlanPath)
	if err != nil {
		return err, nil
	}

	return nil, &Data
}
	
