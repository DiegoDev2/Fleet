package main

// Godap - Complete TUI (terminal user interface) for LDAP
// Homepage: https://github.com/Macmod/godap

import (
	"fmt"
	
	"os/exec"
)

func installGodap() {
	// Método 1: Descargar y extraer .tar.gz
	godap_tar_url := "https://github.com/Macmod/godap/archive/refs/tags/v2.7.3.tar.gz"
	godap_cmd_tar := exec.Command("curl", "-L", godap_tar_url, "-o", "package.tar.gz")
	err := godap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	godap_zip_url := "https://github.com/Macmod/godap/archive/refs/tags/v2.7.3.zip"
	godap_cmd_zip := exec.Command("curl", "-L", godap_zip_url, "-o", "package.zip")
	err = godap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	godap_bin_url := "https://github.com/Macmod/godap/archive/refs/tags/v2.7.3.bin"
	godap_cmd_bin := exec.Command("curl", "-L", godap_bin_url, "-o", "binary.bin")
	err = godap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	godap_src_url := "https://github.com/Macmod/godap/archive/refs/tags/v2.7.3.src.tar.gz"
	godap_cmd_src := exec.Command("curl", "-L", godap_src_url, "-o", "source.tar.gz")
	err = godap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	godap_cmd_direct := exec.Command("./binary")
	err = godap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
