package main

// Xeol - Xcanner for end-of-life software in container images, filesystems, and SBOMs
// Homepage: https://github.com/xeol-io/xeol

import (
	"fmt"
	
	"os/exec"
)

func installXeol() {
	// Método 1: Descargar y extraer .tar.gz
	xeol_tar_url := "https://github.com/xeol-io/xeol/archive/refs/tags/v0.10.0.tar.gz"
	xeol_cmd_tar := exec.Command("curl", "-L", xeol_tar_url, "-o", "package.tar.gz")
	err := xeol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xeol_zip_url := "https://github.com/xeol-io/xeol/archive/refs/tags/v0.10.0.zip"
	xeol_cmd_zip := exec.Command("curl", "-L", xeol_zip_url, "-o", "package.zip")
	err = xeol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xeol_bin_url := "https://github.com/xeol-io/xeol/archive/refs/tags/v0.10.0.bin"
	xeol_cmd_bin := exec.Command("curl", "-L", xeol_bin_url, "-o", "binary.bin")
	err = xeol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xeol_src_url := "https://github.com/xeol-io/xeol/archive/refs/tags/v0.10.0.src.tar.gz"
	xeol_cmd_src := exec.Command("curl", "-L", xeol_src_url, "-o", "source.tar.gz")
	err = xeol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xeol_cmd_direct := exec.Command("./binary")
	err = xeol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
