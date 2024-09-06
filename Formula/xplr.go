package main

// Xplr - Hackable, minimal, fast TUI file explorer
// Homepage: https://github.com/sayanarijit/xplr

import (
	"fmt"
	
	"os/exec"
)

func installXplr() {
	// Método 1: Descargar y extraer .tar.gz
	xplr_tar_url := "https://github.com/sayanarijit/xplr/archive/refs/tags/v0.21.9.tar.gz"
	xplr_cmd_tar := exec.Command("curl", "-L", xplr_tar_url, "-o", "package.tar.gz")
	err := xplr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xplr_zip_url := "https://github.com/sayanarijit/xplr/archive/refs/tags/v0.21.9.zip"
	xplr_cmd_zip := exec.Command("curl", "-L", xplr_zip_url, "-o", "package.zip")
	err = xplr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xplr_bin_url := "https://github.com/sayanarijit/xplr/archive/refs/tags/v0.21.9.bin"
	xplr_cmd_bin := exec.Command("curl", "-L", xplr_bin_url, "-o", "binary.bin")
	err = xplr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xplr_src_url := "https://github.com/sayanarijit/xplr/archive/refs/tags/v0.21.9.src.tar.gz"
	xplr_cmd_src := exec.Command("curl", "-L", xplr_src_url, "-o", "source.tar.gz")
	err = xplr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xplr_cmd_direct := exec.Command("./binary")
	err = xplr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
}
