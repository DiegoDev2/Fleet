package main

// Vifm - Ncurses-based file manager with vi-like keybindings
// Homepage: https://vifm.info/

import (
	"fmt"
	
	"os/exec"
)

func installVifm() {
	// Método 1: Descargar y extraer .tar.gz
	vifm_tar_url := "https://github.com/vifm/vifm/releases/download/v0.13/vifm-0.13.tar.bz2"
	vifm_cmd_tar := exec.Command("curl", "-L", vifm_tar_url, "-o", "package.tar.gz")
	err := vifm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vifm_zip_url := "https://github.com/vifm/vifm/releases/download/v0.13/vifm-0.13.tar.bz2"
	vifm_cmd_zip := exec.Command("curl", "-L", vifm_zip_url, "-o", "package.zip")
	err = vifm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vifm_bin_url := "https://github.com/vifm/vifm/releases/download/v0.13/vifm-0.13.tar.bz2"
	vifm_cmd_bin := exec.Command("curl", "-L", vifm_bin_url, "-o", "binary.bin")
	err = vifm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vifm_src_url := "https://github.com/vifm/vifm/releases/download/v0.13/vifm-0.13.tar.bz2"
	vifm_cmd_src := exec.Command("curl", "-L", vifm_src_url, "-o", "source.tar.gz")
	err = vifm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vifm_cmd_direct := exec.Command("./binary")
	err = vifm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
}
