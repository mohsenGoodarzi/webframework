package acltests
import "testing"
import "github.com/mohsenGoodarzi/webframework/tests/entities"

func TestCreateACL(t *testing.T, result *entities.ACL) { 
	got := &entities.ACL{
		EndPoints:  {'name': 'Kramer',  # string
		'age': 25          # int
	}
	}
	want := "Hello, world"
	if got != want {
	t.Errorf("got %q want %q", got, want)
	} }