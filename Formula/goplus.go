package main

// Goplus - Programming language for engineering, STEM education, and data science
// Homepage: https://goplus.org

import (
	"fmt"
	
	"os/exec"
)

func installGoplus() {
	// Método 1: Descargar y extraer .tar.gz
	goplus_tar_url := "https://github.com/goplus/gop/archive/refs/tags/v1.2.6.tar.gz"
	goplus_cmd_tar := exec.Command("curl", "-L", goplus_tar_url, "-o", "package.tar.gz")
	err := goplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goplus_zip_url := "https://github.com/goplus/gop/archive/refs/tags/v1.2.6.zip"
	goplus_cmd_zip := exec.Command("curl", "-L", goplus_zip_url, "-o", "package.zip")
	err = goplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goplus_bin_url := "https://github.com/goplus/gop/archive/refs/tags/v1.2.6.bin"
	goplus_cmd_bin := exec.Command("curl", "-L", goplus_bin_url, "-o", "binary.bin")
	err = goplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goplus_src_url := "https://github.com/goplus/gop/archive/refs/tags/v1.2.6.src.tar.gz"
	goplus_cmd_src := exec.Command("curl", "-L", goplus_src_url, "-o", "source.tar.gz")
	err = goplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goplus_cmd_direct := exec.Command("./binary")
	err = goplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
