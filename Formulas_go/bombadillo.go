package main

// Bombadillo - Non-web browser, designed for a growing list of protocols
// Homepage: https://bombadillo.colorfield.space/

import (
	"fmt"
	
	"os/exec"
)

func installBombadillo() {
	// Método 1: Descargar y extraer .tar.gz
	bombadillo_tar_url := "https://tildegit.org/sloum/bombadillo/archive/2.4.0.tar.gz"
	bombadillo_cmd_tar := exec.Command("curl", "-L", bombadillo_tar_url, "-o", "package.tar.gz")
	err := bombadillo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bombadillo_zip_url := "https://tildegit.org/sloum/bombadillo/archive/2.4.0.zip"
	bombadillo_cmd_zip := exec.Command("curl", "-L", bombadillo_zip_url, "-o", "package.zip")
	err = bombadillo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bombadillo_bin_url := "https://tildegit.org/sloum/bombadillo/archive/2.4.0.bin"
	bombadillo_cmd_bin := exec.Command("curl", "-L", bombadillo_bin_url, "-o", "binary.bin")
	err = bombadillo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bombadillo_src_url := "https://tildegit.org/sloum/bombadillo/archive/2.4.0.src.tar.gz"
	bombadillo_cmd_src := exec.Command("curl", "-L", bombadillo_src_url, "-o", "source.tar.gz")
	err = bombadillo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bombadillo_cmd_direct := exec.Command("./binary")
	err = bombadillo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
