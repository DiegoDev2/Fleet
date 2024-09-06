package main

// Vtclock - Text-mode fullscreen digital clock
// Homepage: https://github.com/dse/vtclock

import (
	"fmt"
	
	"os/exec"
)

func installVtclock() {
	// Método 1: Descargar y extraer .tar.gz
	vtclock_tar_url := "https://github.com/dse/vtclock/archive/refs/tags/v0.99.1.tar.gz"
	vtclock_cmd_tar := exec.Command("curl", "-L", vtclock_tar_url, "-o", "package.tar.gz")
	err := vtclock_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vtclock_zip_url := "https://github.com/dse/vtclock/archive/refs/tags/v0.99.1.zip"
	vtclock_cmd_zip := exec.Command("curl", "-L", vtclock_zip_url, "-o", "package.zip")
	err = vtclock_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vtclock_bin_url := "https://github.com/dse/vtclock/archive/refs/tags/v0.99.1.bin"
	vtclock_cmd_bin := exec.Command("curl", "-L", vtclock_bin_url, "-o", "binary.bin")
	err = vtclock_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vtclock_src_url := "https://github.com/dse/vtclock/archive/refs/tags/v0.99.1.src.tar.gz"
	vtclock_cmd_src := exec.Command("curl", "-L", vtclock_src_url, "-o", "source.tar.gz")
	err = vtclock_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vtclock_cmd_direct := exec.Command("./binary")
	err = vtclock_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
