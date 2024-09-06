package main

// Lgeneral - Turn-based strategy engine heavily inspired by Panzer General
// Homepage: https://lgames.sourceforge.io/LGeneral/

import (
	"fmt"
	
	"os/exec"
)

func installLgeneral() {
	// Método 1: Descargar y extraer .tar.gz
	lgeneral_tar_url := "https://downloads.sourceforge.net/lgeneral/lgeneral/lgeneral-1.4.4.tar.gz"
	lgeneral_cmd_tar := exec.Command("curl", "-L", lgeneral_tar_url, "-o", "package.tar.gz")
	err := lgeneral_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lgeneral_zip_url := "https://downloads.sourceforge.net/lgeneral/lgeneral/lgeneral-1.4.4.zip"
	lgeneral_cmd_zip := exec.Command("curl", "-L", lgeneral_zip_url, "-o", "package.zip")
	err = lgeneral_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lgeneral_bin_url := "https://downloads.sourceforge.net/lgeneral/lgeneral/lgeneral-1.4.4.bin"
	lgeneral_cmd_bin := exec.Command("curl", "-L", lgeneral_bin_url, "-o", "binary.bin")
	err = lgeneral_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lgeneral_src_url := "https://downloads.sourceforge.net/lgeneral/lgeneral/lgeneral-1.4.4.src.tar.gz"
	lgeneral_cmd_src := exec.Command("curl", "-L", lgeneral_src_url, "-o", "source.tar.gz")
	err = lgeneral_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lgeneral_cmd_direct := exec.Command("./binary")
	err = lgeneral_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_mixer")
	exec.Command("latte", "install", "sdl2_mixer").Run()
}
