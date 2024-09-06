package main

// Nanoflann - Header-only library for Nearest Neighbor search with KD-trees
// Homepage: https://github.com/jlblancoc/nanoflann

import (
	"fmt"
	
	"os/exec"
)

func installNanoflann() {
	// Método 1: Descargar y extraer .tar.gz
	nanoflann_tar_url := "https://github.com/jlblancoc/nanoflann/archive/refs/tags/v1.6.1.tar.gz"
	nanoflann_cmd_tar := exec.Command("curl", "-L", nanoflann_tar_url, "-o", "package.tar.gz")
	err := nanoflann_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nanoflann_zip_url := "https://github.com/jlblancoc/nanoflann/archive/refs/tags/v1.6.1.zip"
	nanoflann_cmd_zip := exec.Command("curl", "-L", nanoflann_zip_url, "-o", "package.zip")
	err = nanoflann_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nanoflann_bin_url := "https://github.com/jlblancoc/nanoflann/archive/refs/tags/v1.6.1.bin"
	nanoflann_cmd_bin := exec.Command("curl", "-L", nanoflann_bin_url, "-o", "binary.bin")
	err = nanoflann_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nanoflann_src_url := "https://github.com/jlblancoc/nanoflann/archive/refs/tags/v1.6.1.src.tar.gz"
	nanoflann_cmd_src := exec.Command("curl", "-L", nanoflann_src_url, "-o", "source.tar.gz")
	err = nanoflann_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nanoflann_cmd_direct := exec.Command("./binary")
	err = nanoflann_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
}
