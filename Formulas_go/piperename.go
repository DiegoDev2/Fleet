package main

// PipeRename - Rename your files using your favorite text editor
// Homepage: https://github.com/marcusbuffett/pipe-rename

import (
	"fmt"
	
	"os/exec"
)

func installPipeRename() {
	// Método 1: Descargar y extraer .tar.gz
	piperename_tar_url := "https://github.com/marcusbuffett/pipe-rename/archive/refs/tags/1.6.5.tar.gz"
	piperename_cmd_tar := exec.Command("curl", "-L", piperename_tar_url, "-o", "package.tar.gz")
	err := piperename_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	piperename_zip_url := "https://github.com/marcusbuffett/pipe-rename/archive/refs/tags/1.6.5.zip"
	piperename_cmd_zip := exec.Command("curl", "-L", piperename_zip_url, "-o", "package.zip")
	err = piperename_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	piperename_bin_url := "https://github.com/marcusbuffett/pipe-rename/archive/refs/tags/1.6.5.bin"
	piperename_cmd_bin := exec.Command("curl", "-L", piperename_bin_url, "-o", "binary.bin")
	err = piperename_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	piperename_src_url := "https://github.com/marcusbuffett/pipe-rename/archive/refs/tags/1.6.5.src.tar.gz"
	piperename_cmd_src := exec.Command("curl", "-L", piperename_src_url, "-o", "source.tar.gz")
	err = piperename_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	piperename_cmd_direct := exec.Command("./binary")
	err = piperename_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
