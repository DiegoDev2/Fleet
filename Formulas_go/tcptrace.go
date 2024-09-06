package main

// Tcptrace - Analyze tcpdump output
// Homepage: https://web.archive.org/web/20210826120800/http://www.tcptrace.org/

import (
	"fmt"
	
	"os/exec"
)

func installTcptrace() {
	// Método 1: Descargar y extraer .tar.gz
	tcptrace_tar_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/tcptrace/tcptrace-6.6.7.tar.gz"
	tcptrace_cmd_tar := exec.Command("curl", "-L", tcptrace_tar_url, "-o", "package.tar.gz")
	err := tcptrace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcptrace_zip_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/tcptrace/tcptrace-6.6.7.zip"
	tcptrace_cmd_zip := exec.Command("curl", "-L", tcptrace_zip_url, "-o", "package.zip")
	err = tcptrace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcptrace_bin_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/tcptrace/tcptrace-6.6.7.bin"
	tcptrace_cmd_bin := exec.Command("curl", "-L", tcptrace_bin_url, "-o", "binary.bin")
	err = tcptrace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcptrace_src_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/tcptrace/tcptrace-6.6.7.src.tar.gz"
	tcptrace_cmd_src := exec.Command("curl", "-L", tcptrace_src_url, "-o", "source.tar.gz")
	err = tcptrace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcptrace_cmd_direct := exec.Command("./binary")
	err = tcptrace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
