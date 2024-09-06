package main

// Qhull - Computes convex hulls in n dimensions
// Homepage: http://www.qhull.org/

import (
	"fmt"
	
	"os/exec"
)

func installQhull() {
	// Método 1: Descargar y extraer .tar.gz
	qhull_tar_url := "http://www.qhull.org/download/qhull-2020-src-8.0.2.tgz"
	qhull_cmd_tar := exec.Command("curl", "-L", qhull_tar_url, "-o", "package.tar.gz")
	err := qhull_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qhull_zip_url := "http://www.qhull.org/download/qhull-2020-src-8.0.2.tgz"
	qhull_cmd_zip := exec.Command("curl", "-L", qhull_zip_url, "-o", "package.zip")
	err = qhull_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qhull_bin_url := "http://www.qhull.org/download/qhull-2020-src-8.0.2.tgz"
	qhull_cmd_bin := exec.Command("curl", "-L", qhull_bin_url, "-o", "binary.bin")
	err = qhull_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qhull_src_url := "http://www.qhull.org/download/qhull-2020-src-8.0.2.tgz"
	qhull_cmd_src := exec.Command("curl", "-L", qhull_src_url, "-o", "source.tar.gz")
	err = qhull_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qhull_cmd_direct := exec.Command("./binary")
	err = qhull_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
