package main

// Mpg123 - MP3 player for Linux and UNIX
// Homepage: https://www.mpg123.de/

import (
	"fmt"
	
	"os/exec"
)

func installMpg123() {
	// Método 1: Descargar y extraer .tar.gz
	mpg123_tar_url := "https://www.mpg123.de/download/mpg123-1.32.7.tar.bz2"
	mpg123_cmd_tar := exec.Command("curl", "-L", mpg123_tar_url, "-o", "package.tar.gz")
	err := mpg123_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpg123_zip_url := "https://www.mpg123.de/download/mpg123-1.32.7.tar.bz2"
	mpg123_cmd_zip := exec.Command("curl", "-L", mpg123_zip_url, "-o", "package.zip")
	err = mpg123_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpg123_bin_url := "https://www.mpg123.de/download/mpg123-1.32.7.tar.bz2"
	mpg123_cmd_bin := exec.Command("curl", "-L", mpg123_bin_url, "-o", "binary.bin")
	err = mpg123_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpg123_src_url := "https://www.mpg123.de/download/mpg123-1.32.7.tar.bz2"
	mpg123_cmd_src := exec.Command("curl", "-L", mpg123_src_url, "-o", "source.tar.gz")
	err = mpg123_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpg123_cmd_direct := exec.Command("./binary")
	err = mpg123_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
