package main

// Rename - Perl-powered file rename script with many helpful built-ins
// Homepage: http://plasmasturm.org/code/rename

import (
	"fmt"
	
	"os/exec"
)

func installRename() {
	// Método 1: Descargar y extraer .tar.gz
	rename_tar_url := "https://github.com/ap/rename/archive/refs/tags/v1.601.tar.gz"
	rename_cmd_tar := exec.Command("curl", "-L", rename_tar_url, "-o", "package.tar.gz")
	err := rename_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rename_zip_url := "https://github.com/ap/rename/archive/refs/tags/v1.601.zip"
	rename_cmd_zip := exec.Command("curl", "-L", rename_zip_url, "-o", "package.zip")
	err = rename_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rename_bin_url := "https://github.com/ap/rename/archive/refs/tags/v1.601.bin"
	rename_cmd_bin := exec.Command("curl", "-L", rename_bin_url, "-o", "binary.bin")
	err = rename_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rename_src_url := "https://github.com/ap/rename/archive/refs/tags/v1.601.src.tar.gz"
	rename_cmd_src := exec.Command("curl", "-L", rename_src_url, "-o", "source.tar.gz")
	err = rename_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rename_cmd_direct := exec.Command("./binary")
	err = rename_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pod2man")
	exec.Command("latte", "install", "pod2man").Run()
}
