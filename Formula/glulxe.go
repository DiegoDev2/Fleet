package main

// Glulxe - Portable VM like the Z-machine
// Homepage: https://www.eblong.com/zarf/glulx/

import (
	"fmt"
	
	"os/exec"
)

func installGlulxe() {
	// Método 1: Descargar y extraer .tar.gz
	glulxe_tar_url := "https://eblong.com/zarf/glulx/glulxe-061.tar.gz"
	glulxe_cmd_tar := exec.Command("curl", "-L", glulxe_tar_url, "-o", "package.tar.gz")
	err := glulxe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glulxe_zip_url := "https://eblong.com/zarf/glulx/glulxe-061.zip"
	glulxe_cmd_zip := exec.Command("curl", "-L", glulxe_zip_url, "-o", "package.zip")
	err = glulxe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glulxe_bin_url := "https://eblong.com/zarf/glulx/glulxe-061.bin"
	glulxe_cmd_bin := exec.Command("curl", "-L", glulxe_bin_url, "-o", "binary.bin")
	err = glulxe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glulxe_src_url := "https://eblong.com/zarf/glulx/glulxe-061.src.tar.gz"
	glulxe_cmd_src := exec.Command("curl", "-L", glulxe_src_url, "-o", "source.tar.gz")
	err = glulxe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glulxe_cmd_direct := exec.Command("./binary")
	err = glulxe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: glktermw")
	exec.Command("latte", "install", "glktermw").Run()
}
