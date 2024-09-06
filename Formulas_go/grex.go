package main

// Grex - Command-line tool for generating regular expressions
// Homepage: https://github.com/pemistahl/grex

import (
	"fmt"
	
	"os/exec"
)

func installGrex() {
	// Método 1: Descargar y extraer .tar.gz
	grex_tar_url := "https://github.com/pemistahl/grex/archive/refs/tags/v1.4.5.tar.gz"
	grex_cmd_tar := exec.Command("curl", "-L", grex_tar_url, "-o", "package.tar.gz")
	err := grex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grex_zip_url := "https://github.com/pemistahl/grex/archive/refs/tags/v1.4.5.zip"
	grex_cmd_zip := exec.Command("curl", "-L", grex_zip_url, "-o", "package.zip")
	err = grex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grex_bin_url := "https://github.com/pemistahl/grex/archive/refs/tags/v1.4.5.bin"
	grex_cmd_bin := exec.Command("curl", "-L", grex_bin_url, "-o", "binary.bin")
	err = grex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grex_src_url := "https://github.com/pemistahl/grex/archive/refs/tags/v1.4.5.src.tar.gz"
	grex_cmd_src := exec.Command("curl", "-L", grex_src_url, "-o", "source.tar.gz")
	err = grex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grex_cmd_direct := exec.Command("./binary")
	err = grex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
