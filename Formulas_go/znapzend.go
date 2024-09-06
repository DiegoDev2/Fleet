package main

// Znapzend - ZFS backup with remote capabilities and mbuffer integration
// Homepage: https://www.znapzend.org/

import (
	"fmt"
	
	"os/exec"
)

func installZnapzend() {
	// Método 1: Descargar y extraer .tar.gz
	znapzend_tar_url := "https://github.com/oetiker/znapzend/releases/download/v0.23.2/znapzend-0.23.2.tar.gz"
	znapzend_cmd_tar := exec.Command("curl", "-L", znapzend_tar_url, "-o", "package.tar.gz")
	err := znapzend_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	znapzend_zip_url := "https://github.com/oetiker/znapzend/releases/download/v0.23.2/znapzend-0.23.2.zip"
	znapzend_cmd_zip := exec.Command("curl", "-L", znapzend_zip_url, "-o", "package.zip")
	err = znapzend_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	znapzend_bin_url := "https://github.com/oetiker/znapzend/releases/download/v0.23.2/znapzend-0.23.2.bin"
	znapzend_cmd_bin := exec.Command("curl", "-L", znapzend_bin_url, "-o", "binary.bin")
	err = znapzend_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	znapzend_src_url := "https://github.com/oetiker/znapzend/releases/download/v0.23.2/znapzend-0.23.2.src.tar.gz"
	znapzend_cmd_src := exec.Command("curl", "-L", znapzend_src_url, "-o", "source.tar.gz")
	err = znapzend_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	znapzend_cmd_direct := exec.Command("./binary")
	err = znapzend_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
