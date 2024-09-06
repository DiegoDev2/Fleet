package main

// Qsoas - Versatile software for data analysis
// Homepage: https://bip.cnrs.fr/groups/bip06/software/

import (
	"fmt"
	
	"os/exec"
)

func installQsoas() {
	// Método 1: Descargar y extraer .tar.gz
	qsoas_tar_url := "https://bip.cnrs.fr/wp-content/uploads/qsoas/qsoas-3.3.tar.gz"
	qsoas_cmd_tar := exec.Command("curl", "-L", qsoas_tar_url, "-o", "package.tar.gz")
	err := qsoas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qsoas_zip_url := "https://bip.cnrs.fr/wp-content/uploads/qsoas/qsoas-3.3.zip"
	qsoas_cmd_zip := exec.Command("curl", "-L", qsoas_zip_url, "-o", "package.zip")
	err = qsoas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qsoas_bin_url := "https://bip.cnrs.fr/wp-content/uploads/qsoas/qsoas-3.3.bin"
	qsoas_cmd_bin := exec.Command("curl", "-L", qsoas_bin_url, "-o", "binary.bin")
	err = qsoas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qsoas_src_url := "https://bip.cnrs.fr/wp-content/uploads/qsoas/qsoas-3.3.src.tar.gz"
	qsoas_cmd_src := exec.Command("curl", "-L", qsoas_src_url, "-o", "source.tar.gz")
	err = qsoas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qsoas_cmd_direct := exec.Command("./binary")
	err = qsoas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: gsl")
exec.Command("latte", "install", "gsl")
	fmt.Println("Instalando dependencia: mruby")
exec.Command("latte", "install", "mruby")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
}
