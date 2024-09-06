package main

// BoostAT185 - Collection of portable C++ source libraries
// Homepage: https://www.boost.org/

import (
	"fmt"
	
	"os/exec"
)

func installBoostAT185() {
	// Método 1: Descargar y extraer .tar.gz
	boostat185_tar_url := "https://github.com/boostorg/boost/releases/download/boost-1.85.0/boost-1.85.0-b2-nodocs.tar.xz"
	boostat185_cmd_tar := exec.Command("curl", "-L", boostat185_tar_url, "-o", "package.tar.gz")
	err := boostat185_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	boostat185_zip_url := "https://github.com/boostorg/boost/releases/download/boost-1.85.0/boost-1.85.0-b2-nodocs.tar.xz"
	boostat185_cmd_zip := exec.Command("curl", "-L", boostat185_zip_url, "-o", "package.zip")
	err = boostat185_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	boostat185_bin_url := "https://github.com/boostorg/boost/releases/download/boost-1.85.0/boost-1.85.0-b2-nodocs.tar.xz"
	boostat185_cmd_bin := exec.Command("curl", "-L", boostat185_bin_url, "-o", "binary.bin")
	err = boostat185_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	boostat185_src_url := "https://github.com/boostorg/boost/releases/download/boost-1.85.0/boost-1.85.0-b2-nodocs.tar.xz"
	boostat185_cmd_src := exec.Command("curl", "-L", boostat185_src_url, "-o", "source.tar.gz")
	err = boostat185_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	boostat185_cmd_direct := exec.Command("./binary")
	err = boostat185_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
