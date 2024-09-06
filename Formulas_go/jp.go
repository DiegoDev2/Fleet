package main

// Jp - Dead simple terminal plots from JSON data
// Homepage: https://github.com/sgreben/jp

import (
	"fmt"
	
	"os/exec"
)

func installJp() {
	// Método 1: Descargar y extraer .tar.gz
	jp_tar_url := "https://github.com/sgreben/jp/archive/refs/tags/1.1.12.tar.gz"
	jp_cmd_tar := exec.Command("curl", "-L", jp_tar_url, "-o", "package.tar.gz")
	err := jp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jp_zip_url := "https://github.com/sgreben/jp/archive/refs/tags/1.1.12.zip"
	jp_cmd_zip := exec.Command("curl", "-L", jp_zip_url, "-o", "package.zip")
	err = jp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jp_bin_url := "https://github.com/sgreben/jp/archive/refs/tags/1.1.12.bin"
	jp_cmd_bin := exec.Command("curl", "-L", jp_bin_url, "-o", "binary.bin")
	err = jp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jp_src_url := "https://github.com/sgreben/jp/archive/refs/tags/1.1.12.src.tar.gz"
	jp_cmd_src := exec.Command("curl", "-L", jp_src_url, "-o", "source.tar.gz")
	err = jp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jp_cmd_direct := exec.Command("./binary")
	err = jp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
