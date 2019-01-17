Bytes!
## 0.vgimg: consider migrating backend to github.com/fogleman/gg
Sep 17, 2018
```text
 Over the years we've had numerous problems because of API changes and bugs in github.com/llgcode/draw2d (some of which still exist). I'd like to suggest we have a look at github.com/fogleman/gg as a replacement.
```
## 1.plotutil 3rd dashed line only partially renders at certain resolutions
Nov 23, 2015
```text
 I found this result rather confusing a few weeks back thinking I had incomplete data until I discovered it was the plot itself. Here's a sample program and plot of the result
package main

import (
    "log"
    "math"

    "github.com/gonum/plot"
    "github.com/gonum/plot/plotter"
    "github.com/gonum/plot/plotutil"
)

func xyer(out []float64) plotter.XYs {
    n := len(out)
    xys := make(plotter.XYs, n)
    for i, v := range out {
        xys[i].X = float64(i) / float64(n)
        xys[i].Y = v
    }
    return xys
}

func decay() []float64 {
    sig := make([]float64, 1024)
    n := float64(len(sig))
    for i := 0.0; i < n; i++ {
        sig[int(i)] = math.Exp(2 * math.Pi * -(i / n))
    }
    return sig
}

func main() {
    plt, err := plot.New()
    if err != nil {
        log.Fatal(err)
    }
    plt.X.Min, plt.X.Max = 0, 1
    plt.Y.Min, plt.Y.Max = -1, 1

    empty := make([]float64, 1024)
    plotutil.AddLines(plt,
        "empty", xyer(empty),
        "empty", xyer(empty), // comment out and decay plots ok
        "decay", xyer(decay()), // this dashed line is bugged
    )

    plt.Save(1000, 500, "out.png")
    // Increase size and decay plots ok
    // plt.Save(2000, 1000, "out.png")
}
```
```text
 The same fault happens when you save to jpg and tiff, but not to svg. I suspect this is a github.com/llgcode/draw2d bug.
```
```text
 interesting library. I'll see about reproducing specific to draw2d and if it's an issue that can actually be addressed there. Though to some degree, I might expect a drawing library to drop data too small to render in a given frame where as I might expect a plotter to compensate to produce something useful.
```
```text
 A way to approach this might be to capture what it being drawn with recorder.Canvas and replaying sections of the captured actions to figure out what is going on. I may be able to get to this today.
```
```text
 I'm certainly in no rush and only discovered today that an increase in
resolution resolves the issue. It'd likely be 7+ days before I could look
into it and being rather unfamiliar with everything I'd probably not be
very productive :)
On Mon, Nov 23, 2015 at 4:52 PM Dan Kortschak notifications@github.com
wrote:

A way to approach this might be to capture what it being drawn with
recorder.Canvas
https://godoc.org/github.com/gonum/plot/vg/recorder#Canvas and
replaying sections of the captured actions to figure out what is going on.
I may be able to get to this today.
—
Reply to this email directly or view it on GitHub
#231 (comment).
```
```text
 Here is a simpler reproducer; I'm not sure the recorder will help here.
package main

import (
    "log"
    "math"

    "github.com/gonum/plot"
    "github.com/gonum/plot/plotter"
    "github.com/gonum/plot/vg"
)

func decay() plotter.XYs {
    sig := make(plotter.XYs, 1024)
    n := float64(len(sig))
    for i := range sig {
        sig[i].X = float64(i) / n
        sig[i].Y = math.Exp(2 * math.Pi * -(float64(i) / n))
    }
    return sig
}

func main() {
    plt, err := plot.New()
    if err != nil {
        log.Fatal(err)
    }
    plt.X.Min, plt.X.Max = 0, 1
    plt.Y.Min, plt.Y.Max = -1, 1

    l, err := plotter.NewLine(decay())
    if err != nil {
        log.Fatalf("failed to make new line: %v", err)
    }
    l.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}

    plt.Add(l)

    plt.HideX()
    plt.HideY()

    plt.Save(1000, 500, "out.png")
}

If the dash/space length is 2 or 3 you get the effect (2 causes continuous line and 3 causes blank). Above that it is always a continuous line. The point at which break happens is not consistent between the 2 and 3 cases. Interestingly if you have replicas of the lines, they all break at the same part of the curve. The error appears to be in draw2dbase.(*DashVertexConverter).lineTo where alternation between gap and line is calculated, though it's not obvious to me how. \cc @llgcode.
```
```text
 The issue is that the condition here is preventing saving the correct state of the dash for the next call to lineTo. So when the distance moved in a call to lineTo is too small compared to a dash segment, no progression through the dash space is possible. This results in either a blank or a solid line (why all the higher value dash segments above result in solid lines).
This fixes it AFAICS:
func (dasher *DashVertexConverter) lineTo(x, y float64) {
    rest := dasher.dash[dasher.currentDash] - dasher.distance
    for rest < 0 {
        dasher.distance -= dasher.dash[dasher.currentDash]
        dasher.nextDash()
        rest = dasher.dash[dasher.currentDash] - dasher.distance
    }
    dasher.distance += math.Hypot(dasher.x-x, dasher.y-y)
    for dasher.distance >= rest {
        k := rest / dasher.distance
        dasher.x += k * (x - dasher.x)
        dasher.y += k * (y - dasher.y)
        dasher.dashTo(dasher.x, dasher.y)
        dasher.distance -= rest
        dasher.nextDash()
        rest = dasher.dash[dasher.currentDash]
    }

    dasher.dashTo(x, y)
    if dasher.distance >= rest {
        dasher.distance -= rest
        dasher.nextDash()
    }
    dasher.x, dasher.y = x, y
}

func (dasher *DashVertexConverter) nextDash() {
    dasher.currentDash++
    dasher.currentDash %= len(dasher.dash)
}

func (dasher *DashVertexConverter) dashTo(x, y float64) {
    if dasher.currentDash%2 == 0 {
        // line
        dasher.next.LineTo(x, y)
    } else {
        // gap
        dasher.next.End()
        dasher.next.MoveTo(x, y)
    }
}

This is still not correct (the dash lengths look odd in places - actually off by a factor of two when compared to the svg output which is clearly correct: use dashes []vg.Length{72, 72} to see that, outputing to png and svg), but it does retain state and alternate.
```
```text
 Thanks for tracking this down. I think we should open a bug with draw2d.
```
## 2.Out of fixed registers error during installation
Aug 29, 2016
```text
 When I run go get github.com/gonum/plot/ I got this error:
# bitbucket.org/zombiezen/gopdf/pdf
run compiler with -v for register allocation sites
.go/src/bitbucket.org/zombiezen/gopdf/pdf/image.go:89: internal compiler error: out of fixed registers
# rsc.io/pdf
run compiler with -v for register allocation sites
.go/src/rsc.io/pdf/page.go:250: internal compiler error: out of fixed registers
```
```text
 interesting...
could you tell us the output of:
$> go env
$> go version
```
```text
 go env:
GOBIN=""
GOEXE=""
GOHOSTARCH="386"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/kir/.go"
GORACE=""
GOROOT="/usr/lib/go"
GOTOOLDIR="/usr/lib/go/pkg/tool/linux_386"
GO15VENDOREXPERIMENT=""
CC="gcc"
GOGCCFLAGS="-fPIC -m32 -pthread -fmessage-length=0"
CXX="g++"
CGO_ENABLED="1"```
go version:

go version:
go1.5.1 linux/386
```
```text
 golang/go#12840
```
```text
 I don't think this is us (it says it's a compiler error) - and we are not even the package that is causing it (it's happening in zombiezen and rsc's pdf package). Finally, it's not a current go version. I would try to replicate this with 1.7 and then if it persists, file an issue at golang.org/issues.
```
```text
 go version:
go version go1.7 linux/386

go get github.com/gonum/plot/:
# bitbucket.org/zombiezen/gopdf/pdf
run compiler with -v for register allocation sites
.go/src/bitbucket.org/zombiezen/gopdf/pdf/image.go:89: internal compiler error: out of fixed registers

goroutine 1 [running]:
runtime/debug.Stack(0x0, 0x0, 0x0)
    /usr/local/go/src/runtime/debug/stack.go:24 +0x80
cmd/compile/internal/gc.Fatalf(0x8570a2a, 0x16, 0x0, 0x0, 0x0)
    /usr/local/go/src/cmd/compile/internal/gc/subr.go:165 +0x216
cmd/compile/internal/gc.Regalloc(0x193b9e00, 0x18a12b80, 0x193b9d40)
    /usr/local/go/src/cmd/compile/internal/gc/gsubr.go:749 +0x3a1
cmd/compile/internal/x86.ginscmp(0x8558b3e, 0x18a12b80, 0x193b9c80, 0x193b9d40, 0x1, 0x1)
    /usr/local/go/src/cmd/compile/internal/x86/gsubr.go:657 +0x3a4
cmd/compile/internal/gc.Agenr(0x18cc8900, 0x193b9800, 0x193b9680)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:1245 +0x2489
cmd/compile/internal/gc.Igen(0x18cc8900, 0x193b9800, 0x193b9680)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:1718 +0x25a
cmd/compile/internal/gc.cgen_wb(0x18cc8900, 0x193b9680, 0x0)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:523 +0x3ad4
cmd/compile/internal/gc.Cgen(0x18cc8900, 0x193b9680)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.Mgen(0x18cc8900, 0x193b9680, 0x193b95c0)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:906 +0xd9
cmd/compile/internal/gc.cgen_wb(0x18cc8a20, 0x193b95c0, 0x3000)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:460 +0x3318
cmd/compile/internal/gc.Cgen(0x18cc8a20, 0x193b95c0)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.Mgen(0x18cc8a20, 0x193b95c0, 0x193b94a0)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:906 +0xd9
cmd/compile/internal/gc.cgen_wb(0x18cc8c60, 0x193b94a0, 0x193b9400)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:714 +0x2067
cmd/compile/internal/gc.Cgen(0x18cc8c60, 0x193b94a0)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.cgen_wb(0x18cc8d20, 0x193b9320, 0x0)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:664 +0x4257
cmd/compile/internal/gc.Cgen(0x18cc8d20, 0x193b9320)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.Mgen(0x18cc8d20, 0x193b9320, 0x193b92c0)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:906 +0xd9
cmd/compile/internal/gc.cgen_wb(0x18cc8e40, 0x193b92c0, 0x193b9200)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:460 +0x3318
cmd/compile/internal/gc.Cgen(0x18cc8e40, 0x193b92c0)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.cgen_wb(0x18cc8e40, 0x1939b080, 0x19070e00)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:102 +0x5ad
cmd/compile/internal/gc.Cgen_as_wb(0x1939b080, 0x18cc8e40, 0x200)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:1023 +0x138
cmd/compile/internal/gc.Cgen_as(0x1939b080, 0x18cc8e40)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:979 +0x2e
cmd/compile/internal/gc.gen(0x18cc9020)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:926 +0x1be
cmd/compile/internal/gc.Genlist(0x19384f90)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x18cd1e60)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:875 +0xb89
cmd/compile/internal/gc.Genlist(0x19384fb0)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x18cd12c0)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:860 +0x977
cmd/compile/internal/gc.Genlist(0x19384fd0)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.genlegacy(0x18ad0180, 0x1938f8c0, 0x1938f940)
    /usr/local/go/src/cmd/compile/internal/gc/pgen.go:498 +0x56
cmd/compile/internal/gc.compile(0x18d10060)
    /usr/local/go/src/cmd/compile/internal/gc/pgen.go:485 +0xfcd
cmd/compile/internal/gc.funccompile(0x18d10060)
    /usr/local/go/src/cmd/compile/internal/gc/dcl.go:1287 +0x167
cmd/compile/internal/gc.Main()
    /usr/local/go/src/cmd/compile/internal/gc/main.go:467 +0x1bf4
cmd/compile/internal/x86.Main()
    /usr/local/go/src/cmd/compile/internal/x86/galign.go:80 +0x2dc
main.main()
    /usr/local/go/src/cmd/compile/main.go:31 +0x117

# rsc.io/pdf
run compiler with -v for register allocation sites
.go/src/rsc.io/pdf/page.go:250: internal compiler error: out of fixed registers

goroutine 1 [running]:
runtime/debug.Stack(0x0, 0x0, 0x0)
    /usr/local/go/src/runtime/debug/stack.go:24 +0x80
cmd/compile/internal/gc.Fatalf(0x8570a2a, 0x16, 0x0, 0x0, 0x0)
    /usr/local/go/src/cmd/compile/internal/gc/subr.go:165 +0x216
cmd/compile/internal/gc.Regalloc(0x18d6e900, 0x18a12b80, 0x18d6e7e0)
    /usr/local/go/src/cmd/compile/internal/gc/gsubr.go:749 +0x3a1
cmd/compile/internal/x86.ginscmp(0x8558b3e, 0x18a12b80, 0x18d6e6c0, 0x18d6e7e0, 0x1, 0x18aeb374)
    /usr/local/go/src/cmd/compile/internal/x86/gsubr.go:657 +0x3a4
cmd/compile/internal/gc.Agenr(0x190190e0, 0x18d74fc0, 0x18d74480)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:1245 +0x2489
cmd/compile/internal/gc.Igen(0x190190e0, 0x18d74fc0, 0x18d74480)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:1718 +0x25a
cmd/compile/internal/gc.cgen_wb(0x190190e0, 0x18d74480, 0x18d74400)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:523 +0x3ad4
cmd/compile/internal/gc.Cgen(0x190190e0, 0x18d74480)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.cgen_wb(0x190193e0, 0x18d7a2a0, 0x18d7a200)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:741 +0x254e
cmd/compile/internal/gc.Cgen(0x190193e0, 0x18d7a2a0)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.cgen_wb(0x18e3fe60, 0x18d7d500, 0x18d7d500)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:725 +0x225c
cmd/compile/internal/gc.Cgen(0x18e3fe60, 0x18d7d500)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.cgen_wb(0x18e3fe60, 0x18e3f860, 0x8560300)
    /usr/local/go/src/cmd/compile/internal/gc/cgen.go:102 +0x5ad
cmd/compile/internal/gc.Cgen_as_wb(0x18e3f860, 0x18e3fe60, 0x1300)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:1023 +0x138
cmd/compile/internal/gc.Cgen_as(0x18e3f860, 0x18e3fe60)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:979 +0x2e
cmd/compile/internal/gc.gen(0x19019440)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:926 +0x1be
cmd/compile/internal/gc.Genlist(0x1985c4b0)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19018a80)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:875 +0xb89
cmd/compile/internal/gc.Genlist(0x1985c4e0)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x190186c0)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:875 +0xb89
cmd/compile/internal/gc.Genlist(0x1985c570)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19018480)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:875 +0xb89
cmd/compile/internal/gc.Genlist(0x1985c800)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19015f80)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:860 +0x977
cmd/compile/internal/gc.Genlist(0x1985c5b0)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19015b60)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:875 +0xb89
cmd/compile/internal/gc.Genlist(0x1985c6e0)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19015620)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:860 +0x977
cmd/compile/internal/gc.Genlist(0x1985c5e0)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x190153e0)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:860 +0x977
cmd/compile/internal/gc.Genlist(0x1985c620)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19014ea0)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:860 +0x977
cmd/compile/internal/gc.Genlist(0x1985c630)
    /usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.genlegacy(0x18ad2180, 0x19612bc0, 0x19612c40)
    /usr/local/go/src/cmd/compile/internal/gc/pgen.go:498 +0x56
cmd/compile/internal/gc.compile(0x19014660)
    /usr/local/go/src/cmd/compile/internal/gc/pgen.go:485 +0xfcd
cmd/compile/internal/gc.funccompile(0x19014660)
    /usr/local/go/src/cmd/compile/internal/gc/dcl.go:1287 +0x167
cmd/compile/internal/gc.Main()
    /usr/local/go/src/cmd/compile/internal/gc/main.go:467 +0x1bf4
cmd/compile/internal/x86.Main()
    /usr/local/go/src/cmd/compile/internal/x86/galign.go:80 +0x2dc
main.main()
    /usr/local/go/src/cmd/compile/main.go:31 +0x117
```
```text
 If I interpret it correctly, the golang issue I linked to above says that
the compiler bug on 386 won't be fixed until go 1.8.
On Aug 29, 2016 8:52 PM, "Kirill" notifications@github.com wrote:

go version:
go version go1.7 linux/386
go get github.com/gonum/plot/:
bitbucket.org/zombiezen/gopdf/pdf
run compiler with -v for register allocation sites
.go/src/bitbucket.org/zombiezen/gopdf/pdf/image.go:89: internal compiler error: out of fixed registers
goroutine 1 [running]:
runtime/debug.Stack(0x0, 0x0, 0x0)
/usr/local/go/src/runtime/debug/stack.go:24 +0x80
cmd/compile/internal/gc.Fatalf(0x8570a2a, 0x16, 0x0, 0x0, 0x0)
/usr/local/go/src/cmd/compile/internal/gc/subr.go:165 +0x216
cmd/compile/internal/gc.Regalloc(0x193b9e00, 0x18a12b80, 0x193b9d40)
/usr/local/go/src/cmd/compile/internal/gc/gsubr.go:749 +0x3a1
cmd/compile/internal/x86.ginscmp(0x8558b3e, 0x18a12b80, 0x193b9c80, 0x193b9d40, 0x1, 0x1)
/usr/local/go/src/cmd/compile/internal/x86/gsubr.go:657 +0x3a4
cmd/compile/internal/gc.Agenr(0x18cc8900, 0x193b9800, 0x193b9680)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:1245 +0x2489
cmd/compile/internal/gc.Igen(0x18cc8900, 0x193b9800, 0x193b9680)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:1718 +0x25a
cmd/compile/internal/gc.cgen_wb(0x18cc8900, 0x193b9680, 0x0)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:523 +0x3ad4
cmd/compile/internal/gc.Cgen(0x18cc8900, 0x193b9680)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.Mgen(0x18cc8900, 0x193b9680, 0x193b95c0)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:906 +0xd9
cmd/compile/internal/gc.cgen_wb(0x18cc8a20, 0x193b95c0, 0x3000)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:460 +0x3318
cmd/compile/internal/gc.Cgen(0x18cc8a20, 0x193b95c0)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.Mgen(0x18cc8a20, 0x193b95c0, 0x193b94a0)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:906 +0xd9
cmd/compile/internal/gc.cgen_wb(0x18cc8c60, 0x193b94a0, 0x193b9400)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:714 +0x2067
cmd/compile/internal/gc.Cgen(0x18cc8c60, 0x193b94a0)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.cgen_wb(0x18cc8d20, 0x193b9320, 0x0)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:664 +0x4257
cmd/compile/internal/gc.Cgen(0x18cc8d20, 0x193b9320)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.Mgen(0x18cc8d20, 0x193b9320, 0x193b92c0)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:906 +0xd9
cmd/compile/internal/gc.cgen_wb(0x18cc8e40, 0x193b92c0, 0x193b9200)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:460 +0x3318
cmd/compile/internal/gc.Cgen(0x18cc8e40, 0x193b92c0)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.cgen_wb(0x18cc8e40, 0x1939b080, 0x19070e00)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:102 +0x5ad
cmd/compile/internal/gc.Cgen_as_wb(0x1939b080, 0x18cc8e40, 0x200)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:1023 +0x138
cmd/compile/internal/gc.Cgen_as(0x1939b080, 0x18cc8e40)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:979 +0x2e
cmd/compile/internal/gc.gen(0x18cc9020)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:926 +0x1be
cmd/compile/internal/gc.Genlist(0x19384f90)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x18cd1e60)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:875 +0xb89
cmd/compile/internal/gc.Genlist(0x19384fb0)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x18cd12c0)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:860 +0x977
cmd/compile/internal/gc.Genlist(0x19384fd0)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.genlegacy(0x18ad0180, 0x1938f8c0, 0x1938f940)
/usr/local/go/src/cmd/compile/internal/gc/pgen.go:498 +0x56
cmd/compile/internal/gc.compile(0x18d10060)
/usr/local/go/src/cmd/compile/internal/gc/pgen.go:485 +0xfcd
cmd/compile/internal/gc.funccompile(0x18d10060)
/usr/local/go/src/cmd/compile/internal/gc/dcl.go:1287 +0x167
cmd/compile/internal/gc.Main()
/usr/local/go/src/cmd/compile/internal/gc/main.go:467 +0x1bf4
cmd/compile/internal/x86.Main()
/usr/local/go/src/cmd/compile/internal/x86/galign.go:80 +0x2dc
main.main()
/usr/local/go/src/cmd/compile/main.go:31 +0x117
rsc.io/pdf
run compiler with -v for register allocation sites
.go/src/rsc.io/pdf/page.go:250: internal compiler error: out of fixed registers
goroutine 1 [running]:
runtime/debug.Stack(0x0, 0x0, 0x0)
/usr/local/go/src/runtime/debug/stack.go:24 +0x80
cmd/compile/internal/gc.Fatalf(0x8570a2a, 0x16, 0x0, 0x0, 0x0)
/usr/local/go/src/cmd/compile/internal/gc/subr.go:165 +0x216
cmd/compile/internal/gc.Regalloc(0x18d6e900, 0x18a12b80, 0x18d6e7e0)
/usr/local/go/src/cmd/compile/internal/gc/gsubr.go:749 +0x3a1
cmd/compile/internal/x86.ginscmp(0x8558b3e, 0x18a12b80, 0x18d6e6c0, 0x18d6e7e0, 0x1, 0x18aeb374)
/usr/local/go/src/cmd/compile/internal/x86/gsubr.go:657 +0x3a4
cmd/compile/internal/gc.Agenr(0x190190e0, 0x18d74fc0, 0x18d74480)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:1245 +0x2489
cmd/compile/internal/gc.Igen(0x190190e0, 0x18d74fc0, 0x18d74480)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:1718 +0x25a
cmd/compile/internal/gc.cgen_wb(0x190190e0, 0x18d74480, 0x18d74400)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:523 +0x3ad4
cmd/compile/internal/gc.Cgen(0x190190e0, 0x18d74480)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.cgen_wb(0x190193e0, 0x18d7a2a0, 0x18d7a200)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:741 +0x254e
cmd/compile/internal/gc.Cgen(0x190193e0, 0x18d7a2a0)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.cgen_wb(0x18e3fe60, 0x18d7d500, 0x18d7d500)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:725 +0x225c
cmd/compile/internal/gc.Cgen(0x18e3fe60, 0x18d7d500)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:19 +0x2e
cmd/compile/internal/gc.cgen_wb(0x18e3fe60, 0x18e3f860, 0x8560300)
/usr/local/go/src/cmd/compile/internal/gc/cgen.go:102 +0x5ad
cmd/compile/internal/gc.Cgen_as_wb(0x18e3f860, 0x18e3fe60, 0x1300)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:1023 +0x138
cmd/compile/internal/gc.Cgen_as(0x18e3f860, 0x18e3fe60)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:979 +0x2e
cmd/compile/internal/gc.gen(0x19019440)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:926 +0x1be
cmd/compile/internal/gc.Genlist(0x1985c4b0)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19018a80)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:875 +0xb89
cmd/compile/internal/gc.Genlist(0x1985c4e0)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x190186c0)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:875 +0xb89
cmd/compile/internal/gc.Genlist(0x1985c570)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19018480)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:875 +0xb89
cmd/compile/internal/gc.Genlist(0x1985c800)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19015f80)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:860 +0x977
cmd/compile/internal/gc.Genlist(0x1985c5b0)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19015b60)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:875 +0xb89
cmd/compile/internal/gc.Genlist(0x1985c6e0)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19015620)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:860 +0x977
cmd/compile/internal/gc.Genlist(0x1985c5e0)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x190153e0)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:860 +0x977
cmd/compile/internal/gc.Genlist(0x1985c620)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.gen(0x19014ea0)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:860 +0x977
cmd/compile/internal/gc.Genlist(0x1985c630)
/usr/local/go/src/cmd/compile/internal/gc/gen.go:311 +0x6d
cmd/compile/internal/gc.genlegacy(0x18ad2180, 0x19612bc0, 0x19612c40)
/usr/local/go/src/cmd/compile/internal/gc/pgen.go:498 +0x56
cmd/compile/internal/gc.compile(0x19014660)
/usr/local/go/src/cmd/compile/internal/gc/pgen.go:485 +0xfcd
cmd/compile/internal/gc.funccompile(0x19014660)
/usr/local/go/src/cmd/compile/internal/gc/dcl.go:1287 +0x167
cmd/compile/internal/gc.Main()
/usr/local/go/src/cmd/compile/internal/gc/main.go:467 +0x1bf4
cmd/compile/internal/x86.Main()
/usr/local/go/src/cmd/compile/internal/x86/galign.go:80 +0x2dc
main.main()
/usr/local/go/src/cmd/compile/main.go:31 +0x117
—
You are receiving this because you commented.
Reply to this email directly, view it on GitHub
#311 (comment), or mute
the thread
https://github.com/notifications/unsubscribe-auth/ABwJvpkSFV_GV1U1USfcemLad46eyEwzks5qksfxgaJpZM4JvReu
.
```
## 3.plot: The aspect ratio can't be specified
Apr 12, 2018
```text
 What are you trying to do?
I would like to be able to specify the aspect ratio of the plots, for instance to make sure geometric shapes look right and make sure the 1:1 line in scatter plots is at 45 degrees: For example: https://matplotlib.org/gallery/subplots_axes_and_figures/axis_equal_demo.html
I think, in terms of API, this could be accomplished by adding an Aspect field to plot.Plot. With the default 0 value, the aspect ratio would be determined based on the size of the canvas as normal, but if it is set to a positive number (e.g., 1) then the aspect ratio would be forced to that number.
What version of Go and Gonum/plot are you using?

go 1.10
07f4425
Does this issue reproduce with the current master?
yes
```
```text
 I would love this very much, too. I've recently hacked on a simple Go package which allows to draw ellipse only to realize there is no way to specify the aspect ratio, so the ellipses look like circles :-/
```
## 4.plot: fix locations of default ticks
Oct 30, 2017
```text
 This is a continuation of #393.
I'd like to look into this implementation, but in the mean time, change the location of the tick marks.
```
```text
 adding the linked file here as well: I was only able to look at it / download it from my phone but it 404-ed from chromium on my desktop...
2010-TickLabels-InfoVis.pdf
```
```text
 Here is a Go implementation (from the paper and the R code in the labelling package - MIT + file LICENSE | Unlimited according to the R documentation) with modifications to allow capturing the magnitude of the values and demonstrating how this is used.
Please have a play with values of dMin, dMax and m.
I think with this and a log-space equivalent we can resolve the tick labelling situation into a single type again.
```
## 5.vg: consider migrating vg.Length and vg.Point to x/image/math/fixed
Apr 12, 2016
```text
 The package golang.org/x/image/math/fixed defines fixed point integer types such as fixed.Int26_6 and fixed.Int52_12 (and corresponding Point and Rectangle struct types).
migrating vg.Length from:
package vg
type Length float64
to:
package vg
import "golang.org/x/image/math/fixed"
type Length fixed.Int26_6
would (perhaps) ease interoperabiliy with all the other 2D packages.
```
```text
 I don't think we should do this. The image/math/fixed package is there for rasterizing images. Rasterization uses integer-addressed pixels, and fixed-point values make this easier. Vg deals with vector graphics, which work in floating point space. Sure, rounding may happen in the way-way-back end, but all the libraries called by the Canvases provide work in floats. Using fixed-point types would make our lives more difficult.
```
```text
 fair enough.
what about shiny/unit.Unit then?
https://godoc.org/golang.org/x/exp/shiny/unit
(yes, shiny is still experimental...)
```
```text
 I may be OK with that, but I wouldn't be in a huge rush to do it. I'm not bothered because shiny's experimental; it's made by "core go developers" and sits in a github.com/golang repo, so it's guaranteed to be viewed as the standard UI package for Go. The thing that bothers me is that a package like 'unit' isn't shiny specific; it's more generally applicable, for example, gonum/plot could very well use it.  I would prefer to wait to see if they move it out of the shiny directory. It seems strange to create a plot using a seemingly-shiny-specific unit package.
```
## 6.plotter: HeatMap panics if h.Min = -inf
Mar 22, 2018
```text
 While calling plot.Save
gonum.org/v1/plot/plotter.(*HeatMap).Plot(0xc42006a360, 0x13d8220, 0xc42052c420, 0x403e70933333332f, 0x4037d0c666666662, 0x4071f23ecccccccd, 0x4071f3fb9999999a, 0xc42042e000)
	/Users/brendan/Documents/mygo/src/gonum.org/v1/plot/plotter/heat.go:171 +0x1268
