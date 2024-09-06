package main

// Libosmium - Fast and flexible C++ library for working with OpenStreetMap data
// Homepage: https://osmcode.org/libosmium/

import (
	"fmt"
	
	"os/exec"
)

func installLibosmium() {
	// Método 1: Descargar y extraer .tar.gz
	libosmium_tar_url := "https://github.com/osmcode/libosmium/archive/refs/tags/v2.20.0.tar.gz"
	libosmium_cmd_tar := exec.Command("curl", "-L", libosmium_tar_url, "-o", "package.tar.gz")
	err := libosmium_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libosmium_zip_url := "https://github.com/osmcode/libosmium/archive/refs/tags/v2.20.0.zip"
	libosmium_cmd_zip := exec.Command("curl", "-L", libosmium_zip_url, "-o", "package.zip")
	err = libosmium_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libosmium_bin_url := "https://github.com/osmcode/libosmium/archive/refs/tags/v2.20.0.bin"
	libosmium_cmd_bin := exec.Command("curl", "-L", libosmium_bin_url, "-o", "binary.bin")
	err = libosmium_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libosmium_src_url := "https://github.com/osmcode/libosmium/archive/refs/tags/v2.20.0.src.tar.gz"
	libosmium_cmd_src := exec.Command("curl", "-L", libosmium_src_url, "-o", "source.tar.gz")
	err = libosmium_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libosmium_cmd_direct := exec.Command("./binary")
	err = libosmium_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
}
