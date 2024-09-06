package main

// Smartypants - Typography prettifier
// Homepage: https://daringfireball.net/projects/smartypants/

import (
	"fmt"
	
	"os/exec"
)

func installSmartypants() {
	// Método 1: Descargar y extraer .tar.gz
	smartypants_tar_url := "https://daringfireball.net/projects/downloads/SmartyPants_1.5.1.zip"
	smartypants_cmd_tar := exec.Command("curl", "-L", smartypants_tar_url, "-o", "package.tar.gz")
	err := smartypants_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smartypants_zip_url := "https://daringfireball.net/projects/downloads/SmartyPants_1.5.1.zip"
	smartypants_cmd_zip := exec.Command("curl", "-L", smartypants_zip_url, "-o", "package.zip")
	err = smartypants_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smartypants_bin_url := "https://daringfireball.net/projects/downloads/SmartyPants_1.5.1.zip"
	smartypants_cmd_bin := exec.Command("curl", "-L", smartypants_bin_url, "-o", "binary.bin")
	err = smartypants_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smartypants_src_url := "https://daringfireball.net/projects/downloads/SmartyPants_1.5.1.zip"
	smartypants_cmd_src := exec.Command("curl", "-L", smartypants_src_url, "-o", "source.tar.gz")
	err = smartypants_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smartypants_cmd_direct := exec.Command("./binary")
	err = smartypants_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
