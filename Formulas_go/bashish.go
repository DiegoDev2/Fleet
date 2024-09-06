package main

// Bashish - Theme environment for text terminals
// Homepage: https://bashish.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installBashish() {
	// Método 1: Descargar y extraer .tar.gz
	bashish_tar_url := "https://downloads.sourceforge.net/project/bashish/bashish/2.2.4/bashish-2.2.4.tar.gz"
	bashish_cmd_tar := exec.Command("curl", "-L", bashish_tar_url, "-o", "package.tar.gz")
	err := bashish_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bashish_zip_url := "https://downloads.sourceforge.net/project/bashish/bashish/2.2.4/bashish-2.2.4.zip"
	bashish_cmd_zip := exec.Command("curl", "-L", bashish_zip_url, "-o", "package.zip")
	err = bashish_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bashish_bin_url := "https://downloads.sourceforge.net/project/bashish/bashish/2.2.4/bashish-2.2.4.bin"
	bashish_cmd_bin := exec.Command("curl", "-L", bashish_bin_url, "-o", "binary.bin")
	err = bashish_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bashish_src_url := "https://downloads.sourceforge.net/project/bashish/bashish/2.2.4/bashish-2.2.4.src.tar.gz"
	bashish_cmd_src := exec.Command("curl", "-L", bashish_src_url, "-o", "source.tar.gz")
	err = bashish_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bashish_cmd_direct := exec.Command("./binary")
	err = bashish_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dialog")
exec.Command("latte", "install", "dialog")
}
