package main

// Since - Stateful tail: show changes to files since last check
// Homepage: http://welz.org.za/projects/since

import (
	"fmt"
	
	"os/exec"
)

func installSince() {
	// Método 1: Descargar y extraer .tar.gz
	since_tar_url := "http://welz.org.za/projects/since/since-1.1.tar.gz"
	since_cmd_tar := exec.Command("curl", "-L", since_tar_url, "-o", "package.tar.gz")
	err := since_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	since_zip_url := "http://welz.org.za/projects/since/since-1.1.zip"
	since_cmd_zip := exec.Command("curl", "-L", since_zip_url, "-o", "package.zip")
	err = since_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	since_bin_url := "http://welz.org.za/projects/since/since-1.1.bin"
	since_cmd_bin := exec.Command("curl", "-L", since_bin_url, "-o", "binary.bin")
	err = since_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	since_src_url := "http://welz.org.za/projects/since/since-1.1.src.tar.gz"
	since_cmd_src := exec.Command("curl", "-L", since_src_url, "-o", "source.tar.gz")
	err = since_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	since_cmd_direct := exec.Command("./binary")
	err = since_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
