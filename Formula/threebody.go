package main

// ThreeBody - 三体编程语言 Three Body Language written in Rust
// Homepage: https://github.com/rustq/3body-lang

import (
	"fmt"
	
	"os/exec"
)

func installThreeBody() {
	// Método 1: Descargar y extraer .tar.gz
	threebody_tar_url := "https://github.com/rustq/3body-lang/archive/refs/tags/0.6.1.tar.gz"
	threebody_cmd_tar := exec.Command("curl", "-L", threebody_tar_url, "-o", "package.tar.gz")
	err := threebody_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	threebody_zip_url := "https://github.com/rustq/3body-lang/archive/refs/tags/0.6.1.zip"
	threebody_cmd_zip := exec.Command("curl", "-L", threebody_zip_url, "-o", "package.zip")
	err = threebody_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	threebody_bin_url := "https://github.com/rustq/3body-lang/archive/refs/tags/0.6.1.bin"
	threebody_cmd_bin := exec.Command("curl", "-L", threebody_bin_url, "-o", "binary.bin")
	err = threebody_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	threebody_src_url := "https://github.com/rustq/3body-lang/archive/refs/tags/0.6.1.src.tar.gz"
	threebody_cmd_src := exec.Command("curl", "-L", threebody_src_url, "-o", "source.tar.gz")
	err = threebody_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	threebody_cmd_direct := exec.Command("./binary")
	err = threebody_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
