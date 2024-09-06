package main

// Apgdiff - Another PostgreSQL diff tool
// Homepage: https://www.apgdiff.com/

import (
	"fmt"
	
	"os/exec"
)

func installApgdiff() {
	// Método 1: Descargar y extraer .tar.gz
	apgdiff_tar_url := "https://github.com/fordfrog/apgdiff/archive/refs/tags/release_2.7.0.tar.gz"
	apgdiff_cmd_tar := exec.Command("curl", "-L", apgdiff_tar_url, "-o", "package.tar.gz")
	err := apgdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apgdiff_zip_url := "https://github.com/fordfrog/apgdiff/archive/refs/tags/release_2.7.0.zip"
	apgdiff_cmd_zip := exec.Command("curl", "-L", apgdiff_zip_url, "-o", "package.zip")
	err = apgdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apgdiff_bin_url := "https://github.com/fordfrog/apgdiff/archive/refs/tags/release_2.7.0.bin"
	apgdiff_cmd_bin := exec.Command("curl", "-L", apgdiff_bin_url, "-o", "binary.bin")
	err = apgdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apgdiff_src_url := "https://github.com/fordfrog/apgdiff/archive/refs/tags/release_2.7.0.src.tar.gz"
	apgdiff_cmd_src := exec.Command("curl", "-L", apgdiff_src_url, "-o", "source.tar.gz")
	err = apgdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apgdiff_cmd_direct := exec.Command("./binary")
	err = apgdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
exec.Command("latte", "install", "ant")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
