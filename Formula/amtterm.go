package main

// Amtterm - Serial-over-LAN (sol) client for Intel AMT
// Homepage: https://www.kraxel.org/blog/linux/amtterm/

import (
	"fmt"
	
	"os/exec"
)

func installAmtterm() {
	// Método 1: Descargar y extraer .tar.gz
	amtterm_tar_url := "https://www.kraxel.org/releases/amtterm/amtterm-1.7.tar.gz"
	amtterm_cmd_tar := exec.Command("curl", "-L", amtterm_tar_url, "-o", "package.tar.gz")
	err := amtterm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	amtterm_zip_url := "https://www.kraxel.org/releases/amtterm/amtterm-1.7.zip"
	amtterm_cmd_zip := exec.Command("curl", "-L", amtterm_zip_url, "-o", "package.zip")
	err = amtterm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	amtterm_bin_url := "https://www.kraxel.org/releases/amtterm/amtterm-1.7.bin"
	amtterm_cmd_bin := exec.Command("curl", "-L", amtterm_bin_url, "-o", "binary.bin")
	err = amtterm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	amtterm_src_url := "https://www.kraxel.org/releases/amtterm/amtterm-1.7.src.tar.gz"
	amtterm_cmd_src := exec.Command("curl", "-L", amtterm_src_url, "-o", "source.tar.gz")
	err = amtterm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	amtterm_cmd_direct := exec.Command("./binary")
	err = amtterm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: vte3")
	exec.Command("latte", "install", "vte3").Run()
}
