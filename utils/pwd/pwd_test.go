package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$W5aumKM4MvQUbxXO6qLcu.aLR8a83E8FetuFtOHeCGUynvJ2lCbIK", "12341"))
}
