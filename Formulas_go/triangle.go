package main

// Triangle - Convert images to computer generated art using Delaunay triangulation
// Homepage: https://github.com/esimov/triangle

import (
	"fmt"
	
	"os/exec"
)

func installTriangle() {
	// Método 1: Descargar y extraer .tar.gz
	triangle_tar_url := "https://github.com/esimov/triangle/archive/refs/tags/v2.0.0.tar.gz"
	triangle_cmd_tar := exec.Command("curl", "-L", triangle_tar_url, "-o", "package.tar.gz")
	err := triangle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	triangle_zip_url := "https://github.com/esimov/triangle/archive/refs/tags/v2.0.0.zip"
	triangle_cmd_zip := exec.Command("curl", "-L", triangle_zip_url, "-o", "package.zip")
	err = triangle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	triangle_bin_url := "https://github.com/esimov/triangle/archive/refs/tags/v2.0.0.bin"
	triangle_cmd_bin := exec.Command("curl", "-L", triangle_bin_url, "-o", "binary.bin")
	err = triangle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	triangle_src_url := "https://github.com/esimov/triangle/archive/refs/tags/v2.0.0.src.tar.gz"
	triangle_cmd_src := exec.Command("curl", "-L", triangle_src_url, "-o", "source.tar.gz")
	err = triangle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	triangle_cmd_direct := exec.Command("./binary")
	err = triangle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
