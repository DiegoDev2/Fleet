package main

// Grt - Gesture Recognition Toolkit for real-time machine learning
// Homepage: https://nickgillian.com/grt/

import (
	"fmt"
	
	"os/exec"
)

func installGrt() {
	// Método 1: Descargar y extraer .tar.gz
	grt_tar_url := "https://github.com/nickgillian/grt/archive/refs/tags/v0.2.4.tar.gz"
	grt_cmd_tar := exec.Command("curl", "-L", grt_tar_url, "-o", "package.tar.gz")
	err := grt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grt_zip_url := "https://github.com/nickgillian/grt/archive/refs/tags/v0.2.4.zip"
	grt_cmd_zip := exec.Command("curl", "-L", grt_zip_url, "-o", "package.zip")
	err = grt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grt_bin_url := "https://github.com/nickgillian/grt/archive/refs/tags/v0.2.4.bin"
	grt_cmd_bin := exec.Command("curl", "-L", grt_bin_url, "-o", "binary.bin")
	err = grt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grt_src_url := "https://github.com/nickgillian/grt/archive/refs/tags/v0.2.4.src.tar.gz"
	grt_cmd_src := exec.Command("curl", "-L", grt_src_url, "-o", "source.tar.gz")
	err = grt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grt_cmd_direct := exec.Command("./binary")
	err = grt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
