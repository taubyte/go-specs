package domainSpec

import "testing"

var (
	domainName     = "taubyte"
	topLevelDomain = "com"
	rootDomain     = domainName + "." + topLevelDomain
)

func TestDomain(t *testing.T) {
	_path, err := Tns().BasicPath(rootDomain)
	if err != nil {
		t.Error(err)
		return
	}

	expectedPath := PathVariable.String() + "/" + topLevelDomain + "/" + domainName
	if _path.String() != expectedPath {
		t.Errorf("Got `%s` expected `%s`", _path, expectedPath)
		return
	}

}

func TestWebHookPass(t *testing.T) {
	test1 := "hooks.git.taubyte.com"
	test2 := "hooks.git.taubyte.com."
	test3 := "https://hooks.git.taubyte.com"
	test4 := "https://hooks.git.taubyte.com."

	if TaubyteHooksDomain.MatchString(test1) == false {
		t.Errorf("%s did not pass regex match string", test1)
		return
	}

	if TaubyteHooksDomain.MatchString(test2) == false {
		t.Errorf("%s did not pass regex match string", test2)
		return
	}

	if TaubyteHooksDomain.MatchString(test3) == false {
		t.Errorf("%s did not pass regex match string", test3)
		return
	}

	if TaubyteHooksDomain.MatchString(test4) == false {
		t.Errorf("%s did not pass regex match string", test4)
		return
	}
}
