package main

// Jsoncpp - Library for interacting with JSON
// Homepage: https://github.com/open-source-parsers/jsoncpp

import (
	"fmt"
	
	"os/exec"
)

func installJsoncpp() {
	// Método 1: Descargar y extraer .tar.gz
	jsoncpp_tar_url := "https://github.com/open-source-parsers/jsoncpp/archive/refs/tags/1.9.5.tar.gz"
	jsoncpp_cmd_tar := exec.Command("curl", "-L", jsoncpp_tar_url, "-o", "package.tar.gz")
	err := jsoncpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsoncpp_zip_url := "https://github.com/open-source-parsers/jsoncpp/archive/refs/tags/1.9.5.zip"
	jsoncpp_cmd_zip := exec.Command("curl", "-L", jsoncpp_zip_url, "-o", "package.zip")
	err = jsoncpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsoncpp_bin_url := "https://github.com/open-source-parsers/jsoncpp/archive/refs/tags/1.9.5.bin"
	jsoncpp_cmd_bin := exec.Command("curl", "-L", jsoncpp_bin_url, "-o", "binary.bin")
	err = jsoncpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsoncpp_src_url := "https://github.com/open-source-parsers/jsoncpp/archive/refs/tags/1.9.5.src.tar.gz"
	jsoncpp_cmd_src := exec.Command("curl", "-L", jsoncpp_src_url, "-o", "source.tar.gz")
	err = jsoncpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsoncpp_cmd_direct := exec.Command("./binary")
	err = jsoncpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
