package main

// Libcue - Cue sheet parser library for C
// Homepage: https://github.com/lipnitsk/libcue

import (
	"fmt"
	
	"os/exec"
)

func installLibcue() {
	// Método 1: Descargar y extraer .tar.gz
	libcue_tar_url := "https://github.com/lipnitsk/libcue/archive/refs/tags/v2.3.0.tar.gz"
	libcue_cmd_tar := exec.Command("curl", "-L", libcue_tar_url, "-o", "package.tar.gz")
	err := libcue_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcue_zip_url := "https://github.com/lipnitsk/libcue/archive/refs/tags/v2.3.0.zip"
	libcue_cmd_zip := exec.Command("curl", "-L", libcue_zip_url, "-o", "package.zip")
	err = libcue_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcue_bin_url := "https://github.com/lipnitsk/libcue/archive/refs/tags/v2.3.0.bin"
	libcue_cmd_bin := exec.Command("curl", "-L", libcue_bin_url, "-o", "binary.bin")
	err = libcue_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcue_src_url := "https://github.com/lipnitsk/libcue/archive/refs/tags/v2.3.0.src.tar.gz"
	libcue_cmd_src := exec.Command("curl", "-L", libcue_src_url, "-o", "source.tar.gz")
	err = libcue_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcue_cmd_direct := exec.Command("./binary")
	err = libcue_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
