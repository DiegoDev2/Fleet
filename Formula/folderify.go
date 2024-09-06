package main

// Folderify - Generate pixel-perfect macOS folder icons in the native style
// Homepage: https://github.com/lgarron/folderify

import (
	"fmt"
	
	"os/exec"
)

func installFolderify() {
	// Método 1: Descargar y extraer .tar.gz
	folderify_tar_url := "https://github.com/lgarron/folderify/archive/refs/tags/v4.0.0.tar.gz"
	folderify_cmd_tar := exec.Command("curl", "-L", folderify_tar_url, "-o", "package.tar.gz")
	err := folderify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	folderify_zip_url := "https://github.com/lgarron/folderify/archive/refs/tags/v4.0.0.zip"
	folderify_cmd_zip := exec.Command("curl", "-L", folderify_zip_url, "-o", "package.zip")
	err = folderify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	folderify_bin_url := "https://github.com/lgarron/folderify/archive/refs/tags/v4.0.0.bin"
	folderify_cmd_bin := exec.Command("curl", "-L", folderify_bin_url, "-o", "binary.bin")
	err = folderify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	folderify_src_url := "https://github.com/lgarron/folderify/archive/refs/tags/v4.0.0.src.tar.gz"
	folderify_cmd_src := exec.Command("curl", "-L", folderify_src_url, "-o", "source.tar.gz")
	err = folderify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	folderify_cmd_direct := exec.Command("./binary")
	err = folderify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: imagemagick")
	exec.Command("latte", "install", "imagemagick").Run()
}
