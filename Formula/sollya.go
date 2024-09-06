package main

// Sollya - Library for safe floating-point code development
// Homepage: https://www.sollya.org/

import (
	"fmt"
	
	"os/exec"
)

func installSollya() {
	// Método 1: Descargar y extraer .tar.gz
	sollya_tar_url := "https://www.sollya.org/releases/sollya-8.0/sollya-8.0.tar.gz"
	sollya_cmd_tar := exec.Command("curl", "-L", sollya_tar_url, "-o", "package.tar.gz")
	err := sollya_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sollya_zip_url := "https://www.sollya.org/releases/sollya-8.0/sollya-8.0.zip"
	sollya_cmd_zip := exec.Command("curl", "-L", sollya_zip_url, "-o", "package.zip")
	err = sollya_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sollya_bin_url := "https://www.sollya.org/releases/sollya-8.0/sollya-8.0.bin"
	sollya_cmd_bin := exec.Command("curl", "-L", sollya_bin_url, "-o", "binary.bin")
	err = sollya_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sollya_src_url := "https://www.sollya.org/releases/sollya-8.0/sollya-8.0.src.tar.gz"
	sollya_cmd_src := exec.Command("curl", "-L", sollya_src_url, "-o", "source.tar.gz")
	err = sollya_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sollya_cmd_direct := exec.Command("./binary")
	err = sollya_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fplll")
	exec.Command("latte", "install", "fplll").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: mpfi")
	exec.Command("latte", "install", "mpfi").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
}
