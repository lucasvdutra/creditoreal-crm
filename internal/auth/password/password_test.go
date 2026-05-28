package password

import "testing"

func TestHashAndVerify(t *testing.T) {
	t.Parallel()

	hash, err := Hash("senha-segura")
	if err != nil {
		t.Fatalf("hash password: %v", err)
	}
	if !Verify("senha-segura", hash) {
		t.Fatal("expected password to verify")
	}
	if Verify("senha-errada", hash) {
		t.Fatal("expected wrong password to fail")
	}
}
