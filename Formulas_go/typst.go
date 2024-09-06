package main

// Typst - Markup-based typesetting system
// Homepage: https://typst.app/

import (
	"fmt"
	
	"os/exec"
)

func installTypst() {
	// Método 1: Descargar y extraer .tar.gz
	typst_tar_url := "https://github.com/typst/typst/archive/refs/tags/v0.11.1.tar.gz"
	typst_cmd_tar := exec.Command("curl", "-L", typst_tar_url, "-o", "package.tar.gz")
	err := typst_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	typst_zip_url := "https://github.com/typst/typst/archive/refs/tags/v0.11.1.zip"
	typst_cmd_zip := exec.Command("curl", "-L", typst_zip_url, "-o", "package.zip")
	err = typst_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	typst_bin_url := "https://github.com/typst/typst/archive/refs/tags/v0.11.1.bin"
	typst_cmd_bin := exec.Command("curl", "-L", typst_bin_url, "-o", "binary.bin")
	err = typst_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	typst_src_url := "https://github.com/typst/typst/archive/refs/tags/v0.11.1.src.tar.gz"
	typst_cmd_src := exec.Command("curl", "-L", typst_src_url, "-o", "source.tar.gz")
	err = typst_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	typst_cmd_direct := exec.Command("./binary")
	err = typst_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
