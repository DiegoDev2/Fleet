package main

// OneMl - Reboot of ML, unifying its core and (now first-class) module layers
// Homepage: https://people.mpi-sws.org/~rossberg/1ml/

import (
	"fmt"
	
	"os/exec"
)

func installOneMl() {
	// Método 1: Descargar y extraer .tar.gz
	oneml_tar_url := "https://people.mpi-sws.org/~rossberg/1ml/1ml-0.1.zip"
	oneml_cmd_tar := exec.Command("curl", "-L", oneml_tar_url, "-o", "package.tar.gz")
	err := oneml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oneml_zip_url := "https://people.mpi-sws.org/~rossberg/1ml/1ml-0.1.zip"
	oneml_cmd_zip := exec.Command("curl", "-L", oneml_zip_url, "-o", "package.zip")
	err = oneml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oneml_bin_url := "https://people.mpi-sws.org/~rossberg/1ml/1ml-0.1.zip"
	oneml_cmd_bin := exec.Command("curl", "-L", oneml_bin_url, "-o", "binary.bin")
	err = oneml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oneml_src_url := "https://people.mpi-sws.org/~rossberg/1ml/1ml-0.1.zip"
	oneml_cmd_src := exec.Command("curl", "-L", oneml_src_url, "-o", "source.tar.gz")
	err = oneml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oneml_cmd_direct := exec.Command("./binary")
	err = oneml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml")
	exec.Command("latte", "install", "ocaml").Run()
}