gonum.org/v1/plot.(*Plot).Draw(0xc42042e000, 0x13d8120, 0xc420088d80, 0x0, 0x0, 0x4072000000000000, 0x4072000000000000)
	/Users/brendan/Documents/mygo/src/gonum.org/v1/plot/plot.go:165 +0x84c
gonum.org/v1/plot.(*Plot).WriterTo(0xc42042e000, 0x4072000000000000, 0x4072000000000000, 0x1396f61, 0x3, 0x1, 0x1, 0xc42000e790, 0xc42042e398)
	/Users/brendan/Documents/mygo/src/gonum.org/v1/plot/plot.go:445 +0x136
gonum.org/v1/plot.(*Plot).Save(0xc42042e000, 0x4072000000000000, 0x4072000000000000, 0x1396f5b, 0x9, 0x0, 0x0)
	/Users/brendan/Documents/mygo/src/gonum.org/v1/plot/plot.go:471 +0x130

The problem is that the calculation on heat.go171 is fmt.Println(int((v-h.Min)*ps + 0.5)). When h.Min = -inf, then ps = 0, which means the calculation is NaN, and the interpreted integer is a very negative integer which panics on the bound.
I'm not sure what the fix is, Scatter et. al. have errors returned in creation, maybe this should be the case as well? Or maybe the non-inf min should be calculated to allow default special casing.
```
```text
 We should not allow ±Inf or NaN bounds on creation. They make even less sense here than they do in the floats span functions. What were the parameters used in creating the heat map?
```
```text
 In this case, just the default. I had a -inf value in my grid and didn't realize it. This problem, though, isn't manifested until attempting to save the plot. I would suggest that NewHeatMap return (heat, err) giving an error for things like the (effective) data bounds are inf or NaN
```
```text
 SGTM
```
## 7.Feature request: Plot style iterator as a structure
May 27, 2015
```text
 The plotutil.AddLines, etc. is really useful for easily creating distinguishable lines. It would be nice if this could be exposed as an iterator so that it's easier to combine datasets in various combinations.
```
## 8.Change number of ticks based on plot size
Mar 15, 2015
```text
 Original issue 38 created by eaburns on 2012-06-21T12:20:38.000Z:
It's nice to have the number of tick marks computed by DefaultTicks change based on the size of the plot.
SPT (code.google.com/p/caml-spt) uses 1 major tick every 4cm, but no fewer than 2 major ticks, as its hint for how many ticks to attempt.  IMGO, SPT sometimes uses too few ticks.  Maybe 1 tick every 4cm, but no fewer than 3 would be good for Plotinum.
```
```text
 Comment #1 originally posted by eaburns on 2012-12-12T13:50:49.000Z:
<empty>
```
```text
 Comment #2 originally posted by eaburns on 2013-02-20T19:20:37.000Z:
<empty>
```
## 9.vg: soften third-party dependencies
Nov 9, 2017
```text
 At the moment, installing gonum.org/v1/plot/vg/draw has a hard requirement on various non-vendored third-party packages.
I'd like to propose a refactoring of the vg/{draw,vgeps,vgimg,vgpdf,vgsvg} so that implementations of vg.Canvas need to be registered, which the packages could do in their func init().
This would make all third-party dependencies soft requirements and a user might choose not to fulfill some of them in case the respective output formats are not required (e.g. one might want to use PNG but have no need for PDF).
As a side benefit, this would make it easy to incubate new implementations of the vg.Canvas interface, as these no longer need to be hardwired to vg/draw.
(Incidentally, this would fix issues like #238, but that's not the objective here.)
```
```text
 It would be good to see a plan for your implementation of this.
```
```text
 At the moment, I do not have plan, as I am not very familiar with the codebase. But if you are open to the general idea, I might get in deeper into the source and try to devise one in the next couple of weeks.
```
```text
 I'm open to the idea, but it's worth waiting for input from others here. The issues that you may face is handling how plot.WriterTo and plot.Save work and are documented, but I don't see anything that would make handling those issues impossible.
```
```text
 As plot.WriterTo calls draw.NewFormattedCanvas which returns an error my naive approach would be to simply let that error be non-nil in case the requested format has not been registered, which would then trickle down the stack to plot.Save as well. I'd add to the documentation to plot.Save and plot.WriterTo that trying any non-registered format would lead to such an error.
I wonder whether this was an API-breaking change, as on the one hand any using code would still compile while on the other hand the behaviour changes and care must be taken accordingly to import and register any implementations of the plot.Canvasinterface. Does the fact that there are no (SemVer) release tags have an impact? I'd be genuinely interested in the thoughts of the gonum contributors on these questions.
```
```text
 It will be breaking, but we aee breaking things at the moment. We had intended to freeze the API at the beginning of this year, but delays...
```
```text
 @kortschak So far, nobody else chimed in. Should we ping somebody specific to get input from?
```
```text
 I am a bit of two minds here.
on the one hand, being able to slim down the amount of dependencies pulled down by plot.Plot.Save is nice and requiring people to import _ "my/vg/backend" definitely looks like a nice nod to import _ "image/png".
on the other hand, this will break people's code at runtime...
but we could introduce something like gonum.org/v1/plot/vg/vgall that would register all the current backends.
```
```text
 As far as I understood from @kortschak breaking stuff is not showstopper here. I agree that it is very unfortunate that this would not already throw an error during compilation.

but we could introduce something like gonum.org/v1/plot/vg/vgall that would register all the current backends.

Yes, that's what I had in mind, too. An additional mitigation of the break might be to have a .go with a +build !minimal  build constraint that will import vg/vgall for a certain deprecation period. In that case users that want to benefit from the softened dependency chain would need to define -tags minimal at build time for said deprecation period. (I'm not attached to the name minimal here).
What do you think?
```
```text
 I know you wanted to see a plan for this first, but when I was trying to come up with one I was really surprised how little code was actually needed to get this done. Hence, I present my plan in the form of a commit: fawick@8b68616
```
```text
 Come to think of it, the package vgall is not needed. The file import.go could be part of gonum.org/v1/plot itself: fawick/plot@9eb0a40
```
```text
 Broadly I think that is OK, though there are a couple of correctness issues. The registry must be protected by a mutex (see how the font registry is protected) and the use of an appended slice and use of first match means that a re-registration of a format with a new implementation is not effective - a map would be simpler too.

The default use of importing all (doing this in a build tag-protected file in plot as you suggest) does retain runtime compatibility, so I think people should be OK with that.

If you want to put together a PR taking into account these points we can review it.
```
```text
 Also, registering a nil func should do delete (formats, name).
```
```text
 Thanks for the feedback! I wasn't thinking of re-registration, my blueprint was the image package from stdlib. I'll use a map as suggested and submit a PR tomorrow.
```
## 10.Best fit lines
Mar 15, 2015
```text
 Original issue 50 created by eaburns on 2012-08-04T13:09:46.000Z:
Compute a least-squares fit of some points.
```
```text
 Comment #1 originally posted by eaburns on 2012-08-04T13:10:02.000Z:
<empty>
```
```text
 Comment #2 originally posted by eaburns on 2012-08-24T20:49:07.000Z:
<empty>
```
```text
 Comment #3 originally posted by eaburns on 2013-02-20T19:24:51.000Z:
<empty>
```
## 11.Add Thumbnail function to bubbles.go
Mar 21, 2017
```text
 You can't create legends for bubbles. This is the case because Bubbles.go doesn't have a thumbnail function.
```
```text
 #344
```
## 12.plotutil.AddBoxPlots offset field
Mar 15, 2015
```text
 Original issue 97 created by eaburns on 2012-11-05T14:01:29.000Z:
A box plot offset field should be specifiable in the AddBoxPlots function.  In addition to a string argument to name the box data, an float argument preceding the data could specify the offset.
```
```text
 Comment #1 originally posted by eaburns on 2013-02-20T19:25:28.000Z:
<empty>
```
## 13.plotter: NewPolygon should take an interface rather than a list of interfaces
Apr 12, 2018
```text
 Right now, plotter.NewPolygon takes ...XYer as an input. However, it seems like it is difficult to actually get data in that format.
Would it make sense to add an interface something like:
XYers interface{
    Len() int
    LenAt(int) int
    XY(int, int) (x, y float64)
}
That might be easier to implement.
```
## 14.plot/palette: clean up API
Apr 30, 2017
```text
 I am troubled by the size of the ColorMap interface. It can be reduced in size to
// A ColorMap maps scalar values to colors.
type ColorMap interface {
	// At returns the color associated with the given value.
	// If the value is not between Max() and Min(), an error is returned.
	At(float64) (color.Color, error)

	// Min returns the current minimum value of the ColorMap.
	Min() float64

	// Max returns the current maximum value of the ColorMap.
	Max() float64

	// Palette returns a new Palette with the specified number of
	// colors from the ColorMap.
	Palette(colors int) Palette
}

and
// DivergingColorMap maps scalar values to colors that diverge from a central value.
type DivergingColorMap interface {
	ColorMap

	// ConvergePoint returns the value where the diverging colors meet.
	ConvergePoint() float64
}

The Set methods are required because the concrete types are not exported. This can be fixed with minor name juggling on the field names of the concrete types. For example:
// Luminance is a color palette that interpolates between control colors in a way that
// ensures a linear relationship between the luminance of a color and the value it represents.
// NewLuminance must be used to create a usable Luminance.
type Luminance struct {
	// colors are the control colors to be interpolated among.
	// The colors must be monotonically increasing in luminance.
	colors []cieLAB

	// scalars are the scalar control points associated with
	// each item in colors (above). They are monotonically
	// increasing values between zero and one that correspond
	// to the luminance of a given control color in relation
	// to the minimum and maximum luminance among all control
	// colors.
	scalars []float64

	// Alpha represents the opacity of the returned colors in
	// the range (0,1). NewLuminance sets Alpha to 1.
	Alpha float64

	// Range specifies the minimum and maximum values of the
	// range of scalars that can be mapped to colors using this
	// ColorMap.
	Range struct { Min, Max float64 }
}
```
```text
 Makes sense to me.
On this topic, additionally, I'm somewhat confused regarding what the purpose of the palette.Palette interface is. It only has one method, and the method doesn't take any arguments, so anything that is being used as a palette has to know ahead of time what colors it is going to return. Given that, would it make more sense to do type Palette []color.Color or just delete the Palette type altogether and use []color.Color instead?
```
```text
 I suppose one reason is that without a Palette interface we couldn't have a DivergingPalette. That may be enough to justify it, but it still feels to me like a lot of complexity just for that.
```
```text
 palette.Palette is justified by the same argument as fmt.Stringer.
```
```text
 The functions in fmt are set up so they can take either a string or a fmt.Stringer (or anything else). So the existence of fmt.Stringer doesn't prevent people from using plain strings in the functions. In this case, though, we have to choose between accepting []color.Color or palette.Palette; there doesn't seem to be a good way to accept both. I agree, though, that on balance palette.Palette is probably the better choice.
```
## 15.plot: Can't add data labels to the barchart
Oct 14, 2018
```text
 What are you trying to do?
I am trying to add data labels(show the value of each bar on top of it) to a bar chart.
What did you expect to happen?
A field in the BarChart struct which can set it to show the data labels or not and even where to put them.
What actually happened?
Right now there is no field in the BarChart struct 
What version of Go and Gonum/plot are you using?
Go: 1.10.3
Gonum: 5f3c436
```
## 16.Embed bitmapped images with vg.Canvas
Mar 17, 2015
```text
 Original issue 131 created by eaburns on 2013-07-01T01:17:26.000Z:
All of the backing formats support (to a greater or lesser degree) embedding of bit mapped images. It would be useful for the vg.Canvas interface to be able to embed bit mapped images for some applications.
```
```text
 Comment #1 originally posted by eaburns on 2013-07-01T12:10:41.000Z:
<empty>
```
```text
 Comment #2 originally posted by eaburns on 2013-07-07T18:57:04.000Z:
<empty>
```
## 17.Box plot fill color
Mar 15, 2015
```text
 Original issue 76 created by eaburns on 2012-09-03T13:41:43.000Z:
Other packages allow you to fill in box plot boxes with a color.  This could be useful when making grouped box plots where colors can signify similarity between items from separate groups, and it would be very easy to add.
```
```text
 Comment #1 originally posted by eaburns on 2013-02-20T19:24:34.000Z:
<empty>
```
## 18.Rethink plotutil
May 31, 2016
```text
 See discussion: https://groups.google.com/d/msg/gonum-dev/FrGpgA-pZQs/FCy8op2JCwAJ
In short, it seems worth rethinking the plot-util routines to have more gonum-like signatures and possibly to better match common use cases.
```
```text
 could you give one specific example?
e.g.  here is the current API:
func AddLines(plt *plot.Plot, vs ...interface{}) error {...}
it should in fact look like: ...
```
## 19.vg/vgpdf: support for non-ASCII text
Apr 26, 2017
```text
 this is a follow-up of #208.
currently, vg/vgpdf uses a PDF backend that doesn't support non-ASCII text.
either:

contribute support to upstream (bitbucket.org/zombiezen/gopdf/pdf).
this project is in maintainance mode, quite dormant.
switch to github.com/llcode/draw2d/draw2dpdf.
this project (or its sibling, draw2d/draw2dimg) record wrt backward compatibility and API stability is somewhat checkered.
switch to github.com/jung-kurt/gofpdf.

(just listing options at this point)
```
## 20.Plot cuts off lines
May 11, 2015
```text
 See the following code
http://play.golang.org/p/gNd64W9hDY

The maximum of the plot cuts off the peak of the line.
```
```text
 I think that the top of the plot needs to be padded by the line width.
It's ugly, but without changing the API, the only way that I can see to do this is to add GlyphBoxes for the extreme points in the line plotter. The size of the box should be linewidth/2.
```
## 21.meta: submit gonum/plot to JOSS
Aug 15, 2017
```text
 We should consider submitting a paper about gonum/plot to the journal of open source software:
http://joss.theoj.org/about
```
## 22.vg: harmonise UseDPI input parameter type
Jan 17, 2019
```text
 See #492 (comment).
We should choose a type, int or float64 for all the vg types' UseDPI calls.
```
## 23.plotutil: consider using palette to replace color slices
Feb 11, 2015
```text
 Currently plotutil uses the Color func to return a order set of colours for plots. Consider using a palette.Palette instead.
Issue of whether to take a Palette parameter or continue to use the current Color func.
```
## 24.Package import with `go get gonum.org/v1/plot/...` is broken — 404 returned.
Oct 25, 2017
```text
 I am unable to import this project with go get or dep, which never failed me before.
Call to Parsing meta tags from https://gonum.org/v1/plot?go-get=1 (status code 404) fails.
The page https://www.gonum.org/v1/plot?go-get=1 returns Page not found.
Output:
$ go get -v gonum.org/v1/plot/...

Fetching https://gonum.org/v1/plot?go-get=1
Parsing meta tags from https://gonum.org/v1/plot?go-get=1 (status code 404)
get "gonum.org/v1/plot": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot?go-get=1
gonum.org/v1/plot (download)
Fetching https://gonum.org/v1/plot/vg?go-get=1
Parsing meta tags from https://gonum.org/v1/plot/vg?go-get=1 (status code 404)
get "gonum.org/v1/plot/vg": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot/vg?go-get=1
get "gonum.org/v1/plot/vg": verifying non-authoritative meta tag
Fetching https://gonum.org/v1/plot?go-get=1
Parsing meta tags from https://gonum.org/v1/plot?go-get=1 (status code 404)
github.com/golang/freetype (download)
Fetching https://golang.org/x/image/math/fixed?go-get=1
Parsing meta tags from https://golang.org/x/image/math/fixed?go-get=1 (status code 200)
get "golang.org/x/image/math/fixed": found meta tag get.metaImport{Prefix:"golang.org/x/image", VCS:"git", RepoRoot:"https://go.googlesource.com/image"} at https://golang.org/x/image/math/fixed?go-get=1
get "golang.org/x/image/math/fixed": verifying non-authoritative meta tag
Fetching https://golang.org/x/image?go-get=1
Parsing meta tags from https://golang.org/x/image?go-get=1 (status code 200)
golang.org/x/image (download)
Fetching https://golang.org/x/image/font?go-get=1
Parsing meta tags from https://golang.org/x/image/font?go-get=1 (status code 200)
get "golang.org/x/image/font": found meta tag get.metaImport{Prefix:"golang.org/x/image", VCS:"git", RepoRoot:"https://go.googlesource.com/image"} at https://golang.org/x/image/font?go-get=1
get "golang.org/x/image/font": verifying non-authoritative meta tag
Fetching https://gonum.org/v1/plot/vg/fonts?go-get=1
Parsing meta tags from https://gonum.org/v1/plot/vg/fonts?go-get=1 (status code 404)
get "gonum.org/v1/plot/vg/fonts": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot/vg/fonts?go-get=1
get "gonum.org/v1/plot/vg/fonts": verifying non-authoritative meta tag
Fetching https://gonum.org/v1/plot/vg/draw?go-get=1
Parsing meta tags from https://gonum.org/v1/plot/vg/draw?go-get=1 (status code 404)
get "gonum.org/v1/plot/vg/draw": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot/vg/draw?go-get=1
get "gonum.org/v1/plot/vg/draw": verifying non-authoritative meta tag
Fetching https://gonum.org/v1/plot/vg/vgeps?go-get=1
Parsing meta tags from https://gonum.org/v1/plot/vg/vgeps?go-get=1 (status code 404)
get "gonum.org/v1/plot/vg/vgeps": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot/vg/vgeps?go-get=1
get "gonum.org/v1/plot/vg/vgeps": verifying non-authoritative meta tag
Fetching https://gonum.org/v1/plot/vg/vgimg?go-get=1
Parsing meta tags from https://gonum.org/v1/plot/vg/vgimg?go-get=1 (status code 404)
get "gonum.org/v1/plot/vg/vgimg": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot/vg/vgimg?go-get=1
get "gonum.org/v1/plot/vg/vgimg": verifying non-authoritative meta tag
github.com/llgcode/draw2d (download)
Fetching https://golang.org/x/image/draw?go-get=1
Parsing meta tags from https://golang.org/x/image/draw?go-get=1 (status code 200)
get "golang.org/x/image/draw": found meta tag get.metaImport{Prefix:"golang.org/x/image", VCS:"git", RepoRoot:"https://go.googlesource.com/image"} at https://golang.org/x/image/draw?go-get=1
get "golang.org/x/image/draw": verifying non-authoritative meta tag
Fetching https://golang.org/x/image/math/f64?go-get=1
Parsing meta tags from https://golang.org/x/image/math/f64?go-get=1 (status code 200)
get "golang.org/x/image/math/f64": found meta tag get.metaImport{Prefix:"golang.org/x/image", VCS:"git", RepoRoot:"https://go.googlesource.com/image"} at https://golang.org/x/image/math/f64?go-get=1
get "golang.org/x/image/math/f64": verifying non-authoritative meta tag
Fetching https://golang.org/x/image/tiff?go-get=1
Parsing meta tags from https://golang.org/x/image/tiff?go-get=1 (status code 200)
get "golang.org/x/image/tiff": found meta tag get.metaImport{Prefix:"golang.org/x/image", VCS:"git", RepoRoot:"https://go.googlesource.com/image"} at https://golang.org/x/image/tiff?go-get=1
get "golang.org/x/image/tiff": verifying non-authoritative meta tag
Fetching https://golang.org/x/image/tiff/lzw?go-get=1
Parsing meta tags from https://golang.org/x/image/tiff/lzw?go-get=1 (status code 200)
get "golang.org/x/image/tiff/lzw": found meta tag get.metaImport{Prefix:"golang.org/x/image", VCS:"git", RepoRoot:"https://go.googlesource.com/image"} at https://golang.org/x/image/tiff/lzw?go-get=1
get "golang.org/x/image/tiff/lzw": verifying non-authoritative meta tag
Fetching https://gonum.org/v1/plot/vg/vgpdf?go-get=1
Parsing meta tags from https://gonum.org/v1/plot/vg/vgpdf?go-get=1 (status code 404)
get "gonum.org/v1/plot/vg/vgpdf": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot/vg/vgpdf?go-get=1
get "gonum.org/v1/plot/vg/vgpdf": verifying non-authoritative meta tag
bitbucket.org/zombiezen/gopdf (download)
go: missing Mercurial command. See https://golang.org/s/gogetcmd
package gonum.org/v1/plot
	imports bitbucket.org/zombiezen/gopdf/pdf: exec: "hg": executable file not found in $PATH
