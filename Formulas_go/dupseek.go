package main

// Dupseek - Interactive program to find and remove duplicate files
// Homepage: http://www.beautylabs.net/software/dupseek.html

import (
	"fmt"
	
	"os/exec"
)

func installDupseek() {
	// Método 1: Descargar y extraer .tar.gz
	dupseek_tar_url := "http://www.beautylabs.net/software/dupseek-1.3.tgz"
	dupseek_cmd_tar := exec.Command("curl", "-L", dupseek_tar_url, "-o", "package.tar.gz")
	err := dupseek_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dupseek_zip_url := "http://www.beautylabs.net/software/dupseek-1.3.tgz"
	dupseek_cmd_zip := exec.Command("curl", "-L", dupseek_zip_url, "-o", "package.zip")
	err = dupseek_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dupseek_bin_url := "http://www.beautylabs.net/software/dupseek-1.3.tgz"
	dupseek_cmd_bin := exec.Command("curl", "-L", dupseek_bin_url, "-o", "binary.bin")
	err = dupseek_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dupseek_src_url := "http://www.beautylabs.net/software/dupseek-1.3.tgz"
	dupseek_cmd_src := exec.Command("curl", "-L", dupseek_src_url, "-o", "source.tar.gz")
	err = dupseek_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dupseek_cmd_direct := exec.Command("./binary")
	err = dupseek_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
