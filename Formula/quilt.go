package main

// Quilt - Work with series of patches
// Homepage: https://savannah.nongnu.org/projects/quilt

import (
	"fmt"
	
	"os/exec"
)

func installQuilt() {
	// Método 1: Descargar y extraer .tar.gz
	quilt_tar_url := "https://download.savannah.gnu.org/releases/quilt/quilt-0.68.tar.gz"
	quilt_cmd_tar := exec.Command("curl", "-L", quilt_tar_url, "-o", "package.tar.gz")
	err := quilt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	quilt_zip_url := "https://download.savannah.gnu.org/releases/quilt/quilt-0.68.zip"
	quilt_cmd_zip := exec.Command("curl", "-L", quilt_zip_url, "-o", "package.zip")
	err = quilt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	quilt_bin_url := "https://download.savannah.gnu.org/releases/quilt/quilt-0.68.bin"
	quilt_cmd_bin := exec.Command("curl", "-L", quilt_bin_url, "-o", "binary.bin")
	err = quilt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	quilt_src_url := "https://download.savannah.gnu.org/releases/quilt/quilt-0.68.src.tar.gz"
	quilt_cmd_src := exec.Command("curl", "-L", quilt_src_url, "-o", "source.tar.gz")
	err = quilt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	quilt_cmd_direct := exec.Command("./binary")
	err = quilt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: diffutils")
	exec.Command("latte", "install", "diffutils").Run()
	fmt.Println("Instalando dependencia: gpatch")
	exec.Command("latte", "install", "gpatch").Run()
}
