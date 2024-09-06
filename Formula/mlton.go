package main

// Mlton - Whole-program, optimizing compiler for Standard ML
// Homepage: http://mlton.org

import (
	"fmt"
	
	"os/exec"
)

func installMlton() {
	// Método 1: Descargar y extraer .tar.gz
	mlton_tar_url := "https://downloads.sourceforge.net/project/mlton/mlton/20210117/mlton-20210117.src.tgz"
	mlton_cmd_tar := exec.Command("curl", "-L", mlton_tar_url, "-o", "package.tar.gz")
	err := mlton_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mlton_zip_url := "https://downloads.sourceforge.net/project/mlton/mlton/20210117/mlton-20210117.src.tgz"
	mlton_cmd_zip := exec.Command("curl", "-L", mlton_zip_url, "-o", "package.zip")
	err = mlton_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mlton_bin_url := "https://downloads.sourceforge.net/project/mlton/mlton/20210117/mlton-20210117.src.tgz"
	mlton_cmd_bin := exec.Command("curl", "-L", mlton_bin_url, "-o", "binary.bin")
	err = mlton_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mlton_src_url := "https://downloads.sourceforge.net/project/mlton/mlton/20210117/mlton-20210117.src.tgz"
	mlton_cmd_src := exec.Command("curl", "-L", mlton_src_url, "-o", "source.tar.gz")
	err = mlton_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mlton_cmd_direct := exec.Command("./binary")
	err = mlton_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
}
