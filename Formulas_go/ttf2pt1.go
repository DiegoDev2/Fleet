package main

// Ttf2pt1 - True Type Font to Postscript Type 1 converter
// Homepage: https://ttf2pt1.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installTtf2pt1() {
	// Método 1: Descargar y extraer .tar.gz
	ttf2pt1_tar_url := "https://downloads.sourceforge.net/project/ttf2pt1/ttf2pt1/3.4.4/ttf2pt1-3.4.4.tgz"
	ttf2pt1_cmd_tar := exec.Command("curl", "-L", ttf2pt1_tar_url, "-o", "package.tar.gz")
	err := ttf2pt1_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ttf2pt1_zip_url := "https://downloads.sourceforge.net/project/ttf2pt1/ttf2pt1/3.4.4/ttf2pt1-3.4.4.tgz"
	ttf2pt1_cmd_zip := exec.Command("curl", "-L", ttf2pt1_zip_url, "-o", "package.zip")
	err = ttf2pt1_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ttf2pt1_bin_url := "https://downloads.sourceforge.net/project/ttf2pt1/ttf2pt1/3.4.4/ttf2pt1-3.4.4.tgz"
	ttf2pt1_cmd_bin := exec.Command("curl", "-L", ttf2pt1_bin_url, "-o", "binary.bin")
	err = ttf2pt1_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ttf2pt1_src_url := "https://downloads.sourceforge.net/project/ttf2pt1/ttf2pt1/3.4.4/ttf2pt1-3.4.4.tgz"
	ttf2pt1_cmd_src := exec.Command("curl", "-L", ttf2pt1_src_url, "-o", "source.tar.gz")
	err = ttf2pt1_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ttf2pt1_cmd_direct := exec.Command("./binary")
	err = ttf2pt1_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
