package main

// Litani - Metabuild system
// Homepage: https://awslabs.github.io/aws-build-accumulator/

import (
	"fmt"
	
	"os/exec"
)

func installLitani() {
	// Método 1: Descargar y extraer .tar.gz
	litani_tar_url := "https://github.com/awslabs/aws-build-accumulator.git"
	litani_cmd_tar := exec.Command("curl", "-L", litani_tar_url, "-o", "package.tar.gz")
	err := litani_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	litani_zip_url := "https://github.com/awslabs/aws-build-accumulator.git"
	litani_cmd_zip := exec.Command("curl", "-L", litani_zip_url, "-o", "package.zip")
	err = litani_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	litani_bin_url := "https://github.com/awslabs/aws-build-accumulator.git"
	litani_cmd_bin := exec.Command("curl", "-L", litani_bin_url, "-o", "binary.bin")
	err = litani_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	litani_src_url := "https://github.com/awslabs/aws-build-accumulator.git"
	litani_cmd_src := exec.Command("curl", "-L", litani_src_url, "-o", "source.tar.gz")
	err = litani_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	litani_cmd_direct := exec.Command("./binary")
	err = litani_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: mandoc")
exec.Command("latte", "install", "mandoc")
	fmt.Println("Instalando dependencia: scdoc")
exec.Command("latte", "install", "scdoc")
	fmt.Println("Instalando dependencia: gnuplot")
exec.Command("latte", "install", "gnuplot")
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
