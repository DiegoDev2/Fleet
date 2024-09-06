package main

// Autocannon - Fast HTTP/1.1 benchmarking tool written in Node.js
// Homepage: https://github.com/mcollina/autocannon

import (
	"fmt"
	
	"os/exec"
)

func installAutocannon() {
	// Método 1: Descargar y extraer .tar.gz
	autocannon_tar_url := "https://registry.npmjs.org/autocannon/-/autocannon-7.15.0.tgz"
	autocannon_cmd_tar := exec.Command("curl", "-L", autocannon_tar_url, "-o", "package.tar.gz")
	err := autocannon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autocannon_zip_url := "https://registry.npmjs.org/autocannon/-/autocannon-7.15.0.tgz"
	autocannon_cmd_zip := exec.Command("curl", "-L", autocannon_zip_url, "-o", "package.zip")
	err = autocannon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autocannon_bin_url := "https://registry.npmjs.org/autocannon/-/autocannon-7.15.0.tgz"
	autocannon_cmd_bin := exec.Command("curl", "-L", autocannon_bin_url, "-o", "binary.bin")
	err = autocannon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autocannon_src_url := "https://registry.npmjs.org/autocannon/-/autocannon-7.15.0.tgz"
	autocannon_cmd_src := exec.Command("curl", "-L", autocannon_src_url, "-o", "source.tar.gz")
	err = autocannon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autocannon_cmd_direct := exec.Command("./binary")
	err = autocannon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
