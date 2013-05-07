package plugin

import (
	"cgl.tideland.biz/asserts"
	"os/exec"
	"testing"
)

func TestCommand_NoExist(t *testing.T) {
	assert := asserts.NewTestingAsserts(t, true)

	_, err := Command(exec.Command("i-should-never-ever-ever-exist"))
	assert.NotNil(err, "should have an error")
}

func TestCommand_Good(t *testing.T) {
	assert := asserts.NewTestingAsserts(t, true)

	command, err := Command(helperProcess("command"))
	assert.Nil(err, "should start command properly")

	result := command.Synopsis()
	assert.Equal(result, "1", "should return result")
}
