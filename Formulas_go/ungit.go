package main

// Ungit - Easiest way to use Git. On any platform. Anywhere
// Homepage: https://github.com/FredrikNoren/ungit

import (
	"fmt"
	
	"os/exec"
)

func installUngit() {
	// Método 1: Descargar y extraer .tar.gz
	ungit_tar_url := "https://registry.npmjs.org/ungit/-/ungit-1.5.27.tgz"
	ungit_cmd_tar := exec.Command("curl", "-L", ungit_tar_url, "-o", "package.tar.gz")
	err := ungit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ungit_zip_url := "https://registry.npmjs.org/ungit/-/ungit-1.5.27.tgz"
	ungit_cmd_zip := exec.Command("curl", "-L", ungit_zip_url, "-o", "package.zip")
	err = ungit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ungit_bin_url := "https://registry.npmjs.org/ungit/-/ungit-1.5.27.tgz"
	ungit_cmd_bin := exec.Command("curl", "-L", ungit_bin_url, "-o", "binary.bin")
	err = ungit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ungit_src_url := "https://registry.npmjs.org/ungit/-/ungit-1.5.27.tgz"
	ungit_cmd_src := exec.Command("curl", "-L", ungit_src_url, "-o", "source.tar.gz")
	err = ungit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ungit_cmd_direct := exec.Command("./binary")
	err = ungit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
