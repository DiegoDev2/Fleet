package main

// Navi - Interactive cheatsheet tool for the command-line
// Homepage: https://github.com/denisidoro/navi

import (
	"fmt"
	
	"os/exec"
)

func installNavi() {
	// Método 1: Descargar y extraer .tar.gz
	navi_tar_url := "https://github.com/denisidoro/navi/archive/refs/tags/v2.23.0.tar.gz"
	navi_cmd_tar := exec.Command("curl", "-L", navi_tar_url, "-o", "package.tar.gz")
	err := navi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	navi_zip_url := "https://github.com/denisidoro/navi/archive/refs/tags/v2.23.0.zip"
	navi_cmd_zip := exec.Command("curl", "-L", navi_zip_url, "-o", "package.zip")
	err = navi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	navi_bin_url := "https://github.com/denisidoro/navi/archive/refs/tags/v2.23.0.bin"
	navi_cmd_bin := exec.Command("curl", "-L", navi_bin_url, "-o", "binary.bin")
	err = navi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	navi_src_url := "https://github.com/denisidoro/navi/archive/refs/tags/v2.23.0.src.tar.gz"
	navi_cmd_src := exec.Command("curl", "-L", navi_src_url, "-o", "source.tar.gz")
	err = navi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	navi_cmd_direct := exec.Command("./binary")
	err = navi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: fzf")
exec.Command("latte", "install", "fzf")
}
