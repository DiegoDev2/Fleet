package main

// Ctl - Programming language for digital color management
// Homepage: https://github.com/ampas/CTL

import (
	"fmt"
	
	"os/exec"
)

func installCtl() {
	// Método 1: Descargar y extraer .tar.gz
	ctl_tar_url := "https://github.com/ampas/CTL/archive/refs/tags/ctl-1.5.3.tar.gz"
	ctl_cmd_tar := exec.Command("curl", "-L", ctl_tar_url, "-o", "package.tar.gz")
	err := ctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ctl_zip_url := "https://github.com/ampas/CTL/archive/refs/tags/ctl-1.5.3.zip"
	ctl_cmd_zip := exec.Command("curl", "-L", ctl_zip_url, "-o", "package.zip")
	err = ctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ctl_bin_url := "https://github.com/ampas/CTL/archive/refs/tags/ctl-1.5.3.bin"
	ctl_cmd_bin := exec.Command("curl", "-L", ctl_bin_url, "-o", "binary.bin")
	err = ctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ctl_src_url := "https://github.com/ampas/CTL/archive/refs/tags/ctl-1.5.3.src.tar.gz"
	ctl_cmd_src := exec.Command("curl", "-L", ctl_src_url, "-o", "source.tar.gz")
	err = ctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ctl_cmd_direct := exec.Command("./binary")
	err = ctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: aces_container")
exec.Command("latte", "install", "aces_container")
	fmt.Println("Instalando dependencia: imath")
exec.Command("latte", "install", "imath")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: openexr")
exec.Command("latte", "install", "openexr")
}
