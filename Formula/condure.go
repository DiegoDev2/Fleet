package main

// Condure - HTTP/WebSocket connection manager
// Homepage: https://github.com/fanout/condure

import (
	"fmt"
	
	"os/exec"
)

func installCondure() {
	// Método 1: Descargar y extraer .tar.gz
	condure_tar_url := "https://github.com/fanout/condure/archive/refs/tags/1.10.1.tar.gz"
	condure_cmd_tar := exec.Command("curl", "-L", condure_tar_url, "-o", "package.tar.gz")
	err := condure_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	condure_zip_url := "https://github.com/fanout/condure/archive/refs/tags/1.10.1.zip"
	condure_cmd_zip := exec.Command("curl", "-L", condure_zip_url, "-o", "package.zip")
	err = condure_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	condure_bin_url := "https://github.com/fanout/condure/archive/refs/tags/1.10.1.bin"
	condure_cmd_bin := exec.Command("curl", "-L", condure_bin_url, "-o", "binary.bin")
	err = condure_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	condure_src_url := "https://github.com/fanout/condure/archive/refs/tags/1.10.1.src.tar.gz"
	condure_cmd_src := exec.Command("curl", "-L", condure_src_url, "-o", "source.tar.gz")
	err = condure_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	condure_cmd_direct := exec.Command("./binary")
	err = condure_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: cython")
	exec.Command("latte", "install", "cython").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
}
