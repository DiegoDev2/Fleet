package main

// Xz - General-purpose data compression with high compression ratio
// Homepage: https://tukaani.org/xz/

import (
	"fmt"
	
	"os/exec"
)

func installXz() {
	// Método 1: Descargar y extraer .tar.gz
	xz_tar_url := "https://github.com/tukaani-project/xz/releases/download/v5.6.2/xz-5.6.2.tar.gz"
	xz_cmd_tar := exec.Command("curl", "-L", xz_tar_url, "-o", "package.tar.gz")
	err := xz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xz_zip_url := "https://github.com/tukaani-project/xz/releases/download/v5.6.2/xz-5.6.2.zip"
	xz_cmd_zip := exec.Command("curl", "-L", xz_zip_url, "-o", "package.zip")
	err = xz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xz_bin_url := "https://github.com/tukaani-project/xz/releases/download/v5.6.2/xz-5.6.2.bin"
	xz_cmd_bin := exec.Command("curl", "-L", xz_bin_url, "-o", "binary.bin")
	err = xz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xz_src_url := "https://github.com/tukaani-project/xz/releases/download/v5.6.2/xz-5.6.2.src.tar.gz"
	xz_cmd_src := exec.Command("curl", "-L", xz_src_url, "-o", "source.tar.gz")
	err = xz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xz_cmd_direct := exec.Command("./binary")
	err = xz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
