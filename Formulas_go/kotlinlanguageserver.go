package main

// KotlinLanguageServer - Intelligent Kotlin support for any editor/IDE using the Language Server Protocol
// Homepage: https://github.com/fwcd/kotlin-language-server

import (
	"fmt"
	
	"os/exec"
)

func installKotlinLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	kotlinlanguageserver_tar_url := "https://github.com/fwcd/kotlin-language-server/archive/refs/tags/1.3.12.tar.gz"
	kotlinlanguageserver_cmd_tar := exec.Command("curl", "-L", kotlinlanguageserver_tar_url, "-o", "package.tar.gz")
	err := kotlinlanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kotlinlanguageserver_zip_url := "https://github.com/fwcd/kotlin-language-server/archive/refs/tags/1.3.12.zip"
	kotlinlanguageserver_cmd_zip := exec.Command("curl", "-L", kotlinlanguageserver_zip_url, "-o", "package.zip")
	err = kotlinlanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kotlinlanguageserver_bin_url := "https://github.com/fwcd/kotlin-language-server/archive/refs/tags/1.3.12.bin"
	kotlinlanguageserver_cmd_bin := exec.Command("curl", "-L", kotlinlanguageserver_bin_url, "-o", "binary.bin")
	err = kotlinlanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kotlinlanguageserver_src_url := "https://github.com/fwcd/kotlin-language-server/archive/refs/tags/1.3.12.src.tar.gz"
	kotlinlanguageserver_cmd_src := exec.Command("curl", "-L", kotlinlanguageserver_src_url, "-o", "source.tar.gz")
	err = kotlinlanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kotlinlanguageserver_cmd_direct := exec.Command("./binary")
	err = kotlinlanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gradle")
exec.Command("latte", "install", "gradle")
	fmt.Println("Instalando dependencia: openjdk@21")
exec.Command("latte", "install", "openjdk@21")
}
