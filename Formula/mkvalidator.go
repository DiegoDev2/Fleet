package main

// Mkvalidator - Tool to verify Matroska and WebM files for spec conformance
// Homepage: https://www.matroska.org/downloads/mkvalidator.html

import (
	"fmt"
	
	"os/exec"
)

func installMkvalidator() {
	// Método 1: Descargar y extraer .tar.gz
	mkvalidator_tar_url := "https://downloads.sourceforge.net/project/matroska/mkvalidator/mkvalidator-0.6.0.tar.bz2"
	mkvalidator_cmd_tar := exec.Command("curl", "-L", mkvalidator_tar_url, "-o", "package.tar.gz")
	err := mkvalidator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mkvalidator_zip_url := "https://downloads.sourceforge.net/project/matroska/mkvalidator/mkvalidator-0.6.0.tar.bz2"
	mkvalidator_cmd_zip := exec.Command("curl", "-L", mkvalidator_zip_url, "-o", "package.zip")
	err = mkvalidator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mkvalidator_bin_url := "https://downloads.sourceforge.net/project/matroska/mkvalidator/mkvalidator-0.6.0.tar.bz2"
	mkvalidator_cmd_bin := exec.Command("curl", "-L", mkvalidator_bin_url, "-o", "binary.bin")
	err = mkvalidator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mkvalidator_src_url := "https://downloads.sourceforge.net/project/matroska/mkvalidator/mkvalidator-0.6.0.tar.bz2"
	mkvalidator_cmd_src := exec.Command("curl", "-L", mkvalidator_src_url, "-o", "source.tar.gz")
	err = mkvalidator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mkvalidator_cmd_direct := exec.Command("./binary")
	err = mkvalidator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
