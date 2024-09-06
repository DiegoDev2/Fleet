package main

// Lighthouse - Rust Ethereum 2.0 Client
// Homepage: https://lighthouse.sigmaprime.io/

import (
	"fmt"
	
	"os/exec"
)

func installLighthouse() {
	// Método 1: Descargar y extraer .tar.gz
	lighthouse_tar_url := "https://github.com/sigp/lighthouse/archive/refs/tags/v5.3.0.tar.gz"
	lighthouse_cmd_tar := exec.Command("curl", "-L", lighthouse_tar_url, "-o", "package.tar.gz")
	err := lighthouse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lighthouse_zip_url := "https://github.com/sigp/lighthouse/archive/refs/tags/v5.3.0.zip"
	lighthouse_cmd_zip := exec.Command("curl", "-L", lighthouse_zip_url, "-o", "package.zip")
	err = lighthouse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lighthouse_bin_url := "https://github.com/sigp/lighthouse/archive/refs/tags/v5.3.0.bin"
	lighthouse_cmd_bin := exec.Command("curl", "-L", lighthouse_bin_url, "-o", "binary.bin")
	err = lighthouse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lighthouse_src_url := "https://github.com/sigp/lighthouse/archive/refs/tags/v5.3.0.src.tar.gz"
	lighthouse_cmd_src := exec.Command("curl", "-L", lighthouse_src_url, "-o", "source.tar.gz")
	err = lighthouse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lighthouse_cmd_direct := exec.Command("./binary")
	err = lighthouse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
