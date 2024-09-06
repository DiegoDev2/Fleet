package main

// Baresip - Modular SIP useragent
// Homepage: https://github.com/baresip/baresip

import (
	"fmt"
	
	"os/exec"
)

func installBaresip() {
	// Método 1: Descargar y extraer .tar.gz
	baresip_tar_url := "https://github.com/baresip/baresip/archive/refs/tags/v3.15.0.tar.gz"
	baresip_cmd_tar := exec.Command("curl", "-L", baresip_tar_url, "-o", "package.tar.gz")
	err := baresip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	baresip_zip_url := "https://github.com/baresip/baresip/archive/refs/tags/v3.15.0.zip"
	baresip_cmd_zip := exec.Command("curl", "-L", baresip_zip_url, "-o", "package.zip")
	err = baresip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	baresip_bin_url := "https://github.com/baresip/baresip/archive/refs/tags/v3.15.0.bin"
	baresip_cmd_bin := exec.Command("curl", "-L", baresip_bin_url, "-o", "binary.bin")
	err = baresip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	baresip_src_url := "https://github.com/baresip/baresip/archive/refs/tags/v3.15.0.src.tar.gz"
	baresip_cmd_src := exec.Command("curl", "-L", baresip_src_url, "-o", "source.tar.gz")
	err = baresip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	baresip_cmd_direct := exec.Command("./binary")
	err = baresip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libre")
exec.Command("latte", "install", "libre")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
