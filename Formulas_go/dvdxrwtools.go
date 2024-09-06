package main

// DvdxrwTools - DVD+-RW/R tools
// Homepage: https://fy.chalmers.se/~appro/linux/DVD+RW/

import (
	"fmt"
	
	"os/exec"
)

func installDvdxrwTools() {
	// Método 1: Descargar y extraer .tar.gz
	dvdxrwtools_tar_url := "https://fy.chalmers.se/~appro/linux/DVD+RW/tools/dvd+rw-tools-7.1.tar.gz"
	dvdxrwtools_cmd_tar := exec.Command("curl", "-L", dvdxrwtools_tar_url, "-o", "package.tar.gz")
	err := dvdxrwtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dvdxrwtools_zip_url := "https://fy.chalmers.se/~appro/linux/DVD+RW/tools/dvd+rw-tools-7.1.zip"
	dvdxrwtools_cmd_zip := exec.Command("curl", "-L", dvdxrwtools_zip_url, "-o", "package.zip")
	err = dvdxrwtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dvdxrwtools_bin_url := "https://fy.chalmers.se/~appro/linux/DVD+RW/tools/dvd+rw-tools-7.1.bin"
	dvdxrwtools_cmd_bin := exec.Command("curl", "-L", dvdxrwtools_bin_url, "-o", "binary.bin")
	err = dvdxrwtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dvdxrwtools_src_url := "https://fy.chalmers.se/~appro/linux/DVD+RW/tools/dvd+rw-tools-7.1.src.tar.gz"
	dvdxrwtools_cmd_src := exec.Command("curl", "-L", dvdxrwtools_src_url, "-o", "source.tar.gz")
	err = dvdxrwtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dvdxrwtools_cmd_direct := exec.Command("./binary")
	err = dvdxrwtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
