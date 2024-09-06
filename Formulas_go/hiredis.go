package main

// Hiredis - Minimalistic client for Redis
// Homepage: https://github.com/redis/hiredis

import (
	"fmt"
	
	"os/exec"
)

func installHiredis() {
	// Método 1: Descargar y extraer .tar.gz
	hiredis_tar_url := "https://github.com/redis/hiredis/archive/refs/tags/v1.2.0.tar.gz"
	hiredis_cmd_tar := exec.Command("curl", "-L", hiredis_tar_url, "-o", "package.tar.gz")
	err := hiredis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hiredis_zip_url := "https://github.com/redis/hiredis/archive/refs/tags/v1.2.0.zip"
	hiredis_cmd_zip := exec.Command("curl", "-L", hiredis_zip_url, "-o", "package.zip")
	err = hiredis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hiredis_bin_url := "https://github.com/redis/hiredis/archive/refs/tags/v1.2.0.bin"
	hiredis_cmd_bin := exec.Command("curl", "-L", hiredis_bin_url, "-o", "binary.bin")
	err = hiredis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hiredis_src_url := "https://github.com/redis/hiredis/archive/refs/tags/v1.2.0.src.tar.gz"
	hiredis_cmd_src := exec.Command("curl", "-L", hiredis_src_url, "-o", "source.tar.gz")
	err = hiredis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hiredis_cmd_direct := exec.Command("./binary")
	err = hiredis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
