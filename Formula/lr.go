package main

// Lr - File list utility with features from ls(1), find(1), stat(1), and du(1)
// Homepage: https://github.com/leahneukirchen/lr

import (
	"fmt"
	
	"os/exec"
)

func installLr() {
	// Método 1: Descargar y extraer .tar.gz
	lr_tar_url := "https://github.com/leahneukirchen/lr/archive/refs/tags/v1.6.tar.gz"
	lr_cmd_tar := exec.Command("curl", "-L", lr_tar_url, "-o", "package.tar.gz")
	err := lr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lr_zip_url := "https://github.com/leahneukirchen/lr/archive/refs/tags/v1.6.zip"
	lr_cmd_zip := exec.Command("curl", "-L", lr_zip_url, "-o", "package.zip")
	err = lr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lr_bin_url := "https://github.com/leahneukirchen/lr/archive/refs/tags/v1.6.bin"
	lr_cmd_bin := exec.Command("curl", "-L", lr_bin_url, "-o", "binary.bin")
	err = lr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lr_src_url := "https://github.com/leahneukirchen/lr/archive/refs/tags/v1.6.src.tar.gz"
	lr_cmd_src := exec.Command("curl", "-L", lr_src_url, "-o", "source.tar.gz")
	err = lr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lr_cmd_direct := exec.Command("./binary")
	err = lr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
