package main

// AdaUrl - WHATWG-compliant and fast URL parser written in modern C++
// Homepage: https://github.com/ada-url/ada

import (
	"fmt"
	
	"os/exec"
)

func installAdaUrl() {
	// Método 1: Descargar y extraer .tar.gz
	adaurl_tar_url := "https://github.com/ada-url/ada/archive/refs/tags/v2.9.2.tar.gz"
	adaurl_cmd_tar := exec.Command("curl", "-L", adaurl_tar_url, "-o", "package.tar.gz")
	err := adaurl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	adaurl_zip_url := "https://github.com/ada-url/ada/archive/refs/tags/v2.9.2.zip"
	adaurl_cmd_zip := exec.Command("curl", "-L", adaurl_zip_url, "-o", "package.zip")
	err = adaurl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	adaurl_bin_url := "https://github.com/ada-url/ada/archive/refs/tags/v2.9.2.bin"
	adaurl_cmd_bin := exec.Command("curl", "-L", adaurl_bin_url, "-o", "binary.bin")
	err = adaurl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	adaurl_src_url := "https://github.com/ada-url/ada/archive/refs/tags/v2.9.2.src.tar.gz"
	adaurl_cmd_src := exec.Command("curl", "-L", adaurl_src_url, "-o", "source.tar.gz")
	err = adaurl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	adaurl_cmd_direct := exec.Command("./binary")
	err = adaurl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
