package main

// Libmatroska - Extensible, open standard container format for audio/video
// Homepage: https://www.matroska.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibmatroska() {
	// Método 1: Descargar y extraer .tar.gz
	libmatroska_tar_url := "https://dl.matroska.org/downloads/libmatroska/libmatroska-1.7.1.tar.xz"
	libmatroska_cmd_tar := exec.Command("curl", "-L", libmatroska_tar_url, "-o", "package.tar.gz")
	err := libmatroska_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmatroska_zip_url := "https://dl.matroska.org/downloads/libmatroska/libmatroska-1.7.1.tar.xz"
	libmatroska_cmd_zip := exec.Command("curl", "-L", libmatroska_zip_url, "-o", "package.zip")
	err = libmatroska_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmatroska_bin_url := "https://dl.matroska.org/downloads/libmatroska/libmatroska-1.7.1.tar.xz"
	libmatroska_cmd_bin := exec.Command("curl", "-L", libmatroska_bin_url, "-o", "binary.bin")
	err = libmatroska_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmatroska_src_url := "https://dl.matroska.org/downloads/libmatroska/libmatroska-1.7.1.tar.xz"
	libmatroska_cmd_src := exec.Command("curl", "-L", libmatroska_src_url, "-o", "source.tar.gz")
	err = libmatroska_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmatroska_cmd_direct := exec.Command("./binary")
	err = libmatroska_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libebml")
exec.Command("latte", "install", "libebml")
}