Fetching https://gonum.org/v1/plot/vg/vgsvg?go-get=1
Parsing meta tags from https://gonum.org/v1/plot/vg/vgsvg?go-get=1 (status code 404)
get "gonum.org/v1/plot/vg/vgsvg": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot/vg/vgsvg?go-get=1
get "gonum.org/v1/plot/vg/vgsvg": verifying non-authoritative meta tag
github.com/ajstarks/svgo (download)
Fetching https://gonum.org/v1/plot?go-get=1
Parsing meta tags from https://gonum.org/v1/plot?go-get=1 (status code 404)
get "gonum.org/v1/plot": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot?go-get=1
Fetching https://gonum.org/v1/plot/plotter?go-get=1
Parsing meta tags from https://gonum.org/v1/plot/plotter?go-get=1 (status code 404)
get "gonum.org/v1/plot/plotter": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot/plotter?go-get=1
get "gonum.org/v1/plot/plotter": verifying non-authoritative meta tag
Fetching https://gonum.org/v1/plot/palette?go-get=1
Parsing meta tags from https://gonum.org/v1/plot/palette?go-get=1 (status code 404)
get "gonum.org/v1/plot/palette": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot/palette?go-get=1
get "gonum.org/v1/plot/palette": verifying non-authoritative meta tag
Fetching https://gonum.org/v1/plot/tools/bezier?go-get=1
Parsing meta tags from https://gonum.org/v1/plot/tools/bezier?go-get=1 (status code 404)
get "gonum.org/v1/plot/tools/bezier": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot/tools/bezier?go-get=1
get "gonum.org/v1/plot/tools/bezier": verifying non-authoritative meta tag
Fetching https://rsc.io/pdf?go-get=1
Parsing meta tags from https://rsc.io/pdf?go-get=1 (status code 200)
get "rsc.io/pdf": found meta tag get.metaImport{Prefix:"rsc.io/pdf", VCS:"git", RepoRoot:"https://github.com/rsc/pdf"} at https://rsc.io/pdf?go-get=1
rsc.io/pdf (download)
```
```text
 I just tried this here in a clean GOPATH and it works (it complains with 404's and probably shouldn't, @btracey), but it still downloads and builds (you can see that it downloads since otherwise it wouldn't pull the dependencies in your pasted output).
Your paste does miss the bitbucket.org/zombiezen/gopdf dep, but that's because you don't have hg installed. Can you try this again after installing mercurial?
```
```text
 Oh no, how embarrassing😖 Cleaned up GOPATH and installed mercurial and now it works.
Both dep init -v and go get -v gonum.org/v1/plot/... ran through without issues. Two hours ago dep init -v got stuck for minutes multiple times at (1) try gonum.org/v1/plot@master.
go get -v gonum.org/v1/plot/... showed 404s and exited non-zero and I somehow overlooked the failed hg dependency.
I've done some testing with a fresh Ubuntu 16.04 VM.
When mercurial is not installed, dep init hangs at this position:
$ dep init -v

Getting direct dependencies...
Checked 1 directories for packages.
Found 1 direct dependencies.
Root project is "github.com/denisbrodbeck/entro"
 1 transitively valid internal packages
 1 external packages imported from 1 projects
(0)   ✓ select (root)
(1)	? attempt gonum.org/v1/plot with 1 pkgs; 4 versions to try
(1)	    try gonum.org/v1/plot@master
When mercurial is installed, dep runs through.
$ dep init -v

Getting direct dependencies...
Checked 4 directories for packages.
Found 1 direct dependencies.
Root project is "github.com/denisbrodbeck/entro"
 3 transitively valid internal packages
 1 external packages imported from 1 projects
(0)   ✓ select (root)
(1)	? attempt gonum.org/v1/plot with 1 pkgs; 4 versions to try
(1)	    try gonum.org/v1/plot@master
(1)	✓ select gonum.org/v1/plot@master w/8 pkgs
(2)	? attempt golang.org/x/image with 2 pkgs; 1 versions to try
(2)	    try golang.org/x/image@master
(2)	✓ select golang.org/x/image@master w/3 pkgs
(3)	? attempt github.com/golang/freetype with 2 pkgs; 3 versions to try
(3)	    try github.com/golang/freetype@master
(3)	✓ select github.com/golang/freetype@master w/3 pkgs
(4)	? revisit golang.org/x/image to add 1 pkgs
(4)	  ✓ include 2 more pkgs from golang.org/x/image@master
(4)	? attempt bitbucket.org/zombiezen/gopdf with 1 pkgs; 6 versions to try
(5)	    try bitbucket.org/zombiezen/gopdf@default
(5)	✓ select bitbucket.org/zombiezen/gopdf@default w/1 pkgs
(5)	? attempt github.com/llgcode/draw2d with 2 pkgs; 5 versions to try
(6)	    try github.com/llgcode/draw2d@master
(6)	✓ select github.com/llgcode/draw2d@master w/3 pkgs
(6)	? revisit golang.org/x/image to add 2 pkgs
(7)	  ✓ include 2 more pkgs from golang.org/x/image@master
(6)	? attempt github.com/ajstarks/svgo with 1 pkgs; 9 versions to try
(8)	    try github.com/ajstarks/svgo@master
(8)	✓ select github.com/ajstarks/svgo@master w/1 pkgs
(7)	? revisit github.com/golang/freetype to add 1 pkgs
(9)	  ✓ include 1 more pkgs from github.com/golang/freetype@master
  ✓ found solution with 22 packages from 6 projects

Solver wall times by segment:
     b-source-exists: 16.637893064s
     b-list-versions:  9.386182938s
         b-list-pkgs:  541.701378ms
              b-gmal:  374.514969ms
  b-deduce-proj-root:  275.168535ms
         select-atom:    4.245835ms
            new-atom:    2.256841ms
             satisfy:    2.161001ms
         select-root:     180.629µs
            add-atom:      22.268µs
               other:      21.778µs

  TOTAL: 27.224349236s

  Using master as constraint for direct dep gonum.org/v1/plot
  Locking in master (def97c8) for direct dep gonum.org/v1/plot
  Locking in master (f7e31b4) for transitive dep golang.org/x/image
  Locking in master (e2365df) for transitive dep github.com/golang/freetype
  Locking in default (1c63dc6) for transitive dep bitbucket.org/zombiezen/gopdf
  Locking in master (dcbfbe5) for transitive dep github.com/llgcode/draw2d
  Locking in master (ac87889) for transitive dep github.com/ajstarks/svgo
(1/6) Wrote github.com/golang/freetype@master
(2/6) Wrote github.com/ajstarks/svgo@master
(3/6) Wrote github.com/llgcode/draw2d@master
(4/6) Wrote gonum.org/v1/plot@master
(5/6) Wrote golang.org/x/image@master
(6/6) Wrote bitbucket.org/zombiezen/gopdf@default
I guess, I'll open an issue at dep.
@kortschak Thank you for your quick reply and sorry for any inconvenience. It would be great, if your project README would include some hint, that the user needs mercurial installed, or installation might fail (creating a PR right now).
```
```text
 With any luck mercurial will not be required in the near future (#389), but in the Go ecosystem, it pays to have mercurial installed since there is a history there (Go was previously distributed via mercurial).
```
```text
 Could you open an issue on the website with what the behavior should be? The page doesn't exist, so I would expect a 404
```
```text
 Hmm, IMO it's just a minor issue. I've looked up the go get documentation and in the end go get handles this rather clever.
Given the output:
Parsing meta tags from https://gonum.org/v1/plot?go-get=1 (status code 404)
get "gonum.org/v1/plot": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot?go-get=1
gonum.org/v1/plot (download)

We can see, that go get still searches and finds the meta tag in the doc header, even when the server returns a 404.
One could get rid of the 404s, if there's an actual hugo page behind the path /v1/plot/. The ?go-get=1 part shouldn't cause issues.
I suppose, something like hugo new v1/plot.md (without actual content) might suffice.
Maybe for all top level import paths:
hugo new v1/gonum.md
hugo new v1/plot.md
hugo new v1/netlib.md
hugo new v1/hdf5.md
and for each md:
+++
title = "gonum"
date = "2017-10-26T00:00:00"
selected = false
highlight = false
math = false
+++
```
```text
 Others return a refresh to godoc, e.g.
~ $ curl golang.org/x/tools
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="golang.org/x/tools git https://go.googlesource.com/tools">
<meta name="go-source" content="golang.org/x/tools https://github.com/golang/tools/ https://github.com/golang/tools/tree/master{/dir} https://github.com/golang/tools/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/golang.org/x/tools">
</head>
<body>
Nothing to see here; <a href="https://godoc.org/golang.org/x/tools">move along</a>.
</body>
</html>
```
## 25.vg/vgpdf: add ability to create PDFs with multiple pages
May 30, 2018
```text
 it would be great to be able to create a bunch of plots and save each in a new PDF page:
func plot(x,y []float64) (*plot.Plot, error) { ... }

func analysis() {
    for i, sample := range samples {
        p, err := plot(sample.x, sample.y)
        // ... save p to page-i ...
    }
}
```
```text
 (Short backstory: moving many chart report from my normal HTML -> PDF report pipeline due to current PDFer limitations in performance with 80+ charts in a single document).
In my case, I'd ideally want to be able to insert charts as some type of semi-opaque objects in a larger PDF document. To get me through the day, I'll probably need to modify vgpdf locally to expose the underlying pdf writer.  Even if there wasn't a more advanced layout option, I need to put N plots on a single page.
Okay, I know what I want to tackle for a side project: document layout (eep!). That would be really useful when generating reports.
```
```text
 For multiple plots on a single page, you can use draw.Tiles and plot.Align. For multiple plots that are not arranged in a regular grid, you can use draw.Crop.
```
```text
 @ctessum Thanks for the tip. That helped.
In vgpdf, I ended up adding locally the following function:
func (c *Canvas) AddPage() {
	_, h := c.Size()
	c.Pop()
	c.doc.AddPage()
	c.Push()
	c.Translate(vg.Point{0, h})
	c.Scale(1, -1)
}

This seems to work and might be a reasonable API. With didn't need to expose the underlying PDF writer either. I can send a PR if you're interested. It solved my immediate need anyway.
```
```text
 That SGTM.
```
```text
 Would something like NextPage be a better name? Or MoveNewPage? Something might better reflect that the cursor is moving to the next page as it gets added. That is what's happening, correct?
```
```text
 NewPage ? MakeNewPage?
(for me NextPage sounds too much like iteration.)
```
```text
 @kardianos : want to give a stab at sending a PR? (I have a colleague who would be interested in getting this multi page feature.)
```
## 26.plotter: Add capability for colored scatter plots
Apr 16, 2015
```text
 Feature request for having a plotter where data points are {x,y,z} values and the z values are plotted according to color. Scatter plots are useful when the data is not evenly spread throughout the space (so generating a grid for a heat-map is difficult), or when there is noise in the data, so there isn't really a 'surface' to plot.
All of the parts for colored-scatter exist. We can already plot 3-D scattered data where the glyph varies with the third dimension (plotter.Bubbles), and we can plot 3-D data where the color varies with the third dimension (Contour). Thus, in my opinion, the benefit to providing colored scatter greatly outweighs the small marginal maintenance cost of the extra plotter.
```
```text
 Sad to see this was posted 2 years ago. But, I agree as I posted in #345. The HeatMap requires us to interpolate values and associate them with a grid. A palette colored scatter plot allows us to freely assign x / y values in continuous space colored by a z value.
```
```text
 This is a straightforward plotter to write; take the codewalk shown for Bubbles and instead of setting the radius, set the colour based on a colour interpolation as done in HeatMap.
```
## 27.vg: add example for loading/using non-Latin fonts
Apr 11, 2018
```text
 we should probably add an example for loading and then using a non pre-registered font (possibly a non-english/non-latin one) to the vg package.

example with png
example with pdf
example with svg

an example with tex might be complicated...
```
```text
 @sbinet don't know if this is what you are talking about here, but I'd be very interested to get an example of how to get the label of a plot axis showing some greek characters...
```
```text
 nope, that wasn't it :)
what I think you are hinting at is $\eta$, right?
for that kind of plot, one usually relies on the gonum.org/v1/plot/vg/vgtex backend.
otherwise, one can also use UTF-8 characters, such as: η or Ψ.
package main

import (
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	pl, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	pl.Title.Text = "J/Ψ stuff"
	pl.X.Label.Text = "η"
	pl.X.Min = -math.Pi
	pl.X.Max = +math.Pi
	pl.Y.Label.Text = "Stuff (A.U)"

	pl.Add(plotter.NewGrid())
	err = pl.Save(20*vg.Centimeter, 10*vg.Centimeter, "out.png")
	if err != nil {
		log.Fatal(err)
	}
}
```
## 28.A color scale is tied to a single heat map, making it impossible to have multiple heatmaps with the same color scale
Jul 19, 2016
```text
 It could be fixed by making the color-mapper separate from the heat map, and then either requiring the color-mapper to be an argument to NewHeatMap, or requiring GridXYZ to be a matrix of colors rather than a matrix of floats.
Making these separate would also be useful in other areas, such as allowing plots with points or lines that are colored according to some attribute of the data.
I think this is related to #302.
```
```text
 I don't believe this is true. If the same palette is passed to NewHeatmap or used in the struct declaration they will be using the same colour mapping, provided the min and max are conformed to the same range. This is trivial after all the Heatmaps have been created.
func unifyRanges(h []plotter.Heatmap) {
    if len(h) == 0 {
        return
    }
    min := h[0].Min
    max := h[0].Max
    for _, r := range h[1:] {
        min = math.Min(min, r.Min)
        max = math.Max(max, r.Max)
    }
    for _, r := range h {
        r.Min = min
        r.Max = max
    }
}

The legend colour map would just be one of the elements of h.
```
```text
 That makes sense. I think, though, there still might be a rationale for separating the color map from the heat map, as having a standalone color map could enable functionality such as:

colored points or lines (or potentially polygons) as mentioned above
separate color maps for linear and log scaling
separate color maps for categorical and continuous data
making it easier to implement a standard color legend, as the legend could be a method of the color map

Food for thought.
```
```text
 I could see that working, though I don't really understand how it help with some of the things you list, notable, bullet points 1-3. The last bullet point would find a fair amount of additional API to achieve it.
One of the things that needs to be considered here is that plot does (intentionally from @eaburns AFAIUI and certainly from me) a minimal amount of things to get a plot done (in some cases it actually does too much in my view and starts to feel like a framework - debugging things can be difficult because of the amount of data tracing required). This is quite unlike other plotting environments (e.g. ggplot2) which hide most things and perform a lot of magic. These packages make simple/common things easy, but unknown/complicated things virtually impossible. I hope we don't end up in that space.
```
```text
 The functionality in question here is assigning colors to data points. Currently, HeatMap is the only part of the package that does this, but there are a bunch of potential applications of this functionality (eg, 1, 2, 3). There is already at least one  (@btracey)  external plotter that does it independently.
To me, the question is whether every plotter that needs to assign colors to data points should have that logic built into it independently, or whether it should be separated out and implemented once. The answer to that question probably depends whether having a general function/interface for assigning colors to data points is worth the added API surface (which I feel it does because I use that functionality a lot in different contexts, but others may disagree), but I don't think it would necessarily increase the level of "magic".
An example API could be
type ColorMap interface {
  Color(float64) (color.Color, error) // the argument could also be an interface{} for more generality
  Legend(vertical bool) (*plot.Plot, error)
}

type LinearColorMap struct {
    // Palette is the color palette used to render
    // the color map. Palette must not be nil or
    // return a zero length []color.Color.
    Palette palette.Palette

    // Underflow and Overflow are colors used to fill
    // color map elements outside the dynamic range
    // defined by Min and Max.
    Underflow color.Color
    Overflow  color.Color

    // Min and Max define the dynamic range of the
    // color map.
    Min, Max float64
}

func (cm *LinearColorMap) Color(v float64) (color.Color, error) {
  // Logic from heatmap.Plot
}

func (cm *LinearColorMap) Legend(vertical bool) (*plot.Plot, error) {
  // Logic adapted from PR #290
}
```
```text
 OK, I understand now. That looks reasonable. The entanglement that I don't like is how Min and Max are used by the plotter. One approach would be that a plotter will only expand the range of a ColorMap type held by it. This would give a natural result equivalent to my snippet above.
I would not think an interface{} param for Color would be a good idea.
Are we happy with the notion of a plot.Plot being a panel unit? It feels ungainly to me, but I guess that's something I could get over given time.
```
```text
 Having Legend return a plot.Plot allows a lot of customizability with tick marks, titles, etc that would otherwise have to be duplicated, but it might be possible with some refactoring to have a ColorLegend type that includes the necessary parts of a plot.Plot without actually being a plot.Plot and with minimal code duplication. Not sure if that would be worth the work though.
A solution to the entanglement could be to not have the plotters edit the color map at all, but instead require Min and Max to be set by the user or have a ColorMap.ExpandRange(...float64) method.
The reason to have an interface{} param for Color() would be to allow an implementation of categorical color maps, which would assign different colors to categories like "apple farms", "orange farms", etc, and therefore would require a string input instead of a float. This might be a case where the added functionality does not warrant the added confusion, however.
```
```text
 A solution to the entanglement could be to not have the plotters edit the color map at all, but instead require Min and Max to be set by the user or have a ColorMap.ExpandRange(...float64) method.

Mmmm. This is what I was thinking about in my comment above, but I subsequently forgot. Losing the automatic range determination would not be a good idea IMO since then all users would always need to determine ranges - from where and how? ExpandRange feels gross, but I guess would be OK, though not with a variadic signature. Maybe Include(float64).

The reason to have an interface{} param for Color() would be to allow an implementation of categorical color maps, which would assign different colors to categories like "apple farms", "orange farms", etc

These should be distinct type I think.
```
## 29.Make a wiki page on tick marks
Mar 15, 2015
```text
 Original issue 107 created by eaburns on 2012-12-12T18:41:33.000Z:
The default ticks are usually fine, but sometimes you want to specify your own, or maybe you want the default ticks but with commas every 3 place values, etc.
See: http://play.golang.org/p/WqpRbC8Ycw
```
```text
 Comment #1 originally posted by eaburns on 2013-02-20T19:26:16.000Z:
<empty>
```
```text
 Comment #2 originally posted by eaburns on 2013-02-20T19:26:23.000Z:
<empty>
```
## 30.Docs on plotutil functions should describe behavior on multiple calls
Mar 17, 2015
```text
 Original issue 138 created by eaburns on 2014-02-03T00:24:33.000Z:
For example, each call to plotutil.AddLines will restart the color pallet. If you want to cycle through the colors, you have to add all lines in a single call.
```
## 31.Smarter automatic tick mark selection
May 26, 2015
```text
 I know tick mark selection is hard, but I think the current algorithm could be improved. If one generates 2-dimensional random numbers, and plots the result, the tick marks are at 0.3, 0.6, and 0.9. Given that the data is very close to being between 0 and 1, it seems more reasonable to me to have tick marks drawn at 0 and 1. If the minimum of the axis is changed to 0, then the ticks are at least 0, 0.3, 0.6 and 0.9
```
## 32.Replace plotter.XYs with slices
May 31, 2016
```text
 In my experience, most data exists as x, y []float64, and not as []struct{X,Y float64}. plotter should replace these types
See discussion https://groups.google.com/d/msg/gonum-dev/FrGpgA-pZQs/FCy8op2JCwAJ
```
```text
 Now if we had tables...
```
```text
 wrt tables, as far as gonum/plot is concerned, everything is a float64.
so, perhaps we could get away with a float64-based ndim array?
but then, is it really the best option to couple gonum/plot with that specific implementation?
should we give @Kunde21's numgo.Array64 a go?
relying on []float64 seems simpler. row-wise data isn't very cache friendly anyways...
```
```text
 That was not a serious suggestion.
```
```text
 Why not just add the code:
type XYArray struct {
   X, Y []float64
}

func (xy XYArray) Len() int {
   return len(xy.X)
}

func (xy XYArray) XY(i int) (x,y float64) {
   return xy.X[i], xy.Y[i]
}
Then one could just do
vals := XYArray{X: x, Y: y}
and directly add vals to a plotter.
```
```text
 https://godoc.org/github.com/btracey/myplot#VecXY :)
Still, I think the slice representation is more natural for gonum/plot. I think it's more common for users of the code, and it's easier to add in more information (say, z data). Frequently the data does need to exist as a slice even internal to plot, see for example the code for NewErrorPoints
```
```text
 So is the question, then, whether to replace the XYs type with VecXY or XYArray, or whether to replace the XYer interface with a concrete type in function arguments? i.e.
NewScatter(xys XYer) 
would become
NewScatter(x, y []float64)
To me, the first option seems like a good idea. I see a tradeoff with the second option, where making the change would make things easier for a lot use cases (i.e., where the data to be plotted is already in x and y arrays), but it would make things more difficult for some use cases (e.g., plotting geographic locations, which typically are in the form struct{X, Y float64}).
I would lean toward keeping the current signature for the function arguments but replacing XYs with VecXY or XYVector as that seems to be most in line with the 'do less, enable more' philosophy, and allows the user to make decisions about their data format rather than having the decisions made for them by the plotting package.
```
```text
 I think your suggestion is reasonable, but just a bit of devil's advocate: Is it actually true that geographic data usually come as struct{X, Y float64})? Doesn't it usually come as struct{Lat, Long, Altitude, {other data} float64})? It doesn't affect your argument, but it does mean that some kind of type munging is always necessary for that kind of data.
```
```text
 Yes, although I think that supports my point rather than refutes it. Having an interface function argument means that the user can just implement the XYer interface for whatever complicated data type they have (which, as shown above, takes 6 lines of code), rather have having to mung. So, for example:
type Point {
   Lat,Lon,Altitude, other, variables float64
}

type Points []Point

func (p Points) Len() int {
   return len(p)
}

func (p Points) XY(i int) (x,y float64) {
   return p[i].Lon, p[i].Lat
}
```
```text
 Yes, I agree it supports your point.
```
```text
 Maybe I'm confused, because I don't understand the problem. Doesn't
everything take the XYer interface? You can implement that however you
like. That's why I made it an interface.
On Wed, Jun 1, 2016, 12:24 Brendan Tracey notifications@github.com wrote:

Yes, I agree it supports your point.
—
You are receiving this because you are subscribed to this thread.
Reply to this email directly, view it on GitHub
#284 (comment), or mute
the thread
https://github.com/notifications/unsubscribe/AAOXB2z8efGhNToSNjq8M8VeklCO56Fqks5qHbJZgaJpZM4Iq_th
.
```
```text
 I think this ticket should be about providing an easy way to create  XYer from X and Y.
This can be done by adding XYArray type as proposed by @ctessum (the name requires some thought though).
Or via function:
func ZipXY(x []float64, y []float64) XYer

Adding some convenience functions like:
NewScatterXY(x []float64, y []float64) (*Scatter, error)

might be good idea too.
```
```text
 @kortschak wrt tables and what not: I was browsing the interwebs the other day and stumbled upon @aclements' playground for providing a gg-inspired toolkit for Go:
https://godoc.org/github.com/aclements/go-gg
especially: https://godoc.org/github.com/aclements/go-gg/table#Table
It's an interesting read :)
```
```text
 That is something that looks like what we would be likely to implement as part of a data frame package, but not what I was talking about here.
```
```text
 I'm really not sure this is worth it. The code for it is just
type zip struct { x, y []float64 }

func (z zip) Len() int                { return len(z.x) }
func (z zip) XY(i int) (x, y float64) { return z.x[i], z.y[i] }

This is less than implementing sort.Interface.
```
```text
 you could flip it around: "sort" provides implementations for sort.Strings, sort.Float64s and sort.Ints and it's in the stdlib (so the bar for inclusion is higher.)
adding just one exported function (func ZipXY(x, y []float64) XYer) is IMHO hitting the sweet spot.
```
```text
 I have to agree with @sbinet about the usefulness of a ZipXY(x,y []float64) XYer) function.  A portion of the users will have existing slices of structs that can easily implement XYer, but there's a lot of data that starts in the form of one X vector and many Y's.  That ends up requiring a zip-style function in each and every project that doesn't have a clean slice of structs which can implement XYer directly.  Yes, it's a helper function, but it's not even close to an esoteric use.
Edit:  By comparison sort only needs an interface implementations, whereas a set of Y slices requires a new constructor or closure before implementing the XYer interface.
```
```text
 What does the function give you? If the type is exported you get exactly the same functionality with a smaller API surface, unless there is a proposal that ZipXY panics when len(x) != len(y), which the implementation may also do (though later)* and can be made to explicitly via the Len method.
* This is another issue I'd like to address again - that all data are copied by plot.
```
```text
 The copying was a conscious decision. If we don't copy the data, then it can change out from under the plotter — safer to copy it. Is copying really causing a problem for anyone? Either the plotter does not copy and stores only a summary of the data (boxplot), or the plotter copies the data and plots all of it (lines, scatter, etc). In the first case there is no copy, so it's not a problem. In the second case, if you are passing so much data that you just can't copy it, then you will certainly will have an issue when you go to plot that many points.
```
```text
 Yeah; we've had this discussion in the past. All our (gonum) data structures allow things to change underfoot unless they are created by the type/func. Because of this I think it's something worth looking at again. My interest is not size/allocations (which is our drive elsewhere) but flexibility.
```
```text
 As a new user of plot and plotter, I would definitely not want to see this change.  The current XYer approach is elegant and easy to use with a myriad of data types.  I am plotting a lot of things, and in all cases I have a slice or set of much more complicated data structures form which I generate an 'X' and 'Y' for plot.  Data definitely does not exist as X, Y []float64.  I could construct that, but it would be much less user friendly that the current XYer interfaces.  As it is with XYer, I am even using cases with different interfaces pointing the same underlying data structures with different XY methods returning different 'X' and 'Y' values for different plots based on the same underlying data structures.  If there is anything new, try to keep the XYer approach alongside it.
```
```text
 I don't think anyone here is suggesting to remove the XYer interface approach.
```
```text
 A ZipXY(x,y []float64) XYer) would definitely be appreciated, that's actually something that I was just searching for as my Data tends to be in slices of X /y data (one slice for each)
Some convenience functions wouldn't hurt anyone, especially ones you would need to write again and again (or write an own package, which defeats the point of using a plotting package in the first place)
```
## 33.Shading above and below lines
Mar 15, 2015
```text
 Original issue 55 created by eaburns on 2012-08-11T12:42:46.000Z:
It should be possible to add shading above and below lines.
```
```text
 Comment #1 originally posted by eaburns on 2012-08-24T20:49:23.000Z:
<empty>
```
```text
 Comment #2 originally posted by eaburns on 2013-02-20T19:24:10.000Z:
