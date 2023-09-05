package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/netip"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	cxcputhread "github.com/cloudxaas/gocpu/thread"
	cx "github.com/cloudxaas/gocx"
	cxnetip "github.com/cloudxaas/gonet/ip"
	cxstrconv "github.com/cloudxaas/gostrconv"
	cxsysinfo "github.com/cloudxaas/gosysinfo"
	"github.com/hertz-contrib/http2/config"
	"github.com/hertz-contrib/http2/factory"
	sprapp "github.com/cloudxaas/gosprapp/app/sprapp"
	hertzpprof "github.com/hertz-contrib/pprof"
	"github.com/panjf2000/ants/v2"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sys/unix"
)

// TODO, we need to make black list and white list use mmapdb, allow ip list need to be updated with netip.Prefix
// IPStatusLoginedDatabankCache.Set not found?!

func main() {
	//childID := flag.Int("child-id", -1, "prefork child id")

	// With prefork we can only do tls-alpn-01 as for the http-01 challenges we can't
	// guarantee that the http request is received by the same process that stared the challenge.

	// TODO: read certData and keyData from somewhere.
	// For now we'll generate a test certificate to test this code.
	var err error

	if cxcputhread.CPUThread == 0 {


		if frwllcentral.DisableHTTPS == 1 {
			childs := uint16(runtime.GOMAXPROCS(-1)) // Start a child for each CPU.
			ids := make(chan uint16, childs)

			var i uint16
			for i = 0; i < childs; i++ {
				ids <- i + 1
			}

			for id := range ids {
				idP := id
				ants.Submit(func() {
					cmd := exec.Command(os.Args[0], append(os.Args[1:], "-t="+cxstrconv.Uint16toa(idP))...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr

					log.Printf("id: %d", idP)
					err := cmd.Run()
					//time.Sleep(1*time.Second)
					log.Printf("%s %s child %d exited with error: %v\n", os.Args[0], os.Args[1:], idP, err)
					//time.Sleep(100*time.Second)
					ids <- idP // When a child dies we just restart it.
				})

			}
		}

	}

	runtime.GOMAXPROCS(1)
	err = cxcputhread.SetCPUAffinity(cxcputhread.CPUThread)
	if err != nil {
		panic(err)
	}
	//^sp 0.0.1




	//splog.Display(splog.Debug, frwllcentral.LogLevel, "ipStatuXXXXX")
	//ants.Submit(func() {
		hnormal := server.New(
			server.WithHostPorts(":80"),

			server.WithListenConfig(&net.ListenConfig{
				Control: func(network, address string, c syscall.RawConn) error {

					return c.Control(func(fd uintptr) {
						syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
					})

				},
			}),
		)


		hnormal.Use(func(cc context.Context, ctx *app.RequestContext) {
			//log.Printf("kjhkjh323 = %s = %s", cc, ctx)
			HertzHandler(&cc, ctx)
		})

		hnormal.Spin()
	//})

}

func HertzHandler(cc *context.Context, ctx *app.RequestContext) {

	sprapp.MainHTTP()

}
