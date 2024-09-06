package main

// Piknik - Copy/paste anything over the network
// Homepage: https://github.com/jedisct1/piknik

import (
	"fmt"
	
	"os/exec"
)

func installPiknik() {
	// Método 1: Descargar y extraer .tar.gz
	piknik_tar_url := "https://github.com/jedisct1/piknik/archive/refs/tags/0.10.1.tar.gz"
	piknik_cmd_tar := exec.Command("curl", "-L", piknik_tar_url, "-o", "package.tar.gz")
	err := piknik_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	piknik_zip_url := "https://github.com/jedisct1/piknik/archive/refs/tags/0.10.1.zip"
	piknik_cmd_zip := exec.Command("curl", "-L", piknik_zip_url, "-o", "package.zip")
	err = piknik_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	piknik_bin_url := "https://github.com/jedisct1/piknik/archive/refs/tags/0.10.1.bin"
	piknik_cmd_bin := exec.Command("curl", "-L", piknik_bin_url, "-o", "binary.bin")
	err = piknik_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	piknik_src_url := "https://github.com/jedisct1/piknik/archive/refs/tags/0.10.1.src.tar.gz"
	piknik_cmd_src := exec.Command("curl", "-L", piknik_src_url, "-o", "source.tar.gz")
	err = piknik_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	piknik_cmd_direct := exec.Command("./binary")
	err = piknik_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
