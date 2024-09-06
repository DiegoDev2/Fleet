package main

// Bloaty - Size profiler for binaries
// Homepage: https://github.com/google/bloaty

import (
	"fmt"
	
	"os/exec"
)

func installBloaty() {
	// Método 1: Descargar y extraer .tar.gz
	bloaty_tar_url := "https://github.com/google/bloaty/releases/download/v1.1/bloaty-1.1.tar.bz2"
	bloaty_cmd_tar := exec.Command("curl", "-L", bloaty_tar_url, "-o", "package.tar.gz")
	err := bloaty_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bloaty_zip_url := "https://github.com/google/bloaty/releases/download/v1.1/bloaty-1.1.tar.bz2"
	bloaty_cmd_zip := exec.Command("curl", "-L", bloaty_zip_url, "-o", "package.zip")
	err = bloaty_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bloaty_bin_url := "https://github.com/google/bloaty/releases/download/v1.1/bloaty-1.1.tar.bz2"
	bloaty_cmd_bin := exec.Command("curl", "-L", bloaty_bin_url, "-o", "binary.bin")
	err = bloaty_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bloaty_src_url := "https://github.com/google/bloaty/releases/download/v1.1/bloaty-1.1.tar.bz2"
	bloaty_cmd_src := exec.Command("curl", "-L", bloaty_src_url, "-o", "source.tar.gz")
	err = bloaty_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bloaty_cmd_direct := exec.Command("./binary")
	err = bloaty_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: capstone")
	exec.Command("latte", "install", "capstone").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: re2")
	exec.Command("latte", "install", "re2").Run()
}
