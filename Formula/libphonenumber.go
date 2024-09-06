package main

// Libphonenumber - C++ Phone Number library by Google
// Homepage: https://github.com/google/libphonenumber

import (
	"fmt"
	
	"os/exec"
)

func installLibphonenumber() {
	// Método 1: Descargar y extraer .tar.gz
	libphonenumber_tar_url := "https://github.com/google/libphonenumber/archive/refs/tags/v8.13.44.tar.gz"
	libphonenumber_cmd_tar := exec.Command("curl", "-L", libphonenumber_tar_url, "-o", "package.tar.gz")
	err := libphonenumber_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libphonenumber_zip_url := "https://github.com/google/libphonenumber/archive/refs/tags/v8.13.44.zip"
	libphonenumber_cmd_zip := exec.Command("curl", "-L", libphonenumber_zip_url, "-o", "package.zip")
	err = libphonenumber_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libphonenumber_bin_url := "https://github.com/google/libphonenumber/archive/refs/tags/v8.13.44.bin"
	libphonenumber_cmd_bin := exec.Command("curl", "-L", libphonenumber_bin_url, "-o", "binary.bin")
	err = libphonenumber_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libphonenumber_src_url := "https://github.com/google/libphonenumber/archive/refs/tags/v8.13.44.src.tar.gz"
	libphonenumber_cmd_src := exec.Command("curl", "-L", libphonenumber_src_url, "-o", "source.tar.gz")
	err = libphonenumber_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libphonenumber_cmd_direct := exec.Command("./binary")
	err = libphonenumber_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: googletest")
	exec.Command("latte", "install", "googletest").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: re2")
	exec.Command("latte", "install", "re2").Run()
}
