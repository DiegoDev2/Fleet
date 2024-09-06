package main

// Libyaml - YAML Parser
// Homepage: https://github.com/yaml/libyaml

import (
	"fmt"
	
	"os/exec"
)

func installLibyaml() {
	// Método 1: Descargar y extraer .tar.gz
	libyaml_tar_url := "https://github.com/yaml/libyaml/archive/refs/tags/0.2.5.tar.gz"
	libyaml_cmd_tar := exec.Command("curl", "-L", libyaml_tar_url, "-o", "package.tar.gz")
	err := libyaml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libyaml_zip_url := "https://github.com/yaml/libyaml/archive/refs/tags/0.2.5.zip"
	libyaml_cmd_zip := exec.Command("curl", "-L", libyaml_zip_url, "-o", "package.zip")
	err = libyaml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libyaml_bin_url := "https://github.com/yaml/libyaml/archive/refs/tags/0.2.5.bin"
	libyaml_cmd_bin := exec.Command("curl", "-L", libyaml_bin_url, "-o", "binary.bin")
	err = libyaml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libyaml_src_url := "https://github.com/yaml/libyaml/archive/refs/tags/0.2.5.src.tar.gz"
	libyaml_cmd_src := exec.Command("curl", "-L", libyaml_src_url, "-o", "source.tar.gz")
	err = libyaml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libyaml_cmd_direct := exec.Command("./binary")
	err = libyaml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
