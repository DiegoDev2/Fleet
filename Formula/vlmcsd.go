package main

// Vlmcsd - KMS Emulator in C
// Homepage: https://github.com/Wind4/vlmcsd

import (
	"fmt"
	
	"os/exec"
)

func installVlmcsd() {
	// Método 1: Descargar y extraer .tar.gz
	vlmcsd_tar_url := "https://github.com/Wind4/vlmcsd/archive/refs/tags/svn1113.tar.gz"
	vlmcsd_cmd_tar := exec.Command("curl", "-L", vlmcsd_tar_url, "-o", "package.tar.gz")
	err := vlmcsd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vlmcsd_zip_url := "https://github.com/Wind4/vlmcsd/archive/refs/tags/svn1113.zip"
	vlmcsd_cmd_zip := exec.Command("curl", "-L", vlmcsd_zip_url, "-o", "package.zip")
	err = vlmcsd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vlmcsd_bin_url := "https://github.com/Wind4/vlmcsd/archive/refs/tags/svn1113.bin"
	vlmcsd_cmd_bin := exec.Command("curl", "-L", vlmcsd_bin_url, "-o", "binary.bin")
	err = vlmcsd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vlmcsd_src_url := "https://github.com/Wind4/vlmcsd/archive/refs/tags/svn1113.src.tar.gz"
	vlmcsd_cmd_src := exec.Command("curl", "-L", vlmcsd_src_url, "-o", "source.tar.gz")
	err = vlmcsd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vlmcsd_cmd_direct := exec.Command("./binary")
	err = vlmcsd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
