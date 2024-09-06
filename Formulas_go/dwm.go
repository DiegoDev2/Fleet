package main

// Dwm - Dynamic window manager
// Homepage: https://dwm.suckless.org/

import (
	"fmt"
	
	"os/exec"
)

func installDwm() {
	// Método 1: Descargar y extraer .tar.gz
	dwm_tar_url := "https://dl.suckless.org/dwm/dwm-6.5.tar.gz"
	dwm_cmd_tar := exec.Command("curl", "-L", dwm_tar_url, "-o", "package.tar.gz")
	err := dwm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dwm_zip_url := "https://dl.suckless.org/dwm/dwm-6.5.zip"
	dwm_cmd_zip := exec.Command("curl", "-L", dwm_zip_url, "-o", "package.zip")
	err = dwm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dwm_bin_url := "https://dl.suckless.org/dwm/dwm-6.5.bin"
	dwm_cmd_bin := exec.Command("curl", "-L", dwm_bin_url, "-o", "binary.bin")
	err = dwm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dwm_src_url := "https://dl.suckless.org/dwm/dwm-6.5.src.tar.gz"
	dwm_cmd_src := exec.Command("curl", "-L", dwm_src_url, "-o", "source.tar.gz")
	err = dwm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dwm_cmd_direct := exec.Command("./binary")
	err = dwm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dmenu")
exec.Command("latte", "install", "dmenu")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxft")
exec.Command("latte", "install", "libxft")
	fmt.Println("Instalando dependencia: libxinerama")
exec.Command("latte", "install", "libxinerama")
}
