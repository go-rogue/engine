package gui

import "testing"

func TestSettingDefaultGroup(t *testing.T) {
	rbg := NewRadioButtonGroups()
	rbg.SetDefaultGroup(1)

	if rbg.defaultGroup != 1 {
		t.Errorf("got %d, want %d", rbg.defaultGroup, 1)
	}
}

func TestUnSelectDefaultGroup(t *testing.T) {
	rbg := NewRadioButtonGroups()
	rbg.UnSelectGroup(1)
	if rbg.groupSelect[1] != nil {
		t.Errorf("not nilled")
	}
}
