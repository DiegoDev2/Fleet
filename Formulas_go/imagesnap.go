package main

// Imagesnap - Tool to capture still images from an iSight or other video source
// Homepage: https://github.com/rharder/imagesnap

import (
	"fmt"
	
	"os/exec"
)

func installImagesnap() {
	// Método 1: Descargar y extraer .tar.gz
	imagesnap_tar_url := "https://github.com/rharder/imagesnap/archive/refs/tags/0.2.16.tar.gz"
	imagesnap_cmd_tar := exec.Command("curl", "-L", imagesnap_tar_url, "-o", "package.tar.gz")
	err := imagesnap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imagesnap_zip_url := "https://github.com/rharder/imagesnap/archive/refs/tags/0.2.16.zip"
	imagesnap_cmd_zip := exec.Command("curl", "-L", imagesnap_zip_url, "-o", "package.zip")
	err = imagesnap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imagesnap_bin_url := "https://github.com/rharder/imagesnap/archive/refs/tags/0.2.16.bin"
	imagesnap_cmd_bin := exec.Command("curl", "-L", imagesnap_bin_url, "-o", "binary.bin")
	err = imagesnap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imagesnap_src_url := "https://github.com/rharder/imagesnap/archive/refs/tags/0.2.16.src.tar.gz"
	imagesnap_cmd_src := exec.Command("curl", "-L", imagesnap_src_url, "-o", "source.tar.gz")
	err = imagesnap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imagesnap_cmd_direct := exec.Command("./binary")
	err = imagesnap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
