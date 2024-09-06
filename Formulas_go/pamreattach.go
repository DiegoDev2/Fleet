package main

// PamReattach - PAM module for reattaching to the user's GUI (Aqua) session
// Homepage: https://github.com/fabianishere/pam_reattach

import (
	"fmt"
	
	"os/exec"
)

func installPamReattach() {
	// Método 1: Descargar y extraer .tar.gz
	pamreattach_tar_url := "https://github.com/fabianishere/pam_reattach/archive/refs/tags/v1.3.tar.gz"
	pamreattach_cmd_tar := exec.Command("curl", "-L", pamreattach_tar_url, "-o", "package.tar.gz")
	err := pamreattach_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pamreattach_zip_url := "https://github.com/fabianishere/pam_reattach/archive/refs/tags/v1.3.zip"
	pamreattach_cmd_zip := exec.Command("curl", "-L", pamreattach_zip_url, "-o", "package.zip")
	err = pamreattach_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pamreattach_bin_url := "https://github.com/fabianishere/pam_reattach/archive/refs/tags/v1.3.bin"
	pamreattach_cmd_bin := exec.Command("curl", "-L", pamreattach_bin_url, "-o", "binary.bin")
	err = pamreattach_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pamreattach_src_url := "https://github.com/fabianishere/pam_reattach/archive/refs/tags/v1.3.src.tar.gz"
	pamreattach_cmd_src := exec.Command("curl", "-L", pamreattach_src_url, "-o", "source.tar.gz")
	err = pamreattach_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pamreattach_cmd_direct := exec.Command("./binary")
	err = pamreattach_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
