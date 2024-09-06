package main

// Frei0r - Minimalistic plugin API for video effects
// Homepage: https://frei0r.dyne.org/

import (
	"fmt"
	
	"os/exec"
)

func installFrei0r() {
	// Método 1: Descargar y extraer .tar.gz
	frei0r_tar_url := "https://github.com/dyne/frei0r/archive/refs/tags/v2.3.3.tar.gz"
	frei0r_cmd_tar := exec.Command("curl", "-L", frei0r_tar_url, "-o", "package.tar.gz")
	err := frei0r_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	frei0r_zip_url := "https://github.com/dyne/frei0r/archive/refs/tags/v2.3.3.zip"
	frei0r_cmd_zip := exec.Command("curl", "-L", frei0r_zip_url, "-o", "package.zip")
	err = frei0r_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	frei0r_bin_url := "https://github.com/dyne/frei0r/archive/refs/tags/v2.3.3.bin"
	frei0r_cmd_bin := exec.Command("curl", "-L", frei0r_bin_url, "-o", "binary.bin")
	err = frei0r_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	frei0r_src_url := "https://github.com/dyne/frei0r/archive/refs/tags/v2.3.3.src.tar.gz"
	frei0r_cmd_src := exec.Command("curl", "-L", frei0r_src_url, "-o", "source.tar.gz")
	err = frei0r_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	frei0r_cmd_direct := exec.Command("./binary")
	err = frei0r_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
