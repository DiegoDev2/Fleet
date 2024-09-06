package main

// Aspcud - Package dependency solver
// Homepage: https://potassco.org/aspcud/

import (
	"fmt"
	
	"os/exec"
)

func installAspcud() {
	// Método 1: Descargar y extraer .tar.gz
	aspcud_tar_url := "https://github.com/potassco/aspcud/archive/refs/tags/v1.9.6.tar.gz"
	aspcud_cmd_tar := exec.Command("curl", "-L", aspcud_tar_url, "-o", "package.tar.gz")
	err := aspcud_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aspcud_zip_url := "https://github.com/potassco/aspcud/archive/refs/tags/v1.9.6.zip"
	aspcud_cmd_zip := exec.Command("curl", "-L", aspcud_zip_url, "-o", "package.zip")
	err = aspcud_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aspcud_bin_url := "https://github.com/potassco/aspcud/archive/refs/tags/v1.9.6.bin"
	aspcud_cmd_bin := exec.Command("curl", "-L", aspcud_bin_url, "-o", "binary.bin")
	err = aspcud_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aspcud_src_url := "https://github.com/potassco/aspcud/archive/refs/tags/v1.9.6.src.tar.gz"
	aspcud_cmd_src := exec.Command("curl", "-L", aspcud_src_url, "-o", "source.tar.gz")
	err = aspcud_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aspcud_cmd_direct := exec.Command("./binary")
	err = aspcud_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: re2c")
	exec.Command("latte", "install", "re2c").Run()
	fmt.Println("Instalando dependencia: clingo")
	exec.Command("latte", "install", "clingo").Run()
}
