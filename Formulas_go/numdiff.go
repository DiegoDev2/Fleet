package main

// Numdiff - Putative files comparison tool
// Homepage: https://www.nongnu.org/numdiff

import (
	"fmt"
	
	"os/exec"
)

func installNumdiff() {
	// Método 1: Descargar y extraer .tar.gz
	numdiff_tar_url := "https://download.savannah.gnu.org/releases/numdiff/numdiff-5.9.0.tar.gz"
	numdiff_cmd_tar := exec.Command("curl", "-L", numdiff_tar_url, "-o", "package.tar.gz")
	err := numdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	numdiff_zip_url := "https://download.savannah.gnu.org/releases/numdiff/numdiff-5.9.0.zip"
	numdiff_cmd_zip := exec.Command("curl", "-L", numdiff_zip_url, "-o", "package.zip")
	err = numdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	numdiff_bin_url := "https://download.savannah.gnu.org/releases/numdiff/numdiff-5.9.0.bin"
	numdiff_cmd_bin := exec.Command("curl", "-L", numdiff_bin_url, "-o", "binary.bin")
	err = numdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	numdiff_src_url := "https://download.savannah.gnu.org/releases/numdiff/numdiff-5.9.0.src.tar.gz"
	numdiff_cmd_src := exec.Command("curl", "-L", numdiff_src_url, "-o", "source.tar.gz")
	err = numdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	numdiff_cmd_direct := exec.Command("./binary")
	err = numdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
