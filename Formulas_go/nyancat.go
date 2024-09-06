package main

// Nyancat - Renders an animated, color, ANSI-text loop of the Poptart Cat
// Homepage: https://github.com/klange/nyancat

import (
	"fmt"
	
	"os/exec"
)

func installNyancat() {
	// Método 1: Descargar y extraer .tar.gz
	nyancat_tar_url := "https://github.com/klange/nyancat/archive/refs/tags/1.5.2.tar.gz"
	nyancat_cmd_tar := exec.Command("curl", "-L", nyancat_tar_url, "-o", "package.tar.gz")
	err := nyancat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nyancat_zip_url := "https://github.com/klange/nyancat/archive/refs/tags/1.5.2.zip"
	nyancat_cmd_zip := exec.Command("curl", "-L", nyancat_zip_url, "-o", "package.zip")
	err = nyancat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nyancat_bin_url := "https://github.com/klange/nyancat/archive/refs/tags/1.5.2.bin"
	nyancat_cmd_bin := exec.Command("curl", "-L", nyancat_bin_url, "-o", "binary.bin")
	err = nyancat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nyancat_src_url := "https://github.com/klange/nyancat/archive/refs/tags/1.5.2.src.tar.gz"
	nyancat_cmd_src := exec.Command("curl", "-L", nyancat_src_url, "-o", "source.tar.gz")
	err = nyancat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nyancat_cmd_direct := exec.Command("./binary")
	err = nyancat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
