package main

// River - Reverse proxy application, based on the pingora library from Cloudflare
// Homepage: https://github.com/memorysafety/river

import (
	"fmt"
	
	"os/exec"
)

func installRiver() {
	// Método 1: Descargar y extraer .tar.gz
	river_tar_url := "https://github.com/memorysafety/river/archive/refs/tags/v0.5.0.tar.gz"
	river_cmd_tar := exec.Command("curl", "-L", river_tar_url, "-o", "package.tar.gz")
	err := river_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	river_zip_url := "https://github.com/memorysafety/river/archive/refs/tags/v0.5.0.zip"
	river_cmd_zip := exec.Command("curl", "-L", river_zip_url, "-o", "package.zip")
	err = river_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	river_bin_url := "https://github.com/memorysafety/river/archive/refs/tags/v0.5.0.bin"
	river_cmd_bin := exec.Command("curl", "-L", river_bin_url, "-o", "binary.bin")
	err = river_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	river_src_url := "https://github.com/memorysafety/river/archive/refs/tags/v0.5.0.src.tar.gz"
	river_cmd_src := exec.Command("curl", "-L", river_src_url, "-o", "source.tar.gz")
	err = river_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	river_cmd_direct := exec.Command("./binary")
	err = river_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
