package main

// Tcping - TCP connect to the given IP/port combo
// Homepage: https://github.com/mkirchner/tcping

import (
	"fmt"
	
	"os/exec"
)

func installTcping() {
	// Método 1: Descargar y extraer .tar.gz
	tcping_tar_url := "https://github.com/mkirchner/tcping/archive/refs/tags/2.1.0.tar.gz"
	tcping_cmd_tar := exec.Command("curl", "-L", tcping_tar_url, "-o", "package.tar.gz")
	err := tcping_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcping_zip_url := "https://github.com/mkirchner/tcping/archive/refs/tags/2.1.0.zip"
	tcping_cmd_zip := exec.Command("curl", "-L", tcping_zip_url, "-o", "package.zip")
	err = tcping_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcping_bin_url := "https://github.com/mkirchner/tcping/archive/refs/tags/2.1.0.bin"
	tcping_cmd_bin := exec.Command("curl", "-L", tcping_bin_url, "-o", "binary.bin")
	err = tcping_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcping_src_url := "https://github.com/mkirchner/tcping/archive/refs/tags/2.1.0.src.tar.gz"
	tcping_cmd_src := exec.Command("curl", "-L", tcping_src_url, "-o", "source.tar.gz")
	err = tcping_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcping_cmd_direct := exec.Command("./binary")
	err = tcping_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
