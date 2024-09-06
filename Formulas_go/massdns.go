package main

// Massdns - High-performance DNS stub resolver
// Homepage: https://github.com/blechschmidt/massdns

import (
	"fmt"
	
	"os/exec"
)

func installMassdns() {
	// Método 1: Descargar y extraer .tar.gz
	massdns_tar_url := "https://github.com/blechschmidt/massdns/archive/refs/tags/v1.1.0.tar.gz"
	massdns_cmd_tar := exec.Command("curl", "-L", massdns_tar_url, "-o", "package.tar.gz")
	err := massdns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	massdns_zip_url := "https://github.com/blechschmidt/massdns/archive/refs/tags/v1.1.0.zip"
	massdns_cmd_zip := exec.Command("curl", "-L", massdns_zip_url, "-o", "package.zip")
	err = massdns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	massdns_bin_url := "https://github.com/blechschmidt/massdns/archive/refs/tags/v1.1.0.bin"
	massdns_cmd_bin := exec.Command("curl", "-L", massdns_bin_url, "-o", "binary.bin")
	err = massdns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	massdns_src_url := "https://github.com/blechschmidt/massdns/archive/refs/tags/v1.1.0.src.tar.gz"
	massdns_cmd_src := exec.Command("curl", "-L", massdns_src_url, "-o", "source.tar.gz")
	err = massdns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	massdns_cmd_direct := exec.Command("./binary")
	err = massdns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
