package main

// Partio - Particle library for 3D graphics
// Homepage: https://github.com/wdas/partio

import (
	"fmt"
	
	"os/exec"
)

func installPartio() {
	// Método 1: Descargar y extraer .tar.gz
	partio_tar_url := "https://github.com/wdas/partio/archive/refs/tags/v1.17.3.tar.gz"
	partio_cmd_tar := exec.Command("curl", "-L", partio_tar_url, "-o", "package.tar.gz")
	err := partio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	partio_zip_url := "https://github.com/wdas/partio/archive/refs/tags/v1.17.3.zip"
	partio_cmd_zip := exec.Command("curl", "-L", partio_zip_url, "-o", "package.zip")
	err = partio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	partio_bin_url := "https://github.com/wdas/partio/archive/refs/tags/v1.17.3.bin"
	partio_cmd_bin := exec.Command("curl", "-L", partio_bin_url, "-o", "binary.bin")
	err = partio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	partio_src_url := "https://github.com/wdas/partio/archive/refs/tags/v1.17.3.src.tar.gz"
	partio_cmd_src := exec.Command("curl", "-L", partio_src_url, "-o", "source.tar.gz")
	err = partio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	partio_cmd_direct := exec.Command("./binary")
	err = partio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: freeglut")
exec.Command("latte", "install", "freeglut")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
}