Just need to shade the polygon defined by the points of the line and the minimum X and Y values of the plot.  The user can manually use the "painter's algorithm" so that shading doesn't obscure other lines.
This can enable stacked lines plots (requested by JB).
```
## 34.Tick display precision should be configurable
Feb 9, 2017
```text
 Commit 2651c0e added a displayPrecision variable in line 18 (2651c0e#diff-d7f174bedb1bfb5e7a82248d40e9b46fR18).
This variable is used to format float64 ticks. The default value of 4 causes a small-ish number like 20000 to be rendered as 2e+04.
The value should be configurable to allow the user to define if he ever wants scientific notation or not.
I can submit a PR that exposes DefaultDisplayPrecision as a global variable to plot, by changing the const to an exported var but you might have something else in mind. Please let me know how you would like to expose access to that precision without forcing the user to implement his own ticks logic just for the precision.
The approach from the wiki example on the comma separated ticks marks based on the default ticks doesn't work because at that point the tick string has already been generated using the displayPrecision.
```
```text
 I like this idea, although I'm not a huge fan of using a global variable, as a user may want different precision for different plots or axes.
Would it be better to change the Ticker interface to the following:
type Ticker interface {
    // Ticks returns Ticks in a specified range
    Ticks(min, max float64, format func(v float64, prec int) string) []Tick
}
and then add a field to Axis.Tick that allows for the specification of a custom format function? (If the custom format function was nil, then the default would be used.)
```
```text
 Sounds like a plan. I'll give it a shot during the weekend hopefully.
```
```text
 This also allows the possibility of exporting some of these helpers.
```
```text
 I created PR #362.
I expanded the Ticker interface but did not add the field to Axis because the places where the format was used had no knowledge of Axis so Axis couldn't be checked there. I might be missing something.
```
## 35.vg: drop use of github.com/jteeuwen/go-bindata
Mar 22, 2018
```text
 This has made the news lately but I forgot we were using that package in gonum/plot.
we should replace it with something else.
```
```text
 Can you provide background please. I had not heard anything about this.
```
```text
 Sure: https://twitter.com/francesc/status/961249107020001280?lang=en
```
```text
 Wow! That's not good.
```
```text
 https://github.com/rakyll/statik?
```
```text
 I have also come across:
https://github.com/gobuffalo/packr
(I have been using it for my Go based forum. it's serviceable.)
```
## 36.plotter: NewHist panics with valid data
Sep 25, 2016
```text
 Code below. Sorry about the large amount of data.
I think it's a rounding problem. The panic message also needs improvement, it does not currently say why the code is panicking.
package main

import (
    "log"

    "github.com/gonum/plot/plotter"
)

var histPts = plotter.Values{0.37009128387828555, 0.4140510262107388, 0.4613139832428571, 0.4882434393955207, 0.5190228171834, 0.5220961292797093, 0.5240823079550753, 0.5304779063530312, 0.5525206878884541, 0.5658303918662759, 0.5768353094178024, 0.59154499761952, 0.5980386804788334, 0.5992479195652316, 0.605892717785441, 0.6091241120201799, 0.6143402964977033, 0.6159595204094886, 0.620997717676013, 0.6262053858147448, 0.6263241431037279, 0.6285301667294215, 0.6302467763370642, 0.6321707217945572, 0.6325729874560262, 0.6349444820417076, 0.6364659335508787, 0.6379882175555828, 0.6445552383271045, 0.6450027892756025, 0.645065529700484, 0.6456550492945196, 0.6458745585602601, 0.6461070030358175, 0.6461837020995134, 0.6499616248432533, 0.6502830252230443, 0.6502874296465386, 0.6532288210172918, 0.6533637452283967, 0.6542867506290639, 0.6555228522130342, 0.656578406949546, 0.656613524419828, 0.6581968552181601, 0.6586494763948838, 0.6601293357008619, 0.6612057593894385, 0.6612616385532253, 0.6647635310182385, 0.6666273823926949, 0.6671411810300372, 0.6709507632617026, 0.675394410740252, 0.6767163266295901, 0.677202276783287, 0.6773980463493011, 0.6802316091300395, 0.6817851050485755, 0.6822745172983482, 0.685320354206143, 0.6855413736612839, 0.6858649326895024, 0.6863055485934907, 0.6880321357872616, 0.6913273090466604, 0.6916889180597929, 0.6918013762904076, 0.6931076222376964, 0.6957684240267794, 0.6960164256619722, 0.6994398390716152, 0.7037755384798435, 0.7083913528937902, 0.7112072989225611, 0.7132327012884291, 0.7134471142162012, 0.7137300005526057, 0.7139403432167455, 0.714055637561479, 0.7150934792837121, 0.7153405114547448, 0.7166123517245541, 0.7175850204191422, 0.7177373869189501, 0.7184601082121338, 0.7192433925795171, 0.7196944515008858, 0.7205104979859092, 0.7230377303043414, 0.723788558389263, 0.7241941772558803, 0.7249807788936329, 0.7251067747417151, 0.7257717262779011, 0.72601826401791, 0.7268439951817564, 0.7274131939227012, 0.7278416382892511, 0.7280017308314175, 0.7283115728114169, 0.7284015806732758, 0.7313545743885275, 0.732252123678662, 0.7325382414705228, 0.734084764610445, 0.734095930322616, 0.7342280292410742, 0.7343649390557402, 0.7349686784210283, 0.7369663375908161, 0.7369946996254545, 0.7375629294676562, 0.7376432892574913, 0.7377764323864369, 0.7377844600454012, 0.7379429298222665, 0.7381015143678554, 0.7407098165194961, 0.7430111380376085, 0.7431513751520586, 0.7432309222242007, 0.7434927380300719, 0.7437381603891328, 0.7442812046358995, 0.7447429853721923, 0.7451953585881527, 0.7455842536192466, 0.7457799091514833, 0.7460949671618301, 0.7479381317118904, 0.7484231496042498, 0.7485209332107612, 0.7489132940350613, 0.7501800410534861, 0.7506062448287537, 0.7519576655062628, 0.7521916675286381, 0.7527602324308638, 0.7530172630513817, 0.7531189239971227, 0.7538976270133886, 0.7542656516518967, 0.7545928318997064, 0.7553926943730195, 0.7564333143243048, 0.756597662714793, 0.7568179674616046, 0.7569800791694212, 0.7572416730337281, 0.7579526687924429, 0.7592225078259971, 0.7595365973092976, 0.7595680962341955, 0.7597088507963641, 0.7609326048248761, 0.7615965913237999, 0.7617954892496176, 0.7624009037003568, 0.7639760452233711, 0.7640429632408469, 0.765248813879771, 0.7655278103085619, 0.7655539267386605, 0.7657073073155519, 0.7657697621254028, 0.7658593623447354, 0.766055078339458, 0.7663470023288758, 0.7664133442064084, 0.7673325695859763, 0.7683927539166149, 0.7684190194366555, 0.7684916294962503, 0.7688393024576092, 0.7697285635261047, 0.7697489214759077, 0.7704136786001269, 0.7707471424740322, 0.7707698797484479, 0.7714617466543078, 0.7723340562692853, 0.7726658203996475, 0.7730687202068967, 0.7733292075467065, 0.7737130461552625, 0.7738995103914904, 0.7743945540771618, 0.7748829257482917, 0.7748987345596995, 0.7755216448635028, 0.7773755385892163, 0.7778135222077405, 0.7780080584779979, 0.7781514712925933, 0.778576765740641, 0.7787319019503264, 0.7789863542438348, 0.7791664742134062, 0.779320712054238, 0.7800272841576654, 0.7803468381695269, 0.780996674264911, 0.7810459586128747, 0.7813513748258387, 0.7814461164149947, 0.7815897603443585, 0.7821698764403919, 0.782566516750569, 0.7826651441329482, 0.7829798879884534, 0.7829939284322924, 0.7831027211338285, 0.7832704145743478, 0.7835265741503598, 0.7836656483352726, 0.7839962396041258, 0.7851397729839422, 0.7854211968497454, 0.7862430378199742, 0.7863701844982469, 0.7865509351336004, 0.7868760290986077, 0.7875923528838987, 0.7879185809656778, 0.7889956236008694, 0.7897809482646645, 0.7907469498391294, 0.7912781535487777, 0.7919621466284361, 0.7919911591679207, 0.792127816164449, 0.7924666018939444, 0.7925373051602047, 0.792954044642139, 0.7935902799777165, 0.7936880523455437, 0.7941387066983607, 0.794435596150956, 0.794905657408873, 0.7952822942497813, 0.795443226087242, 0.7968095326792123, 0.7971614489908883, 0.7973687832406277, 0.7981561715946262, 0.7987199920805121, 0.7988759758296322, 0.7994786514361047, 0.7996459450260951, 0.7998280897714463, 0.7998875364762515, 0.800119405919345, 0.800425433856703, 0.80052997688357, 0.8007414429941495, 0.8009025428430033, 0.8009613639579662, 0.801880958217935, 0.8018817931642119, 0.8022352024761816, 0.8024863248274366, 0.8025404852867993, 0.8027519233370324, 0.8027810150272162, 0.803178201915209, 0.8032366678495861, 0.8033820973734624, 0.803545451235463, 0.8041128069523834, 0.804166136767419, 0.8043583282648488, 0.8046070437866985, 0.804878167930948, 0.8049449081765667, 0.805472898146863, 0.8055102430945924, 0.8055493604059113, 0.8065303102651824, 0.8075324731250259, 0.8077412643880251, 0.8087345235869297, 0.8087361495022602, 0.8091512168566065, 0.809444186736862, 0.8099055213780464, 0.8102803397453922, 0.8107552428559616, 0.8107784634482, 0.8110093097103809, 0.8114473694340236, 0.8122936365084831, 0.8123263582114478, 0.8129949465804458, 0.813156212973162, 0.8132174103485973, 0.8133502434927534, 0.8135756724291338, 0.8136593066446831, 0.8136800993617057, 0.8140027473663454, 0.8141303229933378, 0.8144018105755876, 0.8145101102058594, 0.815070577586771, 0.815153772822673, 0.8152929487162728, 0.8153307521150877, 0.8154965152423865, 0.8155568801057551, 0.8157571361738646, 0.8158088391734004, 0.8163354531448971, 0.8172569207076946, 0.8174625924438771, 0.8180586239736012, 0.8181115451714269, 0.8188254739106581, 0.8194924439558385, 0.8196088945409766, 0.8196466045485492, 0.8197207842753527, 0.8197819931742407, 0.8205664820436906, 0.820595400454324, 0.8206976661385539, 0.8208150287320395, 0.8209288691150338, 0.8209664218362859, 0.8209712730811488, 0.8209722508142344, 0.8213763325659119, 0.8213852036262769, 0.8214875847922506, 0.8215246476575486, 0.8216130852285379, 0.8216923054226706, 0.8218828788450379, 0.8219897347031045, 0.8220871332949833, 0.8221046612078231, 0.8228255883158105, 0.8229713013805668, 0.8230470656141321, 0.8239020427343992, 0.8241451287505481, 0.8243002887946442, 0.8245076947369117, 0.824523429843803, 0.8246910513538873, 0.8247646904381302, 0.8249239134348647, 0.8251533596323082, 0.8251839932727125, 0.8253720470931227, 0.8254884263704803, 0.8255267246410547, 0.8258433350262078, 0.8258461484835137, 0.8260897433137602, 0.8264174334381408, 0.8265187611531077, 0.8265812950419783, 0.826798066553577, 0.8268032804081901, 0.826990207052198, 0.8270557964318348, 0.8272935661141486, 0.827462702274204, 0.8274980797017314, 0.8281612370372792, 0.8283238378471498, 0.8285090955392441, 0.8287325120575958, 0.8288068499069708, 0.8288635895303368, 0.8291035687902532, 0.8291230978093465, 0.8292520682899518, 0.8293527644324653, 0.8294547653539408, 0.8297739331638786, 0.8299268221162396, 0.830349385204364, 0.8306127094861768, 0.8306464284369489, 0.8307273942543635, 0.8309506472239758, 0.8310811373043521, 0.8311001620721536, 0.8311275486014099, 0.8312518058991031, 0.8316127125187234, 0.8319196774938439, 0.8320987486791988, 0.8321743911231859, 0.8325303088090281, 0.8327666450913612, 0.832928561125999, 0.833159947634944, 0.8331630581431413, 0.8332008869809319, 0.833390856029018, 0.8335187210214337, 0.8335936431650706, 0.8338556313850911, 0.8340119567296367, 0.8345376316705069, 0.8349610759267605, 0.8350062856377357, 0.8351133444756077, 0.8351400203226901, 0.8351717721864724, 0.8354187008267158, 0.8357897112541863, 0.8359830830643472, 0.8363117092956115, 0.836315701688147, 0.836563448198992, 0.8366303463718089, 0.8367444565715763, 0.8368193136658623, 0.8368253034234434, 0.8368596891284763, 0.8371302045268685, 0.8377155179636261, 0.8378434222647695, 0.8378958517087074, 0.8382062721376756, 0.8382623881964519, 0.8384413737659587, 0.8384884926520767, 0.8384997173003437, 0.8385551879072821, 0.8386442990797603, 0.8387377506977064, 0.8391913598979956, 0.8393594596510313, 0.8393942563278147, 0.8395290922156592, 0.8396649010041929, 0.8397911451080177, 0.8398529974597916, 0.8403257943451061, 0.8408580464136605, 0.8415823900827769, 0.8417242493998728, 0.8419791355591059, 0.8420578829975182, 0.8421064309959974, 0.8422032505548623, 0.8430411738430843, 0.8430708746622296, 0.8430921331485852, 0.8431126163375798, 0.8431269081020097, 0.843193907010503, 0.8435597332022584, 0.8436364734785564, 0.8438390638216229, 0.843840735103161, 0.8438428597378985, 0.8441885383352349, 0.84449690050741, 0.8448124288462099, 0.8448266108902197, 0.8448629526309642, 0.8449583041544663, 0.8457229137992535, 0.8457529815885535, 0.8457531418777064, 0.8458602257211606, 0.845989080208456, 0.8460749620512018, 0.8462255430117499, 0.8466913365386814, 0.846936571422461, 0.847064469250011, 0.8471434282792409, 0.8472579559993441, 0.8474660136623108, 0.8474803654782045, 0.8475575618740968, 0.8476904729905124, 0.8477488540715744, 0.847881343998382, 0.848048624047718, 0.8481425738316124, 0.8482853200353464, 0.8482995167612462, 0.8485696018859935, 0.8487285465634614, 0.8488703624105336, 0.8492284634505455, 0.849423915559495, 0.849534964236416, 0.8497070408231402, 0.850383417420627, 0.8505408310209062, 0.8506098346415139, 0.8511510891905767, 0.8514460150859233, 0.8514665175238559, 0.8515551423531876, 0.851646659607841, 0.8516484598454379, 0.851729200426964, 0.8517505771792905, 0.8518390118999979, 0.8519248051232621, 0.8520192260482736, 0.8521203920814477, 0.8521754481746504, 0.8522392839036516, 0.8522446509744317, 0.8523743448167858, 0.8523928639245248, 0.8524525751259496, 0.8526742997723206, 0.853063627261575, 0.8531639955802628, 0.8531929515016162, 0.8535607150406738, 0.8536579486946164, 0.8537411751421201, 0.853768591141381, 0.8542477931899125, 0.8542890705552005, 0.8543189642056412, 0.8543933327803044, 0.8544520148187585, 0.8546855811584642, 0.8549316520567539, 0.8549580276708472, 0.8550140622181057, 0.8550333043534348, 0.8550518151557853, 0.8550533810451193, 0.8550735292064009, 0.8551371673048941, 0.8552310125118952, 0.855447699962244, 0.8558571545878222, 0.8560564033768365, 0.8563473235552794, 0.8564082544858361, 0.8566084368620377, 0.8566427439514143, 0.8571724682703918, 0.8576026777926194, 0.8576996669971807, 0.8581835283088431, 0.8584766681167924, 0.8586545618578021, 0.8587189887606432, 0.8587661343814199, 0.8588256093260273, 0.8588869670748489, 0.8589223752038775, 0.8589946122623278, 0.8592555072356453, 0.8594220766582493, 0.85960817437115, 0.8596261420231839, 0.8597335793643961, 0.8598257474316949, 0.8599288689703741, 0.8600570202862348, 0.8600605695842933, 0.860145305434656, 0.8601983304560357, 0.8602363499370669, 0.8603379634323767, 0.8604070763102323, 0.8604511482358683, 0.8605476828167065, 0.8606848847745475, 0.861428785223191, 0.8615125701261458, 0.8615199542375307, 0.8615474733396296, 0.8615529704656633, 0.8615629445189658, 0.8618753997675506, 0.8618897375881934, 0.8619939856482056, 0.862094709320051, 0.8621547385488705, 0.8623220165756099, 0.8623439802298465, 0.862413725074895, 0.8626124921405864, 0.8626278228638521, 0.8627229020292319, 0.8627309589296465, 0.8627555646558611, 0.8629542710824616, 0.8630272116343318, 0.8630444040862776, 0.8633178555505644, 0.8636176856983364, 0.8637255244522711, 0.8639126060660063, 0.8640023110075192, 0.8641958052370384, 0.8642081438671824, 0.8642219710474388, 0.8644324681878787, 0.8645305228023561, 0.8645850870967132, 0.8648511735551858, 0.8650249148609157, 0.8651755868844336, 0.8654450325998062, 0.8654699441374877, 0.8658507913968333, 0.8659417357106962, 0.8661005228958476, 0.8661460631593458, 0.8663380433942391, 0.8663864762034681, 0.8664609671987252, 0.8667866536992176, 0.8669763225940144, 0.8671411148518195, 0.8671724186069768, 0.8671796282122554, 0.8671926426613827, 0.867296591584493, 0.8673571210390705, 0.8674403142998919, 0.8674715621673922, 0.867516534083199, 0.8678966462798079, 0.8679526053316644, 0.868090332456978, 0.8681003045035149, 0.8683507238934014, 0.8685814852689443, 0.8686181032476649, 0.8689046040631008, 0.8690255637517922, 0.8690285216562849, 0.8691442789438327, 0.8694535055698014, 0.8695476013135951, 0.869596071147928, 0.8697730076676476, 0.8697761575568383, 0.8698127767861631, 0.8700576510322591, 0.8701711363199597, 0.8703214489368698, 0.8705637705464886, 0.8705923614317781, 0.8708471133494977, 0.8710170665363144, 0.871020735968129, 0.871132167878188, 0.8711717707931079, 0.871208398423848, 0.8714189117458847, 0.8716212384234885, 0.8718260671268865, 0.8718567664219965, 0.8719748622558846, 0.8722489396451172, 0.8723501995897022, 0.8725468777005524, 0.8726269526814884, 0.8727001112213236, 0.8728559304800054, 0.8728628378185895, 0.8729119220716078, 0.8730584632643719, 0.8731663384015685, 0.8731998149452969, 0.8735236803698896, 0.8737948777115973, 0.8740452402975779, 0.874121900173617, 0.8745843063416424, 0.874868710110581, 0.8749796863570771, 0.8753713494435339, 0.8755595551612355, 0.8756509908486403, 0.8756625594281028, 0.8757413236720176, 0.8757943696175656, 0.875814581225012, 0.8758584924415656, 0.8761230919027099, 0.8762167032359198, 0.8762422386429601, 0.8763051432888476, 0.8763765377299685, 0.8767141788153984, 0.876765948222858, 0.8767719458590804, 0.8772152592921032, 0.8772474121678439, 0.8772748141074447, 0.8773327749128876, 0.8773329851000832, 0.8773384903414378, 0.8774398892478114, 0.8774416245477022, 0.877478772590188, 0.8775353505247776, 0.8777532702155948, 0.8778067493920788, 0.8779127291671542, 0.878075532094166, 0.8781464403970448, 0.8783285405899106, 0.8783433153971447, 0.8783446698937957, 0.8783827193741066, 0.8783855695910457, 0.8788843897409315, 0.8789492354634206, 0.8790456294984013, 0.8790927183613559, 0.8793460905453334, 0.8794258076651098, 0.8794498146349273, 0.879476712019643, 0.8796954001080769, 0.8798609640366604, 0.8798857442242926, 0.8799288416563251, 0.8801926600925445, 0.8802038092835958, 0.8802600831533601, 0.8804229179208377, 0.8804549111270558, 0.8804774984035258, 0.8804911706011126, 0.880559193067794, 0.8806094933708324, 0.8807275136118302, 0.8808940008405702, 0.8809120424720162, 0.8812535806055259, 0.881672530830342, 0.8817652796690296, 0.8818059581883865, 0.8819416633163023, 0.8820342792837103, 0.8820855484527885, 0.8821599267690183, 0.8821874622189333, 0.882487146945981, 0.8826444035889242, 0.8827879385749431, 0.8828143554630399, 0.8829043405423627, 0.8829277623062332, 0.8829580069826918, 0.8829613598509548, 0.8830050265748722, 0.8830299804629234, 0.88303290294148, 0.8830500576302704, 0.8831191676882719, 0.8832249242429715, 0.8832316886009118, 0.883250173697329, 0.8832551791677244, 0.883288436824215, 0.8834185669008553, 0.8835048191863971, 0.8836723347845887, 0.8837663812071997, 0.8837919457443308, 0.8838466786695739, 0.8840977465266028, 0.884232446509382, 0.8842472793993331, 0.8844136574239463, 0.8845627863691622, 0.8845659730117176, 0.8845935339201634, 0.8846080946370232, 0.8846386739486721, 0.8847675864758693, 0.8847994955900531, 0.8848899108393733, 0.8849751812335801, 0.8850785150672481, 0.8852626136033944, 0.8854435813553121, 0.8856496582918877, 0.8857497680300442, 0.8858212019250566, 0.8859047507851383, 0.8860404546338404, 0.8862323709025474, 0.8864244065185929, 0.886463917276498, 0.8868846985678767, 0.8870016080176161, 0.8871199629010255, 0.8872637283341253, 0.8872936330908403, 0.887322449472402, 0.8873858537303716, 0.8873996570053484, 0.887444245352988, 0.8874601310503893, 0.8874820448333983, 0.8876038120829571, 0.8881166856156809, 0.8882140185359398, 0.8882346407902375, 0.8882756093225708, 0.8883302629677977, 0.888463519237417, 0.888481886474997, 0.8886542807304799, 0.8887488510650841, 0.8888526402181276, 0.888855786557928, 0.888893427922418, 0.8889181785832876, 0.8889303534107227, 0.8889473683723651, 0.8890036381294333, 0.8890551709292727, 0.8891020465843532, 0.8892103335527526, 0.8892227036642887, 0.8892478158562745, 0.8892598155821404, 0.8895518599307326, 0.8895771190355424, 0.88959880473276, 0.8896405941102915, 0.8898622599048108, 0.8902521202227921, 0.8903170080844451, 0.8905849863652436, 0.8905916545782065, 0.8906109431351293, 0.8907115510548655, 0.890812682107368, 0.8908805025120269, 0.8910119533986309, 0.8910286664521676, 0.8911215020722817, 0.891379262754116, 0.8914091210611135, 0.8914761102787024, 0.8914815611924424, 0.891496160889268, 0.8915037765611087, 0.8915599192509149, 0.891725491120668, 0.8917491657446329, 0.8918293661154814, 0.8919685217200014, 0.8921665833761373, 0.8923912417147744, 0.892399561591746, 0.8924403106657421, 0.8924777766471788, 0.8925036896297679, 0.892529478284844, 0.8925460351349442, 0.892749869886965, 0.8927860763932501, 0.8929488067549719, 0.8929495521054378, 0.8930604985034907, 0.8932296065322608, 0.8933822004109739, 0.8935295847339183, 0.8935665108459308, 0.8936322503388232, 0.8936432517754117, 0.8938946637046092, 0.8939525659512384, 0.8939949292873447, 0.894023982579738, 0.8940607807663973, 0.8940694205572901, 0.8941392965677057, 0.8942342369708224, 0.8942483395325678, 0.8944069164586987, 0.8944828316865192, 0.8946211324949973, 0.8946585672413323, 0.8946785329379556, 0.8947590723278964, 0.8947958316387663, 0.8948013829818855, 0.894918553293347, 0.8950261945294239, 0.8952177771195697, 0.8952641124951516, 0.8952674356282746, 0.8956174747878225, 0.8956954071006356, 0.8958048501219713, 0.8958599252216832, 0.8959143514790416, 0.8961117375720287, 0.8962201621006458, 0.8963259657841145, 0.896349468316313, 0.8964027238004669, 0.8964451213468515, 0.8965198863825566, 0.8966476248858074, 0.8966636661952564, 0.8967365991436369, 0.8967750000811775, 0.8968138397813644, 0.8968191477045411, 0.8968217164527688, 0.8969295358814972, 0.8971347294813835, 0.8972181196272289, 0.8972594984633091, 0.8975550345896368, 0.897558941947983, 0.897610319136376, 0.8976930141384273, 0.8977251347401053, 0.8977960107155104, 0.8977963947005283, 0.8978347755598951, 0.8978449427229673, 0.8978863970981764, 0.8980041081507983, 0.8980423136546775, 0.8980519815839629, 0.8980598467204235, 0.8981044325569822, 0.8981983596315155, 0.8983188654192134, 0.8983452542218292, 0.8985736016062156, 0.8986055539023269, 0.8987433206715298, 0.8988314331769385, 0.8988549303556641, 0.8989381920506305, 0.8989600414399485, 0.8990795155654459, 0.8991794366385629, 0.8992093492595279, 0.8993027280224437, 0.8994444501857431, 0.8995000901636462, 0.8997790617534797, 0.8998480999677168, 0.8999610859270193, 0.8999692810741763, 0.9000388032011292, 0.9000720172153309, 0.9000856995499733, 0.9001614248549057, 0.9002195686807954, 0.9002908256501653, 0.9004122385536988, 0.9004412796500414, 0.9004502828892069, 0.9004728228315397, 0.9004757997701367, 0.9005457815059218, 0.9005468711606589, 0.9006428629571317, 0.9006871818926895, 0.9007266073202196, 0.9007376322875325, 0.9007382580947069, 0.9007607754169594, 0.9007626245325914, 0.9008099330261768, 0.9009005848129676, 0.9009437720104929, 0.9011054900411003, 0.9011154860191972, 0.9011448305340376, 0.9012206719142875, 0.9012393578943219, 0.9015221806644996, 0.9016239408568828, 0.901657153220027, 0.901683829302504, 0.9017981113531092, 0.9018247003513116, 0.9018759409237004, 0.9020914621205122, 0.9021044724259978, 0.9022714027329044, 0.9022750802197369, 0.9023648130040518, 0.902435775691739, 0.9025826946279851, 0.9025940031700174, 0.9026267104445234, 0.9026942722464084, 0.9027046003288082, 0.9028698545930436, 0.9029402592250408, 0.9030579245716547, 0.9030978843885821, 0.9031677740091987, 0.9033349070965698, 0.9034012288885386, 0.9034736302878676, 0.9034741782299665, 0.9036995615591406, 0.9037024488766819, 0.9037591106904824, 0.9037955305794874, 0.903972603619466, 0.9040594890470107, 0.9040763120366301, 0.9041013719893668, 0.9041286279900042, 0.9042175408332911, 0.90447787969298, 0.9045042808701983, 0.9045685236276965, 0.9046821520783787, 0.9047368984754401, 0.9048499918178092, 0.9048608388713678, 0.9049982025598596, 0.9051393180787525, 0.9052452705228552, 0.905275631599094, 0.9053881101582391, 0.9054001825179095, 0.9054309681155831, 0.9054431912740109, 0.9055170399650159, 0.9055793345914898, 0.9056106498666036, 0.9056186826386715, 0.9057356378652912, 0.9058555096104104, 0.9059418953976015, 0.9059463712209235, 0.9060466057226425, 0.9060528256405235, 0.9061240614877509, 0.906132517408587, 0.906231042654213, 0.9063188490568553, 0.9063613937952232, 0.9064147986880848, 0.9065028044865927, 0.9065482078990663, 0.9065599554021463, 0.9066978082525229, 0.9067376298762917, 0.9068204283867631, 0.9070047153220474, 0.9071037762348315, 0.907465982357867, 0.9077487420231891, 0.9077880318360859, 0.9078083078164639, 0.9078810113826936, 0.907955926429808, 0.9080373995050295, 0.9081078985935409, 0.9081354708558422, 0.9083481470494945, 0.908367739956524, 0.9084629765316705, 0.9084711652924513, 0.9085213771332642, 0.9085533555365508, 0.9086020044193278, 0.9086586950248347, 0.9087150112200758, 0.9087385916563602, 0.9088101393221153, 0.9088410815491782, 0.9088492373585689, 0.9089275869917988, 0.9089407430713161, 0.9092443723746095, 0.90924884783423, 0.9092610094277958, 0.9093592547427944, 0.9094967648181569, 0.9096098666767337, 0.9096633808741285, 0.9097119637258827, 0.9097323998257035, 0.9098803474330169, 0.9099202388840516, 0.9099796251568267, 0.9099996259373202, 0.9100059541849868, 0.9100203167434419, 0.9100314938857268, 0.9101525983432189, 0.9103405251615045, 0.9104958436124277, 0.9107522884246438, 0.9107797324672775, 0.9109776765663413, 0.9110602547407234, 0.9110905668100749, 0.9113057265450083, 0.9113089609455937, 0.9113466070753563, 0.9113528593468159, 0.9117606507325793, 0.9119284291540675, 0.9119764551669346, 0.9120159484948736, 0.9120344916323376, 0.912057837800758, 0.9120641101235701, 0.9123868862227376, 0.9124573339397962, 0.9124626028873022, 0.9125310626501226, 0.9125948685831018, 0.912597388067907, 0.9126290441047026, 0.9126781086721767, 0.9128038300088753, 0.9128139963530915, 0.9128830440590112, 0.9128902514684728, 0.912998211187383, 0.913019621789437, 0.9131708021018504, 0.9131936261682692, 0.9132102927095807, 0.913243982515692, 0.9132990809463988, 0.9136180990197521, 0.9136817189755134, 0.9138034009328343, 0.9138354050660171, 0.9138585481922862, 0.9138675176930939, 0.9139422981453568, 0.9139853310609833, 0.9140316178913701, 0.9140499756234531, 0.9140647980291079, 0.9141715110217624, 0.9141779279127611, 0.9142000707198968, 0.9143038391185485, 0.9143902324811213, 0.9144086031301988, 0.9145653190866959, 0.9145962709589848, 0.914597687918693, 0.9146375090311705, 0.9146896548400896, 0.9148465678936695, 0.9148594656136783, 0.9148630063988625, 0.9149083480989422, 0.9150206169089473, 0.915096484714436, 0.9151144140421906, 0.9151466130097989, 0.9152651736202579, 0.9153119886773543, 0.9153302697758469, 0.9153687060462802, 0.9153843533067697, 0.9154367338451188, 0.9154601108158288, 0.9155953729869123, 0.9156563245102599, 0.91567560766209, 0.9156945328141255, 0.915745842294244, 0.9158874553654353, 0.9159410400760672, 0.9160866458939129, 0.9161715672979476, 0.9161800722268502, 0.9161811811131799, 0.9161892886386245, 0.9162544059013545, 0.9162577072622582, 0.9162875264311053, 0.9162993336657207, 0.9163591056782017, 0.9164649255216565, 0.9165469642195666, 0.9165703339331759, 0.9166798045186745, 0.9167337871357316, 0.9168075059824775, 0.9168234835154889, 0.9168501595820582, 0.9168760637950073, 0.9169136406508936, 0.9169227502472277, 0.9169502169987307, 0.916974607542693, 0.9171198543336381, 0.9171727452749969, 0.917326745125489, 0.9173473771840404, 0.9173824371308044, 0.9174271118072956, 0.9175558736221503, 0.9175780875540754, 0.917587872286784, 0.9176136221934413, 0.9176535735001363, 0.91767377891954, 0.9177285749769842, 0.9178147208551741, 0.9178172223709054, 0.9179062379703018, 0.9179156108891645, 0.9179837855604046, 0.9180315553000626, 0.9180347208400724, 0.9180905114042371, 0.9181330011182915, 0.9182204894766782, 0.9182355994745618, 0.9183167788633491, 0.9183996137668646, 0.918449663946144, 0.9184557186008798, 0.9185364000280102, 0.9186676374429946, 0.9187672786419792, 0.918767446518151, 0.918934878649511, 0.9189797722401686, 0.9190303716051769, 0.919152044194216, 0.9191529849910947, 0.9191744964776211, 0.9192138592890222, 0.9192904446608134, 0.9193271507765928, 0.9194467236219448, 0.919457218360201, 0.9195177778123832, 0.9195873901272299, 0.9197659529749896, 0.919867347679213, 0.9198713314488184, 0.9199892423543606, 0.9201220323056573, 0.9202946355617942, 0.9205148626370413, 0.9206331522943817, 0.9206898659528738, 0.9207353987415346, 0.9207552179947324, 0.9207633271668126, 0.9208390455117322, 0.9208999092236605, 0.9209860934242784, 0.9210079332628549, 0.9210732087526459, 0.9210828809167785, 0.9211590504399502, 0.9212604632171002, 0.9212880364115608, 0.9212894142264118, 0.9213066345966949, 0.921334316976834, 0.9213549800438445, 0.9214331191356584, 0.9214572130115907, 0.9216264735805703, 0.9216606416692452, 0.9217272121816943, 0.9217370480581276, 0.9217603582259055, 0.9218092641309668, 0.9218456675393653, 0.9219566174709273, 0.9220928272387294, 0.922121054709479, 0.9221557096307486, 0.9221923774348654, 0.9223025691185585, 0.9224059558773997, 0.9224277480060452, 0.9224761744414562, 0.9225060312406296, 0.9225503348463597, 0.9226224092855372, 0.9226229275601471, 0.9226534398368651, 0.922751295124646, 0.9229207051670213, 0.922944594727394, 0.9229589254990763, 0.9230495284874215, 0.9231174816979864, 0.9231914519195034, 0.9232393596216015, 0.9232632076891762, 0.923266602210552, 0.9232779312973577, 0.9232819773504558, 0.923355330069161, 0.9233619070866399, 0.923362302365673, 0.9235060970427544, 0.9235439570679821, 0.9236346869692807, 0.9236406172183482, 0.9237592790613303, 0.9238518831079399, 0.9238809873085232, 0.9238919926082535, 0.923899224205632, 0.9240276657025801, 0.9241398310261876, 0.9241549353702242, 0.9242563621817611, 0.9242923885489365, 0.9243705394641989, 0.9243970690408403, 0.9244057685110945, 0.9244162630554821, 0.9244808671479874, 0.9244809988070406, 0.9245631861621987, 0.9245825257194572, 0.9246271065807509, 0.9246349358757443, 0.9246575314492708, 0.9246944896802959, 0.9247601680006254, 0.9248259976653378, 0.92486212946155, 0.9248832682364074, 0.9248923219109244, 0.9249157251519536, 0.9250064313789236, 0.9252435345054364, 0.9253665353531126, 0.925389001139673, 0.9255391079278718, 0.9255454308606011, 0.925574072576722, 0.9255986984174479, 0.9256570498574123, 0.9256902319700988, 0.9258626595530509, 0.9258752928020438, 0.9259063392353175, 0.9259249855968212, 0.9259678541889714, 0.9260857719044605, 0.9261549064818531, 0.926211506564194, 0.9262640973046408, 0.9262649069809045, 0.9263773248675203, 0.9263912656375478, 0.9263940482134312, 0.9264358963497076, 0.92649961332166, 0.9265206954452517, 0.9265221762586698, 0.9265646437097679, 0.9266094344011458, 0.9266485206847053, 0.9266995455713085, 0.9267634693259423, 0.9267995825670466, 0.9268748450605602, 0.9268822977929032, 0.9269003616158398, 0.9269596808999571, 0.9269712733914836, 0.9269897569241405, 0.9270605803627152, 0.9270663784915757, 0.927101315008059, 0.9271129840967497, 0.9271310752306654, 0.9271537285982792, 0.9272194428053413, 0.9272868485782815, 0.9273095783186346, 0.9275185443554994, 0.9276937771446437, 0.9277637935775231, 0.9278219930414159, 0.9278511331507228, 0.9278875940219509, 0.9279124520008791, 0.9279212417192433, 0.9279361321342333, 0.927963014249193, 0.9281084451301098, 0.9281918065350798, 0.9282585590257395, 0.9282785802858147, 0.9283955117196794, 0.9286098052803764, 0.9286687875602879, 0.9287078408493729, 0.9289690052001526, 0.9290089744627802, 0.9290267967608499, 0.9291353708796242, 0.9291521675741021, 0.9291615808250928, 0.9291977126440507, 0.9292494012041016, 0.929385701784271, 0.9295059735216477, 0.9295349284233149, 0.9295498633797018, 0.9296184475231342, 0.9296203829882603, 0.9296277810575698, 0.9296562551797186, 0.9297314246811162, 0.9297758671296574, 0.9297839192630352, 0.92980358612396, 0.9298285625036453, 0.9299018138965351, 0.9299056138163821, 0.9299218261639064, 0.9299485823850021, 0.9299902466975846, 0.930017150570915, 0.9301968696894565, 0.9302516841208468, 0.9302935573502935, 0.9304340309983403, 0.9306109938136107, 0.9306762147750541, 0.9307994891564056, 0.9308729189622968, 0.9310347009604556, 0.931106021696619, 0.9311106204579558, 0.931157929549993, 0.9311921555717959, 0.9312417048316896, 0.9314488239299068, 0.9315613879415705, 0.9315836171188018, 0.9315873877517838, 0.9316748041595297, 0.9317051940879993, 0.9317869267001108, 0.9318575074097334, 0.9319165914660535, 0.9319855523038428, 0.9320245335421659, 0.9321362303035037, 0.932140060861476, 0.93215048525173, 0.9323772995431827, 0.9324579185404432, 0.932460298069769, 0.9324960076416008, 0.9325217685455104, 0.9325236599111049, 0.9325354281341212, 0.9326061108698472, 0.9326561835467059, 0.9327159097269406, 0.9327174437546235, 0.9327273174350744, 0.9327328417678447, 0.9328103713254874, 0.9329766821725355, 0.9329833520564516, 0.933020498239576, 0.9331119203789336, 0.9331476750984384, 0.9332282420423446, 0.9332848644393155, 0.9333243965293074, 0.9333306650531724, 0.9334255484577224, 0.9334409713460902, 0.9335498828511603, 0.9337449443701149, 0.9338343511172335, 0.9338644629851971, 0.9339248527889532, 0.9339374337397124, 0.9340854149486573, 0.9341915863851433, 0.9342946133189275, 0.9343124850923081, 0.934365763713791, 0.9343741396610236, 0.9344265765222987, 0.9344528786994525, 0.9344654013270742, 0.9345932266344102, 0.934738075631405, 0.9347405769279488, 0.9348338736370565, 0.9349081568044834, 0.9349182114895761, 0.9349617104771487, 0.9350923761390502, 0.9351222243564868, 0.935282653511518, 0.9352866637979184, 0.9353547474926679, 0.9354462849071643, 0.9355154424015221, 0.9356130337533143, 0.9356136557768774, 0.9356313491988106, 0.9356454007883233, 0.9356787996063758, 0.9357177916351579, 0.9358979054551636, 0.9359041913125861, 0.9359065020037746, 0.9359296166839045, 0.9359340919929594, 0.9361946972096088, 0.9362135153764068, 0.9362156805600303, 0.9363027823494313, 0.9363348616472837, 0.9364163826286188, 0.9364420511248883, 0.9364590521392439, 0.9365187314977202, 0.9366039670616986, 0.9366201865039139, 0.9366489719641838, 0.9366522691916727, 0.9366829868095118, 0.9366883674672123, 0.936702634778223, 0.9367680862090212, 0.936852221584458, 0.9369591130243129, 0.936965884983951, 0.9370024492330729, 0.9370453368393918, 0.9370666526899502, 0.9371069309151607, 0.9371403999573211, 0.937181241417333, 0.9372206694000519, 0.9372297449103427, 0.937230707387129, 0.9372701057179565, 0.9372935405349916, 0.9373159722644033, 0.9373524163820713, 0.9375961960335206, 0.937609750580159, 0.9376309262009547, 0.9376618830032115, 0.9376949469950753, 0.9378268203893508, 0.9378305532479374, 0.9379047503119877, 0.937906504209415, 0.9379100856025923, 0.9379893704203883, 0.9380409928964344, 0.9380904583049394, 0.9381151393152198, 0.9381290338344522, 0.938138788208977, 0.9381583453123916, 0.9383071157714156, 0.9383313572427526, 0.938370612396714, 0.9384239951983413, 0.9384613560085401, 0.9385110665513932, 0.9385267946665669, 0.9386417636192435, 0.938657925076354, 0.9386619241991379, 0.9387235606465562, 0.9388097842140812, 0.9388133623607573, 0.9388576937205138, 0.9388720308537516, 0.9389004833960352, 0.9389060516233563, 0.9389220032479353, 0.9389317375184286, 0.9389402904636847, 0.9389539025809797, 0.9390228451662381, 0.939027407338353, 0.9391676328880739, 0.9391876482509334, 0.9392262710025552, 0.9392669070722074, 0.9392673531590474, 0.9393603514353147, 0.939415673026294, 0.9394243067296919, 0.9394790851652879, 0.9395215153552663, 0.9395402188228812, 0.9397224129228788, 0.9397251461747272, 0.9397557360785325, 0.9399013116296066, 0.9399344584766975, 0.9399465490810176, 0.9399569099061754, 0.9400363369943089, 0.9400789146722894, 0.9401138952145726, 0.9401790372391224, 0.9401862805967593, 0.9402562979287923, 0.9402617473787589, 0.9404537315999687, 0.9404565162943547, 0.9405116270634916, 0.9405371756655296, 0.9405582318291096, 0.9405657192735919, 0.9406295219152401, 0.9406411415291467, 0.9407024197455827, 0.9407180522591954, 0.9410704851329037, 0.9413173431404508, 0.9413340971990988, 0.9413360009282753, 0.9413753139827975, 0.9414022031981187, 0.9414101962282698, 0.941422397351747, 0.9414279564087006, 0.9414311097739636, 0.941431482423101, 0.9415008420485369, 0.9415142673556949, 0.9415373638248471, 0.9416078810949121, 0.9416132468781443, 0.9416161532058618, 0.9416180224856437, 0.9416459275514634, 0.941661606815243, 0.941669107501812, 0.9417067805385162, 0.9417760212256703, 0.941859657153247, 0.9418658312281348, 0.941877795416916, 0.9418920802148627, 0.9418935790538405, 0.9419482925962973, 0.9421608610690738, 0.9421848749046317, 0.9422547448407306, 0.9422615231622359, 0.94227945331034, 0.9423072381216497, 0.9423181893922127, 0.9423462117722661, 0.9423487859963646, 0.942352854458562, 0.94237362012863, 0.9423755403736992, 0.9424411686628726, 0.9424983708577938, 0.9425086955564674, 0.9425578744139399, 0.942590005558788, 0.94259573345991, 0.9427293538213022, 0.9427536961037906, 0.9428940373685607, 0.9429076261454205, 0.9429363731440975, 0.9429379271881771, 0.9431439707920684, 0.9431912215338146, 0.9431938216146577, 0.9432715100616383, 0.9432875267652547, 0.9433577237678429, 0.9434114251933284, 0.9434706748232264, 0.9434727990142591, 0.9434934097674395, 0.943528707274506, 0.9435555071565628, 0.943624784382602, 0.9436478972218539, 0.9436846612179972, 0.9437100213727941, 0.9437593464666402, 0.9437621947084807, 0.9437996628322776, 0.9438000381913929, 0.9438160123173852, 0.9438233005361073, 0.9438556063556639, 0.9438801090026008, 0.9439123195745308, 0.9440009614557285, 0.944109511492799, 0.9441306253358421, 0.9441359366323264, 0.9441829684029747, 0.9442147525628383, 0.9442408425925216, 0.9442487075897544, 0.9443415536926741, 0.9444121122756416, 0.9444279436939813, 0.9444509946778911, 0.9445409146930298, 0.9445716467992533, 0.9445761925455458, 0.9445971870946761, 0.9446032416169168, 0.9446497899888033, 0.9446604319334156, 0.9447052494607237, 0.9447096023600537, 0.9447534604842054, 0.9447758635183001, 0.9448147512378657, 0.944859454722575, 0.944868197776392, 0.9448921414957832, 0.9449075442460597, 0.9449183693090303, 0.9451000752066168, 0.9451108056443042, 0.9451187954190929, 0.945142884899279, 0.9451447624172248, 0.9451636340924618, 0.945193012067981, 0.9451941760859306, 0.9452519036137681, 0.9453510543610378, 0.9454395577178988, 0.9454887916618575, 0.9455001725376353, 0.9455231275015387, 0.9455323462448934, 0.9455736948432895, 0.9457346651553141, 0.9457505213681457, 0.9458627771813689, 0.9458899349223852, 0.9459213245566895, 0.9459434190541689, 0.9459445464555631, 0.9460167485245311, 0.9460301899977552, 0.9460481328324981, 0.946165482255151, 0.9461740877543577, 0.9462523176557572, 0.9462548733866685, 0.9462910727170643, 0.9462942906088875, 0.9463651428802442, 0.9464175186648988, 0.9464332197556466, 0.94646057034641, 0.946469800835181, 0.9464813826047876, 0.9464897658500141, 0.9465592687291031, 0.9465832803284131, 0.9466561544893076, 0.9466935542520798, 0.9467428004596705, 0.9467537341797686, 0.9468144561919726, 0.9470111539004638, 0.9471002473393405, 0.9471214138796881, 0.9472085616682465, 0.9473073129780601, 0.94734878398838, 0.9473600039800969, 0.9473718328867605, 0.947373500036016, 0.9474397893963702, 0.9475109435299933, 0.9475146205269906, 0.9476165741353636, 0.9476801314880401, 0.9477337215841715, 0.9478179535105817, 0.9478299456260343, 0.9478941024365669, 0.9478988458495233, 0.9479441614983666, 0.9479948056872946, 0.9480187478163714, 0.9480592531301356, 0.9480710222485973, 0.9481048822438406, 0.9481136980179722, 0.948132465226175, 0.9481346226833854, 0.9482280626222249, 0.9482361735481449, 0.948252315906543, 0.9482670847502951, 0.9482787638129951, 0.9482860692639062, 0.948293737601891, 0.9483794286853134, 0.9483832232644579, 0.9484111776318371, 0.9484156200847355, 0.9485002434633116, 0.9485175309692413, 0.9485612402062782, 0.9485758446133143, 0.9485783090399971, 0.9485784025061568, 0.9486793740037113, 0.9486902993908967, 0.9488048759878364, 0.9488347099359307, 0.9488446332893143, 0.9489137444554898, 0.9489270713420627, 0.9489387352162797, 0.9489421794219866, 0.9489839338770943, 0.9489842172253683, 0.9490490441496506, 0.9490826604199278, 0.9491115546437787, 0.9491866404178078, 0.9491927333946137, 0.9492059743701036, 0.9492509418517832, 0.9492523976141073, 0.9492552212369378, 0.9492690325054732, 0.9492904036987044, 0.9493591986385834, 0.9493994827039643, 0.9494568037466959, 0.9494662356876357, 0.9495125163223964, 0.9495421702940958, 0.9495462914292447, 0.949570251309486, 0.949590682523674, 0.9496105043060183, 0.9496182669469277, 0.9496298828690654, 0.9496412669871764, 0.9497518614879984, 0.9497745244537651, 0.949775606927486, 0.9497952488046132, 0.9498348640808622, 0.9498957239525171, 0.9499192370505889, 0.9499367337111945, 0.9499609183868244, 0.9499937612021143, 0.9500310323766716, 0.9500733331506398, 0.9501090617671407, 0.9501457201804694, 0.9502048845459379, 0.9502063967033537, 0.9502882538252763, 0.9503381662192414, 0.9504929733039085, 0.9505080847915811, 0.9505128105684973, 0.9505160147944156, 0.9505655334139678, 0.9505980772927142, 0.9506462891261069, 0.9506810938205985, 0.9507365802771872, 0.9507878922554701, 0.9508175447255304, 0.9508245999459766, 0.9509061049439146, 0.9509333970747436, 0.9509572052312693, 0.9509736911241665, 0.9509803131054506, 0.9509970328970224, 0.9510461410180503, 0.9511351513695715, 0.9511705438617625, 0.9512291878695626, 0.9512294977892088, 0.9512357668750443, 0.9512985077393585, 0.9513418420661733, 0.9513430004941039, 0.9513544973565933, 0.9514598869179248, 0.9515064120451906, 0.9515161940328318, 0.9515342736977046, 0.9515473789212706, 0.9515511793622367, 0.9515759326628794, 0.9515845945781201, 0.9516010610594029, 0.9516033305260317, 0.9516188080319371, 0.9516487382744279, 0.9516576816757361, 0.9516865090790779, 0.951699305051519, 0.9517177029242248, 0.9517476009404991, 0.9517871351744204, 0.951812467397876, 0.951854635787983, 0.9518848074526804, 0.9519730711703829, 0.9520033014139114, 0.9520534927901947, 0.9520988108021257, 0.9521385316998003, 0.9521534321826955, 0.9521652346271906, 0.9521691297766812, 0.9521766944332695, 0.952240311613236, 0.9522524451572408, 0.9523301014617734, 0.9523935699147413, 0.9525254307298747, 0.9525353665658551, 0.9525430696747605, 0.9525503571243131, 0.9525697670330804, 0.9526100973426211, 0.9526607505859302, 0.9526650950792953, 0.9526826245513823, 0.9526867565840529, 0.9527447351063298, 0.9527613539015408, 0.9527699803175304, 0.9529809696931651, 0.9530557922559412, 0.953068211837877, 0.9532253656951758, 0.953275746208201, 0.9533069951240101, 0.9533428924826288, 0.9533703292591633, 0.9534280352832333, 0.9535156784934063, 0.953580738855863, 0.9535920822439806, 0.953638049666413, 0.9536592124652865, 0.9537094974492443, 0.9537252536486877, 0.9537679201281832, 0.9537686379389108, 0.9538683876101214, 0.9539274143355211, 0.9540246852771128, 0.9540780400812244, 0.9540841371138334, 0.9540902460190261, 0.954128530258453, 0.9542046701367418, 0.9542342293893201, 0.9542622454296953, 0.9542763607491802, 0.9542927343621428, 0.9543099230204287, 0.9543286081485447, 0.9543569940594305, 0.9544676454320201, 0.9544829853707294, 0.9544834511355129, 0.9545130563391785, 0.9545200248154924, 0.9545237692271652, 0.9545726273007569, 0.9545768499772138, 0.9546499062021864, 0.9547047492714796, 0.9547054599084017, 0.9547075954925491, 0.9547667152309897, 0.9548192151064794, 0.9548736129235511, 0.9549178721414764, 0.9549273346213715, 0.9549417532820786, 0.9549488237788428, 0.954959860170406, 0.9549647327027935, 0.9549882409162224, 0.955039424022249, 0.9550404789307967, 0.9550466581167497, 0.9550509823642236, 0.9550633185297199, 0.9550868730556947, 0.9551566657870356, 0.955185584488972, 0.9551958747641898, 0.9552229778898566, 0.9554141758958443, 0.9554468929634288, 0.9554956296163634, 0.955565548536467, 0.9556136709949057, 0.9556224206305276, 0.9556562366124972, 0.9557298428489538, 0.9557837453871013, 0.9557912981060258, 0.9558421795571077, 0.9559114200181733, 0.9559265707704722, 0.9559401870832438, 0.9559607068159359, 0.9559794037958241, 0.9559836612891706, 0.9560880920864238, 0.9561046382096331, 0.9561262604162368, 0.9561299894548254, 0.9561664580804479, 0.9561860621741566, 0.9562052310808573, 0.9562145742094027, 0.95623206348325, 0.9563620328609144, 0.9564010228873336, 0.9564177228948365, 0.9564199867078936, 0.9564751816967365, 0.9564780526252272, 0.9564964488953362, 0.9565775005161673, 0.9567149842099031, 0.9567167186950778, 0.9567219705389551, 0.9567408558903155, 0.956764206058176, 0.9567670878838456, 0.9568417124363224, 0.9568517907040051, 0.9569033877693103, 0.9569585499766756, 0.9570083461315448, 0.9571536295843888, 0.9571765680707046, 0.9572479527393208, 0.9572629253773459, 0.9572633392571974, 0.9572682322791735, 0.9573031496928149, 0.9573152632332643, 0.9573269789872167, 0.95734064081897, 0.9573467636137742, 0.9574024744228927, 0.9574295242518376, 0.9574600920152155, 0.95748732866377, 0.9574949422925185, 0.9575805874596257, 0.9577652881414134, 0.9578172527863761, 0.9578512917593294, 0.9578986778070804, 0.9579176968540566, 0.9579450444644199, 0.9579602118866873, 0.957991533401883, 0.9580130380475272, 0.9580399264669086, 0.9580564159052016, 0.9580673858758801, 0.9581439427049492, 0.958151830667976, 0.9582692094384242, 0.9583765360304668, 0.9583846822870379, 0.9583984452414265, 0.9584087392606846, 0.9584216444809327, 0.9584575130727535, 0.9584652567573736, 0.9584714857657337, 0.9584790273731697, 0.9585008204464133, 0.9585435745596156, 0.9585600078687534, 0.9585621686683145, 0.9585808520772905, 0.9585861830558811, 0.9586089883268065, 0.9587812987156662, 0.9588083613395192, 0.9588379231959689, 0.9588391958730359, 0.9588811155141747, 0.9589463626779391, 0.9589557172284908, 0.9589834621395138, 0.9590542529412913, 0.9590622060894768, 0.95908127932716, 0.9590857466948045, 0.9590870741738591, 0.9591639631751402, 0.9592082395196802, 0.9592383920485265, 0.9592681019062693, 0.9592907283759935, 0.9593035300570335, 0.9593316919518138, 0.9593460215618606, 0.9594228363526706, 0.9594570080372767, 0.9594877246610252, 0.9595082654195043, 0.959514050949722, 0.9595300251157038, 0.9595453453919769, 0.9597118647153354, 0.9597304971228489, 0.9597554895116966, 0.9599412438738847, 0.9600079055593272, 0.9600504244254926, 0.9600846904471971, 0.9600848766243093, 0.9601614754641983, 0.9601744767195046, 0.9601927560970508, 0.9602348629252206, 0.9602678945687361, 0.9602927420147498, 0.9603526268655797, 0.9603554429205924, 0.9604832875932825, 0.9605329650268994, 0.9605710997216337, 0.9606382651143898, 0.960639507878, 0.9606604661526019, 0.9606648619894624, 0.9606749330462976, 0.9606784917124125, 0.9607106241686588, 0.9607875007619932, 0.960807500766107, 0.9608426770583807, 0.9608622284503432, 0.9609183337033349, 0.9609476449852934, 0.960956435501028, 0.960963656383047, 0.9610325962011194, 0.9610380195379952, 0.9610621582809117, 0.9610714990526296, 0.9610857810647067, 0.9611104625885342, 0.9611384631161721, 0.9611461789562444, 0.9611565046469427, 0.9611582116719167, 0.9611652143274124, 0.9611900490886216, 0.9611911761342409, 0.9612127541588575, 0.9612596670868374, 0.9612596843325175, 0.9612903885199183, 0.9614158886920006, 0.961496459753897, 0.9615094805768719, 0.9615154367887285, 0.96153280658585, 0.9615419227003039, 0.9615514602886381, 0.961569335697087, 0.9616192705877101, 0.9616409677516536, 0.9616477296079979, 0.9617065793723595, 0.9617345180348702, 0.9617461538238118, 0.9617942412755938, 0.9618263856156634, 0.961837006578856, 0.9618675705939135, 0.9619124450617588, 0.9619725903863103, 0.9620282613267749, 0.9620544250811476, 0.9621198427121304, 0.9621322748421224, 0.9621340306864852, 0.9621580539129811, 0.962216495062933, 0.9622290854085496, 0.9622442540359214, 0.9623211967659144, 0.9623606038343099, 0.9623693169483201, 0.9623758807582389, 0.9624093366438232, 0.9624239677213933, 0.9624627744412398, 0.9624975338371744, 0.9625103942319803, 0.9625159414647332, 0.9625222442238933, 0.9625406300298265, 0.9625838686838715, 0.9626134037105847, 0.9626814308532522, 0.9627986738962518, 0.9628151451010075, 0.9628310095949099, 0.9628784308985117, 0.9628807884908605, 0.9629925751395695, 0.9630338870769971, 0.9630398132927511, 0.9630472493396448, 0.9630932603878867, 0.9631361474816036, 0.9631485450762024, 0.963170050910186, 0.9631930952316892, 0.9631968522204082, 0.9632076228832109, 0.9633109820914384, 0.9633115898056211, 0.9633145845968831, 0.9633264307749656, 0.9633377731302188, 0.9633447546956418, 0.9634413260871387, 0.9634836838454253, 0.9634871892682287, 0.9634940584842533, 0.963501950170206, 0.963610271757765, 0.9637100370131355, 0.9637250015760674, 0.9637489302104048, 0.9637546571924407, 0.963781307725343, 0.9637822658624974, 0.9637999677008836, 0.9638348364009954, 0.9638987423937181, 0.9639237544583039, 0.9639806325886671, 0.963993590372036, 0.9639962759679537, 0.9640027142890156, 0.9640202700979532, 0.9640254458447816, 0.9640874472336054, 0.9641591097942024, 0.9641702886803519, 0.9641892380344176, 0.9642319959249779, 0.9642428695275892, 0.9642707243061864, 0.9643845166042307, 0.9644771611875823, 0.9645111317537689, 0.9645181831214651, 0.9645336805325502, 0.9645922249752547, 0.9645961313538838, 0.96459875638788, 0.9646138074842461, 0.964620191298275, 0.9646875725901515, 0.9647178321533517, 0.9647188140659346, 0.9647289814913927, 0.9647380657268947, 0.964799116543997, 0.9647994764450923, 0.9648197319524664, 0.9648463310179751, 0.9648624437720312, 0.9648720477126753, 0.9648917174106814, 0.9652436814593222, 0.965257770518097, 0.9652798024836221, 0.9653604801008702, 0.9653751354324285, 0.9654058260228123, 0.9654155873066919, 0.9654471765807652, 0.9655129519639982, 0.9655570320661089, 0.9655612473389251, 0.9656377030620485, 0.9657646634080672, 0.9657647859833557, 0.9657818398862672, 0.965791961921597, 0.9658497329453021, 0.9659521225307841, 0.9659748501442021, 0.9660092423385788, 0.9660118395295985, 0.9660210840424022, 0.9660538985606253, 0.9660586040144049, 0.966079037614569, 0.9661207627202986, 0.9661340638608398, 0.96616335237046, 0.9661823349585431, 0.9661848333935683, 0.9661873682339991, 0.9661904614262417, 0.9661968753605484, 0.9662029879047622, 0.966239254652289, 0.9662619436200711, 0.9662665659985646, 0.9663117623450669, 0.9663636274206432, 0.9663681502857799, 0.9663795975634426, 0.9665304846161401, 0.9665541278131322, 0.9665545620309802, 0.9665576620580881, 0.9666571351691994, 0.9666586953260236, 0.9666982803381913, 0.9666997619516771, 0.9667220470638805, 0.9667318882815358, 0.9667457290937, 0.9667743353544515, 0.9668001875346088, 0.966841454541149, 0.9668555883779645, 0.966951195165278, 0.9670101033709437, 0.9670391176669507, 0.9670469032552443, 0.9670718864945087, 0.9670983482211705, 0.9671174858100399, 0.9671966354982815, 0.9672156394811141, 0.9672498161346418, 0.9672606232691592, 0.9672681686097953, 0.9672861599002733, 0.9672861774398857, 0.9672883871958945, 0.9673881222456326, 0.967399694961475, 0.9674105752259635, 0.9674294312536207, 0.9674454424663458, 0.967450025341743, 0.9674971356244887, 0.9675312964016614, 0.9675406964962757, 0.9675434477991093, 0.967562361577045, 0.9675645006764081, 0.9675650662332768, 0.9675656624151008, 0.9676186397830956, 0.9676318816741007, 0.9676439791181539, 0.9676629009397937, 0.9677265667282527, 0.9677801950804845, 0.9678004420028864, 0.9678353159026711, 0.96788045356497, 0.9678874470566958, 0.9679245242765468, 0.9679663234527314, 0.9679892472068735, 0.9679992732364133, 0.9680059495376775, 0.9680718193673373, 0.9681153053083932, 0.9682006133735741, 0.9682301256306539, 0.9683075666539561, 0.9684071162516715, 0.96841612490387, 0.968429680896739, 0.9684362472148983, 0.9684374676766673, 0.9684739462324646, 0.9684763652510697, 0.9685661613595525, 0.9685734432548221, 0.9686633006661749, 0.968768087811817, 0.9688266351126202, 0.9688423788158856, 0.9688480682577282, 0.9689001133009596, 0.9689395605494106, 0.9690634441252807, 0.9690985929357687, 0.969104472551274, 0.9691137206204722, 0.9691197405146593, 0.969119861792576, 0.9691472957395791, 0.9691503201163112, 0.9692428531307664, 0.969284719043415, 0.9692902226617464, 0.9692968173902671, 0.969355601482007, 0.969378571132666, 0.9695393566875375, 0.9695583697797204, 0.9695913643186923, 0.9696181249771979, 0.9696256963517933, 0.969683563638414, 0.9696855800177768, 0.9696917354816316, 0.9697155686367559, 0.9697235004050107, 0.9697255741906567, 0.9697441852815175, 0.9697897815855169, 0.9698216798060443, 0.9698609339013341, 0.9698886665155378, 0.9699190429349516, 0.9699327631403399, 0.969947513084616, 0.9699619022510162, 0.969963728950228, 0.9699775984139063, 0.9700513796980907, 0.970055990935275, 0.9700578841825035, 0.9700578880496585, 0.9700816259446163, 0.9701347835769367, 0.9701733403359064, 0.9702397237702757, 0.9702546018987657, 0.9702871914773925, 0.970333983290835, 0.9703556492188664, 0.9703727435653642, 0.9703900918659635, 0.9704969122401113, 0.9705299894070143, 0.9705669344652795, 0.9705768446084391, 0.9706087132866288, 0.970632770057926, 0.9706331960995127, 0.9707243950370044, 0.9707399435269022, 0.9708217851405513, 0.9708236848121677, 0.9708245843311859, 0.970844076016648, 0.9708817988883891, 0.9708886135158106, 0.9709010843696521, 0.9709934560896488, 0.9710216408799489, 0.9710567845952048, 0.9710968190219769, 0.9711207379844029, 0.9711251117973485, 0.9711660068520126, 0.971170372750885, 0.9711961246525481, 0.9712106938028315, 0.9712226289711513, 0.9712429005316152, 0.9712598526359352, 0.9712859987957888, 0.9712908972103856, 0.9713958601595529, 0.9714285444177172, 0.9714482774535254, 0.9714740381922143, 0.9714889839171161, 0.9715114440417956, 0.9715477029106205, 0.9715842822952828, 0.9715860184267412, 0.9715919698329049, 0.9716404013852316, 0.9716467390342178, 0.971647790801872, 0.9717094771513519, 0.9717241804421421, 0.9717391427456394, 0.9717425266269586, 0.9717569496491963, 0.9717612938978988, 0.971823387178847, 0.9718288457728191, 0.9718380863450115, 0.9718479138201862, 0.9719023970971606, 0.9719686589885946, 0.9719751209935061, 0.9719794473412023, 0.9719967696612216, 0.9719971579815069, 0.9720015032534517, 0.9720100077896507, 0.9720235345236002, 0.9720239848530523, 0.9720296465484025, 0.9720514415398869, 0.9720779874406954, 0.9720876556001854, 0.9721227166897974, 0.9721233731863114, 0.9721904009720872, 0.9722642675294333, 0.9722850189615696, 0.9722918826340785, 0.9723089622906868, 0.9723135240374383, 0.9723537982384178, 0.9723705016513214, 0.9723780293865035, 0.9723838104414064, 0.9724203123219383, 0.9724302386397563, 0.9724508155000453, 0.9724535816892173, 0.9724590873987822, 0.9724677665202547, 0.9724688291084953, 0.9724804279311193, 0.9725119275632391, 0.9725617284184166, 0.972571629119565, 0.9725886404592597, 0.9726559183766779, 0.9726845570844691, 0.9727596144213579, 0.9727806882984028, 0.9728288971291376, 0.9728584318540153, 0.9728617745076895, 0.9728756373435808, 0.9728864247346359, 0.9729130396355343, 0.9729381202451698, 0.9729769895190566, 0.9730773133762648, 0.9730884407490831, 0.9731036422837251, 0.9731139239915316, 0.9731363426684574, 0.9731589717235032, 0.9731642101470314, 0.9731748767826403, 0.9732083604086631, 0.9732473773752741, 0.9732636812545825, 0.9732671391134752, 0.9732789549003582, 0.9733024250636337, 0.9733067762071461, 0.9733306465788188, 0.9733430503730957, 0.9733689106302066, 0.9733718207619856, 0.9734830309220487, 0.9735420726517877, 0.9735459981743481, 0.9735827809000254, 0.9735860595574403, 0.9735889311972519, 0.9735940083011347, 0.9736140960160488, 0.9736243259021877, 0.9736255438458199, 0.973661060388444, 0.9736749941125623, 0.9737038932376211, 0.9737184539878533, 0.9737239496079565, 0.9737499195154878, 0.9737530065294292, 0.9737621416097788, 0.9737856689223382, 0.9737897903813011, 0.973806160906015, 0.9738187733132323, 0.9738262526739818, 0.9738437430615592, 0.9739511190264125, 0.9740170266934703, 0.9740564430921218, 0.974104036239052, 0.9741125042234586, 0.9741441043015272, 0.9741636540205191, 0.9742157498898387, 0.9742827907176534, 0.9743045127047172, 0.9743150684859595, 0.9743402075163297, 0.9743529793062725, 0.9743559023060698, 0.9743705285222466, 0.9743857682920467, 0.9743943883087296, 0.9743969543643546, 0.9745123400212296, 0.9745485934752847, 0.9745925961778665, 0.974599531583126, 0.9746558767733285, 0.9746693982145206, 0.9746931397503297, 0.9746991987966013, 0.9747405438543859, 0.9747500561340001, 0.9747561786806817, 0.9747856936168321, 0.9747940485727898, 0.9747951772558056, 0.9747962353116798, 0.9748021071645613, 0.9748029637586986, 0.9748085835180346, 0.9748858925180786, 0.9748920686271656, 0.9749058803299557, 0.9750092021059671, 0.9750140946008705, 0.9750184781465966, 0.9750602053813948, 0.9751952106955372, 0.9751998237457363, 0.9752155022975546, 0.9752452370853191, 0.9752709747828219, 0.9752792508481224, 0.9752806730075221, 0.9753030284367136, 0.9753169102464897, 0.9753455947433519, 0.9753891387287976, 0.975401181036272, 0.9754494117935808, 0.9754624209519124, 0.9755028203795514, 0.9755385834910398, 0.9755958362714697, 0.9755996826722396, 0.9756006596828573, 0.9756032922703286, 0.9756071538604586, 0.9756133542314844, 0.9756789953923639, 0.9757369807257877, 0.9757889420505426, 0.9758052884043069, 0.9758210818378174, 0.975838977396889, 0.9758399313338926, 0.9758467163548521, 0.9759065725334755, 0.9759139536257244, 0.9759204959375152, 0.9759387244798525, 0.9759540275595252, 0.9759868443148402, 0.9759883176764355, 0.976006601874387, 0.9760265223492391, 0.9760540793225444, 0.9760698626883133, 0.9760968614631562, 0.976144171198306, 0.9761627634017006, 0.9761770049514774, 0.9761866900065832, 0.9761908989745112, 0.9762028557507169, 0.9762063375079559, 0.9762247528825756, 0.9762286895213015, 0.9762603554457838, 0.9762625403068925, 0.9762665429242096, 0.9762910263028003, 0.9763088341346273, 0.9763435035385817, 0.9763644861263165, 0.9764015633760533, 0.9764040491962128, 0.9764391811443469, 0.9764725402795877, 0.9764735925540942, 0.976506901210899, 0.9765538134132474, 0.976554994731937, 0.9765590571014614, 0.9765799815380365, 0.9765984005446577, 0.9766129018532225, 0.9766286066094608, 0.976684488722294, 0.97675393420217, 0.9767686953727727, 0.9767699232333379, 0.9767935072016608, 0.9767962228924528, 0.9767988568356035, 0.97680070052413, 0.976822926423923, 0.9768325763541252, 0.9768439045036919, 0.9768905765007733, 0.9769340215416601, 0.9769477660555186, 0.9769678339705592, 0.97697889029983, 0.9769831913567404, 0.9770219888514792, 0.9770445528344992, 0.977063566880842, 0.9770754372005713, 0.9771674805226257, 0.9772415019580685, 0.9772975655741979, 0.9773133944655595, 0.9773207267282847, 0.9773253935413475, 0.9773320475828182, 0.9773658395224718, 0.9773762864131493, 0.977378635734069, 0.9774420721011416, 0.9774571732250242, 0.9774689059111475, 0.9775134774025728, 0.9775285150821473, 0.9775583088356687, 0.9775641137916103, 0.9776018370316665, 0.977605881732632, 0.9776499932331117, 0.9776842102269413, 0.9776892065275072, 0.9777044219924181, 0.977744354062564, 0.9777520979070817, 0.9777773981376904, 0.9777975954576357, 0.9778028247625914, 0.9779238032899041, 0.9779491692300577, 0.977954245012428, 0.9779863923447476, 0.9779916294792929, 0.9780231746900195, 0.9780505298714347, 0.9780610295957323, 0.978097919642584, 0.9781782437937129, 0.9781862560795151, 0.9781927863279197, 0.9782011605711086, 0.9782563699771539, 0.978338951331894, 0.9783506674461502, 0.9783675246079084, 0.9783683595712219, 0.9784126096688929, 0.978417247129466, 0.9784346488127444, 0.9784793702428063, 0.9785041077346257, 0.978512333220191, 0.9785357617624225, 0.9785999128718393, 0.9786221596023145, 0.9786397177249495, 0.9786720190946535, 0.9787282830105676, 0.9787595291713831, 0.9788323571358698, 0.9788633860795096, 0.9788701814397877, 0.9788786432365788, 0.9789066868199818, 0.9789087258726924, 0.9789099378726448, 0.9789252230919067, 0.9789451762093605, 0.9789503918003195, 0.9789649420480808, 0.9789710701661114, 0.9789774919085003, 0.9789929079667665, 0.9790663284196114, 0.9790724205950091, 0.9790891214337726, 0.9791135476783811, 0.9791188293165045, 0.9791390578938989, 0.9791433895096534, 0.9791498488895309, 0.9791522226553209, 0.979152484020181, 0.9791720080323969, 0.9791796557842893, 0.9792325349796214, 0.9792598573563771, 0.9792995920755018, 0.9793348832042431, 0.9793782414137274, 0.979405032163908, 0.9794365633351046, 0.9794567388582667, 0.9794712960454021, 0.9794716134716795, 0.9795207888202941, 0.9795222038094943, 0.9795306287936691, 0.9795979729463293, 0.9796153119163898, 0.9796488961347385, 0.979680600296011, 0.9796985978612945, 0.9797118068771188, 0.9797206842742551, 0.9797209443737561, 0.9797667514104333, 0.9797844326982799, 0.9798129924156982, 0.9798280916709646, 0.9798572709448626, 0.9798913937407719, 0.9798924695713119, 0.9799270313783993, 0.97992756111541, 0.9799481606189847, 0.9799584136302567, 0.9800053746871846, 0.9800065778532031, 0.9800156438784582, 0.9800221882032427, 0.9800606156331079, 0.9800767085526425, 0.9800769413620689, 0.9801026620351767, 0.9801535083561542, 0.9801812606294159, 0.9801850811626027, 0.9802406031542146, 0.9802503265176163, 0.9802605684239178, 0.9802626702884039, 0.9802652577451891, 0.9803359302728047, 0.980339491370984, 0.9803444173574458, 0.9803956624594938, 0.9803980608034755, 0.9804067808673991, 0.9804335038033135, 0.9804413368773874, 0.9804745433878224, 0.9805907546613177, 0.9806047490520743, 0.9806282624589219, 0.980692108053649, 0.9807324443542762, 0.9807390495470768, 0.9807804939589901, 0.9807922246720726, 0.9808153997439678, 0.9808169071661659, 0.9808432733445073, 0.9808545334089267, 0.9809000474783794, 0.9809370600456774, 0.980945084210022, 0.9810829855954488, 0.9812247083387554, 0.9812312682566292, 0.9812632598179658, 0.9813034685247786, 0.9813403001123867, 0.9814025611805095, 0.9814122651685557, 0.9814241597095501, 0.9814268113454722, 0.981479260276702, 0.981517358871408, 0.9815250972528505, 0.9815271697511522, 0.9815579876311517, 0.981588581797699, 0.981593220902154, 0.9815994383886498, 0.9816135376809586, 0.981619148976616, 0.9816282103642592, 0.9816284830836429, 0.9816289959732014, 0.9816652770622969, 0.9816986739289381, 0.9817045400719064, 0.9817243526974135, 0.9817475713177242, 0.9817601787903418, 0.9818742996952038, 0.9819330673573237, 0.9819380578027177, 0.981938430327989, 0.9819494665990898, 0.9819931426460706, 0.9820001214645888, 0.9820098118351783, 0.9820177184726473, 0.9821424932879811, 0.9821776916389446, 0.9821979956648562, 0.9822149482920274, 0.9822329747901891, 0.9822642689519109, 0.9822725343996093, 0.982273304035935, 0.982282168824769, 0.9822836307466064, 0.9823088015586557, 0.982312113982565, 0.9823219771231868, 0.9823939237493359, 0.982411149845164, 0.982434724854934, 0.9824619870801536, 0.9824732246690618, 0.9824779475013483, 0.9824853412733292, 0.9824950435842117, 0.982511937625769, 0.9825424279041419, 0.982550486982997, 0.9825827974012439, 0.9825848447180607, 0.9826780000198326, 0.9827014321525722, 0.9827348692977652, 0.9827628665500259, 0.9827984442496333, 0.9828302978373187, 0.982841086656131, 0.9828492157049598, 0.9828573416559225, 0.9828623784788499, 0.9828726953033984, 0.9828874637920048, 0.9829008315691324, 0.9829448515465518, 0.9829461008098083, 0.9829496286972986, 0.9829522661405868, 0.9829761808694222, 0.9829985807505058, 0.9830014630164045, 0.9830345265114782, 0.9830433711325565, 0.9830468501373908, 0.9830516693819873, 0.9830668985395038, 0.9830935292508312, 0.9830966700403726, 0.9831224572212671, 0.9831320504668202, 0.9831477684693019, 0.9831537371915801, 0.9831729646673257, 0.9832027369948558, 0.9832264459633677, 0.98323267213812, 0.9832762160186704, 0.9833054517865101, 0.9833406137923854, 0.9833884908388293, 0.9833926635142083, 0.9834074146055561, 0.9834089983263248, 0.9834132631925991, 0.9834785354092023, 0.9834974782623556, 0.983555537318161, 0.9835667153317249, 0.9835690170333314, 0.9836381933058165, 0.9837104339364766, 0.9837630565577639, 0.9837721635954858, 0.9837863510504724, 0.9837994950781116, 0.9838242645980414, 0.9838286936672889, 0.983844749000283, 0.983846093502757, 0.9838578226868426, 0.9838638393720684, 0.9838722854745553, 0.98389631608153, 0.9839334773348328, 0.9840012242942551, 0.9840141629259156, 0.9840442505195991, 0.9840454417595985, 0.9840530547555744, 0.9840567521093083, 0.9840631201214795, 0.9840646214609327, 0.9840737708921989, 0.9840783672019935, 0.9841427788197599, 0.9841512792326688, 0.9841516063434432, 0.9841580271583067, 0.9841769387958358, 0.9842216283063945, 0.9842253081247607, 0.9842410802165296, 0.9842487082387744, 0.9842635154027195, 0.9842712669696657, 0.9842770650299779, 0.9843022800403001, 0.9843040051680406, 0.9843336439897893, 0.9843341021237372, 0.9843430989853774, 0.9843891984155207, 0.9843923624085732, 0.9844090797545664, 0.9844225137271823, 0.9844377499614123, 0.9844560681579264, 0.9844652340273703, 0.984505792707506, 0.9845279730256091, 0.9845352706967373, 0.9845596086313013, 0.9845786214952921, 0.9845791741420967, 0.9845894715301545, 0.9846004295106303, 0.9846047831911172, 0.984617547549848, 0.9846212712124437, 0.9846486246072823, 0.9846593188591296, 0.9847808284768521, 0.9848508016213758, 0.9849163223746077, 0.9849484287875083, 0.9849522794115082, 0.9849629068719479, 0.9849680373600286, 0.9850376687631008, 0.9850554384893816, 0.9850989946767493, 0.9851272294327588, 0.9851382983493299, 0.985167431990445, 0.9851738895559972, 0.9851750964514049, 0.9851800171172058, 0.985238696544115, 0.9852936714687051, 0.9853057802176384, 0.9853258152117512, 0.9853340646689577, 0.9853887743592465, 0.9853996325829315, 0.9854008915221982, 0.9854205427928188, 0.985426742320809, 0.9854527973430129, 0.9854781866240226, 0.9854821465181206, 0.9854873493967129, 0.985523499889785, 0.9855255201795319, 0.98554527275639, 0.9855555815063219, 0.9855795745781118, 0.9855887220218197, 0.9855892781119233, 0.9855997315975827, 0.9856482684679682, 0.9856564166607548, 0.985684585478109, 0.9857613263935109, 0.9857707781753271, 0.9857915804293305, 0.9858176556864523, 0.9858291114873176, 0.985843433756157, 0.9858445620848335, 0.9858930167312857, 0.9859303734240367, 0.9859350240489283, 0.9859448862195358, 0.985964395416929, 0.9860398815533151, 0.986068736269349, 0.9860775471921069, 0.9860842461933148, 0.9861016707912423, 0.9861033270102143, 0.9861173192352967, 0.9861282611570321, 0.9861590531278466, 0.9861674840261615, 0.9861945870013813, 0.9862011953574108, 0.9862183562098027, 0.9862275882895871, 0.9862346046478631, 0.9862502926132284, 0.9862846212556204, 0.986302268165397, 0.9863213066242059, 0.9863290919748545, 0.9863341368679305, 0.9863443418634501, 0.9863476747932898, 0.986354912403375, 0.9863655334153025, 0.9863825387357231, 0.9863847501040421, 0.98639372356464, 0.9864044835644127, 0.9864170953894539, 0.9864295761104239, 0.9864330479978899, 0.9864382533816092, 0.9864422185627176, 0.9864618803613832, 0.986495527549455, 0.9865024889192693, 0.9865073786737276, 0.9865381014823454, 0.9865381127629536, 0.986615562070086, 0.9866358272549346, 0.9867129423316473, 0.9867224153230866, 0.9867313328234274, 0.9867381238090357, 0.9867547791709108, 0.9867755012799563, 0.9867817898301502, 0.9868287547373163, 0.9868290277724737, 0.9868557096073481, 0.9868732909211464, 0.9868737371657662, 0.9868825136634891, 0.9868921095834104, 0.9869354652461705, 0.9869475116175975, 0.9869496399741806, 0.9869634688873384, 0.9869649012467709, 0.9869833076081723, 0.9869846989591643, 0.9869942065007387, 0.9870065571970913, 0.9870090484702714, 0.9870599169589861, 0.9870618811013341, 0.9870746383658507, 0.9871160900794839, 0.9871165362469043, 0.9871526128601775, 0.9871754010594657, 0.9871861632949033, 0.9872445142421068, 0.9872717624724969, 0.9873183327144244, 0.987357589639222, 0.9873621700766458, 0.9873852141273759, 0.9874105189378191, 0.987422209436565, 0.9874480705048834, 0.9874554949863162, 0.9874781110427988, 0.9875335239071866, 0.9875467832695376, 0.9875591263159397, 0.9875687913447445, 0.987603770022113, 0.9876325476507228, 0.98763298235865, 0.9876797918222925, 0.987701358069917, 0.9877263894517888, 0.9877572779809, 0.9877818702383336, 0.9878325233795306, 0.9878384849073918, 0.9878463503929305, 0.9878613715112382, 0.9878841353361875, 0.9878957286730571, 0.9879380435117046, 0.9879484143956369, 0.9879501647480213, 0.9879560734176615, 0.9879575204925337, 0.9879669077914559, 0.9879780822556204, 0.9880121079484282, 0.9880286638167577, 0.9880658827521135, 0.9880957294919125, 0.9881372017960703, 0.9881562651409008, 0.9882159555971148, 0.9882161271869206, 0.9882492754533706, 0.9883081227461665, 0.9883118690599296, 0.9883365880658, 0.9883532357923849, 0.9883904163272238, 0.9884039875804332, 0.9884077511878507, 0.9884293776919785, 0.9884302007487957, 0.9884380232670134, 0.9884593634634409, 0.9884627916149238, 0.9884977407824306, 0.9885151834887624, 0.9885477573861452, 0.988550570707904, 0.988586947272066, 0.9885887223014728, 0.9885989736298335, 0.9886055739840836, 0.9886318158655283, 0.9886411611314413, 0.9886466842265653, 0.988681547286741, 0.988749652796652, 0.9887573341496361, 0.9887723772238762, 0.988810357491302, 0.9888539682691928, 0.9888553405198117, 0.9888755620286497, 0.9889025310851198, 0.988943221195535, 0.9889443985040768, 0.9889552101860873, 0.9889761322669454, 0.98897681669605, 0.9889831609747334, 0.9889852653172831, 0.9890206534638756, 0.9890212775041222, 0.9890414814511167, 0.9890423468126973, 0.9890529762908732, 0.9890763134006918, 0.9890953792402546, 0.9890993292143292, 0.9891888620685519, 0.989205407689689, 0.9892216605621281, 0.9892659207583173, 0.9892941361136969, 0.9893231188257156, 0.989354134051891, 0.9893850304620792, 0.9894065192046918, 0.9894278269721256, 0.9894288309687166, 0.9894393720667216, 0.9894580738982879, 0.9894672086766995, 0.9895212835436519, 0.9895478906138485, 0.9895516776193137, 0.9895666235444442, 0.9895670902633974, 0.9895755078218375, 0.9895833971921422, 0.989595086477035, 0.9895968212791827, 0.9895983689743937, 0.9896093500112303, 0.9896169853235564, 0.9896261029049761, 0.9896396634568763, 0.9897003177208282, 0.9897219069320704, 0.9897223484959589, 0.9897381083710555, 0.989741654562408, 0.9897485765289287, 0.989777127877514, 0.989777310227015, 0.9898255377315216, 0.9898296940878749, 0.9898442970121168, 0.9898576335091114, 0.9898596764078261, 0.9898771943356514, 0.9898883968141258, 0.9899055897559973, 0.9899078700393024, 0.9899263501584793, 0.9899501980347486, 0.989972355344451, 0.990003459461411, 0.9900243876422173, 0.9900376767758706, 0.990064988408636, 0.9900686426617219, 0.9900719896969149, 0.9900864774231635, 0.9901441317566206, 0.9901482124354823, 0.990150271490113, 0.9901573402637156, 0.9901934900758771, 0.9901958483316925, 0.990196776432308, 0.9902088668556416, 0.9902374019184061, 0.9902598542551476, 0.9902796521751817, 0.9902868517901511, 0.9903257005114903, 0.9903976881496013, 0.99043440726153, 0.9904638416850221, 0.9905069642632398, 0.9905145456164582, 0.9905495674528199, 0.990583720838008, 0.99058516961224, 0.9905861275563129, 0.9905862804392924, 0.9905891502096679, 0.990605371541552, 0.9906296366857382, 0.9906298062195219, 0.9906350716095947, 0.9906389314468392, 0.9906685306531987, 0.9907308053999127, 0.99074071724771, 0.9907475663076603, 0.9907626801399934, 0.9907764064655902, 0.9907850745396609, 0.9907916155886167, 0.9908513300134959, 0.9908695508048644, 0.9908709481252491, 0.9908757904710183, 0.9908920904478837, 0.9909020493726317, 0.9909057322692777, 0.9909116002579147, 0.9909244813889161, 0.9910172340252729, 0.9910182489959037, 0.99105245152219, 0.9910779178655263, 0.9910794560498011, 0.9910905788899488, 0.991100086128003, 0.9911143415929917, 0.9911481836285649, 0.9911605405643688, 0.9912038714744053, 0.9912040380540665, 0.9912088334131032, 0.9912136692156472, 0.9912331844748722, 0.9912476815182214, 0.9912481058668879, 0.9912829850792403, 0.9912835949201974, 0.9913265170833931, 0.991335041344943, 0.9913362403705922, 0.9913743139876192, 0.9913915285731812, 0.9914358200636254, 0.9914451525089693, 0.991495511037795, 0.9915419935107952, 0.9915425648203714, 0.9915614534281568, 0.9915813788394081, 0.9916030294459852, 0.991610225541073, 0.9916261939069311, 0.9916472753955392, 0.9916520169858953, 0.9917127216998185, 0.9917215814063794, 0.991740409332209, 0.991744292139269, 0.9917544471178139, 0.9917919094098665, 0.9917986782778103, 0.9917998633113585, 0.9918009706813143, 0.9918426107402463, 0.9918578252095881, 0.9918730861445965, 0.9919047120347876, 0.9919411134805343, 0.9919500158312102, 0.9919533830916832, 0.9919792974858711, 0.9919845988233034, 0.9919969368581529, 0.992007897583788, 0.9920079827994549, 0.9920190467230833, 0.9920533911262663, 0.9921151392053774, 0.9921255153714374, 0.9921283352937434, 0.9921338853759087, 0.9921395328376663, 0.9921442095265716, 0.9921472893820512, 0.9921550068854007, 0.9921860942070475, 0.9921911064188761, 0.9922004890253799, 0.9922227137380856, 0.9922558829660485, 0.9922764776536789, 0.9922877830850693, 0.9922931861850414, 0.9923048307253078, 0.9923363895396593, 0.9923452630529577, 0.9923636397647396, 0.9924127467121983, 0.9924169704639278, 0.9924208215837242, 0.9924296654910257, 0.992470506490747, 0.9925142808251155, 0.9925241546318636, 0.9925253677176745, 0.992569529173595, 0.9926531080794653, 0.99267334766138, 0.9926734825749873, 0.9927207738313589, 0.9927239818111252, 0.992736980586828, 0.9927484361467191, 0.9927912772124883, 0.992801294935346, 0.9928158847001358, 0.992826193710918, 0.9928289522538516, 0.9928372485347441, 0.9928513187548957, 0.9928921739341594, 0.9929221714210554, 0.9929452736244007, 0.992987655116033, 0.9929880944649928, 0.9929931949329779, 0.9930092189924261, 0.9930123202287164, 0.9930276350374211, 0.9930314902749442, 0.9930417595689996, 0.9930651741579776, 0.9931619607500415, 0.993163361541648, 0.9931676338407759, 0.9931788892961718, 0.9931899689290259, 0.993206992871684, 0.9932077203568285, 0.9932437167539094, 0.9932692563951294, 0.9932997368588293, 0.993300428950903, 0.9933067041512603, 0.9933100939066681, 0.9933102384960533, 0.993314680333113, 0.9933236323520405, 0.9933317712012761, 0.9933569602569122, 0.9933610676571036, 0.9934025121876127, 0.9934036715422266, 0.9934043194059176, 0.9934269929133979, 0.9934503276188815, 0.9934516842250881, 0.9934616638530384, 0.9934777893973654, 0.9935236107977394, 0.9935325446432524, 0.9935345756351328, 0.9935421776840672, 0.9935621390561887, 0.9935634379176284, 0.9935858001230233, 0.9936184549971943, 0.993630343390039, 0.9936434534410984, 0.993644413334592, 0.9936670299715387, 0.9936807300740135, 0.9937041036893212, 0.9937132603369635, 0.9937461909175019, 0.9937521817489451, 0.9937644294986809, 0.9937668248241321, 0.9937869443157894, 0.993823267378999, 0.9938248173005482, 0.9938797603966981, 0.9938937525185902, 0.9939291643166261, 0.993936427922094, 0.9940159394737256, 0.9940190770131903, 0.994019212984673, 0.994035432268001, 0.9940439722943234, 0.9940531719969457, 0.9940568767805426, 0.9940645670401232, 0.9940648119424601, 0.9940650805965074, 0.9940757028240749, 0.9940834894766876, 0.9940971468109064, 0.9941069452089172, 0.99410770010465, 0.9941275454542412, 0.9941341985023825, 0.9941374954350708, 0.9941878312153072, 0.994204148573009, 0.9942244761975039, 0.9942458966222093, 0.994271769806207, 0.9942826607629761, 0.9942901690811283, 0.9942924256387119, 0.994306144105794, 0.994311730332458, 0.9943196111767051, 0.9943217799257353, 0.9943343635035118, 0.9943361412984275, 0.994361114075414, 0.9943729334601499, 0.9943950555412752, 0.9944022309669074, 0.994441262840953, 0.9944811419674119, 0.9944853413165329, 0.9945055341635024, 0.9945464210228275, 0.994550174885437, 0.9945538325822139, 0.9945856176859396, 0.9945952700218633, 0.9945953801327109, 0.9946153114657795, 0.9946153380006271, 0.9947094047517029, 0.9947194350000519, 0.99473210892934, 0.9947428547341606, 0.9947536450800756, 0.9947622387022454, 0.9947836693888198, 0.9947944014809244, 0.994802891564267, 0.9948541572248799, 0.9948651386527775, 0.9949110010314348, 0.9949155253176754, 0.9949171533820622, 0.9949400699715789, 0.9949446496742762, 0.9949539095273311, 0.9949690838490998, 0.9949828704152296, 0.9949931146208395, 0.9949948070550075, 0.9949978457278358, 0.9950038357329944, 0.9950588019074944, 0.995093633033207, 0.99510748507832, 0.9951110078429095, 0.995112733829427, 0.9951157044971473, 0.9951166876636696, 0.9951385624166085, 0.9951492041106699, 0.995160274515675, 0.9951814148554605, 0.9951931871845918, 0.9952007845901801, 0.995236409672196, 0.9952385031520968, 0.9952387033568683, 0.995246815532462, 0.9952536666476337, 0.9952560067650493, 0.9952575162679319, 0.9952757076520341, 0.9952825920851083, 0.9953109014162957, 0.9953210764264591, 0.995328033890664, 0.9953394184666986, 0.9953728310477151, 0.9953740349530839, 0.995382602678678, 0.9953832393015501, 0.9954115484931152, 0.9954174108439475, 0.9954474566307039, 0.9954676112188878, 0.9954678225239568, 0.9954878611917016, 0.9955337435406095, 0.9955616475853163, 0.995611460590303, 0.995618727983464, 0.9956344070884481, 0.9956399481656425, 0.9956487907318635, 0.9956495884663665, 0.9956521497217489, 0.9956591973548394, 0.9956657819143757, 0.995670018440439, 0.9956740798620887, 0.9957243782744012, 0.9957461379300427, 0.9957539606861256, 0.9957578215685751, 0.9957687195389364, 0.9957780074944644, 0.9957868108935221, 0.9957882547138944, 0.9957948506973198, 0.99581023489683, 0.9958111610036621, 0.9958421139210414, 0.995878192423295, 0.995881250809641, 0.9958855608935744, 0.9959468669847165, 0.9959938522961862, 0.9959953895970294, 0.9960081802235743, 0.9960311519746324, 0.9960373439866527, 0.9960375951293051, 0.9960942126245268, 0.9960989551300371, 0.9961135272104013, 0.9961166297050372, 0.9961215743124966, 0.9961235825580869, 0.9961366470370148, 0.9961500334790427, 0.9961558667703267, 0.9961718050630141, 0.9962012634146794, 0.9962224688089104, 0.9962595912434073, 0.9963355563619954, 0.9963359575480986, 0.9963362959676091, 0.9963392789922908, 0.9963508948319125, 0.9963968350250336, 0.9964027223354571, 0.9964456415943853, 0.9964675317673778, 0.996489159319584, 0.9964928227715882, 0.9964999378924053, 0.9965679916527421, 0.996568865289134, 0.9965834187028396, 0.9966100337985909, 0.9966292206191285, 0.9966400329438477, 0.9966477219012135, 0.9966801916501109, 0.9966890285519248, 0.9967081445495432, 0.9967118232427625, 0.9967160815825266, 0.9967210928366158, 0.9967378364011603, 0.9967435795511018, 0.9967647061494711, 0.9967727728056732, 0.996782060422222, 0.9967841940231592, 0.9968181242947156, 0.9968220058696953, 0.9968672320463016, 0.996915888373137, 0.9969163242016619, 0.9969350466361836, 0.9970132323996294, 0.9970207431309482, 0.9970209586691996, 0.9970307825987519, 0.9970620180101334, 0.9970651546935749, 0.9970776001678237, 0.9970843760889104, 0.9970859807428395, 0.9971017052267576, 0.9971129886315305, 0.997175741923337, 0.997176031957525, 0.9971967956415542, 0.9972154058713543, 0.9972201813930387, 0.9972244369724796, 0.9972507069371241, 0.997257259878623, 0.9972637389088563, 0.9972722264919419, 0.9972779650832487, 0.9972825129794756, 0.9972982089394336, 0.9972998785086701, 0.9973203847276692, 0.9973238967040756, 0.9973692469213212, 0.9973694336994601, 0.9973754177503974, 0.9973813026655509, 0.9973848514606637, 0.9973918219780645, 0.9973940110097764, 0.9974258892197454, 0.9974347463921478, 0.9974374535730965, 0.9974416573592381, 0.9974921534071005, 0.9975252586398061, 0.9975729019347664, 0.997572933736076, 0.9975777245624554, 0.9975904092854954, 0.9976390903376218, 0.9976514818470122, 0.9976536396108957, 0.9976564095386771, 0.9976568246970428, 0.997664429343239, 0.9976696850219572, 0.997676772442949, 0.9976902192931975, 0.9976999492131258, 0.9977226784386088, 0.9977697549774431, 0.997799368045472, 0.9978071294777576, 0.9978205121950283, 0.9978574043585231, 0.9978614266982052, 0.9978884955576551, 0.9978993176746607, 0.997901389624498, 0.9979326170926589, 0.9979471457499806, 0.9979614403947348, 0.9980144803566364, 0.9980165227430449, 0.9980484320667407, 0.998052152422137, 0.9980573717146954, 0.9980631496133513, 0.9980797318610276, 0.9980875558311249, 0.9980898553616332, 0.9981021907002118, 0.9981257118231973, 0.9981298125539478, 0.9981386004746968, 0.9981492032855914, 0.9981605746109568, 0.9981681472668767, 0.9982301089893904, 0.9982402216875023, 0.9982428034979592, 0.9982435595006854, 0.9982469903868614, 0.9982574845189778, 0.9982645696082962, 0.9982672454351943, 0.9982683412514435, 0.9982741448963894, 0.9983010198404842, 0.9983334586093352, 0.9983386292834893, 0.9983780020991303, 0.9983784517432845, 0.9983839427446535, 0.9983992647292796, 0.9984046659984589, 0.9984088820051076, 0.9984472132073251, 0.9984626879144287, 0.9984799806883894, 0.9984958861559684, 0.998506440233432, 0.9985114396373892, 0.9985223471420417, 0.9985560617388031, 0.9985610804004761, 0.9985692424931564, 0.9986168705593462, 0.9986171988203916, 0.9986276844160015, 0.9986295398572715, 0.9986320774817891, 0.9986395209240037, 0.9986684601351098, 0.9986904190298073, 0.9987089184321606, 0.998763110968577, 0.998764239932817, 0.9987886357090913, 0.9988070449876448, 0.9988577741685626, 0.9988616952818618, 0.9988646758630043, 0.9988914040739408, 0.9989084067610147, 0.9989208551467612, 0.9989245255580269, 0.998942572537078, 0.9989522679540557, 0.9989581774797772, 0.9989708217970791, 0.998971191064458, 0.9989800549832349, 0.9989827149271489, 0.9989831638568877, 0.9989895815307108, 0.9989906931081749, 0.9990060673915764, 0.9990315712057458, 0.999061825097626, 0.9990819259261868, 0.999087488504983, 0.9990992320559369, 0.9991628483997863, 0.9991830258557198, 0.9991936964394147, 0.9992005289857726, 0.9992054105997598, 0.9992185331203736, 0.9992209246338364, 0.9992230833075099, 0.99923380334359, 0.9992539644021143, 0.9992658292144372, 0.9992695698021263, 0.9992919652999792, 0.9993039762435914, 0.9993088296762328, 0.9993199950272463, 0.999343328940148, 0.9993541379065509, 0.9993707590131655, 0.999408874454121, 0.999419503477395, 0.9994255856217801, 0.9994273517459368, 0.9994289209818816, 0.9994296861240966, 0.9994476187674446, 0.9994546487269155, 0.9994778433649176, 0.999495630048794, 0.9994959792414109, 0.9995033698784506, 0.9995075240658116, 0.9995218778066639, 0.999528250924943, 0.9995503246951432, 0.9995788125620011, 0.9995824307445936, 0.9995967268645278, 0.9996074296106077, 0.9996082672893828, 0.9996142434136415, 0.9996160589989077, 0.9996184664743765, 0.9996205151760955, 0.9996237122464391, 0.9996345928079494, 0.999639351764008, 0.99965394466798, 0.9996725708102738, 0.9996757932960264, 0.999678999047043, 0.9996880564950656, 0.9996932382785454, 0.9997168663840683, 0.9997344683431565, 0.9997369058859139, 0.9997415924404776, 0.9997469355429057, 0.999768109819428, 0.999772598422505, 0.9997876331089625, 0.999797256361496, 0.9997990913703342, 0.9998005276116827, 0.9998052223420377, 0.9998052613148759, 0.9998123426249832, 0.9998369910585354, 0.99984514974885, 0.9998473060487628, 0.9998522156449182, 0.9998546384571526, 0.9998712760087048, 0.9998718666157408, 0.9998802301957754, 0.9998835966055248, 0.9998868690992941, 0.9998965252860067, 0.999913668435777, 0.9999271250990928, 0.9999348384526503, 0.9999478851816374, 0.9999490928358528, 0.9999499122181468, 0.9999548578998524, 0.9999562005779112, 0.999957026687172, 0.9999588735929167, 0.9999674073481407, 0.9999675118900714, 0.9999737311218474, 0.9999740384110919, 0.9999759515877571, 0.9999762220738502, 0.9999906813277756, 0.9999921707465845, 0.9999948500978305, 0.9999998195483736, 0.9999998515579356, 0.9999999999962059, 0.9999999999999788, 0.9999999999999853, 0.999999999999989, 0.9999999999999933, 0.9999999999999938, 0.9999999999999948, 0.9999999999999948, 0.9999999999999952, 0.9999999999999954, 0.9999999999999956, 0.9999999999999956, 0.9999999999999957, 0.9999999999999961, 0.9999999999999961, 0.9999999999999962, 0.9999999999999962, 0.9999999999999963, 0.9999999999999963, 0.9999999999999966, 0.9999999999999967, 0.9999999999999969, 0.999999999999997, 0.999999999999997, 0.9999999999999971, 0.9999999999999973, 0.9999999999999974, 0.9999999999999974, 0.9999999999999976, 0.9999999999999977, 0.9999999999999977, 0.9999999999999977, 0.9999999999999978, 0.9999999999999978, 0.9999999999999979, 0.999999999999998, 0.999999999999998, 0.999999999999998, 0.999999999999998, 0.999999999999998, 0.999999999999998, 0.999999999999998, 0.9999999999999981, 0.9999999999999981, 0.9999999999999981, 0.9999999999999981, 0.9999999999999981, 0.9999999999999981, 0.9999999999999982, 0.9999999999999982, 0.9999999999999983, 0.9999999999999983, 0.9999999999999983, 0.9999999999999983, 0.9999999999999984, 0.9999999999999984, 0.9999999999999984, 0.9999999999999984, 0.9999999999999984, 0.9999999999999984, 0.9999999999999984, 0.9999999999999986, 0.9999999999999986, 0.9999999999999986, 0.9999999999999986, 0.9999999999999986, 0.9999999999999986, 0.9999999999999986, 0.9999999999999987, 0.9999999999999987, 0.9999999999999987, 0.9999999999999987, 0.9999999999999987, 0.9999999999999987, 0.9999999999999987, 0.9999999999999987, 0.9999999999999987, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999988, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.9999999999999989, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.999999999999999, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999991, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999992, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999993, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999994, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999996, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999997, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999998, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999, 0.9999999999999999}

func main() {
    _, err := plotter.NewHist(histPts, 100)
    if err != nil {
        log.Fatal(err)
    }
}
```
```text
 I assume it's the panic on line 182. Do you have a stack trace?
```
```text
 Yes, it's that panic.
panic: 0.9999999999999998, xmin=0.37009128387828555, xmax=0.9999999999999999, w=0.006299087161217143, bin=100, n=100


goroutine 1 [running]:
panic(0x20ebc0, 0xc4200728f0)
    /Users/brendan/gover/go/src/runtime/panic.go:500 +0x1a1
github.com/gonum/plot/plotter.binPoints(0x364d00, 0xc4200728a0, 0x64, 0xc4200728a0, 0x10, 0x10, 0x626401)
    /Users/brendan/Documents/mygo/src/github.com/gonum/plot/plotter/histogram.go:184 +0x5cd
github.com/gonum/plot/plotter.NewHistogram(0x364d00, 0xc4200728a0, 0x64, 0x364d00, 0xc4200728a0, 0xc42006c480)
    /Users/brendan/Documents/mygo/src/github.com/gonum/plot/plotter/histogram.go:50 +0xd0
github.com/gonum/plot/plotter.NewHist(0x364c80, 0xc42006c480, 0x64, 0x364c80, 0xc42006c480, 0xc420072868)
    /Users/brendan/Documents/mygo/src/github.com/gonum/plot/plotter/histogram.go:63 +0x8c
main.main()
    /Users/brendan/Documents/mygo/src/Testing/histpanic/histpanic.go:12 +0x67

My temporary fix is to add
if bin == n {
    bin = n - 1
}

because the code around it means all of the points around it have to be in-bounds anyway, but there's probably a better long-term fix.
```
## 37.internal/cmpimg: export as gonum.org/v1/plot/cmpimg
Aug 2, 2018
```text
 Comparing plot outputs (PNG, PDF, JPG, ...) for tests is useful not only in gonum/plot but also for clients using gonum/plot, outside of this repository.
we should consider exporting gonum.org/v1/plot/internal/cmpimg.
perhaps simply as gonum.org/v1/plot/cmpimg ?
```
```text
 I would be in favor of something like that.
```
## 38.plot: make GlyphBox support plot area-dependent sized glyphs
Feb 12, 2015
```text
 At the moment, a plot.GlyphBox represents a zero-sized point in the data space surrounded by an pair of offsets in the plot space. There is no way to represent the extent of glyphs that have a size dependent on the size of the plot.
This is due to the fact that the GlyphBoxes method is not passed any information regarding the plot area.
This can be fixed by adding a second pair of normalised x,y coordinates.
type GlyphBox struct {
    // The glyph location in normalized coordinates.
    MinX, MinY float64
    MaxX, MaxY float64

    // Rect is the offset of the glyph's minimum drawing
    // point relative to the glyph location and its size.
    Rect
}

The details of types etc need to be fleshed out.
```
## 39.Only error about missing fonts if text is actually drawn
Mar 17, 2015
```text
 Original issue 132 created by eaburns on 2013-07-08T11:26:19.000Z:
Currently, plot.New returns an error if it cannot find the default font.  There are two concerns here: 1) if the default font will remain unexported then this might as well be a panic, and 2) we should support drawing plots that do not contain any text if the fonts aren't actually installed.
My plan is: export the defaultFont variable, leave the error return on plot.New, and change plot and vg so that, if the fonts aren't installed, they only error if/when text is actually drawn.
```
```text
 Comment #1 originally posted by eaburns on 2014-12-08T11:38:06.000Z:
