package main

// Cubelib - Cube, is a performance report explorer for Scalasca and Score-P
// Homepage: https://scalasca.org/software/cube-4.x/download.html

import (
	"fmt"
	
	"os/exec"
)

func installCubelib() {
	// Método 1: Descargar y extraer .tar.gz
	cubelib_tar_url := "https://apps.fz-juelich.de/scalasca/releases/cube/4.8/dist/cubelib-4.8.2.tar.gz"
	cubelib_cmd_tar := exec.Command("curl", "-L", cubelib_tar_url, "-o", "package.tar.gz")
	err := cubelib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cubelib_zip_url := "https://apps.fz-juelich.de/scalasca/releases/cube/4.8/dist/cubelib-4.8.2.zip"
	cubelib_cmd_zip := exec.Command("curl", "-L", cubelib_zip_url, "-o", "package.zip")
	err = cubelib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cubelib_bin_url := "https://apps.fz-juelich.de/scalasca/releases/cube/4.8/dist/cubelib-4.8.2.bin"
	cubelib_cmd_bin := exec.Command("curl", "-L", cubelib_bin_url, "-o", "binary.bin")
	err = cubelib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cubelib_src_url := "https://apps.fz-juelich.de/scalasca/releases/cube/4.8/dist/cubelib-4.8.2.src.tar.gz"
	cubelib_cmd_src := exec.Command("curl", "-L", cubelib_src_url, "-o", "source.tar.gz")
	err = cubelib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cubelib_cmd_direct := exec.Command("./binary")
	err = cubelib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
