package main

// Flann - Fast Library for Approximate Nearest Neighbors
// Homepage: https://github.com/flann-lib/flann

import (
	"fmt"
	
	"os/exec"
)

func installFlann() {
	// Método 1: Descargar y extraer .tar.gz
	flann_tar_url := "https://github.com/flann-lib/flann/archive/refs/tags/1.9.2.tar.gz"
	flann_cmd_tar := exec.Command("curl", "-L", flann_tar_url, "-o", "package.tar.gz")
	err := flann_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flann_zip_url := "https://github.com/flann-lib/flann/archive/refs/tags/1.9.2.zip"
	flann_cmd_zip := exec.Command("curl", "-L", flann_zip_url, "-o", "package.zip")
	err = flann_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flann_bin_url := "https://github.com/flann-lib/flann/archive/refs/tags/1.9.2.bin"
	flann_cmd_bin := exec.Command("curl", "-L", flann_bin_url, "-o", "binary.bin")
	err = flann_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flann_src_url := "https://github.com/flann-lib/flann/archive/refs/tags/1.9.2.src.tar.gz"
	flann_cmd_src := exec.Command("curl", "-L", flann_src_url, "-o", "source.tar.gz")
	err = flann_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flann_cmd_direct := exec.Command("./binary")
	err = flann_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
}