Yes, please export the defaultFont variable. That would make customizing text style much easier.
```
```text
 Comment #2 originally posted by eaburns on 2014-12-20T14:43:04.000Z:
See issue 147 for exporting defaultFont.
```
## 40.imports path
Oct 31, 2018
```text
 Please, change imports correct paths to github.com/gonum/plot' from gonum.org/v1/plot
gonum.org/v1/plot has 404
and problems:
dep init
```
```text
 @btracey Can you look into this?
$ go get -u -v gonum.org/v1/plot
Fetching https://gonum.org/v1/plot?go-get=1
Parsing meta tags from https://gonum.org/v1/plot?go-get=1 (status code 404)
get "gonum.org/v1/plot": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot?go-get=1
gonum.org/v1/plot (download)
Fetching https://gonum.org/v1/plot/vg?go-get=1
Parsing meta tags from https://gonum.org/v1/plot/vg?go-get=1 (status code 404)
get "gonum.org/v1/plot/vg": found meta tag get.metaImport{Prefix:"gonum.org/v1/plot", VCS:"git", RepoRoot:"https://github.com/gonum/plot"} at https://gonum.org/v1/plot/vg?go-get=1
get "gonum.org/v1/plot/vg": verifying non-authoritative meta tag
Fetching https://gonum.org/v1/plot?go-get=1
Parsing meta tags from https://gonum.org/v1/plot?go-get=1 (status code 404)
github.com/golang/freetype (download)
Fetching https://golang.org/x/image/math/fixed?go-get=1
Parsing meta tags from https://golang.org/x/image/math/fixed?go-get=1 (status code 200)
<snip>

