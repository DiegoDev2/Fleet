package main

// Libfyaml - Fully feature complete YAML parser and emitter
// Homepage: https://github.com/pantoniou/libfyaml

import (
	"fmt"
	
	"os/exec"
)

func installLibfyaml() {
	// Método 1: Descargar y extraer .tar.gz
	libfyaml_tar_url := "https://github.com/pantoniou/libfyaml/releases/download/v0.9/libfyaml-0.9.tar.gz"
	libfyaml_cmd_tar := exec.Command("curl", "-L", libfyaml_tar_url, "-o", "package.tar.gz")
	err := libfyaml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfyaml_zip_url := "https://github.com/pantoniou/libfyaml/releases/download/v0.9/libfyaml-0.9.zip"
	libfyaml_cmd_zip := exec.Command("curl", "-L", libfyaml_zip_url, "-o", "package.zip")
	err = libfyaml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfyaml_bin_url := "https://github.com/pantoniou/libfyaml/releases/download/v0.9/libfyaml-0.9.bin"
	libfyaml_cmd_bin := exec.Command("curl", "-L", libfyaml_bin_url, "-o", "binary.bin")
	err = libfyaml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfyaml_src_url := "https://github.com/pantoniou/libfyaml/releases/download/v0.9/libfyaml-0.9.src.tar.gz"
	libfyaml_cmd_src := exec.Command("curl", "-L", libfyaml_src_url, "-o", "source.tar.gz")
	err = libfyaml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfyaml_cmd_direct := exec.Command("./binary")
	err = libfyaml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
