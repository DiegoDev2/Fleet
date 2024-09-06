package main

// YamlCpp - C++ YAML parser and emitter for YAML 1.2 spec
// Homepage: https://github.com/jbeder/yaml-cpp

import (
	"fmt"
	
	"os/exec"
)

func installYamlCpp() {
	// Método 1: Descargar y extraer .tar.gz
	yamlcpp_tar_url := "https://github.com/jbeder/yaml-cpp/archive/refs/tags/0.8.0.tar.gz"
	yamlcpp_cmd_tar := exec.Command("curl", "-L", yamlcpp_tar_url, "-o", "package.tar.gz")
	err := yamlcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yamlcpp_zip_url := "https://github.com/jbeder/yaml-cpp/archive/refs/tags/0.8.0.zip"
	yamlcpp_cmd_zip := exec.Command("curl", "-L", yamlcpp_zip_url, "-o", "package.zip")
	err = yamlcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yamlcpp_bin_url := "https://github.com/jbeder/yaml-cpp/archive/refs/tags/0.8.0.bin"
	yamlcpp_cmd_bin := exec.Command("curl", "-L", yamlcpp_bin_url, "-o", "binary.bin")
	err = yamlcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yamlcpp_src_url := "https://github.com/jbeder/yaml-cpp/archive/refs/tags/0.8.0.src.tar.gz"
	yamlcpp_cmd_src := exec.Command("curl", "-L", yamlcpp_src_url, "-o", "source.tar.gz")
	err = yamlcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yamlcpp_cmd_direct := exec.Command("./binary")
	err = yamlcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