It looks like gonum.org is behaving oddly. Maybe this is related to the cert problems we are having with tests?
```
```text
 @jedi108 Changing the import paths is not the correct fix here.
```
```text
 Here is a working website
Last login: Wed Oct 31 21:58:53 on ttys004
brendan:~$ curl http://golang.org/x/image/tiff/lzw
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="golang.org/x/image git https://go.googlesource.com/image">
<meta name="go-source" content="golang.org/x/image https://github.com/golang/image/ https://github.com/golang/image/tree/master{/dir} https://github.com/golang/image/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/golang.org/x/image/tiff/lzw">
</head>
<body>
Nothing to see here; <a href="https://godoc.org/golang.org/x/image/tiff/lzw">move along</a>.
</body>
</html>

Our website has similar tags at the top, but then all of the 404 stuff.
I wonder if Hugo changed their 404 resolution, so we need to somehow make those pages exist?
brendan:~$ curl https://www.gonum.org/v1/plot
<!DOCTYPE html>
<html lang="en-us">
<head>

  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">

  
  <meta name="go-import" content="gonum.org/v1/gonum git https://github.com/gonum/gonum">
  <meta name="go-import" content="gonum.org/v1/hdf5 git https://github.com/gonum/hdf5">
  <meta name="go-import" content="gonum.org/v1/netlib git https://github.com/gonum/netlib">
  <meta name="go-import" content="gonum.org/v1/plot git https://github.com/gonum/plot">
  <meta name="go-import" content="gonum.org/v1/tools git https://github.com/gonum/tools">

  
  <meta name="go-source" content="gonum.org/v1/gonum https://github.com/gonum/gonum/ https://github.com/gonum/gonum/tree/master{/dir} https://github.com/gonum/gonum/blob/master{/dir}/{file}#L{line}"/>
  <meta name="go-source" content="gonum.org/v1/hdf5 https://github.com/gonum/hdf5/ https://github.com/gonum/hdf5/tree/master{/dir} https://github.com/gonum/hdf5/blob/master{/dir}/{file}#L{line}">
  <meta name="go-source" content="gonum.org/v1/netlib https://github.com/gonum/netlib/ https://github.com/gonum/netlib/tree/master{/dir} https://github.com/gonum/netlib/blob/master{/dir}/{file}#L{line}"/>
  <meta name="go-source" content="gonum.org/v1/plot https://github.com/gonum/plot/ https://github.com/gonum/plot/tree/master{/dir} https://github.com/gonum/plot/blob/master{/dir}/{file}#L{line}"/>
  <meta name="go-source" content="gonum.org/v1/tools https://github.com/gonum/tools/ https://github.com/gonum/tools/tree/master{/dir} https://github.com/gonum/tools/blob/master{/dir}/{file}#L{line}">

  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="theme" content="hugo-academic">
  <meta name="generator" content="Hugo 0.27" />
  <meta name="author" content="Gonum Numerical Packages">
  <meta name="description" content="Consistent, composable, and comprehensible scientific code">

  
  
  
    
  
  
    
    
    <link rel='stylesheet' href='https://d33wubrfki0l68.cloudfront.net/css/384624f7b31abdea686b0923d4771c7cee288772/css/highlight.min.css'/>
