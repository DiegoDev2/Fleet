package main

// Librist - Reliable Internet Stream Transport (RIST)
// Homepage: https://code.videolan.org/rist/

import (
	"fmt"
	
	"os/exec"
)

func installLibrist() {
	// Método 1: Descargar y extraer .tar.gz
	librist_tar_url := "https://code.videolan.org/rist/librist/-/archive/v0.2.10/librist-v0.2.10.tar.gz"
	librist_cmd_tar := exec.Command("curl", "-L", librist_tar_url, "-o", "package.tar.gz")
	err := librist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librist_zip_url := "https://code.videolan.org/rist/librist/-/archive/v0.2.10/librist-v0.2.10.zip"
	librist_cmd_zip := exec.Command("curl", "-L", librist_zip_url, "-o", "package.zip")
	err = librist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librist_bin_url := "https://code.videolan.org/rist/librist/-/archive/v0.2.10/librist-v0.2.10.bin"
	librist_cmd_bin := exec.Command("curl", "-L", librist_bin_url, "-o", "binary.bin")
	err = librist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librist_src_url := "https://code.videolan.org/rist/librist/-/archive/v0.2.10/librist-v0.2.10.src.tar.gz"
	librist_cmd_src := exec.Command("curl", "-L", librist_src_url, "-o", "source.tar.gz")
	err = librist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librist_cmd_direct := exec.Command("./binary")
	err = librist_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: cjson")
exec.Command("latte", "install", "cjson")
	fmt.Println("Instalando dependencia: libmicrohttpd")
exec.Command("latte", "install", "libmicrohttpd")
	fmt.Println("Instalando dependencia: mbedtls")
exec.Command("latte", "install", "mbedtls")
}
