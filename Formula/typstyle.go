package main

// Typstyle - Beautiful and reliable typst code formatter
// Homepage: https://enter-tainer.github.io/typstyle/

import (
	"fmt"
	
	"os/exec"
)

func installTypstyle() {
	// Método 1: Descargar y extraer .tar.gz
	typstyle_tar_url := "https://github.com/Enter-tainer/typstyle/archive/refs/tags/v0.11.32.tar.gz"
	typstyle_cmd_tar := exec.Command("curl", "-L", typstyle_tar_url, "-o", "package.tar.gz")
	err := typstyle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	typstyle_zip_url := "https://github.com/Enter-tainer/typstyle/archive/refs/tags/v0.11.32.zip"
	typstyle_cmd_zip := exec.Command("curl", "-L", typstyle_zip_url, "-o", "package.zip")
	err = typstyle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	typstyle_bin_url := "https://github.com/Enter-tainer/typstyle/archive/refs/tags/v0.11.32.bin"
	typstyle_cmd_bin := exec.Command("curl", "-L", typstyle_bin_url, "-o", "binary.bin")
	err = typstyle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	typstyle_src_url := "https://github.com/Enter-tainer/typstyle/archive/refs/tags/v0.11.32.src.tar.gz"
	typstyle_cmd_src := exec.Command("curl", "-L", typstyle_src_url, "-o", "source.tar.gz")
	err = typstyle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	typstyle_cmd_direct := exec.Command("./binary")
	err = typstyle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
