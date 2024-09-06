package main

// G2 - Friendly git client
// Homepage: https://orefalo.github.io/g2/

import (
	"fmt"
	
	"os/exec"
)

func installG2() {
	// Método 1: Descargar y extraer .tar.gz
	g2_tar_url := "https://github.com/orefalo/g2/archive/refs/tags/v1.1.tar.gz"
	g2_cmd_tar := exec.Command("curl", "-L", g2_tar_url, "-o", "package.tar.gz")
	err := g2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	g2_zip_url := "https://github.com/orefalo/g2/archive/refs/tags/v1.1.zip"
	g2_cmd_zip := exec.Command("curl", "-L", g2_zip_url, "-o", "package.zip")
	err = g2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	g2_bin_url := "https://github.com/orefalo/g2/archive/refs/tags/v1.1.bin"
	g2_cmd_bin := exec.Command("curl", "-L", g2_bin_url, "-o", "binary.bin")
	err = g2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	g2_src_url := "https://github.com/orefalo/g2/archive/refs/tags/v1.1.src.tar.gz"
	g2_cmd_src := exec.Command("curl", "-L", g2_src_url, "-o", "source.tar.gz")
	err = g2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	g2_cmd_direct := exec.Command("./binary")
	err = g2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
