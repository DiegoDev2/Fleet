package main

// Nspr - Platform-neutral API for system-level and libc-like functions
// Homepage: https://hg.mozilla.org/projects/nspr

import (
	"fmt"
	
	"os/exec"
)

func installNspr() {
	// Método 1: Descargar y extraer .tar.gz
	nspr_tar_url := "https://archive.mozilla.org/pub/nspr/releases/v4.35/src/nspr-4.35.tar.gz"
	nspr_cmd_tar := exec.Command("curl", "-L", nspr_tar_url, "-o", "package.tar.gz")
	err := nspr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nspr_zip_url := "https://archive.mozilla.org/pub/nspr/releases/v4.35/src/nspr-4.35.zip"
	nspr_cmd_zip := exec.Command("curl", "-L", nspr_zip_url, "-o", "package.zip")
	err = nspr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nspr_bin_url := "https://archive.mozilla.org/pub/nspr/releases/v4.35/src/nspr-4.35.bin"
	nspr_cmd_bin := exec.Command("curl", "-L", nspr_bin_url, "-o", "binary.bin")
	err = nspr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nspr_src_url := "https://archive.mozilla.org/pub/nspr/releases/v4.35/src/nspr-4.35.src.tar.gz"
	nspr_cmd_src := exec.Command("curl", "-L", nspr_src_url, "-o", "source.tar.gz")
	err = nspr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nspr_cmd_direct := exec.Command("./binary")
	err = nspr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
