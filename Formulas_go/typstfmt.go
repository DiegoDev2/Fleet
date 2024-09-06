package main

// Typstfmt - Formatter for typst
// Homepage: https://github.com/astrale-sharp/typstfmt

import (
	"fmt"
	
	"os/exec"
)

func installTypstfmt() {
	// Método 1: Descargar y extraer .tar.gz
	typstfmt_tar_url := "https://github.com/astrale-sharp/typstfmt/archive/refs/tags/0.2.10.tar.gz"
	typstfmt_cmd_tar := exec.Command("curl", "-L", typstfmt_tar_url, "-o", "package.tar.gz")
	err := typstfmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	typstfmt_zip_url := "https://github.com/astrale-sharp/typstfmt/archive/refs/tags/0.2.10.zip"
	typstfmt_cmd_zip := exec.Command("curl", "-L", typstfmt_zip_url, "-o", "package.zip")
	err = typstfmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	typstfmt_bin_url := "https://github.com/astrale-sharp/typstfmt/archive/refs/tags/0.2.10.bin"
	typstfmt_cmd_bin := exec.Command("curl", "-L", typstfmt_bin_url, "-o", "binary.bin")
	err = typstfmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	typstfmt_src_url := "https://github.com/astrale-sharp/typstfmt/archive/refs/tags/0.2.10.src.tar.gz"
	typstfmt_cmd_src := exec.Command("curl", "-L", typstfmt_src_url, "-o", "source.tar.gz")
	err = typstfmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	typstfmt_cmd_direct := exec.Command("./binary")
	err = typstfmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
