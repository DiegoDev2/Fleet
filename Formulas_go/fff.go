package main

// Fff - Simple file manager written in bash
// Homepage: https://github.com/dylanaraps/fff

import (
	"fmt"
	
	"os/exec"
)

func installFff() {
	// Método 1: Descargar y extraer .tar.gz
	fff_tar_url := "https://github.com/dylanaraps/fff/archive/refs/tags/2.2.tar.gz"
	fff_cmd_tar := exec.Command("curl", "-L", fff_tar_url, "-o", "package.tar.gz")
	err := fff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fff_zip_url := "https://github.com/dylanaraps/fff/archive/refs/tags/2.2.zip"
	fff_cmd_zip := exec.Command("curl", "-L", fff_zip_url, "-o", "package.zip")
	err = fff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fff_bin_url := "https://github.com/dylanaraps/fff/archive/refs/tags/2.2.bin"
	fff_cmd_bin := exec.Command("curl", "-L", fff_bin_url, "-o", "binary.bin")
	err = fff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fff_src_url := "https://github.com/dylanaraps/fff/archive/refs/tags/2.2.src.tar.gz"
	fff_cmd_src := exec.Command("curl", "-L", fff_src_url, "-o", "source.tar.gz")
	err = fff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fff_cmd_direct := exec.Command("./binary")
	err = fff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
