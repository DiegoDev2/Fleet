package main

// Ghr - Upload multiple artifacts to GitHub Release in parallel
// Homepage: https://tcnksm.github.io/ghr

import (
	"fmt"
	
	"os/exec"
)

func installGhr() {
	// Método 1: Descargar y extraer .tar.gz
	ghr_tar_url := "https://github.com/tcnksm/ghr/archive/refs/tags/v0.16.2.tar.gz"
	ghr_cmd_tar := exec.Command("curl", "-L", ghr_tar_url, "-o", "package.tar.gz")
	err := ghr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghr_zip_url := "https://github.com/tcnksm/ghr/archive/refs/tags/v0.16.2.zip"
	ghr_cmd_zip := exec.Command("curl", "-L", ghr_zip_url, "-o", "package.zip")
	err = ghr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghr_bin_url := "https://github.com/tcnksm/ghr/archive/refs/tags/v0.16.2.bin"
	ghr_cmd_bin := exec.Command("curl", "-L", ghr_bin_url, "-o", "binary.bin")
	err = ghr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghr_src_url := "https://github.com/tcnksm/ghr/archive/refs/tags/v0.16.2.src.tar.gz"
	ghr_cmd_src := exec.Command("curl", "-L", ghr_src_url, "-o", "source.tar.gz")
	err = ghr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghr_cmd_direct := exec.Command("./binary")
	err = ghr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
