package main

// ClearlooksPhenix - GTK+3 port of the Clearlooks Theme
// Homepage: https://github.com/jpfleury/clearlooks-phenix

import (
	"fmt"
	
	"os/exec"
)

func installClearlooksPhenix() {
	// Método 1: Descargar y extraer .tar.gz
	clearlooksphenix_tar_url := "https://github.com/jpfleury/clearlooks-phenix/archive/refs/tags/7.0.1.tar.gz"
	clearlooksphenix_cmd_tar := exec.Command("curl", "-L", clearlooksphenix_tar_url, "-o", "package.tar.gz")
	err := clearlooksphenix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clearlooksphenix_zip_url := "https://github.com/jpfleury/clearlooks-phenix/archive/refs/tags/7.0.1.zip"
	clearlooksphenix_cmd_zip := exec.Command("curl", "-L", clearlooksphenix_zip_url, "-o", "package.zip")
	err = clearlooksphenix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clearlooksphenix_bin_url := "https://github.com/jpfleury/clearlooks-phenix/archive/refs/tags/7.0.1.bin"
	clearlooksphenix_cmd_bin := exec.Command("curl", "-L", clearlooksphenix_bin_url, "-o", "binary.bin")
	err = clearlooksphenix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clearlooksphenix_src_url := "https://github.com/jpfleury/clearlooks-phenix/archive/refs/tags/7.0.1.src.tar.gz"
	clearlooksphenix_cmd_src := exec.Command("curl", "-L", clearlooksphenix_src_url, "-o", "source.tar.gz")
	err = clearlooksphenix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clearlooksphenix_cmd_direct := exec.Command("./binary")
	err = clearlooksphenix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
}
