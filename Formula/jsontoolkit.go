package main

// Jsontoolkit - Swiss-army knife library for expressive JSON programming in modern C++
// Homepage: https://jsontoolkit.sourcemeta.com/

import (
	"fmt"
	
	"os/exec"
)

func installJsontoolkit() {
	// Método 1: Descargar y extraer .tar.gz
	jsontoolkit_tar_url := "https://github.com/sourcemeta/jsontoolkit/archive/refs/tags/v2.0.0.tar.gz"
	jsontoolkit_cmd_tar := exec.Command("curl", "-L", jsontoolkit_tar_url, "-o", "package.tar.gz")
	err := jsontoolkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsontoolkit_zip_url := "https://github.com/sourcemeta/jsontoolkit/archive/refs/tags/v2.0.0.zip"
	jsontoolkit_cmd_zip := exec.Command("curl", "-L", jsontoolkit_zip_url, "-o", "package.zip")
	err = jsontoolkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsontoolkit_bin_url := "https://github.com/sourcemeta/jsontoolkit/archive/refs/tags/v2.0.0.bin"
	jsontoolkit_cmd_bin := exec.Command("curl", "-L", jsontoolkit_bin_url, "-o", "binary.bin")
	err = jsontoolkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsontoolkit_src_url := "https://github.com/sourcemeta/jsontoolkit/archive/refs/tags/v2.0.0.src.tar.gz"
	jsontoolkit_cmd_src := exec.Command("curl", "-L", jsontoolkit_src_url, "-o", "source.tar.gz")
	err = jsontoolkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsontoolkit_cmd_direct := exec.Command("./binary")
	err = jsontoolkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