```
```text
 Maybe to do that on-the-fly website generation we need to use headless bundles? https://gohugo.io/content-management/page-bundles/#headless-bundle
```
```text
 here is what I did for go-hep.org:
https://github.com/go-hep/go-hep.org/blob/master/main.go
I put a net/http server in front and served the blog+news alongside the go-doc/go-get stuff.
```
```text
 i try command
dep ensure
The following errors occurred while deducing packages:
  * "gonum.org/v1/plot/vg/vgimg": unable to deduce repository and source type for "gonum.org/v1/plot": unable to read metadata: unable to fetch raw metadata: failed HTTP request to URL "http://gonum.org/v1/plot?go-get=1": Get https://www.gonum.org/v1/plot?go-get=1: dial tcp 167.99.129.42:443: connect: connection refused
  * "gonum.org/v1/plot/vg/draw": unable to deduce repository and source type for "gonum.org/v1/plot": unable to read metadata: unable to fetch raw metadata: failed HTTP request to URL "http://gonum.org/v1/plot?go-get=1": Get https://www.gonum.org/v1/plot?go-get=1: dial tcp 167.99.129.42:443: connect: connection refused
  * "gonum.org/v1/plot": unable to deduce repository and source type for "gonum.org/v1/plot": unable to read metadata: unable to fetch raw metadata: failed HTTP request to URL "http://gonum.org/v1/plot?go-get=1": Get https://www.gonum.org/v1/plot?go-get=1: dial tcp 167.99.129.42:443: connect: connection refused

validateParams: could not deduce external imports' project roots
```
```text
 Hmmm... I'll try to look into this. If anyone knows how to help, any would
be appreciated.
…
On Wed, Nov 14, 2018, 9:55 AM Vadim Tsurkov ***@***.*** wrote:
 i try command

 dep ensure

 The following errors occurred while deducing packages:
   * "gonum.org/v1/plot/vg/vgimg": unable to deduce repository and source type for "gonum.org/v1/plot": unable to read metadata: unable to fetch raw metadata: failed HTTP request to URL "http://gonum.org/v1/plot?go-get=1": Get https://www.gonum.org/v1/plot?go-get=1: dial tcp 167.99.129.42:443: connect: connection refused
   * "gonum.org/v1/plot/vg/draw": unable to deduce repository and source type for "gonum.org/v1/plot": unable to read metadata: unable to fetch raw metadata: failed HTTP request to URL "http://gonum.org/v1/plot?go-get=1": Get https://www.gonum.org/v1/plot?go-get=1: dial tcp 167.99.129.42:443: connect: connection refused
   * "gonum.org/v1/plot": unable to deduce repository and source type for "gonum.org/v1/plot": unable to read metadata: unable to fetch raw metadata: failed HTTP request to URL "http://gonum.org/v1/plot?go-get=1": Get https://www.gonum.org/v1/plot?go-get=1: dial tcp 167.99.129.42:443: connect: connection refused

 validateParams: could not deduce external imports' project roots

 —
 You are receiving this because you were mentioned.
 Reply to this email directly, view it on GitHub
 <#481 (comment)>, or mute
 the thread
 <https://github.com/notifications/unsubscribe-auth/ADgqW7HLonhpr41hleAYOsKIfNBBKb_1ks5uu-ijgaJpZM4YFTNt>
 .
```
```text
 With the new module support as of 1.11.x, have you considered enabling GO111MODULE=on, from default GO111MODULE=auto, as a temporary measure? Though this might be causing all kinds of issues converging from dep to go. I followed these steps to get the module support going for one of my libraries:
Moving from DEP to VGO (go mod)

export GO111MODULE=on
go mod init
go vendor
go build
```
## 41.plot/palette: Allow Radial to set the critical index
Nov 16, 2016
```text
 It's common to want a specific diverging point in a radial map. For example, 0 may be a critical value in many applications, or 1 in aerospace applications for the Mach number. It would be nice to have a construction provided that allows that.
```
```text
 I don't think this is the responsibility of palette. Note that palette.Diverging is an interface, so the user can define any type of colour mapping symmetry. The actual mapping of data values to colours happens in the plotter that is rendering the value-colour relationship. So to a diverging palette around a defined value (0 or 1 as per the example), the Min and Max values of the plotter can be set to so that the desired origin lies half way between them. Am I missing something here? It seems to me that loading the responsibility onto the palette types excessively parameterises those types.
```
```text
 I meant the modifying the Radial function within palette, or providing a second function similar to it.
```
```text
 What would that look like and how would the information propagate to the plotter (at the moment the critical value is not used by any plotter)?
```
```text
 I see. I didn't sufficiently understand how it worked. I'll ponder.
When I did this with my own implementation in the past, I set the critical value, and inferred the min/max from the data. There was then an individual interpolation on either end.
```
```text
 A pull request I'm currently working on has this ability built in for different type of diverging color maps. See here.
```
## 42.Make it easier to create a Heat Map with a Legend
Mar 14, 2018
```text
 Is it possible to make a utility function (maybe in plotutil?) for making a simple heat map with a legend? Seems like the easiest way at the moment is to copy and paste the current code in heat_test.go, but even that seems to be difficult to get to work right.
```
```text
 Do you want to suggest API for this?
```
```text
 I think a more general-purpose solution (or perhaps a prerequisite to the proposal above) might be to add (or export) the capability to draw a Legend on an arbitrary canvas. Right now it's only allowed to draw it on the same canvas as the plot
```
```text
 It looks like it might be as simple as exporting Legend.draw and maybe makeLegend
```
```text
 My thought was maybe something with plotutil that returns the heat map and the legend, but it sounds like @ctessum has a much clearer idea.
```
```text
 So with #436, it is now possible to create a utility function because we now know how tall and wide a legend is and therefore know how much room to leave for it. I don't have a strong idea of what such a function should look like, though, so I will leave that to someone else.
```
## 43.Padding below X axis?
Jul 10, 2016
```text
 I've added custom labels to my X axis but the text is getting cut off (the bottom half of the "y" in "May" isn't visible). Is there a way to add padding to the bottom of the graph?
```
```text
 You can use draw.Crop: https://godoc.org/github.com/gonum/plot/vg/draw#Crop.
However, if this is happening it is probably a bug so if you are willing to submit an example that reproduces the problem that would be appreciated.
```
```text
 Thanks!
I will put together an example. What information would you need?
```
```text
 Attaching an example of a graph.
```
```text
 We'd need the code that generated it too.
```
## 44.Feature request: Superscript/subscript support in text strings
Oct 5, 2015
```text
 Useful for plots that include chemical names, variables, etc. (example: R2).
A possible way to implement this would be to support a subset of Markdown or HTML syntax (for subscript and superscript, and possibly but less importantly bold and italics). There already exist Go parsers for both, so the tricky part would be implementing the text sizing and offsetting in all of the vg backends.
A more powerful alternative would be to support Latex syntax, but I feel like the difficulty to usefulness ratio would be too high.
```
```text
 This will require some research. I think that it's more complex than simply drawing smaller characters that are either raised or lowered. With smallcaps, for example, fonts have entire separate glyph sets. Simply drawing smaller versions of normal capital letters is wrong and looks bad. I suspect that superscripts and subscripts are the same. If so, we should do them right.
```
```text
 @eaburns, I was thinking similar things - this is probably related to @vdobler's CSS proposal.
