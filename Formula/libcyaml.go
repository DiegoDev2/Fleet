package main

// Libcyaml - C library for reading and writing YAML
// Homepage: https://github.com/tlsa/libcyaml

import (
	"fmt"
	
	"os/exec"
)

func installLibcyaml() {
	// Método 1: Descargar y extraer .tar.gz
	libcyaml_tar_url := "https://github.com/tlsa/libcyaml/archive/refs/tags/v1.4.2.tar.gz"
	libcyaml_cmd_tar := exec.Command("curl", "-L", libcyaml_tar_url, "-o", "package.tar.gz")
	err := libcyaml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcyaml_zip_url := "https://github.com/tlsa/libcyaml/archive/refs/tags/v1.4.2.zip"
	libcyaml_cmd_zip := exec.Command("curl", "-L", libcyaml_zip_url, "-o", "package.zip")
	err = libcyaml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcyaml_bin_url := "https://github.com/tlsa/libcyaml/archive/refs/tags/v1.4.2.bin"
	libcyaml_cmd_bin := exec.Command("curl", "-L", libcyaml_bin_url, "-o", "binary.bin")
	err = libcyaml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcyaml_src_url := "https://github.com/tlsa/libcyaml/archive/refs/tags/v1.4.2.src.tar.gz"
	libcyaml_cmd_src := exec.Command("curl", "-L", libcyaml_src_url, "-o", "source.tar.gz")
	err = libcyaml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcyaml_cmd_direct := exec.Command("./binary")
	err = libcyaml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
}
