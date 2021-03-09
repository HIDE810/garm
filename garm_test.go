package garm

import "testing"

func TestArm(t *testing.T) {

	var err error

	_, err = ArmToHex("mov r0, r0")
	if err != nil {
		t.Fatal(err)
	}

	_, err = HexToArm("E1A00000")
	if err != nil {
		t.Fatal(err)
	}

	_, err = ThumbToHex("mov r8, r8")
	if err != nil {
		t.Fatal(err)
	}

	_, err = HexToThumb("46C0")
	if err != nil {
		t.Fatal(err)
	}
}
