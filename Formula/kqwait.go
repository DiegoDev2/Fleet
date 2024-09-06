package main

// Kqwait - Wait for events on files or directories on macOS
// Homepage: https://github.com/sschober/kqwait

import (
	"fmt"
	
	"os/exec"
)

func installKqwait() {
	// Método 1: Descargar y extraer .tar.gz
	kqwait_tar_url := "https://github.com/sschober/kqwait/archive/refs/tags/kqwait-v1.0.3.tar.gz"
	kqwait_cmd_tar := exec.Command("curl", "-L", kqwait_tar_url, "-o", "package.tar.gz")
	err := kqwait_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kqwait_zip_url := "https://github.com/sschober/kqwait/archive/refs/tags/kqwait-v1.0.3.zip"
	kqwait_cmd_zip := exec.Command("curl", "-L", kqwait_zip_url, "-o", "package.zip")
	err = kqwait_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kqwait_bin_url := "https://github.com/sschober/kqwait/archive/refs/tags/kqwait-v1.0.3.bin"
	kqwait_cmd_bin := exec.Command("curl", "-L", kqwait_bin_url, "-o", "binary.bin")
	err = kqwait_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kqwait_src_url := "https://github.com/sschober/kqwait/archive/refs/tags/kqwait-v1.0.3.src.tar.gz"
	kqwait_cmd_src := exec.Command("curl", "-L", kqwait_src_url, "-o", "source.tar.gz")
	err = kqwait_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kqwait_cmd_direct := exec.Command("./binary")
	err = kqwait_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
