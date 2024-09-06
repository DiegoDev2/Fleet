package main

// Kimwituxx - Tool for processing trees (i.e. terms)
// Homepage: https://savannah.nongnu.org/projects/kimwitu-pp

import (
	"fmt"
	
	"os/exec"
)

func installKimwituxx() {
	// Método 1: Descargar y extraer .tar.gz
	kimwituxx_tar_url := "https://download.savannah.gnu.org/releases/kimwitu-pp/kimwitu++-2.3.13.tar.gz"
	kimwituxx_cmd_tar := exec.Command("curl", "-L", kimwituxx_tar_url, "-o", "package.tar.gz")
	err := kimwituxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kimwituxx_zip_url := "https://download.savannah.gnu.org/releases/kimwitu-pp/kimwitu++-2.3.13.zip"
	kimwituxx_cmd_zip := exec.Command("curl", "-L", kimwituxx_zip_url, "-o", "package.zip")
	err = kimwituxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kimwituxx_bin_url := "https://download.savannah.gnu.org/releases/kimwitu-pp/kimwitu++-2.3.13.bin"
	kimwituxx_cmd_bin := exec.Command("curl", "-L", kimwituxx_bin_url, "-o", "binary.bin")
	err = kimwituxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kimwituxx_src_url := "https://download.savannah.gnu.org/releases/kimwitu-pp/kimwitu++-2.3.13.src.tar.gz"
	kimwituxx_cmd_src := exec.Command("curl", "-L", kimwituxx_src_url, "-o", "source.tar.gz")
	err = kimwituxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kimwituxx_cmd_direct := exec.Command("./binary")
	err = kimwituxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
