package main

// Gptsync - GPT and MBR partition tables synchronization tool
// Homepage: https://refit.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGptsync() {
	// Método 1: Descargar y extraer .tar.gz
	gptsync_tar_url := "https://downloads.sourceforge.net/project/refit/rEFIt/0.14/refit-src-0.14.tar.gz"
	gptsync_cmd_tar := exec.Command("curl", "-L", gptsync_tar_url, "-o", "package.tar.gz")
	err := gptsync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gptsync_zip_url := "https://downloads.sourceforge.net/project/refit/rEFIt/0.14/refit-src-0.14.zip"
	gptsync_cmd_zip := exec.Command("curl", "-L", gptsync_zip_url, "-o", "package.zip")
	err = gptsync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gptsync_bin_url := "https://downloads.sourceforge.net/project/refit/rEFIt/0.14/refit-src-0.14.bin"
	gptsync_cmd_bin := exec.Command("curl", "-L", gptsync_bin_url, "-o", "binary.bin")
	err = gptsync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gptsync_src_url := "https://downloads.sourceforge.net/project/refit/rEFIt/0.14/refit-src-0.14.src.tar.gz"
	gptsync_cmd_src := exec.Command("curl", "-L", gptsync_src_url, "-o", "source.tar.gz")
	err = gptsync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gptsync_cmd_direct := exec.Command("./binary")
	err = gptsync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
