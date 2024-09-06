package main

// Diffr - LCS based diff highlighting tool to ease code review from your terminal
// Homepage: https://github.com/mookid/diffr

import (
	"fmt"
	
	"os/exec"
)

func installDiffr() {
	// Método 1: Descargar y extraer .tar.gz
	diffr_tar_url := "https://github.com/mookid/diffr/archive/refs/tags/v0.1.5.tar.gz"
	diffr_cmd_tar := exec.Command("curl", "-L", diffr_tar_url, "-o", "package.tar.gz")
	err := diffr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diffr_zip_url := "https://github.com/mookid/diffr/archive/refs/tags/v0.1.5.zip"
	diffr_cmd_zip := exec.Command("curl", "-L", diffr_zip_url, "-o", "package.zip")
	err = diffr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diffr_bin_url := "https://github.com/mookid/diffr/archive/refs/tags/v0.1.5.bin"
	diffr_cmd_bin := exec.Command("curl", "-L", diffr_bin_url, "-o", "binary.bin")
	err = diffr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diffr_src_url := "https://github.com/mookid/diffr/archive/refs/tags/v0.1.5.src.tar.gz"
	diffr_cmd_src := exec.Command("curl", "-L", diffr_src_url, "-o", "source.tar.gz")
	err = diffr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diffr_cmd_direct := exec.Command("./binary")
	err = diffr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: diffutils")
exec.Command("latte", "install", "diffutils")
}
