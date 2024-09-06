package main

// Tinymist - Language server for Typst
// Homepage: https://github.com/Myriad-Dreamin/tinymist

import (
	"fmt"
	
	"os/exec"
)

func installTinymist() {
	// Método 1: Descargar y extraer .tar.gz
	tinymist_tar_url := "https://github.com/Myriad-Dreamin/tinymist/archive/refs/tags/v0.11.20.tar.gz"
	tinymist_cmd_tar := exec.Command("curl", "-L", tinymist_tar_url, "-o", "package.tar.gz")
	err := tinymist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tinymist_zip_url := "https://github.com/Myriad-Dreamin/tinymist/archive/refs/tags/v0.11.20.zip"
	tinymist_cmd_zip := exec.Command("curl", "-L", tinymist_zip_url, "-o", "package.zip")
	err = tinymist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tinymist_bin_url := "https://github.com/Myriad-Dreamin/tinymist/archive/refs/tags/v0.11.20.bin"
	tinymist_cmd_bin := exec.Command("curl", "-L", tinymist_bin_url, "-o", "binary.bin")
	err = tinymist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tinymist_src_url := "https://github.com/Myriad-Dreamin/tinymist/archive/refs/tags/v0.11.20.src.tar.gz"
	tinymist_cmd_src := exec.Command("curl", "-L", tinymist_src_url, "-o", "source.tar.gz")
	err = tinymist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tinymist_cmd_direct := exec.Command("./binary")
	err = tinymist_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
