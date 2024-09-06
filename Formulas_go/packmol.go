package main

// Packmol - Packing optimization for molecular dynamics simulations
// Homepage: https://www.ime.unicamp.br/~martinez/packmol/

import (
	"fmt"
	
	"os/exec"
)

func installPackmol() {
	// Método 1: Descargar y extraer .tar.gz
	packmol_tar_url := "https://github.com/m3g/packmol/archive/refs/tags/v20.15.1.tar.gz"
	packmol_cmd_tar := exec.Command("curl", "-L", packmol_tar_url, "-o", "package.tar.gz")
	err := packmol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	packmol_zip_url := "https://github.com/m3g/packmol/archive/refs/tags/v20.15.1.zip"
	packmol_cmd_zip := exec.Command("curl", "-L", packmol_zip_url, "-o", "package.zip")
	err = packmol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	packmol_bin_url := "https://github.com/m3g/packmol/archive/refs/tags/v20.15.1.bin"
	packmol_cmd_bin := exec.Command("curl", "-L", packmol_bin_url, "-o", "binary.bin")
	err = packmol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	packmol_src_url := "https://github.com/m3g/packmol/archive/refs/tags/v20.15.1.src.tar.gz"
	packmol_cmd_src := exec.Command("curl", "-L", packmol_src_url, "-o", "source.tar.gz")
	err = packmol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	packmol_cmd_direct := exec.Command("./binary")
	err = packmol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
