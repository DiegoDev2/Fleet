package main

// Ice - Comprehensive RPC framework
// Homepage: https://zeroc.com

import (
	"fmt"
	
	"os/exec"
)

func installIce() {
	// Método 1: Descargar y extraer .tar.gz
	ice_tar_url := "https://github.com/zeroc-ice/ice/archive/refs/tags/v3.7.10.tar.gz"
	ice_cmd_tar := exec.Command("curl", "-L", ice_tar_url, "-o", "package.tar.gz")
	err := ice_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ice_zip_url := "https://github.com/zeroc-ice/ice/archive/refs/tags/v3.7.10.zip"
	ice_cmd_zip := exec.Command("curl", "-L", ice_zip_url, "-o", "package.zip")
	err = ice_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ice_bin_url := "https://github.com/zeroc-ice/ice/archive/refs/tags/v3.7.10.bin"
	ice_cmd_bin := exec.Command("curl", "-L", ice_bin_url, "-o", "binary.bin")
	err = ice_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ice_src_url := "https://github.com/zeroc-ice/ice/archive/refs/tags/v3.7.10.src.tar.gz"
	ice_cmd_src := exec.Command("curl", "-L", ice_src_url, "-o", "source.tar.gz")
	err = ice_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ice_cmd_direct := exec.Command("./binary")
	err = ice_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lmdb")
exec.Command("latte", "install", "lmdb")
	fmt.Println("Instalando dependencia: mcpp")
exec.Command("latte", "install", "mcpp")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
