package main

// Ktmpl - Parameterized templates for Kubernetes manifests
// Homepage: https://github.com/jimmycuadra/ktmpl

import (
	"fmt"
	
	"os/exec"
)

func installKtmpl() {
	// Método 1: Descargar y extraer .tar.gz
	ktmpl_tar_url := "https://github.com/jimmycuadra/ktmpl/archive/refs/tags/0.9.1.tar.gz"
	ktmpl_cmd_tar := exec.Command("curl", "-L", ktmpl_tar_url, "-o", "package.tar.gz")
	err := ktmpl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ktmpl_zip_url := "https://github.com/jimmycuadra/ktmpl/archive/refs/tags/0.9.1.zip"
	ktmpl_cmd_zip := exec.Command("curl", "-L", ktmpl_zip_url, "-o", "package.zip")
	err = ktmpl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ktmpl_bin_url := "https://github.com/jimmycuadra/ktmpl/archive/refs/tags/0.9.1.bin"
	ktmpl_cmd_bin := exec.Command("curl", "-L", ktmpl_bin_url, "-o", "binary.bin")
	err = ktmpl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ktmpl_src_url := "https://github.com/jimmycuadra/ktmpl/archive/refs/tags/0.9.1.src.tar.gz"
	ktmpl_cmd_src := exec.Command("curl", "-L", ktmpl_src_url, "-o", "source.tar.gz")
	err = ktmpl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ktmpl_cmd_direct := exec.Command("./binary")
	err = ktmpl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
