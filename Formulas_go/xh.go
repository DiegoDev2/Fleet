package main

// Xh - Friendly and fast tool for sending HTTP requests
// Homepage: https://github.com/ducaale/xh

import (
	"fmt"
	
	"os/exec"
)

func installXh() {
	// Método 1: Descargar y extraer .tar.gz
	xh_tar_url := "https://github.com/ducaale/xh/archive/refs/tags/v0.22.2.tar.gz"
	xh_cmd_tar := exec.Command("curl", "-L", xh_tar_url, "-o", "package.tar.gz")
	err := xh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xh_zip_url := "https://github.com/ducaale/xh/archive/refs/tags/v0.22.2.zip"
	xh_cmd_zip := exec.Command("curl", "-L", xh_zip_url, "-o", "package.zip")
	err = xh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xh_bin_url := "https://github.com/ducaale/xh/archive/refs/tags/v0.22.2.bin"
	xh_cmd_bin := exec.Command("curl", "-L", xh_bin_url, "-o", "binary.bin")
	err = xh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xh_src_url := "https://github.com/ducaale/xh/archive/refs/tags/v0.22.2.src.tar.gz"
	xh_cmd_src := exec.Command("curl", "-L", xh_src_url, "-o", "source.tar.gz")
	err = xh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xh_cmd_direct := exec.Command("./binary")
	err = xh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
