package main

// Gcviewer - Java garbage collection visualization tool
// Homepage: https://github.com/chewiebug/GCViewer

import (
	"fmt"
	
	"os/exec"
)

func installGcviewer() {
	// Método 1: Descargar y extraer .tar.gz
	gcviewer_tar_url := "https://downloads.sourceforge.net/project/gcviewer/gcviewer-1.36.jar"
	gcviewer_cmd_tar := exec.Command("curl", "-L", gcviewer_tar_url, "-o", "package.tar.gz")
	err := gcviewer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gcviewer_zip_url := "https://downloads.sourceforge.net/project/gcviewer/gcviewer-1.36.jar"
	gcviewer_cmd_zip := exec.Command("curl", "-L", gcviewer_zip_url, "-o", "package.zip")
	err = gcviewer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gcviewer_bin_url := "https://downloads.sourceforge.net/project/gcviewer/gcviewer-1.36.jar"
	gcviewer_cmd_bin := exec.Command("curl", "-L", gcviewer_bin_url, "-o", "binary.bin")
	err = gcviewer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gcviewer_src_url := "https://downloads.sourceforge.net/project/gcviewer/gcviewer-1.36.jar"
	gcviewer_cmd_src := exec.Command("curl", "-L", gcviewer_src_url, "-o", "source.tar.gz")
	err = gcviewer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gcviewer_cmd_direct := exec.Command("./binary")
	err = gcviewer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
