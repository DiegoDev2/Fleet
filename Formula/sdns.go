package main

// Sdns - Privacy important, fast, recursive dns resolver server with dnssec support
// Homepage: https://sdns.dev

import (
	"fmt"
	
	"os/exec"
)

func installSdns() {
	// Método 1: Descargar y extraer .tar.gz
	sdns_tar_url := "https://github.com/semihalev/sdns/archive/refs/tags/v1.3.7.tar.gz"
	sdns_cmd_tar := exec.Command("curl", "-L", sdns_tar_url, "-o", "package.tar.gz")
	err := sdns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdns_zip_url := "https://github.com/semihalev/sdns/archive/refs/tags/v1.3.7.zip"
	sdns_cmd_zip := exec.Command("curl", "-L", sdns_zip_url, "-o", "package.zip")
	err = sdns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdns_bin_url := "https://github.com/semihalev/sdns/archive/refs/tags/v1.3.7.bin"
	sdns_cmd_bin := exec.Command("curl", "-L", sdns_bin_url, "-o", "binary.bin")
	err = sdns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdns_src_url := "https://github.com/semihalev/sdns/archive/refs/tags/v1.3.7.src.tar.gz"
	sdns_cmd_src := exec.Command("curl", "-L", sdns_src_url, "-o", "source.tar.gz")
	err = sdns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdns_cmd_direct := exec.Command("./binary")
	err = sdns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
