package main

// Utf8cpp - UTF-8 with C++ in a Portable Way
// Homepage: https://github.com/nemtrif/utfcpp

import (
	"fmt"
	
	"os/exec"
)

func installUtf8cpp() {
	// Método 1: Descargar y extraer .tar.gz
	utf8cpp_tar_url := "https://github.com/nemtrif/utfcpp/archive/refs/tags/v4.0.5.tar.gz"
	utf8cpp_cmd_tar := exec.Command("curl", "-L", utf8cpp_tar_url, "-o", "package.tar.gz")
	err := utf8cpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	utf8cpp_zip_url := "https://github.com/nemtrif/utfcpp/archive/refs/tags/v4.0.5.zip"
	utf8cpp_cmd_zip := exec.Command("curl", "-L", utf8cpp_zip_url, "-o", "package.zip")
	err = utf8cpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	utf8cpp_bin_url := "https://github.com/nemtrif/utfcpp/archive/refs/tags/v4.0.5.bin"
	utf8cpp_cmd_bin := exec.Command("curl", "-L", utf8cpp_bin_url, "-o", "binary.bin")
	err = utf8cpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	utf8cpp_src_url := "https://github.com/nemtrif/utfcpp/archive/refs/tags/v4.0.5.src.tar.gz"
	utf8cpp_cmd_src := exec.Command("curl", "-L", utf8cpp_src_url, "-o", "source.tar.gz")
	err = utf8cpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	utf8cpp_cmd_direct := exec.Command("./binary")
	err = utf8cpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
