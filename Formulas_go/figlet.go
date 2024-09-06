package main

// Figlet - Banner-like program prints strings as ASCII art
// Homepage: http://www.figlet.org/

import (
	"fmt"
	
	"os/exec"
)

func installFiglet() {
	// Método 1: Descargar y extraer .tar.gz
	figlet_tar_url := "http://ftp.figlet.org/pub/figlet/program/unix/figlet-2.2.5.tar.gz"
	figlet_cmd_tar := exec.Command("curl", "-L", figlet_tar_url, "-o", "package.tar.gz")
	err := figlet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	figlet_zip_url := "http://ftp.figlet.org/pub/figlet/program/unix/figlet-2.2.5.zip"
	figlet_cmd_zip := exec.Command("curl", "-L", figlet_zip_url, "-o", "package.zip")
	err = figlet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	figlet_bin_url := "http://ftp.figlet.org/pub/figlet/program/unix/figlet-2.2.5.bin"
	figlet_cmd_bin := exec.Command("curl", "-L", figlet_bin_url, "-o", "binary.bin")
	err = figlet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	figlet_src_url := "http://ftp.figlet.org/pub/figlet/program/unix/figlet-2.2.5.src.tar.gz"
	figlet_cmd_src := exec.Command("curl", "-L", figlet_src_url, "-o", "source.tar.gz")
	err = figlet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	figlet_cmd_direct := exec.Command("./binary")
	err = figlet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
