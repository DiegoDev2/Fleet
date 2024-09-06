package main

// Re2c - Generate C-based recognizers from regular expressions
// Homepage: https://re2c.org

import (
	"fmt"
	
	"os/exec"
)

func installRe2c() {
	// Método 1: Descargar y extraer .tar.gz
	re2c_tar_url := "https://github.com/skvadrik/re2c/releases/download/3.1/re2c-3.1.tar.xz"
	re2c_cmd_tar := exec.Command("curl", "-L", re2c_tar_url, "-o", "package.tar.gz")
	err := re2c_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	re2c_zip_url := "https://github.com/skvadrik/re2c/releases/download/3.1/re2c-3.1.tar.xz"
	re2c_cmd_zip := exec.Command("curl", "-L", re2c_zip_url, "-o", "package.zip")
	err = re2c_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	re2c_bin_url := "https://github.com/skvadrik/re2c/releases/download/3.1/re2c-3.1.tar.xz"
	re2c_cmd_bin := exec.Command("curl", "-L", re2c_bin_url, "-o", "binary.bin")
	err = re2c_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	re2c_src_url := "https://github.com/skvadrik/re2c/releases/download/3.1/re2c-3.1.tar.xz"
	re2c_cmd_src := exec.Command("curl", "-L", re2c_src_url, "-o", "source.tar.gz")
	err = re2c_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	re2c_cmd_direct := exec.Command("./binary")
	err = re2c_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
