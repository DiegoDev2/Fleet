package main

// GnuGo - Plays the game of Go
// Homepage: https://www.gnu.org/software/gnugo/gnugo.html

import (
	"fmt"
	
	"os/exec"
)

func installGnuGo() {
	// Método 1: Descargar y extraer .tar.gz
	gnugo_tar_url := "https://ftp.gnu.org/gnu/gnugo/gnugo-3.8.tar.gz"
	gnugo_cmd_tar := exec.Command("curl", "-L", gnugo_tar_url, "-o", "package.tar.gz")
	err := gnugo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnugo_zip_url := "https://ftp.gnu.org/gnu/gnugo/gnugo-3.8.zip"
	gnugo_cmd_zip := exec.Command("curl", "-L", gnugo_zip_url, "-o", "package.zip")
	err = gnugo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnugo_bin_url := "https://ftp.gnu.org/gnu/gnugo/gnugo-3.8.bin"
	gnugo_cmd_bin := exec.Command("curl", "-L", gnugo_bin_url, "-o", "binary.bin")
	err = gnugo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnugo_src_url := "https://ftp.gnu.org/gnu/gnugo/gnugo-3.8.src.tar.gz"
	gnugo_cmd_src := exec.Command("curl", "-L", gnugo_src_url, "-o", "source.tar.gz")
	err = gnugo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnugo_cmd_direct := exec.Command("./binary")
	err = gnugo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
