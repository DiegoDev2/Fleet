package main

// Ksh93 - KornShell, ksh93
// Homepage: https://github.com/ksh93/ksh

import (
	"fmt"
	
	"os/exec"
)

func installKsh93() {
	// Método 1: Descargar y extraer .tar.gz
	ksh93_tar_url := "https://github.com/ksh93/ksh/archive/refs/tags/v1.0.10.tar.gz"
	ksh93_cmd_tar := exec.Command("curl", "-L", ksh93_tar_url, "-o", "package.tar.gz")
	err := ksh93_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ksh93_zip_url := "https://github.com/ksh93/ksh/archive/refs/tags/v1.0.10.zip"
	ksh93_cmd_zip := exec.Command("curl", "-L", ksh93_zip_url, "-o", "package.zip")
	err = ksh93_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ksh93_bin_url := "https://github.com/ksh93/ksh/archive/refs/tags/v1.0.10.bin"
	ksh93_cmd_bin := exec.Command("curl", "-L", ksh93_bin_url, "-o", "binary.bin")
	err = ksh93_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ksh93_src_url := "https://github.com/ksh93/ksh/archive/refs/tags/v1.0.10.src.tar.gz"
	ksh93_cmd_src := exec.Command("curl", "-L", ksh93_src_url, "-o", "source.tar.gz")
	err = ksh93_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ksh93_cmd_direct := exec.Command("./binary")
	err = ksh93_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