```
```text
 I think doing it right requires using fonts that include superscript and subscript glyphs. The unicode standard supports superscript and subscript for a small subset of characters (https://en.wikipedia.org/wiki/Unicode_subscripts_and_superscripts). These characters will already work in the vg backends except pdf (which doesn't seem to support unicode). Some fonts include more characters but I don't know how common it is.
So I think the common strategy is to use the correct glyphs if available, and fall back to an adjustment of size, alignment, and (maybe) thickness if they are not available: https://en.wikipedia.org/wiki/Subscript_and_superscript#Desktop_publishing
It might make sense to implement the fallback first, because it doesn't require knowledge of the opentype standard and will be required anyway, and then implement the "professional" solution later.
```
```text
 My guess is that, if we implement the fallback first, we'll never go back and do it the correct way. That doesn't mean that we shouldn't. However, it would be nice if this was all in its own package (using vg to draw markdown or something), where someone was dedicated to maintaining it and find it right. Also, that way, the package could be used for more than just gonum/plot.
```
```text
 Sorry, I'm on my phone: s/find it right/doing it right/
```
## 45.Default labeling for X axis in timeseries is sparse.
Jul 15, 2016
```text
 When you render a timeseries, no matter how wide you render it, there are only three labels (dates) drawn on the x axis. Probably should include more by default.  (Is there any way to change this programatically?)
```
```text
 I agree with the general issue -- the algorithm for selecting tick labels could be better. For real numbers it could do a much better job of trying to find powers of 10 and other round numbers.
```
```text
 @zaddok the distribution of ticks is not dependent on the spatial arrangement of the axis, but rather only the numerical span of the axis. You can see this here. You can see here how changing the start and end dates impact on the tick locations.
Remember that DefaultTicks is just that, a default. Maybe we should have a better heuristic with more knobs, but I think that we will probably still keep an extremely simple default as exists now.
```
```text
 Thanks! I do strongly agree with the general principle of keeping code as simple as possible.  I'll propose my reasons for the idea that, in this case, it may be warranted to improve the default behaviour:

If the default behaviour is to create a graph that no-one would ever want, perhaps the default is not adequate.
The default behaviour is possibly not user friendly for people new to the library. It would be more user friendly if the built in behaviour works out of the box. Although this could be mitigated by improving the example code for timeseries in the documentation.

Anyway, the tool is awesome either way. You guys have done some great work. This is just my feedback/reaction to the tool as a new user.
```
```text
 Yeah, you are right. The timeseries (and indeed ticks in general) have PRs for changing behaviour that should make things better.
```
```text
 Aaah great, I'll have a look at the PR's.
Im still learning the code, my initial reaction however, is that perhaps the SuggestedTicks in the default ticker:
func (DefaultTicks) Ticks(min, max float64) (ticks []Tick) {
   const SuggestedTicks = 3

could be exposed, so that it could be initialised like this:
p.X.Tick.Marker = plot.UnixTimeTicks{Format: "Jan 2", SuggestedTicks: 5}
```
```text
 Furthermore, I am wondering, Would it be worth submitting a patch for a different ticker options, i.e. MonthlyTimeTicks, or WeeklyTimeTicks.
I'm going to do it anyway, just not sure if people would be receptive to receiving it into the main code. How many people using this library, do we think would have to go ahead and roll their own MonthlyTimeTicks, or WeeklyTimeTicks variants?
```
```text
 They sound reasonable to me.
```
```text
 The issue isn't the number of labels, but label density. Having too many labels in little space makes the plot very noisy (for example look at the y-axis in the plot on the original post here). Having too few makes it hard to comprehend the scale (the x-axis of the same plot).
The underlying problem here is that the ticker code doesn't know the final size of the plot. If it did, it would select the number of ticks and the number of labels based on the density. (I think Tufty has recommendations for tick density. We should check.)
The current algorithm is based off an algorithm that I hand-tuned a long time ago when I was interested in 3in x 3in plots for two-column academic papers. I certainly wouldn't say that it is a useless default; it's just not tuned for large plots. That is why I made the tick selection code configurable. More generally, this is why I made lots of things configurable (ticks, vg backbends, plotters, etc.): it's not possible to provide defaults that everyone will like.
That said, if there is a better default, let's use it.
```
```text
 I'be created my own alternative to the built in unix time series. I'll submit it tomorrow with some sample images to see if people like it. I suspect it'll fill a use case for others as well, if useful then great, if not, no worries.
(And yes, code to measure the widths of the labels would be ideal, but That's no small amount of work for me.)
```
```text
 Hi. Any hints for this?
```
## 46.Write a wiki article on padding code
Mar 15, 2015
```text
 Original issue 73 created by eaburns on 2012-08-24T13:13:39.000Z:
A wiki article could help demystify the plot's padX and padY functions a lot.
X transform is x * (da.Max - da.Min) + da.Min:
l.X * (newMax - newMin) + newMin = da.Min - l.Min.X (we subtract l.Min.X because it's given as a negative value).  So, the X location of our left-most glyph will transform it to a location that has room for its left-most edge to lie directly on da.Min.
We also want:
r.X * (newMax - newMin) + newMin = da.Max - (r.Min.X + r.Size.X).   This is the same as above, but we are dealing with the right-most edge of the right-most glyph and we want it to end up at da.Max.
```
```text
 Comment #1 originally posted by eaburns on 2013-02-20T19:24:28.000Z:
<empty>
```
## 47.Request for utility to place error bars
Mar 17, 2015
```text
 Original issue 137 created by eaburns on 2014-01-22T19:52:20.000Z:
When you place a lot of data points with error bars you can end up with quite a bit of clutter. It would be nice to have a utility function of some sort to reduce the number of error bars placed on the plot and possibly have them staggered so that they don't overlap.
A very silly example of cluttered error bars is attached.
```
## 48.Enable Glyphs to have an outline
Jul 15, 2016
```text
 Just in case anyone is wanting to ask why we would do this, 1. other libraries do it, and 2. In my case I need to make my application create a graph that looks like the current ones (with an outline). Consider for example the javascript library, chart.js
A patch for this could be done without altering the API. Here is my current solution:
--- a/vg/draw/canvas.go
+++ b/vg/draw/canvas.go
@@ -55,6 +55,12 @@ type GlyphStyle struct {
        // Radius specifies the size of the glyph's radius.
        Radius vg.Length

+       // Width of the line around the glyph. If set, LineColour is used to outline the glyph
+       Width vg.Length
+
+       // Outline glyph with this color if Width is set.
+       LineColor color.Color
+
        // Shape draws the shape of the glyph.
        Shape GlyphDrawer
 }
@@ -106,6 +112,11 @@ func (CircleGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
        p.Arc(pt, sty.Radius, 0, 2*math.Pi)
        p.Close()
        c.Fill(p)
+
+       if sty.Width > 0 {
+               c.SetLineStyle(LineStyle{Color: sty.LineColor, Width: sty.Width})
+               c.Stroke(p)
+       }
 }

 // RingGlyph is a glyph that draws the outline of a circle.


This simply add's two fields, if you want to set an outline, you can set the new Width field, and the new LineColor field.
Patch: #298
```
```text
 This seems fine. But to be clear, if we do it, the reason isn't that other packages allow it, so we should to; the reason is because we think that it's a good thing to have. I don't want to start adding lots of features just because other plot packages have them.
```
```text
 Excellent, thanks. I agree libraries should avoid feature bloat. I was providing the reference to provide evidence that this enhancement wasn't just a wacky idea I dreamed up.  Not to suggest we should copy or do things just because someone else has. Thanks. I'll fix and submit a new patch tomorrow.
```
## 49.Why is `Plot.plotters` private ?
Sep 3, 2016
```text
 I'm super happy to use the plotutil plot-making functions, but I'd also want to be able to tweak the plotters after the fact, simply changing colors or line size.. I see that the Plot.plotters list of Plotters is internal only.. could we expose that somehow ? with a Plotters() function or directly opening the Plotters field ?
```
```text
 To have made a plot.Plot you must have had each of the plotters in your hand at some time. Can you not just modify them where you hold them (they are all pointer types). On the other hand, @eaburns, why is there an add rather than just appending to an exported slice?
```
```text
 I just read the source for Add and it's pretty clear that exposing the retainer plotters would be a bad idea; the plot calculates the min and max of each plotter's x and y values - user changes to these would break behaviour.

@abourget, you can still do what I said in retaining your pointers and mutating the plotters that way, but I don't think we should be encouraging that.
```
```text
 Dan, exactly. That's why.
```
```text
 Couldn't the plotutil Add functions return the plotters?
Le sam. 3 sept. 2016 09:03, Ethan Burns notifications@github.com a écrit :

Dan, exactly. That's why.
—
You are receiving this because you were mentioned.
Reply to this email directly, view it on GitHub
#312 (comment), or mute
the thread
https://github.com/notifications/unsubscribe-auth/AAFs8FM8SjOg93UdAY66OipdcrdHcFQrks5qmW_9gaJpZM4J0NL7
.
```
```text
 Maybe, but I forgot what those functions do, and I'm just on my phone right now.
```
## 50.plot: Can't change background and line color
Oct 5, 2018
```text
 What are you trying to do?
Changing the background and line color of a plot.
What did you do?

graph, err := plot.New()
if err != nil {
	return err
}

line, err := plotter.NewLine(dataPoints)
if err != nil {
	return err
}

graph.HideAxes()
graph.BackgroundColor = color.RGBA{A: 0}
line.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 237, A: 255}

err = plotutil.AddLines(graph, line)
if err != nil {
	return err
}

err = graph.Save(6*vg.Centimeter, vg.Centimeter, path)
if err != nil {
	return err
}

What did you expect to happen?
The background to be transparent and the line color to be blue.
What actually happened?
The image is created but the background is still white and the line color isn't blue.
What version of Go and Gonum/plot are you using?
Go: 1.8
gonum: 5f3c436
```
```text
 line is already a plot.Plotter.
plotutil.AddLines expects at least one plotter.XYer, and will create a new plot.Plotter value from that:


  
    
      plot/plotutil/add.go
    
    
         Line 192
      in
      5f3c436
    
    
    
    

        
          
           case plotter.XYer: 
        
    
  



so you'd just have to replace plotutil.AddLines(...) with: p.Add(line).
```
```text
 Thank you very much, that helped!
Do you know why the background isn't changing to transparent? Are there any restrictions to this? I'm saving the graph as a png which should support it
```
```text
 not sure whether this is a gonum/plot issue or one with the vg backend for png...
```
```text
 Maybe this is related?
https://github.com/gonum/plot/blob/master/vg/vgimg/vgimg.go#L104
```
```text
 definitely!
```
## 51.Reconsider using relative box and bar widths.
Mar 17, 2015
```text
 Original issue 133 created by eaburns on 2013-07-17T13:22:42.000Z:
Currently plotter.BoxPlot and plotter.BarChart (maybe others) use absolute values for their widths.  It can be hard to predict a good width, and currently one either needs to do very roundabout things (see this discussion: https://groups.google.com/d/topic/plotinum-discuss/WD02Ty30G2c/discussion) or to have a good guess.  Perhaps using a relative width would be better.  Can we make a nice API that supports both?
```
## 52.Feature Request: Easy way to cycle through colors
Sep 29, 2015
```text
 It's a very common desire to have the lines in a plot have different colors. This is not too difficult to accomplish if there are a small number of lines added in one place, but if there are a series of operations that each add to the plot (say, plot is passed along a channel to different cases) it's hard to write each module so that it can add a unique color without caring where it is in the chain.
The easiest design from the perspective of a user is that plot would have a "ColorIterator"
type ColorIterator interface {
    NextColor() color.RGBA
}

This would then be embedded by plot and set to a default iterator with plot.New. A subroutine could then choose the color of a line (for example) with a call to p.NextColor().
I don't know if this is the best design with respect to the rest of the package, some design along those lines would be very nice.
```
```text
 There is something that goes most of the way to this in plotutil. This uses globals and probably shouldn't to allow concurrent use, but if that were fixed would that satisfy your use cases?
```
```text
 Perhaps change the plotutil functions like so:
// Colors returns a function that cycles through the given colors.
// Each call to the returned function returns the next color in the slice,
// wrapping around after the last color is returned.
func Colors(colors []color.Color) func()(color.Color) {
    var i int
    return func() {
        c := DefaultColors[i%len(colors)]
        i++
        return c
    }
}

Then the user only needs to create a color function, and doesn't need to track the index themselves. It also fixes Dan's concern about globals.
c := plotutil.Colors(plotutil.SoftColors)
l0.Color = c()
l1.Color = c()
// etc.
```
```text
 @btracey, after re-reading your suggestion, it seems like you don't actually just want a simpler way to get the next color, but some way to track style information across complex control flow. Perhaps plotutil could include such a thing; something that holds a "next" function for all of colors, shapes, and dashes:
type plotBuilder {
    p *plot.Plot
    nextColor func()color.Color
    nextShape func()draw.GlyphDrawer
    nextDash func()[]vg.Length
}

// Add adds the plotter to the plot.
// If the plotter has a Color field, it is set to the nextColor in the builder. [use reflection]
// If the plotter has a ...
func (b *plotBuilder) Add(p plotter.Plotter) {
   // ...
}

Do folks think this is worth adding? In my opinion, it's not. I'd much prefer to keep the API smaller. Users who need such a thing can easily implement something for their own needs. It won't need to be as general as something in plotutil, so it can be much simpler.
```
## 53.Rename variables in padding code
Mar 15, 2015
```text
 Original issue 72 created by eaburns on 2012-08-24T13:12:20.000Z:
n could be newMin and m could be newMax.
```
```text
 Comment #1 originally posted by eaburns on 2013-02-20T19:22:07.000Z:
<empty>
```
## 54.plotutil.AddBarCharts
Mar 15, 2015
```text
 Original issue 96 created by eaburns on 2012-11-05T14:00:15.000Z:
Should make a plotutil function for adding bar charts.  It should also allow one to add grouped bar charts.  This may be possible by allowing not only a string parameter to name a bar chart, but also by allowing an int parameter to precede the data which specifies its offset.
```
```text
 Comment #1 originally posted by eaburns on 2013-02-20T19:25:14.000Z:
<empty>
```
## 55.Stacked bar chart example
Mar 15, 2015
```text
 Original issue 90 created by eaburns on 2012-09-23T21:38:19.000Z:
Put one up on the wiki.
```
```text
 Comment #1 originally posted by eaburns on 2013-02-20T19:27:48.000Z:
<empty>
```
## 56.Different behavior for vgpdf paths
Mar 15, 2015
```text
 Original issue 128 created by eaburns on 2013-06-08T12:44:54.000Z:
All other vg backends seem to insert an implicit Move to the start of a new path, and they seem to insert Lines between disjoint path components.
For example, if you draw two arcs with the same center point but different radii then call p.Close, vgpdf doesn't draw anything (because the path has a discontinuity) but all other backends draw a sector.
```
```text
 Comment #1 originally posted by eaburns on 2013-06-08T13:15:27.000Z:
<empty>
```
```text
 Comment #2 originally posted by eaburns on 2013-07-07T18:57:13.000Z:
<empty>
```
## 57.p.X.Label.Text
Jan 9, 2019
```text
 p.X.Label.Text  This variable does not support Chinese
```
```text
 Thanks for the report, but, could you be a bit more explicit?
What did you try? (With some code :))
What did you expect?
What did you get instead?
```
## 58.Allow each tick label to have its own TextStyle
Mar 15, 2015
```text
 Original issue 21 created by eaburns on 2012-06-13T19:38:10.000Z:
Add a *TextStyle field to Tick.  If it is nil then use the Axis.Tick.Label.TextStyle, otherwise use the pointed to style.
```
```text
 Comment #1 originally posted by eaburns on 2013-02-20T19:24:18.000Z:
<empty>
```
## 59.Rethink Palette/Contour/Heatmap interface
Feb 17, 2015
```text
 Palettes specify a list of colors. Strictly this is always a specific list, though may of the implemented palettes can create a long list of colors. These are being used at present to choose colors as part of Heatmaps and contours. In these cases, the z-axis values can be a real number, and a sort of interpolation function is being used to get the specific palette color. There are some special cases introduced, for example OverflowColor, which allow warning of specific colors.
In my mind, it's much better instead to have
// ColorFunc associates each float64 value with a color.Color
type ColorFunc func(float64) color.Color

A ColorFunc could use a palette as its underlying color choice. It could use a "threshholded palette", where it uses a palette for some intermediate range and uses warning colors outside of that range. There are many other possible uses, for example, returning a third warning color when the value is NaN, or having a grayscale image with a specific range of float64 returned as red to highlight their locations. It's much more powerful to have a function rather than just a palette with linear indexing, it's more composable because the special cases are located with the choice of color and not the choice of plotter, and (to me at least) it's more conceptually aligned with what's happening.
```
## 60.If the axis scale is LogScale, the major tic spacing should also be log scale
May 27, 2015
```text
 For example, on my example plot, the X axis is log scale, but the major tick marks are at "200, 400, 600" and "1e5, 2e5, 3e5, 4e5, 5e5". On a log scale this makes the labels closer together. It's better to have, say, powers of 10 like the automatic in matlab
```
## 61.plot/vg: consider using golang.org/x/image/font/sfnt for parsing fonts
Nov 23, 2016
```text
 right now, we are using "github.com/golang/freetype" and "github.com/golang/freetype/truetype".
They work fine, but the license is a dual one (BSD-like or GPL2+) : https://github.com/golang/freetype/blob/master/LICENSE
We should perhaps consider use instead the new:
https://godoc.org/golang.org/x/image/font/sfnt
(which is released under the same license than Go itself (BSD-3))
for this to happen, though, the following methods will need to be implemented for sfnt.Font:

sfnt.Font.Bounds
sfnt.Font.Index
sfnt.Font.Kern
sfnt.Font.HMetric
```
```text
 https://go-review.googlesource.com/c/37410/ added sfnt.Font.Kern.
```
```text
 https://go-review.googlesource.com/39670 added sfnt.Font.Bounds.
```
```text
 https://go-review.googlesource.com/#/c/67330/
is adding metrics to font.
```
```text
 freetype/truetype.Font.Index is spelled sfnt.GlyphIndex.
```
## 62.Y-axis precision for LogScale Ticks
Sep 27, 2017
```text
 Now it is not scientific notation. But when the values are quite large numbers, the plot might look ugly being  overloaded with zeros. And, as a thought, should we make the level of precision constant for the LogScale?
```
## 63.Double Y-axis possible?
Sep 13, 2017
```text
 I've got some data which I want to present using a double Y-axis. I can't find any examples or mentions of it anywhere though.
Is it possible to use a double Y-axis with Gonum plots?
```
```text
 Please see #235.
```
```text
 Do I understand it correctly that because you think dual y-axis plots are not good you're not going to implement them? I've scanned through the arguments in the papers and I agree that a double axis has  many drawbacks; relative comparisons as well as intersections are meaningless.
But despite those problems, I still need a double axis to be able to compare trends. In the plots I often look at I mostly don't even care about the values on the y-axis, I just look at whether the lines move up or down, that's all. And for my use case a double axis is perfect. The alternative (what I'm doing now) is drawing two plots and constantly flipping between them really fast to "simulate" an overlay, which is ridiculous to say the least.
In conclusion, I think a dual y-axis is like all data representations; it can be gravely misused, but for the proper use case (like mine) it's awesome.
Do you know of any alternative plotting lib for go which does support a dual y-axis?
```
```text
 One way to compare trends in data sets is to divide the points in each
dataset by its average. Then you can plot them both on the same scale.
…
On Thu, Sep 14, 2017, 6:00 PM kramer65 ***@***.***> wrote:
 Do I understand it correctly that because you think dual y-axis plots are
 not good you're not going to implement them? I've scanned through the
 arguments in the papers and I agree that a double axis has many drawbacks;
 relative comparisons as well as intersections are meaningless.

 But despite those problems, I still need a double axis to be able to
 compare trends. In the plots I often look at I don't even care about the
 values on the y-axis, I just look at whether the lines move up or down,
 that's all. And for my use case a double axis is perfect. The alternative
 (what I'm doing now) is drawing two plots and constantly flipping between
 them really fast to "simulate" an overlay, which is ridiculous to say the
 least.

 In conclusion, I think a dual y-axis is like all data representations; it
 can be gravely misused, but for the proper use case (like mine) it's
 awesome.

 Do you know of any alternative plotting lib for go which does support a
 dual y-axis?

 —
 You are receiving this because you are subscribed to this thread.
 Reply to this email directly, view it on GitHub
 <#387 (comment)>, or mute
 the thread
 <https://github.com/notifications/unsubscribe-auth/AF6AASjOxUUtC_q79BwvRFvs_ZMdMDRAks5sicxHgaJpZM4PWqDD>
 .
```
```text
 Since you don't care about the values on the aces, I would suggest just normalise to the range of the data and plot those normalised series.
But, yes, this is not something we are going to implement. Though note that in #235 (comment) there is a pathway to a state that would allow it without making it easy.
```
```text
 @ctessum - that is quite a good tip, thanks!
I still think double y-axis should be implemented, but I can use this trick to work around it.
```
## 64.vgeps: implement vg.Canvas.DrawImage
Apr 15, 2016
```text
 right now, vgeps.Canvas.DrawImage is just panic-ing.
here is the start of an implementation:
// DrawImage implements the vg.Canvas.DrawImage method.
func (c *Canvas) DrawImage(rect vg.Rectangle, img image.Image) {
    var (
        width   = rect.Size().X.Dots(DPI)
        height  = rect.Size().Y.Dots(DPI)
        iwidth  = img.Bounds().Dx()
        iheight = img.Bounds().Dy()
    )
    c.Push()
    fmt.Fprintf(c.buf, "\n%%%% Image data\n")
    c.Translate(rect.Min)
    fmt.Fprintf(c.buf, "%.*g %.*g scale     %% scale pixels to points\n", pr, width, pr, height)
    fmt.Fprintf(c.buf, "%d %d 8 [%d 0 0 %d 0 0]     ", iwidth, iheight, iwidth, iheight)
    fmt.Fprintf(c.buf, "%%%% width height colordepth transform\n")
    fmt.Fprintf(c.buf, "/datasource currentfile /ASCII85Decode filter /RunLengthDecode filter def\n")
    fmt.Fprintf(c.buf, "/datastring %d string def   ", iwidth*3)
    fmt.Fprintf(c.buf, "%% %d = width * color components\n", iwidth*3)
    fmt.Fprintf(c.buf, "{datasource datastring readstring pop}\n")
    fmt.Fprintf(c.buf, "false       %%%% false == single data source (rgb)\n")
    fmt.Fprintf(c.buf, "3   %%%% number of color components\n")
    fmt.Fprintf(c.buf, "colorimage\n")
    err := encodeEPSData(c.buf, img)
    c.Pop()
    if err != nil {
        panic(fmt.Errorf("vgeps: error encoding EPS data: %v", err))
    }
}

func encodeEPSData(w io.Writer, img image.Image) error {
    var (
        err  error
        cs   = new(bytes.Buffer)
        pix  [3]byte
        imax = img.Bounds().Dx()
        jmax = img.Bounds().Dy()
    )
    if err != nil {
        return err
    }
    for j := 0; j < jmax; j++ {
        for i := 0; i < imax; i++ {
            r, g, b, a := img.At(i, j).RGBA()
            switch a {
            case 0:
                pix = [3]byte{}
            default:
                pix[0] = uint8((r * 65535 / a) >> 8)
                pix[1] = uint8((g * 65535 / a) >> 8)
                pix[2] = uint8((b * 65535 / a) >> 8)
            }
            _, err = cs.Write(pix[:])
            if err != nil {
                return err
            }
        }
    }
    rle := rlEncode(cs.String())
    r := bytes.NewReader([]byte(rle))
    enc := ascii85.NewEncoder(w)
    _, err = io.Copy(enc, r)
    if err != nil {
        return err
    }

    err = enc.Close()
    if err != nil {
        return err
    }

    _, err = w.Write([]byte("~>\n"))
    return err
}

func rlEncode(s string) string {
    count := 1
    var o bytes.Buffer
    var i int
    for i = 0; i < len(s)-1; i++ {
        if s[i] == s[i+1] {
            count++
        } else {
            o.WriteString(fmt.Sprintf("%d%c", count, s[i]))
            count = 1
        }
    }
    o.WriteString(fmt.Sprintf("%d%c", count, s[i]))
    return o.String()
}
this was mostly extracted from:
https://github.com/eaburns/search/blob/master/graphics/image.cc#L102
https://github.com/eaburns/search/blob/master/utils/encode.cc#L19
```
## 65.How to LogScale in reverse?
Oct 6, 2016
```text
 I have duration in seconds on the Y scale and percentages on the X scale but when I set LogScale I get the top and right parts of the graph squished. How can I log scale in reverse so the later values are higher resolution instead?
```
```text
 That's not achievable in the general sense, at least not without further configuration. To do this you need to have a known upper bound, otherwise you are attempting to increase resolution as to also increase the size of the function range. Are you sure you really want to do this? Can you explain your end goal?
If you do have an upper bound, the trivial solution is to plot the difference from that upper bound, thus inverting the sign of your function. An alternative is to implement your own Normalizer and Ticker types that know your upper bound. To get that to work well, you will probably need to copy the plot.precisionOf and plot.formatFloatTick functions (maybe we should export them to aid this kind of endeavour.
```
```text
 The log scale it does currently looks like this on the X axis and compresses the tail.
-----------------------------------------------
0                    30 |        50 |   70 | 90          

But when measuring latency, showing the 90-100 percentiles in greater detail is more important than 0-30 percentiles so I would rather something that compresses the head like this:
-----------------------------------------------
0 | 30 | 50    | 70        | 90
```
```text
 Yeah, that makes sense. Either of the approaches I mention above would work for you.
```
```text
 If you do have an upper bound, the trivial solution is to plot the difference from that upper bound, thus inverting the sign of your function.

Could you explain or give an example? I'm not sure I understand.
```
```text
 The trivial approach is to convert the value you are plotting from a latency percentile one hundred minus the latency percentile (I can't think of a good name for this - which is probably a long-term argument against it). The alternative approach is to fork the code here and here to implement a Ticker and a Normalizer that does what you want. Actually, thinking more, you don't need the Ticker, since you can just use a ConstantTicks - then it's just the Normalizer, which is trivial. Something like this:
// ReverseLogScale can be used as the value of an Axis.Scale function to
// set the axis to a log scale.
type ReverseLogScale struct{}

// Normalize returns the fractional logarithmic distance of
// x between min and max.
func (ReverseLogScale) Normalize(min, max, x float64) float64 {
    logMin := math.Log(min)
    return 1 - ((math.Log(max-x) - logMin) / (math.Log(max) - logMin))
}

Demonstrated here.
```
```text
 Oh interesting, thank you! Although when I run this with the plotter the process never completes. It must be causing some endless loop I'm suspecting.
```
```text
 Try with an explicit 100 for the uses of max in Normalize, or set the axis max to 100.
```
```text
 It exits now correctly but there's no values or ticks along the axis and there's no visible line drawn.
```
```text
 Debugging code remotely without being able to see it is sort of tricky. If you can't sort this out, you should probably put it up somewhere for people to look at.
```
```text
 Making a plot with unusual axis scaling may be more confusing than helpful for the audience anyway. Another option could be to make two plots and put them side-by-side, where one plot has an axis spanning 0-100 percentile, and the other plot has an axis spanning 90-100 or whatever. To make multiple plots in one image you can use draw.Tiles{}.
```
```text
 I've seen this kind of scaling in this domain in the past, so I think this is a perfectly reasonable thing to do.

I am interested in seeing why this is not working though (and investigating the previous non-termination behaviour). To look into those I need a minimal reproducer.
```
```text
 This is an example of what I'm talking about:
https://twitter.com/benalexau/status/790782407301435392
```
## 66.all: put ExampleFoo in pkg_test instead of pkg
Jan 8, 2019
```text
 ie: instead of:
package plotter

func ExampleScatter_color() {
	// ...
	sc, err := NewScatter(scatterData)
	if err != nil {
		log.Panic(err)
	}
}
have the example be written as:
package plotter_test

func ExampleScatter_color() {
	// ...
	sc, err := plotter.NewScatter(scatterData)
	if err != nil {
		log.Panic(err)
	}
}
so the thing we want to explain is properly qualified.
```
## 67.Default tick marks should be more "balanced"
Mar 15, 2015
```text
 Original issue 114 created by eaburns on 2013-02-20T19:08:16.000Z:
In the plot below, the tick marks are "top heavy" (or "right heavy" if you consider the X axis).  If the major and minor ticks were swapped then the major tick marks would look more balanced.  It may be easy to get more balanced ticks by looking at the mean value of the major ticks and comparing it to the center of the axis.  The closer the mean is to the center, the more balanced the ticks are.
package main

import (
    "code.google.com/p/plotinum/plot"
    "code.google.com/p/plotinum/plotter"
    "code.google.com/p/plotinum/plotutil"
)

var (
    newSched = speedup{
        { 1, 23.857 },
        { 2, 12.183 },
        { 4, 6.557 },
        { 8, 4.338 },
        { 16, 3.478 },
        { 32, 2.385 },
    }
    oldSched = speedup{
        { 1, 23.163},
        { 2, 22.182},
        { 4, 25.568 },
        { 8, 31.372},
        { 16, 27.275 },
        { 32,26.074 },
    }
)

func main() {
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    plotutil.AddLinePoints(p, "new", newSched, "old", oldSched)
    p.Add(plotter.NewFunction(func(x float64)float64{ return x }))

    p.Title.Text = "Speedup of schedulers"
    p.X.Label.Text = "proc"
    p.Y.Label.Text = "speedup (of real time)"
    p.Legend.Top = true
    p.Legend.Left = true
    p.Y.Max = 32
    p.Save(4, 4, "speedup.svg")
}

type speedup plotter.XYs

func (s speedup) Len() int {
    return plotter.XYs(s).Len()
}

func (s speedup) XY(i int) (float64, float64) {
    x, y := plotter.XYs(s).XY(i)
    _, y0 := plotter.XYs(s).XY(0)
    return x, y0/y
}
```
```text
 Comment #1 originally posted by eaburns on 2013-02-20T19:09:07.000Z:
Hopefully someone else wants to play with this.
```
## 68.plotter: Line does not respect X-axis range
Dec 19, 2018
```text
 when a plotter.Line is given a fill-color (spelled ShadeColor for now) and one restricts the X axis range after the facts, the fill color bleeds out of the plot.
// +build ignore

package main

import (
	"image/color"
	"log"
	"math/rand"

	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	rnd := rand.New(rand.NewSource(1))

	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	randomPoints := func(n int) plotter.XYs {
		pts := make(plotter.XYs, n)
		for i := range pts {
			pts[i].X = float64(i)
			pts[i].Y = float64(10 * rnd.Float64())
		}
		return pts
	}

	n := 20
	data := randomPoints(n)

	tp := hplot.NewTiledPlot(draw.Tiles{Cols: 3, Rows: 3})
	for i := 0; i < 9; i++ {
		if i != 4 {
			tp.Plots[i] = nil
			continue
		}
		p := tp.Plots[i]
		p.X.Label.Text = "X"
		p.Y.Label.Text = "Y"
		p.Add(plotter.NewGrid())

		line, points, err := plotter.NewLinePoints(data)
		if err != nil {
			log.Panic(err)
		}
		line.Color = color.RGBA{G: 255, A: 255}
		line.ShadeColor = &line.Color
		points.Shape = draw.CircleGlyph{}
		points.Color = color.RGBA{R: 255, A: 255}

		p.Add(line)
		p.Add(points)
	}

	tp.Plots[4].X.Min = 3
	tp.Plots[4].X.Max = 13
	err := tp.Save(30*vg.Centimeter, 25*vg.Centimeter, "out.png")
	if err != nil {
		log.Panic(err)
	}
}
this is what we get:

this is what we want:
```
## 69.OpenGL back-end for vg
Mar 15, 2015
```text
 Original issue 110 created by eaburns on 2013-02-13T16:46:54.000Z:
This should be pretty easy, but I am not sure of a great way to do dashed lines in OpenGL, maybe someone with shader experience [Steve] could figure it out.
An OpenGL back-end would make it much easier to make a quality GUI front-end for Plotinum---something that has been requested a few times at this point.
```
```text
 Comment #1 originally posted by eaburns on 2013-02-20T19:26:48.000Z:
<empty>
```
## 70.Actual tests
Mar 15, 2015
```text
 Original issue 29 created by eaburns on 2012-06-15T15:39:34.000Z:
Currently, the tests don't really test anything.  They are just used as a convenient way to look at sample plots as I code.  There should be some actual tests.
```
```text
 Comment #1 originally posted by eaburns on 2012-09-21T12:23:33.000Z:
<empty>
```
```text
 Comment #2 originally posted by eaburns on 2013-03-12T15:36:51.000Z:
Whoever can feel free to add some of these.  We really should have a bunch more tests before the code grows any bigger.
```
```text
 Comment #3 originally posted by eaburns on 2013-06-08T13:16:23.000Z:
<empty>
```
```text
 Comment #4 originally posted by eaburns on 2013-07-07T18:54:06.000Z:
<empty>
```
```text
 Comment #5 originally posted by eaburns on 2013-07-07T18:54:18.000Z:
<empty>
```
## 71.Make a gallery page
Mar 15, 2015
```text
 Original issue 91 created by eaburns on 2012-09-23T21:39:52.000Z:
This would be like the examples page but would just contain a bunch of pretty images.  The point would be to post nice looking images to show off Plotinum's capabilities but without posting the details of where all of the data came from and how the plot with created.
```
```text
 Comment #1 originally posted by eaburns on 2013-06-08T13:16:03.000Z:
<empty>
```
## 72.plot/vg: consider using Go-font for fonts
Nov 17, 2016
```text
 more details here: https://blog.golang.org/go-fonts
(the advantage would be license. but we should probably check whether it fits gonum/plot rendering wrt PS)
```
```text
 The blog post said that they are mostly metric equivalent, so it should work. Let's try it.
```
```text
 hum... not sure the gofonts have everything we want:
    FontMap = map[string]string{
        "Courier":             "GoMono-Regular",
        "Courier-Bold":        "GoMono-Bold",
        "Courier-Oblique":     "GoMono-Italic",
        "Courier-BoldOblique": "GoMono-BoldItalic",

        "Helvetica":             "GoSans-Regular",
        "Helvetica-Bold":        "GoSans-Bold",
        "Helvetica-Oblique":     "GoSans-Italic",
        "Helvetica-BoldOblique": "GoSans-BoldItalic",

        "Times-Roman":      "GoSerif-Regular",
        "Times-Bold":       "GoSerif-Bold",
        "Times-Italic":     "GoSerif-Italic",
        "Times-BoldItalic": "GoSerif-BoldItalic",
    }
GoMono-xyz is, obviously, "golang.org/x/image/font/gofont/gomono{,bold,bolditalic,italic}"
GoSans-xyz is probably, "golang.org/x/image/font/gofont/go{regular,bold,bolditalic,italic}"
but what about the GoSerif-xyz ones?
```
```text
 Yeah, hm, those have no match. :-(
```
```text
 I've filed golang/go#17964
```
