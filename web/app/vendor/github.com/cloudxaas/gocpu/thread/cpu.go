package cxcputhread

import(
	"runtime"
	
	"golang.org/x/sys/unix"
	"github.com/cloudxaas/gohash/fnv1a32"
	flag "github.com/spf13/pflag"
	
)
var (
	//Support up to 65534 cpu cores
	//because of the way other functions wraps it, 
        //"0" value denotes or master
	CPUThread uint16
)

func init() {
	//CPUThread = *flag.Uint16P("CPUThread", "t", 0, "prefork child id")
	flag.Uint16VarP(&CPUThread, "CPUThread", "t", 0, "prefork child id")
	flag.Parse()
}

//cpu core 0 = 1
//cpu core 1 = 2
func SetCPUAffinity(cpu uint16) error {
	var newMask unix.CPUSet
	newMask.Set(int(cpu)-1)
	return unix.SchedSetaffinity(0, &newMask)
}

//hash based on fnv1a32, getting hashed cpu id
func CPUHash(k []byte) uint16 {
	return uint16(fnv1a32.Hash(k) % uint32(runtime.NumCPU()))
}

func IsCurrentCPUID(id uint16) uint8 {
	if id == CPUThread - 1 {
		return 1
	}
	return 0
}
