package main

// Libebur128 - Library implementing the EBU R128 loudness standard
// Homepage: https://github.com/jiixyj/libebur128

import (
	"fmt"
	
	"os/exec"
)

func installLibebur128() {
	// Método 1: Descargar y extraer .tar.gz
	libebur128_tar_url := "https://github.com/jiixyj/libebur128/archive/refs/tags/v1.2.6.tar.gz"
	libebur128_cmd_tar := exec.Command("curl", "-L", libebur128_tar_url, "-o", "package.tar.gz")
	err := libebur128_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libebur128_zip_url := "https://github.com/jiixyj/libebur128/archive/refs/tags/v1.2.6.zip"
	libebur128_cmd_zip := exec.Command("curl", "-L", libebur128_zip_url, "-o", "package.zip")
	err = libebur128_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libebur128_bin_url := "https://github.com/jiixyj/libebur128/archive/refs/tags/v1.2.6.bin"
	libebur128_cmd_bin := exec.Command("curl", "-L", libebur128_bin_url, "-o", "binary.bin")
	err = libebur128_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libebur128_src_url := "https://github.com/jiixyj/libebur128/archive/refs/tags/v1.2.6.src.tar.gz"
	libebur128_cmd_src := exec.Command("curl", "-L", libebur128_src_url, "-o", "source.tar.gz")
	err = libebur128_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libebur128_cmd_direct := exec.Command("./binary")
	err = libebur128_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: speex")
exec.Command("latte", "install", "speex")
}
