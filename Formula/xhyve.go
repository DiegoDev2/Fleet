package main

// Xhyve - Lightweight macOS virtualization solution based on FreeBSD's bhyve
// Homepage: https://github.com/machyve/xhyve

import (
	"fmt"
	
	"os/exec"
)

func installXhyve() {
	// Método 1: Descargar y extraer .tar.gz
	xhyve_tar_url := "https://github.com/machyve/xhyve/archive/refs/tags/v0.2.0.tar.gz"
	xhyve_cmd_tar := exec.Command("curl", "-L", xhyve_tar_url, "-o", "package.tar.gz")
	err := xhyve_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xhyve_zip_url := "https://github.com/machyve/xhyve/archive/refs/tags/v0.2.0.zip"
	xhyve_cmd_zip := exec.Command("curl", "-L", xhyve_zip_url, "-o", "package.zip")
	err = xhyve_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xhyve_bin_url := "https://github.com/machyve/xhyve/archive/refs/tags/v0.2.0.bin"
	xhyve_cmd_bin := exec.Command("curl", "-L", xhyve_bin_url, "-o", "binary.bin")
	err = xhyve_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xhyve_src_url := "https://github.com/machyve/xhyve/archive/refs/tags/v0.2.0.src.tar.gz"
	xhyve_cmd_src := exec.Command("curl", "-L", xhyve_src_url, "-o", "source.tar.gz")
	err = xhyve_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xhyve_cmd_direct := exec.Command("./binary")
	err = xhyve_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
