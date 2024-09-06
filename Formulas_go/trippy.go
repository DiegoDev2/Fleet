package main

// Trippy - Network diagnostic tool, inspired by mtr
// Homepage: https://trippy.cli.rs/

import (
	"fmt"
	
	"os/exec"
)

func installTrippy() {
	// Método 1: Descargar y extraer .tar.gz
	trippy_tar_url := "https://github.com/fujiapple852/trippy/archive/refs/tags/0.11.0.tar.gz"
	trippy_cmd_tar := exec.Command("curl", "-L", trippy_tar_url, "-o", "package.tar.gz")
	err := trippy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trippy_zip_url := "https://github.com/fujiapple852/trippy/archive/refs/tags/0.11.0.zip"
	trippy_cmd_zip := exec.Command("curl", "-L", trippy_zip_url, "-o", "package.zip")
	err = trippy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trippy_bin_url := "https://github.com/fujiapple852/trippy/archive/refs/tags/0.11.0.bin"
	trippy_cmd_bin := exec.Command("curl", "-L", trippy_bin_url, "-o", "binary.bin")
	err = trippy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trippy_src_url := "https://github.com/fujiapple852/trippy/archive/refs/tags/0.11.0.src.tar.gz"
	trippy_cmd_src := exec.Command("curl", "-L", trippy_src_url, "-o", "source.tar.gz")
	err = trippy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trippy_cmd_direct := exec.Command("./binary")
	err = trippy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
