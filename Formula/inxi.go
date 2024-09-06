package main

// Inxi - Full featured CLI system information tool
// Homepage: https://smxi.org/docs/inxi.htm

import (
	"fmt"
	
	"os/exec"
)

func installInxi() {
	// Método 1: Descargar y extraer .tar.gz
	inxi_tar_url := "https://codeberg.org/smxi/inxi/archive/3.3.36-1.tar.gz"
	inxi_cmd_tar := exec.Command("curl", "-L", inxi_tar_url, "-o", "package.tar.gz")
	err := inxi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inxi_zip_url := "https://codeberg.org/smxi/inxi/archive/3.3.36-1.zip"
	inxi_cmd_zip := exec.Command("curl", "-L", inxi_zip_url, "-o", "package.zip")
	err = inxi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inxi_bin_url := "https://codeberg.org/smxi/inxi/archive/3.3.36-1.bin"
	inxi_cmd_bin := exec.Command("curl", "-L", inxi_bin_url, "-o", "binary.bin")
	err = inxi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inxi_src_url := "https://codeberg.org/smxi/inxi/archive/3.3.36-1.src.tar.gz"
	inxi_cmd_src := exec.Command("curl", "-L", inxi_src_url, "-o", "source.tar.gz")
	err = inxi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inxi_cmd_direct := exec.Command("./binary")
	err = inxi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
