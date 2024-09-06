package main

// Miruo - Pretty-print TCP session monitor/analyzer
// Homepage: https://github.com/KLab/miruo/

import (
	"fmt"
	
	"os/exec"
)

func installMiruo() {
	// Método 1: Descargar y extraer .tar.gz
	miruo_tar_url := "https://github.com/KLab/miruo/archive/refs/tags/0.9.6b.tar.gz"
	miruo_cmd_tar := exec.Command("curl", "-L", miruo_tar_url, "-o", "package.tar.gz")
	err := miruo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	miruo_zip_url := "https://github.com/KLab/miruo/archive/refs/tags/0.9.6b.zip"
	miruo_cmd_zip := exec.Command("curl", "-L", miruo_zip_url, "-o", "package.zip")
	err = miruo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	miruo_bin_url := "https://github.com/KLab/miruo/archive/refs/tags/0.9.6b.bin"
	miruo_cmd_bin := exec.Command("curl", "-L", miruo_bin_url, "-o", "binary.bin")
	err = miruo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	miruo_src_url := "https://github.com/KLab/miruo/archive/refs/tags/0.9.6b.src.tar.gz"
	miruo_cmd_src := exec.Command("curl", "-L", miruo_src_url, "-o", "source.tar.gz")
	err = miruo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	miruo_cmd_direct := exec.Command("./binary")
	err = miruo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
