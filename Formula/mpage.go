package main

// Mpage - Many to one page printing utility
// Homepage: https://mesa.nl/pub/mpage/README

import (
	"fmt"
	
	"os/exec"
)

func installMpage() {
	// Método 1: Descargar y extraer .tar.gz
	mpage_tar_url := "https://mesa.nl/pub/mpage/mpage-2.5.8.tgz"
	mpage_cmd_tar := exec.Command("curl", "-L", mpage_tar_url, "-o", "package.tar.gz")
	err := mpage_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpage_zip_url := "https://mesa.nl/pub/mpage/mpage-2.5.8.tgz"
	mpage_cmd_zip := exec.Command("curl", "-L", mpage_zip_url, "-o", "package.zip")
	err = mpage_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpage_bin_url := "https://mesa.nl/pub/mpage/mpage-2.5.8.tgz"
	mpage_cmd_bin := exec.Command("curl", "-L", mpage_bin_url, "-o", "binary.bin")
	err = mpage_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpage_src_url := "https://mesa.nl/pub/mpage/mpage-2.5.8.tgz"
	mpage_cmd_src := exec.Command("curl", "-L", mpage_src_url, "-o", "source.tar.gz")
	err = mpage_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpage_cmd_direct := exec.Command("./binary")
	err = mpage_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
