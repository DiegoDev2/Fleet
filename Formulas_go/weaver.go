package main

// Weaver - Command-line tool for Weaver
// Homepage: https://github.com/scribd/Weaver

import (
	"fmt"
	
	"os/exec"
)

func installWeaver() {
	// Método 1: Descargar y extraer .tar.gz
	weaver_tar_url := "https://github.com/scribd/Weaver/archive/refs/tags/1.1.5.tar.gz"
	weaver_cmd_tar := exec.Command("curl", "-L", weaver_tar_url, "-o", "package.tar.gz")
	err := weaver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	weaver_zip_url := "https://github.com/scribd/Weaver/archive/refs/tags/1.1.5.zip"
	weaver_cmd_zip := exec.Command("curl", "-L", weaver_zip_url, "-o", "package.zip")
	err = weaver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	weaver_bin_url := "https://github.com/scribd/Weaver/archive/refs/tags/1.1.5.bin"
	weaver_cmd_bin := exec.Command("curl", "-L", weaver_bin_url, "-o", "binary.bin")
	err = weaver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	weaver_src_url := "https://github.com/scribd/Weaver/archive/refs/tags/1.1.5.src.tar.gz"
	weaver_cmd_src := exec.Command("curl", "-L", weaver_src_url, "-o", "source.tar.gz")
	err = weaver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	weaver_cmd_direct := exec.Command("./binary")
	err = weaver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
