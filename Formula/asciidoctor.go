package main

// Asciidoctor - Text processor and publishing toolchain for AsciiDoc
// Homepage: https://asciidoctor.org/

import (
	"fmt"
	
	"os/exec"
)

func installAsciidoctor() {
	// Método 1: Descargar y extraer .tar.gz
	asciidoctor_tar_url := "https://github.com/asciidoctor/asciidoctor/archive/refs/tags/v2.0.23.tar.gz"
	asciidoctor_cmd_tar := exec.Command("curl", "-L", asciidoctor_tar_url, "-o", "package.tar.gz")
	err := asciidoctor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asciidoctor_zip_url := "https://github.com/asciidoctor/asciidoctor/archive/refs/tags/v2.0.23.zip"
	asciidoctor_cmd_zip := exec.Command("curl", "-L", asciidoctor_zip_url, "-o", "package.zip")
	err = asciidoctor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asciidoctor_bin_url := "https://github.com/asciidoctor/asciidoctor/archive/refs/tags/v2.0.23.bin"
	asciidoctor_cmd_bin := exec.Command("curl", "-L", asciidoctor_bin_url, "-o", "binary.bin")
	err = asciidoctor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asciidoctor_src_url := "https://github.com/asciidoctor/asciidoctor/archive/refs/tags/v2.0.23.src.tar.gz"
	asciidoctor_cmd_src := exec.Command("curl", "-L", asciidoctor_src_url, "-o", "source.tar.gz")
	err = asciidoctor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asciidoctor_cmd_direct := exec.Command("./binary")
	err = asciidoctor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
}
