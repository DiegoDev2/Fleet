package main

// Vile - Vi Like Emacs Editor
// Homepage: https://invisible-island.net/vile/

import (
	"fmt"
	
	"os/exec"
)

func installVile() {
	// Método 1: Descargar y extraer .tar.gz
	vile_tar_url := "https://invisible-island.net/archives/vile/current/vile-9.8z.tgz"
	vile_cmd_tar := exec.Command("curl", "-L", vile_tar_url, "-o", "package.tar.gz")
	err := vile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vile_zip_url := "https://invisible-island.net/archives/vile/current/vile-9.8z.tgz"
	vile_cmd_zip := exec.Command("curl", "-L", vile_zip_url, "-o", "package.zip")
	err = vile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vile_bin_url := "https://invisible-island.net/archives/vile/current/vile-9.8z.tgz"
	vile_cmd_bin := exec.Command("curl", "-L", vile_bin_url, "-o", "binary.bin")
	err = vile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vile_src_url := "https://invisible-island.net/archives/vile/current/vile-9.8z.tgz"
	vile_cmd_src := exec.Command("curl", "-L", vile_src_url, "-o", "source.tar.gz")
	err = vile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vile_cmd_direct := exec.Command("./binary")
	err = vile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
