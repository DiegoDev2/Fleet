package main

// Twm - Tab Window Manager for X Window System
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installTwm() {
	// Método 1: Descargar y extraer .tar.gz
	twm_tar_url := "https://www.x.org/releases/individual/app/twm-1.0.12.tar.xz"
	twm_cmd_tar := exec.Command("curl", "-L", twm_tar_url, "-o", "package.tar.gz")
	err := twm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	twm_zip_url := "https://www.x.org/releases/individual/app/twm-1.0.12.tar.xz"
	twm_cmd_zip := exec.Command("curl", "-L", twm_zip_url, "-o", "package.zip")
	err = twm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	twm_bin_url := "https://www.x.org/releases/individual/app/twm-1.0.12.tar.xz"
	twm_cmd_bin := exec.Command("curl", "-L", twm_bin_url, "-o", "binary.bin")
	err = twm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	twm_src_url := "https://www.x.org/releases/individual/app/twm-1.0.12.tar.xz"
	twm_cmd_src := exec.Command("curl", "-L", twm_src_url, "-o", "source.tar.gz")
	err = twm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	twm_cmd_direct := exec.Command("./binary")
	err = twm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libxmu")
	exec.Command("latte", "install", "libxmu").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
}
