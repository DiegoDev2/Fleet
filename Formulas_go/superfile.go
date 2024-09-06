package main

// Superfile - Modern and pretty fancy file manager for the terminal
// Homepage: https://github.com/yorukot/superfile

import (
	"fmt"
	
	"os/exec"
)

func installSuperfile() {
	// Método 1: Descargar y extraer .tar.gz
	superfile_tar_url := "https://github.com/yorukot/superfile/archive/refs/tags/v1.1.4.tar.gz"
	superfile_cmd_tar := exec.Command("curl", "-L", superfile_tar_url, "-o", "package.tar.gz")
	err := superfile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	superfile_zip_url := "https://github.com/yorukot/superfile/archive/refs/tags/v1.1.4.zip"
	superfile_cmd_zip := exec.Command("curl", "-L", superfile_zip_url, "-o", "package.zip")
	err = superfile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	superfile_bin_url := "https://github.com/yorukot/superfile/archive/refs/tags/v1.1.4.bin"
	superfile_cmd_bin := exec.Command("curl", "-L", superfile_bin_url, "-o", "binary.bin")
	err = superfile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	superfile_src_url := "https://github.com/yorukot/superfile/archive/refs/tags/v1.1.4.src.tar.gz"
	superfile_cmd_src := exec.Command("curl", "-L", superfile_src_url, "-o", "source.tar.gz")
	err = superfile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	superfile_cmd_direct := exec.Command("./binary")
	err = superfile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
