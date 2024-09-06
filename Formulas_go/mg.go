package main

// Mg - Small Emacs-like editor
// Homepage: https://github.com/ibara/mg

import (
	"fmt"
	
	"os/exec"
)

func installMg() {
	// Método 1: Descargar y extraer .tar.gz
	mg_tar_url := "https://github.com/ibara/mg/releases/download/mg-7.3/mg-7.3.tar.gz"
	mg_cmd_tar := exec.Command("curl", "-L", mg_tar_url, "-o", "package.tar.gz")
	err := mg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mg_zip_url := "https://github.com/ibara/mg/releases/download/mg-7.3/mg-7.3.zip"
	mg_cmd_zip := exec.Command("curl", "-L", mg_zip_url, "-o", "package.zip")
	err = mg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mg_bin_url := "https://github.com/ibara/mg/releases/download/mg-7.3/mg-7.3.bin"
	mg_cmd_bin := exec.Command("curl", "-L", mg_bin_url, "-o", "binary.bin")
	err = mg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mg_src_url := "https://github.com/ibara/mg/releases/download/mg-7.3/mg-7.3.src.tar.gz"
	mg_cmd_src := exec.Command("curl", "-L", mg_src_url, "-o", "source.tar.gz")
	err = mg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mg_cmd_direct := exec.Command("./binary")
	err = mg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
