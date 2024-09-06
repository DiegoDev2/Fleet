package main

// Sersniff - Program to tunnel/sniff between 2 serial ports
// Homepage: https://www.earth.li/projectpurple/progs/sersniff.html

import (
	"fmt"
	
	"os/exec"
)

func installSersniff() {
	// Método 1: Descargar y extraer .tar.gz
	sersniff_tar_url := "https://www.earth.li/projectpurple/files/sersniff-0.0.5.tar.gz"
	sersniff_cmd_tar := exec.Command("curl", "-L", sersniff_tar_url, "-o", "package.tar.gz")
	err := sersniff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sersniff_zip_url := "https://www.earth.li/projectpurple/files/sersniff-0.0.5.zip"
	sersniff_cmd_zip := exec.Command("curl", "-L", sersniff_zip_url, "-o", "package.zip")
	err = sersniff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sersniff_bin_url := "https://www.earth.li/projectpurple/files/sersniff-0.0.5.bin"
	sersniff_cmd_bin := exec.Command("curl", "-L", sersniff_bin_url, "-o", "binary.bin")
	err = sersniff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sersniff_src_url := "https://www.earth.li/projectpurple/files/sersniff-0.0.5.src.tar.gz"
	sersniff_cmd_src := exec.Command("curl", "-L", sersniff_src_url, "-o", "source.tar.gz")
	err = sersniff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sersniff_cmd_direct := exec.Command("./binary")
	err = sersniff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
