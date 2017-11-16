package differ

import (
	"runtime"
	"testing"
)

func TestNew(t *testing.T) {

}

func TestDiffer_Scan(t *testing.T) {

}

func TestDiffer_Count(t *testing.T) {

}

func TestDiffer_FileMd5(t *testing.T) {

}

func TestDiffer_Sames(t *testing.T) {

}

func TestScanner_ChunksAsCPUNumber(t *testing.T) {
	d := New("/usr/local/var/www/jh-framework/jh")
	d.Scan()
	chunks := d.ChunksAsCPUNumber()

	if len(chunks) != runtime.NumCPU() {
		t.Error("error")
	} else {
		t.Log("PASS")
	}
}
