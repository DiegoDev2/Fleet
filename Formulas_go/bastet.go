package main

// Bastet - Bastard Tetris
// Homepage: https://fph.altervista.org/prog/bastet.html

import (
	"fmt"
	
	"os/exec"
)

func installBastet() {
	// Método 1: Descargar y extraer .tar.gz
	bastet_tar_url := "https://github.com/fph/bastet/archive/refs/tags/0.43.2.tar.gz"
	bastet_cmd_tar := exec.Command("curl", "-L", bastet_tar_url, "-o", "package.tar.gz")
	err := bastet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bastet_zip_url := "https://github.com/fph/bastet/archive/refs/tags/0.43.2.zip"
	bastet_cmd_zip := exec.Command("curl", "-L", bastet_zip_url, "-o", "package.zip")
	err = bastet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bastet_bin_url := "https://github.com/fph/bastet/archive/refs/tags/0.43.2.bin"
	bastet_cmd_bin := exec.Command("curl", "-L", bastet_bin_url, "-o", "binary.bin")
	err = bastet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bastet_src_url := "https://github.com/fph/bastet/archive/refs/tags/0.43.2.src.tar.gz"
	bastet_cmd_src := exec.Command("curl", "-L", bastet_src_url, "-o", "source.tar.gz")
	err = bastet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bastet_cmd_direct := exec.Command("./binary")
	err = bastet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
}
