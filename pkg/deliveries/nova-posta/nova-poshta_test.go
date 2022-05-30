package nova_posta

import (
	"testing"
)

func TestTrackingDocument(t *testing.T) {
	np:=NewNovaPoshta("")
	np.TrackingDocument("","")
}
