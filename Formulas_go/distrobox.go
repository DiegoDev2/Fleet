package main

// Distrobox - Use any Linux distribution inside your terminal
// Homepage: https://distrobox.privatedns.org/

import (
	"fmt"
	
	"os/exec"
)

func installDistrobox() {
	// Método 1: Descargar y extraer .tar.gz
	distrobox_tar_url := "https://github.com/89luca89/distrobox/archive/refs/tags/1.7.2.1.tar.gz"
	distrobox_cmd_tar := exec.Command("curl", "-L", distrobox_tar_url, "-o", "package.tar.gz")
	err := distrobox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	distrobox_zip_url := "https://github.com/89luca89/distrobox/archive/refs/tags/1.7.2.1.zip"
	distrobox_cmd_zip := exec.Command("curl", "-L", distrobox_zip_url, "-o", "package.zip")
	err = distrobox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	distrobox_bin_url := "https://github.com/89luca89/distrobox/archive/refs/tags/1.7.2.1.bin"
	distrobox_cmd_bin := exec.Command("curl", "-L", distrobox_bin_url, "-o", "binary.bin")
	err = distrobox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	distrobox_src_url := "https://github.com/89luca89/distrobox/archive/refs/tags/1.7.2.1.src.tar.gz"
	distrobox_cmd_src := exec.Command("curl", "-L", distrobox_src_url, "-o", "source.tar.gz")
	err = distrobox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	distrobox_cmd_direct := exec.Command("./binary")
	err = distrobox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
