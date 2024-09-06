package main

// Golo - Lightweight dynamic language for the JVM
// Homepage: https://golo-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installGolo() {
	// Método 1: Descargar y extraer .tar.gz
	golo_tar_url := "https://github.com/eclipse-archived/golo-lang/releases/download/release%2F3.4.0/golo-3.4.0.zip"
	golo_cmd_tar := exec.Command("curl", "-L", golo_tar_url, "-o", "package.tar.gz")
	err := golo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	golo_zip_url := "https://github.com/eclipse-archived/golo-lang/releases/download/release%2F3.4.0/golo-3.4.0.zip"
	golo_cmd_zip := exec.Command("curl", "-L", golo_zip_url, "-o", "package.zip")
	err = golo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	golo_bin_url := "https://github.com/eclipse-archived/golo-lang/releases/download/release%2F3.4.0/golo-3.4.0.zip"
	golo_cmd_bin := exec.Command("curl", "-L", golo_bin_url, "-o", "binary.bin")
	err = golo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	golo_src_url := "https://github.com/eclipse-archived/golo-lang/releases/download/release%2F3.4.0/golo-3.4.0.zip"
	golo_cmd_src := exec.Command("curl", "-L", golo_src_url, "-o", "source.tar.gz")
	err = golo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	golo_cmd_direct := exec.Command("./binary")
	err = golo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}
