package main

// Librem - Toolkit library for real-time audio and video processing
// Homepage: https://github.com/baresip/rem

import (
	"fmt"
	
	"os/exec"
)

func installLibrem() {
	// Método 1: Descargar y extraer .tar.gz
	librem_tar_url := "https://github.com/baresip/rem/archive/refs/tags/v2.12.0.tar.gz"
	librem_cmd_tar := exec.Command("curl", "-L", librem_tar_url, "-o", "package.tar.gz")
	err := librem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librem_zip_url := "https://github.com/baresip/rem/archive/refs/tags/v2.12.0.zip"
	librem_cmd_zip := exec.Command("curl", "-L", librem_zip_url, "-o", "package.zip")
	err = librem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librem_bin_url := "https://github.com/baresip/rem/archive/refs/tags/v2.12.0.bin"
	librem_cmd_bin := exec.Command("curl", "-L", librem_bin_url, "-o", "binary.bin")
	err = librem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librem_src_url := "https://github.com/baresip/rem/archive/refs/tags/v2.12.0.src.tar.gz"
	librem_cmd_src := exec.Command("curl", "-L", librem_src_url, "-o", "source.tar.gz")
	err = librem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librem_cmd_direct := exec.Command("./binary")
	err = librem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libre")
exec.Command("latte", "install", "libre")
}
