package main

// Jrsonnet - Rust implementation of Jsonnet language
// Homepage: https://github.com/CertainLach/jrsonnet

import (
	"fmt"
	
	"os/exec"
)

func installJrsonnet() {
	// Método 1: Descargar y extraer .tar.gz
	jrsonnet_tar_url := "https://github.com/CertainLach/jrsonnet/archive/refs/tags/v0.4.2.tar.gz"
	jrsonnet_cmd_tar := exec.Command("curl", "-L", jrsonnet_tar_url, "-o", "package.tar.gz")
	err := jrsonnet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jrsonnet_zip_url := "https://github.com/CertainLach/jrsonnet/archive/refs/tags/v0.4.2.zip"
	jrsonnet_cmd_zip := exec.Command("curl", "-L", jrsonnet_zip_url, "-o", "package.zip")
	err = jrsonnet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jrsonnet_bin_url := "https://github.com/CertainLach/jrsonnet/archive/refs/tags/v0.4.2.bin"
	jrsonnet_cmd_bin := exec.Command("curl", "-L", jrsonnet_bin_url, "-o", "binary.bin")
	err = jrsonnet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jrsonnet_src_url := "https://github.com/CertainLach/jrsonnet/archive/refs/tags/v0.4.2.src.tar.gz"
	jrsonnet_cmd_src := exec.Command("curl", "-L", jrsonnet_src_url, "-o", "source.tar.gz")
	err = jrsonnet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jrsonnet_cmd_direct := exec.Command("./binary")
	err = jrsonnet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
