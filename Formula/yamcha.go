package main

// Yamcha - NLP text chunker using Support Vector Machines
// Homepage: http://chasen.org/~taku/software/yamcha/

import (
	"fmt"
	
	"os/exec"
)

func installYamcha() {
	// Método 1: Descargar y extraer .tar.gz
	yamcha_tar_url := "http://chasen.org/~taku/software/yamcha/src/yamcha-0.33.tar.gz"
	yamcha_cmd_tar := exec.Command("curl", "-L", yamcha_tar_url, "-o", "package.tar.gz")
	err := yamcha_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yamcha_zip_url := "http://chasen.org/~taku/software/yamcha/src/yamcha-0.33.zip"
	yamcha_cmd_zip := exec.Command("curl", "-L", yamcha_zip_url, "-o", "package.zip")
	err = yamcha_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yamcha_bin_url := "http://chasen.org/~taku/software/yamcha/src/yamcha-0.33.bin"
	yamcha_cmd_bin := exec.Command("curl", "-L", yamcha_bin_url, "-o", "binary.bin")
	err = yamcha_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yamcha_src_url := "http://chasen.org/~taku/software/yamcha/src/yamcha-0.33.src.tar.gz"
	yamcha_cmd_src := exec.Command("curl", "-L", yamcha_src_url, "-o", "source.tar.gz")
	err = yamcha_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yamcha_cmd_direct := exec.Command("./binary")
	err = yamcha_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tinysvm")
	exec.Command("latte", "install", "tinysvm").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
